package audit

import "nfxid/pkgs/errx"

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
