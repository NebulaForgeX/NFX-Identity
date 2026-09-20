package images

import (
	"context"

	asseterrs "nfxidentity/errors/src/asset"
	imgsQuery "nfxidentity/modules/asset/query/images"

	"github.com/google/uuid"
)

type ConfirmUploadInput struct {
	AccountID uuid.UUID
	ImageID   uuid.UUID
}

type ConfirmUploadOutput struct {
	Image *imgsQuery.ImageVO
}

func (s *Service) ConfirmUpload(ctx context.Context, in ConfirmUploadInput) (*ConfirmUploadOutput, error) {
	img, err := s.imageRepo.Get.ByID(ctx, in.ImageID)
	if err != nil {
		return nil, err
	}
	if img.UploaderID() != in.AccountID {
		return nil, asseterrs.ErrAssetNotFound
	}
	st, err := s.storage.Stat(ctx, img.FilePath())
	if err != nil {
		return nil, asseterrs.ErrObjectMissing.WithCause(err)
	}
	if err := img.UpdateFileMeta("", st.Size, st.ContentType); err != nil {
		return nil, err
	}
	if err := s.imageRepo.Update.Generic(ctx, img); err != nil {
		return nil, err
	}
	vo, err := s.imageQuery.Single.ByID(ctx, in.ImageID)
	if err != nil {
		return nil, err
	}
	return &ConfirmUploadOutput{Image: vo}, nil
}
