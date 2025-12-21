# ServiceCore

![ProjectLanguageCard]![ProjectLicense]![BuildStateCard]

## 介绍

ServiceCore 是一个基于 Go 语言的框架，提供了服务框架，用于开发分布式服务。

## 功能

- 日志记录
- 配置管理
- 缓存操作
- Http响应封装

## 使用

```shell
go get half-nothing.cn/service-core
```

## 测试

```shell
go test ./...
```

## 构建标签

- database: 数据库操作
- http: Http服务
- httpjwt: Http服务JWT验证
- grpc: Grpc服务
- permission: 权限工具
- event_bus: 事件总线支持
- telemetry: 可观测性支持

## 命令行参数与环境变量一览

| 命令行参数              | 环境变量               | 描述                 | 默认值                                       |
|:-------------------|:-------------------|:-------------------|:------------------------------------------|
| no_logs            | NO_LOGS            | 禁用日志输出到文件          | false                                     |
| auto_migrate       | AUTO_MIGRATE       | 自动迁移数据库(不要在生产环境使用) | false                                     |
| config             | CONFIG_FILE_PATH   | 配置文件路径             | "config.yaml"                             |
| broadcast_port     | BROADCAST_PORT     | 服务发现广播端口           | 9999                                      |
| heartbeat_interval | HEARTBEAT_INTERVAL | 心跳间隔               | "30s"                                     |
| service_timeout    | SERVICE_TIMEOUT    | 服务超时时间             | "90s"                                     |
| cleanup_interval   | CLEANUP_INTERVAL   | 清理间隔               | "30s"                                     |
| reconnect_timeout  | RECONNECT_TIMEOUT  | 重连超时时间             | "30s"                                     |
| eth_name           | ETH_NAME           | 以太网接口名称            | "Ethernet"(windows) / "eth0"(linux/macos) |
| http_timeout       | HTTP_TIMEOUT       | Http请求超时时间         | "30s"                                     |
| gzip_level         | GZIP_LEVEL         | Gzip压缩等级           | 5                                         |

## 开源协议

MIT License

Copyright © 2025 Half_nothing

无附加条款。

[ProjectLanguageCard]: https://img.shields.io/github/languages/top/FSD-Universe/service-core?style=for-the-badge&logo=github

[ProjectLicense]: https://img.shields.io/badge/License-MIT-blue?style=for-the-badge&logo=github

[BuildStateCard]: https://img.shields.io/github/actions/workflow/status/FSD-Universe/service-core/go-test.yml?style=for-the-badge&logo=github&label=UnitTests
