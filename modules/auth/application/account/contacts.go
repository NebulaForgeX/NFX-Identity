package account

import (
	"context"
	"strings"
	"time"

	"nfxidentity/enums"
	"nfxidentity/errors/src/auth"
	"nfxidentity/errors/src/sys"
	accountinternal "nfxidentity/modules/auth/application/account/internal"
	authmail "nfxidentity/modules/auth/infrastructure/email"

	"github.com/google/uuid"
)

func (s *Service) IssueAndStoreVerification(ctx context.Context, to, lang string) error {
	code := accountinternal.RandomCode()
	if err := authmail.SendVerificationEmail(ctx, s.mail, to, code, lang); err != nil {
		return err
	}
	if s.redis == nil {
		return auth.ErrVerificationCodeSaveFailed
	}
	return s.redis.Set(ctx, accountinternal.VerifyKey(to), code, 10*time.Minute).Err()
}

func (s *Service) ConsumeVerificationCode(ctx context.Context, email, code string) error {
	code = strings.TrimSpace(code)
	if code == "" {
		return auth.ErrVerificationCodeWrong
	}
	if s.checkFn != nil {
		if s.checkFn(ctx, email, code) {
			return nil
		}
		return auth.ErrVerificationCodeWrong
	}
	return auth.ErrVerificationCodeExpired
}

func (s *Service) redisCheckCode(ctx context.Context, email, code string) bool {
	if s.redis == nil {
		return false
	}
	got, err := s.redis.Get(ctx, accountinternal.VerifyKey(email)).Result()
	if err != nil {
		return false
	}
	ok := strings.EqualFold(strings.TrimSpace(got), strings.TrimSpace(code))
	if ok {
		_ = s.redis.Del(ctx, accountinternal.VerifyKey(email)).Err()
	}
	return ok
}

func (s *Service) PrimaryContacts(ctx context.Context, accountID uuid.UUID) (string, string) {
	var emailAddr, phoneNum string
	if emails, err := s.emails.List.ByAccountID(ctx, accountID); err == nil {
		for _, item := range emails {
			if item.IsPrimary {
				emailAddr = item.Email
				break
			}
		}
		if emailAddr == "" && len(emails) > 0 {
			emailAddr = emails[0].Email
		}
	}
	if phones, err := s.phones.List.ByAccountID(ctx, accountID); err == nil {
		for _, item := range phones {
			if item.IsPrimary {
				phoneNum = item.Phone
				break
			}
		}
		if phoneNum == "" && len(phones) > 0 {
			phoneNum = phones[0].Phone
		}
	}
	return emailAddr, phoneNum
}

func (s *Service) PrimaryContactsExport(ctx context.Context, accountID uuid.UUID) (string, error) {
	emailAddr, _ := s.PrimaryContacts(ctx, accountID)
	if emailAddr == "" {
		return "", sys.ErrNotFound
	}
	return emailAddr, nil
}

func (s *Service) InvalidateFull(ctx context.Context, accountID, profileID uuid.UUID) error {
	return nil
}

func (s *Service) RequireOwner(ctx context.Context, accountID uuid.UUID) error {
	rows, err := s.profiles.AuthorityList.ByAccountID(ctx, accountID)
	if err != nil {
		return err
	}
	for _, row := range rows {
		for _, role := range row.AuthorityRoles {
			if role == enums.AuthAuthorityRoleOwner {
				return nil
			}
		}
	}
	return auth.ErrInvalidAccountPermission
}

func (s *Service) AccountIDByEmail(ctx context.Context, address string) (uuid.UUID, error) {
	row, err := s.repoFactory.Email(accountinternal.None()).Get.ByEmail(ctx, strings.ToLower(strings.TrimSpace(address)))
	if err != nil {
		return uuid.Nil, sys.ErrNotFound
	}
	return row.AccountID(), nil
}

func (s *Service) AccountIDByPhone(ctx context.Context, number string) (uuid.UUID, error) {
	row, err := s.repoFactory.Phone(accountinternal.None()).Get.ByPhone(ctx, strings.TrimSpace(number))
	if err != nil {
		return uuid.Nil, sys.ErrNotFound
	}
	return row.AccountID(), nil
}
