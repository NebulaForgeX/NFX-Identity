package auth

import (
	userPermissionViews "nfxid/modules/permission/application/user_permission/views"
)

// LoginCmd 登录命令（支持多种方式）
type LoginCmd struct {
	// 登录方式：username, email, phone, email_code
	Type       string // "password" or "code"
	Identifier string // username, email 或 phone
	Password   string // 密码（当 Type=password 时）
	Code       string // 验证码（当 Type=code 时，仅支持邮箱验证码）
}

// LoginResponse 登录响应
type LoginResponse struct {
	AccessToken  string                                    `json:"access_token"`
	RefreshToken string                                    `json:"refresh_token"`
	UserID       string                                    `json:"user_id"`
	Username     string                                    `json:"username"`
	Email        string                                    `json:"email"`
	Phone        string                                    `json:"phone"`
	Permissions  []*userPermissionViews.UserPermissionView `json:"permissions"`
	PermissionTags []string                                `json:"permission_tags"` // 权限标签列表，前端用于控制显示
}

