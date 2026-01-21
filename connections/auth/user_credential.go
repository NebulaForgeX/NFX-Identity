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
