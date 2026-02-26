package handler

import (
	"context"

	userBadgeApp "nfxidentity/modules/directory/application/user_badges"
	"nfxidentity/modules/directory/interfaces/grpc/mapper"
	"nfxidentity/pkgs/errx"
	userbadgepb "nfxidentity/protos/gen/directory/user_badge"

	"github.com/google/uuid"
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
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userBadgeView, err := h.userBadgeAppSvc.GetUserBadge(ctx, userBadgeID)
	if err != nil {
		return nil, err
	}

	userBadge := mapper.UserBadgeROToProto(&userBadgeView)
	return &userbadgepb.GetUserBadgeByIDResponse{UserBadge: userBadge}, nil
}

// GetUserBadgesByUserID 根据用户ID获取用户徽章列表
func (h *UserBadgeHandler) GetUserBadgesByUserID(
	ctx context.Context,
	req *userbadgepb.GetUserBadgesByUserIDRequest,
) (*userbadgepb.GetUserBadgesByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userBadgeViews, err := h.userBadgeAppSvc.GetUserBadgesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	userBadges := mapper.UserBadgeListROToProto(userBadgeViews)
	return &userbadgepb.GetUserBadgesByUserIDResponse{UserBadges: userBadges}, nil
}

// GetUserBadgesByBadgeID 根据徽章ID获取用户徽章列表
func (h *UserBadgeHandler) GetUserBadgesByBadgeID(
	ctx context.Context,
	req *userbadgepb.GetUserBadgesByBadgeIDRequest,
) (*userbadgepb.GetUserBadgesByBadgeIDResponse, error) {
	badgeID, err := uuid.Parse(req.BadgeId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userBadgeViews, err := h.userBadgeAppSvc.GetUserBadgesByBadgeID(ctx, badgeID)
	if err != nil {
		return nil, err
	}

	userBadges := mapper.UserBadgeListROToProto(userBadgeViews)
	return &userbadgepb.GetUserBadgesByBadgeIDResponse{UserBadges: userBadges}, nil
}
