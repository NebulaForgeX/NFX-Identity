package signup

import (
	"context"
	"strings"
	"time"

	"nfxidentity/errors/src/auth"
	signupinternal "nfxidentity/modules/auth/application/signup/internal"
	authmail "nfxidentity/modules/auth/infrastructure/email"
)

func (s *Service) storeVerificationCode(ctx context.Context, email, code string) error {
	if s.redis == nil {
		return auth.ErrVerificationCodeSaveFailed
	}
	if err := s.redis.Set(ctx, signupinternal.VerifyKey(email), code, 10*time.Minute).Err(); err != nil {
		return auth.ErrVerificationCodeSaveFailed.WithCause(err)
	}
	return nil
}

func (s *Service) lookupVerificationCode(ctx context.Context, email string) (string, error) {
	if s.redis == nil {
		return "", auth.ErrVerificationCodeExpired
	}
	got, err := s.redis.Get(ctx, signupinternal.VerifyKey(email)).Result()
	if err != nil {
		return "", auth.ErrVerificationCodeExpired
	}
	return got, nil
}

func (s *Service) deleteVerificationCode(ctx context.Context, email string) error {
	if s.redis == nil {
		return nil
	}
	return s.redis.Del(ctx, signupinternal.VerifyKey(email)).Err()
}

func (s *Service) redisCheckCode(ctx context.Context, email, code string) bool {
	got, err := s.lookupVerificationCode(ctx, email)
	if err != nil {
		return false
	}
	ok := strings.EqualFold(strings.TrimSpace(got), strings.TrimSpace(code))
	if ok {
		_ = s.deleteVerificationCode(ctx, email)
	}
	return ok
}

func (s *Service) consumeVerificationCode(ctx context.Context, email, code string) error {
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

func (s *Service) issueAndStoreVerification(ctx context.Context, to, lang string) error {
	code := signupinternal.RandomCode()
	if err := authmail.SendVerificationEmail(ctx, s.mail, to, code, lang); err != nil {
		return err
	}
	return s.storeVerificationCode(ctx, to, code)
}
