package access

import (
	"context"
	"fmt"

	applicationroleassignmentpb "nfxid/protos/gen/access/application_role_assignment"
)

// ApplicationRoleAssignmentClient 应用角色分配 gRPC 客户端
type ApplicationRoleAssignmentClient struct {
	client applicationroleassignmentpb.ApplicationRoleAssignmentServiceClient
}

// NewApplicationRoleAssignmentClient 创建 ApplicationRoleAssignment 客户端
func NewApplicationRoleAssignmentClient(client applicationroleassignmentpb.ApplicationRoleAssignmentServiceClient) *ApplicationRoleAssignmentClient {
	return &ApplicationRoleAssignmentClient{client: client}
}

// GetApplicationRoleAssignmentByID 根据ID获取应用角色分配
func (c *ApplicationRoleAssignmentClient) GetApplicationRoleAssignmentByID(ctx context.Context, id string) (*applicationroleassignmentpb.ApplicationRoleAssignment, error) {
	resp, err := c.client.GetApplicationRoleAssignmentByID(ctx, &applicationroleassignmentpb.GetApplicationRoleAssignmentByIDRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.ApplicationRoleAssignment, nil
}

// GetApplicationRoleAssignmentByUserAndApplication 根据用户与应用获取应用角色分配
func (c *ApplicationRoleAssignmentClient) GetApplicationRoleAssignmentByUserAndApplication(ctx context.Context, userID, applicationID string) (*applicationroleassignmentpb.ApplicationRoleAssignment, error) {
	resp, err := c.client.GetApplicationRoleAssignmentByUserAndApplication(ctx, &applicationroleassignmentpb.GetApplicationRoleAssignmentByUserAndApplicationRequest{
		UserId:        userID,
		ApplicationId: applicationID,
	})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.ApplicationRoleAssignment, nil
}

// ListApplicationRoleAssignmentsByUserID 按用户ID列出应用角色分配
func (c *ApplicationRoleAssignmentClient) ListApplicationRoleAssignmentsByUserID(ctx context.Context, userID string) ([]*applicationroleassignmentpb.ApplicationRoleAssignment, error) {
	resp, err := c.client.ListApplicationRoleAssignmentsByUserID(ctx, &applicationroleassignmentpb.ListApplicationRoleAssignmentsByUserIDRequest{UserId: userID})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.ApplicationRoleAssignments, nil
}

// ListApplicationRoleAssignmentsByApplicationID 按应用ID列出应用角色分配
func (c *ApplicationRoleAssignmentClient) ListApplicationRoleAssignmentsByApplicationID(ctx context.Context, applicationID string) ([]*applicationroleassignmentpb.ApplicationRoleAssignment, error) {
	resp, err := c.client.ListApplicationRoleAssignmentsByApplicationID(ctx, &applicationroleassignmentpb.ListApplicationRoleAssignmentsByApplicationIDRequest{ApplicationId: applicationID})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.ApplicationRoleAssignments, nil
}

// BatchGetApplicationRoleAssignments 批量获取应用角色分配
func (c *ApplicationRoleAssignmentClient) BatchGetApplicationRoleAssignments(ctx context.Context, ids []string) ([]*applicationroleassignmentpb.ApplicationRoleAssignment, error) {
	resp, err := c.client.BatchGetApplicationRoleAssignments(ctx, &applicationroleassignmentpb.BatchGetApplicationRoleAssignmentsRequest{Ids: ids})
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}
	return resp.ApplicationRoleAssignments, nil
}
