package sessions

import (
	"context"
	"time"
	sessionCommands "nfxid/modules/auth/application/sessions/commands"
	sessionDomain "nfxid/modules/auth/domain/sessions"

	"github.com/google/uuid"
)

// CreateSession 创建会话
func (s *Service) CreateSession(ctx context.Context, cmd sessionCommands.CreateSessionCmd) (uuid.UUID, error) {
	// Parse expires at
	expiresAt, err := time.Parse(time.RFC3339, cmd.ExpiresAt)
	if err != nil {
		return uuid.Nil, err
	}

	// Create domain entity
	session, err := sessionDomain.NewSession(sessionDomain.NewSessionParams{
		SessionID:         cmd.SessionID,
		TenantID:          cmd.TenantID,
		UserID:            cmd.UserID,
		AppID:             cmd.AppID,
		ClientID:          cmd.ClientID,
		ExpiresAt:         expiresAt,
		IP:                cmd.IP,
		UAHash:            cmd.UAHash,
		DeviceID:          cmd.DeviceID,
		DeviceFingerprint: cmd.DeviceFingerprint,
		DeviceName:        cmd.DeviceName,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.sessionRepo.Create.New(ctx, session); err != nil {
		return uuid.Nil, err
	}

	return session.ID(), nil
}
