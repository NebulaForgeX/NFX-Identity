package refresh_tokens

import (
	"context"
	"time"
	refreshTokenCommands "nfxid/modules/auth/application/refresh_tokens/commands"
	refreshTokenDomain "nfxid/modules/auth/domain/refresh_tokens"

	"github.com/google/uuid"
)

// CreateRefreshToken 创建刷新令牌
func (s *Service) CreateRefreshToken(ctx context.Context, cmd refreshTokenCommands.CreateRefreshTokenCmd) (uuid.UUID, error) {
	// Parse expires at
	expiresAt, err := time.Parse(time.RFC3339, cmd.ExpiresAt)
	if err != nil {
		return uuid.Nil, err
	}

	// Create domain entity
	refreshToken, err := refreshTokenDomain.NewRefreshToken(refreshTokenDomain.NewRefreshTokenParams{
		TokenID:   cmd.TokenID,
		UserID:    cmd.UserID,
		AppID:     cmd.AppID,
		ClientID:  cmd.ClientID,
		SessionID: cmd.SessionID,
		ExpiresAt: expiresAt,
		DeviceID:  cmd.DeviceID,
		IP:        cmd.IP,
		UAHash:    cmd.UAHash,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.refreshTokenRepo.Create.New(ctx, refreshToken); err != nil {
		return uuid.Nil, err
	}

	return refreshToken.ID(), nil
}
