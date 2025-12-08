//go:build event_bus

// Package message
package message

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
	"half-nothing.cn/service-core/interfaces/bus"
	"half-nothing.cn/service-core/interfaces/logger"
)

// AsyncEventBus 异步事件总线结构体，负责管理事件发布、订阅与处理流程。
//
// 字段说明：
// - logger: 用于记录运行时日志的日志适配器
// - channelSize: 消息通道的缓冲区大小
// - messageCh: 存储待处理事件的消息通道
// - subscribeMap: 记录每种事件类型对应的订阅者列表
// - wg: 控制协程生命周期的等待组
// - lock: 读写锁，保护共享资源（如subscribeMap）的安全访问
// - context: 上下文对象，用于控制事件循环的启停
type AsyncEventBus[T comparable] struct {
	logger       logger.Interface
	channelSize  int
	messageCh    chan *bus.Event[T]
	subscribeMap map[T][]bus.Subscriber[T]
	wg           sync.WaitGroup
	lock         sync.RWMutex
	context      context.Context
	cancel       context.CancelFunc
}

// NewAsyncEventBus 创建一个新的异步事件总线实例
//
// 参数:
//
//	lg: 日志记录器接口，用于记录事件总线的相关日志
//	channelSize: 通道大小，控制事件队列的缓冲区大小
//
// 返回值:
//
//	*AsyncEventBus: 返回初始化后的异步事件总线指针
func NewAsyncEventBus[T comparable](
	lg logger.Interface,
	channelSize int,
) *AsyncEventBus[T] {
	return &AsyncEventBus[T]{
		logger:       logger.NewLoggerAdapter(lg, "event-bus"),
		channelSize:  channelSize,
		messageCh:    nil,
		subscribeMap: make(map[T][]bus.Subscriber[T]),
		wg:           sync.WaitGroup{},
		lock:         sync.RWMutex{},
		context:      nil,
		cancel:       nil,
	}
}

// Start 启动异步事件总线服务
//
// 参数:
//
//	ctx: 父级上下文，用于派生出不受取消影响的新上下文
//
// 功能描述:
//   - 初始化消息通道和上下文环境
//   - 启动后台协程监听并分发事件到各订阅者进行并发处理
func (eventBus *AsyncEventBus[T]) Start(ctx context.Context) {
	eventBus.lock.Lock()
	if eventBus.context != nil && eventBus.cancel != nil {
		eventBus.logger.Error("event bus already running")
		eventBus.lock.Unlock()
		return
	}
	eventBus.context, eventBus.cancel = context.WithCancel(ctx)
	eventBus.logger.Info("event bus started")
	eventBus.messageCh = make(chan *bus.Event[T], eventBus.channelSize)
	eventBus.lock.Unlock()
	eventBus.wg.Add(1)
	go func() {
		defer eventBus.wg.Done()
		for {
			select {
			case event := <-eventBus.messageCh:
				eventBus.wg.Add(1)
				go func() {
					defer eventBus.wg.Done()
					_ = eventBus.handleEvent(event)
				}()
			case <-eventBus.context.Done():
				return
			}
		}
	}()
}

// Stop 停止异步事件总线的服务
//
// 功能描述:
//   - 关闭消息通道
//   - 等待所有正在执行的任务完成
//   - 清理上下文引用
func (eventBus *AsyncEventBus[T]) Stop() {
	eventBus.lock.Lock()
	if eventBus.context == nil || eventBus.cancel == nil {
		eventBus.logger.Error("event bus not running")
		eventBus.lock.Unlock()
		return
	}
	eventBus.lock.Unlock()
	eventBus.logger.Info("event bus stopping")
	eventBus.cancel()
	close(eventBus.messageCh)
	eventBus.wg.Wait()
	eventBus.context = nil
	eventBus.messageCh = nil
	eventBus.logger.Info("event bus stopped")
}

// Publish 发布一个事件至事件总线中
//
// 参数:
//
//	event: 待发布的事件对象
//
// 功能描述:
//   - 将事件发送到内部消息通道以供后续消费
func (eventBus *AsyncEventBus[T]) Publish(event *bus.Event[T]) {
	if eventBus.context == nil || eventBus.cancel == nil || eventBus.messageCh == nil {
		eventBus.logger.Error("event bus not running")
		return
	}
	eventBus.messageCh <- event
}

// Subscribe 注册指定类型的事件订阅者
//
// 参数:
//
//	eventType: 要订阅的事件类型字符串标识
//	handler: 实现了Subscriber接口的具体处理器函数
//
// 功能描述:
//   - 在subscribeMap中将该处理器加入对应事件类型的订阅者列表
func (eventBus *AsyncEventBus[T]) Subscribe(eventType T, handler bus.Subscriber[T]) {
	eventBus.lock.Lock()
	defer eventBus.lock.Unlock()
	if subscribers, exist := eventBus.subscribeMap[eventType]; !exist {
		eventBus.subscribeMap[eventType] = []bus.Subscriber[T]{handler}
	} else {
		eventBus.subscribeMap[eventType] = append(subscribers, handler)
	}
}

// Shutdown 安全关闭事件总线，并设置超时限制
//
// 参数:
//
//	ctx: 外部传入的上下文，用于传递取消信号或截止时间
//
// 返回值:
//
//	error: 若在设定时间内未能成功停止则返回错误；否则返回nil
func (eventBus *AsyncEventBus[T]) Shutdown(ctx context.Context) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	cleanFinish := make(chan struct{})

	go func() {
		eventBus.Stop()
		close(cleanFinish)
	}()

	select {
	case <-timeoutCtx.Done():
		return timeoutCtx.Err()
	case <-cleanFinish:
		return nil
	}
}

// handleEvent 处理单个事件的所有订阅者回调逻辑
//
// 参数:
//
//	event: 需要被处理的事件对象
//
// 返回值:
//
//	error: 所有订阅者的执行结果合并后可能产生的错误信息
//
// 功能描述:
//   - 根据事件类型查找所有已注册的订阅者
//   - 使用errgroup并发调用这些订阅者方法
//   - 收集并返回首个发生的错误（如果有）
func (eventBus *AsyncEventBus[T]) handleEvent(event *bus.Event[T]) error {
	eventBus.lock.RLock()
	subscribers, exist := eventBus.subscribeMap[event.Type]
	eventBus.lock.RUnlock()
	if !exist {
		eventBus.logger.Warnf("No subscribers for message type %+v", event.Type)
		return fmt.Errorf("no subscribers for message type %+v", event.Type)
	}
	var eg errgroup.Group
	for _, subscriber := range subscribers {
		eg.Go(func() error { return subscriber(event) })
	}
	if err := eg.Wait(); err != nil {
		eventBus.logger.Errorf("Error in handling message type %+v: %s", event.Type, err.Error())
		return err
	}
	return nil
}
