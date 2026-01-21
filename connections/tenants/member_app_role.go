package tenants

import (
	"context"
	"fmt"

	memberapprolepb "nfxid/protos/gen/tenants/member_app_role"
)

// MemberAppRoleClient MemberAppRole 客户端
type MemberAppRoleClient struct {
	client memberapprolepb.MemberAppRoleServiceClient
}

// NewMemberAppRoleClient 创建 MemberAppRole 客户端
func NewMemberAppRoleClient(client memberapprolepb.MemberAppRoleServiceClient) *MemberAppRoleClient {
	return &MemberAppRoleClient{client: client}
}

// GetMemberAppRoleByID 根据ID获取成员应用角色
func (c *MemberAppRoleClient) GetMemberAppRoleByID(ctx context.Context, id string) (*memberapprolepb.MemberAppRole, error) {
	req := &memberapprolepb.GetMemberAppRoleByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetMemberAppRoleByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.MemberAppRole, nil
}

// GetMemberAppRolesByMemberID 根据成员ID获取成员应用角色列表
func (c *MemberAppRoleClient) GetMemberAppRolesByMemberID(ctx context.Context, memberID string, appID *string) ([]*memberapprolepb.MemberAppRole, error) {
	req := &memberapprolepb.GetMemberAppRolesByMemberIDRequest{
		MemberId: memberID,
		AppId:    appID,
	}

	resp, err := c.client.GetMemberAppRolesByMemberID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.MemberAppRoles, nil
}