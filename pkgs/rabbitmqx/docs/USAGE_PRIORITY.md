# 消息优先级 / Message Priority

RabbitMQ 支持消息优先级，优先级高的消息会优先被消费。

## 配置优先级队列

首先，需要在队列配置中启用优先级支持：

```toml
[rabbitmq.queue]
    max_priority = 10  # 队列最大优先级（0-255），必须 > 0 才能使用优先级
```

**注意**：
- `max_priority` 必须大于 0 才能使用优先级功能
- 优先级范围：0-255，数值越大优先级越高
- 如果队列未配置优先级（`max_priority = 0`），消息优先级会被忽略

## 发布带优先级的消息

### 方式一：使用 WithPriority 选项（推荐）

```go
import (
    "context"
    "nfxid/pkgs/rabbitmqx/messaging"
)

// 发布高优先级消息
err := messaging.PublishMessage(ctx, publisher, urgentMessage,
    messaging.WithPriority(10),  // 设置优先级为 10（最高）
)

// 发布普通优先级消息
err := messaging.PublishMessage(ctx, publisher, normalMessage,
    messaging.WithPriority(5),  // 设置优先级为 5（中等）
)

// 发布低优先级消息
err := messaging.PublishMessage(ctx, publisher, lowPriorityMessage,
    messaging.WithPriority(1),  // 设置优先级为 1（最低）
)
```

### 方式二：通过元数据设置

```go
err := messaging.PublishMessage(ctx, publisher, msg,
    messaging.WithMetadata(map[string]string{
        "x-priority": "10",  // 优先级（0-255）
    }),
)
```

## 完整示例

### 1. 配置示例

```toml
[rabbitmq]
    uri = "amqp://guest:guest@localhost:5672/"

    [rabbitmq.queue]
        durable = true
        max_priority = 10  # 启用优先级队列，最大优先级为 10

    [rabbitmq.producer_exchanges]
        notification = { exchange = "notifications", routing_key = "notification" }

    [rabbitmq.consumer_queues]
        notification = { queue = "notification-queue", binding_key = "notification" }
```

### 2. 发布不同优先级的消息

```go
package main

import (
    "context"
    "nfxid/pkgs/rabbitmqx"
    "nfxid/pkgs/rabbitmqx/messaging"
)

type NotificationMessage struct {
    Type    string `json:"type"`
    Content string `json:"content"`
}

func (NotificationMessage) RoutingKey() messaging.MessageKey {
    return "notification"
}

func main() {
    cfg := loadConfig()  // 加载配置
    publisher, _ := rabbitmqx.NewPublisher(cfg)
    ctx := context.Background()

    // 发布紧急通知（高优先级）
    urgent := NotificationMessage{
        Type:    "urgent",
        Content: "系统故障，请立即处理！",
    }
    messaging.PublishMessage(ctx, publisher, urgent,
        messaging.WithPriority(10),  // 最高优先级
    )

    // 发布普通通知（中等优先级）
    normal := NotificationMessage{
        Type:    "normal",
        Content: "新功能已上线",
    }
    messaging.PublishMessage(ctx, publisher, normal,
        messaging.WithPriority(5),  // 中等优先级
    )

    // 发布低优先级通知
    low := NotificationMessage{
        Type:    "info",
        Content: "系统维护通知",
    }
    messaging.PublishMessage(ctx, publisher, low,
        messaging.WithPriority(1),  // 低优先级
    )
}
```

### 3. 消费优先级消息

消息会按照优先级从高到低被消费：

```go
// 注册处理器
messaging.RegisterHandler(router, func(ctx context.Context, msg NotificationMessage, rawMsg *message.Message) error {
    log.Printf("Received notification: %s - %s", msg.Type, msg.Content)
    // 高优先级的消息会先被处理
    return nil
})
```

## 优先级工作原理

1. **发布消息时**：通过 `WithPriority` 设置消息优先级
2. **队列存储**：队列按照优先级排序，高优先级消息在前
3. **消费顺序**：消费者优先消费高优先级消息

```
队列状态：
[优先级 10] 消息 A
[优先级 10] 消息 B
[优先级 5]  消息 C
[优先级 1]  消息 D

消费顺序：A → B → C → D
```

## 注意事项

1. **队列必须配置优先级**：如果 `max_priority = 0`，所有消息优先级都会被忽略
2. **优先级范围**：0-255，建议使用较小的范围（如 0-10）以便管理
3. **性能影响**：优先级队列的性能略低于普通队列，因为需要维护排序
4. **优先级继承**：如果消息未设置优先级，使用队列的默认优先级（通常为 0）

## 最佳实践

1. **合理设置优先级范围**：
   ```toml
   max_priority = 10  # 使用 0-10 的范围，而不是 0-255
   ```

2. **定义优先级常量**：
   ```go
   const (
       PriorityUrgent = 10
       PriorityHigh   = 7
       PriorityNormal = 5
       PriorityLow    = 3
       PriorityLowest = 1
   )

   messaging.PublishMessage(ctx, publisher, msg,
       messaging.WithPriority(PriorityUrgent),
   )
   ```

3. **结合业务场景**：
   - 紧急告警：优先级 10
   - 重要通知：优先级 7
   - 普通消息：优先级 5
   - 低优先级消息：优先级 1

## 相关文档

- [基本用法](./USAGE_BASIC.md) - 了解基本发布和订阅
- [配置详解](./USAGE_CONFIG.md) - 查看完整配置选项
