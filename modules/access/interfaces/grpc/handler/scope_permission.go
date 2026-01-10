package handler

import (
	"context"

	scopePermissionApp "nfxid/modules/access/application/scope_permissions"
	"nfxid/modules/access/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	scopepermissionpb "nfxid/protos/gen/access/scope_permission"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ScopePermissionHandler struct {
	scopepermissionpb.UnimplementedScopePermissionServiceServer
	scopePermissionAppSvc *scopePermissionApp.Service
}

func NewScopePermissionHandler(scopePermissionAppSvc *scopePermissionApp.Service) *ScopePermissionHandler {
	return &ScopePermissionHandler{
		scopePermissionAppSvc: scopePermissionAppSvc,
	}
}

// GetScopePermissionByID 根据ID获取范围权限关联
func (h *ScopePermissionHandler) GetScopePermissionByID(ctx context.Context, req *scopepermissionpb.GetScopePermissionByIDRequest) (*scopepermissionpb.GetScopePermissionByIDResponse, error) {
	scopePermissionID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid scope_permission_id: %v", err)
	}

	scopePermissionView, err := h.scopePermissionAppSvc.GetScopePermission(ctx, scopePermissionID)
	if err != nil {
		logx.S().Errorf("failed to get scope permission by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "scope permission not found: %v", err)
	}

	scopePermission := mapper.ScopePermissionROToProto(&scopePermissionView)
	return &scopepermissionpb.GetScopePermissionByIDResponse{ScopePermission: scopePermission}, nil
}

// GetPermissionsByScope 根据范围获取权限列表
func (h *ScopePermissionHandler) GetPermissionsByScope(ctx context.Context, req *scopepermissionpb.GetPermissionsByScopeRequest) (*scopepermissionpb.GetPermissionsByScopeResponse, error) {
	scopePermissionViews, err := h.scopePermissionAppSvc.GetScopePermissionsByScope(ctx, req.Scope)
	if err != nil {
		logx.S().Errorf("failed to get scope permissions by scope: %v", err)
		return nil, status.Errorf(codes.NotFound, "scope permissions not found: %v", err)
	}

	scopePermissions := mapper.ScopePermissionListROToProto(scopePermissionViews)
	return &scopepermissionpb.GetPermissionsByScopeResponse{ScopePermissions: scopePermissions}, nil
}

// GetScopesByPermission 根据权限获取范围列表
func (h *ScopePermissionHandler) GetScopesByPermission(ctx context.Context, req *scopepermissionpb.GetScopesByPermissionRequest) (*scopepermissionpb.GetScopesByPermissionResponse, error) {
	permissionID, err := uuid.Parse(req.PermissionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid permission_id: %v", err)
	}

	scopePermissionViews, err := h.scopePermissionAppSvc.GetScopePermissionsByPermissionID(ctx, permissionID)
	if err != nil {
		logx.S().Errorf("failed to get scope permissions by permission_id: %v", err)
		return nil, status.Errorf(codes.NotFound, "scope permissions not found: %v", err)
	}

	scopePermissions := mapper.ScopePermissionListROToProto(scopePermissionViews)
	return &scopepermissionpb.GetScopesByPermissionResponse{ScopePermissions: scopePermissions}, nil
}
