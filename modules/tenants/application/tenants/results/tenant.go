package results

import (
	"time"

	"nfxid/modules/tenants/domain/tenants"

	"github.com/google/uuid"
)

type TenantRO struct {
	ID            uuid.UUID
	TenantID      string
	Name          string
	DisplayName   *string
	Status        tenants.TenantStatus
	PrimaryDomain *string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Metadata      map[string]interface{}
}

// TenantMapper 将 Domain Tenant 转换为 Application TenantRO
func TenantMapper(t *tenants.Tenant) TenantRO {
	if t == nil {
		return TenantRO{}
	}

	return TenantRO{
		ID:            t.ID(),
		TenantID:      t.TenantID(),
		Name:          t.Name(),
		DisplayName:   t.DisplayName(),
		Status:        t.Status(),
		PrimaryDomain: t.PrimaryDomain(),
		CreatedAt:     t.CreatedAt(),
		UpdatedAt:     t.UpdatedAt(),
		DeletedAt:     t.DeletedAt(),
		Metadata:      t.Metadata(),
	}
}
