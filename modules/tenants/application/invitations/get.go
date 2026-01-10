package invitations

import (
	"context"
	invitationDomain "nfxid/modules/tenants/domain/invitations"
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

// GetInvitationsByTenantID 根据租户ID获取邀请列表
func (s *Service) GetInvitationsByTenantID(ctx context.Context, tenantID uuid.UUID, status *invitationDomain.InvitationStatus) ([]invitationResult.InvitationRO, error) {
	var domainEntities []*invitationDomain.Invitation
	var err error
	
	if status != nil {
		domainEntities, err = s.invitationRepo.Get.ByTenantIDAndStatus(ctx, tenantID, *status)
	} else {
		domainEntities, err = s.invitationRepo.Get.ByTenantID(ctx, tenantID)
	}
	
	if err != nil {
		return nil, err
	}
	
	results := make([]invitationResult.InvitationRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = invitationResult.InvitationMapper(entity)
	}
	return results, nil
}
