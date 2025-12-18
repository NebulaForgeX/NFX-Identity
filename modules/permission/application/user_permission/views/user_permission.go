package views

import (
	"time"

	userPermissionDomainViews "nfxid/modules/permission/domain/user_permission/views"

	"github.com/google/uuid"
)

type UserPermissionView struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	PermissionID uuid.UUID `json:"permission_id"`
	Tag          string    `json:"tag"`  // Permission tag for convenience
	Name         string    `json:"name"` // Permission name for convenience
	Category     string    `json:"category"`
	CreatedAt    time.Time `json:"created_at"`
}

// UserPermissionViewMapper 将 Domain UserPermissionView 转换为 Application UserPermissionView
func UserPermissionViewMapper(v userPermissionDomainViews.UserPermissionView) UserPermissionView {
	return UserPermissionView{
		ID:           v.ID,
		UserID:       v.UserID,
		PermissionID: v.PermissionID,
		Tag:          v.Tag,
		Name:         v.Name,
		Category:     v.Category,
		CreatedAt:    v.CreatedAt,
	}
}
