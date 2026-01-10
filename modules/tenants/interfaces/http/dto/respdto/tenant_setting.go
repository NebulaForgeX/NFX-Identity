package respdto

import (
	"time"

	tenantSettingAppResult "nfxid/modules/tenants/application/tenant_settings/results"

	"github.com/google/uuid"
)

type TenantSettingDTO struct {
	ID                 uuid.UUID              `json:"id"`
	TenantID           uuid.UUID              `json:"tenant_id"`
	EnforceMFA         bool                   `json:"enforce_mfa"`
	AllowedEmailDomains []string              `json:"allowed_email_domains,omitempty"`
	SessionTTLMinutes  *int                   `json:"session_ttl_minutes,omitempty"`
	PasswordPolicy     map[string]interface{} `json:"password_policy,omitempty"`
	LoginPolicy        map[string]interface{} `json:"login_policy,omitempty"`
	MFAPolicy          map[string]interface{} `json:"mfa_policy,omitempty"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
	UpdatedBy          *uuid.UUID             `json:"updated_by,omitempty"`
}

// TenantSettingROToDTO converts application TenantSettingRO to response DTO
func TenantSettingROToDTO(v *tenantSettingAppResult.TenantSettingRO) *TenantSettingDTO {
	if v == nil {
		return nil
	}

	return &TenantSettingDTO{
		ID:                 v.ID,
		TenantID:           v.TenantID,
		EnforceMFA:         v.EnforceMFA,
		AllowedEmailDomains: v.AllowedEmailDomains,
		SessionTTLMinutes:  v.SessionTTLMinutes,
		PasswordPolicy:     v.PasswordPolicy,
		LoginPolicy:        v.LoginPolicy,
		MFAPolicy:          v.MFAPolicy,
		CreatedAt:          v.CreatedAt,
		UpdatedAt:          v.UpdatedAt,
		UpdatedBy:          v.UpdatedBy,
	}
}

// TenantSettingListROToDTO converts list of TenantSettingRO to DTOs
func TenantSettingListROToDTO(results []tenantSettingAppResult.TenantSettingRO) []TenantSettingDTO {
	dtos := make([]TenantSettingDTO, len(results))
	for i, v := range results {
		if dto := TenantSettingROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
