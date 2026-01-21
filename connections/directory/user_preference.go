package directory

import (
	"context"
	"fmt"

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