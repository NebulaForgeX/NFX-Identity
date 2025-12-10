package views

import (
	"time"

	userDomainViews "nebulaid/modules/auth/domain/user/views"

	"github.com/google/uuid"
)

type UserView struct {
	ID          uuid.UUID  `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone"`
	Status      string     `json:"status"`
	IsVerified  bool       `json:"is_verified"`
	LastLoginAt *time.Time `json:"last_login_at"`
	RoleID      *uuid.UUID `json:"role_id"`
	Role        *RoleView  `json:"role,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// UserViewMapper 将 Domain UserView 转换为 Application UserView
func UserViewMapper(v userDomainViews.UserView) UserView {
	return UserView{
		ID:          v.ID,
		Username:    v.Username,
		Email:       v.Email,
		Phone:       v.Phone,
		Status:      v.Status,
		IsVerified:  v.IsVerified,
		LastLoginAt: v.LastLoginAt,
		RoleID:      v.RoleID,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}
}
