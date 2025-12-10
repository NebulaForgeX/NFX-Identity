package image_type

import (
	"context"

	"github.com/google/uuid"
)

type DeleteImageTypeCmd struct {
	ID uuid.UUID
}

func (s *Service) DeleteImageType(ctx context.Context, cmd DeleteImageTypeCmd) error {
	return s.imageTypeRepo.Delete(ctx, cmd.ID)
}

