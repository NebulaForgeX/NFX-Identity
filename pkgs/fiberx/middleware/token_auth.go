package middleware

import (
	"strings"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

var (
	ErrInvalidAuthHeader = errx.Unauthorized("INVALID_AUTH_HEADER", "missing or invalid Authorization header")
	ErrInvalidToken      = errx.Unauthorized("INVALID_TOKEN", "invalid or expired token")
)

func TokenAuth(verifier token.Verifier) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return fiberx.ErrorFromErrx(c, ErrInvalidAuthHeader)
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := verifier.Verify(c.Context(), tokenStr)
		if err != nil {
			return fiberx.ErrorFromErrx(c, ErrInvalidToken.WithCause(err))
		}
		userID, err := uuid.Parse(claims.Registered.Subject)
		if err != nil {
			return fiberx.ErrorFromErrx(c, ErrInvalidToken.WithCause(err))
		}
		c.SetContext(fiberx.WithUserID(c.Context(), userID))
		return c.Next()
	}
}
