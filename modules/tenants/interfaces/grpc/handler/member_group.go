package handler

import (
	"context"

	memberGroupApp "nfxid/modules/tenants/application/member_groups"
	"nfxid/modules/tenants/interfaces/grpc/mapper"
	membergrouppb "nfxid/protos/gen/tenants/member_group"
	"nfxid/pkgs/errx"

	"github.com/google/uuid"
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
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	memberGroupView, err := h.memberGroupAppSvc.GetMemberGroup(ctx, memberGroupID)
	if err != nil {
		return nil, err
	}

	memberGroup := mapper.MemberGroupROToProto(&memberGroupView)
	return &membergrouppb.GetMemberGroupByIDResponse{MemberGroup: memberGroup}, nil
}

// GetMemberGroupsByMemberID 根据成员ID获取成员组列表
func (h *MemberGroupHandler) GetMemberGroupsByMemberID(ctx context.Context, req *membergrouppb.GetMemberGroupsByMemberIDRequest) (*membergrouppb.GetMemberGroupsByMemberIDResponse, error) {
	memberID, err := uuid.Parse(req.MemberId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	memberGroupViews, err := h.memberGroupAppSvc.GetMemberGroupsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}

	memberGroups := mapper.MemberGroupListROToProto(memberGroupViews)
	return &membergrouppb.GetMemberGroupsByMemberIDResponse{MemberGroups: memberGroups}, nil
}

// GetMemberGroupsByGroupID 根据组ID获取成员组列表
func (h *MemberGroupHandler) GetMemberGroupsByGroupID(ctx context.Context, req *membergrouppb.GetMemberGroupsByGroupIDRequest) (*membergrouppb.GetMemberGroupsByGroupIDResponse, error) {
	groupID, err := uuid.Parse(req.GroupId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	memberGroupViews, err := h.memberGroupAppSvc.GetMemberGroupsByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}

	memberGroups := mapper.MemberGroupListROToProto(memberGroupViews)
	return &membergrouppb.GetMemberGroupsByGroupIDResponse{MemberGroups: memberGroups}, nil
}
