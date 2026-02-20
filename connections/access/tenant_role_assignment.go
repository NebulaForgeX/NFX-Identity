package access

import (
	"context"
	"fmt"
	tenantroleassignmentpb "nfxid/protos/gen/access/tenant_role_assignment"
)

type TenantRoleAssignmentClient struct {
	client tenantroleassignmentpb.TenantRoleAssignmentServiceClient
}

func NewTenantRoleAssignmentClient(client tenantroleassignmentpb.TenantRoleAssignmentServiceClient) *TenantRoleAssignmentClient {
	return &TenantRoleAssignmentClient{client: client}
}

func (c *TenantRoleAssignmentClient) GetTenantRoleAssignmentByID(ctx context.Context, id string) (*tenantroleassignmentpb.TenantRoleAssignment, error) {
	resp, err := c.client.GetTenantRoleAssignmentByID(ctx, &tenantroleassignmentpb.GetTenantRoleAssignmentByIDRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.TenantRoleAssignment, nil
}

func (c *TenantRoleAssignmentClient) GetTenantRoleAssignmentByUserAndTenant(ctx context.Context, userID, tenantID string) (*tenantroleassignmentpb.TenantRoleAssignment, error) {
	resp, err := c.client.GetTenantRoleAssignmentByUserAndTenant(ctx, &tenantroleassignmentpb.GetTenantRoleAssignmentByUserAndTenantRequest{UserId: userID, TenantId: tenantID})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.TenantRoleAssignment, nil
}

func (c *TenantRoleAssignmentClient) ListTenantRoleAssignmentsByUserID(ctx context.Context, userID string) ([]*tenantroleassignmentpb.TenantRoleAssignment, error) {
	resp, err := c.client.ListTenantRoleAssignmentsByUserID(ctx, &tenantroleassignmentpb.ListTenantRoleAssignmentsByUserIDRequest{UserId: userID})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.TenantRoleAssignments, nil
}

func (c *TenantRoleAssignmentClient) ListTenantRoleAssignmentsByTenantID(ctx context.Context, tenantID string) ([]*tenantroleassignmentpb.TenantRoleAssignment, error) {
	resp, err := c.client.ListTenantRoleAssignmentsByTenantID(ctx, &tenantroleassignmentpb.ListTenantRoleAssignmentsByTenantIDRequest{TenantId: tenantID})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.TenantRoleAssignments, nil
}

func (c *TenantRoleAssignmentClient) BatchGetTenantRoleAssignments(ctx context.Context, ids []string) ([]*tenantroleassignmentpb.TenantRoleAssignment, error) {
	resp, err := c.client.BatchGetTenantRoleAssignments(ctx, &tenantroleassignmentpb.BatchGetTenantRoleAssignmentsRequest{Ids: ids})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.TenantRoleAssignments, nil
}
