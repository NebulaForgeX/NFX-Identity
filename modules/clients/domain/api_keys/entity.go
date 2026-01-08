package api_keys

import (
	"time"

	"github.com/google/uuid"
)

type APIKeyStatus string

const (
	APIKeyStatusActive  APIKeyStatus = "active"
	APIKeyStatusRevoked APIKeyStatus = "revoked"
	APIKeyStatusExpired APIKeyStatus = "expired"
)

type APIKey struct {
	state APIKeyState
}

type APIKeyState struct {
	ID          uuid.UUID
	KeyID       string
	AppID       uuid.UUID
	KeyHash     string
	HashAlg     string
	Name        string
	Status      APIKeyStatus
	ExpiresAt   *time.Time
	CreatedAt   time.Time
	RevokedAt   *time.Time
	RevokedBy   *uuid.UUID
	RevokeReason *string
	LastUsedAt  *time.Time
	CreatedBy   *uuid.UUID
	Metadata    map[string]interface{}
}

func (ak *APIKey) ID() uuid.UUID               { return ak.state.ID }
func (ak *APIKey) KeyID() string               { return ak.state.KeyID }
func (ak *APIKey) AppID() uuid.UUID            { return ak.state.AppID }
func (ak *APIKey) KeyHash() string             { return ak.state.KeyHash }
func (ak *APIKey) HashAlg() string             { return ak.state.HashAlg }
func (ak *APIKey) Name() string                { return ak.state.Name }
func (ak *APIKey) Status() APIKeyStatus        { return ak.state.Status }
func (ak *APIKey) ExpiresAt() *time.Time       { return ak.state.ExpiresAt }
func (ak *APIKey) CreatedAt() time.Time        { return ak.state.CreatedAt }
func (ak *APIKey) RevokedAt() *time.Time       { return ak.state.RevokedAt }
func (ak *APIKey) RevokedBy() *uuid.UUID       { return ak.state.RevokedBy }
func (ak *APIKey) RevokeReason() *string       { return ak.state.RevokeReason }
func (ak *APIKey) LastUsedAt() *time.Time      { return ak.state.LastUsedAt }
func (ak *APIKey) CreatedBy() *uuid.UUID       { return ak.state.CreatedBy }
func (ak *APIKey) Metadata() map[string]interface{} { return ak.state.Metadata }
