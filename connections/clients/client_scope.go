package clients

import (
	"context"
	"fmt"

	clientscopepb "nfxid/protos/gen/clients/client_scope"
)

// ClientScopeClient ClientScope 客户端
type ClientScopeClient struct {
	client clientscopepb.ClientScopeServiceClient
}

// NewClientScopeClient 创建 ClientScope 客户端
func NewClientScopeClient(client clientscopepb.ClientScopeServiceClient) *ClientScopeClient {
	return &ClientScopeClient{client: client}
}

// GetClientScopeByID 根据ID获取客户端作用域
func (c *ClientScopeClient) GetClientScopeByID(ctx context.Context, id string) (*clientscopepb.ClientScope, error) {
	req := &clientscopepb.GetClientScopeByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetClientScopeByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ClientScope, nil
}

// GetClientScopesByAppID 根据应用ID获取作用域列表
func (c *ClientScopeClient) GetClientScopesByAppID(ctx context.Context, appID string) ([]*clientscopepb.ClientScope, error) {
	req := &clientscopepb.GetClientScopesByAppIDRequest{
		AppId: appID,
	}

	resp, err := c.client.GetClientScopesByAppID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ClientScopes, nil
}