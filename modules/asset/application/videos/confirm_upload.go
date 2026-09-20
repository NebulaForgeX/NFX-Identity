package videos

import (
	"context"

	asseterrs "nfxidentity/errors/src/asset"
	vidsQuery "nfxidentity/modules/asset/query/videos"

	"github.com/google/uuid"
)

type ConfirmUploadInput struct {
	AccountID uuid.UUID
	VideoID   uuid.UUID
}

type ConfirmUploadOutput struct {
	Video *vidsQuery.VideoVO
}

func (s *Service) ConfirmUpload(ctx context.Context, in ConfirmUploadInput) (*ConfirmUploadOutput, error) {
	row, err := s.videoRepo.Get.ByID(ctx, in.VideoID)
	if err != nil {
		return nil, err
	}
	if row.UploaderID() != in.AccountID {
		return nil, asseterrs.ErrAssetNotFound
	}
	st, err := s.storage.Stat(ctx, row.FilePath())
	if err != nil {
		return nil, asseterrs.ErrObjectMissing.WithCause(err)
	}
	if err := row.UpdateFileMeta("", st.Size, st.ContentType); err != nil {
		return nil, err
	}
	if err := s.videoRepo.Update.Generic(ctx, row); err != nil {
		return nil, err
	}
	vo, err := s.videoQuery.Single.ByID(ctx, in.VideoID)
	if err != nil {
		return nil, err
	}
	return &ConfirmUploadOutput{Video: vo}, nil
}
