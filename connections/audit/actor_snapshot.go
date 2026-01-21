package audit

import (
	"context"
	"fmt"

	actorsnapshotpb "nfxid/protos/gen/audit/actor_snapshot"
)

// ActorSnapshotClient ActorSnapshot 客户端
type ActorSnapshotClient struct {
	client actorsnapshotpb.ActorSnapshotServiceClient
}

// NewActorSnapshotClient 创建 ActorSnapshot 客户端
func NewActorSnapshotClient(client actorsnapshotpb.ActorSnapshotServiceClient) *ActorSnapshotClient {
	return &ActorSnapshotClient{client: client}
}

// GetActorSnapshotByID 根据ID获取操作者快照
func (c *ActorSnapshotClient) GetActorSnapshotByID(ctx context.Context, id string) (*actorsnapshotpb.ActorSnapshot, error) {
	req := &actorsnapshotpb.GetActorSnapshotByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetActorSnapshotByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ActorSnapshot, nil
}

// GetActorSnapshotsByActor 根据操作者获取快照列表
func (c *ActorSnapshotClient) GetActorSnapshotsByActor(ctx context.Context, actorType actorsnapshotpb.AuditActorType, actorID string, limit *int32) ([]*actorsnapshotpb.ActorSnapshot, error) {
	req := &actorsnapshotpb.GetActorSnapshotsByActorRequest{
		ActorType: actorType,
		ActorId:   actorID,
		Limit:     limit,
	}

	resp, err := c.client.GetActorSnapshotsByActor(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ActorSnapshots, nil
}