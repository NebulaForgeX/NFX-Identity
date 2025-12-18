package image

import (
	"context"

	"github.com/google/uuid"
)

type DeleteImageCmd struct {
	ID uuid.UUID
}

func (s *Service) DeleteImage(ctx context.Context, cmd DeleteImageCmd) error {
	img, err := s.imageRepo.Get.ByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	img.Delete()
	return s.imageRepo.Update.Generic(ctx, img)
}

type DeleteImageByStoragePathCmd struct {
	StoragePath string
}

func (s *Service) DeleteImageByStoragePath(ctx context.Context, cmd DeleteImageByStoragePathCmd) error {
	return s.imageRepo.Delete.ByStoragePath(ctx, cmd.StoragePath)
}
