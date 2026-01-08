package mfa_factors

import (
	"time"

	"github.com/google/uuid"
)

type MFAType string

const (
	MFATypeTOTP      MFAType = "totp"
	MFATypeSMS       MFAType = "sms"
	MFATypeEmail     MFAType = "email"
	MFATypeWebAuthn  MFAType = "webauthn"
	MFATypeBackupCode MFAType = "backup_code"
)

type MFAFactor struct {
	state MFAFactorState
}

type MFAFactorState struct {
	ID               uuid.UUID
	FactorID         string
	TenantID         uuid.UUID
	UserID           uuid.UUID
	Type             MFAType
	SecretEncrypted  *string
	Phone            *string
	Email            *string
	Name             *string
	Enabled          bool
	CreatedAt        time.Time
	LastUsedAt       *time.Time
	RecoveryCodesHash *string
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}

func (mf *MFAFactor) ID() uuid.UUID            { return mf.state.ID }
func (mf *MFAFactor) FactorID() string         { return mf.state.FactorID }
func (mf *MFAFactor) TenantID() uuid.UUID      { return mf.state.TenantID }
func (mf *MFAFactor) UserID() uuid.UUID        { return mf.state.UserID }
func (mf *MFAFactor) Type() MFAType            { return mf.state.Type }
func (mf *MFAFactor) SecretEncrypted() *string { return mf.state.SecretEncrypted }
func (mf *MFAFactor) Phone() *string           { return mf.state.Phone }
func (mf *MFAFactor) Email() *string           { return mf.state.Email }
func (mf *MFAFactor) Name() *string            { return mf.state.Name }
func (mf *MFAFactor) Enabled() bool            { return mf.state.Enabled }
func (mf *MFAFactor) CreatedAt() time.Time     { return mf.state.CreatedAt }
func (mf *MFAFactor) LastUsedAt() *time.Time   { return mf.state.LastUsedAt }
func (mf *MFAFactor) RecoveryCodesHash() *string { return mf.state.RecoveryCodesHash }
func (mf *MFAFactor) UpdatedAt() time.Time     { return mf.state.UpdatedAt }
func (mf *MFAFactor) DeletedAt() *time.Time    { return mf.state.DeletedAt }
