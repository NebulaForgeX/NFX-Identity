package authorization_code

import (
	"context"
	authorizationCodeCommands "nfxid/modules/permission/application/authorization_code/commands"
)

func (s *Service) UseAuthorizationCode(ctx context.Context, cmd authorizationCodeCommands.UseAuthorizationCodeCmd) error {
	ac, err := s.authorizationCodeRepo.Get.ByCode(ctx, cmd.Code)
	if err != nil {
		return err
	}

	if err := ac.Use(); err != nil {
		return err
	}

	if err := s.authorizationCodeRepo.Update.Generic(ctx, ac); err != nil {
		return err
	}

	return nil
}
