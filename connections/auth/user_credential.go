package auth

import (
	"context"
	"fmt"

	usercredentialpb "nfxid/protos/gen/auth/user_credential"
)

// UserCredentialClient UserCredential 客户端
type UserCredentialClient struct {
	client usercredentialpb.UserCredentialServiceClient
}

// NewUserCredentialClient 创建 UserCredential 客户端
func NewUserCredentialClient(client usercredentialpb.UserCredentialServiceClient) *UserCredentialClient {
	return &UserCredentialClient{client: client}
}

// CreateUserCredential 创建用户凭证
func (c *UserCredentialClient) CreateUserCredential(ctx context.Context, userID, password string, tenantID *string, mustChangePassword bool) error {
	req := &usercredentialpb.CreateUserCredentialRequest{
		UserId:             userID,
		TenantId:           tenantID,
		CredentialType:     usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_PASSWORD,
		Password:           password,
		MustChangePassword: mustChangePassword,
	}

	_, err := c.client.CreateUserCredential(ctx, req)
	if err != nil {
		return fmt.Errorf("gRPC call failed: %w", err)
	}

	return nil
}

// GetUserCredentialByID 根据ID获取用户凭证
func (c *UserCredentialClient) GetUserCredentialByID(ctx context.Context, id string) (*usercredentialpb.UserCredential, error) {
	req := &usercredentialpb.GetUserCredentialByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetUserCredentialByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserCredential, nil
}

// GetUserCredentialByUserID 根据用户ID获取用户凭证
func (c *UserCredentialClient) GetUserCredentialByUserID(ctx context.Context, userID string, tenantID *string) (*usercredentialpb.UserCredential, error) {
	req := &usercredentialpb.GetUserCredentialByUserIDRequest{
		UserId:   userID,
		TenantId: tenantID,
	}

	resp, err := c.client.GetUserCredentialByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserCredential, nil
}

// BatchGetUserCredentials 批量获取用户凭证
func (c *UserCredentialClient) BatchGetUserCredentials(ctx context.Context, ids []string) ([]*usercredentialpb.UserCredential, error) {
	req := &usercredentialpb.BatchGetUserCredentialsRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetUserCredentials(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserCredentials, nil
}
