package invitations

import (
	invitationDomain "nfxid/modules/tenants/domain/invitations"
)

type Service struct {
	invitationRepo *invitationDomain.Repo
}

func NewService(
	invitationRepo *invitationDomain.Repo,
) *Service {
	return &Service{
		invitationRepo: invitationRepo,
	}
}
