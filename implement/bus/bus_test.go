// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package message
package message

import (
	"context"
	"testing"
	"time"

	"half-nothing.cn/service-core/interfaces/bus"
	"half-nothing.cn/service-core/testutils"
)

func TestAsyncEventBus(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()
	eventBus := NewAsyncEventBus[string](
		testutils.NewFakeLogger(t),
		128,
	)

	eventBus.Subscribe("test", func(message *bus.Event[string]) error {
		t.Log("received message: ", message)
		return nil
	})
	eventBus.Subscribe("test2", func(message *bus.Event[string]) error {
		t.Error("should not receive this message")
		return nil
	})

	eventBus.Start(context.Background())
	eventBus.Publish(&bus.Event[string]{Type: "test"})
	eventBus.Publish(&bus.Event[string]{Type: "test4"})
	time.Sleep(time.Millisecond * 100)
	eventBus.Stop()
}

func BenchmarkAsyncEventBus_Publish(b *testing.B) {
	eventBus := NewAsyncEventBus[string](
		testutils.NewFakeLogger(b),
		128,
	)
	eventBus.Subscribe("test", func(message *bus.Event[string]) error {
		return nil
	})
	eventBus.Start(context.Background())
	for i := 0; i < 3; i++ {
		eventBus.Publish(&bus.Event[string]{Type: "test"})
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		eventBus.Publish(&bus.Event[string]{Type: "test"})
	}

	b.StopTimer()

	eventBus.Stop()
}
