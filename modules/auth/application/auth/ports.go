package auth

import "context"

// UserResolver 解析用户标识（邮箱/手机）→ userID、username。
// 由 infrastructure 实现（如 Directory gRPC 适配器）。
type UserResolver interface {
	ResolveByEmail(ctx context.Context, email string) (UserInfo, error)
	ResolveByPhone(ctx context.Context, phone string) (UserInfo, error)
}

// UserInfo 登录流程中解析到的用户信息（Email/Phone 在 use case 内根据登录方式填入）
type UserInfo struct {
	UserID   string
	Username string
	Email    string
	Phone    string
}

// TokenIssuer 签发与刷新 JWT。
// 由 infrastructure 实现（如 tokenx 适配器）。
type TokenIssuer interface {
	IssuePair(userID, username, email, phone, roleID string) (access, refresh string, err error)
	RefreshPair(refreshToken string) (access, refresh string, err error)
}
