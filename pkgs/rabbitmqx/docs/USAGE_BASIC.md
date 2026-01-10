# 基本用法 / Basic Usage

## 1. 配置

在 `example.toml` 或你的配置文件中添加 RabbitMQ 配置：

```toml
[rabbitmq]
    uri = "amqp://guest:guest@localhost:5672/"
    client_id = "nfxid-service"

    [rabbitmq.producer]
        confirm_delivery = true
        default_priority = 0

    [rabbitmq.consumer]
        queue_name = ""
        consumer_tag = "nfxid-consumer"
        prefetch_count = 10

    [rabbitmq.exchange]
        name = ""
        type = "topic"
        durable = true

    [rabbitmq.queue]
        durable = true
        max_priority = 0

    # ProducerExchanges 映射消息键到 Exchange 和 RoutingKey
    [rabbitmq.producer_exchanges]
        user_created = { exchange = "user-events", routing_key = "user.created" }
        order_paid = { exchange = "order-events", routing_key = "order.paid" }

    # ConsumerQueues 映射消息键到 Queue 和 BindingKey
    [rabbitmq.consumer_queues]
        user_created = { queue = "user-queue", binding_key = "user.created" }
        order_paid = { queue = "order-queue", binding_key = "order.paid" }
```

## 2. 创建 Publisher（发布者）

```go
package main

import (
    "nfxid/pkgs/rabbitmqx"
    "nfxid/pkgs/rabbitmqx/messaging"
)

func createPublisher(cfg *rabbitmqx.Config) (*messaging.BusPublisher, error) {
    // 验证配置
    if err := cfg.Validate(); err != nil {
        return nil, fmt.Errorf("invalid rabbitmq config: %w", err)
    }

    // 创建 Publisher
    publisher, err := rabbitmqx.NewPublisher(cfg)
    if err != nil {
        return nil, fmt.Errorf("failed to create rabbitmq publisher: %w", err)
    }

    return publisher, nil
}
```

## 3. 创建 Subscriber（订阅者）

```go
func createSubscriber(cfg *rabbitmqx.Config) (*messaging.BusSubscriber, error) {
    // 验证配置
    if err := cfg.Validate(); err != nil {
        return nil, fmt.Errorf("invalid rabbitmq config: %w", err)
    }

    // 创建 Subscriber
    subscriber, err := rabbitmqx.NewSubscriber(cfg)
    if err != nil {
        return nil, fmt.Errorf("failed to create rabbitmq subscriber: %w", err)
    }

    return subscriber, nil
}
```

## 4. 定义消息类型

```go
package access

import "nfxid/pkgs/rabbitmqx/messaging"

// 方式一：自动生成 MessageType（推荐）
type GrantsInvalidateCacheMessage struct {
    ID string `json:"id"`
}

func (GrantsInvalidateCacheMessage) RoutingKey() messaging.MessageKey {
    return "access"  // 返回消息键
}

// 方式二：自定义 MessageType
type UserCreatedMessage struct {
    UserID string `json:"user_id"`
    Name   string `json:"name"`
}

func (UserCreatedMessage) MessageType() messaging.MessageType {
    return "custom.user.created"
}

func (UserCreatedMessage) RoutingKey() messaging.MessageKey {
    return "custom"
}
```

## 5. 发布消息

```go
package main

import (
    "context"
    "nfxid/pkgs/rabbitmqx/messaging"
)

func publishMessage(ctx context.Context, publisher *messaging.BusPublisher) error {
    msg := access.GrantsInvalidateCacheMessage{
        ID: "grant-123",
    }

    // 发布消息（自动根据消息类型找到对应的 Exchange + RoutingKey）
    err := messaging.PublishMessage(ctx, publisher, msg)
    if err != nil {
        return fmt.Errorf("failed to publish message: %w", err)
    }

    return nil
}
```

## 6. 创建 MessageRouter 并注册处理器

```go
package eventbus

import (
    "context"
    "time"
    "nfxid/pkgs/rabbitmqx"
    "nfxid/pkgs/rabbitmqx/messaging"
    "nfxid/pkgs/logx"
    wmMiddleware "github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

type Router struct {
    *messaging.MessageRouter
}

func NewRouter(cfg *rabbitmqx.Config) (*Router, error) {
    // 创建订阅者
    sub, err := rabbitmqx.NewSubscriber(cfg)
    if err != nil {
        return nil, err
    }

    // 创建路由器
    router, err := messaging.NewMessageRouter(sub, messaging.MessageRouterConfig{
        CloseTimeout: 10 * time.Second,
        Logger:       logx.NewZapWatermillLogger(logx.L()),
    })
    if err != nil {
        return nil, err
    }

    // 添加中间件（可选）
    router.AddMiddleware(
        wmMiddleware.CorrelationID,
        wmMiddleware.Recoverer,
        wmMiddleware.Retry{
            MaxRetries:      3,
            InitialInterval: 200 * time.Millisecond,
            MaxInterval:     2 * time.Second,
            Multiplier:      2.0,
        }.Middleware,
        wmMiddleware.Timeout(10*time.Second),
    )

    return &Router{
        MessageRouter: router,
    }, nil
}

// 注册消息处理器
func (r *Router) RegisterHandlers() {
    // 注册处理器（自动根据消息类型路由）
    messaging.RegisterHandler(r.MessageRouter, r.handleUserCreated)
    messaging.RegisterHandler(r.MessageRouter, r.handleUserUpdated)
}

// 消息处理器示例
func (r *Router) handleUserCreated(ctx context.Context, msg UserCreatedMessage, rawMsg *message.Message) error {
    logx.S().Infof("Received user created message: %+v", msg)
    // 处理消息逻辑
    return nil
}

func (r *Router) handleUserUpdated(ctx context.Context, msg UserUpdatedMessage, rawMsg *message.Message) error {
    logx.S().Infof("Received user updated message: %+v", msg)
    // 处理消息逻辑
    return nil
}

// 启动路由器
func (r *Router) Run(ctx context.Context) error {
    logx.S().Info("Starting RabbitMQ message router...")
    return r.MessageRouter.Run(ctx)
}

// 关闭路由器
func (r *Router) Close() error {
    return r.MessageRouter.Close()
}
```

## 7. 完整示例

```go
package server

import (
    "context"
    "nfxid/pkgs/rabbitmqx"
    "nfxid/pkgs/rabbitmqx/messaging"
)

type Deps struct {
    RabbitMQConfig *rabbitmqx.Config
}

func NewServer(deps Deps) (*Server, error) {
    // 创建 Publisher
    publisher, err := rabbitmqx.NewPublisher(deps.RabbitMQConfig)
    if err != nil {
        return nil, fmt.Errorf("failed to create publisher: %w", err)
    }

    // 创建 Router
    router, err := NewRouter(deps.RabbitMQConfig)
    if err != nil {
        return nil, fmt.Errorf("failed to create router: %w", err)
    }

    // 注册处理器
    router.RegisterHandlers()

    return &Server{
        publisher: publisher,
        router:    router,
    }, nil
}

type Server struct {
    publisher *messaging.BusPublisher
    router    *Router
}

func (s *Server) Start(ctx context.Context) error {
    // 启动消息路由器
    go func() {
        if err := s.router.Run(ctx); err != nil {
            logx.S().Errorf("Router error: %v", err)
        }
    }()

    // 发布消息示例
    msg := UserCreatedMessage{ID: "123", Email: "test@example.com"}
    if err := messaging.PublishMessage(ctx, s.publisher, msg); err != nil {
        return err
    }

    return nil
}

func (s *Server) Stop() error {
    if err := s.router.Close(); err != nil {
        return err
    }
    return nil
}
```

## 下一步

- 了解如何设置消息优先级：[USAGE_PRIORITY.md](./USAGE_PRIORITY.md)
- 了解如何实现广播发送：[USAGE_BROADCAST.md](./USAGE_BROADCAST.md)
- 查看完整配置选项：[USAGE_CONFIG.md](./USAGE_CONFIG.md)
