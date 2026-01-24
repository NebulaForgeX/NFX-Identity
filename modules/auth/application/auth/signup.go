package auth

import (
	"context"
	"fmt"
	"strings"
	"time"

	authCommands "nfxid/modules/auth/application/auth/commands"
	authResults "nfxid/modules/auth/application/auth/results"
	refreshTokenDomain "nfxid/modules/auth/domain/refresh_tokens"
	userCredentialDomain "nfxid/modules/auth/domain/user_credentials"
	userCredentialAppCommands "nfxid/modules/auth/application/user_credentials/commands"
	"nfxid/connections/directory/dto"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Signup 用户注册（邮箱+密码+验证码）
func (s *Service) Signup(ctx context.Context, cmd authCommands.SignupCmd) (authResults.SignupResult, error) {
	if s.emailService == nil {
		return authResults.SignupResult{}, fmt.Errorf("email service not configured")
	}
	if s.cache == nil {
		return authResults.SignupResult{}, fmt.Errorf("cache not configured")
	}

	email := strings.TrimSpace(cmd.Email)
	if email == "" {
		return authResults.SignupResult{}, fmt.Errorf("email is required")
	}
	if cmd.Password == "" {
		return authResults.SignupResult{}, fmt.Errorf("password is required")
	}
	if cmd.VerificationCode == "" {
		return authResults.SignupResult{}, fmt.Errorf("verification code is required")
	}

	// 验证验证码
	key := fmt.Sprintf("verification_code:email:%s", email)
	storedCode, err := s.cache.Client().Get(ctx, key).Result()
	if err != nil {
		return authResults.SignupResult{}, ErrInvalidVerificationCode
	}
	if storedCode != cmd.VerificationCode {
		return authResults.SignupResult{}, ErrInvalidVerificationCode
	}

	// 检查邮箱是否已被注册
	userEmail, err := s.grpcClients.DirectoryClient.UserEmail.GetUserEmailByEmail(ctx, email)
	var userIDStr string
	var userID uuid.UUID
	var username string

	if err == nil && userEmail != nil {
		// 邮箱已存在
		if userEmail.IsVerified {
			// 邮箱已验证，提示用户登录
			return authResults.SignupResult{}, ErrEmailAlreadyVerified
		}
		// 邮箱存在但未验证，继续完成注册（使用现有用户）
		userIDStr = userEmail.UserId
		userID, err = uuid.Parse(userIDStr)
		if err != nil {
			return authResults.SignupResult{}, fmt.Errorf("invalid user ID: %w", err)
		}
		
		// 获取现有用户信息
		user, err := s.grpcClients.DirectoryClient.User.GetUserByID(ctx, userIDStr)
		if err != nil {
			return authResults.SignupResult{}, fmt.Errorf("failed to get existing user: %w", err)
		}
		username = user.Username
		// 注意：邮箱验证状态暂时不更新（因为 connections 中没有更新方法）
		// 验证码已经验证了邮箱的所有权，所以我们可以继续流程
	} else {
		// 邮箱不存在，创建新用户
		// 生成唯一用户名（使用 nfxid-<uuid> 格式避免邮箱前缀重复）
		userUUID, err := uuid.NewV7()
		if err != nil {
			return authResults.SignupResult{}, fmt.Errorf("failed to generate user UUID: %w", err)
		}
		username = fmt.Sprintf("nfxid-%s", userUUID.String())

		// 1. 创建用户（新用户，状态为 active，已验证）
		userIDStr, err = s.grpcClients.DirectoryClient.User.CreateUser(ctx, username, "active", true)
		if err != nil {
			return authResults.SignupResult{}, fmt.Errorf("failed to create user: %w", err)
		}

		userID, err = uuid.Parse(userIDStr)
		if err != nil {
			return authResults.SignupResult{}, fmt.Errorf("invalid user ID returned: %w", err)
		}

		// 2. 创建用户邮箱（主邮箱，已验证）
		_, err = s.grpcClients.DirectoryClient.UserEmail.CreateUserEmail(ctx, &dto.CreateUserEmailDTO{
			UserID:            userIDStr,
			Email:             email,
			IsPrimary:         true,
			IsVerified:        true,
			VerificationToken: nil,
		})
		if err != nil {
			return authResults.SignupResult{}, fmt.Errorf("failed to create user email: %w", err)
		}

		// 3. 创建用户资料（空资料）- 仅新用户
		_, err = s.grpcClients.DirectoryClient.UserProfile.CreateUserProfileDefault(ctx, userIDStr)
		if err != nil {
			// 资料创建失败不影响注册流程，继续
		}

		// 4. 创建用户偏好（默认值）- 仅新用户
		_, err = s.grpcClients.DirectoryClient.UserPreference.CreateUserPreferenceDefault(ctx, userIDStr)
		if err != nil {
			// 偏好创建失败不影响注册流程，继续
		}
	}

	// 检查用户是否已有凭证
	hasCredential, err := s.credRepo.Check.ByUserID(ctx, userID)
	if err == nil && hasCredential {
		// 用户已有凭证，不允许重复注册
		return authResults.SignupResult{}, fmt.Errorf("user already has credentials, please login")
	}

	// 5. 创建用户凭证（密码）- 新用户或未验证用户都没有凭证
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		return authResults.SignupResult{}, fmt.Errorf("failed to hash password: %w", err)
	}

	passwordHash := string(hashedPassword)
	hashAlg := "bcrypt"
	createCredentialCmd := userCredentialAppCommands.CreateUserCredentialCmd{
		UserID:             userID,
		CredentialType:     userCredentialDomain.CredentialTypePassword,
		PasswordHash:       &passwordHash,
		HashAlg:            &hashAlg,
		HashParams:         nil,
		Status:             userCredentialDomain.CredentialStatusActive,
		MustChangePassword: false,
	}
	_, err = s.userCredentialAppSvc.CreateUserCredential(ctx, createCredentialCmd)
	if err != nil {
		return authResults.SignupResult{}, fmt.Errorf("failed to create user credential: %w", err)
	}

	// 6. 删除验证码（已使用）
	_ = s.cache.Client().Del(ctx, key).Err()

	// 7. 签发 Token（自动登录）
	// 生成 refresh token ID
	refreshTokenID, err := uuid.NewV7()
	if err != nil {
		return authResults.SignupResult{}, err
	}
	refreshTokenIDStr := refreshTokenID.String()

	// 生成 Token 对
	access, refresh, err := s.tokenIssuer.IssuePairWithRefreshID(userIDStr, username, email, "", "", "", refreshTokenIDStr)
	if err != nil {
		return authResults.SignupResult{}, err
	}

	// 创建 refresh_tokens 表记录
	expiresAt := time.Now().Add(time.Duration(s.refreshTokenTTL) * time.Second)
	refreshTokenEntity, err := refreshTokenDomain.NewRefreshToken(refreshTokenDomain.NewRefreshTokenParams{
		TokenID:   refreshTokenIDStr,
		UserID:    userID,
		ExpiresAt: expiresAt,
		IP:        nil, // 注册时没有 IP
	})
	if err == nil {
		_ = s.refreshTokenRepo.Create.New(ctx, refreshTokenEntity)
	}

	return authResults.SignupResult{
		AccessToken:  access,
		RefreshToken: refresh,
		ExpiresIn:    s.expiresInSec,
		UserID:       userIDStr,
	}, nil
}

