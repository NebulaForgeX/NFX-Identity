# 配置详解 / Configuration Guide

完整的 RabbitMQ 配置选项说明。

## 配置示例

```toml
[rabbitmq]
    uri = "amqp://guest:guest@localhost:5672/"
    client_id = "nfxid-service"

    [rabbitmq.producer]
        mandatory = false
        immediate = false
        content_type = "application/json"
        delivery_mode = 2
        confirm_delivery = true
        channel_pool_size = 10
        transactional = false
        default_priority = 0

    [rabbitmq.consumer]
        queue_name = ""
        consumer_tag = "nfxid-consumer"
        auto_ack = false
        exclusive = false
        no_local = false
        no_wait = false
        no_requeue_on_nack = false
        prefetch_count = 10
        prefetch_size = 0
        qos_global = false

    [rabbitmq.connection]
        [rabbitmq.connection.tls]
            enabled = false
            insecure_skip_verify = false
            cert_file = ""
            key_file = ""
            ca_file = ""
            server_name = ""

        [rabbitmq.connection.reconnect]
            enabled = true
            initial_interval = "1s"
            randomization_factor = 0.5
            multiplier = 1.5
            max_interval = "30s"

        [rabbitmq.connection.amqp]
            vhost = "/"
            heartbeat = 10
            locale = "en_US"
            channel_max = 0
            frame_size = 0

    [rabbitmq.exchange]
        name = ""
        type = "topic"
        durable = true
        auto_delete = false
        internal = false
        no_wait = false

    [rabbitmq.queue]
        durable = true
        auto_delete = false
        exclusive = false
        no_wait = false
        max_priority = 0
        message_ttl = 0
        max_length = 0
        max_length_bytes = 0
        dead_letter_exchange = ""
        dead_letter_routing_key = ""

    [rabbitmq.queue_bind]
        routing_key = ""
        no_wait = false

    [rabbitmq.producer_exchanges]
        user_created = { exchange = "user-events", routing_key = "user.created" }
        order_paid = { exchange = "order-events", routing_key = "order.paid" }

    [rabbitmq.consumer_queues]
        user_created = { queue = "user-queue", binding_key = "user.created" }
        order_paid = { queue = "order-queue", binding_key = "order.paid" }
```

## 配置项说明

### 基础配置

#### `uri`
- **类型**：`string`
- **说明**：RabbitMQ 连接 URI
- **格式**：`amqp://user:password@host:port/vhost`
- **示例**：`amqp://guest:guest@localhost:5672/`

#### `client_id`
- **类型**：`string`
- **说明**：客户端标识符
- **用途**：用于日志和监控

### Producer 配置

#### `mandatory`
- **类型**：`bool`
- **默认值**：`false`
- **说明**：如果为 true，消息无法路由时会返回错误

#### `immediate`
- **类型**：`bool`
- **默认值**：`false`
- **说明**：如果为 true，消息无法立即投递给消费者时会返回错误

#### `confirm_delivery`
- **类型**：`bool`
- **默认值**：`false`
- **说明**：是否等待服务器确认（Publisher Confirms）

#### `channel_pool_size`
- **类型**：`int`
- **默认值**：`0`
- **说明**：通道池大小，0 表示不池化

#### `default_priority`
- **类型**：`uint8`
- **默认值**：`0`
- **说明**：默认消息优先级（0-255），0 表示不设置优先级

### Consumer 配置

#### `queue_name`
- **类型**：`string`
- **默认值**：`""`
- **说明**：队列名称，如果为空则根据消息键自动生成

#### `prefetch_count`
- **类型**：`int`
- **默认值**：`0`
- **说明**：预取数量，0 表示不限制

#### `prefetch_size`
- **类型**：`int`
- **默认值**：`0`
- **说明**：预取大小（字节），0 表示不限制

#### `qos_global`
- **类型**：`bool`
- **默认值**：`false`
- **说明**：QOS 是否全局应用

### Exchange 配置

#### `name`
- **类型**：`string`
- **默认值**：`""`
- **说明**：Exchange 名称，如果为空则根据消息键自动生成

#### `type`
- **类型**：`string`
- **默认值**：`"topic"`
- **可选值**：`"direct"`, `"topic"`, `"fanout"`, `"headers"`
- **说明**：Exchange 类型

#### `durable`
- **类型**：`bool`
- **默认值**：`true`
- **说明**：是否持久化，服务器重启后仍然存在

#### `auto_delete`
- **类型**：`bool`
- **默认值**：`false`
- **说明**：当没有绑定队列时自动删除

### Queue 配置

#### `durable`
- **类型**：`bool`
- **默认值**：`true`
- **说明**：是否持久化

#### `max_priority`
- **类型**：`int`
- **默认值**：`0`
- **说明**：队列最大优先级（0-255），0 表示不启用优先级队列

#### `message_ttl`
- **类型**：`int`
- **默认值**：`0`
- **说明**：消息 TTL（毫秒），0 表示不设置 TTL

#### `max_length`
- **类型**：`int`
- **默认值**：`0`
- **说明**：队列最大消息数，0 表示不限制

#### `max_length_bytes`
- **类型**：`int`
- **默认值**：`0`
- **说明**：队列最大字节数，0 表示不限制

#### `dead_letter_exchange`
- **类型**：`string`
- **默认值**：`""`
- **说明**：死信交换机名称

#### `dead_letter_routing_key`
- **类型**：`string`
- **默认值**：`""`
- **说明**：死信路由键

### 路由配置

#### `producer_exchanges`
- **类型**：`map[string]ProducerRouting`
- **说明**：映射消息键到 Exchange 和 RoutingKey
- **格式**：
  ```toml
  [rabbitmq.producer_exchanges]
      message_key = { exchange = "exchange-name", routing_key = "routing.key" }
  ```

#### `consumer_queues`
- **类型**：`map[string]ConsumerBinding`
- **说明**：映射消息键到 Queue 和 BindingKey
- **格式**：
  ```toml
  [rabbitmq.consumer_queues]
      message_key = { queue = "queue-name", binding_key = "binding.key" }
  ```

### TLS 配置

#### `enabled`
- **类型**：`bool`
- **默认值**：`false`
- **说明**：是否启用 TLS

#### `insecure_skip_verify`
- **类型**：`bool`
- **默认值**：`false`
- **说明**：是否跳过证书验证（仅用于开发环境）

### 重连配置

#### `enabled`
- **类型**：`bool`
- **默认值**：`true`
- **说明**：是否启用自动重连

#### `initial_interval`
- **类型**：`string`
- **默认值**：`"1s"`
- **说明**：初始重连间隔

#### `max_interval`
- **类型**：`string`
- **默认值**：`"30s"`
- **说明**：最大重连间隔

## 相关文档

- [基本用法](./USAGE_BASIC.md) - 快速上手指南
- [Exchange 详解](./USAGE_EXCHANGE.md) - Exchange 类型和使用场景
