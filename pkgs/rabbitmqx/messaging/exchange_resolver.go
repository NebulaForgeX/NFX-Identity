package messaging

import "fmt"

// MessageKey 消息键类型，用于标识消息的逻辑键。
// 例如: "access", "clients", "directory"
type MessageKey = string

// ExchangeResolver Exchange 解析器。
// 用于在消息键（MessageKey）和 Exchange + RoutingKey 之间进行映射。
type ExchangeResolver struct {
	KeyToExchange map[MessageKey]ExchangeRouting // 消息键到 Exchange + RoutingKey 的映射
	ExchangeToKey map[string]MessageKey         // Exchange+RoutingKey 组合到消息键的映射（用于反向查找）
}

// ExchangeRouting 包含 Exchange 名称和 RoutingKey
type ExchangeRouting struct {
	Exchange   string
	RoutingKey string
}

// NewExchangeResolver 创建新的 Exchange 解析器。
// 根据提供的键值对映射创建双向映射。
//
// 参数:
//   - keyToRouting: 事件键到 ExchangeRouting 的映射
//
// 返回:
//   - *ExchangeResolver: Exchange 解析器实例
//   - error: 如果 Exchange 或 RoutingKey 为空或存在重复，返回错误
//
// 示例:
//
//	resolver, err := messaging.NewExchangeResolver(map[messaging.MessageKey]messaging.ExchangeRouting{
//	    "user_created": {Exchange: "user-events", RoutingKey: "user.created"},
//	    "order_paid":   {Exchange: "order-events", RoutingKey: "order.paid"},
//	})
func NewExchangeResolver(keyToRouting map[MessageKey]ExchangeRouting) (*ExchangeResolver, error) {
	exchangeToKey := make(map[string]MessageKey)
	for k, routing := range keyToRouting {
		if routing.Exchange == "" {
			return nil, fmt.Errorf("exchange name is empty for key %q", k)
		}
		if routing.RoutingKey == "" {
			return nil, fmt.Errorf("routing key is empty for key %q", k)
		}
		// 使用 Exchange:RoutingKey 作为唯一键
		key := routing.Exchange + ":" + routing.RoutingKey
		if _, dup := exchangeToKey[key]; dup {
			return nil, fmt.Errorf("duplicate exchange:routing_key combination %q", key)
		}
		exchangeToKey[key] = k
	}
	return &ExchangeResolver{
		KeyToExchange: keyToRouting,
		ExchangeToKey: exchangeToKey,
	}, nil
}

// GetExchange 根据消息键获取 Exchange 和 RoutingKey。
//
// 参数:
//   - key: 消息键，例如 "user_created"
//
// 返回:
//   - ExchangeRouting: Exchange 和 RoutingKey，如果找到
//   - bool: 是否找到对应的 Exchange 和 RoutingKey
//
// 示例:
//
//	routing, ok := resolver.GetExchange("user_created")
//	if ok {
//	    fmt.Println("Exchange:", routing.Exchange, "RoutingKey:", routing.RoutingKey)
//	}
func (e *ExchangeResolver) GetExchange(key MessageKey) (ExchangeRouting, bool) {
	v, ok := e.KeyToExchange[key]
	return v, ok
}

// GetKey 根据 Exchange 和 RoutingKey 获取消息键。
//
// 参数:
//   - exchange: Exchange 名称
//   - routingKey: RoutingKey
//
// 返回:
//   - MessageKey: 消息键，如果找到
//   - bool: 是否找到对应的消息键
func (e *ExchangeResolver) GetKey(exchange, routingKey string) (MessageKey, bool) {
	key := exchange + ":" + routingKey
	v, ok := e.ExchangeToKey[key]
	return v, ok
}
