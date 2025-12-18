package handler

import (
	"context"

	imageApp "nfxid/modules/image/application/image"
	imageAppQueries "nfxid/modules/image/application/image/queries"
	"nfxid/modules/image/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	imagepb "nfxid/protos/gen/image/image"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ImageHandler struct {
	imagepb.UnimplementedImageServiceServer
	imageAppSvc *imageApp.Service
}

func NewImageHandler(imageAppSvc *imageApp.Service) *ImageHandler {
	return &ImageHandler{
		imageAppSvc: imageAppSvc,
	}
}

// GetImageByID 根据ID获取图片
func (h *ImageHandler) GetImageByID(ctx context.Context, req *imagepb.GetImageByIDRequest) (*imagepb.GetImageByIDResponse, error) {
	id, err := uuid.Parse(req.ImageId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid image_id: %v", err)
	}

	imageView, err := h.imageAppSvc.GetImage(ctx, id)
	if err != nil {
		logx.S().Errorf("failed to get image by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "image not found: %v", err)
	}

	image := mapper.ImageViewToProto(&imageView)
	return &imagepb.GetImageByIDResponse{Image: image}, nil
}

// BatchGetImages 批量获取图片
func (h *ImageHandler) BatchGetImages(ctx context.Context, req *imagepb.BatchGetImagesRequest) (*imagepb.BatchGetImagesResponse, error) {
	imageIDs := make([]uuid.UUID, 0, len(req.ImageIds))
	for _, idStr := range req.ImageIds {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		imageIDs = append(imageIDs, id)
	}

	images := make([]*imagepb.Image, 0, len(imageIDs))
	errorById := make(map[string]string)

	for _, imageID := range imageIDs {
		imageView, err := h.imageAppSvc.GetImage(ctx, imageID)
		if err != nil {
			errorById[imageID.String()] = err.Error()
			continue
		}
		image := mapper.ImageViewToProto(&imageView)
		images = append(images, image)
	}

	return &imagepb.BatchGetImagesResponse{
		Images:    images,
		ErrorById: errorById,
	}, nil
}

// GetImagesByUserID 根据用户ID获取图片列表
func (h *ImageHandler) GetImagesByUserID(ctx context.Context, req *imagepb.GetImagesByUserIDRequest) (*imagepb.GetImagesByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	listQuery := imageAppQueries.ImageListQuery{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
		UserID:   &userID,
	}

	if req.IsPublic != nil {
		isPublic := *req.IsPublic
		listQuery.IsPublic = &isPublic
	}
	if req.TypeId != nil {
		typeID, err := uuid.Parse(*req.TypeId)
		if err == nil {
			listQuery.TypeID = &typeID
		}
	}

	domainQuery := listQuery.ToDomainListQuery()
	result, err := h.imageAppSvc.GetImageList(ctx, domainQuery)
	if err != nil {
		logx.S().Errorf("failed to get images by user_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get images: %v", err)
	}

	images := make([]*imagepb.Image, len(result.Items))
	for i, imageView := range result.Items {
		images[i] = mapper.ImageViewToProto(&imageView)
	}

	return &imagepb.GetImagesByUserIDResponse{
		Images: images,
		Total:  int32(result.Total),
	}, nil
}

// GetImagesByTypeID 根据类型ID获取图片列表
func (h *ImageHandler) GetImagesByTypeID(ctx context.Context, req *imagepb.GetImagesByTypeIDRequest) (*imagepb.GetImagesByTypeIDResponse, error) {
	typeID, err := uuid.Parse(req.TypeId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid type_id: %v", err)
	}

	listQuery := imageAppQueries.ImageListQuery{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
		TypeID:   &typeID,
	}

	if req.IsPublic != nil {
		isPublic := *req.IsPublic
		listQuery.IsPublic = &isPublic
	}

	domainQuery := listQuery.ToDomainListQuery()
	result, err := h.imageAppSvc.GetImageList(ctx, domainQuery)
	if err != nil {
		logx.S().Errorf("failed to get images by type_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get images: %v", err)
	}

	images := make([]*imagepb.Image, len(result.Items))
	for i, imageView := range result.Items {
		images[i] = mapper.ImageViewToProto(&imageView)
	}

	return &imagepb.GetImagesByTypeIDResponse{
		Images: images,
		Total:  int32(result.Total),
	}, nil
}

// GetImagesBySourceDomain 根据来源域获取图片列表
func (h *ImageHandler) GetImagesBySourceDomain(ctx context.Context, req *imagepb.GetImagesBySourceDomainRequest) (*imagepb.GetImagesBySourceDomainResponse, error) {
	sourceDomain := req.SourceDomain
	listQuery := imageAppQueries.ImageListQuery{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
	}
	listQuery.SourceDomain = &sourceDomain

	if req.IsPublic != nil {
		isPublic := *req.IsPublic
		listQuery.IsPublic = &isPublic
	}

	domainQuery := listQuery.ToDomainListQuery()
	result, err := h.imageAppSvc.GetImageList(ctx, domainQuery)
	if err != nil {
		logx.S().Errorf("failed to get images by source_domain: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get images: %v", err)
	}

	images := make([]*imagepb.Image, len(result.Items))
	for i, imageView := range result.Items {
		images[i] = mapper.ImageViewToProto(&imageView)
	}

	return &imagepb.GetImagesBySourceDomainResponse{
		Images: images,
		Total:  int32(result.Total),
	}, nil
}

// GetAllImages 获取所有图片列表
func (h *ImageHandler) GetAllImages(ctx context.Context, req *imagepb.GetAllImagesRequest) (*imagepb.GetAllImagesResponse, error) {
	listQuery := imageAppQueries.ImageListQuery{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
	}

	if req.UserId != nil {
		userID, err := uuid.Parse(*req.UserId)
		if err == nil {
			listQuery.UserID = &userID
		}
	}
	if req.TypeId != nil {
		typeID, err := uuid.Parse(*req.TypeId)
		if err == nil {
			listQuery.TypeID = &typeID
		}
	}
	if req.SourceDomain != nil {
		listQuery.SourceDomain = req.SourceDomain
	}
	if req.IsPublic != nil {
		isPublic := *req.IsPublic
		listQuery.IsPublic = &isPublic
	}
	if req.Search != nil {
		listQuery.Search = *req.Search
	}

	domainQuery := listQuery.ToDomainListQuery()
	result, err := h.imageAppSvc.GetImageList(ctx, domainQuery)
	if err != nil {
		logx.S().Errorf("failed to get all images: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get images: %v", err)
	}

	images := make([]*imagepb.Image, len(result.Items))
	for i, imageView := range result.Items {
		images[i] = mapper.ImageViewToProto(&imageView)
	}

	return &imagepb.GetAllImagesResponse{
		Images: images,
		Total:  int32(result.Total),
	}, nil
}

// SearchImagesByTags 根据标签搜索图片（暂时返回空结果，等待标签功能实现）
func (h *ImageHandler) SearchImagesByTags(ctx context.Context, req *imagepb.SearchImagesByTagsRequest) (*imagepb.SearchImagesByTagsResponse, error) {
	// TODO: 实现标签搜索功能
	logx.S().Warn("SearchImagesByTags not implemented yet")
	return &imagepb.SearchImagesByTagsResponse{
		Images: []*imagepb.Image{},
		Total:  0,
	}, nil
}
