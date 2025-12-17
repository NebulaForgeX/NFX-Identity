package views

import (
	"time"

	"github.com/google/uuid"
)

type BadgeView struct {
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
