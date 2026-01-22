package image_tags

import (
	"context"
	imageTagCommands "nfxid/modules/image/application/image_tags/commands"
	imageTagDomain "nfxid/modules/image/domain/image_tags"
)

// UpdateImageTag 更新图片标签
func (s *Service) UpdateImageTag(ctx context.Context, cmd imageTagCommands.UpdateImageTagCmd) error {
	// Get domain entity
	imageTag, err := s.imageTagRepo.Get.ByID(ctx, cmd.ImageTagID)
	if err != nil {
		return err
	}

	// Get current state
	currentState := imageTagDomain.ImageTagState{
		ID:         imageTag.ID(),
		ImageID:    imageTag.ImageID(),
		Tag:        imageTag.Tag(),
		Confidence: imageTag.Confidence(),
		CreatedAt:  imageTag.CreatedAt(),
	}

	// Update fields if provided
	if cmd.Tag != "" {
		currentState.Tag = cmd.Tag
	}
	if cmd.Confidence != nil {
		currentState.Confidence = cmd.Confidence
	}

	// Create updated entity
	updatedImageTag := imageTagDomain.NewImageTagFromState(currentState)

	// Save to repository
	return s.imageTagRepo.Update.Generic(ctx, updatedImageTag)
}
