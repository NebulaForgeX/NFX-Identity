package results

import (
	"time"

	"nfxid/modules/access/domain/actions"

	"github.com/google/uuid"
)

type ActionRO struct {
	ID          uuid.UUID
	Key         string
	Service     string
	Status      string
	Name        string
	Description *string
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func ActionMapper(a *actions.Action) ActionRO {
	if a == nil {
		return ActionRO{}
	}
	return ActionRO{
		ID:          a.ID(),
		Key:         a.Key(),
		Service:     a.Service(),
		Status:      a.Status(),
		Name:        a.Name(),
		Description: a.Description(),
		IsSystem:    a.IsSystem(),
		CreatedAt:   a.CreatedAt(),
		UpdatedAt:   a.UpdatedAt(),
		DeletedAt:   a.DeletedAt(),
	}
}
