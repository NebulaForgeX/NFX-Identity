package handler

import (
	"errors"
	"strings"

	authApp "nfxid/modules/auth/application/auth"
	authCommands "nfxid/modules/auth/application/auth/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
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
func (h *AuthHandler) LoginByEmail(c *fiber.Ctx) error {
	if h.authSvc == nil {
		return httpresp.Error(c, fiber.StatusServiceUnavailable, "login not configured")
	}
	var req reqdto.LoginByEmailRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}
	if req.Email == "" {
		return httpresp.Error(c, fiber.StatusBadRequest, "email is required")
	}
	if req.Password == "" {
		return httpresp.Error(c, fiber.StatusBadRequest, "password is required")
	}

	// 获取客户端 IP
	clientIP := c.IP()
	if forwardedIP := c.Get("X-Forwarded-For"); forwardedIP != "" {
		clientIP = forwardedIP
	}

	cmd := authCommands.LoginByEmailCmd{
		Email:    strings.TrimSpace(req.Email),
		Password: req.Password,
		IP:       &clientIP,
	}

	res, err := h.authSvc.LoginByEmail(c.Context(), cmd)
	if err != nil {
		if errors.Is(err, authApp.ErrInvalidCredentials) {
			return httpresp.Error(c, fiber.StatusUnauthorized, "invalid email or password")
		}
		if errors.Is(err, authApp.ErrAccountLocked) {
			return httpresp.Error(c, fiber.StatusForbidden, "account is locked due to too many failed login attempts")
		}
		return httpresp.Error(c, fiber.StatusInternalServerError, "login failed")
	}

	return httpresp.Success(c, fiber.StatusOK, "Login successful", httpresp.SuccessOptions{
		Data: respdto.LoginResponseDTO{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
			ExpiresIn:    res.ExpiresIn,
			UserID:       res.UserID,
		},
	})
}

// LoginByPhone 处理 POST /auth/login/phone
func (h *AuthHandler) LoginByPhone(c *fiber.Ctx) error {
	if h.authSvc == nil {
		return httpresp.Error(c, fiber.StatusServiceUnavailable, "login not configured")
	}
	var req reqdto.LoginByPhoneRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}
	if req.Phone == "" {
		return httpresp.Error(c, fiber.StatusBadRequest, "phone is required")
	}
	if req.CountryCode == "" {
		return httpresp.Error(c, fiber.StatusBadRequest, "country_code is required")
	}
	if req.Password == "" {
		return httpresp.Error(c, fiber.StatusBadRequest, "password is required")
	}

	// 获取客户端 IP
	clientIP := c.IP()
	if forwardedIP := c.Get("X-Forwarded-For"); forwardedIP != "" {
		clientIP = forwardedIP
	}

	cmd := authCommands.LoginByPhoneCmd{
		CountryCode: strings.TrimSpace(req.CountryCode),
		Phone:       strings.TrimSpace(req.Phone),
		Password:    req.Password,
		IP:          &clientIP,
	}

	res, err := h.authSvc.LoginByPhone(c.Context(), cmd)
	if err != nil {
		if errors.Is(err, authApp.ErrInvalidCredentials) {
			return httpresp.Error(c, fiber.StatusUnauthorized, "invalid phone or password")
		}
		if errors.Is(err, authApp.ErrAccountLocked) {
			return httpresp.Error(c, fiber.StatusForbidden, "account is locked due to too many failed login attempts")
		}
		return httpresp.Error(c, fiber.StatusInternalServerError, "login failed")
	}

	return httpresp.Success(c, fiber.StatusOK, "Login successful", httpresp.SuccessOptions{
		Data: respdto.LoginResponseDTO{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
			ExpiresIn:    res.ExpiresIn,
			UserID:       res.UserID,
		},
	})
}

// Refresh 处理 POST /auth/refresh
func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	if h.authSvc == nil {
		return httpresp.Error(c, fiber.StatusServiceUnavailable, "refresh not configured")
	}
	var req reqdto.RefreshRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}
	if req.RefreshToken == "" {
		return httpresp.Error(c, fiber.StatusBadRequest, "refresh_token is required")
	}

	// 获取客户端 IP
	clientIP := c.IP()
	if forwardedIP := c.Get("X-Forwarded-For"); forwardedIP != "" {
		clientIP = forwardedIP
	}

	res, err := h.authSvc.Refresh(c.Context(), req.RefreshToken, &clientIP)
	if err != nil {
		if errors.Is(err, authApp.ErrInvalidRefreshToken) {
			return httpresp.Error(c, fiber.StatusUnauthorized, "invalid or expired refresh token")
		}
		if errors.Is(err, authApp.ErrAccountLocked) {
			return httpresp.Error(c, fiber.StatusForbidden, "account is locked")
		}
		return httpresp.Error(c, fiber.StatusInternalServerError, "refresh failed")
	}

	return httpresp.Success(c, fiber.StatusOK, "Token refreshed", httpresp.SuccessOptions{
		Data: respdto.RefreshResponseDTO{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
			ExpiresIn:    res.ExpiresIn,
		},
	})
}
