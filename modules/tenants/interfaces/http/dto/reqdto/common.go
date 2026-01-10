package reqdto

import (
	"github.com/google/uuid"
)

type ByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}
