package profile_occupation

import (
	"context"
	occupationCommands "nfxid/modules/auth/application/profile_occupation/commands"
)

func (s *Service) UpdateOccupation(ctx context.Context, cmd occupationCommands.UpdateOccupationCmd) error {
	o, err := s.occupationRepo.Get.ByID(ctx, cmd.OccupationID)
	if err != nil {
		return err
	}

	if err := o.EnsureEditable(cmd.Editable); err != nil {
		return err
	}

	if err := o.Update(cmd.Editable); err != nil {
		return err
	}

	if err := s.occupationRepo.Update.Generic(ctx, o); err != nil {
		return err
	}

	return nil
}
