package images

import (
	"context"
	"time"
	imageCommands "nfxid/modules/image/application/images/commands"
	imageDomain "nfxid/modules/image/domain/images"
)

// UpdateImage 更新图片
func (s *Service) UpdateImage(ctx context.Context, cmd imageCommands.UpdateImageCmd) error {
	// Get domain entity
	image, err := s.imageRepo.Get.ByID(ctx, cmd.ImageID)
	if err != nil {
		return err
	}

	// Get current state
	currentState := imageDomain.ImageState{
		ID:              image.ID(),
		TypeID:          image.TypeID(),
		UserID:          image.UserID(),
		TenantID:        image.TenantID(),
		AppID:           image.AppID(),
		SourceDomain:    image.SourceDomain(),
		Filename:        image.Filename(),
		OriginalFilename: image.OriginalFilename(),
		MimeType:        image.MimeType(),
		Size:            image.Size(),
		Width:           image.Width(),
		Height:          image.Height(),
		StoragePath:     image.StoragePath(),
		URL:             image.URL(),
		IsPublic:        image.IsPublic(),
		Metadata:        image.Metadata(),
		CreatedAt:       image.CreatedAt(),
		UpdatedAt:       image.UpdatedAt(),
		DeletedAt:       image.DeletedAt(),
	}

	// Update fields if provided
	if cmd.TypeID != nil {
		currentState.TypeID = cmd.TypeID
	}
	if cmd.UserID != nil {
		currentState.UserID = cmd.UserID
	}
	if cmd.TenantID != nil {
		currentState.TenantID = cmd.TenantID
	}
	if cmd.AppID != nil {
		currentState.AppID = cmd.AppID
	}
	if cmd.SourceDomain != nil {
		currentState.SourceDomain = cmd.SourceDomain
	}
	if cmd.Filename != nil {
		currentState.Filename = *cmd.Filename
	}
	if cmd.OriginalFilename != nil {
		currentState.OriginalFilename = *cmd.OriginalFilename
	}
	if cmd.MimeType != nil {
		currentState.MimeType = *cmd.MimeType
	}
	if cmd.Size != nil {
		currentState.Size = *cmd.Size
	}
	if cmd.Width != nil {
		currentState.Width = cmd.Width
	}
	if cmd.Height != nil {
		currentState.Height = cmd.Height
	}
	if cmd.StoragePath != nil {
		currentState.StoragePath = *cmd.StoragePath
	}
	if cmd.URL != nil {
		currentState.URL = cmd.URL
	}
	if cmd.IsPublic != nil {
		currentState.IsPublic = *cmd.IsPublic
	}
	if cmd.Metadata != nil {
		currentState.Metadata = cmd.Metadata
	}

	// Update timestamp
	currentState.UpdatedAt = time.Now().UTC()

	// Create updated entity
	updatedImage := imageDomain.NewImageFromState(currentState)

	// Save to repository
	return s.imageRepo.Update.Generic(ctx, updatedImage)
}

// UpdateImageURL 更新图片URL
func (s *Service) UpdateImageURL(ctx context.Context, cmd imageCommands.UpdateImageURLCmd) error {
	// Get domain entity
	image, err := s.imageRepo.Get.ByID(ctx, cmd.ImageID)
	if err != nil {
		return err
	}

	// Update URL using domain behavior
	if err := image.UpdateURL(cmd.URL); err != nil {
		return err
	}

	// Save to repository
	return s.imageRepo.Update.UpdateURL(ctx, cmd.ImageID, cmd.URL)
}

// UpdateImagePublic 更新图片公开状态
func (s *Service) UpdateImagePublic(ctx context.Context, cmd imageCommands.UpdateImagePublicCmd) error {
	// Get domain entity
	image, err := s.imageRepo.Get.ByID(ctx, cmd.ImageID)
	if err != nil {
		return err
	}

	// Update public status using domain behavior
	if err := image.UpdatePublic(cmd.IsPublic); err != nil {
		return err
	}

	// Save to repository
	return s.imageRepo.Update.UpdatePublic(ctx, cmd.ImageID, cmd.IsPublic)
}

