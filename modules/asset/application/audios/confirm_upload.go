package audios

import (
	"context"

	asseterrs "nfxidentity/errors/src/asset"
	audsQuery "nfxidentity/modules/asset/query/audios"

	"github.com/google/uuid"
)

type ConfirmUploadInput struct {
	AccountID uuid.UUID
	AudioID   uuid.UUID
}

type ConfirmUploadOutput struct {
	Audio *audsQuery.AudioVO
}

func (s *Service) ConfirmUpload(ctx context.Context, in ConfirmUploadInput) (*ConfirmUploadOutput, error) {
	row, err := s.audioRepo.Get.ByID(ctx, in.AudioID)
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
	if err := s.audioRepo.Update.Generic(ctx, row); err != nil {
		return nil, err
	}
	vo, err := s.audioQuery.Single.ByID(ctx, in.AudioID)
	if err != nil {
		return nil, err
	}
	return &ConfirmUploadOutput{Audio: vo}, nil
}
