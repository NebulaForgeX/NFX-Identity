package videos

import (
	"context"

	vidsQuery "nfxidentity/modules/asset/query/videos"

	"github.com/google/uuid"
)

type GetInput struct {
	ID uuid.UUID
}

func (s *Service) Get(ctx context.Context, in GetInput) (*vidsQuery.VideoVO, error) {
	return s.videoQuery.Single.ByID(ctx, in.ID)
}

func (s *Service) GetMany(ctx context.Context, ids []string) ([]vidsQuery.VideoVO, error) {
	out := make([]vidsQuery.VideoVO, 0, len(ids))
	for _, raw := range ids {
		id, err := uuid.Parse(raw)
		if err != nil {
			continue
		}
		vo, err := s.Get(ctx, GetInput{ID: id})
		if err != nil {
			continue
		}
		out = append(out, *vo)
	}
	return out, nil
}
