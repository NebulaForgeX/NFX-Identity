package media

import (
	"context"

	authconn "nfxidentity/connections/auth"
	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/pkgs/fiberx"

	"github.com/google/uuid"
)

func EnsureOwnedProfile(ctx context.Context, authClient *authconn.Client, accountID uuid.UUID) error {
	if authClient == nil || authClient.Account == nil {
		return authErr.ErrProfileNotOwned
	}
	profileID, ok := fiberx.ProfileIDFromContext(ctx)
	if !ok || profileID == uuid.Nil {
		return authErr.ErrProfileNotOwned
	}
	scope, ok := fiberx.ProfileScopeFromContext(ctx)
	if !ok || scope == "" {
		return authErr.ErrProfileNotOwned
	}
	allowed, err := authClient.Account.EnsureOwnedProfile(ctx, accountID, profileID, scope)
	if err != nil {
		return err
	}
	if !allowed {
		return authErr.ErrProfileNotOwned
	}
	return nil
}
