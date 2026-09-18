package platform

import (
	"context"
	"nfxidentity/errors/src/auth"
	"nfxidentity/errors/src/sys"
	"strings"
	"time"

	"nfxidentity/modules/auth/domain/phone"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
)

func (s *Service) ListPhones(ctx context.Context, accountID uuid.UUID) ([]map[string]any, int, error) {
	rows, err := s.phones.List.ByAccountID(ctx, accountID)
	if err != nil {
		return nil, 0, err
	}
	return phoneMaps(rows), len(rows), nil
}

func (s *Service) CreatePhone(ctx context.Context, accountID uuid.UUID, number string) (string, error) {
	number = strings.TrimSpace(number)
	if number == "" {
		return "", auth.ErrInvalidPhone
	}
	now := time.Now()
	id := uuid.New()
	if err := s.repoFactory.Phone(none()).Create.New(ctx, phone.NewPhoneFromState(phone.PhoneState{
		ID: id, AccountID: accountID, Phone: number, CreatedAt: now, UpdatedAt: now,
	})); err != nil {
		return "", auth.ErrPhoneAlreadyExists
	}
	return id.String(), nil
}

func (s *Service) SendPhoneVerificationCode(ctx context.Context, accountID, phoneID uuid.UUID) error {
	row, err := s.repoFactory.Phone(none()).Get.ByID(ctx, phoneID)
	if err != nil || row.AccountID() != accountID {
		return sys.ErrNotFound
	}
	return sys.ErrUnimplemented
}

func (s *Service) VerifyPhone(ctx context.Context, accountID, phoneID uuid.UUID, code string) error {
	row, err := s.repoFactory.Phone(none()).Get.ByID(ctx, phoneID)
	if err != nil || row.AccountID() != accountID {
		return sys.ErrNotFound
	}
	if err := s.consumeVerificationCode(ctx, "phone:"+row.Phone(), code); err != nil {
		return err
	}
	row.MarkVerified(time.Now())
	return s.repoFactory.Phone(none()).Update.Generic(ctx, row)
}

func (s *Service) UpdatePhone(ctx context.Context, accountID, phoneID uuid.UUID, number string) error {
	number = strings.TrimSpace(number)
	row, err := s.repoFactory.Phone(none()).Get.ByID(ctx, phoneID)
	if err != nil || row.AccountID() != accountID {
		return sys.ErrNotFound
	}
	row.ChangeNumber(number)
	return s.repoFactory.Phone(none()).Update.Generic(ctx, row)
}

func (s *Service) SetPrimaryPhone(ctx context.Context, accountID, phoneID uuid.UUID) error {
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		phoneRepo := s.repoFactory.Phone(uow)
		rows, err := phoneRepo.Get.ByAccountID(ctx, accountID)
		if err != nil {
			return err
		}
		found := false
		for _, row := range rows {
			if row.ID() == phoneID {
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

func (s *Service) DeletePhone(ctx context.Context, accountID, phoneID uuid.UUID) error {
	row, err := s.repoFactory.Phone(none()).Get.ByID(ctx, phoneID)
	if err != nil || row.AccountID() != accountID {
		return auth.ErrPhoneNotDeletable
	}
	if row.IsPrimary() {
		return auth.ErrPhoneNotDeletable
	}
	row.SoftDelete(time.Now())
	return s.repoFactory.Phone(none()).Update.Generic(ctx, row)
}

func (s *Service) AccountIDByPhone(ctx context.Context, number string) (uuid.UUID, error) {
	row, err := s.repoFactory.Phone(none()).Get.ByPhone(ctx, strings.TrimSpace(number))
	if err != nil {
		return uuid.Nil, sys.ErrNotFound
	}
	return row.AccountID(), nil
}
