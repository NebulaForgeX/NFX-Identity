package directory

import (
	"errors"
	"nfxid/messages"
	"nfxid/pkgs/rabbitmqx/messaging"
)

// UserUpdateMessage 用户更新消息
// MessageType 会自动从类型名生成: "user_update"
// RoutingKey 通过嵌入 messages.DirectoryTopic 自动提供: "directory"
type UserUpdateMessage struct {
	messages.DirectoryTopic
	UserID    string                 `json:"user_id"`    // 用户 ID
	UpdatedAt string                 `json:"updated_at"` // 更新时间
	Changes   map[string]interface{} `json:"changes"`    // 变更的字段
}

// Validate 验证消息数据的有效性
func (m UserUpdateMessage) Validate() error {
	if m.UserID == "" {
		return errors.New("user_id is required")
	}
	return nil
}

// =============== Cache Invalidation Messages ===============
// 当用户更新时，需要通知其他 7 个服务清除相关缓存
// 
// 正确的方式：使用 Fanout Exchange 或 Topic Exchange
// - 发送一条消息到一个 Exchange
// - Exchange 自动将消息路由到所有绑定的 Queue
// - 这才是 RabbitMQ 的核心优势！

// UserCacheInvalidateMessage 用户缓存清除消息（广播给所有服务）
// MessageType 会自动从类型名生成: "user_cache_invalidate"
// 
// 使用方式：
// 1. 配置一个 Fanout Exchange（推荐用于广播）
// 2. 所有服务的 Queue 都绑定到这个 Exchange
// 3. 只发送一条消息，RabbitMQ 自动分发到所有绑定的 Queue
type UserCacheInvalidateMessage struct {
	UserID        string `json:"user_id"`         // 要清除缓存的用户 ID
	Prefix        string `json:"prefix"`          // Cache prefix，例如 "user"
	Namespace     string `json:"namespace"`       // Cache namespace，可选，例如 tenantId
	Reason        string `json:"reason"`          // 清除原因，例如 "user_updated"
	InvalidatedAt string `json:"invalidated_at"`  // 清除时间
}

// RoutingKey 返回消息的路由键
// 对于 Fanout Exchange，RoutingKey 会被忽略，所有绑定的队列都会收到消息
// 对于 Topic Exchange，可以使用通配符 BindingKey 匹配
func (UserCacheInvalidateMessage) RoutingKey() messaging.MessageKey {
	// 返回一个通用的键，在配置中映射到 Fanout Exchange 或 Topic Exchange
	// 配置示例：
	// [rabbitmq.producer_exchanges]
	//     user_cache_invalidate = { exchange = "cache-broadcast", routing_key = "" }
	return "user_cache_invalidate"
}

// Validate 验证消息数据的有效性
func (m UserCacheInvalidateMessage) Validate() error {
	if m.UserID == "" {
		return errors.New("user_id is required")
	}
	if m.Prefix == "" {
		return errors.New("prefix is required")
	}
	return nil
}
