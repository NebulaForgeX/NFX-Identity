package handler

import (
	"nfxidentity/errors/src/sys"
	"nfxidentity/pkgs/errx"
	"nfxidentity/pkgs/fiberx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func wrap(c fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}
	if e := errx.AsError(err); e != nil {
		return fiberx.ErrorFromErrx(c, e)
	}
	return fiberx.ErrorFromErrx(c, sys.ErrInternal.WithCause(err))
}

func accountID(c fiber.Ctx) (uuid.UUID, error) {
	id, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok {
		return uuid.UUID{}, sys.ErrInvalidToken
	}
	return id, nil
}

func profileID(c fiber.Ctx) (uuid.UUID, error) {
	id, ok := fiberx.ProfileIDFromContext(c.Context())
	if !ok {
		return uuid.UUID{}, sys.ErrInvalidToken
	}
	return id, nil
}
