package handler

import (
	"context"

	imageTypeApp "nebulaid/modules/image/application/image_type"
	imageTypeAppQueries "nebulaid/modules/image/application/image_type/queries"
	"nebulaid/modules/image/interfaces/grpc/mapper"
	"nebulaid/pkgs/logx"
	imagetypepb "nebulaid/protos/gen/image/image_type"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ImageTypeHandler struct {
	imagetypepb.UnimplementedImageTypeServiceServer
	imageTypeAppSvc *imageTypeApp.Service
}

func NewImageTypeHandler(imageTypeAppSvc *imageTypeApp.Service) *ImageTypeHandler {
	return &ImageTypeHandler{
		imageTypeAppSvc: imageTypeAppSvc,
	}
}

// GetImageTypeByID 根据ID获取图片类型
func (h *ImageTypeHandler) GetImageTypeByID(ctx context.Context, req *imagetypepb.GetImageTypeByIDRequest) (*imagetypepb.GetImageTypeByIDResponse, error) {
	id, err := uuid.Parse(req.TypeId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid type_id: %v", err)
	}

	imageTypeView, err := h.imageTypeAppSvc.GetImageType(ctx, id)
	if err != nil {
		logx.S().Errorf("failed to get image type by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "image type not found: %v", err)
	}

	imageType := mapper.ImageTypeViewToProto(&imageTypeView)
	return &imagetypepb.GetImageTypeByIDResponse{ImageType: imageType}, nil
}

// GetImageTypeByKey 根据Key获取图片类型
func (h *ImageTypeHandler) GetImageTypeByKey(ctx context.Context, req *imagetypepb.GetImageTypeByKeyRequest) (*imagetypepb.GetImageTypeByKeyResponse, error) {
	imageTypeView, err := h.imageTypeAppSvc.GetImageTypeByKey(ctx, req.Key)
	if err != nil {
		logx.S().Errorf("failed to get image type by key: %v", err)
		return nil, status.Errorf(codes.NotFound, "image type not found: %v", err)
	}

	imageType := mapper.ImageTypeViewToProto(&imageTypeView)
	return &imagetypepb.GetImageTypeByKeyResponse{ImageType: imageType}, nil
}

// BatchGetImageTypes 批量获取图片类型
func (h *ImageTypeHandler) BatchGetImageTypes(ctx context.Context, req *imagetypepb.BatchGetImageTypesRequest) (*imagetypepb.BatchGetImageTypesResponse, error) {
	typeIDs := make([]uuid.UUID, 0, len(req.TypeIds))
	for _, idStr := range req.TypeIds {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		typeIDs = append(typeIDs, id)
	}

	imageTypes := make([]*imagetypepb.ImageType, 0, len(typeIDs))
	errorById := make(map[string]string)

	for _, typeID := range typeIDs {
		imageTypeView, err := h.imageTypeAppSvc.GetImageType(ctx, typeID)
		if err != nil {
			errorById[typeID.String()] = err.Error()
			continue
		}
		imageType := mapper.ImageTypeViewToProto(&imageTypeView)
		imageTypes = append(imageTypes, imageType)
	}

	return &imagetypepb.BatchGetImageTypesResponse{
		ImageTypes: imageTypes,
		ErrorById:  errorById,
	}, nil
}

// GetAllImageTypes 获取所有图片类型列表
func (h *ImageTypeHandler) GetAllImageTypes(ctx context.Context, req *imagetypepb.GetAllImageTypesRequest) (*imagetypepb.GetAllImageTypesResponse, error) {
	listQuery := imageTypeAppQueries.ImageTypeListQuery{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
	}

	if req.Search != nil {
		listQuery.Search = *req.Search
	}
	if req.IsSystem != nil {
		isSystem := *req.IsSystem
		listQuery.IsSystem = &isSystem
	}

	result, err := h.imageTypeAppSvc.GetImageTypeList(ctx, listQuery)
	if err != nil {
		logx.S().Errorf("failed to get all image types: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get image types: %v", err)
	}

	imageTypes := make([]*imagetypepb.ImageType, len(result.Items))
	for i, imageTypeView := range result.Items {
		imageTypes[i] = mapper.ImageTypeViewToProto(&imageTypeView)
	}

	return &imagetypepb.GetAllImageTypesResponse{
		ImageTypes: imageTypes,
		Total:      int32(result.Total),
	}, nil
}
