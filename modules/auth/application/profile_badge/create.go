package profile_badge

import (
	"context"
	profileBadgeCommands "nebulaid/modules/auth/application/profile_badge/commands"
	profileBadgeDomain "nebulaid/modules/auth/domain/profile_badge"
	profileBadgeDomainErrors "nebulaid/modules/auth/domain/profile_badge/errors"
)

func (s *Service) CreateProfileBadge(ctx context.Context, cmd profileBadgeCommands.CreateProfileBadgeCmd) (*profileBadgeDomain.ProfileBadge, error) {
	// 检查关联是否已存在
	if exists, _ := s.profileBadgeRepo.Exists(ctx, cmd.ProfileID, cmd.BadgeID); exists {
		return nil, profileBadgeDomainErrors.ErrProfileBadgeAlreadyExists
	}

	// 使用 domain factory 创建实体
	pb, err := profileBadgeDomain.NewProfileBadge(profileBadgeDomain.NewProfileBadgeParams{
		ProfileID: cmd.ProfileID,
		BadgeID:   cmd.BadgeID,
		Editable:  cmd.Editable,
	})
	if err != nil {
		return nil, err
	}

	if err := s.profileBadgeRepo.Create(ctx, pb); err != nil {
		return nil, err
	}

	return pb, nil
}
