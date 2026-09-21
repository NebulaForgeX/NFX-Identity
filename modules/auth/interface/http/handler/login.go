package handler

import (
	"strings"

	authErr "nfxidentity/errors/src/auth"
	sysErr "nfxidentity/errors/src/sys"
	authmsg "nfxidentity/messages/src/auth"
	login "nfxidentity/modules/auth/application/login"
	"nfxidentity/modules/auth/interface/http/dto/reqdto"
	"nfxidentity/pkgs/errx"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type LoginHandler struct {
	login    *login.Service
	validate *fiberx.StructValidator
}

func NewLoginHandler(loginSvc *login.Service) *LoginHandler {
	return &LoginHandler{
		login:    loginSvc,
		validate: fiberx.NewStructValidator(),
	}
}

func (h *LoginHandler) WithEmail(c fiber.Ctx) error {
	var req reqdto.LoginWithEmailRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))
	req.Password = strings.TrimSpace(req.Password)
	req.DeviceID = strings.TrimSpace(req.DeviceID)
	if req.Email == "" {
		return authErr.ErrEmailAddressRequired
	}
	if req.Password == "" {
		return authErr.ErrAccountPasswordRequired
	}
	if req.DeviceID == "" {
		return authErr.ErrAccountDeviceIDRequired
	}
	if err := h.validate.Validate(&req); err != nil {
		return err
	}

	result, err := h.login.WithEmail(c.Context(), login.WithEmailInput{
		Email:    req.Email,
		Password: req.Password,
		DeviceID: req.DeviceID,
	})
	if err != nil {
		return err
	}
	return fiberx.OK(c, authmsg.LOGIN_SUCCESS, httpx.SuccessOptions{Data: result})
}

func (h *LoginHandler) WithPhone(c fiber.Ctx) error {
	var req reqdto.LoginWithPhoneRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	req.Phone = strings.TrimSpace(req.Phone)
	req.Password = strings.TrimSpace(req.Password)
	req.DeviceID = strings.TrimSpace(req.DeviceID)
	if err := h.validate.Validate(&req); err != nil {
		return err
	}
	result, err := h.login.WithPhone(c.Context(), login.WithPhoneInput{
		Phone:    req.Phone,
		Password: req.Password,
		DeviceID: req.DeviceID,
	})
	if err != nil {
		return err
	}
	return fiberx.OK(c, authmsg.LOGIN_SUCCESS, httpx.SuccessOptions{Data: result})
}

func (h *LoginHandler) SelectProfile(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok {
		return sysErr.ErrInvalidToken
	}

	var req reqdto.SelectProfileRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	if err := h.validate.Validate(&req); err != nil {
		return err
	}

	if req.ProfileID == uuid.Nil {
		return authErr.ErrForgerProfileNotFound
	}

	currentProfileID, hasCurrentProfile := fiberx.ProfileIDFromContext(c.Context())
	completeLogin := !hasCurrentProfile || currentProfileID == uuid.Nil
	loginEmail, _ := fiberx.LoginEmailFromContext(c.Context())

	result, err := h.login.SelectProfile(c.Context(), login.SelectProfileInput{
		AccountID:     accountID,
		ProfileID:     req.ProfileID,
		Kind:          req.Kind,
		DeviceID:      req.DeviceID,
		CompleteLogin: completeLogin,
		LoginEmail:    loginEmail,
	})
	if err != nil {
		return err
	}
	return fiberx.OK(c, authmsg.PROFILE_SELECTED, httpx.SuccessOptions{Data: result})
}

func (h *LoginHandler) Refresh(c fiber.Ctx) error {
	var req reqdto.RefreshTokensRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	req.RefreshToken = strings.TrimSpace(req.RefreshToken)
	req.DeviceID = strings.TrimSpace(req.DeviceID)
	if req.RefreshToken == "" {
		return authErr.ErrInvalidRefreshToken
	}
	if req.DeviceID == "" {
		return authErr.ErrAccountDeviceIDRequired
	}
	if err := h.validate.Validate(&req); err != nil {
		return err
	}

	result, err := h.login.Refresh(c.Context(), login.RefreshInput{
		RefreshToken: req.RefreshToken,
		DeviceID:     req.DeviceID,
	})
	if err != nil {
		return err
	}
	return fiberx.OK(c, authmsg.TOKENS_REFRESHED, httpx.SuccessOptions{Data: result})
}

func (h *LoginHandler) Logout(c fiber.Ctx) error {
	var req reqdto.LogoutRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	if err := h.login.Logout(c.Context(), strings.TrimSpace(req.RefreshToken)); err != nil {
		return err
	}
	return fiberx.OK(c, authmsg.LOGOUT_SUCCESS, httpx.SuccessOptions{Data: nil})
}
