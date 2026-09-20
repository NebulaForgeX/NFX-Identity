package audios

import (
	"context"

	audsQuery "nfxidentity/modules/asset/query/audios"

	"github.com/google/uuid"
)

type ConfirmUploadsInput struct {
	AccountID uuid.UUID
	AudioIDs  []uuid.UUID
}

type ConfirmUploadsOutput struct {
	Audios []*audsQuery.AudioVO
}

func (s *Service) ConfirmUploads(ctx context.Context, in ConfirmUploadsInput) (*ConfirmUploadsOutput, error) {
	rows := make([]*audsQuery.AudioVO, 0, len(in.AudioIDs))
	for _, id := range in.AudioIDs {
		out, err := s.ConfirmUpload(ctx, ConfirmUploadInput{AccountID: in.AccountID, AudioID: id})
		if err != nil {
			return nil, err
		}
		rows = append(rows, out.Audio)
	}
	return &ConfirmUploadsOutput{Audios: rows}, nil
}
