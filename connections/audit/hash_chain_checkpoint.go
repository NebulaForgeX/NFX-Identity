package audit

import (
	"context"
	"fmt"

	hashchaincheckpointpb "nfxid/protos/gen/audit/hash_chain_checkpoint"
)

// HashChainCheckpointClient HashChainCheckpoint 客户端
type HashChainCheckpointClient struct {
	client hashchaincheckpointpb.HashChainCheckpointServiceClient
}

// NewHashChainCheckpointClient 创建 HashChainCheckpoint 客户端
func NewHashChainCheckpointClient(client hashchaincheckpointpb.HashChainCheckpointServiceClient) *HashChainCheckpointClient {
	return &HashChainCheckpointClient{client: client}
}

// GetHashChainCheckpointByID 根据ID获取检查点
func (c *HashChainCheckpointClient) GetHashChainCheckpointByID(ctx context.Context, id string) (*hashchaincheckpointpb.HashChainCheckpoint, error) {
	req := &hashchaincheckpointpb.GetHashChainCheckpointByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetHashChainCheckpointByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.HashChainCheckpoint, nil
}

// GetHashChainCheckpointByCheckpointID 根据检查点ID获取检查点
func (c *HashChainCheckpointClient) GetHashChainCheckpointByCheckpointID(ctx context.Context, checkpointID string) (*hashchaincheckpointpb.HashChainCheckpoint, error) {
	req := &hashchaincheckpointpb.GetHashChainCheckpointByCheckpointIDRequest{
		CheckpointId: checkpointID,
	}

	resp, err := c.client.GetHashChainCheckpointByCheckpointID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.HashChainCheckpoint, nil
}