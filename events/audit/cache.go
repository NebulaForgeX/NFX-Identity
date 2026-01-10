package audit

import (
	"nfxid/events"
)

// ActorSnapshotsInvalidateCacheEvent ActorSnapshots 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.AuditTopic 自动提供
// Cache key 格式: {prefix[:namespace]}:entity:{id}
type ActorSnapshotsInvalidateCacheEvent struct {
	events.AuditTopic
	ID        string `json:"id"`         // 要清除的 Actor Snapshot ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "actor_snapshot"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// EventRetentionPoliciesInvalidateCacheEvent EventRetentionPolicies 缓存清除事件
type EventRetentionPoliciesInvalidateCacheEvent struct {
	events.AuditTopic
	ID        string `json:"id"`         // 要清除的 Event Retention Policy ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "event_retention_policy"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// EventSearchIndexInvalidateCacheEvent EventSearchIndex 缓存清除事件
type EventSearchIndexInvalidateCacheEvent struct {
	events.AuditTopic
	ID        string `json:"id"`         // 要清除的 Event Search Index ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "event_search_index"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// EventsInvalidateCacheEvent Events 缓存清除事件
type EventsInvalidateCacheEvent struct {
	events.AuditTopic
	ID        string `json:"id"`         // 要清除的 Event ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "event"
	Namespace string `json:"namespace"` // Cache namespace，可选
}

// HashChainCheckpointsInvalidateCacheEvent HashChainCheckpoints 缓存清除事件
type HashChainCheckpointsInvalidateCacheEvent struct {
	events.AuditTopic
	ID        string `json:"id"`         // 要清除的 Hash Chain Checkpoint ID
	Prefix    string `json:"prefix"`     // Cache prefix，例如 "hash_chain_checkpoint"
	Namespace string `json:"namespace"` // Cache namespace，可选
}
