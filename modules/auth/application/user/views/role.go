package views

import (
	userDomainViews "nebulaid/modules/auth/domain/user/views"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type RoleView struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	Permissions *datatypes.JSON `json:"permissions"`
	IsSystem    bool            `json:"is_system"`
}

// RoleViewMapper 将 Domain RoleView 转换为 Application RoleView
func RoleViewMapper(v userDomainViews.RoleView) RoleView {
	return RoleView{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
		Permissions: v.Permissions,
		IsSystem:    v.IsSystem,
	}
}
