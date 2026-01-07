package handler

import (
	"context"
	"nfxid/enums"

	permissionApp "nfxid/modules/permission/application/permission"
	permissionAppCommands "nfxid/modules/permission/application/permission/commands"
	"nfxid/modules/permission/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	permissionpb "nfxid/protos/gen/permission/permission"

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

// GetPermissionByID 根据ID获取权限
func (h *PermissionHandler) GetPermissionByID(ctx context.Context, req *permissionpb.GetPermissionByIDRequest) (*permissionpb.GetPermissionByIDResponse, error) {
	permissionID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid permission_id: %v", err)
	}

	permissionView, err := h.permissionAppSvc.GetPermission(ctx, permissionAppCommands.GetPermissionCmd{
		ID: permissionID,
	})
	if err != nil {
		logx.S().Errorf("failed to get permission by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "permission not found: %v", err)
	}

	permission := mapper.PermissionViewToProto(permissionView)
	return &permissionpb.GetPermissionByIDResponse{Permission: permission}, nil
}

// GetPermissionByTag 根据Tag获取权限
func (h *PermissionHandler) GetPermissionByTag(ctx context.Context, req *permissionpb.GetPermissionByTagRequest) (*permissionpb.GetPermissionByTagResponse, error) {
	permissionView, err := h.permissionAppSvc.GetPermissionByTag(ctx, permissionAppCommands.GetPermissionByTagCmd{
		Tag: req.Tag,
	})
	if err != nil {
		logx.S().Errorf("failed to get permission by tag: %v", err)
		return nil, status.Errorf(codes.NotFound, "permission not found: %v", err)
	}

	permission := mapper.PermissionViewToProto(permissionView)
	return &permissionpb.GetPermissionByTagResponse{Permission: permission}, nil
}

// GetPermissionsByTags 根据Tag列表批量获取权限
func (h *PermissionHandler) GetPermissionsByTags(ctx context.Context, req *permissionpb.GetPermissionsByTagsRequest) (*permissionpb.GetPermissionsByTagsResponse, error) {
	permissions, err := h.permissionAppSvc.ListPermissions(ctx, permissionAppCommands.ListPermissionsCmd{})
	if err != nil {
		logx.S().Errorf("failed to get permissions by tags: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get permissions: %v", err)
	}

	// Filter by tags
	tagMap := make(map[string]bool)
	for _, tag := range req.Tags {
		tagMap[tag] = true
	}

	result := make([]*permissionpb.Permission, 0)
	for _, p := range permissions {
		if tagMap[p.Tag] {
			result = append(result, mapper.PermissionViewToProto(p))
		}
	}

	return &permissionpb.GetPermissionsByTagsResponse{Permissions: result}, nil
}

// GetPermissionsByCategory 根据Category获取权限列表
func (h *PermissionHandler) GetPermissionsByCategory(ctx context.Context, req *permissionpb.GetPermissionsByCategoryRequest) (*permissionpb.GetPermissionsByCategoryResponse, error) {
	permissions, err := h.permissionAppSvc.ListPermissions(ctx, permissionAppCommands.ListPermissionsCmd{
		Category: enums.PermissionCategory(req.Category), // Convert string to enum
	})
	if err != nil {
		logx.S().Errorf("failed to get permissions by category: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get permissions: %v", err)
	}

	result := make([]*permissionpb.Permission, len(permissions))
	for i, p := range permissions {
		result[i] = mapper.PermissionViewToProto(p)
	}

	return &permissionpb.GetPermissionsByCategoryResponse{Permissions: result}, nil
}

// GetAllPermissions 获取所有权限列表（支持分页和过滤）
func (h *PermissionHandler) GetAllPermissions(ctx context.Context, req *permissionpb.GetAllPermissionsRequest) (*permissionpb.GetAllPermissionsResponse, error) {
	cmd := permissionAppCommands.ListPermissionsCmd{}
	if req.Category != nil {
		cmd.Category = enums.PermissionCategory(*req.Category) // Convert string to enum
	}

	permissions, err := h.permissionAppSvc.ListPermissions(ctx, cmd)
	if err != nil {
		logx.S().Errorf("failed to get all permissions: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get permissions: %v", err)
	}

	// Apply pagination if needed
	total := int32(len(permissions))
	start := int32(0)
	end := total

	if req.Page > 0 && req.PageSize > 0 {
		start = (req.Page - 1) * req.PageSize
		end = start + req.PageSize
		if end > total {
			end = total
		}
	}

	result := make([]*permissionpb.Permission, 0)
	for i := start; i < end; i++ {
		result = append(result, mapper.PermissionViewToProto(permissions[i]))
	}

	return &permissionpb.GetAllPermissionsResponse{
		Permissions: result,
		Total:       total,
	}, nil
}
