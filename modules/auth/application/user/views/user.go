package views

import (
	"encoding/json"
	"time"

	userDomainViews "nfxid/modules/auth/domain/user/views"

	"github.com/google/uuid"
)

type UserView struct {
	ID           uuid.UUID  `json:"id"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	Phone        *string    `json:"phone,omitempty"`
	PasswordHash string     `json:"password_hash"`
	Status       string     `json:"status"`
	IsVerified   bool       `json:"is_verified"`
	LastLoginAt  *time.Time `json:"last_login_at"`
	Roles        []RoleView `json:"roles,omitempty"` // Array of roles (from user_with_role_view)
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// UserViewMapper 将 Domain UserView 转换为 Application UserView
func UserViewMapper(v userDomainViews.UserView) UserView {
	// Parse roles from JSON
	var roles []RoleView
	if len(v.Roles) > 0 {
		// Unmarshal roles JSON array from user_with_role_view
		var roleData []struct {
			RoleID          string          `json:"role_id"`
			RoleName        string          `json:"role_name"`
			RoleDescription *string         `json:"role_description"`
			RolePermissions json.RawMessage `json:"role_permissions"`
		}
		if err := json.Unmarshal(v.Roles, &roleData); err == nil {
			roles = make([]RoleView, 0, len(roleData))
			for _, rd := range roleData {
				roleID, _ := uuid.Parse(rd.RoleID)
				role := RoleView{
					ID:          roleID,
					Name:        rd.RoleName,
					Description: rd.RoleDescription,
					IsSystem:    false, // Default, can be enhanced if needed
				}
				roles = append(roles, role)
			}
		}
	}

	return UserView{
		ID:           v.ID,
		Username:     v.Username,
		Email:        v.Email,
		Phone:        v.Phone,
		PasswordHash: v.PasswordHash,
		Status:       v.Status,
		IsVerified:   v.IsVerified,
		LastLoginAt:  v.LastLoginAt,
		Roles:        roles,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
	}
}
