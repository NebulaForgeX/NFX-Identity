package handler

import (
	"context"

	userPermissionApp "nfxid/modules/permission/application/user_permission"
	userPermissionAppCommands "nfxid/modules/permission/application/user_permission/commands"
	"nfxid/modules/permission/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	userpermissionpb "nfxid/protos/gen/permission/user_permission"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserPermissionHandler struct {
	userpermissionpb.UnimplementedUserPermissionServiceServer
	userPermissionAppSvc *userPermissionApp.Service
}

func NewUserPermissionHandler(userPermissionAppSvc *userPermissionApp.Service) *UserPermissionHandler {
	return &UserPermissionHandler{
		userPermissionAppSvc: userPermissionAppSvc,
	}
}

// GetUserPermissionsByUserID 根据UserID获取用户权限列表
func (h *UserPermissionHandler) GetUserPermissionsByUserID(ctx context.Context, req *userpermissionpb.GetUserPermissionsByUserIDRequest) (*userpermissionpb.GetUserPermissionsByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	includePermission := req.IncludePermission != nil && *req.IncludePermission

	permissions, err := h.userPermissionAppSvc.GetUserPermissions(ctx, userPermissionAppCommands.GetUserPermissionsCmd{
		UserID: userID,
	})
	if err != nil {
		logx.S().Errorf("failed to get user permissions: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get user permissions: %v", err)
	}

	result := make([]*userpermissionpb.UserPermission, len(permissions))
	for i, p := range permissions {
		result[i] = mapper.UserPermissionViewToProto(p, includePermission)
	}

	return &userpermissionpb.GetUserPermissionsByUserIDResponse{UserPermissions: result}, nil
}

// GetPermissionTagsByUserID 根据UserID获取权限标签列表
func (h *UserPermissionHandler) GetPermissionTagsByUserID(ctx context.Context, req *userpermissionpb.GetPermissionTagsByUserIDRequest) (*userpermissionpb.GetPermissionTagsByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	tags, err := h.userPermissionAppSvc.GetUserPermissionTags(ctx, userPermissionAppCommands.GetUserPermissionsCmd{
		UserID: userID,
	})
	if err != nil {
		logx.S().Errorf("failed to get user permission tags: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get user permission tags: %v", err)
	}

	return &userpermissionpb.GetPermissionTagsByUserIDResponse{Tags: tags}, nil
}

// CheckUserPermission 检查用户是否拥有指定权限标签
func (h *UserPermissionHandler) CheckUserPermission(ctx context.Context, req *userpermissionpb.CheckUserPermissionRequest) (*userpermissionpb.CheckUserPermissionResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	hasPermission, err := h.userPermissionAppSvc.CheckPermission(ctx, userPermissionAppCommands.CheckPermissionCmd{
		UserID: userID,
		Tag:    req.Tag,
	})
	if err != nil {
		logx.S().Errorf("failed to check user permission: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to check permission: %v", err)
	}

	return &userpermissionpb.CheckUserPermissionResponse{HasPermission: hasPermission}, nil
}

