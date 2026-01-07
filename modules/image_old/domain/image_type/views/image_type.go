package views

import (
	"time"

	"github.com/google/uuid"
)

type ImageTypeView struct {
	ID          uuid.UUID
	Key         string
	Description *string
	MaxWidth    *int
	MaxHeight   *int
	AspectRatio *string
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

