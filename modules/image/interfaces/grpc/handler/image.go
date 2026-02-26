package handler

import (
	"context"

	imageApp "nfxidentity/modules/image/application/images"
	imageCommands "nfxidentity/modules/image/application/images/commands"
	"nfxidentity/modules/image/interfaces/grpc/mapper"
	"nfxidentity/pkgs/errx"
	imagepb "nfxidentity/protos/gen/image/image"

	"github.com/google/uuid"
)

type ImageHandler struct {
	imagepb.UnimplementedImageServiceServer
	appSvc *imageApp.Service
}

func NewImageHandler(appSvc *imageApp.Service) *ImageHandler {
	return &ImageHandler{appSvc: appSvc}
}

func (h *ImageHandler) GetImageByID(ctx context.Context, req *imagepb.GetImageByIDRequest) (*imagepb.GetImageByIDResponse, error) {
	if req == nil || req.Id == "" {
		return nil, errx.InvalidArg("ID_REQUIRED", "id is required")
	}
	imageID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	ro, err := h.appSvc.GetImage(ctx, imageID)
	if err != nil {
		return nil, err
	}
	return &imagepb.GetImageByIDResponse{Image: mapper.ImageROToProto(ro)}, nil
}

func (h *ImageHandler) GetImageByImageID(ctx context.Context, req *imagepb.GetImageByImageIDRequest) (*imagepb.GetImageByImageIDResponse, error) {
	if req == nil || req.ImageId == "" {
		return nil, errx.InvalidArg("IMAGE_ID_REQUIRED", "image_id is required")
	}
	imageID, err := uuid.Parse(req.ImageId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	ro, err := h.appSvc.GetImage(ctx, imageID)
	if err != nil {
		return nil, err
	}
	return &imagepb.GetImageByImageIDResponse{Image: mapper.ImageROToProto(ro)}, nil
}

func (h *ImageHandler) BatchGetImages(ctx context.Context, req *imagepb.BatchGetImagesRequest) (*imagepb.BatchGetImagesResponse, error) {
	if req == nil || len(req.Ids) == 0 {
		return &imagepb.BatchGetImagesResponse{Images: nil}, nil
	}
	images := make([]*imagepb.Image, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		imageID, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		ro, err := h.appSvc.GetImage(ctx, imageID)
		if err != nil {
			continue
		}
		images = append(images, mapper.ImageROToProto(ro))
	}
	return &imagepb.BatchGetImagesResponse{Images: images}, nil
}

func (h *ImageHandler) MoveImage(ctx context.Context, req *imagepb.MoveImageRequest) (*imagepb.MoveImageResponse, error) {
	if req == nil || req.Id == "" {
		return nil, errx.InvalidArg("ID_REQUIRED", "id is required")
	}
	if req.TargetType == "" {
		return nil, errx.InvalidArg("TARGET_TYPE_REQUIRED", "target_type is required")
	}
	if req.TargetType != "avatar" && req.TargetType != "background" {
		return nil, errx.InvalidArg("INVALID_TARGET_TYPE", "target_type must be 'avatar' or 'background'")
	}
	imageID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	ro, err := h.appSvc.MoveImage(ctx, imageID, req.TargetType)
	if err != nil {
		return nil, err
	}
	return &imagepb.MoveImageResponse{Image: mapper.ImageROToProto(ro)}, nil
}

func (h *ImageHandler) DeleteImage(ctx context.Context, req *imagepb.DeleteImageRequest) (*imagepb.DeleteImageResponse, error) {
	if req == nil || req.Id == "" {
		return nil, errx.InvalidArg("ID_REQUIRED", "id is required")
	}
	imageID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	cmd := imageCommands.DeleteImageCmd{ImageID: imageID}
	if err := h.appSvc.DeleteImage(ctx, cmd); err != nil {
		return nil, err
	}
	return &imagepb.DeleteImageResponse{}, nil
}

func (h *ImageHandler) ClearStorageData(ctx context.Context, req *imagepb.ClearStorageDataRequest) (*imagepb.ClearStorageDataResponse, error) {
	if err := h.appSvc.ClearStorageData(ctx); err != nil {
		msg := err.Error()
		return &imagepb.ClearStorageDataResponse{Success: false, ErrorMessage: &msg}, nil
	}
	return &imagepb.ClearStorageDataResponse{Success: true}, nil
}
