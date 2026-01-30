package auth

import (
	"context"
	"strings"
	"time"

	"nfxid/constants"
	authCommands "nfxid/modules/auth/application/auth/commands"
	authResults "nfxid/modules/auth/application/auth/results"
	accountLockoutDomain "nfxid/modules/auth/domain/account_lockouts"
	loginAttemptDomain "nfxid/modules/auth/domain/login_attempts"
	refreshTokenDomain "nfxid/modules/auth/domain/refresh_tokens"
	grantpb "nfxid/protos/gen/access/grant"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// LoginByPhone 手机号登录：解析用户 → 校验密码 → 记录尝试 → 签发 Token
func (s *Service) LoginByPhone(ctx context.Context, cmd authCommands.LoginByPhoneCmd) (authResults.LoginResult, error) {
	if cmd.Password == "" {
		return authResults.LoginResult{}, ErrInvalidCredentials
	}

	phone := strings.TrimSpace(cmd.Phone)
	countryCode := strings.TrimSpace(cmd.CountryCode)
	if phone == "" || countryCode == "" {
		return authResults.LoginResult{}, ErrInvalidCredentials
	}

	identifier := countryCode + phone

	// 解析用户信息
	up, err := s.grpcClients.DirectoryClient.UserPhone.GetUserPhoneByCountryCodeAndPhone(ctx, countryCode, phone)
	if err != nil {
		s.recordFailedLoginByPhone(ctx, identifier, nil, loginAttemptDomain.FailureCodeUserNotFound, cmd.IP)
		return authResults.LoginResult{}, ErrInvalidCredentials
	}
	if up == nil || up.UserId == "" {
		s.recordFailedLoginByPhone(ctx, identifier, nil, loginAttemptDomain.FailureCodeUserNotFound, cmd.IP)
		return authResults.LoginResult{}, ErrInvalidCredentials
	}

	// 获取 username
	u, err := s.grpcClients.DirectoryClient.User.GetUserByID(ctx, up.UserId)
	username := ""
	if err == nil && u != nil {
		username = u.Username
	}

	uid, err := uuid.Parse(up.UserId)
	if err != nil {
		s.recordFailedLoginByPhone(ctx, identifier, nil, loginAttemptDomain.FailureCodeUserNotFound, cmd.IP)
		return authResults.LoginResult{}, ErrInvalidCredentials
	}

	// 检查账户是否被锁定
	locked, err := s.accountLockoutRepo.Check.IsLocked(ctx, uid)
	if err != nil {
		// 检查失败不影响登录流程，继续
	}
	if locked {
		s.recordFailedLoginByPhone(ctx, identifier, &uid, loginAttemptDomain.FailureCodeLocked, cmd.IP)
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
			s.recordFailedLoginByPhone(ctx, identifier, &uid, loginAttemptDomain.FailureCodeLocked, cmd.IP)
			return authResults.LoginResult{}, ErrAccountLocked
		}
	}

	// 获取用户凭证
	uc, err := s.credRepo.Get.ByUserID(ctx, uid)
	if err != nil {
		s.recordFailedLoginByPhone(ctx, identifier, &uid, loginAttemptDomain.FailureCodeUserNotFound, cmd.IP)
		return authResults.LoginResult{}, ErrInvalidCredentials
	}
	if uc.PasswordHash() == nil || *uc.PasswordHash() == "" {
		s.recordFailedLoginByPhone(ctx, identifier, &uid, loginAttemptDomain.FailureCodeUserNotFound, cmd.IP)
		return authResults.LoginResult{}, ErrInvalidCredentials
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(*uc.PasswordHash()), []byte(cmd.Password)); err != nil {
		s.recordFailedLoginByPhone(ctx, identifier, &uid, loginAttemptDomain.FailureCodeBadPassword, cmd.IP)
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
	s.clearFailedAttemptsByPhone(ctx, uid)

	// 记录成功登录
	s.recordSuccessfulLoginByPhone(ctx, identifier, &uid, cmd.IP)

	// 更新最后成功登录时间
	_ = s.credRepo.Update.UpdateLastSuccessLogin(ctx, uid)

	// 签发 Token
	if username == "" {
		username = "user"
	}

	// 获取用户的邮箱信息（如果有的话）
	var email string
	userEmails, err := s.grpcClients.DirectoryClient.UserEmail.GetUserEmailsByUserID(ctx, up.UserId)
	if err == nil && len(userEmails) > 0 {
		// 使用第一个邮箱（通常是主邮箱）
		ue := userEmails[0]
		if ue.Email != "" {
			email = ue.Email
		}
	}

	// 获取用户的角色信息（NewGRPCClients 已保证 AccessClient 非 nil）
	var roleID string
	grants, err := s.grpcClients.AccessClient.Grant.GetGrantsBySubject(ctx, "user", up.UserId, nil)
	if err == nil {
		for _, grant := range grants {
			if grant.GrantType == grantpb.AccessGrantType_ACCESS_GRANT_TYPE_ROLE && grant.RevokedAt == nil {
				roleID = grant.GrantRefId
				break
			}
		}
	}

	// 生成 refresh token ID（用于 refresh_tokens 表）
	refreshTokenID, err := uuid.NewV7()
	if err != nil {
		return authResults.LoginResult{}, err
	}
	refreshTokenIDStr := refreshTokenID.String()

	// 生成 Token 对（refresh token 包含 token_id/jti）
	access, refresh, err := s.tokenIssuer.IssuePairWithRefreshID(up.UserId, username, email, phone, countryCode, roleID, refreshTokenIDStr)
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
		UserID:       up.UserId,
	}, nil
}

// recordFailedLoginByPhone 记录失败的登录尝试
func (s *Service) recordFailedLoginByPhone(ctx context.Context, identifier string, userID *uuid.UUID, failureCode loginAttemptDomain.FailureCode, ip *string) {
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

// recordSuccessfulLoginByPhone 记录成功的登录尝试
func (s *Service) recordSuccessfulLoginByPhone(ctx context.Context, identifier string, userID *uuid.UUID, ip *string) {
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

// clearFailedAttemptsByPhone 清空用户的失败登录记录（删除最近的失败记录）
func (s *Service) clearFailedAttemptsByPhone(ctx context.Context, userID uuid.UUID) {
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
