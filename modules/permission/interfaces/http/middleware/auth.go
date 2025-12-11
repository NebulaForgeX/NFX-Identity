package middleware

import (
	"nfxid/pkgs/tokenx"

	"github.com/gofiber/fiber/v2"
)

// AccessTokenMiddleware 验证访问令牌的中间件
func AccessTokenMiddleware(tokenx *tokenx.Tokenx) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is required",
			})
		}

		// 提取 token（支持 "Bearer <token>" 格式）
		token := authHeader
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			token = authHeader[7:]
		}

		// 验证 token
		claims, err := tokenx.VerifyAccessToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// 将用户信息存储到 locals 中
		c.Locals("user_id", claims.UserID)
		c.Locals("username", claims.Username)
		c.Locals("email", claims.Email)
		c.Locals("phone", claims.Phone)

		return c.Next()
	}
}

