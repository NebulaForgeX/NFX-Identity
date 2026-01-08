package hash_chain_checkpoints

import "errors"

var (
	ErrHashChainCheckpointNotFound      = errors.New("hash chain checkpoint not found")
	ErrCheckpointIDRequired             = errors.New("checkpoint id is required")
	ErrPartitionDateRequired            = errors.New("partition date is required")
	ErrCheckpointHashRequired           = errors.New("checkpoint hash is required")
	ErrCheckpointIDAlreadyExists        = errors.New("checkpoint id already exists")
)
