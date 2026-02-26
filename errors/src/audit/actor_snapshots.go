package audit

import "nfxidentity/pkgs/errx"

const (
	CodeActorSnapshotNotFound      = "ACTOR_SNAPSHOT_NOT_FOUND"
	CodeSnapshotAtRequired         = "SNAPSHOT_AT_REQUIRED"
	CodeActorSnapshotAlreadyExists = "ACTOR_SNAPSHOT_ALREADY_EXISTS"
)

var (
	ErrActorSnapshotNotFound      = errx.NotFound(CodeActorSnapshotNotFound, "actor snapshot not found")
	ErrSnapshotAtRequired         = errx.InvalidArg(CodeSnapshotAtRequired, "snapshot at is required")
	ErrActorSnapshotAlreadyExists = errx.Conflict(CodeActorSnapshotAlreadyExists, "actor snapshot already exists")
)

/*
!ACTOR_SNAPSHOT_NOT_FOUND
*en<actor snapshot not found>
*zh<操作者快照不存在>
*fr<instantané d'acteur introuvable>

!SNAPSHOT_AT_REQUIRED
*en<snapshot at required>
*zh<快照时间为必填>
*fr<date d'instantané requise>

!ACTOR_SNAPSHOT_ALREADY_EXISTS
*en<actor snapshot already exists>
*zh<操作者快照已存在>
*fr<instantané d'acteur existe déjà>

*/
