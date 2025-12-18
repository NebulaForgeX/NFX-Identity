package authorization_code

import (
	"context"
	authorizationCodeCommands "nfxid/modules/permission/application/authorization_code/commands"
)

func (s *Service) UseAuthorizationCode(ctx context.Context, cmd authorizationCodeCommands.UseAuthorizationCodeCmd) error {
	// 使用 ByCodeAndIncrement 方法进行原子性检查和更新
	_, err := s.authorizationCodeRepo.Check.ByCodeAndIncrement(ctx, cmd.Code)
	if err != nil {
		return err
	}
	return nil
}
