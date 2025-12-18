package authorization_code

import (
	"context"
	authorizationCodeCommands "nfxid/modules/permission/application/authorization_code/commands"
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"
)

func (s *Service) GetAuthorizationCode(ctx context.Context, cmd authorizationCodeCommands.GetAuthorizationCodeCmd) (*authorizationCodeDomain.AuthorizationCode, error) {
	return s.authorizationCodeRepo.Get.ByID(ctx, cmd.ID)
}

func (s *Service) GetAuthorizationCodeByCode(ctx context.Context, cmd authorizationCodeCommands.GetAuthorizationCodeByCodeCmd) (*authorizationCodeDomain.AuthorizationCode, error) {
	return s.authorizationCodeRepo.Get.ByCode(ctx, cmd.Code)
}
