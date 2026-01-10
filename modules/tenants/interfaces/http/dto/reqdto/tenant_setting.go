package reqdto

import (
	tenantSettingAppCommands "nfxid/modules/tenants/application/tenant_settings/commands"

	"github.com/google/uuid"
)

type TenantSettingCreateRequestDTO struct {
	TenantID           uuid.UUID              `json:"tenant_id" validate:"required,uuid"`
	EnforceMFA         bool                   `json:"enforce_mfa"`
	AllowedEmailDomains []string              `json:"allowed_email_domains,omitempty"`
	SessionTTLMinutes  *int                   `json:"session_ttl_minutes,omitempty"`
	PasswordPolicy     map[string]interface{} `json:"password_policy,omitempty"`
	LoginPolicy        map[string]interface{} `json:"login_policy,omitempty"`
	MFAPolicy          map[string]interface{} `json:"mfa_policy,omitempty"`
}

type TenantSettingUpdateRequestDTO struct {
	ID                 uuid.UUID              `params:"id" validate:"required,uuid"`
	EnforceMFA         *bool                  `json:"enforce_mfa,omitempty"`
	AllowedEmailDomains []string              `json:"allowed_email_domains,omitempty"`
	SessionTTLMinutes  *int                   `json:"session_ttl_minutes,omitempty"`
	PasswordPolicy     map[string]interface{} `json:"password_policy,omitempty"`
	LoginPolicy        map[string]interface{} `json:"login_policy,omitempty"`
	MFAPolicy          map[string]interface{} `json:"mfa_policy,omitempty"`
	UpdatedBy          *uuid.UUID             `json:"updated_by,omitempty"`
}

func (r *TenantSettingCreateRequestDTO) ToCreateCmd() tenantSettingAppCommands.CreateTenantSettingCmd {
	return tenantSettingAppCommands.CreateTenantSettingCmd{
		TenantID:           r.TenantID,
		EnforceMFA:         r.EnforceMFA,
		AllowedEmailDomains: r.AllowedEmailDomains,
		SessionTTLMinutes:  r.SessionTTLMinutes,
		PasswordPolicy:     r.PasswordPolicy,
		LoginPolicy:        r.LoginPolicy,
		MFAPolicy:          r.MFAPolicy,
	}
}

func (r *TenantSettingUpdateRequestDTO) ToUpdateCmd() tenantSettingAppCommands.UpdateTenantSettingCmd {
	cmd := tenantSettingAppCommands.UpdateTenantSettingCmd{
		TenantSettingID:     r.ID,
		AllowedEmailDomains: r.AllowedEmailDomains,
		SessionTTLMinutes:  r.SessionTTLMinutes,
		PasswordPolicy:     r.PasswordPolicy,
		LoginPolicy:        r.LoginPolicy,
		MFAPolicy:          r.MFAPolicy,
		UpdatedBy:          r.UpdatedBy,
	}

	if r.EnforceMFA != nil {
		cmd.EnforceMFA = r.EnforceMFA
	}

	return cmd
}
