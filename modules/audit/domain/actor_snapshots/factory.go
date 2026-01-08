package actor_snapshots

import (
	"time"

	"github.com/google/uuid"
)

type NewActorSnapshotParams struct {
	ActorType    ActorType
	ActorID      uuid.UUID
	DisplayName  *string
	Email        *string
	ClientName   *string
	TenantID     *uuid.UUID
	SnapshotAt   time.Time
	SnapshotData map[string]interface{}
}

func NewActorSnapshot(p NewActorSnapshotParams) (*ActorSnapshot, error) {
	if err := validateActorSnapshotParams(p); err != nil {
		return nil, err
	}

	snapshotAt := p.SnapshotAt
	if snapshotAt.IsZero() {
		snapshotAt = time.Now().UTC()
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewActorSnapshotFromState(ActorSnapshotState{
		ID:           id,
		ActorType:    p.ActorType,
		ActorID:      p.ActorID,
		DisplayName:  p.DisplayName,
		Email:        p.Email,
		ClientName:   p.ClientName,
		TenantID:     p.TenantID,
		SnapshotAt:   snapshotAt,
		SnapshotData: p.SnapshotData,
		CreatedAt:    now,
	}), nil
}

func NewActorSnapshotFromState(st ActorSnapshotState) *ActorSnapshot {
	return &ActorSnapshot{state: st}
}

func validateActorSnapshotParams(p NewActorSnapshotParams) error {
	if p.ActorType == "" {
		return ErrActorTypeRequired
	}
	validActorTypes := map[ActorType]struct{}{
		ActorTypeUser:    {},
		ActorTypeService: {},
		ActorTypeSystem:  {},
		ActorTypeAdmin:   {},
	}
	if _, ok := validActorTypes[p.ActorType]; !ok {
		return ErrInvalidActorType
	}
	if p.ActorID == uuid.Nil {
		return ErrActorIDRequired
	}
	return nil
}
