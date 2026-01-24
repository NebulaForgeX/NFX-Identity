package auth

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	authCommands "nfxid/modules/auth/application/auth/commands"
	emailPkg "nfxid/pkgs/email"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// SendVerificationCode 发送邮箱验证码
func (s *Service) SendVerificationCode(ctx context.Context, cmd authCommands.SendVerificationCodeCmd) error {
	if s.emailService == nil {
		return fmt.Errorf("email service not configured")
	}
	if s.cache == nil {
		return fmt.Errorf("cache not configured")
	}

	email := cmd.Email
	if email == "" {
		return fmt.Errorf("email is required")
	}

	// 检查邮箱是否已被注册
	userEmail, err := s.grpcClients.DirectoryClient.UserEmail.GetUserEmailByEmail(ctx, email)
	if err == nil && userEmail != nil {
		// 邮箱已存在，检查是否已验证
		if userEmail.IsVerified {
			// 邮箱已验证，提示用户登录
			return ErrEmailAlreadyVerified
		}
		// 邮箱存在但未验证，允许继续注册流程（重新发送验证码）
	}

	// 生成6位数字验证码
	code := generateVerificationCode(6)

	// 存储验证码到 Redis（5分钟有效期）
	key := fmt.Sprintf("verification_code:email:%s", email)
	if err := s.cache.Client().Set(ctx, key, code, 5*time.Minute).Err(); err != nil {
		return fmt.Errorf("failed to store verification code: %w", err)
	}

	// 发送邮件（使用模板）
	subject := "NFX-Identity 注册验证码"
	body := emailPkg.BuildVerificationEmailHTML(code)

	err = s.emailService.Send(emailPkg.EmailMessage{
		To:      []string{email},
		Subject: subject,
		Body:    body,
		IsHTML:  true,
	})
	if err != nil {
		return fmt.Errorf("failed to send verification code email: %w", err)
	}

	return nil
}

// generateVerificationCode 生成指定长度的数字验证码
func generateVerificationCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < length; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	return code
}
