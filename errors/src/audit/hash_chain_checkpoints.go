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

/*
!HASH_CHAIN_CHECKPOINT_NOT_FOUND
*en<hash chain checkpoint not found>
*zh<哈希链检查点不存在>
*fr<point de contrôle de chaîne de hachage introuvable>

!CHECKPOINT_ID_REQUIRED
*en<checkpoint id required>
*zh<检查点 ID 为必填>
*fr<id de point de contrôle requis>

!PARTITION_DATE_REQUIRED
*en<partition date required>
*zh<分区日期为必填>
*fr<date de partition requise>

!CHECKPOINT_HASH_REQUIRED
*en<checkpoint hash required>
*zh<检查点哈希为必填>
*fr<hachage du point de contrôle requis>

!CHECKPOINT_ID_ALREADY_EXISTS
*en<checkpoint id already exists>
*zh<检查点 ID 已存在>
*fr<id de point de contrôle existe déjà>

*/
