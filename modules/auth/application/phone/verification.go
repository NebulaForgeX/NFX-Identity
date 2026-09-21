package phone

import (
	"context"
	"strings"

	"nfxidentity/errors/src/auth"
	phoneinternal "nfxidentity/modules/auth/application/phone/internal"
)

func (s *Service) ConsumeVerificationCode(ctx context.Context, key, code string) error {
	code = strings.TrimSpace(code)
	if code == "" {
		return auth.ErrVerificationCodeWrong
	}
	if s.checkFn != nil {
		if s.checkFn(ctx, key, code) {
			return nil
		}
		return auth.ErrVerificationCodeWrong
	}
	return auth.ErrVerificationCodeExpired
}

func (s *Service) redisCheckCode(ctx context.Context, key, code string) bool {
	if s.redis == nil {
		return false
	}
	got, err := s.redis.Get(ctx, phoneinternal.VerifyKey(key)).Result()
	if err != nil {
		return false
	}
	ok := strings.EqualFold(strings.TrimSpace(got), strings.TrimSpace(code))
	if ok {
		_ = s.redis.Del(ctx, phoneinternal.VerifyKey(key)).Err()
	}
	return ok
}
