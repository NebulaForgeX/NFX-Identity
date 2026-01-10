package handler

import (
	"context"

	groupApp "nfxid/modules/tenants/application/groups"
	"nfxid/modules/tenants/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	grouppb "nfxid/protos/gen/tenants/group"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GroupHandler struct {
	grouppb.UnimplementedGroupServiceServer
	groupAppSvc *groupApp.Service
}

func NewGroupHandler(groupAppSvc *groupApp.Service) *GroupHandler {
	return &GroupHandler{
		groupAppSvc: groupAppSvc,
	}
}

// GetGroupByID 根据ID获取组
func (h *GroupHandler) GetGroupByID(ctx context.Context, req *grouppb.GetGroupByIDRequest) (*grouppb.GetGroupByIDResponse, error) {
	groupID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid group_id: %v", err)
	}

	groupView, err := h.groupAppSvc.GetGroup(ctx, groupID)
	if err != nil {
		logx.S().Errorf("failed to get group by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "group not found: %v", err)
	}

	group := mapper.GroupROToProto(&groupView)
	return &grouppb.GetGroupByIDResponse{Group: group}, nil
}

// GetGroupsByTenantID 根据租户ID获取组列表
func (h *GroupHandler) GetGroupsByTenantID(ctx context.Context, req *grouppb.GetGroupsByTenantIDRequest) (*grouppb.GetGroupsByTenantIDResponse, error) {
	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tenant_id: %v", err)
	}

	var parentID *uuid.UUID
	if req.ParentId != nil && *req.ParentId != "" {
		parsed, err := uuid.Parse(*req.ParentId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid parent_id: %v", err)
		}
		parentID = &parsed
	}

	groupViews, err := h.groupAppSvc.GetGroupsByTenantID(ctx, tenantID, parentID)
	if err != nil {
		logx.S().Errorf("failed to get groups by tenant_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get groups: %v", err)
	}

	groups := mapper.GroupListROToProto(groupViews)
	return &grouppb.GetGroupsByTenantIDResponse{Groups: groups}, nil
}

// BatchGetGroups 批量获取组
func (h *GroupHandler) BatchGetGroups(ctx context.Context, req *grouppb.BatchGetGroupsRequest) (*grouppb.BatchGetGroupsResponse, error) {
	groupIDs := make([]uuid.UUID, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		groupIDs = append(groupIDs, id)
	}

	groups := make([]*grouppb.Group, 0, len(groupIDs))
	for _, groupID := range groupIDs {
		groupView, err := h.groupAppSvc.GetGroup(ctx, groupID)
		if err != nil {
			logx.S().Warnf("failed to get group %s: %v", groupID, err)
			continue
		}
		group := mapper.GroupROToProto(&groupView)
		groups = append(groups, group)
	}

	return &grouppb.BatchGetGroupsResponse{Groups: groups}, nil
}
