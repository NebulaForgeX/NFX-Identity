package profile

import (
	"context"
	profileCommands "nebulaid/modules/auth/application/profile/commands"
)

func (s *Service) UpdateProfile(ctx context.Context, cmd profileCommands.UpdateProfileCmd) error {
	p, err := s.profileRepo.GetByID(ctx, cmd.ProfileID)
	if err != nil {
		return err
	}

	if err := p.EnsureEditable(cmd.Editable); err != nil {
		return err
	}

	if err := p.Update(cmd.Editable); err != nil {
		return err
	}

	if err := s.profileRepo.Update(ctx, p); err != nil {
		return err
	}

	return nil
}
