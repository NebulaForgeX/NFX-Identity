package handler

import (
	"context"

	memberApp "nfxid/modules/tenants/application/members"
	memberDomain "nfxid/modules/tenants/domain/members"
	"nfxid/modules/tenants/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	memberpb "nfxid/protos/gen/tenants/member"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MemberHandler struct {
	memberpb.UnimplementedMemberServiceServer
	memberAppSvc *memberApp.Service
}

func NewMemberHandler(memberAppSvc *memberApp.Service) *MemberHandler {
	return &MemberHandler{
		memberAppSvc: memberAppSvc,
	}
}

// GetMemberByID 根据ID获取成员
func (h *MemberHandler) GetMemberByID(ctx context.Context, req *memberpb.GetMemberByIDRequest) (*memberpb.GetMemberByIDResponse, error) {
	memberID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid member_id: %v", err)
	}

	memberView, err := h.memberAppSvc.GetMember(ctx, memberID)
	if err != nil {
		logx.S().Errorf("failed to get member by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "member not found: %v", err)
	}

	member := mapper.MemberROToProto(&memberView)
	return &memberpb.GetMemberByIDResponse{Member: member}, nil
}

// GetMemberByUserID 根据用户ID和租户ID获取成员
func (h *MemberHandler) GetMemberByUserID(ctx context.Context, req *memberpb.GetMemberByUserIDRequest) (*memberpb.GetMemberByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tenant_id: %v", err)
	}

	memberView, err := h.memberAppSvc.GetMemberByUserID(ctx, userID, tenantID)
	if err != nil {
		logx.S().Errorf("failed to get member by user_id: %v", err)
		return nil, status.Errorf(codes.NotFound, "member not found: %v", err)
	}

	member := mapper.MemberROToProto(&memberView)
	return &memberpb.GetMemberByUserIDResponse{Member: member}, nil
}

// GetMembersByTenantID 根据租户ID获取成员列表
func (h *MemberHandler) GetMembersByTenantID(ctx context.Context, req *memberpb.GetMembersByTenantIDRequest) (*memberpb.GetMembersByTenantIDResponse, error) {
	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tenant_id: %v", err)
	}

	var memberStatus *memberDomain.MemberStatus
	if req.Status != nil {
		statusVal := protoMemberStatusToDomain(*req.Status)
		memberStatus = &statusVal
	}

	memberViews, err := h.memberAppSvc.GetMembersByTenantID(ctx, tenantID, memberStatus)
	if err != nil {
		logx.S().Errorf("failed to get members by tenant_id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get members: %v", err)
	}

	members := mapper.MemberListROToProto(memberViews)
	return &memberpb.GetMembersByTenantIDResponse{Members: members}, nil
}

// BatchGetMembers 批量获取成员
func (h *MemberHandler) BatchGetMembers(ctx context.Context, req *memberpb.BatchGetMembersRequest) (*memberpb.BatchGetMembersResponse, error) {
	memberIDs := make([]uuid.UUID, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		memberIDs = append(memberIDs, id)
	}

	members := make([]*memberpb.Member, 0, len(memberIDs))
	for _, memberID := range memberIDs {
		memberView, err := h.memberAppSvc.GetMember(ctx, memberID)
		if err != nil {
			logx.S().Warnf("failed to get member %s: %v", memberID, err)
			continue
		}
		member := mapper.MemberROToProto(&memberView)
		members = append(members, member)
	}

	return &memberpb.BatchGetMembersResponse{Members: members}, nil
}

// protoMemberStatusToDomain 将 proto TenantsMemberStatus 转换为 domain MemberStatus
func protoMemberStatusToDomain(status memberpb.TenantsMemberStatus) memberDomain.MemberStatus {
	switch status {
	case memberpb.TenantsMemberStatus_TENANTS_MEMBER_STATUS_INVITED:
		return memberDomain.MemberStatusInvited
	case memberpb.TenantsMemberStatus_TENANTS_MEMBER_STATUS_ACTIVE:
		return memberDomain.MemberStatusActive
	case memberpb.TenantsMemberStatus_TENANTS_MEMBER_STATUS_SUSPENDED:
		return memberDomain.MemberStatusSuspended
	case memberpb.TenantsMemberStatus_TENANTS_MEMBER_STATUS_REMOVED:
		return memberDomain.MemberStatusRemoved
	default:
		return memberDomain.MemberStatusInvited
	}
}
