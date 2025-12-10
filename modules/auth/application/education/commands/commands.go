package commands

import (
	educationDomain "nebulaid/modules/auth/domain/education"

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
