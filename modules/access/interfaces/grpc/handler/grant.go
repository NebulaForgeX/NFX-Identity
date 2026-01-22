package handler

import (
	"context"
	"time"

	grantApp "nfxid/modules/access/application/grants"
	grantAppCommands "nfxid/modules/access/application/grants/commands"
	grantDomain "nfxid/modules/access/domain/grants"
	"nfxid/modules/access/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	grantpb "nfxid/protos/gen/access/grant"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrantHandler struct {
	grantpb.UnimplementedGrantServiceServer
	grantAppSvc *grantApp.Service
}

func NewGrantHandler(grantAppSvc *grantApp.Service) *GrantHandler {
	return &GrantHandler{
		grantAppSvc: grantAppSvc,
	}
}

// CreateGrant 创建授权
func (h *GrantHandler) CreateGrant(ctx context.Context, req *grantpb.CreateGrantRequest) (*grantpb.CreateGrantResponse, error) {
	// 解析主体ID
	subjectID, err := uuid.Parse(req.SubjectId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid subject_id: %v", err)
	}

	// 解析授权引用ID
	grantRefID, err := uuid.Parse(req.GrantRefId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid grant_ref_id: %v", err)
	}

	// 转换 protobuf 类型到 domain 类型
	subjectType := protoToSubjectType(req.SubjectType)
	grantType := protoToGrantType(req.GrantType)
	effect := protoToGrantEffect(req.Effect)

	// 解析可选字段
	var tenantID, appID, resourceID *uuid.UUID
	if req.TenantId != nil && *req.TenantId != "" {
		id, err := uuid.Parse(*req.TenantId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid tenant_id: %v", err)
		}
		tenantID = &id
	}
	if req.AppId != nil && *req.AppId != "" {
		id, err := uuid.Parse(*req.AppId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid app_id: %v", err)
		}
		appID = &id
	}
	if req.ResourceId != nil && *req.ResourceId != "" {
		id, err := uuid.Parse(*req.ResourceId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid resource_id: %v", err)
		}
		resourceID = &id
	}

	// 处理过期时间
	var expiresAt *string
	if req.ExpiresAt != nil {
		expiresAtStr := req.ExpiresAt.AsTime().Format(time.RFC3339)
		expiresAt = &expiresAtStr
	}

	// 解析创建者ID（如果提供）
	var createdBy *uuid.UUID
	if req.CreatedBy != nil && *req.CreatedBy != "" {
		createdByID, err := uuid.Parse(*req.CreatedBy)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid created_by: %v", err)
		}
		createdBy = &createdByID
	}

	// 创建命令
	cmd := grantAppCommands.CreateGrantCmd{
		SubjectType:  subjectType,
		SubjectID:    subjectID,
		GrantType:    grantType,
		GrantRefID:   grantRefID,
		TenantID:     tenantID,
		AppID:        appID,
		ResourceType: req.ResourceType,
		ResourceID:   resourceID,
		Effect:       effect,
		ExpiresAt:    expiresAt,
		CreatedBy:    createdBy,
	}

	// 调用应用服务创建授权
	grantID, err := h.grantAppSvc.CreateGrant(ctx, cmd)
	if err != nil {
		logx.S().Errorf("failed to create grant: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create grant: %v", err)
	}

	// 获取创建的授权
	grantView, err := h.grantAppSvc.GetGrant(ctx, grantID)
	if err != nil {
		logx.S().Errorf("failed to get created grant: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get created grant: %v", err)
	}

	// 转换为 protobuf 响应
	grant := mapper.GrantROToProto(&grantView)
	return &grantpb.CreateGrantResponse{Grant: grant}, nil
}

// GetGrantByID 根据ID获取授权
func (h *GrantHandler) GetGrantByID(ctx context.Context, req *grantpb.GetGrantByIDRequest) (*grantpb.GetGrantByIDResponse, error) {
	grantID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid grant_id: %v", err)
	}

	grantView, err := h.grantAppSvc.GetGrant(ctx, grantID)
	if err != nil {
		logx.S().Errorf("failed to get grant by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "grant not found: %v", err)
	}

	grant := mapper.GrantROToProto(&grantView)
	return &grantpb.GetGrantByIDResponse{Grant: grant}, nil
}

// GetGrantsBySubject 根据主体获取授权列表
func (h *GrantHandler) GetGrantsBySubject(ctx context.Context, req *grantpb.GetGrantsBySubjectRequest) (*grantpb.GetGrantsBySubjectResponse, error) {
	subjectID, err := uuid.Parse(req.SubjectId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid subject_id: %v", err)
	}

	subjectType := protoToSubjectType(req.SubjectType)
	grantViews, err := h.grantAppSvc.GetGrantsBySubject(ctx, subjectType, subjectID)
	if err != nil {
		logx.S().Errorf("failed to get grants by subject: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get grants: %v", err)
	}

	protoGrants := mapper.GrantListROToProto(grantViews)
	return &grantpb.GetGrantsBySubjectResponse{Grants: protoGrants}, nil
}

// BatchGetGrants 批量获取授权
func (h *GrantHandler) BatchGetGrants(ctx context.Context, req *grantpb.BatchGetGrantsRequest) (*grantpb.BatchGetGrantsResponse, error) {
	grantIDs := make([]uuid.UUID, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		grantIDs = append(grantIDs, id)
	}

	grants := make([]*grantpb.Grant, 0, len(grantIDs))
	for _, grantID := range grantIDs {
		grantView, err := h.grantAppSvc.GetGrant(ctx, grantID)
		if err != nil {
			logx.S().Warnf("failed to get grant %s: %v", grantID, err)
			continue
		}
		grant := mapper.GrantROToProto(&grantView)
		grants = append(grants, grant)
	}

	return &grantpb.BatchGetGrantsResponse{Grants: grants}, nil
}

func protoToSubjectType(pt grantpb.AccessSubjectType) grantDomain.SubjectType {
	switch pt {
	case grantpb.AccessSubjectType_ACCESS_SUBJECT_TYPE_USER:
		return grantDomain.SubjectTypeUser
	case grantpb.AccessSubjectType_ACCESS_SUBJECT_TYPE_CLIENT:
		return grantDomain.SubjectTypeClient
	default:
		return grantDomain.SubjectTypeUser
	}
}

func protoToGrantType(gt grantpb.AccessGrantType) grantDomain.GrantType {
	switch gt {
	case grantpb.AccessGrantType_ACCESS_GRANT_TYPE_ROLE:
		return grantDomain.GrantTypeRole
	case grantpb.AccessGrantType_ACCESS_GRANT_TYPE_PERMISSION:
		return grantDomain.GrantTypePermission
	default:
		return grantDomain.GrantTypeRole
	}
}

func protoToGrantEffect(ge grantpb.AccessGrantEffect) grantDomain.GrantEffect {
	switch ge {
	case grantpb.AccessGrantEffect_ACCESS_GRANT_EFFECT_ALLOW:
		return grantDomain.GrantEffectAllow
	case grantpb.AccessGrantEffect_ACCESS_GRANT_EFFECT_DENY:
		return grantDomain.GrantEffectDeny
	default:
		return grantDomain.GrantEffectAllow
	}
}
