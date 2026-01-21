package directory

import (
	"context"
	"fmt"

	userpb "nfxid/protos/gen/directory/user"
)

// UserClient User 客户端
type UserClient struct {
	client userpb.UserServiceClient
}

// NewUserClient 创建 User 客户端
func NewUserClient(client userpb.UserServiceClient) *UserClient {
	return &UserClient{client: client}
}

// CreateUser 创建用户
func (c *UserClient) CreateUser(ctx context.Context, username string, status string, isVerified bool) (string, error) {
	// 转换状态枚举
	var userStatus userpb.DirectoryUserStatus
	switch status {
	case "active":
		userStatus = userpb.DirectoryUserStatus_DIRECTORY_USER_STATUS_ACTIVE
	case "pending":
		userStatus = userpb.DirectoryUserStatus_DIRECTORY_USER_STATUS_PENDING
	case "deactive":
		userStatus = userpb.DirectoryUserStatus_DIRECTORY_USER_STATUS_DEACTIVE
	default:
		userStatus = userpb.DirectoryUserStatus_DIRECTORY_USER_STATUS_ACTIVE
	}

	req := &userpb.CreateUserRequest{
		Username:   username,
		Status:     userStatus,
		IsVerified: isVerified,
	}

	resp, err := c.client.CreateUser(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.User.Id, nil
}

// GetUserByID 根据ID获取用户
func (c *UserClient) GetUserByID(ctx context.Context, id string) (*userpb.User, error) {
	req := &userpb.GetUserByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetUserByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.User, nil
}

// GetUserByUsername 根据用户名获取用户
func (c *UserClient) GetUserByUsername(ctx context.Context, username string) (*userpb.User, error) {
	req := &userpb.GetUserByUsernameRequest{
		Username: username,
	}

	resp, err := c.client.GetUserByUsername(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.User, nil
}

// BatchGetUsers 批量获取用户
func (c *UserClient) BatchGetUsers(ctx context.Context, ids []string) ([]*userpb.User, error) {
	req := &userpb.BatchGetUsersRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetUsers(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Users, nil
}
