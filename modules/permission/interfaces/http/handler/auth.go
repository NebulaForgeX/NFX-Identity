package handler

import (
	authApp "nfxid/modules/permission/application/auth"
	"nfxid/modules/permission/interfaces/http/dto/reqdto"
	"nfxid/modules/permission/interfaces/http/dto/respdto"
	"nfxid/pkgs/logx"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	appSvc *authApp.Service
}

func NewAuthHandler(appSvc *authApp.Service) *AuthHandler {
	return &AuthHandler{
		appSvc: appSvc,
	}
}

// Login 登录（支持用户名、邮箱、手机号密码登录，以及邮箱验证码登录）
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req reqdto.AuthLoginRequestDTO
	if err := c.BodyParser(&req); err != nil {
		logx.S().Errorf("❌ Failed to parse login request: %v", err)
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	logx.S().Infof("🔐 Login attempt for identifier: %s (type: %s)", req.Identifier, req.Type)

	cmd := req.ToLoginCmd()
	result, err := h.appSvc.Login(c.Context(), cmd)
	if err != nil {
		logx.S().Errorf("❌ Login failed for identifier %s: %v", req.Identifier, err)
		return httpresp.Error(c, fiber.StatusUnauthorized, "Login failed: "+err.Error())
	}

	logx.S().Infof("✅ Login successful for identifier: %s", req.Identifier)
	return httpresp.Success(c, fiber.StatusOK, "Login successful", httpresp.SuccessOptions{Data: respdto.LoginResponseToDTO(result)})
}

