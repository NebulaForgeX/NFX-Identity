package profile

import (
	"context"

	"github.com/google/uuid"
)

type DeleteProfileCmd struct {
	ProfileID uuid.UUID
}

func (s *Service) DeleteProfile(ctx context.Context, cmd DeleteProfileCmd) error {
	p, err := s.profileRepo.GetByID(ctx, cmd.ProfileID)
	if err != nil {
		return err
	}

	if err := p.Delete(); err != nil {
		return err
	}

	if err := s.profileRepo.Update(ctx, p); err != nil {
		return err
	}

	return nil
}

