package user

import "nfxid/modules/auth/application/user/views"

// LoginCmd 登录命令
type LoginCmd struct {
	Identifier string // username, email 或 phone
	Password   string
}

// RefreshCmd 刷新 Token 命令
type RefreshCmd struct {
	RefreshToken string
}

// LoginResponse 登录响应
type LoginResponse struct {
	AccessToken  string
	RefreshToken string
	User         views.UserView
}

// RefreshResponse 刷新 Token 响应
type RefreshResponse struct {
	AccessToken  string
	RefreshToken string
}
