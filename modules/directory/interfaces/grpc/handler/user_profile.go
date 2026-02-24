package handler

import (
	"context"

	userProfileApp "nfxid/modules/directory/application/user_profiles"
	userProfileAppCommands "nfxid/modules/directory/application/user_profiles/commands"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	"nfxid/pkgs/errx"
	userprofilepb "nfxid/protos/gen/directory/user_profile"

	"github.com/google/uuid"
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

// CreateUserProfile 创建用户资料
func (h *UserProfileHandler) CreateUserProfile(
	ctx context.Context,
	req *userprofilepb.CreateUserProfileRequest,
) (*userprofilepb.CreateUserProfileResponse, error) {
	// 解析用户ID（id 直接引用 users.id）
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	// 转换 protobuf Struct 到 map[string]interface{}
	var socialLinks map[string]interface{}
	if req.SocialLinks != nil {
		socialLinks = req.SocialLinks.AsMap()
	}

	var skills map[string]interface{}
	if req.Skills != nil {
		skills = req.Skills.AsMap()
	}

	// 处理生日
	var birthday *string
	if req.Birthday != nil {
		birthdayStr := req.Birthday.AsTime().Format("2006-01-02")
		birthday = &birthdayStr
	}

	// 转换 Age 从 *int32 到 *int
	var age *int
	if req.Age != nil {
		ageVal := int(*req.Age)
		age = &ageVal
	}

	// 创建命令
	cmd := userProfileAppCommands.CreateUserProfileCmd{
		UserID:      userID,
		Role:        req.Role,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Nickname:    req.Nickname,
		DisplayName: req.DisplayName,
		Bio:         req.Bio,
		Birthday:    birthday,
		Age:         age,
		Gender:      req.Gender,
		Location:    req.Location,
		Website:     req.Website,
		Github:      req.Github,
		SocialLinks: socialLinks,
		Skills:      skills,
	}

	// 调用应用服务创建用户资料
	userProfileID, err := h.userProfileAppSvc.CreateUserProfile(ctx, cmd)
	if err != nil {
		return nil, err
	}

	// 获取创建的用户资料
	userProfileView, err := h.userProfileAppSvc.GetUserProfile(ctx, userProfileID)
	if err != nil {
		return nil, err
	}

	// 转换为 protobuf 响应
	userProfile := mapper.UserProfileROToProto(&userProfileView)
	return &userprofilepb.CreateUserProfileResponse{UserProfile: userProfile}, nil
}

// GetUserProfileByID 根据ID获取用户资料
func (h *UserProfileHandler) GetUserProfileByID(
	ctx context.Context,
	req *userprofilepb.GetUserProfileByIDRequest,
) (*userprofilepb.GetUserProfileByIDResponse, error) {
	userProfileID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userProfileView, err := h.userProfileAppSvc.GetUserProfile(ctx, userProfileID)
	if err != nil {
		return nil, err
	}

	userProfile := mapper.UserProfileROToProto(&userProfileView)
	return &userprofilepb.GetUserProfileByIDResponse{UserProfile: userProfile}, nil
}

// GetUserProfileByUserID 根据用户ID获取用户资料
func (h *UserProfileHandler) GetUserProfileByUserID(
	ctx context.Context,
	req *userprofilepb.GetUserProfileByUserIDRequest,
) (*userprofilepb.GetUserProfileByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userProfileView, err := h.userProfileAppSvc.GetUserProfileByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	userProfile := mapper.UserProfileROToProto(&userProfileView)
	return &userprofilepb.GetUserProfileByUserIDResponse{UserProfile: userProfile}, nil
}

// BatchGetUserProfiles 批量获取用户资料
func (h *UserProfileHandler) BatchGetUserProfiles(
	ctx context.Context,
	req *userprofilepb.BatchGetUserProfilesRequest,
) (*userprofilepb.BatchGetUserProfilesResponse, error) {
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
			continue
		}
		userProfile := mapper.UserProfileROToProto(&userProfileView)
		userProfiles = append(userProfiles, userProfile)
	}

	return &userprofilepb.BatchGetUserProfilesResponse{UserProfiles: userProfiles}, nil
}
