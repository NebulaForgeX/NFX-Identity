package platform

import (
	"context"
	"strings"
	"time"

	"nfxidentity/modules/auth/domain/email"
	"nfxidentity/pkgs/errx"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
)

func (s *Service) ListEmails(ctx context.Context, accountID uuid.UUID) ([]map[string]any, int, error) {
	rows, err := s.emails.List.ByAccountID(ctx, accountID)
	if err != nil {
		return nil, 0, err
	}
	return emailMaps(rows), len(rows), nil
}

func (s *Service) CreateEmail(ctx context.Context, accountID uuid.UUID, address string) (string, error) {
	address = strings.ToLower(strings.TrimSpace(address))
	if address == "" {
		return "", errx.InvalidArg("INVALID_EMAIL", "email required")
	}
	now := time.Now()
	id := uuid.New()
	if err := s.repos.Email(none()).Create.New(ctx, email.NewFromState(email.EmailState{
		ID: id, AccountID: accountID, Address: address, CreatedAt: now, UpdatedAt: now,
	})); err != nil {
		return "", errx.Conflict("EMAIL_TAKEN", "email already registered")
	}
	return id.String(), nil
}

func (s *Service) SendEmailVerificationCode(ctx context.Context, accountID, emailID uuid.UUID) error {
	row, err := s.repos.Email(none()).Get.ByID(ctx, emailID)
	if err != nil || row.AccountID() != accountID {
		return ErrNotFound
	}
	s.StoreVerificationCode(ctx, row.Address(), RandomCode())
	return nil
}

func (s *Service) VerifyEmail(ctx context.Context, accountID, emailID uuid.UUID, code string) error {
	row, err := s.repos.Email(none()).Get.ByID(ctx, emailID)
	if err != nil || row.AccountID() != accountID {
		return ErrNotFound
	}
	if !s.checkVerificationCode(ctx, row.Address(), code) {
		return errx.InvalidArg("INVALID_VERIFICATION_CODE", "invalid verification code")
	}
	row.MarkVerified(time.Now())
	return s.repos.Email(none()).Update.Generic(ctx, row)
}

func (s *Service) UpdateEmail(ctx context.Context, accountID, emailID uuid.UUID, address string) error {
	address = strings.ToLower(strings.TrimSpace(address))
	row, err := s.repos.Email(none()).Get.ByID(ctx, emailID)
	if err != nil || row.AccountID() != accountID {
		return ErrNotFound
	}
	row.ChangeAddress(address)
	return s.repos.Email(none()).Update.Generic(ctx, row)
}

func (s *Service) SetPrimaryEmail(ctx context.Context, accountID, emailID uuid.UUID) error {
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		repo := s.repos.Email(uow)
		rows, err := repo.Get.ByAccountID(ctx, accountID)
		if err != nil {
			return err
		}
		found := false
		for _, row := range rows {
			if row.ID() == emailID {
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

func (s *Service) DeleteEmail(ctx context.Context, accountID, emailID uuid.UUID) error {
	row, err := s.repos.Email(none()).Get.ByID(ctx, emailID)
	if err != nil || row.AccountID() != accountID {
		return errx.FailedPrecond("EMAIL_NOT_DELETABLE", "cannot delete primary or missing email")
	}
	if row.IsPrimary() {
		return errx.FailedPrecond("EMAIL_NOT_DELETABLE", "cannot delete primary or missing email")
	}
	row.SoftDelete(time.Now())
	return s.repos.Email(none()).Update.Generic(ctx, row)
}

func (s *Service) AccountIDByEmail(ctx context.Context, address string) (uuid.UUID, error) {
	row, err := s.repos.Email(none()).Get.ByAddress(ctx, strings.ToLower(strings.TrimSpace(address)))
	if err != nil {
		return uuid.Nil, ErrNotFound
	}
	return row.AccountID(), nil
}
