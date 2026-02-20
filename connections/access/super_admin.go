package access

import (
	"context"
	"fmt"
	superadminpb "nfxid/protos/gen/access/super_admin"
)

type SuperAdminClient struct {
	client superadminpb.SuperAdminServiceClient
}

func NewSuperAdminClient(client superadminpb.SuperAdminServiceClient) *SuperAdminClient {
	return &SuperAdminClient{client: client}
}

func (c *SuperAdminClient) GetSuperAdminByUserID(ctx context.Context, userID string) (*superadminpb.SuperAdmin, error) {
	resp, err := c.client.GetSuperAdminByUserID(ctx, &superadminpb.GetSuperAdminByUserIDRequest{UserId: userID})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.SuperAdmin, nil
}

func (c *SuperAdminClient) ListSuperAdmins(ctx context.Context, limit, offset *int32) ([]*superadminpb.SuperAdmin, error) {
	req := &superadminpb.ListSuperAdminsRequest{}
	if limit != nil {
		req.Limit = limit
	}
	if offset != nil {
		req.Offset = offset
	}
	resp, err := c.client.ListSuperAdmins(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.SuperAdmins, nil
}
