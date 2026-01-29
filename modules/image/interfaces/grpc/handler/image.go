package handler

import (
	"context"

	imageApp "nfxid/modules/image/application/images"
	imageCommands "nfxid/modules/image/application/images/commands"
	"nfxid/modules/image/interfaces/grpc/mapper"
	imagepb "nfxid/protos/gen/image/image"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ImageHandler 实现 image.ImageService gRPC 服务
type ImageHandler struct {
	imagepb.UnimplementedImageServiceServer
	appSvc *imageApp.Service
}

// NewImageHandler 创建 Image gRPC handler
func NewImageHandler(appSvc *imageApp.Service) *ImageHandler {
	return &ImageHandler{appSvc: appSvc}
}

// GetImageByID 根据 ID 获取图片
func (h *ImageHandler) GetImageByID(ctx context.Context, req *imagepb.GetImageByIDRequest) (*imagepb.GetImageByIDResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	imageID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid image id: "+err.Error())
	}
	ro, err := h.appSvc.GetImage(ctx, imageID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "image not found: "+err.Error())
	}
	return &imagepb.GetImageByIDResponse{Image: mapper.ImageROToProto(ro)}, nil
}

// GetImageByImageID 根据图片 ID 获取图片（与 GetImageByID 同义，兼容调用方）
func (h *ImageHandler) GetImageByImageID(ctx context.Context, req *imagepb.GetImageByImageIDRequest) (*imagepb.GetImageByImageIDResponse, error) {
	if req == nil || req.ImageId == "" {
		return nil, status.Error(codes.InvalidArgument, "image_id is required")
	}
	imageID, err := uuid.Parse(req.ImageId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid image_id: "+err.Error())
	}
	ro, err := h.appSvc.GetImage(ctx, imageID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "image not found: "+err.Error())
	}
	return &imagepb.GetImageByImageIDResponse{Image: mapper.ImageROToProto(ro)}, nil
}

// BatchGetImages 批量获取图片
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

// MoveImage 移动图片（从 tmp 移动到目标目录）
func (h *ImageHandler) MoveImage(ctx context.Context, req *imagepb.MoveImageRequest) (*imagepb.MoveImageResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	if req.TargetType == "" {
		return nil, status.Error(codes.InvalidArgument, "target_type is required")
	}
	if req.TargetType != "avatar" && req.TargetType != "background" {
		return nil, status.Error(codes.InvalidArgument, "target_type must be 'avatar' or 'background'")
	}

	imageID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid image id: "+err.Error())
	}

	ro, err := h.appSvc.MoveImage(ctx, imageID, req.TargetType)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to move image: "+err.Error())
	}

	return &imagepb.MoveImageResponse{Image: mapper.ImageROToProto(ro)}, nil
}

// DeleteImage 删除图片
func (h *ImageHandler) DeleteImage(ctx context.Context, req *imagepb.DeleteImageRequest) (*imagepb.DeleteImageResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	imageID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid image id: "+err.Error())
	}

	cmd := imageCommands.DeleteImageCmd{ImageID: imageID}
	if err := h.appSvc.DeleteImage(ctx, cmd); err != nil {
		return nil, status.Error(codes.Internal, "failed to delete image: "+err.Error())
	}

	return &imagepb.DeleteImageResponse{}, nil
}
