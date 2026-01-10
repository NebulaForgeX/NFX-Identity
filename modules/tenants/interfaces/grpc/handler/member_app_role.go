package handler

import (
	"context"

	memberAppRoleApp "nfxid/modules/tenants/application/member_app_roles"
	"nfxid/modules/tenants/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	memberapprolepb "nfxid/protos/gen/tenants/member_app_role"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MemberAppRoleHandler struct {
	memberapprolepb.UnimplementedMemberAppRoleServiceServer
	memberAppRoleAppSvc *memberAppRoleApp.Service
}

func NewMemberAppRoleHandler(memberAppRoleAppSvc *memberAppRoleApp.Service) *MemberAppRoleHandler {
	return &MemberAppRoleHandler{
		memberAppRoleAppSvc: memberAppRoleAppSvc,
	}
}

// GetMemberAppRoleByID 根据ID获取成员应用角色
func (h *MemberAppRoleHandler) GetMemberAppRoleByID(ctx context.Context, req *memberapprolepb.GetMemberAppRoleByIDRequest) (*memberapprolepb.GetMemberAppRoleByIDResponse, error) {
	memberAppRoleID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid member_app_role_id: %v", err)
	}

	memberAppRoleView, err := h.memberAppRoleAppSvc.GetMemberAppRole(ctx, memberAppRoleID)
	if err != nil {
		logx.S().Errorf("failed to get member app role by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "member app role not found: %v", err)
	}

	memberAppRole := mapper.MemberAppRoleROToProto(&memberAppRoleView)
	return &memberapprolepb.GetMemberAppRoleByIDResponse{MemberAppRole: memberAppRole}, nil
}

// GetMemberAppRolesByMemberID 根据成员ID获取成员应用角色列表
func (h *MemberAppRoleHandler) GetMemberAppRolesByMemberID(ctx context.Context, req *memberapprolepb.GetMemberAppRolesByMemberIDRequest) (*memberapprolepb.GetMemberAppRolesByMemberIDResponse, error) {
	memberID, err := uuid.Parse(req.MemberId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid member_id: %v", err)
	}

	var appID *uuid.UUID
	if req.AppId != nil && *req.AppId != "" {
		parsed, err := uuid.Parse(*req.AppId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid app_id: %v", err)
		}
		appID = &parsed
	}

	memberAppRoleViews, err := h.memberAppRoleAppSvc.GetMemberAppRolesByMemberID(ctx, memberID, appID)
	if err != nil {
		logx.S().Errorf("failed to get member app roles by member_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get member app roles: %v", err)
	}

	memberAppRoles := mapper.MemberAppRoleListROToProto(memberAppRoleViews)
	return &memberapprolepb.GetMemberAppRolesByMemberIDResponse{MemberAppRoles: memberAppRoles}, nil
}
