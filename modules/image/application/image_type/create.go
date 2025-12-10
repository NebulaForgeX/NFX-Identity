package image_type

import (
	"context"
	imageTypeDomain "nebulaid/modules/image/domain/image_type"
)

type CreateImageTypeCmd struct {
	Key         string
	Description *string
	MaxWidth    *int
	MaxHeight   *int
	AspectRatio *string
	IsSystem    bool
}

func (s *Service) CreateImageType(ctx context.Context, cmd CreateImageTypeCmd) (*imageTypeDomain.ImageType, error) {
	editable := imageTypeDomain.ImageTypeEditable{
		Key:         cmd.Key,
		Description: cmd.Description,
		MaxWidth:    cmd.MaxWidth,
		MaxHeight:   cmd.MaxHeight,
		AspectRatio: cmd.AspectRatio,
		IsSystem:    cmd.IsSystem,
	}

	it := imageTypeDomain.NewImageType(editable)
	if err := it.Validate(); err != nil {
		return nil, err
	}

	if err := s.imageTypeRepo.Create(ctx, it); err != nil {
		return nil, err
	}

	return it, nil
}
