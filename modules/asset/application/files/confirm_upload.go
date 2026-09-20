package files

import (
	"context"

	asseterrs "nfxidentity/errors/src/asset"
	flsQuery "nfxidentity/modules/asset/query/files"

	"github.com/google/uuid"
)

type ConfirmUploadInput struct {
	AccountID uuid.UUID
	FileID    uuid.UUID
}

type ConfirmUploadOutput struct {
	File *flsQuery.FileVO
}

func (s *Service) ConfirmUpload(ctx context.Context, in ConfirmUploadInput) (*ConfirmUploadOutput, error) {
	row, err := s.fileRepo.Get.ByID(ctx, in.FileID)
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
	if err := s.fileRepo.Update.Generic(ctx, row); err != nil {
		return nil, err
	}
	vo, err := s.fileQuery.Single.ByID(ctx, in.FileID)
	if err != nil {
		return nil, err
	}
	return &ConfirmUploadOutput{File: vo}, nil
}
