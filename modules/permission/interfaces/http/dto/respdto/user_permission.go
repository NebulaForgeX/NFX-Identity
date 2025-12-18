package respdto

import (
	userPermissionViews "nfxid/modules/permission/application/user_permission/views"
	"time"

	"github.com/google/uuid"
)

type UserPermissionDTO struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	PermissionID uuid.UUID `json:"permission_id"`
	Tag          string    `json:"tag"`
	Name         string    `json:"name"`
	Category     string    `json:"category"`
	CreatedAt    time.Time `json:"created_at"`
}

func UserPermissionViewToDTO(v *userPermissionViews.UserPermissionView) *UserPermissionDTO {
	if v == nil {
		return nil
	}

	return &UserPermissionDTO{
		ID:           v.ID,
		UserID:       v.UserID,
		PermissionID: v.PermissionID,
		Tag:          v.Tag,
		Name:         v.Name,
		Category:     string(v.Category), // Convert enum to string for JSON
		CreatedAt:    v.CreatedAt,
	}
}
