package access

import (
	"context"
	"fmt"
	"time"

	grantpb "nfxid/protos/gen/access/grant"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GrantClient Grant 客户端
type GrantClient struct {
	client grantpb.GrantServiceClient
}

// NewGrantClient 创建 Grant 客户端
func NewGrantClient(client grantpb.GrantServiceClient) *GrantClient {
	return &GrantClient{client: client}
}

// CreateGrantOptions 创建授权的选项
type CreateGrantOptions struct {
	TenantID      *string
	AppID         *string
	ResourceType  *string
	ResourceID    *string
	Effect        *string // "allow" or "deny", default is "allow"
	ExpiresAt     *time.Time
	CreatedBy     *string
}

// CreateGrant 创建授权
func (c *GrantClient) CreateGrant(ctx context.Context, subjectType, subjectID, grantType, grantRefID string, opts *CreateGrantOptions) (string, error) {
	// 转换主体类型枚举
	var subjType grantpb.AccessSubjectType
	switch subjectType {
	case "user":
		subjType = grantpb.AccessSubjectType_ACCESS_SUBJECT_TYPE_USER
	case "client":
		subjType = grantpb.AccessSubjectType_ACCESS_SUBJECT_TYPE_CLIENT
	default:
		return "", fmt.Errorf("invalid subject type: %s", subjectType)
	}

	// 转换授权类型枚举
	var gType grantpb.AccessGrantType
	switch grantType {
	case "role":
		gType = grantpb.AccessGrantType_ACCESS_GRANT_TYPE_ROLE
	case "permission":
		gType = grantpb.AccessGrantType_ACCESS_GRANT_TYPE_PERMISSION
	default:
		return "", fmt.Errorf("invalid grant type: %s", grantType)
	}

	// 转换授权效果枚举
	effect := grantpb.AccessGrantEffect_ACCESS_GRANT_EFFECT_ALLOW
	if opts != nil && opts.Effect != nil {
		switch *opts.Effect {
		case "allow":
			effect = grantpb.AccessGrantEffect_ACCESS_GRANT_EFFECT_ALLOW
		case "deny":
			effect = grantpb.AccessGrantEffect_ACCESS_GRANT_EFFECT_DENY
		default:
			return "", fmt.Errorf("invalid effect: %s", *opts.Effect)
		}
	}

	req := &grantpb.CreateGrantRequest{
		SubjectType: subjType,
		SubjectId:   subjectID,
		GrantType:   gType,
		GrantRefId:  grantRefID,
		Effect:      effect,
	}

	if opts != nil {
		req.TenantId = opts.TenantID
		req.AppId = opts.AppID
		req.ResourceType = opts.ResourceType
		req.ResourceId = opts.ResourceID
		req.CreatedBy = opts.CreatedBy
		if opts.ExpiresAt != nil {
			req.ExpiresAt = timestamppb.New(*opts.ExpiresAt)
		}
	}

	resp, err := c.client.CreateGrant(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Grant.Id, nil
}

// GetGrantByID 根据ID获取授权
func (c *GrantClient) GetGrantByID(ctx context.Context, id string) (*grantpb.Grant, error) {
	req := &grantpb.GetGrantByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetGrantByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Grant, nil
}

// GetGrantsBySubject 根据主体获取授权列表
func (c *GrantClient) GetGrantsBySubject(ctx context.Context, subjectType, subjectID string, tenantID *string) ([]*grantpb.Grant, error) {
	var subjType grantpb.AccessSubjectType
	switch subjectType {
	case "user":
		subjType = grantpb.AccessSubjectType_ACCESS_SUBJECT_TYPE_USER
	case "client":
		subjType = grantpb.AccessSubjectType_ACCESS_SUBJECT_TYPE_CLIENT
	default:
		return nil, fmt.Errorf("invalid subject type: %s", subjectType)
	}

	req := &grantpb.GetGrantsBySubjectRequest{
		SubjectType: subjType,
		SubjectId:   subjectID,
		TenantId:    tenantID,
	}

	resp, err := c.client.GetGrantsBySubject(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Grants, nil
}

// BatchGetGrants 批量获取授权
func (c *GrantClient) BatchGetGrants(ctx context.Context, ids []string) ([]*grantpb.Grant, error) {
	req := &grantpb.BatchGetGrantsRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetGrants(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Grants, nil
}