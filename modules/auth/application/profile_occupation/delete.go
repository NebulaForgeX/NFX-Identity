package profile_occupation

import (
	"context"

	"github.com/google/uuid"
)

type DeleteOccupationCmd struct {
	OccupationID uuid.UUID
}

func (s *Service) DeleteOccupation(ctx context.Context, cmd DeleteOccupationCmd) error {
	o, err := s.occupationRepo.GetByID(ctx, cmd.OccupationID)
	if err != nil {
		return err
	}

	if err := o.Delete(); err != nil {
		return err
	}

	if err := s.occupationRepo.Update(ctx, o); err != nil {
		return err
	}

	return nil
}

