# Exchange 类型详解 / Exchange Type Guide

## 概述

`rabbitmqx` 提供了类型安全的 `ExchangeType` 来管理 RabbitMQ 的 Exchange 类型，支持基本类型和插件类型。

## ExchangeType 类型

### 类型定义

```go
package messaging

// ExchangeType Exchange 类型
type ExchangeType string
```

### 基本类型（RabbitMQ 内置）

#### ExchangeTypeDirect

**值**：`"direct"`

**特点**：
- 精确匹配 RoutingKey
- 完全匹配才路由

**使用场景**：点对点通信、精确路由

**示例**：
```toml
[rabbitmq.producer_exchanges]
    order_paid = {
        exchange = "order-events",
        routing_key = "order.paid",
        type = "direct"
    }
```

#### ExchangeTypeTopic（默认）

**值**：`"topic"`

**特点**：
- 支持通配符匹配（`*` 和 `#`）
- `*` 匹配一个单词
- `#` 匹配零个或多个单词
- 最灵活的路由方式

**使用场景**：需要灵活路由、按消息类型路由

**示例**：
```toml
[rabbitmq.producer_exchanges]
    user_events = {
        exchange = "user-events",
        routing_key = "user.*",
        type = "topic"
    }
```

#### ExchangeTypeFanout

**值**：`"fanout"`

**特点**：
- 忽略 RoutingKey
- 广播到所有绑定的队列
- 一条消息 → 所有队列

**使用场景**：系统通知、配置更新、缓存失效

**示例**：
```toml
[rabbitmq.producer_exchanges]
    cache_invalidate = {
        exchange = "cache-broadcast",
        routing_key = "",  # Fanout Exchange 忽略 RoutingKey
        type = "fanout"
    }
```

详细说明请参考：[USAGE_BROADCAST.md](./USAGE_BROADCAST.md)

#### ExchangeTypeHeaders

**值**：`"headers"`

**特点**：
- 根据消息头（Headers）匹配
- 忽略 RoutingKey
- 支持 `x-match` 参数（`all` 或 `any`）

**使用场景**：基于消息头进行复杂路由

**示例**：
```toml
[rabbitmq.producer_exchanges]
    header_routed = {
        exchange = "headers-exchange",
        routing_key = "",  # Headers Exchange 忽略 RoutingKey
        type = "headers"
    }
```

### 插件类型（需要安装插件）

#### ExchangeTypeDelayedMessage

**值**：`"x-delayed-message"`

**需要插件**：`rabbitmq-delayed-message-exchange`

**安装**：
```bash
rabbitmq-plugins enable rabbitmq_delayed_message_exchange
```

**特点**：
- 支持延迟消息投递
- 通过 `x-delay` 消息头指定延迟时间（毫秒）
- 底层使用其他 Exchange 类型（默认 topic）

**配置示例**：
```toml
[rabbitmq.producer_exchanges]
    delayed_notification = {
        exchange = "delayed-events",
        routing_key = "notification",
        type = "x-delayed-message"
    }
```

**使用示例**：
```go
// 发布延迟 5 秒的消息
messaging.PublishMessage(ctx, publisher, msg,
    messaging.WithMetadata(map[string]string{
        "x-delay": "5000",  // 延迟 5 秒（毫秒）
    }),
)
```

#### ExchangeTypeConsistentHash

**值**：`"x-consistent-hash"`

**需要插件**：`rabbitmq-consistent-hash-exchange`

**安装**：
```bash
rabbitmq-plugins enable rabbitmq_consistent_hash_exchange
```

**特点**：
- 根据 RoutingKey 的哈希值路由消息
- 实现负载均衡
- 相同 RoutingKey 总是路由到同一个队列

**配置示例**：
```toml
[rabbitmq.producer_exchanges]
    load_balanced = {
        exchange = "hash-exchange",
        routing_key = "task",
        type = "x-consistent-hash"
    }
```

#### ExchangeTypeSharding

**值**：`"x-sharding"`

**需要插件**：`rabbitmq-sharding`

**安装**：
```bash
rabbitmq-plugins enable rabbitmq_sharding
```

**特点**：
- 将消息分片到多个队列
- 提高并发处理能力

**配置示例**：
```toml
[rabbitmq.producer_exchanges]
    sharded = {
        exchange = "sharded-exchange",
        routing_key = "data",
        type = "x-sharding"
    }
```

#### ExchangeTypeModulusHash

**值**：`"x-modulus-hash"`

**需要插件**：`rabbitmq-modulus-hash-exchange`

**特点**：根据 RoutingKey 的模数路由消息

#### ExchangeTypeRandom

**值**：`"x-random"`

**需要插件**：`rabbitmq-random-exchange`

**特点**：随机路由消息到绑定的队列

#### ExchangeTypeJMS

**值**：`"x-jms-topic"`

**需要插件**：`rabbitmq-jms-topic-exchange`

**特点**：支持 JMS 主题语义

## ✅ 为每个消息配置不同的 Exchange 类型

### 核心功能

**✅ 可以在发送消息时动态指定 Exchange 类型，而不是在配置中预先设置！**

### 方式一：在发送消息时动态指定（✅ 推荐）

**在发送消息的那一刻指定 Exchange 类型**，不需要在配置中预先设置：

```go
import (
    "context"
    "nfxid/messages/directory"
    "nfxid/pkgs/rabbitmqx/messaging"
)

// 发送到 Fanout Exchange（广播）
err := messaging.PublishMessage(ctx, publisher, cacheMsg,
    messaging.WithExchangeType(messaging.ExchangeTypeFanout),  // ✅ 在发送时指定
)

// 发送到 Topic Exchange
err := messaging.PublishMessage(ctx, publisher, updateMsg,
    messaging.WithExchangeType(messaging.ExchangeTypeTopic),  // ✅ 在发送时指定
)

// 发送到延迟消息 Exchange（需要插件）
err := messaging.PublishMessage(ctx, publisher, delayedMsg,
    messaging.WithExchangeType(messaging.ExchangeTypeDelayedMessage),  // ✅ 在发送时指定
    messaging.WithMetadata(map[string]string{"x-delay": "5000"}),
)
```

**优点**：
- ✅ **灵活性最高**：每次发送消息时可以指定不同的 Exchange 类型
- ✅ **不需要预先配置**：不需要在配置文件中指定 Exchange 类型
- ✅ **自动声明**：如果 Exchange 不存在，自动创建（使用指定的类型）
- ✅ **类型检查**：如果 Exchange 已存在但类型不匹配，会返回错误

### 方式二：在配置中指定 type（向后兼容）

如果需要在配置中预先指定 Exchange 类型，仍然支持：

```toml
[rabbitmq.exchange]
    name = "nfxid-events"
    type = "topic"  # 默认类型
    durable = true

[rabbitmq.producer_exchanges]
    # 使用 Topic Exchange
    directory = { 
        exchange = "nfxid-events", 
        routing_key = "directory.user.update",
        type = "topic"  # ✅ 可选，指定 Exchange 类型
    }
    
    # 使用 Fanout Exchange（广播）
    user_cache_invalidate = { 
        exchange = "cache-broadcast", 
        routing_key = "",
        type = "fanout"  # ✅ 指定为 Fanout Exchange
    }
```

**注意**：如果同时在配置中指定了 `type` 和在发送消息时使用了 `WithExchangeType`，**发送时的类型会覆盖配置中的类型**。

### 工作原理

#### 方式一：发送时动态指定

1. **发送消息时指定类型**：使用 `WithExchangeType()` 选项
2. **检查 Exchange**：检查 Exchange 是否存在
3. **自动声明**：如果不存在，自动创建（使用指定的类型）
4. **类型验证**：如果已存在但类型不匹配，返回错误

#### 方式二：配置中指定

1. **收集 Exchange 配置**：在创建 Publisher 时，收集所有 `ProducerExchanges` 中指定的 Exchange 名称和类型
2. **预先声明 Exchange**：可选，在创建 Publisher 时预先声明所有需要的 Exchange
3. **类型优先级**：
   - 如果发送时使用 `WithExchangeType`，优先使用发送时指定的类型
   - 如果 `ProducerRouting.Type` 不为空，使用配置中的类型
   - 如果 `ProducerRouting.Type` 为空，使用全局 `ExchangeConfig.Type`
   - 如果全局 `ExchangeConfig.Type` 也为空，使用默认值 `"topic"`

### 自动声明 Exchange

**发送时动态声明**（推荐）：
```go
// 第一次发送消息时，自动声明 Exchange
err := messaging.PublishMessage(ctx, publisher, msg,
    messaging.WithExchangeType(messaging.ExchangeTypeFanout),
)
// ✅ 如果 cache-broadcast Exchange 不存在，自动创建（fanout 类型）
```

**预先声明**（可选）：
```go
publisher, err := rabbitmqx.NewPublisher(cfg)
// ✅ 可选：预先声明所有配置中的 Exchange
// 如果预先声明失败，会在发送消息时自动声明
```

## ExchangeType 方法

### IsValid()

检查 Exchange 类型是否有效：

```go
exchangeType := messaging.ExchangeType("fanout")
if exchangeType.IsValid() {
    // 类型有效
}
```

### IsPluginType()

检查是否为插件类型：

```go
exchangeType := messaging.ExchangeType("x-delayed-message")
if exchangeType.IsPluginType() {
    // 是插件类型，需要安装对应插件
    // 安装：rabbitmq-plugins enable rabbitmq_delayed_message_exchange
}
```

### IsBasicType()

检查是否为基本类型：

```go
exchangeType := messaging.ExchangeType("topic")
if exchangeType.IsBasicType() {
    // 是基本类型，RabbitMQ 内置支持
}
```

## 完整使用示例

### 方式一：发送时动态指定（✅ 推荐）

```go
import (
    "context"
    "nfxid/messages/directory"
    "nfxid/pkgs/rabbitmqx"
    "nfxid/pkgs/rabbitmqx/messaging"
)

// 配置中只需要指定 Exchange 名称和 RoutingKey，不需要指定 type
cfg := &rabbitmqx.Config{
    URI: "amqp://guest:guest@localhost:5672/",
    Exchange: rabbitmqx.ExchangeConfig{
        Durable: true,
    },
    ProducerExchanges: map[messaging.MessageKey]rabbitmqx.ProducerRouting{
        "directory": {
            Exchange:   "nfxid-events",
            RoutingKey: "directory.user.update",
            // ✅ 不需要指定 type，在发送时指定
        },
        "user_cache_invalidate": {
            Exchange:   "cache-broadcast",
            RoutingKey: "",
            // ✅ 不需要指定 type，在发送时指定
        },
    },
}

publisher, _ := rabbitmqx.NewPublisher(cfg)

// 发送消息时动态指定 Exchange 类型
ctx := context.Background()

// 发送到 Topic Exchange
updateMsg := directory.UserUpdateMessage{UserID: "user-123"}
err := messaging.PublishMessage(ctx, publisher, updateMsg,
    messaging.WithExchangeType(messaging.ExchangeTypeTopic),  // ✅ 发送时指定
)

// 发送到 Fanout Exchange（广播）
cacheMsg := directory.UserCacheInvalidateMessage{UserID: "user-123"}
err = messaging.PublishMessage(ctx, publisher, cacheMsg,
    messaging.WithExchangeType(messaging.ExchangeTypeFanout),  // ✅ 发送时指定
)

// 发送到延迟消息 Exchange
delayedMsg := directory.NotificationMessage{Message: "Hello"}
err = messaging.PublishMessage(ctx, publisher, delayedMsg,
    messaging.WithExchangeType(messaging.ExchangeTypeDelayedMessage),  // ✅ 发送时指定
    messaging.WithMetadata(map[string]string{"x-delay": "5000"}),
)
```

### 方式二：配置中指定（向后兼容）

```toml
[rabbitmq]
    uri = "amqp://guest:guest@localhost:5672/"
    client_id = "nfxid-service"
    
    # 全局 Exchange 配置（用于默认 Exchange）
    [rabbitmq.exchange]
        name = "nfxid-events"
        type = "topic"  # 默认类型
        durable = true
        auto_delete = false
    
    # 为不同的消息配置不同的 Exchange 类型
    [rabbitmq.producer_exchanges]
        # Topic Exchange（用于普通消息）
        directory = { 
            exchange = "nfxid-events", 
            routing_key = "directory.user.update",
            type = "topic"  # ✅ 可选，覆盖全局配置
        }
        
        # Fanout Exchange（用于广播缓存清除）
        user_cache_invalidate = { 
            exchange = "cache-broadcast", 
            routing_key = "",
            type = "fanout"  # ✅ 指定为 Fanout Exchange
        }
        
        # Direct Exchange（用于精确路由）
        order_paid = {
            exchange = "order-events",
            routing_key = "order.paid",
            type = "direct"  # ✅ 指定为 Direct Exchange
        }
        
        # 延迟消息 Exchange（需要插件）
        delayed_notification = {
            exchange = "delayed-events",
            routing_key = "notification",
            type = "x-delayed-message"  # ✅ 插件类型
        }
```

**注意**：如果同时在配置中指定了 `type` 和在发送消息时使用了 `WithExchangeType`，**发送时的类型会覆盖配置中的类型**。

## 最佳实践

1. **✅ 在发送消息时指定 Exchange 类型**（推荐）：
   ```go
   // 每次发送时可以指定不同的 Exchange 类型
   messaging.PublishMessage(ctx, publisher, msg,
       messaging.WithExchangeType(messaging.ExchangeTypeFanout),
   )
   ```

2. **使用类型常量**：在代码中使用 `messaging.ExchangeTypeTopic` 等常量，而不是字符串

3. **配置中指定类型**（可选）：如果大多数消息使用相同的 Exchange 类型，可以在配置中指定

4. **插件类型检查**：使用 `IsPluginType()` 检查是否为插件类型，确保已安装对应插件

5. **自动声明**：依赖自动声明功能，不需要手动创建 Exchange

6. **类型一致性**：确保同一个 Exchange 名称始终使用相同的类型，避免类型冲突

## 相关文档

- [Exchange 详解](./USAGE_EXCHANGE.md) - Exchange 的基本概念和使用场景
- [广播发送](./USAGE_BROADCAST.md) - Fanout Exchange 的详细说明
- [配置详解](./USAGE_CONFIG.md) - 完整的配置选项说明
