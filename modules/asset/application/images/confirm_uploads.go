package images

import (
	"context"

	imgsQuery "nfxidentity/modules/asset/query/images"

	"github.com/google/uuid"
)

type ConfirmUploadsInput struct {
	AccountID uuid.UUID
	ImageIDs  []uuid.UUID
}

type ConfirmUploadsOutput struct {
	Images []*imgsQuery.ImageVO
}

func (s *Service) ConfirmUploads(ctx context.Context, in ConfirmUploadsInput) (*ConfirmUploadsOutput, error) {
	images := make([]*imgsQuery.ImageVO, 0, len(in.ImageIDs))
	for _, imageID := range in.ImageIDs {
		out, err := s.ConfirmUpload(ctx, ConfirmUploadInput{
			AccountID: in.AccountID,
			ImageID:   imageID,
		})
		if err != nil {
			return nil, err
		}
		images = append(images, out.Image)
	}
	return &ConfirmUploadsOutput{Images: images}, nil
}
