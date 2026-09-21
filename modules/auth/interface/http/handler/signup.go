package handler

import (
	authErr "nfxidentity/errors/src/auth"
	authmsg "nfxidentity/messages/src/auth"
	signup "nfxidentity/modules/auth/application/signup"
	"nfxidentity/modules/auth/interface/http/dto/reqdto"
	"nfxidentity/pkgs/errx"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/httpx"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type SignupHandler struct {
	signup   *signup.Service
	validate *fiberx.StructValidator
}

func NewSignupHandler(signupSvc *signup.Service) *SignupHandler {
	return &SignupHandler{
		signup:   signupSvc,
		validate: fiberx.NewStructValidator(),
	}
}

func (h *SignupHandler) BySendingCode(c fiber.Ctx) error {
	var req reqdto.SendVerificationCodeRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	if err := h.validate.Validate(&req); err != nil {
		return err
	}
	email := strings.ToLower(strings.TrimSpace(req.Email))
	if email == "" {
		return authErr.ErrEmailAddressRequired
	}
	if err := h.signup.BySendingCode(c.Context(), email, string(req.Lang)); err != nil {
		return err
	}
	return fiberx.OK(c, authmsg.VERIFICATION_CODE_SENT, httpx.SuccessOptions{})
}

func (h *SignupHandler) WithEmail(c fiber.Ctx) error {
	var req reqdto.SignupRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	if err := h.validate.Validate(&req); err != nil {
		return err
	}

	email := strings.ToLower(strings.TrimSpace(req.Email))
	if email == "" {
		return authErr.ErrEmailAddressRequired
	}
	password := strings.TrimSpace(req.Password)
	if password == "" {
		return authErr.ErrAccountPasswordRequired
	}
	verificationCode := strings.ToUpper(strings.TrimSpace(req.VerificationCode))
	if verificationCode == "" {
		return authErr.ErrVerificationCodeWrong
	}
	req.DeviceID = strings.TrimSpace(req.DeviceID)
	if req.DeviceID == "" {
		return authErr.ErrAccountDeviceIDRequired
	}
	if req.SignupPlatform == "" {
		return authErr.ErrAccountSignupPlatformInvalid
	}
	result, err := h.signup.WithEmail(c.Context(), signup.WithEmailInput{
		Email:            email,
		Password:         password,
		VerificationCode: verificationCode,
		Lang:             req.Lang,
		DeviceID:         req.DeviceID,
		SignupPlatform:   req.SignupPlatform,
	})
	if err != nil {
		return err
	}
	return fiberx.OK(c, authmsg.SIGNUP_SUCCESS, httpx.SuccessOptions{Data: result})
}
