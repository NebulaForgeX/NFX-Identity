package auth

import (
	"context"
	"fmt"
	"strings"

	authCommands "nfxid/modules/permission/application/auth/commands"
	authorizationCodeAppCommands "nfxid/modules/permission/application/authorization_code/commands"
	userPermissionAppCommands "nfxid/modules/permission/application/user_permission/commands"
	userPermissionViews "nfxid/modules/permission/application/user_permission/views"

	"github.com/google/uuid"
)

// Register 注册（用于 Identity-Admin 平台）
// 流程：
// 1. 调用 auth service 检查用户是否存在，验证码是否正确
// 2. 如果 auth 返回错误，直接返回给前端
// 3. 如果 auth 检查通过，检查授权码是否有效（并增加 used_count）
// 4. 授权码有效后，再次调用 auth service 创建用户和 profile
func (s *Service) Register(ctx context.Context, cmd authCommands.RegisterCmd) (*authCommands.LoginResponse, error) {
	// 步骤 1: 调用 auth service 检查用户是否存在，验证码是否正确
	userExists, codeValid, errorMsg, err := s.authGRPCClient.CheckUserAndVerificationCode(ctx, cmd.Email, cmd.VerificationCode)
	if err != nil {
		return nil, fmt.Errorf("failed to check user and verification code: %w", err)
	}

	// 步骤 2: 如果 auth 返回错误，直接返回给前端
	if userExists {
		return nil, ErrUserAlreadyExists
	}
	if !codeValid {
		if errorMsg == "invalid_verification_code" {
			return nil, ErrInvalidVerificationCode
		}
		return nil, fmt.Errorf("verification code check failed: %s", errorMsg)
	}

	// 步骤 3: 如果 auth 检查通过，检查授权码是否有效（并增加 used_count）
	err = s.authorizationCodeSvc.UseAuthorizationCode(ctx, authorizationCodeAppCommands.UseAuthorizationCodeCmd{
		Code: cmd.AuthorizationCode,
	})
	if err != nil {
		return nil, fmt.Errorf("authorization code validation failed: %w", err)
	}

	// 步骤 4: 授权码有效后，再次调用 auth service 创建用户和 profile
	// 生成默认用户名（使用邮箱前缀）
	username := strings.Split(cmd.Email, "@")[0]
	userID, createdUsername, createdEmail, err := s.authGRPCClient.CreateUserWithProfile(ctx, cmd.Email, cmd.Password, username, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create user with profile: %w", err)
	}

	// 获取用户权限（新用户应该没有权限）
	userUUID, _ := uuid.Parse(userID)
	permissions, err := s.userPermissionSvc.GetUserPermissions(ctx, userPermissionAppCommands.GetUserPermissionsCmd{
		UserID: userUUID,
	})
	if err != nil {
		// 如果查询权限失败，返回空权限列表（新用户可能没有权限）
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
		createdUsername,
		createdEmail,
		"", // phone 为空
		"", // role_id 不再使用
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	return &authCommands.LoginResponse{
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		UserID:         userID,
		Username:       createdUsername,
		Email:          createdEmail,
		Phone:          "",
		Permissions:    permissions,
		PermissionTags: permissionTags,
	}, nil
}
