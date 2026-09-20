package audios

import (
	"context"

	asseterrs "nfxidentity/errors/src/asset"

	"github.com/google/uuid"
)

type DeleteInput struct {
	AccountID uuid.UUID
	AudioID   uuid.UUID
}

func (s *Service) Delete(ctx context.Context, in DeleteInput) error {
	img, err := s.audioRepo.Get.ByID(ctx, in.AudioID)
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
	return s.audioRepo.Update.Generic(ctx, img)
}
