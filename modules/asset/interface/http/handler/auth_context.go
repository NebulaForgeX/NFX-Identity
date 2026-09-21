package handler

import (
	"nfxidentity/connections/auth"
	"nfxidentity/enums"
	asseterrs "nfxidentity/errors/src/asset"
	authErr "nfxidentity/errors/src/auth"
	sysErr "nfxidentity/errors/src/sys"
	"nfxidentity/pkgs/fiberx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func uploadContext(c fiber.Ctx) (accountID, profileID uuid.UUID, profileScope enums.AuthProfileScope, err error) {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok || accountID == uuid.Nil {
		return uuid.Nil, uuid.Nil, "", sysErr.ErrInvalidToken
	}
	profileID, ok = fiberx.ProfileIDFromContext(c.Context())
	if !ok || profileID == uuid.Nil {
		return uuid.Nil, uuid.Nil, "", asseterrs.ErrInvalidProfileID
	}
	profileScope, ok = fiberx.ProfileScopeFromContext(c.Context())
	if !ok {
		return uuid.Nil, uuid.Nil, "", sysErr.ErrInvalidToken
	}
	return accountID, profileID, profileScope, nil
}

func ensureOwnedProfile(c fiber.Ctx, authClient *auth.Client, accountID uuid.UUID) error {
	if authClient == nil || authClient.Account == nil {
		return authErr.ErrProfileNotOwned
	}
	profileID, ok := fiberx.ProfileIDFromContext(c.Context())
	if !ok || profileID == uuid.Nil {
		return authErr.ErrProfileNotOwned
	}
	scope, ok := fiberx.ProfileScopeFromContext(c.Context())
	if !ok {
		return authErr.ErrProfileNotOwned
	}
	allowed, err := authClient.Account.EnsureOwnedProfile(c.Context(), accountID, profileID, scope)
	if err != nil {
		return err
	}
	if !allowed {
		return authErr.ErrProfileNotOwned
	}
	return nil
}
