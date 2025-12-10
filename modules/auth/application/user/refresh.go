package user

import (
	"context"

	userCommands "nebulaid/modules/auth/application/user/commands"
	userDomainErrors "nebulaid/modules/auth/domain/user/errors"

	"github.com/google/uuid"
)

// RefreshToken 刷新 Token
func (s *Service) RefreshToken(ctx context.Context, cmd userCommands.RefreshCmd) (*userCommands.RefreshResponse, error) {
	if cmd.RefreshToken == "" {
		return nil, userDomainErrors.ErrUserPasswordRequired
	}

	// 验证 refresh token
	claims, err := s.tokenx.VerifyRefreshToken(cmd.RefreshToken)
	if err != nil {
		return nil, userDomainErrors.ErrInvalidCredentials
	}

	// 验证用户是否存在且活跃
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return nil, userDomainErrors.ErrInvalidCredentials
	}

	entity, err := s.userRepo.GetByID(ctx, userID)
	if err != nil || entity == nil {
		return nil, userDomainErrors.ErrUserNotFound
	}

	if !entity.IsActive() {
		return nil, userDomainErrors.ErrUserInactive
	}

	// 生成新的 Token 对
	roleID := ""
	if entity.Editable().RoleID != nil {
		roleID = entity.Editable().RoleID.String()
	}

	newAccessToken, newRefreshToken, err := s.tokenx.GenerateTokenPair(
		entity.ID().String(),
		entity.Editable().Username,
		entity.Editable().Email,
		entity.Editable().Phone,
		roleID,
	)
	if err != nil {
		return nil, err
	}

	return &userCommands.RefreshResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
