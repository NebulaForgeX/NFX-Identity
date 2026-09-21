package reqdto

import (
	"strings"
	"time"

	"nfxidentity/constants"
	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	"nfxidentity/pkgs/patch"

	"github.com/google/uuid"
)

// CreateAuthorityProfile is the request body for POST /auth/me/authority-profiles.
type CreateAuthorityProfile struct {
	DisplayName     string                    `json:"display_name"`
	ProfileLanguage enums.AuthProfileLanguage `json:"profile_language"`
}

func (r CreateAuthorityProfile) Normalize() CreateAuthorityProfile {
	return CreateAuthorityProfile{
		DisplayName:     strings.TrimSpace(r.DisplayName),
		ProfileLanguage: r.ProfileLanguage,
	}
}

func (r CreateAuthorityProfile) Validate() error {
	if r.DisplayName == "" {
		return authErr.ErrAuthorityProfileFieldInvalid
	}
	if len(r.DisplayName) > 150 {
		return authErr.ErrAuthorityProfileFieldInvalid
	}
	if r.ProfileLanguage != "" && !constants.AuthLanguage.Valid(r.ProfileLanguage) {
		return authErr.ErrAuthorityProfileLanguageInvalid
	}
	return nil
}

// DeleteAuthorityProfileURI binds the profile id path param for DELETE /auth/me/authority-profiles/:profileId.
type DeleteAuthorityProfileURI struct {
	ProfileID uuid.UUID `uri:"profileId"`
}

// ConfirmAuthorityProfileAvatar is the request body for PUT /auth/me/authority-profile/avatars.
type ConfirmAuthorityProfileAvatar struct {
	ImageID string `json:"image_id"`
}

// PatchAuthorityProfile is the request body for PATCH /auth/me/authority-profile.
type PatchAuthorityProfile struct {
	ProfileLanguage patch.Field[enums.AuthProfileLanguage] `json:"profile_language"`
	DisplayName     patch.Field[string]                    `json:"display_name"`
	FirstName       patch.Field[string]                    `json:"first_name"`
	LastName        patch.Field[string]                    `json:"last_name"`
	Country         patch.Field[string]                    `json:"country"`
	City            patch.Field[string]                    `json:"city"`
	Gender          patch.Field[string]                    `json:"gender"`
	Birthday        patch.Field[time.Time]                 `json:"birthday"`
	Website         patch.Field[string]                    `json:"website"`
	Timezone        patch.Field[string]                    `json:"timezone"`
	Bio             patch.Field[string]                    `json:"bio"`
}

func (r PatchAuthorityProfile) ToPatch() profileDomain.AuthorityProfilePatch {
	return profileDomain.AuthorityProfilePatch{
		ProfileLanguage: r.ProfileLanguage.PatchField,
		DisplayName:     normalizeOptionalStringPatch(r.DisplayName.PatchField),
		FirstName:       normalizeOptionalStringPatch(r.FirstName.PatchField),
		LastName:        normalizeOptionalStringPatch(r.LastName.PatchField),
		Country:         normalizeOptionalStringPatch(r.Country.PatchField),
		City:            normalizeOptionalStringPatch(r.City.PatchField),
		Gender:          normalizeOptionalStringPatch(r.Gender.PatchField),
		Birthday:        r.Birthday.PatchField,
		Website:         normalizeOptionalStringPatch(r.Website.PatchField),
		Timezone:        normalizeOptionalStringPatch(r.Timezone.PatchField),
		Bio:             normalizeOptionalStringPatch(r.Bio.PatchField),
	}
}

// PatchAuthorityProfileSettings is the request body for PATCH /auth/me/authority-profile-settings.
type PatchAuthorityProfileSettings struct {
	LoginNotification patch.Field[bool] `json:"login_notification"`
}

func (r PatchAuthorityProfileSettings) ToPatch() profileDomain.AuthorityProfileSettingsPatch {
	return profileDomain.AuthorityProfileSettingsPatch{
		LoginNotification: r.LoginNotification.PatchField,
	}
}
