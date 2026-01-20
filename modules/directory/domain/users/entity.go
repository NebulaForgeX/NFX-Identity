package users

import (
	"time"

	"github.com/google/uuid"
)

type UserStatus string

const (
	UserStatusPending  UserStatus = "pending"
	UserStatusActive   UserStatus = "active"
	UserStatusDeactive UserStatus = "deactive"
)

type User struct {
	state UserState
}

type UserState struct {
	ID          uuid.UUID
	Username    string
	Status      UserStatus
	IsVerified  bool
	LastLoginAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (u *User) ID() uuid.UUID           { return u.state.ID }
func (u *User) Username() string        { return u.state.Username }
func (u *User) Status() UserStatus      { return u.state.Status }
func (u *User) IsVerified() bool        { return u.state.IsVerified }
func (u *User) LastLoginAt() *time.Time { return u.state.LastLoginAt }
func (u *User) CreatedAt() time.Time    { return u.state.CreatedAt }
func (u *User) UpdatedAt() time.Time    { return u.state.UpdatedAt }
func (u *User) DeletedAt() *time.Time   { return u.state.DeletedAt }
