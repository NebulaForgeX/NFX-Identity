package image_types

import (
	"context"
	imageTypeCommands "nfxid/modules/image/application/image_types/commands"
	imageTypeDomain "nfxid/modules/image/domain/image_types"

	"github.com/google/uuid"
)

// CreateImageType 创建图片类型
func (s *Service) CreateImageType(ctx context.Context, cmd imageTypeCommands.CreateImageTypeCmd) (uuid.UUID, error) {
	// Create domain entity
	imageType, err := imageTypeDomain.NewImageType(imageTypeDomain.NewImageTypeParams{
		Key:         cmd.Key,
		Description: cmd.Description,
		MaxWidth:    cmd.MaxWidth,
		MaxHeight:   cmd.MaxHeight,
		AspectRatio: cmd.AspectRatio,
		IsSystem:    cmd.IsSystem,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.imageTypeRepo.Create.New(ctx, imageType); err != nil {
		return uuid.Nil, err
	}

	return imageType.ID(), nil
}
