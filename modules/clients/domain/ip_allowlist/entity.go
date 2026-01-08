package ip_allowlist

import (
	"time"

	"github.com/google/uuid"
)

type AllowlistStatus string

const (
	AllowlistStatusActive   AllowlistStatus = "active"
	AllowlistStatusDisabled AllowlistStatus = "disabled"
	AllowlistStatusRevoked  AllowlistStatus = "revoked"
)

type IPAllowlist struct {
	state IPAllowlistState
}

type IPAllowlistState struct {
	ID          uuid.UUID
	RuleID      string
	AppID       uuid.UUID
	CIDR        string
	Description *string
	Status      AllowlistStatus
	CreatedAt   time.Time
	CreatedBy   *uuid.UUID
	UpdatedAt   time.Time
	UpdatedBy   *uuid.UUID
	RevokedAt   *time.Time
	RevokedBy   *uuid.UUID
	RevokeReason *string
}

func (ip *IPAllowlist) ID() uuid.UUID              { return ip.state.ID }
func (ip *IPAllowlist) RuleID() string              { return ip.state.RuleID }
func (ip *IPAllowlist) AppID() uuid.UUID            { return ip.state.AppID }
func (ip *IPAllowlist) CIDR() string                { return ip.state.CIDR }
func (ip *IPAllowlist) Description() *string        { return ip.state.Description }
func (ip *IPAllowlist) Status() AllowlistStatus     { return ip.state.Status }
func (ip *IPAllowlist) CreatedAt() time.Time        { return ip.state.CreatedAt }
func (ip *IPAllowlist) CreatedBy() *uuid.UUID       { return ip.state.CreatedBy }
func (ip *IPAllowlist) UpdatedAt() time.Time        { return ip.state.UpdatedAt }
func (ip *IPAllowlist) UpdatedBy() *uuid.UUID       { return ip.state.UpdatedBy }
func (ip *IPAllowlist) RevokedAt() *time.Time       { return ip.state.RevokedAt }
func (ip *IPAllowlist) RevokedBy() *uuid.UUID       { return ip.state.RevokedBy }
func (ip *IPAllowlist) RevokeReason() *string       { return ip.state.RevokeReason }
