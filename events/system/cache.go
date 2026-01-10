package system

import (
	"nfxid/events"
)

// SystemStateInvalidateCacheEvent SystemState 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.SystemTopic 自动提供
// Cache key 格式: {prefix[:namespace]}:entity:{id}
type SystemStateInvalidateCacheEvent struct {
	events.SystemTopic
	ID        string `json:"id"`         // 要清除的 System State ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "system_state"
	Namespace string `json:"namespace"` // Cache namespace，可选
}
