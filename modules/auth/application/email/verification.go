package email

import (
	"context"
	"strings"
	"time"

	"nfxidentity/errors/src/auth"
	emailinternal "nfxidentity/modules/auth/application/email/internal"
	authmail "nfxidentity/modules/auth/infrastructure/email"
)

func (s *Service) IssueAndStoreVerification(ctx context.Context, to, lang string) error {
	code := emailinternal.RandomCode()
	if err := authmail.SendVerificationEmail(ctx, s.mail, to, code, lang); err != nil {
		return err
	}
	if s.redis == nil {
		return auth.ErrVerificationCodeSaveFailed
	}
	return s.redis.Set(ctx, emailinternal.VerifyKey(to), code, 10*time.Minute).Err()
}

func (s *Service) ConsumeVerificationCode(ctx context.Context, emailAddr, code string) error {
	code = strings.TrimSpace(code)
	if code == "" {
		return auth.ErrVerificationCodeWrong
	}
	if s.checkFn != nil {
		if s.checkFn(ctx, emailAddr, code) {
			return nil
		}
		return auth.ErrVerificationCodeWrong
	}
	return auth.ErrVerificationCodeExpired
}

func (s *Service) redisCheckCode(ctx context.Context, emailAddr, code string) bool {
	if s.redis == nil {
		return false
	}
	got, err := s.redis.Get(ctx, emailinternal.VerifyKey(emailAddr)).Result()
	if err != nil {
		return false
	}
	ok := strings.EqualFold(strings.TrimSpace(got), strings.TrimSpace(code))
	if ok {
		_ = s.redis.Del(ctx, emailinternal.VerifyKey(emailAddr)).Err()
	}
	return ok
}
