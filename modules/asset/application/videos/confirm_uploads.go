package videos

import (
	"context"

	vidsQuery "nfxidentity/modules/asset/query/videos"

	"github.com/google/uuid"
)

type ConfirmUploadsInput struct {
	AccountID uuid.UUID
	VideoIDs  []uuid.UUID
}

type ConfirmUploadsOutput struct {
	Videos []*vidsQuery.VideoVO
}

func (s *Service) ConfirmUploads(ctx context.Context, in ConfirmUploadsInput) (*ConfirmUploadsOutput, error) {
	rows := make([]*vidsQuery.VideoVO, 0, len(in.VideoIDs))
	for _, id := range in.VideoIDs {
		out, err := s.ConfirmUpload(ctx, ConfirmUploadInput{AccountID: in.AccountID, VideoID: id})
		if err != nil {
			return nil, err
		}
		rows = append(rows, out.Video)
	}
	return &ConfirmUploadsOutput{Videos: rows}, nil
}
