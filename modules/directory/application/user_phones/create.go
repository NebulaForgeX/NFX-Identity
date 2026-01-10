package user_phones

import (
	"context"
	"time"
	userPhoneCommands "nfxid/modules/directory/application/user_phones/commands"
	userPhoneDomain "nfxid/modules/directory/domain/user_phones"

	"github.com/google/uuid"
)

// CreateUserPhone 创建用户手机号
func (s *Service) CreateUserPhone(ctx context.Context, cmd userPhoneCommands.CreateUserPhoneCmd) (uuid.UUID, error) {
	// Check if phone already exists
	if exists, _ := s.userPhoneRepo.Check.ByPhone(ctx, cmd.Phone); exists {
		return uuid.Nil, userPhoneDomain.ErrPhoneAlreadyExists
	}

	var verificationExpiresAt *time.Time
	if cmd.VerificationExpiresAt != nil && *cmd.VerificationExpiresAt != "" {
		parsed, err := time.Parse(time.RFC3339, *cmd.VerificationExpiresAt)
		if err != nil {
			return uuid.Nil, err
		}
		verificationExpiresAt = &parsed
	}

	// Create domain entity
	userPhone, err := userPhoneDomain.NewUserPhone(userPhoneDomain.NewUserPhoneParams{
		UserID:                cmd.UserID,
		Phone:                 cmd.Phone,
		CountryCode:           cmd.CountryCode,
		IsPrimary:             cmd.IsPrimary,
		IsVerified:            cmd.IsVerified,
		VerificationCode:      cmd.VerificationCode,
		VerificationExpiresAt: verificationExpiresAt,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.userPhoneRepo.Create.New(ctx, userPhone); err != nil {
		return uuid.Nil, err
	}

	return userPhone.ID(), nil
}
