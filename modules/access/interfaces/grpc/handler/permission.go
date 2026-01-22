package handler

import (
	"context"

	permissionApp "nfxid/modules/access/application/permissions"
	permissionAppCommands "nfxid/modules/access/application/permissions/commands"
	"nfxid/modules/access/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	permissionpb "nfxid/protos/gen/access/permission"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PermissionHandler struct {
	permissionpb.UnimplementedPermissionServiceServer
	permissionAppSvc *permissionApp.Service
}

func NewPermissionHandler(permissionAppSvc *permissionApp.Service) *PermissionHandler {
	return &PermissionHandler{
		permissionAppSvc: permissionAppSvc,
	}
}

// CreatePermission 创建权限
func (h *PermissionHandler) CreatePermission(ctx context.Context, req *permissionpb.CreatePermissionRequest) (*permissionpb.CreatePermissionResponse, error) {
	// 创建命令
	cmd := permissionAppCommands.CreatePermissionCmd{
		Key:         req.Key,
		Name:        req.Name,
		Description: req.Description,
		IsSystem:    req.IsSystem,
	}

	// 调用应用服务创建权限
	permissionID, err := h.permissionAppSvc.CreatePermission(ctx, cmd)
	if err != nil {
		logx.S().Errorf("failed to create permission: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create permission: %v", err)
	}

	// 获取创建的权限
	permissionView, err := h.permissionAppSvc.GetPermission(ctx, permissionID)
	if err != nil {
		logx.S().Errorf("failed to get created permission: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get created permission: %v", err)
	}

	// 转换为 protobuf 响应
	permission := mapper.PermissionROToProto(&permissionView)
	return &permissionpb.CreatePermissionResponse{Permission: permission}, nil
}

// GetPermissionByID 根据ID获取权限
func (h *PermissionHandler) GetPermissionByID(ctx context.Context, req *permissionpb.GetPermissionByIDRequest) (*permissionpb.GetPermissionByIDResponse, error) {
	permissionID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid permission_id: %v", err)
	}

	permissionView, err := h.permissionAppSvc.GetPermission(ctx, permissionID)
	if err != nil {
		logx.S().Errorf("failed to get permission by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "permission not found: %v", err)
	}

	permission := mapper.PermissionROToProto(&permissionView)
	return &permissionpb.GetPermissionByIDResponse{Permission: permission}, nil
}

// GetPermissionByKey 根据Key获取权限
func (h *PermissionHandler) GetPermissionByKey(ctx context.Context, req *permissionpb.GetPermissionByKeyRequest) (*permissionpb.GetPermissionByKeyResponse, error) {
	permissionView, err := h.permissionAppSvc.GetPermissionByKey(ctx, req.Key)
	if err != nil {
		logx.S().Errorf("failed to get permission by key: %v", err)
		return nil, status.Errorf(codes.NotFound, "permission not found: %v", err)
	}

	permission := mapper.PermissionROToProto(&permissionView)
	return &permissionpb.GetPermissionByKeyResponse{Permission: permission}, nil
}

// GetAllPermissions 获取所有权限列表
func (h *PermissionHandler) GetAllPermissions(ctx context.Context, req *permissionpb.GetAllPermissionsRequest) (*permissionpb.GetAllPermissionsResponse, error) {
	// 注意：domain repository 目前没有 GetAll 方法，需要先实现 repository 层的 GetAll 方法
	// 然后才能根据 is_system 进行过滤
	// 目前返回空列表，待 repository 层实现后补充
	permissions := []*permissionpb.Permission{}
	return &permissionpb.GetAllPermissionsResponse{Permissions: permissions}, nil
}

// BatchGetPermissions 批量获取权限
func (h *PermissionHandler) BatchGetPermissions(ctx context.Context, req *permissionpb.BatchGetPermissionsRequest) (*permissionpb.BatchGetPermissionsResponse, error) {
	permissionIDs := make([]uuid.UUID, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		permissionIDs = append(permissionIDs, id)
	}

	permissions := make([]*permissionpb.Permission, 0, len(permissionIDs))
	for _, permissionID := range permissionIDs {
		permissionView, err := h.permissionAppSvc.GetPermission(ctx, permissionID)
		if err != nil {
			logx.S().Warnf("failed to get permission %s: %v", permissionID, err)
			continue
		}
		permission := mapper.PermissionROToProto(&permissionView)
		permissions = append(permissions, permission)
	}

	return &permissionpb.BatchGetPermissionsResponse{Permissions: permissions}, nil
}
