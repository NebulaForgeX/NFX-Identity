package respdto

import (
	authCommands "nfxid/modules/permission/application/auth/commands"
	userPermissionViews "nfxid/modules/permission/application/user_permission/views"
)

type LoginResponseDTO struct {
	AccessToken    string                  `json:"access_token"`
	RefreshToken   string                  `json:"refresh_token"`
	UserID         string                  `json:"user_id"`
	Username       string                  `json:"username"`
	Email          string                  `json:"email"`
	Phone          string                  `json:"phone"`
	Permissions    []UserPermissionViewDTO `json:"permissions"`
	PermissionTags []string                `json:"permission_tags"`
}

type UserPermissionViewDTO struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	PermissionID string `json:"permission_id"`
	Tag          string `json:"tag"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	CreatedAt    string `json:"created_at"`
}

func LoginResponseToDTO(resp *authCommands.LoginResponse) *LoginResponseDTO {
	if resp == nil {
		return nil
	}

	permissions := make([]UserPermissionViewDTO, len(resp.Permissions))
	for i, p := range resp.Permissions {
		permissions[i] = UserPermissionViewToAuthDTO(p)
	}

	return &LoginResponseDTO{
		AccessToken:    resp.AccessToken,
		RefreshToken:   resp.RefreshToken,
		UserID:         resp.UserID,
		Username:       resp.Username,
		Email:          resp.Email,
		Phone:          resp.Phone,
		Permissions:    permissions,
		PermissionTags: resp.PermissionTags,
	}
}

func UserPermissionViewToAuthDTO(v *userPermissionViews.UserPermissionView) UserPermissionViewDTO {
	if v == nil {
		return UserPermissionViewDTO{}
	}

	return UserPermissionViewDTO{
		ID:           v.ID.String(),
		UserID:       v.UserID.String(),
		PermissionID: v.PermissionID.String(),
		Tag:          v.Tag,
		Name:         v.Name,
		Category:     string(v.Category), // Convert enum to string for JSON
		CreatedAt:    v.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
