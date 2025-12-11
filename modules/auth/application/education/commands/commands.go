package commands

import (
	educationDomain "nfxid/modules/auth/domain/education"

	"github.com/google/uuid"
)

type CreateEducationCmd struct {
	ProfileID uuid.UUID
	Editable  educationDomain.EducationEditable
}

type UpdateEducationCmd struct {
	EducationID uuid.UUID
	Editable    educationDomain.EducationEditable
}
