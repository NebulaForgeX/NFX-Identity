package respdto

import (
	"encoding/json"
	"time"

	roleAppViews "nfxid/modules/auth/application/role/views"

	"github.com/google/uuid"
)

type RoleDTO struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Permissions []string   `json:"permissions"`
	IsSystem    bool       `json:"is_system"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// RoleViewToDTO converts application RoleView to response DTO
func RoleViewToDTO(v *roleAppViews.RoleView) *RoleDTO {
	if v == nil {
		return nil
	}

	permissions := []string{}
	if v.Permissions != nil {
		var perms []interface{}
		if err := json.Unmarshal(*v.Permissions, &perms); err == nil {
			permissions = make([]string, 0, len(perms))
			for _, p := range perms {
				if perm, ok := p.(string); ok {
					permissions = append(permissions, perm)
				}
			}
		} else {
			// 如果 unmarshal 失败，尝试作为字符串数组直接解析
			var permsStr []string
			if err2 := json.Unmarshal(*v.Permissions, &permsStr); err2 == nil {
				permissions = permsStr
			}
		}
	}

	return &RoleDTO{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
		Permissions: permissions,
		IsSystem:    v.IsSystem,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		DeletedAt:   v.DeletedAt,
	}
}

// RoleListViewToDTO converts list of RoleView to DTOs
func RoleListViewToDTO(views []roleAppViews.RoleView) []RoleDTO {
	dtos := make([]RoleDTO, len(views))
	for i, v := range views {
		if dto := RoleViewToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
