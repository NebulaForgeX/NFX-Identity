package hash_chain_checkpoints

import (
	auditErr "nfxidentity/errors/src/audit"
)

func (hcc *HashChainCheckpoint) Validate() error {
	if hcc.CheckpointID() == "" {
		return auditErr.ErrCheckpointIDRequired
	}
	if hcc.CheckpointHash() == "" {
		return auditErr.ErrCheckpointHashRequired
	}
	return nil
}
