package videos

import (
	"context"

	asseterrs "nfxidentity/errors/src/asset"

	"github.com/google/uuid"
)

type DeleteInput struct {
	AccountID uuid.UUID
	VideoID   uuid.UUID
}

func (s *Service) Delete(ctx context.Context, in DeleteInput) error {
	img, err := s.videoRepo.Get.ByID(ctx, in.VideoID)
	if err != nil {
		return err
	}
	if img.UploaderID() != in.AccountID {
		return asseterrs.ErrInvalidUploaderID
	}
	_ = s.storage.Remove(ctx, img.FilePath())
	if err := img.Delete(); err != nil {
		return err
	}
	return s.videoRepo.Update.Generic(ctx, img)
}
