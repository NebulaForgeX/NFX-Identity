package member_groups

import (
	"context"
	memberGroupCommands "nfxidentity/modules/tenants/application/member_groups/commands"
	memberGroupDomain "nfxidentity/modules/tenants/domain/member_groups"
	tenantsErr "nfxidentity/errors/src/tenants"

	"github.com/google/uuid"
)

// CreateMemberGroup 创建成员组
func (s *Service) CreateMemberGroup(ctx context.Context, cmd memberGroupCommands.CreateMemberGroupCmd) (uuid.UUID, error) {
	// Check if member group already exists
	if exists, _ := s.memberGroupRepo.Check.ByMemberIDAndGroupID(ctx, cmd.MemberID, cmd.GroupID); exists {
		return uuid.Nil, tenantsErr.ErrMemberGroupAlreadyExists
	}

	// Create domain entity
	memberGroup, err := memberGroupDomain.NewMemberGroup(memberGroupDomain.NewMemberGroupParams{
		MemberID:   cmd.MemberID,
		GroupID:    cmd.GroupID,
		AssignedBy: cmd.AssignedBy,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.memberGroupRepo.Create.New(ctx, memberGroup); err != nil {
		return uuid.Nil, err
	}

	return memberGroup.ID(), nil
}
