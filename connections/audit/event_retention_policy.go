package audit

import (
	"context"
	"fmt"

	eventretentionpolicypb "nfxid/protos/gen/audit/event_retention_policy"
)

// EventRetentionPolicyClient EventRetentionPolicy 客户端
type EventRetentionPolicyClient struct {
	client eventretentionpolicypb.EventRetentionPolicyServiceClient
}

// NewEventRetentionPolicyClient 创建 EventRetentionPolicy 客户端
func NewEventRetentionPolicyClient(client eventretentionpolicypb.EventRetentionPolicyServiceClient) *EventRetentionPolicyClient {
	return &EventRetentionPolicyClient{client: client}
}

// GetEventRetentionPolicyByID 根据ID获取策略
func (c *EventRetentionPolicyClient) GetEventRetentionPolicyByID(ctx context.Context, id string) (*eventretentionpolicypb.EventRetentionPolicy, error) {
	req := &eventretentionpolicypb.GetEventRetentionPolicyByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetEventRetentionPolicyByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.EventRetentionPolicy, nil
}

// GetEventRetentionPolicyByName 根据策略名称获取策略
func (c *EventRetentionPolicyClient) GetEventRetentionPolicyByName(ctx context.Context, policyName string) (*eventretentionpolicypb.EventRetentionPolicy, error) {
	req := &eventretentionpolicypb.GetEventRetentionPolicyByNameRequest{
		PolicyName: policyName,
	}

	resp, err := c.client.GetEventRetentionPolicyByName(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.EventRetentionPolicy, nil
}

// GetAllEventRetentionPolicies 获取所有策略列表
func (c *EventRetentionPolicyClient) GetAllEventRetentionPolicies(ctx context.Context, tenantID *string) ([]*eventretentionpolicypb.EventRetentionPolicy, error) {
	req := &eventretentionpolicypb.GetAllEventRetentionPoliciesRequest{
		TenantId: tenantID,
	}

	resp, err := c.client.GetAllEventRetentionPolicies(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.EventRetentionPolicies, nil
}