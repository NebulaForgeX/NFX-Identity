package views

import (
	"time"

	userRoleDomainViews "nfxid/modules/auth/domain/user_role/views"

	"github.com/google/uuid"
)

type UserRoleView struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	RoleID    uuid.UUID `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}

// UserRoleViewMapper 将 Domain UserRoleView 转换为 Application UserRoleView
func UserRoleViewMapper(v userRoleDomainViews.UserRoleView) UserRoleView {
	return UserRoleView{
		ID:        v.ID,
		UserID:    v.UserID,
		RoleID:    v.RoleID,
		CreatedAt: v.CreatedAt,
	}
}

