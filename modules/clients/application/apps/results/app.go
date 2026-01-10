package results

import (
	"time"

	"nfxid/modules/clients/domain/apps"

	"github.com/google/uuid"
)

type AppRO struct {
	ID          uuid.UUID
	AppID       string
	TenantID    uuid.UUID
	Name        string
	Description *string
	Type        apps.AppType
	Status      apps.AppStatus
	Environment apps.Environment
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   *uuid.UUID
	UpdatedBy   *uuid.UUID
	Metadata    map[string]interface{}
	DeletedAt   *time.Time
}

// AppMapper 将 Domain App 转换为 Application AppRO
func AppMapper(a *apps.App) AppRO {
	if a == nil {
		return AppRO{}
	}

	return AppRO{
		ID:          a.ID(),
		AppID:       a.AppID(),
		TenantID:    a.TenantID(),
		Name:        a.Name(),
		Description: a.Description(),
		Type:        a.Type(),
		Status:      a.Status(),
		Environment: a.Environment(),
		CreatedAt:   a.CreatedAt(),
		UpdatedAt:   a.UpdatedAt(),
		CreatedBy:   a.CreatedBy(),
		UpdatedBy:   a.UpdatedBy(),
		Metadata:    a.Metadata(),
		DeletedAt:   a.DeletedAt(),
	}
}
