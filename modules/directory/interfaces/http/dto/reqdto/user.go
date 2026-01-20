package reqdto

import (
	userAppCommands "nfxid/modules/directory/application/users/commands"
	userDomain "nfxid/modules/directory/domain/users"

	"github.com/google/uuid"
)

type UserCreateRequestDTO struct {
	Username   string    `json:"username" validate:"required"`
	Status     string    `json:"status,omitempty"`
	IsVerified bool      `json:"is_verified,omitempty"`
}

type UserUpdateStatusRequestDTO struct {
	ID     uuid.UUID `params:"id" validate:"required,uuid"`
	Status string    `json:"status" validate:"required"`
}

type UserUpdateUsernameRequestDTO struct {
	ID       uuid.UUID `params:"id" validate:"required,uuid"`
	Username string    `json:"username" validate:"required"`
}

type UserVerifyRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type UserByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type UserByUsernameRequestDTO struct {
	Username string `params:"username" validate:"required"`
}

func (r *UserCreateRequestDTO) ToCreateCmd() userAppCommands.CreateUserCmd {
	cmd := userAppCommands.CreateUserCmd{
		Username:   r.Username,
		IsVerified: r.IsVerified,
	}

	// Parse status
	if r.Status != "" {
		cmd.Status = userDomain.UserStatus(r.Status)
	} else {
		cmd.Status = userDomain.UserStatusActive
	}

	return cmd
}

func (r *UserUpdateStatusRequestDTO) ToUpdateStatusCmd() userAppCommands.UpdateUserStatusCmd {
	return userAppCommands.UpdateUserStatusCmd{
		UserID: r.ID,
		Status: userDomain.UserStatus(r.Status),
	}
}

func (r *UserUpdateUsernameRequestDTO) ToUpdateUsernameCmd() userAppCommands.UpdateUsernameCmd {
	return userAppCommands.UpdateUsernameCmd{
		UserID:   r.ID,
		Username: r.Username,
	}
}

func (r *UserVerifyRequestDTO) ToVerifyCmd() userAppCommands.VerifyUserCmd {
	return userAppCommands.VerifyUserCmd{
		UserID: r.ID,
	}
}
