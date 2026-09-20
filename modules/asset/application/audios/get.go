package audios

import (
	"context"

	audsQuery "nfxidentity/modules/asset/query/audios"

	"github.com/google/uuid"
)

type GetInput struct {
	ID uuid.UUID
}

func (s *Service) Get(ctx context.Context, in GetInput) (*audsQuery.AudioVO, error) {
	return s.audioQuery.Single.ByID(ctx, in.ID)
}

func (s *Service) GetMany(ctx context.Context, ids []string) ([]audsQuery.AudioVO, error) {
	out := make([]audsQuery.AudioVO, 0, len(ids))
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
