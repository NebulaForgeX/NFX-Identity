package directory

import (
	"context"
	"fmt"

	usereducationpb "nfxid/protos/gen/directory/user_education"
)

// UserEducationClient UserEducation 客户端
type UserEducationClient struct {
	client usereducationpb.UserEducationServiceClient
}

// NewUserEducationClient 创建 UserEducation 客户端
func NewUserEducationClient(client usereducationpb.UserEducationServiceClient) *UserEducationClient {
	return &UserEducationClient{client: client}
}

// GetUserEducationByID 根据ID获取用户教育
func (c *UserEducationClient) GetUserEducationByID(ctx context.Context, id string) (*usereducationpb.UserEducation, error) {
	req := &usereducationpb.GetUserEducationByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetUserEducationByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserEducation, nil
}

// GetUserEducationsByUserID 根据用户ID获取用户教育列表
func (c *UserEducationClient) GetUserEducationsByUserID(ctx context.Context, userID string) ([]*usereducationpb.UserEducation, error) {
	req := &usereducationpb.GetUserEducationsByUserIDRequest{
		UserId: userID,
	}

	resp, err := c.client.GetUserEducationsByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserEducations, nil
}