package results

import (
	"time"

	"nfxid/modules/directory/domain/badges"

	"github.com/google/uuid"
)

type BadgeRO struct {
	ID          uuid.UUID
	Name        string
	Description *string
	IconURL     *string
	Color       *string
	Category    *string
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// BadgeMapper 将 Domain Badge 转换为 Application BadgeRO
func BadgeMapper(b *badges.Badge) BadgeRO {
	if b == nil {
		return BadgeRO{}
	}

	return BadgeRO{
		ID:          b.ID(),
		Name:        b.Name(),
		Description: b.Description(),
		IconURL:     b.IconURL(),
		Color:       b.Color(),
		Category:    b.Category(),
		IsSystem:    b.IsSystem(),
		CreatedAt:   b.CreatedAt(),
		UpdatedAt:   b.UpdatedAt(),
		DeletedAt:   b.DeletedAt(),
	}
}
