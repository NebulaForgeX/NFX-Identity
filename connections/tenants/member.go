package tenants

import (
	"context"
	"fmt"

	memberpb "nfxid/protos/gen/tenants/member"
)

// MemberClient Member 客户端
type MemberClient struct {
	client memberpb.MemberServiceClient
}

// NewMemberClient 创建 Member 客户端
func NewMemberClient(client memberpb.MemberServiceClient) *MemberClient {
	return &MemberClient{client: client}
}

// GetMemberByID 根据ID获取成员
func (c *MemberClient) GetMemberByID(ctx context.Context, id string) (*memberpb.Member, error) {
	req := &memberpb.GetMemberByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetMemberByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Member, nil
}

// GetMemberByUserID 根据用户ID和租户ID获取成员
func (c *MemberClient) GetMemberByUserID(ctx context.Context, userID, tenantID string) (*memberpb.Member, error) {
	req := &memberpb.GetMemberByUserIDRequest{
		UserId:   userID,
		TenantId: tenantID,
	}

	resp, err := c.client.GetMemberByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Member, nil
}

// GetMembersByTenantID 根据租户ID获取成员列表
func (c *MemberClient) GetMembersByTenantID(ctx context.Context, tenantID string, status *memberpb.TenantsMemberStatus) ([]*memberpb.Member, error) {
	req := &memberpb.GetMembersByTenantIDRequest{
		TenantId: tenantID,
		Status:   status,
	}

	resp, err := c.client.GetMembersByTenantID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Members, nil
}

// BatchGetMembers 批量获取成员
func (c *MemberClient) BatchGetMembers(ctx context.Context, ids []string) ([]*memberpb.Member, error) {
	req := &memberpb.BatchGetMembersRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetMembers(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Members, nil
}