package members

import (
	"context"
	memberCommands "nfxid/modules/tenants/application/members/commands"
)

// UpdateMemberStatus 更新成员状态
func (s *Service) UpdateMemberStatus(ctx context.Context, cmd memberCommands.UpdateMemberStatusCmd) error {
	// Get domain entity
	member, err := s.memberRepo.Get.ByMemberID(ctx, cmd.MemberID)
	if err != nil {
		return err
	}

	// Update status domain entity
	if err := member.UpdateStatus(cmd.Status); err != nil {
		return err
	}

	// Save to repository
	return s.memberRepo.Update.Status(ctx, cmd.MemberID, cmd.Status)
}

// JoinMember 成员加入
func (s *Service) JoinMember(ctx context.Context, cmd memberCommands.JoinMemberCmd) error {
	// Get domain entity
	member, err := s.memberRepo.Get.ByMemberID(ctx, cmd.MemberID)
	if err != nil {
		return err
	}

	// Join domain entity
	if err := member.Join(); err != nil {
		return err
	}

	// Save to repository
	return s.memberRepo.Update.Join(ctx, cmd.MemberID)
}

// LeaveMember 成员离开
func (s *Service) LeaveMember(ctx context.Context, cmd memberCommands.LeaveMemberCmd) error {
	// Get domain entity
	member, err := s.memberRepo.Get.ByMemberID(ctx, cmd.MemberID)
	if err != nil {
		return err
	}

	// Leave domain entity
	if err := member.Leave(); err != nil {
		return err
	}

	// Save to repository
	return s.memberRepo.Update.Leave(ctx, cmd.MemberID)
}

// SuspendMember 暂停成员
func (s *Service) SuspendMember(ctx context.Context, cmd memberCommands.SuspendMemberCmd) error {
	// Get domain entity
	member, err := s.memberRepo.Get.ByMemberID(ctx, cmd.MemberID)
	if err != nil {
		return err
	}

	// Suspend domain entity
	if err := member.Suspend(); err != nil {
		return err
	}

	// Save to repository
	return s.memberRepo.Update.Generic(ctx, member)
}
