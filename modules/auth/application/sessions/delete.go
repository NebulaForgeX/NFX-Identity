package sessions

import (
	"context"
	sessionCommands "nfxid/modules/auth/application/sessions/commands"
)

// DeleteSession 删除会话
func (s *Service) DeleteSession(ctx context.Context, cmd sessionCommands.DeleteSessionCmd) error {
	// Delete from repository (hard delete)
	return s.sessionRepo.Delete.ByID(ctx, cmd.SessionID)
}
