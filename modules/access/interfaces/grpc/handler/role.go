package handler

import (
	"context"

	roleApp "nfxid/modules/access/application/roles"
	roleAppCommands "nfxid/modules/access/application/roles/commands"
	roleDomain "nfxid/modules/access/domain/roles"
	"nfxid/modules/access/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	rolepb "nfxid/protos/gen/access/role"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoleHandler struct {
	rolepb.UnimplementedRoleServiceServer
	roleAppSvc *roleApp.Service
}

func NewRoleHandler(roleAppSvc *roleApp.Service) *RoleHandler {
	return &RoleHandler{
		roleAppSvc: roleAppSvc,
	}
}

// CreateRole 创建角色
func (h *RoleHandler) CreateRole(ctx context.Context, req *rolepb.CreateRoleRequest) (*rolepb.CreateRoleResponse, error) {
	// 转换 protobuf AccessScopeType 到 domain ScopeType
	scopeType := protoScopeTypeToDomain(req.ScopeType)

	// 创建命令
	cmd := roleAppCommands.CreateRoleCmd{
		Key:         req.Key,
		Name:        req.Name,
		Description: req.Description,
		ScopeType:   scopeType,
		IsSystem:    req.IsSystem,
	}

	// 调用应用服务创建角色
	roleID, err := h.roleAppSvc.CreateRole(ctx, cmd)
	if err != nil {
		logx.S().Errorf("failed to create role: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create role: %v", err)
	}

	// 获取创建的角色
	roleView, err := h.roleAppSvc.GetRole(ctx, roleID)
	if err != nil {
		logx.S().Errorf("failed to get created role: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get created role: %v", err)
	}

	// 转换为 protobuf 响应
	role := mapper.RoleROToProto(&roleView)
	return &rolepb.CreateRoleResponse{Role: role}, nil
}

// protoScopeTypeToDomain 将 protobuf AccessScopeType 转换为 domain ScopeType
func protoScopeTypeToDomain(scopeType rolepb.AccessScopeType) roleDomain.ScopeType {
	switch scopeType {
	case rolepb.AccessScopeType_ACCESS_SCOPE_TYPE_TENANT:
		return roleDomain.ScopeTypeTenant
	case rolepb.AccessScopeType_ACCESS_SCOPE_TYPE_APP:
		return roleDomain.ScopeTypeApp
	case rolepb.AccessScopeType_ACCESS_SCOPE_TYPE_GLOBAL:
		return roleDomain.ScopeTypeGlobal
	default:
		return roleDomain.ScopeTypeTenant // 默认值
	}
}

// GetRoleByID 根据ID获取角色
func (h *RoleHandler) GetRoleByID(ctx context.Context, req *rolepb.GetRoleByIDRequest) (*rolepb.GetRoleByIDResponse, error) {
	roleID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid role_id: %v", err)
	}

	roleView, err := h.roleAppSvc.GetRole(ctx, roleID)
	if err != nil {
		logx.S().Errorf("failed to get role by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "role not found: %v", err)
	}

	role := mapper.RoleROToProto(&roleView)
	return &rolepb.GetRoleByIDResponse{Role: role}, nil
}

// GetRoleByKey 根据Key获取角色
func (h *RoleHandler) GetRoleByKey(ctx context.Context, req *rolepb.GetRoleByKeyRequest) (*rolepb.GetRoleByKeyResponse, error) {
	roleView, err := h.roleAppSvc.GetRoleByKey(ctx, req.Key)
	if err != nil {
		logx.S().Errorf("failed to get role by key: %v", err)
		return nil, status.Errorf(codes.NotFound, "role not found: %v", err)
	}

	role := mapper.RoleROToProto(&roleView)
	return &rolepb.GetRoleByKeyResponse{Role: role}, nil
}

// GetAllRoles 获取所有角色列表
func (h *RoleHandler) GetAllRoles(ctx context.Context, req *rolepb.GetAllRolesRequest) (*rolepb.GetAllRolesResponse, error) {
	// 注意：domain repository 目前没有 GetAll 方法，需要先实现 repository 层的 GetAll 方法
	// 然后才能根据 scope_type 和 is_system 进行过滤
	// 目前返回空列表，待 repository 层实现后补充
	roles := []*rolepb.Role{}
	return &rolepb.GetAllRolesResponse{Roles: roles}, nil
}

// BatchGetRoles 批量获取角色
func (h *RoleHandler) BatchGetRoles(ctx context.Context, req *rolepb.BatchGetRolesRequest) (*rolepb.BatchGetRolesResponse, error) {
	roleIDs := make([]uuid.UUID, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		roleIDs = append(roleIDs, id)
	}

	roles := make([]*rolepb.Role, 0, len(roleIDs))
	for _, roleID := range roleIDs {
		roleView, err := h.roleAppSvc.GetRole(ctx, roleID)
		if err != nil {
			logx.S().Warnf("failed to get role %s: %v", roleID, err)
			continue
		}
		role := mapper.RoleROToProto(&roleView)
		roles = append(roles, role)
	}

	return &rolepb.BatchGetRolesResponse{Roles: roles}, nil
}
