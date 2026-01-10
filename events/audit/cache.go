package audit

import (
	"nfxid/events"
)

// ActorSnapshotsInvalidateCacheEvent ActorSnapshots 缓存清除事件
// EventType 会自动从类型名生成，TopicKey 通过嵌入 events.AuditTopic 自动提供
type ActorSnapshotsInvalidateCacheEvent struct {
	events.AuditTopic
	ID string `json:"id"` // 要清除的 Actor Snapshot ID
}

// EventRetentionPoliciesInvalidateCacheEvent EventRetentionPolicies 缓存清除事件
type EventRetentionPoliciesInvalidateCacheEvent struct {
	events.AuditTopic
	ID string `json:"id"` // 要清除的 Event Retention Policy ID
}

// EventSearchIndexInvalidateCacheEvent EventSearchIndex 缓存清除事件
type EventSearchIndexInvalidateCacheEvent struct {
	events.AuditTopic
	ID string `json:"id"` // 要清除的 Event Search Index ID
}

// EventsInvalidateCacheEvent Events 缓存清除事件
type EventsInvalidateCacheEvent struct {
	events.AuditTopic
	ID string `json:"id"` // 要清除的 Event ID
}

// HashChainCheckpointsInvalidateCacheEvent HashChainCheckpoints 缓存清除事件
type HashChainCheckpointsInvalidateCacheEvent struct {
	events.AuditTopic
	ID string `json:"id"` // 要清除的 Hash Chain Checkpoint ID
}
