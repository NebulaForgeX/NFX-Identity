package actor_snapshots

import (
	"time"

	"github.com/google/uuid"
)

type ActorType string

const (
	ActorTypeUser    ActorType = "user"
	ActorTypeService ActorType = "service"
	ActorTypeSystem  ActorType = "system"
	ActorTypeAdmin   ActorType = "admin"
)

type ActorSnapshot struct {
	state ActorSnapshotState
}

type ActorSnapshotState struct {
	ID           uuid.UUID
	ActorType    ActorType
	ActorID      uuid.UUID
	DisplayName  *string
	Email        *string
	ClientName   *string
	TenantID     *uuid.UUID
	SnapshotAt   time.Time
	SnapshotData map[string]interface{}
	CreatedAt    time.Time
}

func (as *ActorSnapshot) ID() uuid.UUID                  { return as.state.ID }
func (as *ActorSnapshot) ActorType() ActorType           { return as.state.ActorType }
func (as *ActorSnapshot) ActorID() uuid.UUID             { return as.state.ActorID }
func (as *ActorSnapshot) DisplayName() *string           { return as.state.DisplayName }
func (as *ActorSnapshot) Email() *string                 { return as.state.Email }
func (as *ActorSnapshot) ClientName() *string            { return as.state.ClientName }
func (as *ActorSnapshot) TenantID() *uuid.UUID           { return as.state.TenantID }
func (as *ActorSnapshot) SnapshotAt() time.Time          { return as.state.SnapshotAt }
func (as *ActorSnapshot) SnapshotData() map[string]interface{} { return as.state.SnapshotData }
func (as *ActorSnapshot) CreatedAt() time.Time           { return as.state.CreatedAt }
