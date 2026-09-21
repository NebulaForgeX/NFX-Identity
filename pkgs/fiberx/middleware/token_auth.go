package middleware

import (
	"strings"

	"nfxidentity/constants"
	"nfxidentity/enums"
	"nfxidentity/errors/src/sys"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func TokenAuth(verifier token.Verifier) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return fiberx.ErrorFromErrx(c, sys.ErrInvalidAuthHeader)
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := verifier.Verify(c.Context(), tokenStr)
		if err != nil {
			return fiberx.ErrorFromErrx(c, sys.ErrInvalidToken.WithCause(err))
		}
		userID, err := uuid.Parse(claims.Registered.Subject)
		if err != nil {
			return fiberx.ErrorFromErrx(c, sys.ErrInvalidToken.WithCause(err))
		}
		ctx := fiberx.WithAccountID(c.Context(), userID)
		if raw, ok := claims.Raw["account_id"].(string); ok {
			if aid, parseErr := uuid.Parse(raw); parseErr == nil {
				ctx = fiberx.WithAccountID(ctx, aid)
			}
		}
		if raw, ok := claims.Raw["profile_id"].(string); ok {
			if pid, parseErr := uuid.Parse(raw); parseErr == nil {
				ctx = fiberx.WithProfileID(ctx, pid)
			}
		}
		if raw, ok := claims.Raw["email"].(string); ok && strings.TrimSpace(raw) != "" {
			ctx = fiberx.WithLoginEmail(ctx, raw)
		}

		rawScope := ""
		if raw, ok := claims.Raw["profile_scope"].(string); ok {
			rawScope = strings.TrimSpace(raw)
		}
		if rawScope != "" {
			profileScope := enums.AuthProfileScope(rawScope)
			if !constants.AuthProfileScope.Valid(profileScope) {
				return fiberx.ErrorFromErrx(c, sys.ErrTokenInvalidProfileScope.WithDetail("profileScope", rawScope))
			}
			ctx = fiberx.WithProfileScope(ctx, profileScope)
		}
		c.SetContext(ctx)
		return c.Next()
	}
}
