package tenant_settings

import (
	"time"

	"github.com/google/uuid"
)

type NewTenantSettingParams struct {
	TenantID           uuid.UUID
	EnforceMFA         bool
	AllowedEmailDomains []string
	SessionTTLMinutes  *int
	PasswordPolicy     map[string]interface{}
	LoginPolicy        map[string]interface{}
	MFAPolicy          map[string]interface{}
}

func NewTenantSetting(p NewTenantSettingParams) (*TenantSetting, error) {
	if err := validateTenantSettingParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewTenantSettingFromState(TenantSettingState{
		ID:                 id,
		TenantID:           p.TenantID,
		EnforceMFA:         p.EnforceMFA,
		AllowedEmailDomains: p.AllowedEmailDomains,
		SessionTTLMinutes:  p.SessionTTLMinutes,
		PasswordPolicy:     p.PasswordPolicy,
		LoginPolicy:        p.LoginPolicy,
		MFAPolicy:          p.MFAPolicy,
		CreatedAt:          now,
		UpdatedAt:          now,
	}), nil
}

func NewTenantSettingFromState(st TenantSettingState) *TenantSetting {
	return &TenantSetting{state: st}
}

func validateTenantSettingParams(p NewTenantSettingParams) error {
	if p.TenantID == uuid.Nil {
		return ErrTenantIDRequired
	}
	return nil
}
