package handler

import (
	"context"

	invitationApp "nfxid/modules/tenants/application/invitations"
	invitationDomain "nfxid/modules/tenants/domain/invitations"
	"nfxid/modules/tenants/interfaces/grpc/mapper"
	"nfxid/pkgs/errx"
	invitationpb "nfxid/protos/gen/tenants/invitation"

	"github.com/google/uuid"
)

type InvitationHandler struct {
	invitationpb.UnimplementedInvitationServiceServer
	invitationAppSvc *invitationApp.Service
}

func NewInvitationHandler(invitationAppSvc *invitationApp.Service) *InvitationHandler {
	return &InvitationHandler{
		invitationAppSvc: invitationAppSvc,
	}
}

// GetInvitationByID 根据ID获取邀请
func (h *InvitationHandler) GetInvitationByID(
	ctx context.Context,
	req *invitationpb.GetInvitationByIDRequest,
) (*invitationpb.GetInvitationByIDResponse, error) {
	invitationID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	invitationView, err := h.invitationAppSvc.GetInvitation(ctx, invitationID)
	if err != nil {
		return nil, err
	}

	invitation := mapper.InvitationROToProto(&invitationView)
	return &invitationpb.GetInvitationByIDResponse{Invitation: invitation}, nil
}

// GetInvitationByInvitationID 根据邀请ID获取邀请
func (h *InvitationHandler) GetInvitationByInvitationID(
	ctx context.Context,
	req *invitationpb.GetInvitationByInvitationIDRequest,
) (*invitationpb.GetInvitationByInvitationIDResponse, error) {
	invitationView, err := h.invitationAppSvc.GetInvitationByInviteID(ctx, req.InvitationId)
	if err != nil {
		return nil, err
	}

	invitation := mapper.InvitationROToProto(&invitationView)
	return &invitationpb.GetInvitationByInvitationIDResponse{Invitation: invitation}, nil
}

// GetInvitationsByTenantID 根据租户ID获取邀请列表
func (h *InvitationHandler) GetInvitationsByTenantID(
	ctx context.Context,
	req *invitationpb.GetInvitationsByTenantIDRequest,
) (*invitationpb.GetInvitationsByTenantIDResponse, error) {
	tenantID, err := uuid.Parse(req.TenantId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	var invitationStatus *invitationDomain.InvitationStatus
	if req.Status != nil {
		statusVal := protoInvitationStatusToDomain(*req.Status)
		invitationStatus = &statusVal
	}

	invitationViews, err := h.invitationAppSvc.GetInvitationsByTenantID(ctx, tenantID, invitationStatus)
	if err != nil {
		return nil, err
	}

	invitations := mapper.InvitationListROToProto(invitationViews)
	return &invitationpb.GetInvitationsByTenantIDResponse{Invitations: invitations}, nil
}

// protoInvitationStatusToDomain 将 proto TenantsInvitationStatus 转换为 domain InvitationStatus
func protoInvitationStatusToDomain(status invitationpb.TenantsInvitationStatus) invitationDomain.InvitationStatus {
	switch status {
	case invitationpb.TenantsInvitationStatus_TENANTS_INVITATION_STATUS_PENDING:
		return invitationDomain.InvitationStatusPending
	case invitationpb.TenantsInvitationStatus_TENANTS_INVITATION_STATUS_ACCEPTED:
		return invitationDomain.InvitationStatusAccepted
	case invitationpb.TenantsInvitationStatus_TENANTS_INVITATION_STATUS_EXPIRED:
		return invitationDomain.InvitationStatusExpired
	case invitationpb.TenantsInvitationStatus_TENANTS_INVITATION_STATUS_REVOKED:
		return invitationDomain.InvitationStatusRevoked
	default:
		return invitationDomain.InvitationStatusPending
	}
}
