package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	state UserState
	// events []eventbus.TypedEvent // Reserved for future event sourcing
}

type UserState struct {
	ID          uuid.UUID
	Editable    UserEditable
	Status      string // pending, active, deactive
	IsVerified  bool
	LastLoginAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type UserEditable struct {
	Username string
	Email    string
	Phone    string
	Password string
}

func (u *User) ID() uuid.UUID              { return u.state.ID }
func (u *User) Editable() UserEditable     { return u.state.Editable }
func (u *User) Status() string             { return u.state.Status }
func (u *User) IsVerified() bool            { return u.state.IsVerified }
func (u *User) LastLoginAt() *time.Time    { return u.state.LastLoginAt }
func (u *User) CreatedAt() time.Time       { return u.state.CreatedAt }
func (u *User) UpdatedAt() time.Time       { return u.state.UpdatedAt }
func (u *User) DeletedAt() *time.Time      { return u.state.DeletedAt }

func (u *User) IsActive() bool {
	return u.Status() == "active" && u.DeletedAt() == nil
}

