package files

import (
	"context"

	"github.com/google/uuid"
)

type PrepareUploadItemInput struct {
	FileName string
	MimeType string
}

type PrepareUploadsInput struct {
	AccountID uuid.UUID
	Items     []PrepareUploadItemInput
}

type PrepareUploadsOutput struct {
	Results []PrepareUploadOutput
}

func (s *Service) PrepareUploads(ctx context.Context, in PrepareUploadsInput) (*PrepareUploadsOutput, error) {
	results := make([]PrepareUploadOutput, 0, len(in.Items))
	for _, item := range in.Items {
		out, err := s.PrepareUpload(ctx, PrepareUploadInput{
			AccountID: in.AccountID,
			FileName:  item.FileName,
			MimeType:  item.MimeType,
		})
		if err != nil {
			return nil, err
		}
		results = append(results, *out)
	}
	return &PrepareUploadsOutput{Results: results}, nil
}
