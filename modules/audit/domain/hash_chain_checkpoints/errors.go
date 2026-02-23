package hash_chain_checkpoints

import "nfxid/pkgs/errx"

var (
	ErrHashChainCheckpointNotFound = errx.NotFound("HASH_CHAIN_CHECKPOINT_NOT_FOUND", "hash chain checkpoint not found")
	ErrCheckpointIDRequired        = errx.InvalidArg("CHECKPOINT_ID_REQUIRED", "checkpoint id is required")
	ErrPartitionDateRequired       = errx.InvalidArg("PARTITION_DATE_REQUIRED", "partition date is required")
	ErrCheckpointHashRequired      = errx.InvalidArg("CHECKPOINT_HASH_REQUIRED", "checkpoint hash is required")
	ErrCheckpointIDAlreadyExists   = errx.Conflict("CHECKPOINT_ID_ALREADY_EXISTS", "checkpoint id already exists")
)
