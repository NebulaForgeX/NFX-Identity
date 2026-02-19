package audit

import (
	actorsnapshotpb "nfxid/protos/gen/audit/actor_snapshot"
	eventpb "nfxid/protos/gen/audit/event"
	eventretentionpolicypb "nfxid/protos/gen/audit/event_retention_policy"
	eventsearchindexpb "nfxid/protos/gen/audit/event_search_index"
	hashchaincheckpointpb "nfxid/protos/gen/audit/hash_chain_checkpoint"
)

// Client Audit 服务客户端
type Client struct {
	Event                *EventClient
	EventSearchIndex     *EventSearchIndexClient
	ActorSnapshot        *ActorSnapshotClient
	HashChainCheckpoint  *HashChainCheckpointClient
	EventRetentionPolicy *EventRetentionPolicyClient
}

// NewClient 创建 Audit 客户端
func NewClient(
	eventClient eventpb.EventServiceClient,
	eventSearchIndexClient eventsearchindexpb.EventSearchIndexServiceClient,
	actorSnapshotClient actorsnapshotpb.ActorSnapshotServiceClient,
	hashChainCheckpointClient hashchaincheckpointpb.HashChainCheckpointServiceClient,
	eventRetentionPolicyClient eventretentionpolicypb.EventRetentionPolicyServiceClient,
) *Client {
	return &Client{
		Event:                NewEventClient(eventClient),
		EventSearchIndex:     NewEventSearchIndexClient(eventSearchIndexClient),
		ActorSnapshot:        NewActorSnapshotClient(actorSnapshotClient),
		HashChainCheckpoint:   NewHashChainCheckpointClient(hashChainCheckpointClient),
		EventRetentionPolicy: NewEventRetentionPolicyClient(eventRetentionPolicyClient),
	}
}
