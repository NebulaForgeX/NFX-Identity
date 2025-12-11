package middleware

import (
	"nfxid/pkgs/netx/httpresp"
	"nfxid/pkgs/tokenx"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// AccessTokenMiddleware 验证 Access Token 的中间件
func AccessTokenMiddleware(tokenx *tokenx.Tokenx) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return httpresp.Error(c, fiber.StatusUnauthorized, "Missing authorization header")
		}

		// 提取 token (Bearer <token>)
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return httpresp.Error(c, fiber.StatusUnauthorized, "Invalid authorization header format")
		}

		tokenString := parts[1]
		claims, err := tokenx.VerifyAccessToken(tokenString)
		if err != nil {
			return httpresp.Error(c, fiber.StatusUnauthorized, "Invalid or expired token: "+err.Error())
		}

		// 将 claims 存储到 context 中
		c.Locals("user_id", claims.UserID)
		c.Locals("username", claims.Username)
		c.Locals("email", claims.Email)
		c.Locals("phone", claims.Phone)
		c.Locals("role_id", claims.RoleID)
		c.Locals("token_claims", claims)

		return c.Next()
	}
}
