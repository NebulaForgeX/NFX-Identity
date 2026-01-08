package hash_chain_checkpoints

func (hcc *HashChainCheckpoint) Validate() error {
	if hcc.CheckpointID() == "" {
		return ErrCheckpointIDRequired
	}
	if hcc.CheckpointHash() == "" {
		return ErrCheckpointHashRequired
	}
	return nil
}
