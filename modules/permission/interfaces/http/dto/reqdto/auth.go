package reqdto

import (
	authCommands "nfxid/modules/permission/application/auth/commands"
)

type AuthLoginRequestDTO struct {
	Type       string `json:"type" validate:"required,oneof=password code"` // "password" or "code"
	Identifier string `json:"identifier" validate:"required"`                // username, email 或 phone
	Password   string `json:"password,omitempty"`                            // 密码（当 type=password 时）
	Code       string `json:"code,omitempty"`                                // 验证码（当 type=code 时）
}

func (dto *AuthLoginRequestDTO) ToLoginCmd() authCommands.LoginCmd {
	return authCommands.LoginCmd{
		Type:       dto.Type,
		Identifier: dto.Identifier,
		Password:   dto.Password,
		Code:       dto.Code,
	}
}

