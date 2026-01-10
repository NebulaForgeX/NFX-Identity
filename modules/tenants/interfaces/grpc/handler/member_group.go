package handler

import (
	"context"

	memberGroupApp "nfxid/modules/tenants/application/member_groups"
	"nfxid/modules/tenants/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	membergrouppb "nfxid/protos/gen/tenants/member_group"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MemberGroupHandler struct {
	membergrouppb.UnimplementedMemberGroupServiceServer
	memberGroupAppSvc *memberGroupApp.Service
}

func NewMemberGroupHandler(memberGroupAppSvc *memberGroupApp.Service) *MemberGroupHandler {
	return &MemberGroupHandler{
		memberGroupAppSvc: memberGroupAppSvc,
	}
}

// GetMemberGroupByID 根据ID获取成员组
func (h *MemberGroupHandler) GetMemberGroupByID(ctx context.Context, req *membergrouppb.GetMemberGroupByIDRequest) (*membergrouppb.GetMemberGroupByIDResponse, error) {
	memberGroupID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid member_group_id: %v", err)
	}

	memberGroupView, err := h.memberGroupAppSvc.GetMemberGroup(ctx, memberGroupID)
	if err != nil {
		logx.S().Errorf("failed to get member group by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "member group not found: %v", err)
	}

	memberGroup := mapper.MemberGroupROToProto(&memberGroupView)
	return &membergrouppb.GetMemberGroupByIDResponse{MemberGroup: memberGroup}, nil
}

// GetMemberGroupsByMemberID 根据成员ID获取成员组列表
func (h *MemberGroupHandler) GetMemberGroupsByMemberID(ctx context.Context, req *membergrouppb.GetMemberGroupsByMemberIDRequest) (*membergrouppb.GetMemberGroupsByMemberIDResponse, error) {
	memberID, err := uuid.Parse(req.MemberId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid member_id: %v", err)
	}

	memberGroupViews, err := h.memberGroupAppSvc.GetMemberGroupsByMemberID(ctx, memberID)
	if err != nil {
		logx.S().Errorf("failed to get member groups by member_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get member groups: %v", err)
	}

	memberGroups := mapper.MemberGroupListROToProto(memberGroupViews)
	return &membergrouppb.GetMemberGroupsByMemberIDResponse{MemberGroups: memberGroups}, nil
}

// GetMemberGroupsByGroupID 根据组ID获取成员组列表
func (h *MemberGroupHandler) GetMemberGroupsByGroupID(ctx context.Context, req *membergrouppb.GetMemberGroupsByGroupIDRequest) (*membergrouppb.GetMemberGroupsByGroupIDResponse, error) {
	groupID, err := uuid.Parse(req.GroupId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid group_id: %v", err)
	}

	memberGroupViews, err := h.memberGroupAppSvc.GetMemberGroupsByGroupID(ctx, groupID)
	if err != nil {
		logx.S().Errorf("failed to get member groups by group_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get member groups: %v", err)
	}

	memberGroups := mapper.MemberGroupListROToProto(memberGroupViews)
	return &membergrouppb.GetMemberGroupsByGroupIDResponse{MemberGroups: memberGroups}, nil
}
