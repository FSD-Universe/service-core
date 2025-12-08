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
- permission: 权限管理
- activity: 活动管理
- controller: 管制员管理
- announcement: 公告管理
- audit: 审计管理
- flightplan: 飞行计划管理
- history: 历史管理
- image: 图片管理
- user: 用户管理
- ticket: 工单管理
- event_bus: 事件总线支持

## 开源协议

MIT License

Copyright © 2025 Half_nothing

无附加条款。

[ProjectLanguageCard]: https://img.shields.io/github/languages/top/FSD-Universe/service-core?style=for-the-badge&logo=github

[ProjectLicense]: https://img.shields.io/badge/License-MIT-blue?style=for-the-badge&logo=github

[BuildStateCard]: https://img.shields.io/github/actions/workflow/status/FSD-Universe/service-core/go-test.yml?style=for-the-badge&logo=github&label=UnitTests
