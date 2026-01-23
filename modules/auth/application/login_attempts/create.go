package login_attempts

import (
	"context"
	loginAttemptCommands "nfxid/modules/auth/application/login_attempts/commands"
	loginAttemptDomain "nfxid/modules/auth/domain/login_attempts"

	"github.com/google/uuid"
)

// CreateLoginAttempt 创建登录尝试
func (s *Service) CreateLoginAttempt(ctx context.Context, cmd loginAttemptCommands.CreateLoginAttemptCmd) (uuid.UUID, error) {
	// Create domain entity
	loginAttempt, err := loginAttemptDomain.NewLoginAttempt(loginAttemptDomain.NewLoginAttemptParams{
		Identifier:        cmd.Identifier,
		UserID:            cmd.UserID,
		IP:                cmd.IP,
		UAHash:            cmd.UAHash,
		DeviceFingerprint: cmd.DeviceFingerprint,
		Success:           cmd.Success,
		FailureCode:       cmd.FailureCode,
		MFARequired:       cmd.MFARequired,
		MFAVerified:       cmd.MFAVerified,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.loginAttemptRepo.Create.New(ctx, loginAttempt); err != nil {
		return uuid.Nil, err
	}

	return loginAttempt.ID(), nil
}
