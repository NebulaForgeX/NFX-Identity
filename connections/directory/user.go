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
