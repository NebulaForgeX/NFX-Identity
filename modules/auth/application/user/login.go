package user

import (
	"context"
	"strings"

	"nfxid/events"
	userCommands "nfxid/modules/auth/application/user/commands"
	"nfxid/modules/auth/application/user/views"
	userDomain "nfxid/modules/auth/domain/user"
	userDomainErrors "nfxid/modules/auth/domain/user/errors"
	"nfxid/pkgs/eventbus"
	"nfxid/pkgs/safeexec"

	"golang.org/x/crypto/bcrypt"
)

// Login 登录（支持用户名、邮箱或手机号）
func (s *Service) Login(ctx context.Context, cmd userCommands.LoginCmd) (*userCommands.LoginResponse, error) {
	if cmd.Identifier == "" {
		return nil, userDomainErrors.ErrUserEmailRequired
	}
	if cmd.Password == "" {
		return nil, userDomainErrors.ErrUserPasswordRequired
	}

	// 根据 identifier 查找用户
	var entity *userDomain.User
	var err error

	identifier := strings.TrimSpace(cmd.Identifier)

	// 尝试按用户名查找
	if entity, err = s.userRepo.Get.ByUsername(ctx, identifier); err == nil && entity != nil {
		// 找到用户，继续验证密码
	} else if strings.Contains(identifier, "@") {
		// 包含 @，按邮箱查找
		entity, err = s.userRepo.Get.ByEmail(ctx, identifier)
	} else {
		// 按手机号查找
		entity, err = s.userRepo.Get.ByPhone(ctx, identifier)
	}

	if err != nil || entity == nil {
		return nil, userDomainErrors.ErrInvalidCredentials
	}

	// 检查用户状态
	if !entity.IsActive() {
		return nil, userDomainErrors.ErrUserInactive
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(entity.Editable().Password), []byte(cmd.Password)); err != nil {
		return nil, userDomainErrors.ErrInvalidCredentials
	}

	// 生成 Token 对
	// Note: roleID is deprecated, passing empty string since we now support multiple roles
	// Future: tokenx should support roles array
	phone := ""
	if entity.Editable().Phone != nil {
		phone = *entity.Editable().Phone
	}
	accessToken, refreshToken, err := s.tokenx.GenerateTokenPair(
		entity.ID().String(),
		entity.Editable().Username,
		entity.Editable().Email,
		phone,
		"", // roleID is deprecated, using empty string
	)
	if err != nil {
		return nil, err
	}

	// 获取用户视图（Domain View）
	domainView, err := s.userQuery.Single.ByID(ctx, entity.ID())
	if err != nil {
		return nil, err
	}

	// 转换为 Application View
	userView := views.UserViewMapper(*domainView)

	// 更新最后登录时间
	entity.UpdateLastLogin()
	if err := s.userRepo.Update.Generic(ctx, entity); err != nil {
		// 记录错误但不阻止登录
	}

	// 发布登录成功事件（Auth -> Auth 内部事件）
	safeexec.SafeGo(func() error {
		event := events.AuthToAuth_SuccessEvent{
			Operation: "auth.login",
			EntityID:  entity.ID().String(),
			UserID:    entity.ID().String(),
			Details: map[string]interface{}{
				"identifier": identifier,
				"username":   entity.Editable().Username,
			},
		}
		return eventbus.PublishEvent(context.Background(), s.busPublisher, event)
	})

	return &userCommands.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         userView,
	}, nil
}
