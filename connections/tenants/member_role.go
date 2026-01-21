package tenants

import (
	"context"
	"fmt"

	memberrolepb "nfxid/protos/gen/tenants/member_role"
)

// MemberRoleClient MemberRole 客户端
type MemberRoleClient struct {
	client memberrolepb.MemberRoleServiceClient
}

// NewMemberRoleClient 创建 MemberRole 客户端
func NewMemberRoleClient(client memberrolepb.MemberRoleServiceClient) *MemberRoleClient {
	return &MemberRoleClient{client: client}
}

// GetMemberRoleByID 根据ID获取成员角色
func (c *MemberRoleClient) GetMemberRoleByID(ctx context.Context, id string) (*memberrolepb.MemberRole, error) {
	req := &memberrolepb.GetMemberRoleByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetMemberRoleByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.MemberRole, nil
}

// GetMemberRolesByMemberID 根据成员ID获取成员角色列表
func (c *MemberRoleClient) GetMemberRolesByMemberID(ctx context.Context, memberID string) ([]*memberrolepb.MemberRole, error) {
	req := &memberrolepb.GetMemberRolesByMemberIDRequest{
		MemberId: memberID,
	}

	resp, err := c.client.GetMemberRolesByMemberID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.MemberRoles, nil
}