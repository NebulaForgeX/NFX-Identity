package commands

import (
	occupationDomain "nfxid/modules/auth/domain/occupation"

	"github.com/google/uuid"
)

type CreateOccupationCmd struct {
	ProfileID uuid.UUID
	Editable  occupationDomain.OccupationEditable
}

type UpdateOccupationCmd struct {
	OccupationID uuid.UUID
	Editable     occupationDomain.OccupationEditable
}
