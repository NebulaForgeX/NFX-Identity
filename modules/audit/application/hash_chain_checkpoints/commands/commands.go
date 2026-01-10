package commands

import (
	"github.com/google/uuid"
)

// CreateHashChainCheckpointCmd 创建哈希链检查点命令
type CreateHashChainCheckpointCmd struct {
	CheckpointID       string
	TenantID           *uuid.UUID
	PartitionDate      string
	CheckpointHash     string
	PrevCheckpointHash *string
	EventCount         int
	FirstEventID       *string
	LastEventID        *string
	CreatedBy          *string
}

// DeleteHashChainCheckpointCmd 删除哈希链检查点命令
type DeleteHashChainCheckpointCmd struct {
	HashChainCheckpointID uuid.UUID
}
