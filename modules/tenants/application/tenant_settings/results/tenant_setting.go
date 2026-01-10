package results

import (
	"time"

	"nfxid/modules/tenants/domain/tenant_settings"

	"github.com/google/uuid"
)

type TenantSettingRO struct {
	ID                 uuid.UUID
	TenantID           uuid.UUID
	EnforceMFA         bool
	AllowedEmailDomains []string
	SessionTTLMinutes  *int
	PasswordPolicy     map[string]interface{}
	LoginPolicy        map[string]interface{}
	MFAPolicy          map[string]interface{}
	CreatedAt          time.Time
	UpdatedAt          time.Time
	UpdatedBy          *uuid.UUID
}

// TenantSettingMapper 将 Domain TenantSetting 转换为 Application TenantSettingRO
func TenantSettingMapper(ts *tenant_settings.TenantSetting) TenantSettingRO {
	if ts == nil {
		return TenantSettingRO{}
	}

	return TenantSettingRO{
		ID:                 ts.ID(),
		TenantID:           ts.TenantID(),
		EnforceMFA:         ts.EnforceMFA(),
		AllowedEmailDomains: ts.AllowedEmailDomains(),
		SessionTTLMinutes:  ts.SessionTTLMinutes(),
		PasswordPolicy:     ts.PasswordPolicy(),
		LoginPolicy:        ts.LoginPolicy(),
		MFAPolicy:          ts.MFAPolicy(),
		CreatedAt:            ts.CreatedAt(),
		UpdatedAt:         ts.UpdatedAt(),
		UpdatedBy:         ts.UpdatedBy(),
	}
}
