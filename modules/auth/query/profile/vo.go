package profile

import (
	"github.com/google/uuid"
	"nfxidentity/enums"
	"time"
)

//* =============================== Authority profile VO =============================== !//
// AuthorityProfileItemVO is one row from auth.ListAuthorityProfileItems (login list, search, social enrichment).
type AuthorityProfileItemVO struct {
	ProfileID         uuid.UUID                 `json:"profile_id"`
	AccountID         uuid.UUID                 `json:"account_id"`
	AuthorityRoles    []enums.AuthAuthorityRole `json:"authority_roles"`
	DisplayName       *string                   `json:"display_name"`
	ProfileLanguage   enums.AuthProfileLanguage `json:"profile_language"`
	City              *string                   `json:"city"`
	Country           *string                   `json:"country"`
	Website           *string                   `json:"website"`
	Timezone          *string                   `json:"timezone"`
	Birthday          *time.Time                `json:"birthday"`
	AvatarImageID     *uuid.UUID                `json:"avatar_image_id"`
	BackgroundImageID *uuid.UUID                `json:"background_image_id"`
	CreatedAt         time.Time                 `json:"created_at"`
}

//* =============================== Forger profile VO =============================== !//
// ForgerProfileItemVO is one row from auth.ListForgerProfileItems (login list, search, social enrichment).
type ForgerProfileItemVO struct {
	ProfileID         uuid.UUID                 `json:"profile_id"`
	AccountID         uuid.UUID                 `json:"account_id"`
	ForgerRoles    []enums.AuthForgerRole `json:"forger_roles"`
	DisplayName       *string                   `json:"display_name"`
	ProfileLanguage   enums.AuthProfileLanguage `json:"profile_language"`
	City              *string                   `json:"city"`
	Country           *string                   `json:"country"`
	Website           *string                   `json:"website"`
	Timezone          *string                   `json:"timezone"`
	Birthday          *time.Time                `json:"birthday"`
	AvatarImageID     *uuid.UUID                `json:"avatar_image_id"`
	BackgroundImageID *uuid.UUID                `json:"background_image_id"`
	CreatedAt         time.Time                 `json:"created_at"`
}
