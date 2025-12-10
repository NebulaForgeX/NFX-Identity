package profile_badge

import (
	"context"
	profileBadgeCommands "nebulaid/modules/auth/application/profile_badge/commands"
)

func (s *Service) UpdateProfileBadge(ctx context.Context, cmd profileBadgeCommands.UpdateProfileBadgeCmd) error {
	pb, err := s.profileBadgeRepo.GetByID(ctx, cmd.ProfileBadgeID)
	if err != nil {
		return err
	}

	if err := pb.EnsureEditable(cmd.Editable); err != nil {
		return err
	}

	if err := pb.Update(cmd.Editable); err != nil {
		return err
	}

	if err := s.profileBadgeRepo.Update(ctx, pb); err != nil {
		return err
	}

	return nil
}
