package handler

import (
	"context"

	badgeApp "nfxid/modules/directory/application/badges"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	badgepb "nfxid/protos/gen/directory/badge"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BadgeHandler struct {
	badgepb.UnimplementedBadgeServiceServer
	badgeAppSvc *badgeApp.Service
}

func NewBadgeHandler(badgeAppSvc *badgeApp.Service) *BadgeHandler {
	return &BadgeHandler{
		badgeAppSvc: badgeAppSvc,
	}
}

// GetBadgeByID 根据ID获取徽章
func (h *BadgeHandler) GetBadgeByID(ctx context.Context, req *badgepb.GetBadgeByIDRequest) (*badgepb.GetBadgeByIDResponse, error) {
	badgeID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid badge_id: %v", err)
	}

	badgeView, err := h.badgeAppSvc.GetBadge(ctx, badgeID)
	if err != nil {
		logx.S().Errorf("failed to get badge by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "badge not found: %v", err)
	}

	badge := mapper.BadgeROToProto(&badgeView)
	return &badgepb.GetBadgeByIDResponse{Badge: badge}, nil
}

// GetBadgeByName 根据名称获取徽章
func (h *BadgeHandler) GetBadgeByName(ctx context.Context, req *badgepb.GetBadgeByNameRequest) (*badgepb.GetBadgeByNameResponse, error) {
	badgeView, err := h.badgeAppSvc.GetBadgeByName(ctx, req.Name)
	if err != nil {
		logx.S().Errorf("failed to get badge by name: %v", err)
		return nil, status.Errorf(codes.NotFound, "badge not found: %v", err)
	}

	badge := mapper.BadgeROToProto(&badgeView)
	return &badgepb.GetBadgeByNameResponse{Badge: badge}, nil
}

// GetAllBadges 获取所有徽章列表
func (h *BadgeHandler) GetAllBadges(ctx context.Context, req *badgepb.GetAllBadgesRequest) (*badgepb.GetAllBadgesResponse, error) {
	var category *string
	if req.Category != nil {
		category = req.Category
	}

	var isSystem *bool
	if req.IsSystem != nil {
		isSystem = req.IsSystem
	}

	badgeViews, err := h.badgeAppSvc.GetAllBadges(ctx, category, isSystem)
	if err != nil {
		logx.S().Errorf("failed to get all badges: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get badges: %v", err)
	}

	badges := mapper.BadgeListROToProto(badgeViews)
	return &badgepb.GetAllBadgesResponse{Badges: badges}, nil
}

// BatchGetBadges 批量获取徽章
func (h *BadgeHandler) BatchGetBadges(ctx context.Context, req *badgepb.BatchGetBadgesRequest) (*badgepb.BatchGetBadgesResponse, error) {
	badgeIDs := make([]uuid.UUID, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		badgeIDs = append(badgeIDs, id)
	}

	badges := make([]*badgepb.Badge, 0, len(badgeIDs))
	for _, badgeID := range badgeIDs {
		badgeView, err := h.badgeAppSvc.GetBadge(ctx, badgeID)
		if err != nil {
			logx.S().Warnf("failed to get badge %s: %v", badgeID, err)
			continue
		}
		badge := mapper.BadgeROToProto(&badgeView)
		badges = append(badges, badge)
	}

	return &badgepb.BatchGetBadgesResponse{Badges: badges}, nil
}
