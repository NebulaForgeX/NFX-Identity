package directory

import (
	"context"
	"fmt"

	"nfxid/connections/directory/dto"
	userprofilepb "nfxid/protos/gen/directory/user_profile"
)

// UserProfileClient UserProfile 客户端
type UserProfileClient struct {
	client userprofilepb.UserProfileServiceClient
}

// NewUserProfileClient 创建 UserProfile 客户端
func NewUserProfileClient(client userprofilepb.UserProfileServiceClient) *UserProfileClient {
	return &UserProfileClient{client: client}
}

// CreateUserProfile 创建用户资料（完整参数）
func (c *UserProfileClient) CreateUserProfile(ctx context.Context, createDTO *dto.CreateUserProfileDTO) (string, error) {
	req, err := createDTO.ToCreateUserProfileRequest()
	if err != nil {
		return "", fmt.Errorf("failed to convert DTO to request: %w", err)
	}

	resp, err := c.client.CreateUserProfile(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserProfile.Id, nil
}

// CreateUserProfileDefault 创建用户资料（默认值，用于系统初始化）
func (c *UserProfileClient) CreateUserProfileDefault(ctx context.Context, userID string) (string, error) {
	createDTO := &dto.CreateUserProfileDTO{
		ID: userID, // id 直接引用 users.id
		// 其他字段为 nil，使用默认值
	}

	return c.CreateUserProfile(ctx, createDTO)
}

// GetUserProfileByID 根据ID获取用户资料
func (c *UserProfileClient) GetUserProfileByID(ctx context.Context, id string) (*userprofilepb.UserProfile, error) {
	req := &userprofilepb.GetUserProfileByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetUserProfileByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserProfile, nil
}

// GetUserProfileByUserID 根据用户ID获取用户资料
func (c *UserProfileClient) GetUserProfileByUserID(ctx context.Context, userID string) (*userprofilepb.UserProfile, error) {
	req := &userprofilepb.GetUserProfileByUserIDRequest{
		UserId: userID,
	}

	resp, err := c.client.GetUserProfileByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserProfile, nil
}

// BatchGetUserProfiles 批量获取用户资料
func (c *UserProfileClient) BatchGetUserProfiles(ctx context.Context, ids []string) ([]*userprofilepb.UserProfile, error) {
	req := &userprofilepb.BatchGetUserProfilesRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetUserProfiles(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserProfiles, nil
}
