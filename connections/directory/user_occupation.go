package directory

import (
	"context"
	"fmt"

	useroccupationpb "nfxid/protos/gen/directory/user_occupation"
)

// UserOccupationClient UserOccupation 客户端
type UserOccupationClient struct {
	client useroccupationpb.UserOccupationServiceClient
}

// NewUserOccupationClient 创建 UserOccupation 客户端
func NewUserOccupationClient(client useroccupationpb.UserOccupationServiceClient) *UserOccupationClient {
	return &UserOccupationClient{client: client}
}

// GetUserOccupationByID 根据ID获取用户职业
func (c *UserOccupationClient) GetUserOccupationByID(ctx context.Context, id string) (*useroccupationpb.UserOccupation, error) {
	req := &useroccupationpb.GetUserOccupationByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetUserOccupationByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserOccupation, nil
}

// GetUserOccupationsByUserID 根据用户ID获取用户职业列表
func (c *UserOccupationClient) GetUserOccupationsByUserID(ctx context.Context, userID string, isCurrent *bool) ([]*useroccupationpb.UserOccupation, error) {
	req := &useroccupationpb.GetUserOccupationsByUserIDRequest{
		UserId: userID,
		IsCurrent: isCurrent,
	}

	resp, err := c.client.GetUserOccupationsByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserOccupations, nil
}