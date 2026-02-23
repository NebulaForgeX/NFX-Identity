package actor_snapshots

import "nfxid/pkgs/errx"

var (
	ErrActorSnapshotNotFound     = errx.NotFound("ACTOR_SNAPSHOT_NOT_FOUND", "actor snapshot not found")
	ErrActorTypeRequired         = errx.InvalidArg("ACTOR_TYPE_REQUIRED", "actor type is required")
	ErrActorIDRequired           = errx.InvalidArg("ACTOR_ID_REQUIRED", "actor id is required")
	ErrSnapshotAtRequired        = errx.InvalidArg("SNAPSHOT_AT_REQUIRED", "snapshot at is required")
	ErrInvalidActorType          = errx.InvalidArg("INVALID_ACTOR_TYPE", "invalid actor type")
	ErrActorSnapshotAlreadyExists = errx.Conflict("ACTOR_SNAPSHOT_ALREADY_EXISTS", "actor snapshot already exists")
)
