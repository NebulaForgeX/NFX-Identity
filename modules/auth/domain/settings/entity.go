package settings

import (
	"time"

	"github.com/google/uuid"
)

type Settings struct{ state State }

type State struct {
	ID                uuid.UUID
	Kind              string
	LoginNotification bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}

func NewFromState(st State) *Settings { return &Settings{state: st} }
func (s *Settings) ID() uuid.UUID                { return s.state.ID }
func (s *Settings) Kind() string                 { return s.state.Kind }
func (s *Settings) LoginNotification() bool      { return s.state.LoginNotification }
func (s *Settings) CreatedAt() time.Time         { return s.state.CreatedAt }
func (s *Settings) UpdatedAt() time.Time         { return s.state.UpdatedAt }
func (s *Settings) State() State                 { return s.state }
func (s *Settings) SetLoginNotification(v bool) {
	s.state.LoginNotification = v
	s.state.UpdatedAt = time.Now()
}
