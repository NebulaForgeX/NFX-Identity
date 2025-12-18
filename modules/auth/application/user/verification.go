package user

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"nfxid/pkgs/email"
)

// SendVerificationCodeCmd 发送验证码命令
type SendVerificationCodeCmd struct {
	Email   string
	Purpose string // "register" 或其他用途
}

// VerifyCodeCmd 验证验证码命令
type VerifyCodeCmd struct {
	Email   string
	Code    string
	Purpose string // "register" 或其他用途
}

// generateCode 生成指定长度的数字验证码
func generateCode(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("code length must be greater than 0")
	}

	code := ""
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %w", err)
		}
		code += num.String()
	}

	return code, nil
}

// SendVerificationCode 发送验证码（发送邮件，存储到 Redis）
func (s *Service) SendVerificationCode(ctx context.Context, cmd SendVerificationCodeCmd) error {
	// 生成 6 位数字验证码
	code, err := generateCode(6)
	if err != nil {
		return fmt.Errorf("failed to generate verification code: %w", err)
	}

	// 构建 Redis key: verification_code:{purpose}:{email}
	key := fmt.Sprintf("verification_code:%s:%s", cmd.Purpose, cmd.Email)

	// 存储到 Redis，有效期 10 分钟（字符串类型，直接使用 Set）
	err = s.cache.Set(ctx, key, code, 10*time.Minute)
	if err != nil {
		return fmt.Errorf("failed to store verification code: %w", err)
	}

	// 发送邮件
	emailMsg := email.EmailMessage{
		To:      []string{cmd.Email},
		Subject: "注册验证码",
		Body:    fmt.Sprintf("您的注册验证码是：%s，有效期 10 分钟。", code),
		IsHTML:  false,
	}

	err = s.emailService.Send(emailMsg)
	if err != nil {
		return fmt.Errorf("failed to send verification email: %w", err)
	}

	return nil
}

// VerifyCode 验证验证码是否正确
func (s *Service) VerifyCode(ctx context.Context, cmd VerifyCodeCmd) error {
	// 构建 Redis key: verification_code:{purpose}:{email}
	key := fmt.Sprintf("verification_code:%s:%s", cmd.Purpose, cmd.Email)

	// 从 Redis 获取存储的验证码
	storedCode, err := s.cache.GetString(ctx, key)
	if err != nil {
		return fmt.Errorf("verification code not found or expired")
	}

	// 比较验证码
	if storedCode != cmd.Code {
		return fmt.Errorf("invalid verification code")
	}

	// 验证成功后删除验证码（一次性使用）
	err = s.cache.Delete(ctx, key)
	if err != nil {
		// 删除失败不影响验证结果，只记录日志
		// 可以考虑使用日志库记录
	}

	return nil
}
