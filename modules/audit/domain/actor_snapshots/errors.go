package actor_snapshots

import "errors"

var (
	ErrActorSnapshotNotFound      = errors.New("actor snapshot not found")
	ErrActorTypeRequired          = errors.New("actor type is required")
	ErrActorIDRequired            = errors.New("actor id is required")
	ErrSnapshotAtRequired         = errors.New("snapshot at is required")
	ErrInvalidActorType           = errors.New("invalid actor type")
	ErrActorSnapshotAlreadyExists = errors.New("actor snapshot already exists")
)
