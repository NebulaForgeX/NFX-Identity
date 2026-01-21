package clients

import (
	"context"
	"fmt"

	apikeypb "nfxid/protos/gen/clients/api_key"
)

// ApiKeyClient ApiKey 客户端
type ApiKeyClient struct {
	client apikeypb.ApiKeyServiceClient
}

// NewApiKeyClient 创建 ApiKey 客户端
func NewApiKeyClient(client apikeypb.ApiKeyServiceClient) *ApiKeyClient {
	return &ApiKeyClient{client: client}
}

// GetApiKeyByID 根据ID获取API密钥
func (c *ApiKeyClient) GetApiKeyByID(ctx context.Context, id string) (*apikeypb.ApiKey, error) {
	req := &apikeypb.GetApiKeyByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetApiKeyByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ApiKey, nil
}

// GetApiKeyByKeyID 根据密钥ID获取API密钥
func (c *ApiKeyClient) GetApiKeyByKeyID(ctx context.Context, keyID string) (*apikeypb.ApiKey, error) {
	req := &apikeypb.GetApiKeyByKeyIDRequest{
		KeyId: keyID,
	}

	resp, err := c.client.GetApiKeyByKeyID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ApiKey, nil
}

// GetApiKeysByAppID 根据应用ID获取API密钥列表
func (c *ApiKeyClient) GetApiKeysByAppID(ctx context.Context, appID string, status *apikeypb.ClientsApiKeyStatus) ([]*apikeypb.ApiKey, error) {
	req := &apikeypb.GetApiKeysByAppIDRequest{
		AppId: appID,
		Status: status,
	}

	resp, err := c.client.GetApiKeysByAppID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ApiKeys, nil
}