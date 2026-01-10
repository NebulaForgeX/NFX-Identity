package handler

import (
	"context"

	rolePermissionApp "nfxid/modules/access/application/role_permissions"
	"nfxid/modules/access/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	rolepermissionpb "nfxid/protos/gen/access/role_permission"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RolePermissionHandler struct {
	rolepermissionpb.UnimplementedRolePermissionServiceServer
	rolePermissionAppSvc *rolePermissionApp.Service
}

func NewRolePermissionHandler(rolePermissionAppSvc *rolePermissionApp.Service) *RolePermissionHandler {
	return &RolePermissionHandler{
		rolePermissionAppSvc: rolePermissionAppSvc,
	}
}

// GetRolePermissionByID 根据ID获取角色权限关联
func (h *RolePermissionHandler) GetRolePermissionByID(ctx context.Context, req *rolepermissionpb.GetRolePermissionByIDRequest) (*rolepermissionpb.GetRolePermissionByIDResponse, error) {
	rolePermissionID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid role_permission_id: %v", err)
	}

	rolePermissionView, err := h.rolePermissionAppSvc.GetRolePermission(ctx, rolePermissionID)
	if err != nil {
		logx.S().Errorf("failed to get role permission by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "role permission not found: %v", err)
	}

	rolePermission := mapper.RolePermissionROToProto(&rolePermissionView)
	return &rolepermissionpb.GetRolePermissionByIDResponse{RolePermission: rolePermission}, nil
}

// GetPermissionsByRole 根据角色获取权限列表
func (h *RolePermissionHandler) GetPermissionsByRole(ctx context.Context, req *rolepermissionpb.GetPermissionsByRoleRequest) (*rolepermissionpb.GetPermissionsByRoleResponse, error) {
	roleID, err := uuid.Parse(req.RoleId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid role_id: %v", err)
	}

	rolePermissionViews, err := h.rolePermissionAppSvc.GetRolePermissionsByRoleID(ctx, roleID)
	if err != nil {
		logx.S().Errorf("failed to get role permissions by role_id: %v", err)
		return nil, status.Errorf(codes.NotFound, "role permissions not found: %v", err)
	}

	rolePermissions := mapper.RolePermissionListROToProto(rolePermissionViews)
	return &rolepermissionpb.GetPermissionsByRoleResponse{RolePermissions: rolePermissions}, nil
}

// GetRolesByPermission 根据权限获取角色列表
func (h *RolePermissionHandler) GetRolesByPermission(ctx context.Context, req *rolepermissionpb.GetRolesByPermissionRequest) (*rolepermissionpb.GetRolesByPermissionResponse, error) {
	permissionID, err := uuid.Parse(req.PermissionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid permission_id: %v", err)
	}

	rolePermissionViews, err := h.rolePermissionAppSvc.GetRolePermissionsByPermissionID(ctx, permissionID)
	if err != nil {
		logx.S().Errorf("failed to get role permissions by permission_id: %v", err)
		return nil, status.Errorf(codes.NotFound, "role permissions not found: %v", err)
	}

	rolePermissions := mapper.RolePermissionListROToProto(rolePermissionViews)
	return &rolepermissionpb.GetRolesByPermissionResponse{RolePermissions: rolePermissions}, nil
}
