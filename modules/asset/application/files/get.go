package files

import (
	"context"

	flsQuery "nfxidentity/modules/asset/query/files"

	"github.com/google/uuid"
)

type GetInput struct {
	ID uuid.UUID
}

func (s *Service) Get(ctx context.Context, in GetInput) (*flsQuery.FileVO, error) {
	return s.fileQuery.Single.ByID(ctx, in.ID)
}

func (s *Service) GetMany(ctx context.Context, ids []string) ([]flsQuery.FileVO, error) {
	out := make([]flsQuery.FileVO, 0, len(ids))
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
