package respdto

import (
	"time"

	actionAppResult "nfxid/modules/access/application/actions/results"

	"github.com/google/uuid"
)

type ActionDTO struct {
	ID          uuid.UUID  `json:"id"`
	Key         string     `json:"key"`
	Service     string     `json:"service"`
	Status      string     `json:"status"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	IsSystem    bool       `json:"is_system"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func ActionROToDTO(v *actionAppResult.ActionRO) *ActionDTO {
	if v == nil {
		return nil
	}
	return &ActionDTO{
		ID:          v.ID,
		Key:         v.Key,
		Service:     v.Service,
		Status:      v.Status,
		Name:        v.Name,
		Description: v.Description,
		IsSystem:    v.IsSystem,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		DeletedAt:   v.DeletedAt,
	}
}
