package user_emails

import (
	"time"

	"github.com/google/uuid"
)

type UserEmail struct {
	state UserEmailState
}

type UserEmailState struct {
	ID                uuid.UUID
	UserID            uuid.UUID
	Email             string
	IsPrimary         bool
	IsVerified        bool
	VerifiedAt        *time.Time
	VerificationToken *string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}

func (ue *UserEmail) ID() uuid.UUID              { return ue.state.ID }
func (ue *UserEmail) UserID() uuid.UUID          { return ue.state.UserID }
func (ue *UserEmail) Email() string              { return ue.state.Email }
func (ue *UserEmail) IsPrimary() bool            { return ue.state.IsPrimary }
func (ue *UserEmail) IsVerified() bool           { return ue.state.IsVerified }
func (ue *UserEmail) VerifiedAt() *time.Time     { return ue.state.VerifiedAt }
func (ue *UserEmail) VerificationToken() *string { return ue.state.VerificationToken }
func (ue *UserEmail) CreatedAt() time.Time       { return ue.state.CreatedAt }
func (ue *UserEmail) UpdatedAt() time.Time       { return ue.state.UpdatedAt }
func (ue *UserEmail) DeletedAt() *time.Time      { return ue.state.DeletedAt }
