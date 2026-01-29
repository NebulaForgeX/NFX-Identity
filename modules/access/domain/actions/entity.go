package actions

import (
	"time"

	"github.com/google/uuid"
)

type Action struct {
	state ActionState
}

type ActionState struct {
	ID          uuid.UUID
	Key         string
	Service     string
	Status      string
	Name        string
	Description *string
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (a *Action) ID() uuid.UUID         { return a.state.ID }
func (a *Action) Key() string           { return a.state.Key }
func (a *Action) Service() string       { return a.state.Service }
func (a *Action) Status() string        { return a.state.Status }
func (a *Action) Name() string          { return a.state.Name }
func (a *Action) Description() *string  { return a.state.Description }
func (a *Action) IsSystem() bool        { return a.state.IsSystem }
func (a *Action) CreatedAt() time.Time  { return a.state.CreatedAt }
func (a *Action) UpdatedAt() time.Time  { return a.state.UpdatedAt }
func (a *Action) DeletedAt() *time.Time { return a.state.DeletedAt }
