# 高级功能 / Advanced Features

RabbitMQ 的高级功能包括延迟队列、死信队列、队列限制等。

## 1. 延迟队列（Delayed Queue / Message TTL）

RabbitMQ 支持消息 TTL（Time To Live），消息在队列中超过指定时间后会被删除或转发到死信队列。

### 配置消息 TTL

```toml
[rabbitmq.queue]
    message_ttl = 60000  # 消息 TTL（毫秒），60000 = 60 秒
```

**工作原理**：
- 消息在队列中停留超过 TTL 时间后，会被删除
- 如果配置了死信队列，过期消息会被转发到死信队列

### 使用场景

#### 场景 1：延迟处理

```toml
# 延迟处理：60 秒后处理
[rabbitmq.queue]
    message_ttl = 60000
    dead_letter_exchange = "processed-exchange"
    dead_letter_routing_key = "processed"
```

#### 场景 2：订单超时取消

```toml
# 订单 30 分钟后自动取消
[rabbitmq.queue]
    message_ttl = 1800000  # 30 分钟
    dead_letter_exchange = "order-timeout"
    dead_letter_routing_key = "order.cancelled"
```

## 2. 死信队列（Dead Letter Queue）

死信队列用于处理无法被正常消费的消息（过期、被拒绝、队列满等）。

### 配置死信队列

```toml
[rabbitmq.queue]
    dead_letter_exchange = "dlx"           # 死信交换机
    dead_letter_routing_key = "dlx.error"  # 死信路由键
```

### 触发条件

- 消息 TTL 过期
- 消息被拒绝且不重新入队
- 队列达到最大长度
- 消息无法路由（mandatory=true 时）

### 使用场景

#### 场景 1：错误处理

```toml
# 主队列
[rabbitmq.queue]
    dead_letter_exchange = "error-exchange"
    dead_letter_routing_key = "error"

# 死信队列（错误队列）
[rabbitmq.consumer_queues]
    error = { queue = "error-queue", binding_key = "error" }
```

#### 场景 2：重试机制

```toml
# 主队列配置 TTL
[rabbitmq.queue]
    message_ttl = 60000  # 60 秒后重试
    dead_letter_exchange = "retry-exchange"
    dead_letter_routing_key = "retry"
```

## 3. 队列限制（Queue Limits）

RabbitMQ 支持限制队列的大小，防止队列无限增长。

### 配置队列限制

```toml
[rabbitmq.queue]
    max_length = 10000      # 队列最大消息数
    max_length_bytes = 104857600  # 队列最大字节数（100MB）
```

**行为**：
- 当队列达到限制时，新消息会被拒绝
- 如果配置了死信队列，会转发到死信队列
- 适用于限流和防止内存溢出

### 使用场景

#### 场景 1：限流

```toml
# 限制队列最多 1000 条消息
[rabbitmq.queue]
    max_length = 1000
    dead_letter_exchange = "overflow-exchange"
    dead_letter_routing_key = "overflow"
```

#### 场景 2：防止内存溢出

```toml
# 限制队列最大 100MB
[rabbitmq.queue]
    max_length_bytes = 104857600
```

## 4. 发布确认（Publisher Confirms）

启用发布确认，确保消息已成功投递。

### 配置

```toml
[rabbitmq.producer]
    confirm_delivery = true  # 启用发布确认
```

启用后，发布消息会等待服务器确认，确保消息已成功投递。

## 5. 通道池（Channel Pool）

使用通道池可以提高性能，避免频繁创建和销毁通道。

### 配置

```toml
[rabbitmq.producer]
    channel_pool_size = 10  # 通道池大小
```

## 6. QOS 预取（QOS Prefetch）

控制消费者预取消息的数量，实现负载均衡。

### 配置

```toml
[rabbitmq.consumer]
    prefetch_count = 10  # 预取数量
    prefetch_size = 0   # 预取大小（字节）
    qos_global = false  # 是否全局应用
```

## 完整示例

### 示例：订单处理系统

```toml
[rabbitmq]
    uri = "amqp://guest:guest@localhost:5672/"

    [rabbitmq.exchange]
        name = "order-events"
        type = "topic"
        durable = true

    [rabbitmq.queue]
        durable = true
        # 延迟队列：30 分钟后处理超时订单
        message_ttl = 1800000
        dead_letter_exchange = "order-timeout"
        dead_letter_routing_key = "order.timeout"
        # 队列限制：最多 10000 条消息
        max_length = 10000
        max_length_bytes = 104857600

    [rabbitmq.producer]
        confirm_delivery = true
        channel_pool_size = 10

    [rabbitmq.consumer]
        prefetch_count = 10

    [rabbitmq.producer_exchanges]
        order_created = { exchange = "order-events", routing_key = "order.created" }
        order_timeout = { exchange = "order-timeout", routing_key = "order.timeout" }

    [rabbitmq.consumer_queues]
        order_created = { queue = "order-queue", binding_key = "order.created" }
        order_timeout = { queue = "order-timeout-queue", binding_key = "order.timeout" }
```

## 相关文档

- [基本用法](./USAGE_BASIC.md) - 快速上手指南
- [消息优先级](./USAGE_PRIORITY.md) - 设置消息优先级
- [配置详解](./USAGE_CONFIG.md) - 完整配置选项
