package image_tags

import (
	"context"
	imageTagCommands "nfxid/modules/image/application/image_tags/commands"
	imageTagDomain "nfxid/modules/image/domain/image_tags"

	"github.com/google/uuid"
)

// CreateImageTag 创建图片标签
func (s *Service) CreateImageTag(ctx context.Context, cmd imageTagCommands.CreateImageTagCmd) (uuid.UUID, error) {
	// Create domain entity
	imageTag, err := imageTagDomain.NewImageTag(imageTagDomain.NewImageTagParams{
		ImageID:    cmd.ImageID,
		Tag:        cmd.Tag,
		Confidence: cmd.Confidence,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.imageTagRepo.Create.New(ctx, imageTag); err != nil {
		return uuid.Nil, err
	}

	return imageTag.ID(), nil
}
