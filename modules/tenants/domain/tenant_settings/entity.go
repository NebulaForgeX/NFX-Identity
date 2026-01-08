package tenant_settings

import (
	"time"

	"github.com/google/uuid"
)

type TenantSetting struct {
	state TenantSettingState
}

type TenantSettingState struct {
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

func (ts *TenantSetting) ID() uuid.UUID                      { return ts.state.ID }
func (ts *TenantSetting) TenantID() uuid.UUID                 { return ts.state.TenantID }
func (ts *TenantSetting) EnforceMFA() bool                    { return ts.state.EnforceMFA }
func (ts *TenantSetting) AllowedEmailDomains() []string       { return ts.state.AllowedEmailDomains }
func (ts *TenantSetting) SessionTTLMinutes() *int             { return ts.state.SessionTTLMinutes }
func (ts *TenantSetting) PasswordPolicy() map[string]interface{} { return ts.state.PasswordPolicy }
func (ts *TenantSetting) LoginPolicy() map[string]interface{} { return ts.state.LoginPolicy }
func (ts *TenantSetting) MFAPolicy() map[string]interface{}   { return ts.state.MFAPolicy }
func (ts *TenantSetting) CreatedAt() time.Time                { return ts.state.CreatedAt }
func (ts *TenantSetting) UpdatedAt() time.Time                { return ts.state.UpdatedAt }
func (ts *TenantSetting) UpdatedBy() *uuid.UUID               { return ts.state.UpdatedBy }
