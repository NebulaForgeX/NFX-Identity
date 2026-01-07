package views

import (
	"nfxid/enums"
	"time"

	permissionDomainViews "nfxid/modules/permission/domain/permission/views"

	"github.com/google/uuid"
)

type PermissionView struct {
	ID          uuid.UUID                `json:"id"`
	Tag         string                   `json:"tag"`
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Category    enums.PermissionCategory `json:"category"`
	IsSystem    bool                     `json:"is_system"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
}

// PermissionViewMapper 将 Domain PermissionView 转换为 Application PermissionView
func PermissionViewMapper(v permissionDomainViews.PermissionView) PermissionView {
	return PermissionView{
		ID:          v.ID,
		Tag:         v.Tag,
		Name:        v.Name,
		Description: v.Description,
		Category:    v.Category,
		IsSystem:    v.IsSystem,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}
}
