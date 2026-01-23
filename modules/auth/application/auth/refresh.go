package auth

import (
	"context"

	authResults "nfxid/modules/auth/application/auth/results"
)

// Refresh 使用 refresh_token 换取新的 access + refresh
func (s *Service) Refresh(ctx context.Context, refreshToken string) (authResults.RefreshResult, error) {
	if refreshToken == "" {
		return authResults.RefreshResult{}, ErrInvalidRefreshToken
	}

	access, refresh, err := s.tokenIssuer.RefreshPair(refreshToken)
	if err != nil {
		return authResults.RefreshResult{}, ErrInvalidRefreshToken
	}

	return authResults.RefreshResult{
		AccessToken:  access,
		RefreshToken: refresh,
		ExpiresIn:    s.expiresInSec,
	}, nil
}
