package member_groups

import (
	"context"
	memberGroupCommands "nfxid/modules/tenants/application/member_groups/commands"
	memberGroupDomain "nfxid/modules/tenants/domain/member_groups"

	"github.com/google/uuid"
)

// CreateMemberGroup 创建成员组
func (s *Service) CreateMemberGroup(ctx context.Context, cmd memberGroupCommands.CreateMemberGroupCmd) (uuid.UUID, error) {
	// Check if member group already exists
	if exists, _ := s.memberGroupRepo.Check.ByMemberIDAndGroupID(ctx, cmd.MemberID, cmd.GroupID); exists {
		return uuid.Nil, memberGroupDomain.ErrMemberGroupAlreadyExists
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
