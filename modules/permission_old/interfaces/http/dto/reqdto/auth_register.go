package reqdto

import (
	authCommands "nfxid/modules/permission/application/auth/commands"
)

type AuthRegisterRequestDTO struct {
	Email            string `json:"email" validate:"required,email"`
	VerificationCode string `json:"verification_code" validate:"required,min=6,max=6"`
	AuthorizationCode string `json:"authorization_code" validate:"required,min=1"`
	Password         string `json:"password" validate:"required,min=6"`
}

func (dto *AuthRegisterRequestDTO) ToRegisterCmd() authCommands.RegisterCmd {
	return authCommands.RegisterCmd{
		Email:            dto.Email,
		VerificationCode: dto.VerificationCode,
		AuthorizationCode: dto.AuthorizationCode,
		Password:         dto.Password,
	}
}
