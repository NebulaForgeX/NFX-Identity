package clients

import (
	"context"
	"fmt"

	ipallowlistpb "nfxid/protos/gen/clients/ip_allowlist"
)

// IpAllowlistClient IpAllowlist 客户端
type IpAllowlistClient struct {
	client ipallowlistpb.IpAllowlistServiceClient
}

// NewIpAllowlistClient 创建 IpAllowlist 客户端
func NewIpAllowlistClient(client ipallowlistpb.IpAllowlistServiceClient) *IpAllowlistClient {
	return &IpAllowlistClient{client: client}
}

// GetIpAllowlistByID 根据ID获取IP白名单
func (c *IpAllowlistClient) GetIpAllowlistByID(ctx context.Context, id string) (*ipallowlistpb.IpAllowlist, error) {
	req := &ipallowlistpb.GetIpAllowlistByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetIpAllowlistByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.IpAllowlist, nil
}

// GetIpAllowlistByRuleID 根据规则ID获取IP白名单
func (c *IpAllowlistClient) GetIpAllowlistByRuleID(ctx context.Context, ruleID string) (*ipallowlistpb.IpAllowlist, error) {
	req := &ipallowlistpb.GetIpAllowlistByRuleIDRequest{
		RuleId: ruleID,
	}

	resp, err := c.client.GetIpAllowlistByRuleID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.IpAllowlist, nil
}

// GetIpAllowlistsByAppID 根据应用ID获取IP白名单列表
func (c *IpAllowlistClient) GetIpAllowlistsByAppID(ctx context.Context, appID string, status *ipallowlistpb.ClientsAllowlistStatus) ([]*ipallowlistpb.IpAllowlist, error) {
	req := &ipallowlistpb.GetIpAllowlistsByAppIDRequest{
		AppId: appID,
		Status: status,
	}

	resp, err := c.client.GetIpAllowlistsByAppID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.IpAllowlists, nil
}