package reqdto

import (
	"github.com/google/uuid"
)

type ByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}
