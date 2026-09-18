package events

import (
	"time"

	"github.com/google/uuid"
)

// SignupSuccessEvent is published after signup records are created.
type SignupSuccessEvent struct {
	AuthTopic
	AccountID uuid.UUID `json:"account_id"`
	Email     string    `json:"email"`
	Lang      string    `json:"lang"`
}

// LoginSuccessEvent is published after the user selects a profile.
type LoginSuccessEvent struct {
	AuthTopic
	AccountID        uuid.UUID `json:"account_id"`
	ProfileID        uuid.UUID `json:"profile_id"`
	ProfileKind      string    `json:"profile_kind"`
	IdentityProvider string    `json:"identity_provider"`
	ProviderSubject  string    `json:"provider_subject"`
	LoginEmail       string    `json:"login_email"`
	LoginAt          time.Time `json:"login_at"`
}
