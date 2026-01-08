package client_credentials

import (
	"time"

	"github.com/google/uuid"
)

type CredentialStatus string

const (
	CredentialStatusActive  CredentialStatus = "active"
	CredentialStatusExpired CredentialStatus = "expired"
	CredentialStatusRevoked CredentialStatus = "revoked"
	CredentialStatusRotating CredentialStatus = "rotating"
)

type ClientCredential struct {
	state ClientCredentialState
}

type ClientCredentialState struct {
	ID          uuid.UUID
	AppID       uuid.UUID
	ClientID    string
	SecretHash  string
	HashAlg     string
	Status      CredentialStatus
	CreatedAt   time.Time
	RotatedAt   *time.Time
	ExpiresAt   *time.Time
	LastUsedAt  *time.Time
	CreatedBy   *uuid.UUID
	RevokedAt   *time.Time
	RevokedBy   *uuid.UUID
	RevokeReason *string
}

func (cc *ClientCredential) ID() uuid.UUID            { return cc.state.ID }
func (cc *ClientCredential) AppID() uuid.UUID         { return cc.state.AppID }
func (cc *ClientCredential) ClientID() string         { return cc.state.ClientID }
func (cc *ClientCredential) SecretHash() string       { return cc.state.SecretHash }
func (cc *ClientCredential) HashAlg() string          { return cc.state.HashAlg }
func (cc *ClientCredential) Status() CredentialStatus { return cc.state.Status }
func (cc *ClientCredential) CreatedAt() time.Time     { return cc.state.CreatedAt }
func (cc *ClientCredential) RotatedAt() *time.Time    { return cc.state.RotatedAt }
func (cc *ClientCredential) ExpiresAt() *time.Time    { return cc.state.ExpiresAt }
func (cc *ClientCredential) LastUsedAt() *time.Time   { return cc.state.LastUsedAt }
func (cc *ClientCredential) CreatedBy() *uuid.UUID    { return cc.state.CreatedBy }
func (cc *ClientCredential) RevokedAt() *time.Time    { return cc.state.RevokedAt }
func (cc *ClientCredential) RevokedBy() *uuid.UUID    { return cc.state.RevokedBy }
func (cc *ClientCredential) RevokeReason() *string    { return cc.state.RevokeReason }
