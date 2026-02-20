package access

import (
	"context"
	"fmt"
	applicationrolepb "nfxid/protos/gen/access/application_role"
)

type ApplicationRoleClient struct {
	client applicationrolepb.ApplicationRoleServiceClient
}

func NewApplicationRoleClient(client applicationrolepb.ApplicationRoleServiceClient) *ApplicationRoleClient {
	return &ApplicationRoleClient{client: client}
}

func (c *ApplicationRoleClient) GetApplicationRoleByID(ctx context.Context, id string) (*applicationrolepb.ApplicationRole, error) {
	resp, err := c.client.GetApplicationRoleByID(ctx, &applicationrolepb.GetApplicationRoleByIDRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.ApplicationRole, nil
}

func (c *ApplicationRoleClient) GetApplicationRoleByApplicationIDAndRoleKey(ctx context.Context, applicationID, roleKey string) (*applicationrolepb.ApplicationRole, error) {
	resp, err := c.client.GetApplicationRoleByApplicationIDAndRoleKey(ctx, &applicationrolepb.GetApplicationRoleByApplicationIDAndRoleKeyRequest{ApplicationId: applicationID, RoleKey: roleKey})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.ApplicationRole, nil
}

func (c *ApplicationRoleClient) ListApplicationRolesByApplicationID(ctx context.Context, applicationID string) ([]*applicationrolepb.ApplicationRole, error) {
	resp, err := c.client.ListApplicationRolesByApplicationID(ctx, &applicationrolepb.ListApplicationRolesByApplicationIDRequest{ApplicationId: applicationID})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.ApplicationRoles, nil
}

func (c *ApplicationRoleClient) BatchGetApplicationRoles(ctx context.Context, ids []string) ([]*applicationrolepb.ApplicationRole, error) {
	resp, err := c.client.BatchGetApplicationRoles(ctx, &applicationrolepb.BatchGetApplicationRolesRequest{Ids: ids})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.ApplicationRoles, nil
}
