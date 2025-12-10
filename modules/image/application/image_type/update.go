package image_type

import (
	"context"
	imageTypeDomain "nebulaid/modules/image/domain/image_type"

	"github.com/google/uuid"
)

type UpdateImageTypeCmd struct {
	ID          uuid.UUID
	Key         string
	Description *string
	MaxWidth    *int
	MaxHeight   *int
	AspectRatio *string
	IsSystem    bool
}

func (s *Service) UpdateImageType(ctx context.Context, cmd UpdateImageTypeCmd) error {
	it, err := s.imageTypeRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	editable := imageTypeDomain.ImageTypeEditable{
		Key:         cmd.Key,
		Description: cmd.Description,
		MaxWidth:    cmd.MaxWidth,
		MaxHeight:   cmd.MaxHeight,
		AspectRatio: cmd.AspectRatio,
		IsSystem:    cmd.IsSystem,
	}

	it.Update(editable)
	if err := it.Validate(); err != nil {
		return err
	}

	return s.imageTypeRepo.Update(ctx, it)
}
