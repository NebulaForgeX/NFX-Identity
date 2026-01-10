# RabbitMQ 消息传递包 / RabbitMQ Messaging Package

`rabbitmqx` 是一个基于 RabbitMQ 的类型安全消息传递包，提供了完整的 RabbitMQ 功能支持，包括 Exchange、Queue、RoutingKey、BindingKey 等原生概念。

## 快速开始 / Quick Start

### 安装

```bash
go get nfxid/pkgs/rabbitmqx
```

### 基本用法

基本用法请参考：[USAGE_BASIC.md](./docs/USAGE_BASIC.md)

### 功能文档

- **基本用法**：[USAGE_BASIC.md](./docs/USAGE_BASIC.md) - 初始化、创建 Publisher/Subscriber、发布和订阅消息
- **广播发送**：[USAGE_BROADCAST.md](./docs/USAGE_BROADCAST.md) - 使用 Fanout Exchange 实现一个 Exchange 发送到多个 Queue
- **消息优先级**：[USAGE_PRIORITY.md](./docs/USAGE_PRIORITY.md) - 如何设置消息优先级
- **配置详解**：[USAGE_CONFIG.md](./docs/USAGE_CONFIG.md) - 完整的配置选项说明
- **Exchange 详解**：[USAGE_EXCHANGE.md](./docs/USAGE_EXCHANGE.md) - Exchange 类型、使用场景和最佳实践
- **高级功能**：[USAGE_ADVANCED.md](./docs/USAGE_ADVANCED.md) - 延迟队列、死信队列、队列限制等

## 核心特性 / Core Features

- ✅ **类型安全**：基于泛型提供编译时类型检查
- ✅ **RabbitMQ 原生**：完全基于 Exchange、Queue、RoutingKey、BindingKey 等原生概念
- ✅ **自动路由**：根据消息类型自动路由到对应的 Exchange/Queue
- ✅ **消息优先级**：支持消息优先级（0-255）
- ✅ **广播支持**：支持 Fanout Exchange 实现一对多广播
- ✅ **延迟队列**：支持消息 TTL 和延迟队列
- ✅ **死信队列**：支持死信队列处理失败消息
- ✅ **队列限制**：支持队列长度和大小限制
- ✅ **自动重连**：支持自动重连和故障恢复

## 与 Kafka 的区别 / Differences from Kafka

RabbitMQ 和 Kafka 的核心区别请参考：[USAGE_EXCHANGE.md](./docs/USAGE_EXCHANGE.md#rabbitmq-vs-kafka核心区别--core-differences)

## 示例代码 / Example Code

### 发布消息

```go
import (
    "context"
    "nfxid/pkgs/rabbitmqx"
    "nfxid/pkgs/rabbitmqx/messaging"
)

// 创建 Publisher
publisher, err := rabbitmqx.NewPublisher(cfg)
if err != nil {
    log.Fatal(err)
}

// 发布消息
msg := MyMessage{ID: "123"}
err = messaging.PublishMessage(ctx, publisher, msg)
```

### 订阅消息

```go
// 创建 Subscriber
subscriber, err := rabbitmqx.NewSubscriber(cfg)
if err != nil {
    log.Fatal(err)
}

// 创建 Router
router, err := messaging.NewMessageRouter(subscriber, messaging.MessageRouterConfig{
    CloseTimeout: 10 * time.Second,
})

// 注册处理器
messaging.RegisterHandler(router, func(ctx context.Context, msg MyMessage, rawMsg *message.Message) error {
    log.Printf("Received: %+v", msg)
    return nil
})

// 启动路由器
router.Run(ctx)
```

## 文档导航 / Documentation Navigation

- [基本用法](./docs/USAGE_BASIC.md) - 快速上手指南
- [广播发送](./docs/USAGE_BROADCAST.md) - 实现一对多消息广播
- [消息优先级](./docs/USAGE_PRIORITY.md) - 设置消息优先级
- [配置详解](./docs/USAGE_CONFIG.md) - 完整配置选项
- [Exchange 详解](./docs/USAGE_EXCHANGE.md) - Exchange 类型和使用场景
- [高级功能](./docs/USAGE_ADVANCED.md) - 延迟队列、死信队列等

## 许可证 / License

[Your License Here]
