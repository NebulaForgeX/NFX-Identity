package directory

import (
	"context"
	"fmt"

	"nfxid/connections/directory/dto"
	userpreferencepb "nfxid/protos/gen/directory/user_preference"
)

// UserPreferenceClient UserPreference 客户端
type UserPreferenceClient struct {
	client userpreferencepb.UserPreferenceServiceClient
}

// NewUserPreferenceClient 创建 UserPreference 客户端
func NewUserPreferenceClient(client userpreferencepb.UserPreferenceServiceClient) *UserPreferenceClient {
	return &UserPreferenceClient{client: client}
}

// CreateUserPreference 创建用户偏好（完整参数）
func (c *UserPreferenceClient) CreateUserPreference(ctx context.Context, createDTO *dto.CreateUserPreferenceDTO) (string, error) {
	req, err := createDTO.ToCreateUserPreferenceRequest()
	if err != nil {
		return "", fmt.Errorf("failed to convert DTO to request: %w", err)
	}

	resp, err := c.client.CreateUserPreference(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserPreference.Id, nil
}

// CreateUserPreferenceDefault 创建用户偏好（默认值，用于系统初始化）
func (c *UserPreferenceClient) CreateUserPreferenceDefault(ctx context.Context, userID string) (string, error) {
	createDTO := &dto.CreateUserPreferenceDTO{
		UserID: userID,
		// 其他字段为 nil，使用默认值
	}

	return c.CreateUserPreference(ctx, createDTO)
}

// GetUserPreferenceByID 根据ID获取用户偏好
func (c *UserPreferenceClient) GetUserPreferenceByID(ctx context.Context, id string) (*userpreferencepb.UserPreference, error) {
	req := &userpreferencepb.GetUserPreferenceByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetUserPreferenceByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserPreference, nil
}

// GetUserPreferenceByUserID 根据用户ID获取用户偏好
func (c *UserPreferenceClient) GetUserPreferenceByUserID(ctx context.Context, userID string) (*userpreferencepb.UserPreference, error) {
	req := &userpreferencepb.GetUserPreferenceByUserIDRequest{
		UserId: userID,
	}

	resp, err := c.client.GetUserPreferenceByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserPreference, nil
}