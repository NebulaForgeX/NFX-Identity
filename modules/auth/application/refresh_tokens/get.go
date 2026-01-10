package refresh_tokens

import (
	"context"
	refreshTokenResult "nfxid/modules/auth/application/refresh_tokens/results"

	"github.com/google/uuid"
)

// GetRefreshToken 根据ID获取刷新令牌
func (s *Service) GetRefreshToken(ctx context.Context, refreshTokenID uuid.UUID) (refreshTokenResult.RefreshTokenRO, error) {
	domainEntity, err := s.refreshTokenRepo.Get.ByID(ctx, refreshTokenID)
	if err != nil {
		return refreshTokenResult.RefreshTokenRO{}, err
	}
	return refreshTokenResult.RefreshTokenMapper(domainEntity), nil
}

// GetRefreshTokenByTokenID 根据TokenID获取刷新令牌
func (s *Service) GetRefreshTokenByTokenID(ctx context.Context, tokenID string) (refreshTokenResult.RefreshTokenRO, error) {
	domainEntity, err := s.refreshTokenRepo.Get.ByTokenID(ctx, tokenID)
	if err != nil {
		return refreshTokenResult.RefreshTokenRO{}, err
	}
	return refreshTokenResult.RefreshTokenMapper(domainEntity), nil
}
