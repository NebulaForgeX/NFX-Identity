package handler

import (
	"context"

	badgeApp "nebulaid/modules/auth/application/badge"
	badgeAppViews "nebulaid/modules/auth/application/badge/views"
	profileBadgeApp "nebulaid/modules/auth/application/profile_badge"
	"nebulaid/modules/auth/interfaces/grpc/mapper"
	"nebulaid/pkgs/logx"
	badgepb "nebulaid/protos/gen/auth/badge"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProfileBadgeHandler struct {
	badgepb.UnimplementedProfileBadgeServiceServer
	profileBadgeAppSvc *profileBadgeApp.Service
	badgeAppSvc        *badgeApp.Service
}

func NewProfileBadgeHandler(
	profileBadgeAppSvc *profileBadgeApp.Service,
	badgeAppSvc *badgeApp.Service,
) *ProfileBadgeHandler {
	return &ProfileBadgeHandler{
		profileBadgeAppSvc: profileBadgeAppSvc,
		badgeAppSvc:        badgeAppSvc,
	}
}

// GetProfileBadgeByID 根据ID获取用户徽章关联
func (h *ProfileBadgeHandler) GetProfileBadgeByID(ctx context.Context, req *badgepb.GetProfileBadgeByIDRequest) (*badgepb.GetProfileBadgeByIDResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid profile_badge_id: %v", err)
	}

	profileBadgeView, err := h.profileBadgeAppSvc.GetProfileBadge(ctx, id)
	if err != nil {
		logx.S().Errorf("failed to get profile badge by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "profile badge not found: %v", err)
	}

	// Get Badge information
	var badgeView *badgeAppViews.BadgeView
	if badge, err := h.badgeAppSvc.GetBadge(ctx, profileBadgeView.BadgeID); err == nil {
		badgeView = &badge
	}

	profileBadge := mapper.ProfileBadgeViewToProto(&profileBadgeView, badgeView)
	return &badgepb.GetProfileBadgeByIDResponse{ProfileBadge: profileBadge}, nil
}

// GetProfileBadgesByProfileID 根据ProfileID获取用户徽章列表
func (h *ProfileBadgeHandler) GetProfileBadgesByProfileID(ctx context.Context, req *badgepb.GetProfileBadgesByProfileIDRequest) (*badgepb.GetProfileBadgesByProfileIDResponse, error) {
	profileID, err := uuid.Parse(req.ProfileId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid profile_id: %v", err)
	}

	profileBadgeViews, err := h.profileBadgeAppSvc.GetProfileBadgesByProfileID(ctx, profileID)
	if err != nil {
		logx.S().Errorf("failed to get profile badges by profile_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get profile badges: %v", err)
	}

	profileBadges := make([]*badgepb.ProfileBadge, len(profileBadgeViews))
	for i, profileBadgeView := range profileBadgeViews {
		// Get Badge information
		var badgeView *badgeAppViews.BadgeView
		if badge, err := h.badgeAppSvc.GetBadge(ctx, profileBadgeView.BadgeID); err == nil {
			badgeView = &badge
		}
		profileBadges[i] = mapper.ProfileBadgeViewToProto(&profileBadgeView, badgeView)
	}

	return &badgepb.GetProfileBadgesByProfileIDResponse{ProfileBadges: profileBadges}, nil
}

// GetProfileBadgesByBadgeID 根据BadgeID获取用户徽章列表
func (h *ProfileBadgeHandler) GetProfileBadgesByBadgeID(ctx context.Context, req *badgepb.GetProfileBadgesByBadgeIDRequest) (*badgepb.GetProfileBadgesByBadgeIDResponse, error) {
	badgeID, err := uuid.Parse(req.BadgeId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid badge_id: %v", err)
	}

	profileBadgeViews, err := h.profileBadgeAppSvc.GetProfileBadgesByBadgeID(ctx, badgeID)
	if err != nil {
		logx.S().Errorf("failed to get profile badges by badge_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get profile badges: %v", err)
	}

	// Get Badge information once
	var badgeView *badgeAppViews.BadgeView
	if badge, err := h.badgeAppSvc.GetBadge(ctx, badgeID); err == nil {
		badgeView = &badge
	}

	profileBadges := make([]*badgepb.ProfileBadge, len(profileBadgeViews))
	for i, profileBadgeView := range profileBadgeViews {
		profileBadges[i] = mapper.ProfileBadgeViewToProto(&profileBadgeView, badgeView)
	}

	return &badgepb.GetProfileBadgesByBadgeIDResponse{ProfileBadges: profileBadges}, nil
}
