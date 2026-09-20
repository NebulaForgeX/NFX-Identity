package files

import (
	"context"

	flsQuery "nfxidentity/modules/asset/query/files"

	"github.com/google/uuid"
)

type ConfirmUploadsInput struct {
	AccountID uuid.UUID
	FileIDs   []uuid.UUID
}

type ConfirmUploadsOutput struct {
	Files []*flsQuery.FileVO
}

func (s *Service) ConfirmUploads(ctx context.Context, in ConfirmUploadsInput) (*ConfirmUploadsOutput, error) {
	rows := make([]*flsQuery.FileVO, 0, len(in.FileIDs))
	for _, id := range in.FileIDs {
		out, err := s.ConfirmUpload(ctx, ConfirmUploadInput{AccountID: in.AccountID, FileID: id})
		if err != nil {
			return nil, err
		}
		rows = append(rows, out.File)
	}
	return &ConfirmUploadsOutput{Files: rows}, nil
}
