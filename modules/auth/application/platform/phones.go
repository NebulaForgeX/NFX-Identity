package platform

import (
	"context"
	"strings"
	"time"

	"nfxidentity/modules/auth/domain/phone"
	"nfxidentity/pkgs/errx"
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
		return "", errx.InvalidArg("INVALID_PHONE", "phone required")
	}
	now := time.Now()
	id := uuid.New()
	if err := s.repos.Phone(none()).Create.New(ctx, phone.NewFromState(phone.PhoneState{
		ID: id, AccountID: accountID, Number: number, CreatedAt: now, UpdatedAt: now,
	})); err != nil {
		return "", errx.Conflict("PHONE_TAKEN", "phone already registered")
	}
	return id.String(), nil
}

func (s *Service) SendPhoneVerificationCode(ctx context.Context, accountID, phoneID uuid.UUID) error {
	row, err := s.repos.Phone(none()).Get.ByID(ctx, phoneID)
	if err != nil || row.AccountID() != accountID {
		return ErrNotFound
	}
	s.StoreVerificationCode(ctx, "phone:"+row.Number(), RandomCode())
	return nil
}

func (s *Service) VerifyPhone(ctx context.Context, accountID, phoneID uuid.UUID, code string) error {
	row, err := s.repos.Phone(none()).Get.ByID(ctx, phoneID)
	if err != nil || row.AccountID() != accountID {
		return ErrNotFound
	}
	if !s.checkVerificationCode(ctx, "phone:"+row.Number(), code) {
		return errx.InvalidArg("INVALID_VERIFICATION_CODE", "invalid verification code")
	}
	row.MarkVerified(time.Now())
	return s.repos.Phone(none()).Update.Generic(ctx, row)
}

func (s *Service) UpdatePhone(ctx context.Context, accountID, phoneID uuid.UUID, number string) error {
	number = strings.TrimSpace(number)
	row, err := s.repos.Phone(none()).Get.ByID(ctx, phoneID)
	if err != nil || row.AccountID() != accountID {
		return ErrNotFound
	}
	row.ChangeNumber(number)
	return s.repos.Phone(none()).Update.Generic(ctx, row)
}

func (s *Service) SetPrimaryPhone(ctx context.Context, accountID, phoneID uuid.UUID) error {
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		repo := s.repos.Phone(uow)
		rows, err := repo.Get.ByAccountID(ctx, accountID)
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
			if err := repo.Update.Generic(ctx, row); err != nil {
				return err
			}
		}
		if !found {
			return ErrNotFound
		}
		return nil
	})
}

func (s *Service) DeletePhone(ctx context.Context, accountID, phoneID uuid.UUID) error {
	row, err := s.repos.Phone(none()).Get.ByID(ctx, phoneID)
	if err != nil || row.AccountID() != accountID {
		return errx.FailedPrecond("PHONE_NOT_DELETABLE", "cannot delete primary or missing phone")
	}
	if row.IsPrimary() {
		return errx.FailedPrecond("PHONE_NOT_DELETABLE", "cannot delete primary or missing phone")
	}
	row.SoftDelete(time.Now())
	return s.repos.Phone(none()).Update.Generic(ctx, row)
}

func (s *Service) AccountIDByPhone(ctx context.Context, number string) (uuid.UUID, error) {
	row, err := s.repos.Phone(none()).Get.ByNumber(ctx, strings.TrimSpace(number))
	if err != nil {
		return uuid.Nil, ErrNotFound
	}
	return row.AccountID(), nil
}
