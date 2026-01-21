package reqdto

import (
	bootstrapCommands "nfxid/modules/system/application/bootstrap/commands"
	systemStateCommands "nfxid/modules/system/application/system_state/commands"

	"github.com/google/uuid"
)

// SystemStateInitializeRequestDTO 系统初始化请求 DTO
type SystemStateInitializeRequestDTO struct {
	Version       string  `json:"version" validate:"omitempty"`
	AdminUsername string  `json:"admin_username" validate:"required"`
	AdminPassword string  `json:"admin_password" validate:"required,min=8"`
	AdminEmail    *string `json:"admin_email,omitempty" validate:"omitempty,email"`
	AdminPhone    *string `json:"admin_phone,omitempty"`
}

// ToBootstrapInitCmd 转换为 BootstrapInitCmd
func (r *SystemStateInitializeRequestDTO) ToBootstrapInitCmd() bootstrapCommands.BootstrapInitCmd {
	version := r.Version
	if version == "" {
		version = "1.0.0" // 默认版本
	}

	return bootstrapCommands.BootstrapInitCmd{
		Version:       version,
		AdminUsername: r.AdminUsername,
		AdminPassword: r.AdminPassword,
		AdminEmail:    r.AdminEmail,
		AdminPhone:    r.AdminPhone,
	}
}

// SystemStateResetRequestDTO 系统重置请求 DTO
type SystemStateResetRequestDTO struct {
	ResetBy uuid.UUID `json:"reset_by" validate:"required,uuid"`
}

// ToResetSystemCmd 转换为 ResetSystemCmd
func (r *SystemStateResetRequestDTO) ToResetSystemCmd() systemStateCommands.ResetSystemCmd {
	return systemStateCommands.ResetSystemCmd{
		ResetBy: r.ResetBy,
	}
}

// SystemStateByIDRequestDTO 根据 ID 获取系统状态请求 DTO
type SystemStateByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}
