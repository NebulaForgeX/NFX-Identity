package respdto

import (
	"time"

	userAppCommands "nfxid/modules/auth/application/user/commands"
	userAppViews "nfxid/modules/auth/application/user/views"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID          uuid.UUID   `json:"id"`
	Username    string      `json:"username"`
	Email       string      `json:"email"`
	Phone       string      `json:"phone"`
	RoleID      *uuid.UUID  `json:"role_id,omitempty"`
	Status      string      `json:"status"`
	IsVerified  bool        `json:"is_verified"`
	LastLoginAt *time.Time  `json:"last_login_at,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Role        *RoleDTO    `json:"role,omitempty"`
	Profile     *ProfileDTO `json:"profile,omitempty"`
}

type LoginResponseDTO struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	User         *UserDTO `json:"user"`
}

type RefreshResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// UserViewToDTO converts application UserView to response DTO
func UserViewToDTO(v *userAppViews.UserView) *UserDTO {
	if v == nil {
		return nil
	}

	dto := &UserDTO{
		ID:          v.ID,
		Username:    v.Username,
		Email:       v.Email,
		Phone:       v.Phone,
		RoleID:      v.RoleID,
		Status:      v.Status,
		IsVerified:  v.IsVerified,
		LastLoginAt: v.LastLoginAt,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}

	if v.Role != nil {
		// Note: userAppViews.RoleView doesn't have CreatedAt, UpdatedAt, DeletedAt
		// So we need to create a minimal RoleDTO
		roleDTO := &RoleDTO{
			ID:          v.Role.ID,
			Name:        v.Role.Name,
			Description: v.Role.Description,
			Permissions: []string{}, // Note: Permissions is *datatypes.JSON, need to unmarshal
			IsSystem:    v.Role.IsSystem,
		}
		dto.Role = roleDTO
	}

	return dto
}

// LoginResponseToDTO converts application LoginResponse to response DTO
func LoginResponseToDTO(resp *userAppCommands.LoginResponse) *LoginResponseDTO {
	if resp == nil {
		return nil
	}

	return &LoginResponseDTO{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		User:         UserViewToDTO(&resp.User),
	}
}

// RefreshResponseToDTO converts application RefreshResponse to response DTO
func RefreshResponseToDTO(resp *userAppCommands.RefreshResponse) *RefreshResponseDTO {
	if resp == nil {
		return nil
	}

	return &RefreshResponseDTO{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}
}

// UserListViewToDTO converts list of UserView to DTOs
func UserListViewToDTO(views []userAppViews.UserView) []UserDTO {
	dtos := make([]UserDTO, len(views))
	for i, v := range views {
		if dto := UserViewToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
