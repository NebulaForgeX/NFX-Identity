package platform

import (
	"context"
	"nfxidentity/errors/src/auth"
	"nfxidentity/errors/src/sys"
	"strings"
	"time"

	"nfxidentity/modules/auth/domain/email"
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
		return "", auth.ErrInvalidEmail
	}
	now := time.Now()
	id := uuid.New()
	if err := s.repoFactory.Email(none()).Create.New(ctx, email.NewFromState(email.EmailState{
		ID: id, AccountID: accountID, Address: address, CreatedAt: now, UpdatedAt: now,
	})); err != nil {
		return "", auth.ErrEmailAlreadyExists
	}
	return id.String(), nil
}

func (s *Service) SendEmailVerificationCode(ctx context.Context, accountID, emailID uuid.UUID, lang string) error {
	row, err := s.repoFactory.Email(none()).Get.ByID(ctx, emailID)
	if err != nil || row.AccountID() != accountID {
		return sys.ErrNotFound
	}
	if err := s.issueAndStoreVerification(ctx, row.Address(), lang); err != nil {
		return err
	}
	return nil
}

func (s *Service) VerifyEmail(ctx context.Context, accountID, emailID uuid.UUID, code string) error {
	row, err := s.repoFactory.Email(none()).Get.ByID(ctx, emailID)
	if err != nil || row.AccountID() != accountID {
		return sys.ErrNotFound
	}
	if err := s.consumeVerificationCode(ctx, row.Address(), code); err != nil {
		return err
	}
	row.MarkVerified(time.Now())
	return s.repoFactory.Email(none()).Update.Generic(ctx, row)
}

func (s *Service) UpdateEmail(ctx context.Context, accountID, emailID uuid.UUID, address string) error {
	address = strings.ToLower(strings.TrimSpace(address))
	row, err := s.repoFactory.Email(none()).Get.ByID(ctx, emailID)
	if err != nil || row.AccountID() != accountID {
		return sys.ErrNotFound
	}
	row.ChangeAddress(address)
	return s.repoFactory.Email(none()).Update.Generic(ctx, row)
}

func (s *Service) SetPrimaryEmail(ctx context.Context, accountID, emailID uuid.UUID) error {
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		emailRepo := s.repoFactory.Email(uow)
		rows, err := emailRepo.Get.ByAccountID(ctx, accountID)
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
			if err := emailRepo.Update.Generic(ctx, row); err != nil {
				return err
			}
		}
		if !found {
			return sys.ErrNotFound
		}
		return nil
	})
}

func (s *Service) DeleteEmail(ctx context.Context, accountID, emailID uuid.UUID) error {
	row, err := s.repoFactory.Email(none()).Get.ByID(ctx, emailID)
	if err != nil || row.AccountID() != accountID {
		return auth.ErrEmailCannotDeletePrimary
	}
	if row.IsPrimary() {
		return auth.ErrEmailCannotDeletePrimary
	}
	row.SoftDelete(time.Now())
	return s.repoFactory.Email(none()).Update.Generic(ctx, row)
}

func (s *Service) AccountIDByEmail(ctx context.Context, address string) (uuid.UUID, error) {
	row, err := s.repoFactory.Email(none()).Get.ByAddress(ctx, strings.ToLower(strings.TrimSpace(address)))
	if err != nil {
		return uuid.Nil, sys.ErrNotFound
	}
	return row.AccountID(), nil
}
