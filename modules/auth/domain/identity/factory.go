package identity

import (
	"strings"
	"time"

	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type NewIdentityParams struct {
	AccountID        uuid.UUID
	IdentityProvider enums.AuthIdentityProvider
	ProviderSubject  string
	PasswordHash     *string
	Metadata         *datatypes.JSON
}

func NewIdentity(p NewIdentityParams) (*Identity, error) {
	if p.AccountID == uuid.Nil {
		return nil, authErr.ErrIdentityAccountIDInvalid
	}
	if err := validateIdentityProvider(p.IdentityProvider); err != nil {
		return nil, err
	}
	subject := strings.TrimSpace(p.ProviderSubject)
	if err := validateProviderSubject(subject); err != nil {
		return nil, err
	}
	if err := validatePasswordHashForProvider(p.IdentityProvider, p.PasswordHash); err != nil {
		return nil, err
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	var hashPtr *string
	if p.PasswordHash != nil {
		h := strings.TrimSpace(*p.PasswordHash)
		hashPtr = &h
	}
	return NewIdentityFromState(IdentityState{
		ID:               id,
		AccountID:        p.AccountID,
		IdentityProvider: p.IdentityProvider,
		ProviderSubject:  subject,
		PasswordHash:     hashPtr,
		Metadata:         p.Metadata,
		CreatedAt:        now,
		UpdatedAt:        now,
	}), nil
}
