package fiberx

import (
	"context"
	"nfxidentity/enums"
	"strings"

	"github.com/google/uuid"
)

type accountIDCtxKey struct{}
type profileIDCtxKey struct{}
type profileScopeCtxKey struct{}
type loginEmailCtxKey struct{}

func WithAccountID(ctx context.Context, accountID uuid.UUID) context.Context {
	return context.WithValue(ctx, accountIDCtxKey{}, accountID)
}

func AccountIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	v := ctx.Value(accountIDCtxKey{})
	if v == nil {
		return uuid.UUID{}, false
	}
	id, ok := v.(uuid.UUID)
	return id, ok
}

func WithProfileID(ctx context.Context, profileID uuid.UUID) context.Context {
	return context.WithValue(ctx, profileIDCtxKey{}, profileID)
}

func ProfileIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	v := ctx.Value(profileIDCtxKey{})
	if v == nil {
		return uuid.UUID{}, false
	}
	id, ok := v.(uuid.UUID)
	return id, ok
}

func WithProfileScope(ctx context.Context, profileScope enums.AuthProfileScope) context.Context {
	return context.WithValue(ctx, profileScopeCtxKey{}, profileScope)
}

func ProfileScopeFromContext(ctx context.Context) (enums.AuthProfileScope, bool) {
	v := ctx.Value(profileScopeCtxKey{})
	if v == nil {
		return "", false
	}
	scope, ok := v.(enums.AuthProfileScope)
	return scope, ok
}

func WithLoginEmail(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, loginEmailCtxKey{}, strings.TrimSpace(email))
}

func LoginEmailFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(loginEmailCtxKey{})
	if v == nil {
		return "", false
	}
	email, ok := v.(string)
	if !ok || strings.TrimSpace(email) == "" {
		return "", false
	}
	return strings.TrimSpace(email), true
}

func AccountProfileIDsFromContext(ctx context.Context) (uuid.UUID, uuid.UUID, bool) {
	accountID, ok := AccountIDFromContext(ctx)
	if !ok {
		return uuid.Nil, uuid.Nil, false
	}
	profileID, ok := ProfileIDFromContext(ctx)
	if !ok {
		return uuid.Nil, uuid.Nil, false
	}
	return accountID, profileID, true
}
