package commands

import (
	"github.com/google/uuid"
)

// CreateTenantSettingCmd 创建租户设置命令
type CreateTenantSettingCmd struct {
	TenantID           uuid.UUID
	EnforceMFA         bool
	AllowedEmailDomains []string
	SessionTTLMinutes  *int
	PasswordPolicy     map[string]interface{}
	LoginPolicy        map[string]interface{}
	MFAPolicy          map[string]interface{}
}

// UpdateTenantSettingCmd 更新租户设置命令
type UpdateTenantSettingCmd struct {
	TenantSettingID    uuid.UUID
	EnforceMFA         *bool
	AllowedEmailDomains []string
	SessionTTLMinutes  *int
	PasswordPolicy     map[string]interface{}
	LoginPolicy        map[string]interface{}
	MFAPolicy          map[string]interface{}
	UpdatedBy          *uuid.UUID
}

// DeleteTenantSettingCmd 删除租户设置命令
type DeleteTenantSettingCmd struct {
	TenantSettingID uuid.UUID
}
