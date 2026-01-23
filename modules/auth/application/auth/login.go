package auth

import (
	"context"
	"strings"
	"time"

	authCommands "nfxid/modules/auth/application/auth/commands"
	authResults "nfxid/modules/auth/application/auth/results"
	accountLockoutDomain "nfxid/modules/auth/domain/account_lockouts"
	loginAttemptDomain "nfxid/modules/auth/domain/login_attempts"
	refreshTokenDomain "nfxid/modules/auth/domain/refresh_tokens"
	"nfxid/constants"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Login 执行登录：检查锁定 → 解析用户 → 校验密码 → 记录尝试 → 签发 Token
func (s *Service) Login(ctx context.Context, cmd authCommands.LoginCmd) (authResults.LoginResult, error) {
	if cmd.LoginType != "email" && cmd.LoginType != "phone" {
		return authResults.LoginResult{}, ErrInvalidCredentials
	}
	if cmd.Password == "" {
		return authResults.LoginResult{}, ErrInvalidCredentials
	}

	// 确定登录标识符（email 或 phone）
	var identifier string
	if cmd.LoginType == "email" {
		if cmd.Email == nil || strings.TrimSpace(*cmd.Email) == "" {
			return authResults.LoginResult{}, ErrInvalidCredentials
		}
		identifier = strings.TrimSpace(*cmd.Email)
	} else {
		if cmd.Phone == nil || strings.TrimSpace(*cmd.Phone) == "" {
			return authResults.LoginResult{}, ErrInvalidCredentials
		}
		phone := strings.TrimSpace(*cmd.Phone)
		cc := ""
		if cmd.CountryCode != nil {
			cc = strings.TrimSpace(*cmd.CountryCode)
		}
		identifier = phone
		if cc != "" {
			identifier = cc + phone
		}
	}

	// 解析用户信息
	var info UserInfo
	var err error
	if cmd.LoginType == "email" {
		info, err = s.userResolver.ResolveByEmail(ctx, identifier)
		if err != nil {
			s.recordFailedLogin(ctx, identifier, nil, loginAttemptDomain.FailureCodeUserNotFound, cmd.IP)
			return authResults.LoginResult{}, ErrInvalidCredentials
		}
		info.Email = identifier
	} else {
		info, err = s.userResolver.ResolveByPhone(ctx, identifier)
		if err != nil {
			s.recordFailedLogin(ctx, identifier, nil, loginAttemptDomain.FailureCodeUserNotFound, cmd.IP)
			return authResults.LoginResult{}, ErrInvalidCredentials
		}
		info.Phone = identifier
	}
	if info.UserID == "" {
		s.recordFailedLogin(ctx, identifier, nil, loginAttemptDomain.FailureCodeUserNotFound, cmd.IP)
		return authResults.LoginResult{}, ErrInvalidCredentials
	}

	uid, err := uuid.Parse(info.UserID)
	if err != nil {
		s.recordFailedLogin(ctx, identifier, &uid, loginAttemptDomain.FailureCodeUserNotFound, cmd.IP)
		return authResults.LoginResult{}, ErrInvalidCredentials
	}

	// 检查账户是否被锁定
	locked, err := s.accountLockoutRepo.Check.IsLocked(ctx, uid)
	if err != nil {
		// 检查失败不影响登录流程，继续
	}
	if locked {
		s.recordFailedLogin(ctx, identifier, &uid, loginAttemptDomain.FailureCodeLocked, cmd.IP)
		return authResults.LoginResult{}, ErrAccountLocked
	}

	// 检查最近的失败次数
	recentAttempts, err := s.loginAttemptRepo.Get.ByUserID(ctx, uid)
	if err == nil {
		failedCount := 0
		now := time.Now()
		// 统计最近 1 小时内的失败次数
		for _, attempt := range recentAttempts {
			if !attempt.Success() && attempt.CreatedAt().After(now.Add(-1*time.Hour)) {
				failedCount++
			}
		}
		if failedCount >= constants.MaxLoginAttempts {
			// 锁定账户
			lockedUntil := now.Add(time.Duration(constants.LockoutDurationMinutes) * time.Minute)
			lockout, err := accountLockoutDomain.NewAccountLockout(accountLockoutDomain.NewAccountLockoutParams{
				UserID:      uid,
				LockedUntil: &lockedUntil,
				LockReason:  accountLockoutDomain.LockReasonTooManyAttempts,
			})
			if err == nil {
				_ = s.accountLockoutRepo.Create.New(ctx, lockout)
			}
			s.recordFailedLogin(ctx, identifier, &uid, loginAttemptDomain.FailureCodeLocked, cmd.IP)
			return authResults.LoginResult{}, ErrAccountLocked
		}
	}

	// 获取用户凭证
	uc, err := s.credRepo.Get.ByUserID(ctx, uid)
	if err != nil {
		s.recordFailedLogin(ctx, identifier, &uid, loginAttemptDomain.FailureCodeUserNotFound, cmd.IP)
		return authResults.LoginResult{}, ErrInvalidCredentials
	}
	if uc.PasswordHash() == nil || *uc.PasswordHash() == "" {
		s.recordFailedLogin(ctx, identifier, &uid, loginAttemptDomain.FailureCodeUserNotFound, cmd.IP)
		return authResults.LoginResult{}, ErrInvalidCredentials
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(*uc.PasswordHash()), []byte(cmd.Password)); err != nil {
		s.recordFailedLogin(ctx, identifier, &uid, loginAttemptDomain.FailureCodeBadPassword, cmd.IP)
		// 检查是否达到最大失败次数
		recentAttempts, err := s.loginAttemptRepo.Get.ByUserID(ctx, uid)
		if err == nil {
			failedCount := 0
			now := time.Now()
			for _, attempt := range recentAttempts {
				if !attempt.Success() && attempt.CreatedAt().After(now.Add(-1*time.Hour)) {
					failedCount++
				}
			}
			if failedCount+1 >= constants.MaxLoginAttempts {
				// 锁定账户
				lockedUntil := now.Add(time.Duration(constants.LockoutDurationMinutes) * time.Minute)
				lockout, err := accountLockoutDomain.NewAccountLockout(accountLockoutDomain.NewAccountLockoutParams{
					UserID:      uid,
					LockedUntil: &lockedUntil,
					LockReason:  accountLockoutDomain.LockReasonTooManyAttempts,
				})
				if err == nil {
					_ = s.accountLockoutRepo.Create.New(ctx, lockout)
				}
			}
		}
		return authResults.LoginResult{}, ErrInvalidCredentials
	}

	// 登录成功：清空失败次数（删除最近的失败记录）
	s.clearFailedAttempts(ctx, uid)

	// 记录成功登录
	s.recordSuccessfulLogin(ctx, identifier, &uid, cmd.IP)

	// 签发 Token
	username := info.Username
	if username == "" {
		username = "user"
	}

	// 生成 refresh token ID（用于 refresh_tokens 表）
	refreshTokenID, err := uuid.NewV7()
	if err != nil {
		return authResults.LoginResult{}, err
	}
	refreshTokenIDStr := refreshTokenID.String()

	// 生成 Token 对（refresh token 包含 token_id/jti）
	access, refresh, err := s.tokenIssuer.IssuePairWithRefreshID(info.UserID, username, info.Email, info.Phone, "", refreshTokenIDStr)
	if err != nil {
		return authResults.LoginResult{}, err
	}

	// 创建 refresh_tokens 表记录
	expiresAt := time.Now().Add(time.Duration(s.refreshTokenTTL) * time.Second)
	refreshTokenEntity, err := refreshTokenDomain.NewRefreshToken(refreshTokenDomain.NewRefreshTokenParams{
		TokenID:   refreshTokenIDStr,
		UserID:    uid,
		ExpiresAt: expiresAt,
		IP:        cmd.IP,
	})
	if err == nil {
		_ = s.refreshTokenRepo.Create.New(ctx, refreshTokenEntity)
	}

	return authResults.LoginResult{
		AccessToken:  access,
		RefreshToken: refresh,
		ExpiresIn:    s.expiresInSec,
		UserID:       info.UserID,
	}, nil
}

// recordFailedLogin 记录失败的登录尝试
func (s *Service) recordFailedLogin(ctx context.Context, identifier string, userID *uuid.UUID, failureCode loginAttemptDomain.FailureCode, ip *string) {
	attempt, err := loginAttemptDomain.NewLoginAttempt(loginAttemptDomain.NewLoginAttemptParams{
		Identifier:  identifier,
		UserID:      userID,
		IP:          ip,
		Success:     false,
		FailureCode: &failureCode,
	})
	if err == nil {
		_ = s.loginAttemptRepo.Create.New(ctx, attempt)
	}
}

// recordSuccessfulLogin 记录成功的登录尝试
func (s *Service) recordSuccessfulLogin(ctx context.Context, identifier string, userID *uuid.UUID, ip *string) {
	attempt, err := loginAttemptDomain.NewLoginAttempt(loginAttemptDomain.NewLoginAttemptParams{
		Identifier: identifier,
		UserID:     userID,
		IP:         ip,
		Success:    true,
	})
	if err == nil {
		_ = s.loginAttemptRepo.Create.New(ctx, attempt)
	}
}

// clearFailedAttempts 清空用户的失败登录记录（删除最近的失败记录）
func (s *Service) clearFailedAttempts(ctx context.Context, userID uuid.UUID) {
	attempts, err := s.loginAttemptRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return
	}
	// 删除最近的失败记录（保留成功记录）
	for _, attempt := range attempts {
		if !attempt.Success() {
			_ = s.loginAttemptRepo.Delete.ByID(ctx, attempt.ID())
		}
	}
}
