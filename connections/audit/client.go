package audit

import (
	actorsnapshotpb "nfxidentity/protos/gen/audit/actor_snapshot"
	eventpb "nfxidentity/protos/gen/audit/event"
	eventretentionpolicypb "nfxidentity/protos/gen/audit/event_retention_policy"
	eventsearchindexpb "nfxidentity/protos/gen/audit/event_search_index"
	hashchaincheckpointpb "nfxidentity/protos/gen/audit/hash_chain_checkpoint"
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
		HashChainCheckpoint:  NewHashChainCheckpointClient(hashChainCheckpointClient),
		EventRetentionPolicy: NewEventRetentionPolicyClient(eventRetentionPolicyClient),
	}
}
