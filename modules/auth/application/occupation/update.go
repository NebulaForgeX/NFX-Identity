package occupation

import (
	"context"
	occupationCommands "nfxid/modules/auth/application/occupation/commands"
)

func (s *Service) UpdateOccupation(ctx context.Context, cmd occupationCommands.UpdateOccupationCmd) error {
	o, err := s.occupationRepo.GetByID(ctx, cmd.OccupationID)
	if err != nil {
		return err
	}

	if err := o.EnsureEditable(cmd.Editable); err != nil {
		return err
	}

	if err := o.Update(cmd.Editable); err != nil {
		return err
	}

	if err := s.occupationRepo.Update(ctx, o); err != nil {
		return err
	}

	return nil
}
