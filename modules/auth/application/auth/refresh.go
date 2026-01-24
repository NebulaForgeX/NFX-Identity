package auth

import (
	"context"
	"time"

	authResults "nfxid/modules/auth/application/auth/results"
	refreshTokenDomain "nfxid/modules/auth/domain/refresh_tokens"

	"github.com/google/uuid"
)

// Refresh 使用 refresh_token 换取新的 access + refresh（带 token rotation）
func (s *Service) Refresh(ctx context.Context, refreshToken string, ip *string) (authResults.RefreshResult, error) {
	if refreshToken == "" {
		return authResults.RefreshResult{}, ErrInvalidRefreshToken
	}

	// 验证 refresh token JWT 并提取 claims（包含 token_id/jti）
	claims, err := s.tokenIssuer.VerifyRefreshToken(refreshToken)
	if err != nil {
		return authResults.RefreshResult{}, ErrInvalidRefreshToken
	}

	// 如果 JWT 中没有 token_id，说明是旧版本的 token，直接刷新（向后兼容）
	if claims.TokenID == "" {
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

	// 查询 refresh_tokens 表，检查 token 是否被撤销
	rt, err := s.refreshTokenRepo.Get.ByTokenID(ctx, claims.TokenID)
	if err != nil {
		// Token 不在表中，可能是旧版本或已被删除
		return authResults.RefreshResult{}, ErrInvalidRefreshToken
	}

	// 检查是否已被撤销
	if rt.RevokedAt() != nil {
		return authResults.RefreshResult{}, ErrInvalidRefreshToken
	}

	// 检查是否过期
	if rt.ExpiresAt().Before(time.Now()) {
		return authResults.RefreshResult{}, ErrInvalidRefreshToken
	}

	// 解析用户ID
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return authResults.RefreshResult{}, ErrInvalidRefreshToken
	}

	// 检查账户是否被锁定
	locked, err := s.accountLockoutRepo.Check.IsLocked(ctx, userID)
	if err == nil && locked {
		// 账户被锁定，撤销所有refresh token
		_ = s.refreshTokenRepo.Update.Revoke(ctx, claims.TokenID, refreshTokenDomain.RevokeReasonAccountLocked)
		return authResults.RefreshResult{}, ErrAccountLocked
	}

	// 生成新的 refresh token ID（用于 rotation）
	newRefreshTokenID, err := uuid.NewV7()
	if err != nil {
		return authResults.RefreshResult{}, err
	}
	newRefreshTokenIDStr := newRefreshTokenID.String()

	// 生成新的 Token 对
	access, newRefresh, err := s.tokenIssuer.IssuePairWithRefreshID(
		claims.UserID,
		claims.Username,
		claims.Email,
		claims.Phone,
		claims.RoleID,
		newRefreshTokenIDStr,
	)
	if err != nil {
		return authResults.RefreshResult{}, err
	}

	// 撤销旧的 refresh token（标记为 rotation）
	_ = s.refreshTokenRepo.Update.Revoke(ctx, claims.TokenID, refreshTokenDomain.RevokeReasonRotation)

	// 创建新的 refresh_tokens 记录
	expiresAt := time.Now().Add(time.Duration(s.refreshTokenTTL) * time.Second)
	oldTokenID := rt.ID()
	newRefreshTokenEntity, err := refreshTokenDomain.NewRefreshToken(refreshTokenDomain.NewRefreshTokenParams{
		TokenID:     newRefreshTokenIDStr,
		UserID:      userID,
		ExpiresAt:   expiresAt,
		RotatedFrom: &oldTokenID,
		IP:          ip,
	})
	if err == nil {
		_ = s.refreshTokenRepo.Create.New(ctx, newRefreshTokenEntity)
	}

	return authResults.RefreshResult{
		AccessToken:  access,
		RefreshToken: newRefresh,
		ExpiresIn:    s.expiresInSec,
	}, nil
}
