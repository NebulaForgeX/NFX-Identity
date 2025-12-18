package auth

import (
	"context"
	"strings"

	authCommands "nfxid/modules/permission/application/auth/commands"
	authorizationCodeApp "nfxid/modules/permission/application/authorization_code"
	userPermissionApp "nfxid/modules/permission/application/user_permission"
	userPermissionAppCommands "nfxid/modules/permission/application/user_permission/commands"
	userPermissionViews "nfxid/modules/permission/application/user_permission/views"
	"nfxid/modules/permission/infrastructure/grpcclient"
	authpb "nfxid/protos/gen/auth/auth"
	"nfxid/pkgs/tokenx"
	"github.com/google/uuid"
)

type Service struct {
	authGRPCClient        *grpcclient.AuthGRPCClient
	userPermissionSvc     *userPermissionApp.Service
	authorizationCodeSvc  *authorizationCodeApp.Service
	tokenx                *tokenx.Tokenx
}

func NewService(
	authGRPCClient *grpcclient.AuthGRPCClient,
	userPermissionSvc *userPermissionApp.Service,
	authorizationCodeSvc *authorizationCodeApp.Service,
	tokenx *tokenx.Tokenx,
) *Service {
	return &Service{
		authGRPCClient:       authGRPCClient,
		userPermissionSvc:    userPermissionSvc,
		authorizationCodeSvc: authorizationCodeSvc,
		tokenx:               tokenx,
	}
}

// Login 登录（支持用户名、邮箱、手机号密码登录，以及邮箱验证码登录）
func (s *Service) Login(ctx context.Context, cmd authCommands.LoginCmd) (*authCommands.LoginResponse, error) {
	if cmd.Identifier == "" {
		return nil, ErrLoginIdentifierRequired
	}

	identifier := strings.TrimSpace(cmd.Identifier)

	// 根据 Type 和 Identifier 调用 auth 服务验证用户
	var auth *authpb.Auth
	var err error

	if cmd.Type == "code" {
		// 邮箱验证码登录（未来实现）
		return nil, ErrEmailCodeNotImplemented
	}

	// 密码登录：通过 gRPC 调用 auth 服务
	if strings.Contains(identifier, "@") {
		// 邮箱登录
		auth, err = s.authGRPCClient.GetAuthByEmail(ctx, identifier, false, false)
		if err != nil {
			return nil, ErrInvalidCredentials
		}
	} else if len(identifier) > 0 && (identifier[0] == '+' || (len(identifier) >= 10 && len(identifier) <= 15)) {
		// 手机号登录
		auth, err = s.authGRPCClient.GetAuthByPhone(ctx, identifier, false, false)
		if err != nil {
			return nil, ErrInvalidCredentials
		}
	} else {
		// 用户名登录
		auth, err = s.authGRPCClient.GetAuthByUsername(ctx, identifier, false, false)
		if err != nil {
			return nil, ErrInvalidCredentials
		}
	}

	// 验证密码
	valid, err := s.authGRPCClient.VerifyPassword(ctx, auth.Id, cmd.Password)
	if err != nil || !valid {
		return nil, ErrInvalidCredentials
	}

	// 检查用户状态
	if auth.Status != authpb.AuthStatus_AUTH_STATUS_ACTIVE {
		return nil, ErrUserInactive
	}

	userID := auth.Id
	username := auth.Username
	email := auth.Email
	phone := auth.Phone

	// 获取用户权限
	userUUID, _ := uuid.Parse(userID)
	permissions, err := s.userPermissionSvc.GetUserPermissions(ctx, userPermissionAppCommands.GetUserPermissionsCmd{
		UserID: userUUID,
	})
	if err != nil {
		// 如果查询权限失败，返回空权限列表（外部用户）
		permissions = []*userPermissionViews.UserPermissionView{}
	}

	// 提取权限标签
	permissionTags := make([]string, 0, len(permissions))
	for _, p := range permissions {
		if p.Tag != "" {
			permissionTags = append(permissionTags, p.Tag)
		}
	}

	// 生成 Token 对
	accessToken, refreshToken, err := s.tokenx.GenerateTokenPair(
		userID,
		username,
		email,
		phone,
		"", // role_id 不再使用
	)
	if err != nil {
		return nil, err
	}

	return &authCommands.LoginResponse{
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		UserID:         userID,
		Username:       username,
		Email:          email,
		Phone:          phone,
		Permissions:    permissions,
		PermissionTags: permissionTags,
	}, nil
}

