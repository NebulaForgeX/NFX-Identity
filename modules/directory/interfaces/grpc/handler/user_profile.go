package handler

import (
	"context"

	userProfileApp "nfxid/modules/directory/application/user_profiles"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	userprofilepb "nfxid/protos/gen/directory/user_profile"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserProfileHandler struct {
	userprofilepb.UnimplementedUserProfileServiceServer
	userProfileAppSvc *userProfileApp.Service
}

func NewUserProfileHandler(userProfileAppSvc *userProfileApp.Service) *UserProfileHandler {
	return &UserProfileHandler{
		userProfileAppSvc: userProfileAppSvc,
	}
}

// GetUserProfileByID 根据ID获取用户资料
func (h *UserProfileHandler) GetUserProfileByID(ctx context.Context, req *userprofilepb.GetUserProfileByIDRequest) (*userprofilepb.GetUserProfileByIDResponse, error) {
	userProfileID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_profile_id: %v", err)
	}

	userProfileView, err := h.userProfileAppSvc.GetUserProfile(ctx, userProfileID)
	if err != nil {
		logx.S().Errorf("failed to get user profile by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user profile not found: %v", err)
	}

	userProfile := mapper.UserProfileROToProto(&userProfileView)
	return &userprofilepb.GetUserProfileByIDResponse{UserProfile: userProfile}, nil
}

// GetUserProfileByUserID 根据用户ID获取用户资料
func (h *UserProfileHandler) GetUserProfileByUserID(ctx context.Context, req *userprofilepb.GetUserProfileByUserIDRequest) (*userprofilepb.GetUserProfileByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	userProfileView, err := h.userProfileAppSvc.GetUserProfileByUserID(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get user profile by user_id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user profile not found: %v", err)
	}

	userProfile := mapper.UserProfileROToProto(&userProfileView)
	return &userprofilepb.GetUserProfileByUserIDResponse{UserProfile: userProfile}, nil
}

// BatchGetUserProfiles 批量获取用户资料
func (h *UserProfileHandler) BatchGetUserProfiles(ctx context.Context, req *userprofilepb.BatchGetUserProfilesRequest) (*userprofilepb.BatchGetUserProfilesResponse, error) {
	userProfileIDs := make([]uuid.UUID, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		userProfileIDs = append(userProfileIDs, id)
	}

	userProfiles := make([]*userprofilepb.UserProfile, 0, len(userProfileIDs))
	for _, userProfileID := range userProfileIDs {
		userProfileView, err := h.userProfileAppSvc.GetUserProfile(ctx, userProfileID)
		if err != nil {
			logx.S().Warnf("failed to get user profile %s: %v", userProfileID, err)
			continue
		}
		userProfile := mapper.UserProfileROToProto(&userProfileView)
		userProfiles = append(userProfiles, userProfile)
	}

	return &userprofilepb.BatchGetUserProfilesResponse{UserProfiles: userProfiles}, nil
}
