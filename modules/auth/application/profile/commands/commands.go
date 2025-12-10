package commands

import (
	profileDomain "nebulaid/modules/auth/domain/profile"

	"github.com/google/uuid"
)

type CreateProfileCmd struct {
	UserID   uuid.UUID
	Editable profileDomain.ProfileEditable
}

type UpdateProfileCmd struct {
	ProfileID uuid.UUID
	Editable  profileDomain.ProfileEditable
}
