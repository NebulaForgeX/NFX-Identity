package audit

import "nfxid/pkgs/errx"

const (
	CodeHashChainCheckpointNotFound = "HASH_CHAIN_CHECKPOINT_NOT_FOUND"
	CodeCheckpointIDRequired        = "CHECKPOINT_ID_REQUIRED"
	CodePartitionDateRequired       = "PARTITION_DATE_REQUIRED"
	CodeCheckpointHashRequired      = "CHECKPOINT_HASH_REQUIRED"
	CodeCheckpointIDAlreadyExists   = "CHECKPOINT_ID_ALREADY_EXISTS"
)

var (
	ErrHashChainCheckpointNotFound = errx.NotFound(CodeHashChainCheckpointNotFound, "hash chain checkpoint not found")
	ErrCheckpointIDRequired        = errx.InvalidArg(CodeCheckpointIDRequired, "checkpoint id is required")
	ErrPartitionDateRequired       = errx.InvalidArg(CodePartitionDateRequired, "partition date is required")
	ErrCheckpointHashRequired      = errx.InvalidArg(CodeCheckpointHashRequired, "checkpoint hash is required")
	ErrCheckpointIDAlreadyExists   = errx.Conflict(CodeCheckpointIDAlreadyExists, "checkpoint id already exists")
)
