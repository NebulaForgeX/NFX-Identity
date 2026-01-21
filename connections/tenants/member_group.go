package tenants

import (
	"context"
	"fmt"

	membergrouppb "nfxid/protos/gen/tenants/member_group"
)

// MemberGroupClient MemberGroup 客户端
type MemberGroupClient struct {
	client membergrouppb.MemberGroupServiceClient
}

// NewMemberGroupClient 创建 MemberGroup 客户端
func NewMemberGroupClient(client membergrouppb.MemberGroupServiceClient) *MemberGroupClient {
	return &MemberGroupClient{client: client}
}

// GetMemberGroupByID 根据ID获取成员组
func (c *MemberGroupClient) GetMemberGroupByID(ctx context.Context, id string) (*membergrouppb.MemberGroup, error) {
	req := &membergrouppb.GetMemberGroupByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetMemberGroupByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.MemberGroup, nil
}

// GetMemberGroupsByMemberID 根据成员ID获取成员组列表
func (c *MemberGroupClient) GetMemberGroupsByMemberID(ctx context.Context, memberID string) ([]*membergrouppb.MemberGroup, error) {
	req := &membergrouppb.GetMemberGroupsByMemberIDRequest{
		MemberId: memberID,
	}

	resp, err := c.client.GetMemberGroupsByMemberID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.MemberGroups, nil
}

// GetMemberGroupsByGroupID 根据组ID获取成员组列表
func (c *MemberGroupClient) GetMemberGroupsByGroupID(ctx context.Context, groupID string) ([]*membergrouppb.MemberGroup, error) {
	req := &membergrouppb.GetMemberGroupsByGroupIDRequest{
		GroupId: groupID,
	}

	resp, err := c.client.GetMemberGroupsByGroupID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.MemberGroups, nil
}