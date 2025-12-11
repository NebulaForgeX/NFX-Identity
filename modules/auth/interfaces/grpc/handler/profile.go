package handler

import (
	"context"

	badgeApp "nfxid/modules/auth/application/badge"
	badgeAppViews "nfxid/modules/auth/application/badge/views"
	profileApp "nfxid/modules/auth/application/profile"
	profileBadgeApp "nfxid/modules/auth/application/profile_badge"
	"nfxid/modules/auth/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	badgepb "nfxid/protos/gen/auth/badge"
	profilepb "nfxid/protos/gen/auth/profile"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProfileHandler struct {
	profilepb.UnimplementedProfileServiceServer
	profileAppSvc      *profileApp.Service
	profileBadgeAppSvc *profileBadgeApp.Service
	badgeAppSvc        *badgeApp.Service
}

func NewProfileHandler(
	profileAppSvc *profileApp.Service,
	profileBadgeAppSvc *profileBadgeApp.Service,
	badgeAppSvc *badgeApp.Service,
) *ProfileHandler {
	return &ProfileHandler{
		profileAppSvc:      profileAppSvc,
		profileBadgeAppSvc: profileBadgeAppSvc,
		badgeAppSvc:        badgeAppSvc,
	}
}

// GetProfileByID 根据ID获取资料
func (h *ProfileHandler) GetProfileByID(ctx context.Context, req *profilepb.GetProfileByIDRequest) (*profilepb.GetProfileByIDResponse, error) {
	profileID, err := uuid.Parse(req.ProfileId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid profile_id: %v", err)
	}

	profileView, err := h.profileAppSvc.GetProfile(ctx, profileID)
	if err != nil {
		logx.S().Errorf("failed to get profile by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "profile not found: %v", err)
	}

	includeBadges := req.IncludeBadges != nil && *req.IncludeBadges
	profile := mapper.ProfileViewToProto(&profileView, includeBadges)

	// 如果需要包含徽章信息
	if includeBadges {
		profileBadgeViews, err := h.profileBadgeAppSvc.GetProfileBadgesByProfileID(ctx, profileView.ID)
		if err == nil && len(profileBadgeViews) > 0 {
			badgeProtos := make([]*badgepb.ProfileBadge, 0, len(profileBadgeViews))
			for _, profileBadgeView := range profileBadgeViews {
				// 获取 badge 信息
				var badgeView *badgeAppViews.BadgeView
				if badge, err := h.badgeAppSvc.GetBadge(ctx, profileBadgeView.BadgeID); err == nil {
					badgeView = &badge
				}
				badgeProto := mapper.ProfileBadgeViewToProto(&profileBadgeView, badgeView)
				badgeProtos = append(badgeProtos, badgeProto)
			}
			profile.Badges = badgeProtos
		}
	}

	return &profilepb.GetProfileByIDResponse{Profile: profile}, nil
}

// GetProfileByUserID 根据用户ID获取资料
func (h *ProfileHandler) GetProfileByUserID(ctx context.Context, req *profilepb.GetProfileByUserIDRequest) (*profilepb.GetProfileByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	profileView, err := h.profileAppSvc.GetProfileByUserID(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get profile by user_id: %v", err)
		return nil, status.Errorf(codes.NotFound, "profile not found: %v", err)
	}

	includeBadges := req.IncludeBadges != nil && *req.IncludeBadges
	profile := mapper.ProfileViewToProto(&profileView, includeBadges)

	// 如果需要包含徽章信息
	if includeBadges {
		profileBadgeViews, err := h.profileBadgeAppSvc.GetProfileBadgesByProfileID(ctx, profileView.ID)
		if err == nil && len(profileBadgeViews) > 0 {
			badgeProtos := make([]*badgepb.ProfileBadge, 0, len(profileBadgeViews))
			for _, profileBadgeView := range profileBadgeViews {
				// 获取 badge 信息
				var badgeView *badgeAppViews.BadgeView
				if badge, err := h.badgeAppSvc.GetBadge(ctx, profileBadgeView.BadgeID); err == nil {
					badgeView = &badge
				}
				badgeProto := mapper.ProfileBadgeViewToProto(&profileBadgeView, badgeView)
				badgeProtos = append(badgeProtos, badgeProto)
			}
			profile.Badges = badgeProtos
		}
	}

	return &profilepb.GetProfileByUserIDResponse{Profile: profile}, nil
}

// BatchGetProfiles 批量获取资料
func (h *ProfileHandler) BatchGetProfiles(ctx context.Context, req *profilepb.BatchGetProfilesRequest) (*profilepb.BatchGetProfilesResponse, error) {
	profiles := make([]*profilepb.Profile, 0)
	errorById := make(map[string]string)

	includeBadges := req.IncludeBadges != nil && *req.IncludeBadges

	// 处理 profile_ids
	for _, idStr := range req.ProfileIds {
		profileID, err := uuid.Parse(idStr)
		if err != nil {
			errorById[idStr] = "invalid profile_id format"
			continue
		}

		profileView, err := h.profileAppSvc.GetProfile(ctx, profileID)
		if err != nil {
			errorById[idStr] = err.Error()
			continue
		}
		profile := mapper.ProfileViewToProto(&profileView, includeBadges)

		// 如果需要包含徽章信息
		if includeBadges {
			profileBadgeViews, err := h.profileBadgeAppSvc.GetProfileBadgesByProfileID(ctx, profileView.ID)
			if err == nil && len(profileBadgeViews) > 0 {
				badgeProtos := make([]*badgepb.ProfileBadge, 0, len(profileBadgeViews))
				for _, profileBadgeView := range profileBadgeViews {
					var badgeView *badgeAppViews.BadgeView
					if badge, err := h.badgeAppSvc.GetBadge(ctx, profileBadgeView.BadgeID); err == nil {
						badgeView = &badge
					}
					badgeProto := mapper.ProfileBadgeViewToProto(&profileBadgeView, badgeView)
					badgeProtos = append(badgeProtos, badgeProto)
				}
				profile.Badges = badgeProtos
			}
		}

		profiles = append(profiles, profile)
	}

	// 处理 user_ids
	for _, idStr := range req.UserIds {
		userID, err := uuid.Parse(idStr)
		if err != nil {
			errorById[idStr] = "invalid user_id format"
			continue
		}

		profileView, err := h.profileAppSvc.GetProfileByUserID(ctx, userID)
		if err != nil {
			errorById[idStr] = err.Error()
			continue
		}
		profile := mapper.ProfileViewToProto(&profileView, includeBadges)

		// 如果需要包含徽章信息
		if includeBadges {
			profileBadgeViews, err := h.profileBadgeAppSvc.GetProfileBadgesByProfileID(ctx, profileView.ID)
			if err == nil && len(profileBadgeViews) > 0 {
				badgeProtos := make([]*badgepb.ProfileBadge, 0, len(profileBadgeViews))
				for _, profileBadgeView := range profileBadgeViews {
					var badgeView *badgeAppViews.BadgeView
					if badge, err := h.badgeAppSvc.GetBadge(ctx, profileBadgeView.BadgeID); err == nil {
						badgeView = &badge
					}
					badgeProto := mapper.ProfileBadgeViewToProto(&profileBadgeView, badgeView)
					badgeProtos = append(badgeProtos, badgeProto)
				}
				profile.Badges = badgeProtos
			}
		}

		profiles = append(profiles, profile)
	}

	return &profilepb.BatchGetProfilesResponse{
		Profiles:  profiles,
		ErrorById: errorById,
	}, nil
}
