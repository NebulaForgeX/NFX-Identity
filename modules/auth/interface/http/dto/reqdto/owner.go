package reqdto

import "github.com/google/uuid"

type UpdateAuthorityRoles struct {
	AuthorityRoles []string `json:"authority_roles"`
}

type OwnerProfileURI struct {
	ProfileID uuid.UUID `uri:"profileId"`
}
