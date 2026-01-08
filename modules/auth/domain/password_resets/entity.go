package password_resets

import (
	"time"

	"github.com/google/uuid"
)

type ResetDelivery string

const (
	ResetDeliveryEmail ResetDelivery = "email"
	ResetDeliverySMS   ResetDelivery = "sms"
)

type ResetStatus string

const (
	ResetStatusIssued  ResetStatus = "issued"
	ResetStatusUsed    ResetStatus = "used"
	ResetStatusExpired ResetStatus = "expired"
	ResetStatusRevoked ResetStatus = "revoked"
)

type PasswordReset struct {
	state PasswordResetState
}

type PasswordResetState struct {
	ID          uuid.UUID
	ResetID     string
	TenantID    uuid.UUID
	UserID      uuid.UUID
	Delivery    ResetDelivery
	CodeHash    string
	ExpiresAt   time.Time
	UsedAt      *time.Time
	RequestedIP *string
	UAHash      *string
	AttemptCount int
	Status      ResetStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (pr *PasswordReset) ID() uuid.UUID            { return pr.state.ID }
func (pr *PasswordReset) ResetID() string          { return pr.state.ResetID }
func (pr *PasswordReset) TenantID() uuid.UUID      { return pr.state.TenantID }
func (pr *PasswordReset) UserID() uuid.UUID        { return pr.state.UserID }
func (pr *PasswordReset) Delivery() ResetDelivery  { return pr.state.Delivery }
func (pr *PasswordReset) CodeHash() string         { return pr.state.CodeHash }
func (pr *PasswordReset) ExpiresAt() time.Time     { return pr.state.ExpiresAt }
func (pr *PasswordReset) UsedAt() *time.Time       { return pr.state.UsedAt }
func (pr *PasswordReset) RequestedIP() *string     { return pr.state.RequestedIP }
func (pr *PasswordReset) UAHash() *string          { return pr.state.UAHash }
func (pr *PasswordReset) AttemptCount() int        { return pr.state.AttemptCount }
func (pr *PasswordReset) Status() ResetStatus      { return pr.state.Status }
func (pr *PasswordReset) CreatedAt() time.Time     { return pr.state.CreatedAt }
func (pr *PasswordReset) UpdatedAt() time.Time     { return pr.state.UpdatedAt }
