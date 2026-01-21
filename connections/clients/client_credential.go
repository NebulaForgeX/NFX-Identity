package clients

import (
	"context"
	"fmt"

	clientcredentialpb "nfxid/protos/gen/clients/client_credential"
)

// ClientCredentialClient ClientCredential 客户端
type ClientCredentialClient struct {
	client clientcredentialpb.ClientCredentialServiceClient
}

// NewClientCredentialClient 创建 ClientCredential 客户端
func NewClientCredentialClient(client clientcredentialpb.ClientCredentialServiceClient) *ClientCredentialClient {
	return &ClientCredentialClient{client: client}
}

// GetClientCredentialByID 根据ID获取客户端凭证
func (c *ClientCredentialClient) GetClientCredentialByID(ctx context.Context, id string) (*clientcredentialpb.ClientCredential, error) {
	req := &clientcredentialpb.GetClientCredentialByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetClientCredentialByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ClientCredential, nil
}

// GetClientCredentialByClientID 根据客户端ID获取客户端凭证
func (c *ClientCredentialClient) GetClientCredentialByClientID(ctx context.Context, clientID string) (*clientcredentialpb.ClientCredential, error) {
	req := &clientcredentialpb.GetClientCredentialByClientIDRequest{
		ClientId: clientID,
	}

	resp, err := c.client.GetClientCredentialByClientID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ClientCredential, nil
}

// GetClientCredentialsByAppID 根据应用ID获取客户端凭证列表
func (c *ClientCredentialClient) GetClientCredentialsByAppID(ctx context.Context, appID string, status *clientcredentialpb.ClientsCredentialStatus) ([]*clientcredentialpb.ClientCredential, error) {
	req := &clientcredentialpb.GetClientCredentialsByAppIDRequest{
		AppId: appID,
		Status: status,
	}

	resp, err := c.client.GetClientCredentialsByAppID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ClientCredentials, nil
}