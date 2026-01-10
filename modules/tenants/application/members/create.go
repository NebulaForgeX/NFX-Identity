package members

import (
	"context"
	memberCommands "nfxid/modules/tenants/application/members/commands"
	memberDomain "nfxid/modules/tenants/domain/members"

	"github.com/google/uuid"
)

// CreateMember 创建成员
func (s *Service) CreateMember(ctx context.Context, cmd memberCommands.CreateMemberCmd) (uuid.UUID, error) {
	// Check if member already exists
	if exists, _ := s.memberRepo.Check.ByTenantIDAndUserID(ctx, cmd.TenantID, cmd.UserID); exists {
		return uuid.Nil, memberDomain.ErrMemberAlreadyExists
	}

	// Create domain entity
	member, err := memberDomain.NewMember(memberDomain.NewMemberParams{
		TenantID:    cmd.TenantID,
		UserID:      cmd.UserID,
		Status:      cmd.Status,
		Source:      cmd.Source,
		CreatedBy:   cmd.CreatedBy,
		ExternalRef: cmd.ExternalRef,
		Metadata:    cmd.Metadata,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.memberRepo.Create.New(ctx, member); err != nil {
		return uuid.Nil, err
	}

	return member.ID(), nil
}
