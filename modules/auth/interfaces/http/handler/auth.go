package handler

import (
	"errors"
	"strings"

	authApp "nfxid/modules/auth/application/auth"
	authCommands "nfxid/modules/auth/application/auth/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

// AuthHandler 认证 HTTP 适配器（登录、刷新 Token）
type AuthHandler struct {
	authSvc *authApp.Service
}

// NewAuthHandler 创建 AuthHandler
func NewAuthHandler(authSvc *authApp.Service) *AuthHandler {
	return &AuthHandler{authSvc: authSvc}
}

// LoginByEmail 处理 POST /auth/login/email
func (h *AuthHandler) LoginByEmail(c fiber.Ctx) error {
	if h.authSvc == nil {
		return fiberx.ErrorFromErrx(c, errx.Internal("SERVICE_UNAVAILABLE", "login not configured"))
	}
	var req reqdto.LoginByEmailRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	if req.Email == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "email is required"))
	}
	if req.Password == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "password is required"))
	}

	clientIP := fiberx.GetIP(c)
	cmd := authCommands.LoginByEmailCmd{
		Email:    strings.TrimSpace(req.Email),
		Password: req.Password,
		IP:       &clientIP,
	}

	res, err := h.authSvc.LoginByEmail(c.Context(), cmd)
	if err != nil {
		if errors.Is(err, authApp.ErrInvalidCredentials) {
			return fiberx.ErrorFromErrx(c, errx.Unauthorized("INVALID_CREDENTIALS", "invalid email or password"))
		}
		if errors.Is(err, authApp.ErrAccountLocked) {
			return fiberx.ErrorFromErrx(c, errx.Forbidden("ACCOUNT_LOCKED", "account is locked due to too many failed login attempts"))
		}
		return err
	}

	return fiberx.OK(c, "Login successful", httpx.SuccessOptions{
		Data: respdto.LoginResponseDTO{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
			ExpiresIn:    res.ExpiresIn,
			UserID:       res.UserID,
		},
	})
}

// LoginByPhone 处理 POST /auth/login/phone
func (h *AuthHandler) LoginByPhone(c fiber.Ctx) error {
	if h.authSvc == nil {
		return fiberx.ErrorFromErrx(c, errx.Internal("SERVICE_UNAVAILABLE", "login not configured"))
	}
	var req reqdto.LoginByPhoneRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	if req.Phone == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "phone is required"))
	}
	if req.CountryCode == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "country_code is required"))
	}
	if req.Password == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "password is required"))
	}

	clientIP := fiberx.GetIP(c)
	cmd := authCommands.LoginByPhoneCmd{
		CountryCode: strings.TrimSpace(req.CountryCode),
		Phone:       strings.TrimSpace(req.Phone),
		Password:    req.Password,
		IP:          &clientIP,
	}

	res, err := h.authSvc.LoginByPhone(c.Context(), cmd)
	if err != nil {
		if errors.Is(err, authApp.ErrInvalidCredentials) {
			return fiberx.ErrorFromErrx(c, errx.Unauthorized("INVALID_CREDENTIALS", "invalid phone or password"))
		}
		if errors.Is(err, authApp.ErrAccountLocked) {
			return fiberx.ErrorFromErrx(c, errx.Forbidden("ACCOUNT_LOCKED", "account is locked due to too many failed login attempts"))
		}
		return err
	}

	return fiberx.OK(c, "Login successful", httpx.SuccessOptions{
		Data: respdto.LoginResponseDTO{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
			ExpiresIn:    res.ExpiresIn,
			UserID:       res.UserID,
		},
	})
}

// Refresh 处理 POST /auth/refresh
func (h *AuthHandler) Refresh(c fiber.Ctx) error {
	if h.authSvc == nil {
		return fiberx.ErrorFromErrx(c, errx.Internal("SERVICE_UNAVAILABLE", "refresh not configured"))
	}
	var req reqdto.RefreshRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	if req.RefreshToken == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "refresh_token is required"))
	}

	clientIP := fiberx.GetIP(c)
	res, err := h.authSvc.Refresh(c.Context(), req.RefreshToken, &clientIP)
	if err != nil {
		if errors.Is(err, authApp.ErrInvalidRefreshToken) {
			return fiberx.ErrorFromErrx(c, errx.Unauthorized("INVALID_REFRESH_TOKEN", "invalid or expired refresh token"))
		}
		if errors.Is(err, authApp.ErrAccountLocked) {
			return fiberx.ErrorFromErrx(c, errx.Forbidden("ACCOUNT_LOCKED", "account is locked"))
		}
		return err
	}

	return fiberx.OK(c, "Token refreshed", httpx.SuccessOptions{
		Data: respdto.RefreshResponseDTO{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
			ExpiresIn:    res.ExpiresIn,
		},
	})
}

// SendVerificationCode 处理 POST /auth/send-verification-code
func (h *AuthHandler) SendVerificationCode(c fiber.Ctx) error {
	if h.authSvc == nil {
		return fiberx.ErrorFromErrx(c, errx.Internal("SERVICE_UNAVAILABLE", "verification code service not configured"))
	}
	var req reqdto.SendVerificationCodeRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	if req.Email == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "email is required"))
	}

	cmd := authCommands.SendVerificationCodeCmd{
		Email: strings.TrimSpace(req.Email),
	}

	err := h.authSvc.SendVerificationCode(c.Context(), cmd)
	if err != nil {
		if errors.Is(err, authApp.ErrEmailAlreadyVerified) {
			return fiberx.ErrorFromErrx(c, errx.Conflict("EMAIL_ALREADY_VERIFIED", "email already verified, please login"))
		}
		return err
	}

	return fiberx.OK(c, "Verification code sent successfully")
}

// Signup 处理 POST /auth/signup
func (h *AuthHandler) Signup(c fiber.Ctx) error {
	if h.authSvc == nil {
		return fiberx.ErrorFromErrx(c, errx.Internal("SERVICE_UNAVAILABLE", "signup not configured"))
	}
	var req reqdto.SignupRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}
	if req.Email == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "email is required"))
	}
	if req.Password == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "password is required"))
	}
	if req.VerificationCode == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "verification code is required"))
	}

	cmd := authCommands.SignupCmd{
		Email:            strings.TrimSpace(req.Email),
		Password:         req.Password,
		VerificationCode: strings.TrimSpace(req.VerificationCode),
	}

	res, err := h.authSvc.Signup(c.Context(), cmd)
	if err != nil {
		if errors.Is(err, authApp.ErrEmailAlreadyVerified) {
			return fiberx.ErrorFromErrx(c, errx.Conflict("EMAIL_ALREADY_VERIFIED", "email already verified, please login"))
		}
		if errors.Is(err, authApp.ErrInvalidVerificationCode) {
			return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_VERIFICATION_CODE", "invalid or expired verification code"))
		}
		if strings.Contains(err.Error(), "already has credentials") {
			return fiberx.ErrorFromErrx(c, errx.Conflict("USER_HAS_CREDENTIALS", "user already has credentials, please login"))
		}
		return err
	}

	return fiberx.Created(c, "Signup successful", httpx.SuccessOptions{
		Data: respdto.SignupResponseDTO{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
			ExpiresIn:    res.ExpiresIn,
			UserID:       res.UserID,
		},
	})
}

// fiber:context-methods migrated
