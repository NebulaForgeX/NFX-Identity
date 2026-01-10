package sessions

import (
	"context"
	sessionCommands "nfxid/modules/auth/application/sessions/commands"
)

// UpdateSessionLastSeen 更新会话最后访问时间
func (s *Service) UpdateSessionLastSeen(ctx context.Context, cmd sessionCommands.UpdateSessionLastSeenCmd) error {
	// Get domain entity
	session, err := s.sessionRepo.Get.BySessionID(ctx, cmd.SessionID)
	if err != nil {
		return err
	}

	// Update last seen domain entity
	if err := session.UpdateLastSeen(); err != nil {
		return err
	}

	// Save to repository
	return s.sessionRepo.Update.UpdateLastSeen(ctx, cmd.SessionID)
}

// RevokeSession 撤销会话
func (s *Service) RevokeSession(ctx context.Context, cmd sessionCommands.RevokeSessionCmd) error {
	// Get domain entity
	session, err := s.sessionRepo.Get.BySessionID(ctx, cmd.SessionID)
	if err != nil {
		return err
	}

	// Revoke domain entity
	if err := session.Revoke(cmd.RevokeReason, cmd.RevokedBy); err != nil {
		return err
	}

	// Save to repository
	return s.sessionRepo.Update.Revoke(ctx, cmd.SessionID, cmd.RevokeReason, cmd.RevokedBy)
}
