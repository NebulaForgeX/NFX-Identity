# Exchange 详解 / Exchange Guide

Exchange（交换机）是 RabbitMQ 的核心概念，负责接收消息并根据路由规则将消息路由到队列。

## Exchange 是什么？

```
Publisher → Exchange → Queue → Consumer
            ↓
         RoutingKey
```

Exchange 是消息路由中心：
1. **接收消息**：Publisher 不直接发送消息到 Queue，而是发送到 Exchange
2. **路由消息**：根据 Exchange 类型和 RoutingKey，将消息路由到匹配的 Queue
3. **解耦发布者和消费者**：Publisher 不需要知道消息最终会到哪个 Queue

## Exchange 类型

RabbitMQ 支持多种 Exchange 类型，包括基本类型和插件类型。`rabbitmqx` 提供了类型安全的 `ExchangeType` 来管理这些类型。

### ExchangeType 类型定义

`rabbitmqx` 使用 `messaging.ExchangeType` 类型来定义 Exchange 类型，类似 `Priority` 类型：

```go
import "nfxid/pkgs/rabbitmqx/messaging"

// 基本类型常量
messaging.ExchangeTypeDirect   // "direct"
messaging.ExchangeTypeTopic    // "topic" (默认)
messaging.ExchangeTypeFanout   // "fanout"
messaging.ExchangeTypeHeaders  // "headers"

// 插件类型常量
messaging.ExchangeTypeDelayedMessage  // "x-delayed-message"
messaging.ExchangeTypeConsistentHash  // "x-consistent-hash"
messaging.ExchangeTypeSharding       // "x-sharding"
// ... 更多插件类型
```

### 基本 Exchange 类型

#### 1. Topic Exchange（主题交换机）- 默认类型

**特点**：
- 支持通配符匹配（`*` 和 `#`）
- `*` 匹配一个单词，`#` 匹配零个或多个单词
- 最灵活的路由方式

**配置示例**：
```toml
[rabbitmq.exchange]
    name = ""
    type = "topic"  # 或使用 messaging.ExchangeTypeTopic
    durable = true
```

**RoutingKey 示例**：
- `"user.created"` - 精确匹配
- `"user.*"` - 匹配 `user.created`, `user.updated` 等
- `"user.#"` - 匹配 `user.created`, `user.updated.profile` 等

**使用场景**：需要灵活路由、按消息类型路由

#### 2. Direct Exchange（直连交换机）

**特点**：
- 精确匹配 RoutingKey
- 完全匹配才路由

**配置示例**：
```toml
[rabbitmq.exchange]
    name = "direct-exchange"
    type = "direct"  # 或使用 messaging.ExchangeTypeDirect
    durable = true
```

**使用场景**：点对点通信、精确路由

#### 3. Fanout Exchange（扇出交换机）

**特点**：
- 忽略 RoutingKey
- 广播到所有绑定的队列

**配置示例**：
```toml
[rabbitmq.exchange]
    name = "broadcast-exchange"
    type = "fanout"  # 或使用 messaging.ExchangeTypeFanout
    durable = true
```

**使用场景**：系统通知、配置更新、缓存失效

详细说明请参考：[USAGE_BROADCAST.md](./USAGE_BROADCAST.md)

#### 4. Headers Exchange（头交换机）

**特点**：
- 根据消息头（Headers）匹配
- 忽略 RoutingKey
- 支持 `x-match` 参数（`all` 或 `any`）

**配置示例**：
```toml
[rabbitmq.exchange]
    name = "headers-exchange"
    type = "headers"  # 或使用 messaging.ExchangeTypeHeaders
    durable = true
```

**使用场景**：基于消息头进行复杂路由

### 插件 Exchange 类型

RabbitMQ 支持通过插件扩展 Exchange 类型。`rabbitmqx` 支持以下插件类型：

#### 1. x-delayed-message（延迟消息 Exchange）

**需要插件**：`rabbitmq-delayed-message-exchange`

**安装**：
```bash
rabbitmq-plugins enable rabbitmq_delayed_message_exchange
```

**特点**：
- 支持延迟消息投递
- 通过 `x-delay` 消息头指定延迟时间（毫秒）
- 底层使用其他 Exchange 类型（如 topic）

**配置示例**：
```toml
[rabbitmq.producer_exchanges]
    delayed_notification = {
        exchange = "delayed-events",
        routing_key = "notification",
        type = "x-delayed-message"  # ✅ 插件类型
    }
```

**使用示例**：
```go
// 发布延迟消息（延迟 5 秒）
messaging.PublishMessage(ctx, publisher, msg,
    messaging.WithMetadata(map[string]string{
        "x-delay": "5000",  // 延迟 5 秒（毫秒）
    }),
)
```

#### 2. x-consistent-hash（一致性哈希 Exchange）

**需要插件**：`rabbitmq-consistent-hash-exchange`

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

#### 3. x-sharding（分片 Exchange）

**需要插件**：`rabbitmq-sharding`

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

#### 4. 其他插件类型

- `x-modulus-hash`：模数哈希 Exchange
- `x-random`：随机 Exchange
- `x-jms-topic`：JMS 主题 Exchange

### ✅ 为每个消息配置不同的 Exchange 类型

**重要功能**：可以为每个消息配置不同的 Exchange 类型，而不是全局配置！

```toml
[rabbitmq.exchange]
    name = "nfxid-events"
    type = "topic"  # 默认类型，如果 ProducerRouting 中未指定 type，使用此值
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
    
    # 使用延迟消息 Exchange（需要插件）
    delayed_notification = {
        exchange = "delayed-events",
        routing_key = "notification",
        type = "x-delayed-message"  # ✅ 插件类型
    }
```

**工作原理**：
1. 在创建 Publisher 时，收集所有 `ProducerExchanges` 中指定的 Exchange 名称和类型
2. 使用 amqp091-go 直接连接 RabbitMQ，自动声明所有需要的 Exchange
3. 每个 Exchange 使用配置中指定的类型（如果未指定，使用全局类型）

**优点**：
- ✅ 配置更清晰
- ✅ 可以在配置中指定 Exchange 类型
- ✅ **自动声明 Exchange**：不需要手动创建
- ✅ **向后兼容**：如果 `type` 为空，使用全局 `ExchangeConfig.Type`

## Exchange 配置

### Exchange 名称（Name）

```toml
[rabbitmq.exchange]
    name = ""  # 如果为空，会根据消息键自动生成
```

**自动生成规则**：
- 如果 `name` 为空，会根据消息键自动生成 exchange 名称

**固定名称**：
```toml
[rabbitmq.exchange]
    name = "my-exchange"  # 所有消息都使用这个 exchange
```

### Exchange 持久化（Durable）

```toml
[rabbitmq.exchange]
    durable = true  # 持久化，服务器重启后仍然存在
```

- `true`：持久化，服务器重启后仍然存在
- `false`：非持久化，服务器重启后删除

### Exchange 自动删除（AutoDelete）

```toml
[rabbitmq.exchange]
    auto_delete = false  # 当没有绑定队列时自动删除
```

- `true`：当没有绑定的队列时自动删除
- `false`：即使没有绑定的队列也保留

## Exchange 与 Queue 的绑定

Exchange 和 Queue 通过 **Binding** 连接：

```toml
[rabbitmq.consumer_queues]
    user_created = { queue = "user-queue", binding_key = "user.created" }
```

**绑定关系**：
```
Exchange (topic) ← BindingKey: "user.*" ← Queue
```

### BindingKey 的作用

**BindingKey（绑定键）** 是 Queue 绑定到 Exchange 时使用的匹配键，用于决定哪些消息会被路由到这个 Queue。

#### 工作流程

```
Publisher → Exchange (RoutingKey: "user.created")
                ↓
           匹配 BindingKey
                ↓
    BindingKey: "user.*" → Queue1 ✅ (匹配)
    BindingKey: "order.*" → Queue2 ❌ (不匹配)
    BindingKey: "user.created" → Queue3 ✅ (精确匹配)
```

#### RoutingKey vs BindingKey

| 概念 | 位置 | 作用 | 示例 |
|------|------|------|------|
| **RoutingKey** | Publisher 发送消息时 | 消息的路由键，标识消息类型 | `"user.created"` |
| **BindingKey** | Queue 绑定 Exchange 时 | 队列的匹配键，决定接收哪些消息 | `"user.*"` |

**匹配规则**：
- Publisher 发送消息时携带 **RoutingKey**（如 `"user.created"`）
- Exchange 根据类型和 RoutingKey 匹配每个 Queue 的 **BindingKey**
- 如果匹配，消息路由到对应的 Queue

#### 不同 Exchange 类型的匹配规则

##### 1. Topic Exchange（主题交换机）

支持通配符匹配：

```toml
[rabbitmq.consumer_queues]
    # 精确匹配
    user_created = { queue = "user-queue", binding_key = "user.created" }
    
    # 通配符匹配：* 匹配一个单词
    user_events = { queue = "user-events-queue", binding_key = "user.*" }
    # 匹配: user.created, user.updated, user.deleted
    
    # 通配符匹配：# 匹配零个或多个单词
    all_events = { queue = "all-events-queue", binding_key = "user.#" }
    # 匹配: user.created, user.updated.profile, user.deleted.reason
```

**示例**：
- RoutingKey: `"user.created"` → 匹配 BindingKey: `"user.created"`, `"user.*"`, `"user.#"`
- RoutingKey: `"user.updated.profile"` → 匹配 BindingKey: `"user.#"`（不匹配 `"user.*"`）

##### 2. Direct Exchange（直连交换机）

精确匹配：

```toml
[rabbitmq.consumer_queues]
    order_paid = { queue = "order-queue", binding_key = "order.paid" }
    # 只有 RoutingKey 完全等于 "order.paid" 才会匹配
```

**示例**：
- RoutingKey: `"order.paid"` → 匹配 BindingKey: `"order.paid"` ✅
- RoutingKey: `"order.cancelled"` → 不匹配 BindingKey: `"order.paid"` ❌

##### 3. Fanout Exchange（扇出交换机）

忽略 BindingKey：

```toml
[rabbitmq.consumer_queues]
    notification = { queue = "service-a-queue", binding_key = "" }
    # Fanout Exchange 忽略 BindingKey，所有绑定的队列都会收到消息
```

**示例**：
- 无论 RoutingKey 是什么，所有绑定的 Queue 都会收到消息
- BindingKey 可以设置为空字符串 `""`

##### 4. Headers Exchange（头交换机）

忽略 BindingKey，根据消息头匹配：

```toml
[rabbitmq.consumer_queues]
    notification = { queue = "notification-queue", binding_key = "" }
    # Headers Exchange 忽略 BindingKey，根据消息头匹配
```

#### 配置示例

```toml
[rabbitmq]
    [rabbitmq.exchange]
        name = "user-events"
        type = "topic"  # 使用 Topic Exchange 支持通配符

    # Publisher 配置：发送消息时使用 RoutingKey
    [rabbitmq.producer_exchanges]
        user_created = { exchange = "user-events", routing_key = "user.created" }
        user_updated = { exchange = "user-events", routing_key = "user.updated" }

    # Consumer 配置：Queue 绑定 Exchange 时使用 BindingKey
    [rabbitmq.consumer_queues]
        # 只接收 user.created 消息
        user_created = { queue = "user-created-queue", binding_key = "user.created" }
        
        # 接收所有 user.* 开头的消息
        user_all = { queue = "user-all-queue", binding_key = "user.*" }
        
        # 接收所有 user 相关的消息（包括子路径）
        user_all_deep = { queue = "user-all-deep-queue", binding_key = "user.#" }
```

#### 实际应用场景

**场景 1：精确路由**
```toml
# 只接收特定类型的消息
user_created = { queue = "user-created-queue", binding_key = "user.created" }
```

**场景 2：模式匹配**
```toml
# 接收所有用户相关事件
user_events = { queue = "user-events-queue", binding_key = "user.*" }
```

**场景 3：广播**
```toml
# Fanout Exchange，忽略 BindingKey
notification = { queue = "service-a-queue", binding_key = "" }
```

#### 注意事项

1. **BindingKey 必须与 Exchange 类型匹配**：
   - Topic Exchange：支持通配符（`*`, `#`）
   - Direct Exchange：必须精确匹配
   - Fanout Exchange：忽略 BindingKey

2. **BindingKey 为空时的行为**：
   - Topic/Direct Exchange：BindingKey 为空时，使用消息键作为默认值
   - Fanout Exchange：BindingKey 为空是正常的（会被忽略）

3. **一个 Queue 可以有多个 BindingKey**：
   - 同一个 Queue 可以绑定到同一个 Exchange 多次，使用不同的 BindingKey
   - 这样可以接收多种类型的消息

## RabbitMQ vs Kafka：核心区别

### 消息模型对比

#### Kafka 模型：
```
Producer → Topic → Partition → Consumer Group → Consumer
```
- **Topic**：消息分类（类似 Exchange）
- **Partition**：Topic 的分区（提高并发）
- **Consumer Group**：消费者组（负载均衡）

#### RabbitMQ 模型：
```
Producer → Exchange → Queue → Consumer
           ↓
        RoutingKey
```
- **Exchange**：消息路由中心（类似 Topic，但更灵活）
- **Queue**：消息存储（类似 Partition，但更灵活）
- **RoutingKey**：路由键（Kafka 没有这个概念）

### 路由能力对比

| 场景 | Kafka | RabbitMQ |
|------|-------|----------|
| **精确路由** | ✅ Topic 名称 | ✅ Direct Exchange |
| **模式匹配** | ❌ 不支持 | ✅ Topic Exchange (`user.*`) |
| **广播** | ✅ 多个 Consumer Group | ✅ Fanout Exchange |
| **多队列路由** | ❌ 需要多个 Topic | ✅ 一个 Exchange → 多个 Queue |
| **动态路由** | ❌ 需要重新创建 Topic | ✅ 动态绑定 Queue |

### 实际场景对比

#### 场景 1：用户事件需要发送到多个服务

**Kafka 方式**：
```go
// 需要为每个服务创建不同的 Topic
publisher.Publish("user-events-auth", event)
publisher.Publish("user-events-notification", event)
publisher.Publish("user-events-analytics", event)
```
问题：需要发布 3 次

**RabbitMQ 方式**：
```toml
[rabbitmq.exchange]
    name = "user-events"
    type = "topic"
```
```go
// 只需发布一次
messaging.PublishMessage(ctx, publisher, UserCreatedEvent{...})
// Exchange 自动路由到所有绑定的 Queue
```
优势：发布一次，自动路由到多个服务

#### 场景 2：需要按消息类型路由

**Kafka 方式**：
```go
// 需要创建多个 Topic
publisher.Publish("user-created", event)
publisher.Publish("user-updated", event)
publisher.Publish("user-deleted", event)
```

**RabbitMQ 方式**：
```toml
[rabbitmq.exchange]
    name = "user-events"
    type = "topic"

[rabbitmq.consumer_queues]
    user_created = { queue = "user-queue", binding_key = "user.created" }
    user_updated = { queue = "user-queue", binding_key = "user.updated" }
```
```go
// 使用同一个 Exchange，通过 RoutingKey 区分
messaging.PublishMessage(ctx, publisher, UserCreatedEvent{...})
messaging.PublishMessage(ctx, publisher, UserUpdatedEvent{...})
```
优势：一个 Exchange，灵活的路由规则

### Exchange vs Topic 总结

| 特性 | Kafka Topic | RabbitMQ Exchange |
|------|-------------|-------------------|
| **作用** | 消息分类存储 | 消息路由中心 |
| **路由能力** | 固定分区路由 | 灵活路由规则 |
| **模式匹配** | ❌ 不支持 | ✅ 支持通配符 |
| **多目标路由** | ❌ 需要多个 Topic | ✅ 一个 Exchange → 多个 Queue |
| **动态绑定** | ❌ 需要重新创建 Topic | ✅ 动态绑定 Queue |
| **消息存储** | ✅ Topic 存储消息 | ❌ Exchange 不存储消息 |
| **消息顺序** | ✅ 分区内有序 | ✅ Queue 内有序 |

## ✅ 为每个消息配置不同的 Exchange 类型

**重要功能**：可以为每个消息配置不同的 Exchange 类型，而不是全局配置！

### 配置方式

在 `ProducerExchanges` 中为每个消息指定 Exchange 类型：

```toml
[rabbitmq.exchange]
    name = "nfxid-events"
    type = "topic"  # 默认类型，如果 ProducerRouting 中未指定 type，使用此值
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
    
    # 使用 Direct Exchange（精确路由）
    order_paid = {
        exchange = "order-events",
        routing_key = "order.paid",
        type = "direct"  # ✅ 指定为 Direct Exchange
    }
    
    # 使用延迟消息 Exchange（需要插件）
    delayed_notification = {
        exchange = "delayed-events",
        routing_key = "notification",
        type = "x-delayed-message"  # ✅ 插件类型
    }
```

### 工作原理

1. **收集 Exchange 配置**：在创建 Publisher 时，收集所有 `ProducerExchanges` 中指定的 Exchange 名称和类型
2. **自动声明 Exchange**：使用 amqp091-go 直接连接 RabbitMQ，自动声明所有需要的 Exchange
3. **类型优先级**：
   - 如果 `ProducerRouting.Type` 不为空，使用指定的类型
   - 如果 `ProducerRouting.Type` 为空，使用全局 `ExchangeConfig.Type`
   - 如果全局 `ExchangeConfig.Type` 也为空，使用默认值 `"topic"`

### 自动声明 Exchange

在创建 Publisher 时，会自动声明所有需要的 Exchange：

```go
publisher, err := rabbitmqx.NewPublisher(cfg)
// ✅ 自动声明所有配置中的 Exchange：
//    - nfxid-events (topic)
//    - cache-broadcast (fanout)
//    - order-events (direct)
//    - delayed-events (x-delayed-message)
```

**日志输出**：
```
✅ Declared exchange: nfxid-events (type: topic, durable: true)
✅ Declared exchange: cache-broadcast (type: fanout, durable: true)
✅ Declared exchange: order-events (type: direct, durable: true)
✅ Declared exchange: delayed-events (type: x-delayed-message, durable: true)
```

### 支持的 Exchange 类型

- **基本类型**：`direct`, `topic`, `fanout`, `headers`
- **插件类型**：`x-delayed-message`, `x-consistent-hash`, `x-sharding`, `x-modulus-hash`, `x-random`, `x-jms-topic`

详细说明请参考：[USAGE_EXCHANGE_TYPE.md](./USAGE_EXCHANGE_TYPE.md)

## 最佳实践

1. **生产环境使用持久化 Exchange**
   ```toml
   [rabbitmq.exchange]
       durable = true
       auto_delete = false
   ```

2. **✅ 在发送消息时动态指定 Exchange 类型**（推荐）
   ```go
   // 每次发送时可以指定不同的 Exchange 类型
   messaging.PublishMessage(ctx, publisher, msg,
       messaging.WithExchangeType(messaging.ExchangeTypeFanout),
   )
   ```

3. **在配置中指定 Exchange 类型**（可选）
   ```toml
   [rabbitmq.producer_exchanges]
       # ✅ 可选：在配置中指定 Exchange 类型
       directory = { 
           exchange = "nfxid-events", 
           routing_key = "directory.user.update",
           type = "topic"  # ✅ 可选，明确指定
       }
   ```

3. **使用 Topic Exchange 实现灵活路由**
   ```toml
   [rabbitmq.exchange]
       type = "topic"  # 默认，最灵活
   ```

4. **为不同服务使用不同的 Exchange**
   ```toml
   # 服务 A
   [rabbitmq.exchange]
       name = "service-a-events"
   
   # 服务 B
   [rabbitmq.exchange]
       name = "service-b-events"
   ```

5. **使用有意义的 Exchange 名称**
   ```toml
   [rabbitmq.exchange]
       name = "user-events"  # 而不是 "exchange1"
   ```

## 相关文档

- [基本用法](./USAGE_BASIC.md) - 快速上手指南
- [广播发送](./USAGE_BROADCAST.md) - Fanout Exchange 广播
- [配置详解](./USAGE_CONFIG.md) - 完整配置选项
- [Exchange 类型](./USAGE_EXCHANGE_TYPE.md) - ✅ ExchangeType 类型详解，支持为每个消息配置不同的 Exchange 类型
