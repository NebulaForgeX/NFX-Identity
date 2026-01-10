package mapper

import (
	invitationAppResult "nfxid/modules/tenants/application/invitations/results"
	invitationDomain "nfxid/modules/tenants/domain/invitations"
	invitationpb "nfxid/protos/gen/tenants/invitation"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// InvitationROToProto 将 InvitationRO 转换为 proto Invitation 消息
func InvitationROToProto(v *invitationAppResult.InvitationRO) *invitationpb.Invitation {
	if v == nil {
		return nil
	}

	invitation := &invitationpb.Invitation{
		Id:         v.ID.String(),
		InviteId:   v.InviteID,
		TenantId:   v.TenantID.String(),
		Email:      v.Email,
		TokenHash:  v.TokenHash,
		ExpiresAt:  timestamppb.New(v.ExpiresAt),
		Status:     invitationStatusToProto(v.Status),
		InvitedBy:  v.InvitedBy.String(),
		InvitedAt:  timestamppb.New(v.InvitedAt),
	}

	if v.AcceptedByUserID != nil {
		acceptedBy := v.AcceptedByUserID.String()
		invitation.AcceptedByUserId = &acceptedBy
	}

	if v.AcceptedAt != nil {
		invitation.AcceptedAt = timestamppb.New(*v.AcceptedAt)
	}

	if v.RevokedBy != nil {
		revokedBy := v.RevokedBy.String()
		invitation.RevokedBy = &revokedBy
	}

	if v.RevokedAt != nil {
		invitation.RevokedAt = timestamppb.New(*v.RevokedAt)
	}

	if v.RevokeReason != nil {
		invitation.RevokeReason = v.RevokeReason
	}

	if v.RoleIDs != nil && len(v.RoleIDs) > 0 {
		roleIDs := make([]string, len(v.RoleIDs))
		for i, roleID := range v.RoleIDs {
			roleIDs[i] = roleID.String()
		}
		invitation.RoleIds = roleIDs
	}

	if v.Metadata != nil && len(v.Metadata) > 0 {
		if metadataStruct, err := structpb.NewStruct(v.Metadata); err == nil {
			invitation.Metadata = metadataStruct
		}
	}

	return invitation
}

// InvitationListROToProto 批量转换 InvitationRO 到 proto Invitation
func InvitationListROToProto(results []invitationAppResult.InvitationRO) []*invitationpb.Invitation {
	invitations := make([]*invitationpb.Invitation, len(results))
	for i, v := range results {
		invitations[i] = InvitationROToProto(&v)
	}
	return invitations
}

// invitationStatusToProto 将 domain InvitationStatus 转换为 proto TenantsInvitationStatus
func invitationStatusToProto(status invitationDomain.InvitationStatus) invitationpb.TenantsInvitationStatus {
	switch status {
	case invitationDomain.InvitationStatusPending:
		return invitationpb.TenantsInvitationStatus_TENANTS_INVITATION_STATUS_PENDING
	case invitationDomain.InvitationStatusAccepted:
		return invitationpb.TenantsInvitationStatus_TENANTS_INVITATION_STATUS_ACCEPTED
	case invitationDomain.InvitationStatusExpired:
		return invitationpb.TenantsInvitationStatus_TENANTS_INVITATION_STATUS_EXPIRED
	case invitationDomain.InvitationStatusRevoked:
		return invitationpb.TenantsInvitationStatus_TENANTS_INVITATION_STATUS_REVOKED
	default:
		return invitationpb.TenantsInvitationStatus_TENANTS_INVITATION_STATUS_UNSPECIFIED
	}
}
