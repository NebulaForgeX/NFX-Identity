package authorization_code

import (
	"context"
	authorizationCodeCommands "nfxid/modules/permission/application/authorization_code/commands"
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"
)

func (s *Service) CreateAuthorizationCode(ctx context.Context, cmd authorizationCodeCommands.CreateAuthorizationCodeCmd) (*authorizationCodeDomain.AuthorizationCode, error) {
	ac, err := authorizationCodeDomain.NewAuthorizationCode(authorizationCodeDomain.NewAuthorizationCodeParams{
		Editable: authorizationCodeDomain.AuthorizationCodeEditable{
			Code:      cmd.Code,
			MaxUses:   cmd.MaxUses,
			UsedCount: 0,
		},
		CreatedBy: cmd.CreatedBy,
		ExpiresAt: cmd.ExpiresAt,
		IsActive:  cmd.IsActive,
	})
	if err != nil {
		return nil, err
	}

	if err := s.authorizationCodeRepo.Create.New(ctx, ac); err != nil {
		return nil, err
	}

	return ac, nil
}
