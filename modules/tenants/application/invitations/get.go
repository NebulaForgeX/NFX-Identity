package invitations

import (
	"context"
	invitationResult "nfxid/modules/tenants/application/invitations/results"

	"github.com/google/uuid"
)

// GetInvitation 根据ID获取邀请
func (s *Service) GetInvitation(ctx context.Context, invitationID uuid.UUID) (invitationResult.InvitationRO, error) {
	domainEntity, err := s.invitationRepo.Get.ByID(ctx, invitationID)
	if err != nil {
		return invitationResult.InvitationRO{}, err
	}
	return invitationResult.InvitationMapper(domainEntity), nil
}

// GetInvitationByInviteID 根据InviteID获取邀请
func (s *Service) GetInvitationByInviteID(ctx context.Context, inviteID string) (invitationResult.InvitationRO, error) {
	domainEntity, err := s.invitationRepo.Get.ByInviteID(ctx, inviteID)
	if err != nil {
		return invitationResult.InvitationRO{}, err
	}
	return invitationResult.InvitationMapper(domainEntity), nil
}
