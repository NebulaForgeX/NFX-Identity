package tenant_settings

import (
	"time"

	"github.com/google/uuid"
)

func (ts *TenantSetting) Update(enforceMFA *bool, allowedEmailDomains []string, sessionTTLMinutes *int, passwordPolicy, loginPolicy, mfaPolicy map[string]interface{}, updatedBy *uuid.UUID) error {
	if enforceMFA != nil {
		ts.state.EnforceMFA = *enforceMFA
	}
	if allowedEmailDomains != nil {
		ts.state.AllowedEmailDomains = allowedEmailDomains
	}
	if sessionTTLMinutes != nil {
		ts.state.SessionTTLMinutes = sessionTTLMinutes
	}
	if passwordPolicy != nil {
		ts.state.PasswordPolicy = passwordPolicy
	}
	if loginPolicy != nil {
		ts.state.LoginPolicy = loginPolicy
	}
	if mfaPolicy != nil {
		ts.state.MFAPolicy = mfaPolicy
	}

	ts.state.UpdatedBy = updatedBy
	ts.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ts *TenantSetting) UpdateEnforceMFA(enforceMFA bool, updatedBy *uuid.UUID) error {
	ts.state.EnforceMFA = enforceMFA
	ts.state.UpdatedBy = updatedBy
	ts.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ts *TenantSetting) UpdatePasswordPolicy(policy map[string]interface{}, updatedBy *uuid.UUID) error {
	if policy == nil {
		return nil
	}
	ts.state.PasswordPolicy = policy
	ts.state.UpdatedBy = updatedBy
	ts.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ts *TenantSetting) UpdateLoginPolicy(policy map[string]interface{}, updatedBy *uuid.UUID) error {
	if policy == nil {
		return nil
	}
	ts.state.LoginPolicy = policy
	ts.state.UpdatedBy = updatedBy
	ts.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ts *TenantSetting) UpdateMFAPolicy(policy map[string]interface{}, updatedBy *uuid.UUID) error {
	if policy == nil {
		return nil
	}
	ts.state.MFAPolicy = policy
	ts.state.UpdatedBy = updatedBy
	ts.state.UpdatedAt = time.Now().UTC()
	return nil
}
