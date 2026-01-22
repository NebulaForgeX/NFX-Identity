package user_credentials

import (
	"time"

	"github.com/google/uuid"
)

type CredentialType string

const (
	CredentialTypePassword  CredentialType = "password"
	CredentialTypePasskey   CredentialType = "passkey"
	CredentialTypeOauthLink CredentialType = "oauth_link"
	CredentialTypeSaml      CredentialType = "saml"
	CredentialTypeLdap      CredentialType = "ldap"
)

type CredentialStatus string

const (
	CredentialStatusActive   CredentialStatus = "active"
	CredentialStatusDisabled CredentialStatus = "disabled"
	CredentialStatusExpired  CredentialStatus = "expired"
)

type UserCredential struct {
	state UserCredentialState
}

type UserCredentialState struct {
	ID                  uuid.UUID
	UserID              uuid.UUID
	CredentialType      CredentialType
	PasswordHash        *string
	HashAlg             *string
	HashParams          map[string]interface{}
	PasswordUpdatedAt   *time.Time
	LastSuccessLoginAt  *time.Time
	Status              CredentialStatus
	MustChangePassword  bool
	Version             int
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time
}

func (uc *UserCredential) ID() uuid.UUID                    { return uc.state.ID }
func (uc *UserCredential) UserID() uuid.UUID                { return uc.state.UserID }
func (uc *UserCredential) CredentialType() CredentialType   { return uc.state.CredentialType }
func (uc *UserCredential) PasswordHash() *string            { return uc.state.PasswordHash }
func (uc *UserCredential) HashAlg() *string                 { return uc.state.HashAlg }
func (uc *UserCredential) HashParams() map[string]interface{} { return uc.state.HashParams }
func (uc *UserCredential) PasswordUpdatedAt() *time.Time    { return uc.state.PasswordUpdatedAt }
func (uc *UserCredential) LastSuccessLoginAt() *time.Time   { return uc.state.LastSuccessLoginAt }
func (uc *UserCredential) Status() CredentialStatus         { return uc.state.Status }
func (uc *UserCredential) MustChangePassword() bool         { return uc.state.MustChangePassword }
func (uc *UserCredential) Version() int                     { return uc.state.Version }
func (uc *UserCredential) CreatedAt() time.Time             { return uc.state.CreatedAt }
func (uc *UserCredential) UpdatedAt() time.Time             { return uc.state.UpdatedAt }
func (uc *UserCredential) DeletedAt() *time.Time            { return uc.state.DeletedAt }
