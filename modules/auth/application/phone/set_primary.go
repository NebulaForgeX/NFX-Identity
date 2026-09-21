package phone

import (
	"context"

	"nfxidentity/errors/src/sys"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
)

type SetPrimaryInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	PhoneID   uuid.UUID
}

func (s *Service) SetPrimary(ctx context.Context, in SetPrimaryInput) error {
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		phoneRepo := s.repoFactory.Phone(uow)
		rows, err := phoneRepo.Get.ByAccountID(ctx, in.AccountID)
		if err != nil {
			return err
		}
		found := false
		for _, row := range rows {
			if row.ID() == in.PhoneID {
				row.SetPrimary(true)
				found = true
			} else {
				row.SetPrimary(false)
			}
			if err := phoneRepo.Update.Generic(ctx, row); err != nil {
				return err
			}
		}
		if !found {
			return sys.ErrNotFound
		}
		return nil
	})
}
