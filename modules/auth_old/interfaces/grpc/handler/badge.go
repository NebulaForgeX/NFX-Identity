package handler

import (
	"context"

	badgeApp "nfxid/modules/auth/application/badge"
	badgeDomain "nfxid/modules/auth/domain/badge"
	"nfxid/modules/auth/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	badgepb "nfxid/protos/gen/auth/badge"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BadgeHandler struct {
	badgepb.UnimplementedBadgeServiceServer
	badgeAppSvc *badgeApp.Service
}

func NewBadgeHandler(badgeAppSvc *badgeApp.Service) *BadgeHandler {
	return &BadgeHandler{badgeAppSvc: badgeAppSvc}
}

// GetBadgeByID 根据ID获取徽章
func (h *BadgeHandler) GetBadgeByID(ctx context.Context, req *badgepb.GetBadgeByIDRequest) (*badgepb.GetBadgeByIDResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid badge_id: %v", err)
	}

	badgeView, err := h.badgeAppSvc.GetBadge(ctx, id)
	if err != nil {
		logx.S().Errorf("failed to get badge by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "badge not found: %v", err)
	}

	badge := mapper.BadgeViewToProto(&badgeView)
	return &badgepb.GetBadgeByIDResponse{Badge: badge}, nil
}

// GetBadgeByName 根据名称获取徽章
func (h *BadgeHandler) GetBadgeByName(ctx context.Context, req *badgepb.GetBadgeByNameRequest) (*badgepb.GetBadgeByNameResponse, error) {
	badgeView, err := h.badgeAppSvc.GetBadgeByName(ctx, req.Name)
	if err != nil {
		logx.S().Errorf("failed to get badge by name: %v", err)
		return nil, status.Errorf(codes.NotFound, "badge not found: %v", err)
	}

	badge := mapper.BadgeViewToProto(&badgeView)
	return &badgepb.GetBadgeByNameResponse{Badge: badge}, nil
}

// GetAllBadges 获取所有徽章列表
func (h *BadgeHandler) GetAllBadges(ctx context.Context, req *badgepb.GetAllBadgesRequest) (*badgepb.GetAllBadgesResponse, error) {
	listQuery := badgeDomain.ListQuery{}

	// Set pagination (convert Page/PageSize to Offset/Limit)
	if req.Page > 0 && req.PageSize > 0 {
		listQuery.Offset = int((req.Page - 1) * req.PageSize)
		listQuery.Limit = int(req.PageSize)
	}

	if req.Search != nil {
		search := *req.Search
		listQuery.Search = &search
	}
	if req.Category != nil {
		category := *req.Category
		listQuery.Category = &category
	}
	if req.IsSystem != nil {
		isSystem := *req.IsSystem
		listQuery.IsSystem = &isSystem
	}

	result, err := h.badgeAppSvc.GetBadgeList(ctx, listQuery)
	if err != nil {
		logx.S().Errorf("failed to get all badges: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get badges: %v", err)
	}

	badges := make([]*badgepb.Badge, len(result.Items))
	for i, badgeView := range result.Items {
		badges[i] = mapper.BadgeViewToProto(&badgeView)
	}

	return &badgepb.GetAllBadgesResponse{
		Badges: badges,
		Total:  int32(result.Total),
	}, nil
}
