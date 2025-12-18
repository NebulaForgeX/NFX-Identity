package authorization_code

import (
	"context"
	authorizationCodeCommands "nfxid/modules/permission/application/authorization_code/commands"
)

func (s *Service) DeleteAuthorizationCode(ctx context.Context, cmd authorizationCodeCommands.DeleteAuthorizationCodeCmd) error {
	ac, err := s.authorizationCodeRepo.Get.ByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err := ac.Delete(); err != nil {
		return err
	}

	if err := s.authorizationCodeRepo.Update.Generic(ctx, ac); err != nil {
		return err
	}

	return nil
}
