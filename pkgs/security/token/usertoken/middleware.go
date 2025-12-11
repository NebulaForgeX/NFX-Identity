package usertoken

import (
	"nfxid/pkgs/netx/httpresp"
	"nfxid/pkgs/security/token"
	"strings"

	"github.com/gofiber/fiber/v2"
	webSocketVerify "github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

func AccessTokenMiddleware(verifier token.Verifier) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return httpresp.Error(c, fiber.StatusUnauthorized, "missing or invalid Authorization header")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := verifier.Verify(c.Context(), tokenStr)
		if err != nil {
			return httpresp.Error(c, fiber.StatusUnauthorized, "invalid or expired token")
		}

		// Inject into context
		userID, err := uuid.Parse(claims.Registered.Subject)
		if err != nil {
			return httpresp.Error(c, fiber.StatusUnauthorized, "invalid token")
		}
		c.Locals("userID", userID)
		return c.Next()
	}
}

func WebSocketJWTMiddleware(verifier token.Verifier) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !webSocketVerify.IsWebSocketUpgrade(c) {
			return fiber.ErrUpgradeRequired
		}
		// ① 先看 Authorization 头
		tokenStr := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		// ② 再看子协议 jwt,<token>
		if tokenStr == "" {
			if p := c.Get("Sec-WebSocket-Protocol"); strings.HasPrefix(p, "jwt,") {
				tokenStr = strings.TrimSpace(strings.TrimPrefix(p, "jwt,"))
			}
		}
		if tokenStr == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing token")
		}
		claims, err := verifier.Verify(c.Context(), tokenStr)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}

		userID, err := uuid.Parse(claims.Registered.Subject)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}
		c.Locals("userID", userID)

		return c.Next()
	}
}
