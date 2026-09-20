package images

import (
	"context"
	"time"

	asseterrs "nfxidentity/errors/src/asset"
)

type FileURLInput struct {
	ObjectKey string
}

func (s *Service) FileURL(ctx context.Context, in FileURLInput) (string, error) {
	u, err := s.storage.PresignGet(ctx, in.ObjectKey, time.Hour)
	if err != nil {
		return "", asseterrs.ErrPresignFailed.WithCause(err)
	}
	return u.String(), nil
}
