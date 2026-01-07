package authorization_code

import (
	"context"
	authorizationCodeCommands "nfxid/modules/permission/application/authorization_code/commands"
)

func (s *Service) ActivateAuthorizationCode(ctx context.Context, cmd authorizationCodeCommands.ActivateAuthorizationCodeCmd) error {
	ac, err := s.authorizationCodeRepo.Get.ByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err := ac.Activate(); err != nil {
		return err
	}

	if err := s.authorizationCodeRepo.Update.Generic(ctx, ac); err != nil {
		return err
	}

	return nil
}
