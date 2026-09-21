package login

import (
	"context"
	"strings"
	"time"

	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/pkgs/tokenx/hashx"
	"nfxidentity/pkgs/transaction"
)

func (s *Service) Logout(ctx context.Context, refreshToken string) error {
	refreshToken = strings.TrimSpace(refreshToken)
	if refreshToken == "" {
		return authErr.ErrInvalidRefreshToken
	}
	row, err := s.repoFactory.RefreshToken(transaction.UoW{}).Get.ByTokenHash(ctx, hashx.SHA256HexString(refreshToken))
	if err != nil {
		return authErr.ErrInvalidRefreshToken
	}
	if err := row.Revoke(time.Now().UTC()); err != nil {
		return err
	}
	return s.repoFactory.RefreshToken(transaction.UoW{}).Update.Generic(ctx, row)
}
