package tenants

import (
	"context"
	"fmt"

	invitationpb "nfxid/protos/gen/tenants/invitation"
)

// InvitationClient Invitation 客户端
type InvitationClient struct {
	client invitationpb.InvitationServiceClient
}

// NewInvitationClient 创建 Invitation 客户端
func NewInvitationClient(client invitationpb.InvitationServiceClient) *InvitationClient {
	return &InvitationClient{client: client}
}

// GetInvitationByID 根据ID获取邀请
func (c *InvitationClient) GetInvitationByID(ctx context.Context, id string) (*invitationpb.Invitation, error) {
	req := &invitationpb.GetInvitationByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetInvitationByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Invitation, nil
}

// GetInvitationByInvitationID 根据邀请ID获取邀请
func (c *InvitationClient) GetInvitationByInvitationID(ctx context.Context, invitationID string) (*invitationpb.Invitation, error) {
	req := &invitationpb.GetInvitationByInvitationIDRequest{
		InvitationId: invitationID,
	}

	resp, err := c.client.GetInvitationByInvitationID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Invitation, nil
}

// GetInvitationsByTenantID 根据租户ID获取邀请列表
func (c *InvitationClient) GetInvitationsByTenantID(ctx context.Context, tenantID string, status *invitationpb.TenantsInvitationStatus) ([]*invitationpb.Invitation, error) {
	req := &invitationpb.GetInvitationsByTenantIDRequest{
		TenantId: tenantID,
		Status:   status,
	}

	resp, err := c.client.GetInvitationsByTenantID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Invitations, nil
}