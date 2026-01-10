# 广播发送 / Broadcast Messaging

RabbitMQ 的 **Fanout Exchange** 可以实现一个 Exchange 发送到多个 Queue，实现广播功能。

## 什么是广播？

广播是指一条消息发送到一个 Exchange，然后自动路由到所有绑定到这个 Exchange 的 Queue。

```
Publisher → Fanout Exchange → Queue1 (所有绑定的队列)
                        → Queue2 (都会收到消息)
                        → Queue3
```

## 配置 Fanout Exchange

### 1. 配置 Exchange 类型为 fanout

```toml
[rabbitmq.exchange]
    name = "broadcast-exchange"  # Exchange 名称
    type = "fanout"              # 关键：使用 fanout 类型
    durable = true
    auto_delete = false
```

### 2. 配置多个 Queue 绑定到同一个 Exchange

```toml
# 配置 ProducerExchanges（发布者配置）
[rabbitmq.producer_exchanges]
    notification = { exchange = "broadcast-exchange", routing_key = "" }  # fanout 忽略 routing_key

# 配置 ConsumerQueues（订阅者配置）
# 多个服务可以绑定到同一个 Exchange
[rabbitmq.consumer_queues]
    # 服务 A 的队列
    notification = { queue = "service-a-notification-queue", binding_key = "" }
    
    # 服务 B 的队列（需要单独配置）
    # 注意：这里需要为每个服务创建不同的配置
```

**注意**：由于配置的限制，如果需要多个服务订阅同一个广播 Exchange，每个服务需要：
1. 使用相同的 Exchange 名称
2. 使用不同的 Queue 名称
3. 在各自的配置中设置 `consumer_queues`

## 完整示例

### 1. 配置示例

```toml
[rabbitmq]
    uri = "amqp://guest:guest@localhost:5672/"

    # 配置 Fanout Exchange
    [rabbitmq.exchange]
        name = "system-notifications"
        type = "fanout"  # 关键：fanout 类型
        durable = true

    # 发布者配置
    [rabbitmq.producer_exchanges]
        notification = { exchange = "system-notifications", routing_key = "" }

    # 服务 A 的订阅配置
    [rabbitmq.consumer_queues]
        notification = { queue = "service-a-queue", binding_key = "" }
```

### 2. 发布广播消息

```go
package main

import (
    "context"
    "nfxid/pkgs/rabbitmqx"
    "nfxid/pkgs/rabbitmqx/messaging"
)

type SystemNotificationMessage struct {
    Type    string `json:"type"`
    Message string `json:"message"`
    Time    int64  `json:"time"`
}

func (SystemNotificationMessage) RoutingKey() messaging.MessageKey {
    return "notification"  // 对应配置中的 producer_exchanges key
}

func main() {
    cfg := loadConfig()
    publisher, _ := rabbitmqx.NewPublisher(cfg)
    ctx := context.Background()

    // 发布广播消息
    notification := SystemNotificationMessage{
        Type:    "system",
        Message: "系统维护通知：将在今晚 22:00 进行系统升级",
        Time:    time.Now().Unix(),
    }

    // 发布到 fanout exchange，所有绑定的队列都会收到
    err := messaging.PublishMessage(ctx, publisher, notification)
    if err != nil {
        log.Fatal(err)
    }
}
```

### 3. 多个服务订阅广播

#### 服务 A 的配置

```toml
[rabbitmq]
    uri = "amqp://guest:guest@localhost:5672/"

    [rabbitmq.exchange]
        name = "system-notifications"
        type = "fanout"
        durable = true

    [rabbitmq.consumer_queues]
        notification = { queue = "service-a-notification-queue", binding_key = "" }
```

#### 服务 B 的配置

```toml
[rabbitmq]
    uri = "amqp://guest:guest@localhost:5672/"

    [rabbitmq.exchange]
        name = "system-notifications"  # 相同的 Exchange
        type = "fanout"
        durable = true

    [rabbitmq.consumer_queues]
        notification = { queue = "service-b-notification-queue", binding_key = "" }  # 不同的 Queue
```

#### 服务 C 的配置

```toml
[rabbitmq]
    uri = "amqp://guest:guest@localhost:5672/"

    [rabbitmq.exchange]
        name = "system-notifications"  # 相同的 Exchange
        type = "fanout"
        durable = true

    [rabbitmq.consumer_queues]
        notification = { queue = "service-c-notification-queue", binding_key = "" }  # 不同的 Queue
```

### 4. 订阅广播消息

```go
package service_a

import (
    "context"
    "nfxid/pkgs/rabbitmqx"
    "nfxid/pkgs/rabbitmqx/messaging"
)

func setupBroadcastSubscriber(cfg *rabbitmqx.Config) error {
    // 创建订阅者
    subscriber, err := rabbitmqx.NewSubscriber(cfg)
    if err != nil {
        return err
    }

    // 创建路由器
    router, err := messaging.NewMessageRouter(subscriber, messaging.MessageRouterConfig{
        CloseTimeout: 10 * time.Second,
    })
    if err != nil {
        return err
    }

    // 注册处理器
    messaging.RegisterHandler(router, func(ctx context.Context, msg SystemNotificationMessage, rawMsg *message.Message) error {
        log.Printf("[Service A] Received broadcast: %s - %s", msg.Type, msg.Message)
        // 服务 A 的处理逻辑
        return nil
    })

    // 启动路由器
    go router.Run(context.Background())
    return nil
}
```

## Fanout Exchange 的特点

1. **忽略 RoutingKey**：Fanout Exchange 会忽略 RoutingKey，所有绑定的队列都会收到消息
2. **广播特性**：一条消息会路由到所有绑定的队列
3. **解耦发布者**：发布者不需要知道有多少个订阅者

## 使用场景

### 场景 1：系统通知广播

```go
// 发布系统维护通知
messaging.PublishMessage(ctx, publisher, SystemMaintenanceNotification{
    StartTime: "2024-01-01 22:00:00",
    Duration:  "2 hours",
})

// 所有服务都会收到：
// - 用户服务（user-service-queue）
// - 订单服务（order-service-queue）
// - 支付服务（payment-service-queue）
// - 通知服务（notification-service-queue）
```

### 场景 2：配置更新广播

```go
// 发布配置更新
messaging.PublishMessage(ctx, publisher, ConfigUpdateMessage{
    ConfigKey: "feature_flags",
    NewValue:  "enabled",
})

// 所有服务都会收到配置更新，自动刷新本地配置
```

### 场景 3：缓存失效广播

```go
// 发布缓存失效通知
messaging.PublishMessage(ctx, publisher, CacheInvalidationMessage{
    CacheKey: "user:123",
    Reason:   "user_updated",
})

// 所有服务的缓存都会被清除
```

## 与其他 Exchange 类型的对比

| Exchange 类型 | 路由方式 | 使用场景 |
|--------------|---------|---------|
| **Fanout** | 忽略 RoutingKey，广播到所有队列 | 系统通知、配置更新、缓存失效 |
| **Topic** | 根据 RoutingKey 模式匹配 | 灵活路由、按消息类型路由 |
| **Direct** | 精确匹配 RoutingKey | 点对点通信、精确路由 |
| **Headers** | 根据消息头匹配 | 复杂的路由规则 |

## 最佳实践

1. **使用有意义的 Exchange 名称**：
   ```toml
   name = "system-notifications"  # 而不是 "exchange1"
   ```

2. **每个服务使用独立的 Queue**：
   ```toml
   # 服务 A
   queue = "service-a-notification-queue"
   
   # 服务 B
   queue = "service-b-notification-queue"
   ```

3. **启用持久化**：
   ```toml
   [rabbitmq.exchange]
       durable = true
   
   [rabbitmq.queue]
       durable = true
   ```

4. **处理消息幂等性**：
   由于广播消息会被多个服务接收，确保处理逻辑是幂等的。

## 注意事项

1. **Fanout Exchange 忽略 RoutingKey**：即使设置了 RoutingKey，Fanout Exchange 也会忽略它
2. **所有队列都会收到消息**：确保所有订阅的服务都能正确处理消息
3. **消息重复**：如果同一个服务有多个消费者实例，每个实例都会收到消息（这是正常的）

## 相关文档

- [基本用法](./USAGE_BASIC.md) - 了解基本发布和订阅
- [Exchange 详解](./USAGE_EXCHANGE.md) - 了解所有 Exchange 类型
- [配置详解](./USAGE_CONFIG.md) - 查看完整配置选项
