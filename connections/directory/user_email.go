package directory

import (
	"context"
	"fmt"

	"nfxid/connections/directory/dto"
	useremailpb "nfxid/protos/gen/directory/user_email"
)

// UserEmailClient UserEmail 客户端
type UserEmailClient struct {
	client useremailpb.UserEmailServiceClient
}

// NewUserEmailClient 创建 UserEmail 客户端
func NewUserEmailClient(client useremailpb.UserEmailServiceClient) *UserEmailClient {
	return &UserEmailClient{client: client}
}

// CreateUserEmail 创建用户邮箱（完整参数）
func (c *UserEmailClient) CreateUserEmail(ctx context.Context, createDTO *dto.CreateUserEmailDTO) (string, error) {
	req := createDTO.ToCreateUserEmailRequest()

	resp, err := c.client.CreateUserEmail(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserEmail.Id, nil
}

// CreateUserEmailDefault 创建用户邮箱（默认值，用于系统初始化）
func (c *UserEmailClient) CreateUserEmailDefault(ctx context.Context, userID, email string) (string, error) {
	createDTO := &dto.CreateUserEmailDTO{
		UserID:            userID,
		Email:             email,
		IsPrimary:         true,
		IsVerified:        true,
		VerificationToken: nil,
	}

	return c.CreateUserEmail(ctx, createDTO)
}

// GetUserEmailByID 根据ID获取用户邮箱
func (c *UserEmailClient) GetUserEmailByID(ctx context.Context, id string) (*useremailpb.UserEmail, error) {
	req := &useremailpb.GetUserEmailByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetUserEmailByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserEmail, nil
}

// GetUserEmailByEmail 根据邮箱地址获取用户邮箱
func (c *UserEmailClient) GetUserEmailByEmail(ctx context.Context, email string) (*useremailpb.UserEmail, error) {
	req := &useremailpb.GetUserEmailByEmailRequest{
		Email: email,
	}

	resp, err := c.client.GetUserEmailByEmail(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserEmail, nil
}

// GetUserEmailsByUserID 根据用户ID获取用户邮箱列表
func (c *UserEmailClient) GetUserEmailsByUserID(ctx context.Context, userID string) ([]*useremailpb.UserEmail, error) {
	req := &useremailpb.GetUserEmailsByUserIDRequest{
		UserId: userID,
	}

	resp, err := c.client.GetUserEmailsByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserEmails, nil
}