package mapper

import (
	"time"

	"nfxid/modules/access/domain/actions"
	"nfxid/pkgs/utils/timex"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ActionModel matches access.actions table (includes service, status; dbgen may not have them)
type ActionModel struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Key         string         `gorm:"type:varchar(255);uniqueIndex;not null"`
	Service     string         `gorm:"type:varchar(255);not null"`
	Status      string         `gorm:"type:varchar(50);not null;default:active"`
	Name        string         `gorm:"type:varchar(255);not null"`
	Description *string        `gorm:"type:text"`
	IsSystem    bool           `gorm:"type:boolean;not null;default:false"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (ActionModel) TableName() string { return "access.actions" }

func ActionDomainToModel(a *actions.Action) *ActionModel {
	if a == nil {
		return nil
	}
	return &ActionModel{
		ID:          a.ID(),
		Key:         a.Key(),
		Service:     a.Service(),
		Status:      a.Status(),
		Name:        a.Name(),
		Description: a.Description(),
		IsSystem:    a.IsSystem(),
		CreatedAt:   a.CreatedAt(),
		UpdatedAt:   a.UpdatedAt(),
		DeletedAt:   timex.TimeToGormDeletedAt(a.DeletedAt()),
	}
}

func ActionModelToDomain(m *ActionModel) *actions.Action {
	if m == nil {
		return nil
	}
	state := actions.ActionState{
		ID:          m.ID,
		Key:         m.Key,
		Service:     m.Service,
		Status:      m.Status,
		Name:        m.Name,
		Description: m.Description,
		IsSystem:    m.IsSystem,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   timex.GormDeletedAtToTime(m.DeletedAt),
	}
	return actions.NewActionFromState(state)
}

func ActionModelToUpdates(m *ActionModel) map[string]any {
	return map[string]any{
		"key":         m.Key,
		"service":     m.Service,
		"status":      m.Status,
		"name":        m.Name,
		"description": m.Description,
		"is_system":   m.IsSystem,
		"updated_at":  m.UpdatedAt,
		"deleted_at":  m.DeletedAt,
	}
}
