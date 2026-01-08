package hash_chain_checkpoints

import (
	"time"

	"github.com/google/uuid"
)

type HashChainCheckpoint struct {
	state HashChainCheckpointState
}

type HashChainCheckpointState struct {
	ID                  uuid.UUID
	CheckpointID        string
	TenantID            *uuid.UUID
	PartitionDate       time.Time
	CheckpointHash      string
	PrevCheckpointHash  *string
	EventCount          int
	FirstEventID        *string
	LastEventID         *string
	CreatedAt           time.Time
	CreatedBy           *string
}

func (hcc *HashChainCheckpoint) ID() uuid.UUID             { return hcc.state.ID }
func (hcc *HashChainCheckpoint) CheckpointID() string      { return hcc.state.CheckpointID }
func (hcc *HashChainCheckpoint) TenantID() *uuid.UUID      { return hcc.state.TenantID }
func (hcc *HashChainCheckpoint) PartitionDate() time.Time  { return hcc.state.PartitionDate }
func (hcc *HashChainCheckpoint) CheckpointHash() string    { return hcc.state.CheckpointHash }
func (hcc *HashChainCheckpoint) PrevCheckpointHash() *string { return hcc.state.PrevCheckpointHash }
func (hcc *HashChainCheckpoint) EventCount() int           { return hcc.state.EventCount }
func (hcc *HashChainCheckpoint) FirstEventID() *string     { return hcc.state.FirstEventID }
func (hcc *HashChainCheckpoint) LastEventID() *string      { return hcc.state.LastEventID }
func (hcc *HashChainCheckpoint) CreatedAt() time.Time      { return hcc.state.CreatedAt }
func (hcc *HashChainCheckpoint) CreatedBy() *string        { return hcc.state.CreatedBy }
