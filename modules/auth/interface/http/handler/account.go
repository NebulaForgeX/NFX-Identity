package handler

import (
	"strings"

	authErr "nfxidentity/errors/src/auth"
	sysErr "nfxidentity/errors/src/sys"
	account "nfxidentity/modules/auth/application/account"
	emailapp "nfxidentity/modules/auth/application/email"
	phoneapp "nfxidentity/modules/auth/application/phone"
	"nfxidentity/modules/auth/interface/http/dto/reqdto"
	"nfxidentity/pkgs/errx"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type AccountHandler struct {
	account  *account.Service
	email    *emailapp.Service
	phone    *phoneapp.Service
	validate *fiberx.StructValidator
}

func NewAccountHandler(accountSvc *account.Service, emailSvc *emailapp.Service, phoneSvc *phoneapp.Service) *AccountHandler {
	return &AccountHandler{
		account:  accountSvc,
		email:    emailSvc,
		phone:    phoneSvc,
		validate: fiberx.NewStructValidator(),
	}
}

func (h *AccountHandler) GetFullAccountInformationWithCommunityProfile(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	profileID, ok := fiberx.ProfileIDFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}

	out, err := h.account.GetFullInformationWithCommunityProfile(c.Context(), account.GetFullInformationWithCommunityProfileInput{AccountID: accountID, ProfileID: profileID})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Account information loaded", httpx.SuccessOptions{Data: out})
}

func (h *AccountHandler) UpdatePreference(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	profileID, ok := fiberx.ProfileIDFromContext(c.Context())
	if !ok || profileID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}

	var req reqdto.UpdatePreference
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	if err := h.account.UpdateCommunityPreference(c.Context(), account.UpdateCommunityPreferenceInput{
		AccountID:  accountID,
		ProfileID:  profileID,
		Preference: req.Preference,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Preference updated", httpx.SuccessOptions{})
}

func (h *AccountHandler) PatchCommunityProfile(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok || accountID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	profileID, ok := fiberx.ProfileIDFromContext(c.Context())
	if !ok || profileID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}

	var req reqdto.PatchCommunityProfile
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	patch := req.ToPatch()
	if err := patch.Validate(); err != nil {
		return err
	}
	if err := h.account.PatchCommunityProfile(c.Context(), account.PatchCommunityProfileInput{
		AccountID: accountID,
		ProfileID: profileID,
		Patch:     patch,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "User profile patched", httpx.SuccessOptions{})
}

func (h *AccountHandler) PatchCommunityProfileSettings(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok || accountID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	profileID, ok := fiberx.ProfileIDFromContext(c.Context())
	if !ok || profileID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}

	var req reqdto.PatchCommunityProfileSettings
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	patch := req.ToPatch()
	if err := patch.Validate(); err != nil {
		return err
	}
	if err := h.account.PatchCommunityProfileSettings(c.Context(), account.PatchCommunityProfileSettingsInput{
		AccountID: accountID,
		ProfileID: profileID,
		Patch:     patch,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "User profile settings patched", httpx.SuccessOptions{})
}

func (h *AccountHandler) ConfirmCommunityProfileAvatar(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.ConfirmCommunityProfileAvatar
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	imageID, err := uuid.Parse(strings.TrimSpace(req.ImageID))
	if err != nil || imageID == uuid.Nil {
		return authErr.ErrInvalidImageID.WithCause(err)
	}
	if err := h.account.ConfirmCommunityProfileAvatar(c.Context(), account.ConfirmCommunityProfileAvatarInput{
		AccountID: accountID,
		ProfileID: profileID,
		ImageID:   imageID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "User profile avatar updated", httpx.SuccessOptions{})
}

func (h *AccountHandler) ClearCommunityProfileAvatar(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	if err := h.account.ClearCommunityProfileAvatar(c.Context(), account.ClearCommunityProfileAvatarInput{
		AccountID: accountID,
		ProfileID: profileID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "User profile avatar cleared", httpx.SuccessOptions{})
}

func (h *AccountHandler) ListCommunityProfiles(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok || accountID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	out, err := h.account.ListCommunityProfiles(c.Context(), account.ListCommunityProfilesInput{AccountID: accountID})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "User profiles loaded", httpx.SuccessOptions{Data: out})
}

func (h *AccountHandler) GetPublicProfileCard(c fiber.Ctx) error {
	var uri reqdto.PublicProfileCardURI
	if err := c.Bind().URI(&uri); err != nil {
		return err
	}
	out, err := h.account.GetCommunityPublicProfileCard(c.Context(), account.GetCommunityPublicProfileCardInput{
		ProfileID: uri.ProfileID,
	})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Public profile card loaded", httpx.SuccessOptions{Data: out})
}

func (h *AccountHandler) SearchCommunityProfiles(c fiber.Ctx) error {
	profileID, ok := fiberx.ProfileIDFromContext(c.Context())
	if !ok || profileID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.SearchProfiles
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	req.Sanitize()
	out, err := h.account.SearchCommunityProfiles(c.Context(), account.SearchCommunityProfilesInput{
		Query: req.ToQuery(profileID),
	})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "User profiles searched", httpx.SuccessOptions{Data: out})
}

func (h *AccountHandler) CreateCommunityProfile(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok || accountID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.CreateCommunityProfile
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	req = req.Normalize()
	if err := req.Validate(); err != nil {
		return err
	}
	out, err := h.account.CreateCommunityProfile(c.Context(), account.CreateCommunityProfileInput{
		AccountID:       accountID,
		DisplayName:     req.DisplayName,
		ProfileLanguage: req.ProfileLanguage,
	})
	if err != nil {
		return err
	}
	return fiberx.Created(c, "User profile created", httpx.SuccessOptions{Data: out})
}

func (h *AccountHandler) DeleteCommunityProfile(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var uri reqdto.DeleteCommunityProfileURI
	if err := c.Bind().URI(&uri); err != nil {
		return err
	}
	if uri.ProfileID == uuid.Nil {
		return authErr.ErrForgerProfileNotFound
	}
	if err := h.account.DeleteCommunityProfile(c.Context(), account.DeleteCommunityProfileInput{
		AccountID:        accountID,
		ProfileID:        uri.ProfileID,
		CurrentProfileID: profileID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "User profile deleted", httpx.SuccessOptions{})
}

func (h *AccountHandler) ConfirmCommunityProfileBackgrounds(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.ConfirmCommunityProfileBackgrounds
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
	if _, err := h.account.ConfirmCommunityProfileBackgrounds(c.Context(), in); err != nil {
		return err
	}
	return fiberx.OK(c, "User profile backgrounds updated", httpx.SuccessOptions{})
}

func (h *AccountHandler) ListEmails(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok || accountID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	out, err := h.email.List(c.Context(), emailapp.ListInput{AccountID: accountID})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Emails loaded", httpx.SuccessOptions{Data: out})
}

func (h *AccountHandler) CreateEmail(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.CreateEmail
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	req = req.Normalize()
	if err := req.Validate(); err != nil {
		return err
	}
	out, err := h.email.Create(c.Context(), emailapp.CreateInput{
		AccountID: accountID,
		ProfileID: profileID,
		Email:     req.Email,
	})
	if err != nil {
		return err
	}
	return fiberx.Created(c, "Email added", httpx.SuccessOptions{Data: out})
}

func (h *AccountHandler) SendEmailVerificationCode(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var uri reqdto.EmailURI
	if err := c.Bind().URI(&uri); err != nil {
		return err
	}
	var req reqdto.SendEmailVerificationCode
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	if err := h.email.SendVerificationCode(c.Context(), emailapp.SendVerificationCodeInput{
		AccountID: accountID,
		ProfileID: profileID,
		EmailID:   uri.EmailID,
		Lang:      req.Lang,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Verification code sent", httpx.SuccessOptions{})
}

func (h *AccountHandler) VerifyEmail(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var uri reqdto.EmailURI
	if err := c.Bind().URI(&uri); err != nil {
		return err
	}
	var req reqdto.VerifyEmail
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	req = req.Normalize()
	if err := req.Validate(); err != nil {
		return err
	}
	if err := h.email.Verify(c.Context(), emailapp.VerifyInput{
		AccountID: accountID,
		ProfileID: profileID,
		EmailID:   uri.EmailID,
		Code:      req.VerificationCode,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Email verified", httpx.SuccessOptions{})
}

func (h *AccountHandler) UpdateEmail(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var uri reqdto.EmailURI
	if err := c.Bind().URI(&uri); err != nil {
		return err
	}
	var req reqdto.UpdateEmail
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	req = req.Normalize()
	if err := req.Validate(); err != nil {
		return err
	}
	if err := h.email.Update(c.Context(), emailapp.UpdateInput{
		AccountID: accountID,
		ProfileID: profileID,
		EmailID:   uri.EmailID,
		Email:     req.Email,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Email updated", httpx.SuccessOptions{})
}

func (h *AccountHandler) SetPrimaryEmail(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var uri reqdto.EmailURI
	if err := c.Bind().URI(&uri); err != nil {
		return err
	}
	if err := h.email.SetPrimary(c.Context(), emailapp.SetPrimaryInput{
		AccountID: accountID,
		ProfileID: profileID,
		EmailID:   uri.EmailID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Primary email updated", httpx.SuccessOptions{})
}

func (h *AccountHandler) DeleteEmail(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var uri reqdto.EmailURI
	if err := c.Bind().URI(&uri); err != nil {
		return err
	}
	if err := h.email.Delete(c.Context(), emailapp.DeleteInput{
		AccountID: accountID,
		ProfileID: profileID,
		EmailID:   uri.EmailID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Email deleted", httpx.SuccessOptions{})
}

func (h *AccountHandler) ListPhones(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok || accountID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	out, err := h.phone.List(c.Context(), phoneapp.ListInput{AccountID: accountID})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Phones loaded", httpx.SuccessOptions{Data: out})
}

func (h *AccountHandler) CreatePhone(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.CreatePhone
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	req = req.Normalize()
	if err := req.Validate(); err != nil {
		return err
	}
	out, err := h.phone.Create(c.Context(), phoneapp.CreateInput{
		AccountID: accountID,
		ProfileID: profileID,
		Phone:     req.Phone,
	})
	if err != nil {
		return err
	}
	return fiberx.Created(c, "Phone added", httpx.SuccessOptions{Data: out})
}

func (h *AccountHandler) SendPhoneVerificationCode(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var uri reqdto.PhoneURI
	if err := c.Bind().URI(&uri); err != nil {
		return err
	}
	if err := h.phone.SendVerificationCode(c.Context(), phoneapp.SendVerificationCodeInput{
		AccountID: accountID,
		ProfileID: profileID,
		PhoneID:   uri.PhoneID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Verification code sent", httpx.SuccessOptions{})
}

func (h *AccountHandler) VerifyPhone(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var uri reqdto.PhoneURI
	if err := c.Bind().URI(&uri); err != nil {
		return err
	}
	var req reqdto.VerifyPhone
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	req = req.Normalize()
	if err := req.Validate(); err != nil {
		return err
	}
	if err := h.phone.Verify(c.Context(), phoneapp.VerifyInput{
		AccountID: accountID,
		ProfileID: profileID,
		PhoneID:   uri.PhoneID,
		Code:      req.VerificationCode,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Phone verified", httpx.SuccessOptions{})
}

func (h *AccountHandler) UpdatePhone(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var uri reqdto.PhoneURI
	if err := c.Bind().URI(&uri); err != nil {
		return err
	}
	var req reqdto.UpdatePhone
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	req = req.Normalize()
	if err := req.Validate(); err != nil {
		return err
	}
	if err := h.phone.Update(c.Context(), phoneapp.UpdateInput{
		AccountID: accountID,
		ProfileID: profileID,
		PhoneID:   uri.PhoneID,
		Phone:     req.Phone,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Phone updated", httpx.SuccessOptions{})
}

func (h *AccountHandler) SetPrimaryPhone(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var uri reqdto.PhoneURI
	if err := c.Bind().URI(&uri); err != nil {
		return err
	}
	if err := h.phone.SetPrimary(c.Context(), phoneapp.SetPrimaryInput{
		AccountID: accountID,
		ProfileID: profileID,
		PhoneID:   uri.PhoneID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Primary phone updated", httpx.SuccessOptions{})
}

func (h *AccountHandler) DeletePhone(c fiber.Ctx) error {
	accountID, profileID, ok := fiberx.AccountProfileIDsFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}
	var uri reqdto.PhoneURI
	if err := c.Bind().URI(&uri); err != nil {
		return err
	}
	if err := h.phone.Delete(c.Context(), phoneapp.DeleteInput{
		AccountID: accountID,
		ProfileID: profileID,
		PhoneID:   uri.PhoneID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Phone deleted", httpx.SuccessOptions{})
}

func (h *AccountHandler) SendChangePasswordVerificationCode(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok || accountID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.SendChangePasswordVerificationCode
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	if err := h.account.SendChangePasswordVerificationCode(c.Context(), account.SendChangePasswordVerificationCodeInput{
		AccountID: accountID,
		Lang:      req.Lang,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Verification code sent", httpx.SuccessOptions{})
}

func (h *AccountHandler) ChangePassword(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok || accountID == uuid.Nil {
		return sysErr.ErrInvalidToken
	}
	var req reqdto.ChangePassword
	if err := c.Bind().Body(&req); err != nil {
		return err
	}
	req = req.Normalize()
	if err := req.Validate(); err != nil {
		return err
	}
	if err := h.account.ChangePassword(c.Context(), account.ChangePasswordInput{
		AccountID:        accountID,
		CurrentPassword:  req.CurrentPassword,
		NewPassword:      req.NewPassword,
		VerificationCode: req.VerificationCode,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Password updated", httpx.SuccessOptions{})
}
