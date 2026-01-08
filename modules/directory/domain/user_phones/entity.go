package user_phones

import (
	"time"

	"github.com/google/uuid"
)

type UserPhone struct {
	state UserPhoneState
}

type UserPhoneState struct {
	ID                    uuid.UUID
	UserID                uuid.UUID
	Phone                 string
	CountryCode           *string
	IsPrimary             bool
	IsVerified            bool
	VerifiedAt            *time.Time
	VerificationCode      *string
	VerificationExpiresAt *time.Time
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             *time.Time
}

func (up *UserPhone) ID() uuid.UUID                    { return up.state.ID }
func (up *UserPhone) UserID() uuid.UUID                { return up.state.UserID }
func (up *UserPhone) Phone() string                    { return up.state.Phone }
func (up *UserPhone) CountryCode() *string             { return up.state.CountryCode }
func (up *UserPhone) IsPrimary() bool                  { return up.state.IsPrimary }
func (up *UserPhone) IsVerified() bool                 { return up.state.IsVerified }
func (up *UserPhone) VerifiedAt() *time.Time           { return up.state.VerifiedAt }
func (up *UserPhone) VerificationCode() *string        { return up.state.VerificationCode }
func (up *UserPhone) VerificationExpiresAt() *time.Time { return up.state.VerificationExpiresAt }
func (up *UserPhone) CreatedAt() time.Time             { return up.state.CreatedAt }
func (up *UserPhone) UpdatedAt() time.Time             { return up.state.UpdatedAt }
func (up *UserPhone) DeletedAt() *time.Time            { return up.state.DeletedAt }
