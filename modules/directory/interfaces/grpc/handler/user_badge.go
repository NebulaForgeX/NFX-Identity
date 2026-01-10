package handler

import (
	"context"

	userBadgeApp "nfxid/modules/directory/application/user_badges"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	userbadgepb "nfxid/protos/gen/directory/user_badge"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserBadgeHandler struct {
	userbadgepb.UnimplementedUserBadgeServiceServer
	userBadgeAppSvc *userBadgeApp.Service
}

func NewUserBadgeHandler(userBadgeAppSvc *userBadgeApp.Service) *UserBadgeHandler {
	return &UserBadgeHandler{
		userBadgeAppSvc: userBadgeAppSvc,
	}
}

// GetUserBadgeByID 根据ID获取用户徽章
func (h *UserBadgeHandler) GetUserBadgeByID(ctx context.Context, req *userbadgepb.GetUserBadgeByIDRequest) (*userbadgepb.GetUserBadgeByIDResponse, error) {
	userBadgeID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_badge_id: %v", err)
	}

	userBadgeView, err := h.userBadgeAppSvc.GetUserBadge(ctx, userBadgeID)
	if err != nil {
		logx.S().Errorf("failed to get user badge by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user badge not found: %v", err)
	}

	userBadge := mapper.UserBadgeROToProto(&userBadgeView)
	return &userbadgepb.GetUserBadgeByIDResponse{UserBadge: userBadge}, nil
}

// GetUserBadgesByUserID 根据用户ID获取用户徽章列表
func (h *UserBadgeHandler) GetUserBadgesByUserID(ctx context.Context, req *userbadgepb.GetUserBadgesByUserIDRequest) (*userbadgepb.GetUserBadgesByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	userBadgeViews, err := h.userBadgeAppSvc.GetUserBadgesByUserID(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get user badges by user_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get user badges: %v", err)
	}

	userBadges := mapper.UserBadgeListROToProto(userBadgeViews)
	return &userbadgepb.GetUserBadgesByUserIDResponse{UserBadges: userBadges}, nil
}

// GetUserBadgesByBadgeID 根据徽章ID获取用户徽章列表
func (h *UserBadgeHandler) GetUserBadgesByBadgeID(ctx context.Context, req *userbadgepb.GetUserBadgesByBadgeIDRequest) (*userbadgepb.GetUserBadgesByBadgeIDResponse, error) {
	badgeID, err := uuid.Parse(req.BadgeId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid badge_id: %v", err)
	}

	userBadgeViews, err := h.userBadgeAppSvc.GetUserBadgesByBadgeID(ctx, badgeID)
	if err != nil {
		logx.S().Errorf("failed to get user badges by badge_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get user badges: %v", err)
	}

	userBadges := mapper.UserBadgeListROToProto(userBadgeViews)
	return &userbadgepb.GetUserBadgesByBadgeIDResponse{UserBadges: userBadges}, nil
}
