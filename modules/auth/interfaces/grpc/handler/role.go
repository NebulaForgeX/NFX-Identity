package handler

import (
	"context"

	roleApp "nfxid/modules/auth/application/role"
	roleDomain "nfxid/modules/auth/domain/role"
	"nfxid/modules/auth/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	rolepb "nfxid/protos/gen/auth/role"

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

// GetRoleByID 根据ID获取角色
func (h *RoleHandler) GetRoleByID(ctx context.Context, req *rolepb.GetRoleByIDRequest) (*rolepb.GetRoleByIDResponse, error) {
	roleID, err := uuid.Parse(req.RoleId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid role_id: %v", err)
	}

	roleView, err := h.roleAppSvc.GetRole(ctx, roleID)
	if err != nil {
		logx.S().Errorf("failed to get role by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "role not found: %v", err)
	}

	role := mapper.RoleViewToProto(&roleView)
	return &rolepb.GetRoleByIDResponse{Role: role}, nil
}

// GetRoleByName 根据名称获取角色
func (h *RoleHandler) GetRoleByName(ctx context.Context, req *rolepb.GetRoleByNameRequest) (*rolepb.GetRoleByNameResponse, error) {
	roleView, err := h.roleAppSvc.GetRoleByName(ctx, req.Name)
	if err != nil {
		logx.S().Errorf("failed to get role by name: %v", err)
		return nil, status.Errorf(codes.NotFound, "role not found: %v", err)
	}

	role := mapper.RoleViewToProto(&roleView)
	return &rolepb.GetRoleByNameResponse{Role: role}, nil
}

// BatchGetRoles 批量获取角色
func (h *RoleHandler) BatchGetRoles(ctx context.Context, req *rolepb.BatchGetRolesRequest) (*rolepb.BatchGetRolesResponse, error) {
	roleIDs := make([]uuid.UUID, 0, len(req.RoleIds))
	for _, idStr := range req.RoleIds {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		roleIDs = append(roleIDs, id)
	}

	roles := make([]*rolepb.Role, 0, len(roleIDs))
	errorById := make(map[string]string)

	for _, roleID := range roleIDs {
		roleView, err := h.roleAppSvc.GetRole(ctx, roleID)
		if err != nil {
			errorById[roleID.String()] = err.Error()
			continue
		}
		role := mapper.RoleViewToProto(&roleView)
		roles = append(roles, role)
	}

	return &rolepb.BatchGetRolesResponse{
		Roles:     roles,
		ErrorById: errorById,
	}, nil
}

// GetAllRoles 获取所有角色列表
func (h *RoleHandler) GetAllRoles(ctx context.Context, req *rolepb.GetAllRolesRequest) (*rolepb.GetAllRolesResponse, error) {
	listQuery := roleDomain.ListQuery{}

	// Set pagination (convert Page/PageSize to Offset/Limit)
	if req.Page > 0 && req.PageSize > 0 {
		listQuery.Offset = int((req.Page - 1) * req.PageSize)
		listQuery.Limit = int(req.PageSize)
	}

	if req.Search != nil {
		search := *req.Search
		listQuery.Search = &search
	}
	if req.IsSystem != nil {
		isSystem := *req.IsSystem
		listQuery.IsSystem = &isSystem
	}

	result, err := h.roleAppSvc.GetRoleList(ctx, listQuery)
	if err != nil {
		logx.S().Errorf("failed to get all roles: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get roles: %v", err)
	}

	roles := make([]*rolepb.Role, len(result.Items))
	for i, roleView := range result.Items {
		roles[i] = mapper.RoleViewToProto(&roleView)
	}

	return &rolepb.GetAllRolesResponse{
		Roles: roles,
		Total: int32(result.Total),
	}, nil
}
