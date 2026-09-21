package handler

import (
	"strings"

	"nfxidentity/constants"
	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	sysErr "nfxidentity/errors/src/sys"
	account "nfxidentity/modules/auth/application/account"
	"nfxidentity/modules/auth/interface/http/dto/reqdto"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type AuthorityAccountHandler struct {
	account *account.Service
}

func NewAuthorityAccountHandler(accountSvc *account.Service) *AuthorityAccountHandler {
	return &AuthorityAccountHandler{account: accountSvc}
}

func (h *AuthorityAccountHandler) GetFullAccountInformationWithAuthorityProfile(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	profileID, ok := fiberx.ProfileIDFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	out, err := h.account.GetFullInformationWithAuthorityProfile(c.Context(), account.GetFullInformationWithAuthorityProfileInput{AccountID: accountID, ProfileID: profileID})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Account information loaded", httpx.SuccessOptions{Data: out})
}

func (h *AuthorityAccountHandler) PatchAuthorityProfile(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok || accountID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	profileID, ok := fiberx.ProfileIDFromContext(c.Context())
	if !ok || profileID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.PatchAuthorityProfile
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	patch := req.ToPatch()
	if err := patch.Validate(); err != nil {
		return err
	}
	if err := h.account.PatchAuthorityProfile(c.Context(), account.PatchAuthorityProfileInput{
		AccountID: accountID,
		ProfileID: profileID,
		Patch:     patch,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Authority profile patched", httpx.SuccessOptions{})
}

func (h *AuthorityAccountHandler) PatchAuthorityProfileSettings(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok || accountID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	profileID, ok := fiberx.ProfileIDFromContext(c.Context())
	if !ok || profileID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.PatchAuthorityProfileSettings
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	patch := req.ToPatch()
	if err := patch.Validate(); err != nil {
		return err
	}
	if err := h.account.PatchAuthorityProfileSettings(c.Context(), account.PatchAuthorityProfileSettingsInput{
		AccountID: accountID,
		ProfileID: profileID,
		Patch:     patch,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Authority profile settings patched", httpx.SuccessOptions{})
}

func (h *AuthorityAccountHandler) ConfirmAuthorityProfileAvatar(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.ConfirmAuthorityProfileAvatar
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	imageID, err := uuid.Parse(strings.TrimSpace(req.ImageID))
	if err != nil || imageID == uuid.Nil {
		return authErr.ErrInvalidImageID.WithCause(err)
	}
	if err := h.account.ConfirmAuthorityProfileAvatar(c.Context(), account.ConfirmAuthorityProfileAvatarInput{
		AccountID: accountID,
		ProfileID: profileID,
		ImageID:   imageID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Authority profile avatar updated", httpx.SuccessOptions{})
}

func (h *AuthorityAccountHandler) ClearAuthorityProfileAvatar(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	if err := h.account.ClearAuthorityProfileAvatar(c.Context(), account.ClearAuthorityProfileAvatarInput{
		AccountID: accountID,
		ProfileID: profileID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Authority profile avatar cleared", httpx.SuccessOptions{})
}

func (h *AuthorityAccountHandler) ListAuthorityProfiles(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok || accountID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	out, err := h.account.ListAuthorityProfiles(c.Context(), account.ListAuthorityProfilesInput{AccountID: accountID})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Authority profiles loaded", httpx.SuccessOptions{Data: out})
}

func (h *AuthorityAccountHandler) SearchAuthorityProfiles(c fiber.Ctx) error {
	profileID, ok := fiberx.ProfileIDFromContext(c.Context())
	if !ok || profileID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	actorScope, ok := fiberx.ProfileScopeFromContext(c.Context())
	if !ok || !constants.AuthProfileScope.Valid(actorScope) {
		return sysErr.ErrInvalidToken
	}
	if actorScope != enums.AuthProfileScopeAuthority {
		return authErr.ErrAuthorityProfileScopeRequired
	}
	var req reqdto.SearchProfiles
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	req.Sanitize()
	out, err := h.account.SearchAuthorityProfiles(c.Context(), account.SearchAuthorityProfilesInput{
		Query: req.ToQuery(profileID),
	})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Authority profiles searched", httpx.SuccessOptions{Data: out})
}

func (h *AuthorityAccountHandler) CreateAuthorityProfile(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok || accountID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.CreateAuthorityProfile
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	req = req.Normalize()
	if err := req.Validate(); err != nil {
		return err
	}
	out, err := h.account.CreateAuthorityProfile(c.Context(), account.CreateAuthorityProfileInput{
		AccountID:       accountID,
		DisplayName:     req.DisplayName,
		ProfileLanguage: req.ProfileLanguage,
	})
	if err != nil {
		return err
	}
	return fiberx.Created(c, "Authority profile created", httpx.SuccessOptions{Data: out})
}

func (h *AuthorityAccountHandler) DeleteAuthorityProfile(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var uri reqdto.DeleteAuthorityProfileURI
	if err := c.Bind().URI(&uri); err != nil {
		return err
	}
	if uri.ProfileID == uuid.Nil {
		return authErr.ErrAuthorityProfileNotFound
	}
	if err := h.account.DeleteAuthorityProfile(c.Context(), account.DeleteAuthorityProfileInput{
		AccountID:        accountID,
		ProfileID:        uri.ProfileID,
		CurrentProfileID: profileID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Authority profile deleted", httpx.SuccessOptions{})
}

func (h *AuthorityAccountHandler) UpdateAuthorityPreference(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.UpdatePreference
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	if err := h.account.UpdateAuthorityPreference(c.Context(), account.UpdateAuthorityPreferenceInput{
		AccountID:  accountID,
		ProfileID:  profileID,
		Preference: req.Preference,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Preference updated", httpx.SuccessOptions{})
}

func (h *AuthorityAccountHandler) ConfirmAuthorityProfileBackgrounds(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.ConfirmAuthorityProfileBackgrounds
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	if err := req.Validate(); err != nil {
		return err
	}
	in, err := req.ToInput(accountID, profileID)
	if err != nil {
		return err
	}
	if _, err := h.account.ConfirmAuthorityProfileBackgrounds(c.Context(), in); err != nil {
		return err
	}
	return fiberx.OK(c, "Authority profile backgrounds updated", httpx.SuccessOptions{})
}
