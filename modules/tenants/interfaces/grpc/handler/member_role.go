package handler

import (
	"context"

	memberRoleApp "nfxid/modules/tenants/application/member_roles"
	"nfxid/modules/tenants/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	memberrolepb "nfxid/protos/gen/tenants/member_role"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MemberRoleHandler struct {
	memberrolepb.UnimplementedMemberRoleServiceServer
	memberRoleAppSvc *memberRoleApp.Service
}

func NewMemberRoleHandler(memberRoleAppSvc *memberRoleApp.Service) *MemberRoleHandler {
	return &MemberRoleHandler{
		memberRoleAppSvc: memberRoleAppSvc,
	}
}

// GetMemberRoleByID 根据ID获取成员角色
func (h *MemberRoleHandler) GetMemberRoleByID(ctx context.Context, req *memberrolepb.GetMemberRoleByIDRequest) (*memberrolepb.GetMemberRoleByIDResponse, error) {
	memberRoleID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid member_role_id: %v", err)
	}

	memberRoleView, err := h.memberRoleAppSvc.GetMemberRole(ctx, memberRoleID)
	if err != nil {
		logx.S().Errorf("failed to get member role by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "member role not found: %v", err)
	}

	memberRole := mapper.MemberRoleROToProto(&memberRoleView)
	return &memberrolepb.GetMemberRoleByIDResponse{MemberRole: memberRole}, nil
}

// GetMemberRolesByMemberID 根据成员ID获取成员角色列表
func (h *MemberRoleHandler) GetMemberRolesByMemberID(ctx context.Context, req *memberrolepb.GetMemberRolesByMemberIDRequest) (*memberrolepb.GetMemberRolesByMemberIDResponse, error) {
	memberID, err := uuid.Parse(req.MemberId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid member_id: %v", err)
	}

	memberRoleViews, err := h.memberRoleAppSvc.GetMemberRolesByMemberID(ctx, memberID)
	if err != nil {
		logx.S().Errorf("failed to get member roles by member_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get member roles: %v", err)
	}

	memberRoles := mapper.MemberRoleListROToProto(memberRoleViews)
	return &memberrolepb.GetMemberRolesByMemberIDResponse{MemberRoles: memberRoles}, nil
}
