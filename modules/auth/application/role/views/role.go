package views

import (
	"time"

	roleDomainViews "nebulaid/modules/auth/domain/role/views"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type RoleView struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	Permissions *datatypes.JSON `json:"permissions"`
	IsSystem    bool            `json:"is_system"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   *time.Time      `json:"deleted_at,omitempty"`
}

// RoleViewMapper 将 Domain RoleView 转换为 Application RoleView
func RoleViewMapper(v roleDomainViews.RoleView) RoleView {
	return RoleView{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
		Permissions: v.Permissions,
		IsSystem:    v.IsSystem,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		DeletedAt:   v.DeletedAt,
	}
}
