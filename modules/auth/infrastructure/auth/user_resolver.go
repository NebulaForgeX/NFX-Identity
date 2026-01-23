package auth

import (
	"context"
	"fmt"

	authApp "nfxid/modules/auth/application/auth"
	authGrpc "nfxid/modules/auth/infrastructure/grpc"
)

// UserResolver 实现 application/auth.UserResolver，委托 Directory gRPC
type UserResolver struct {
	dir *authGrpc.DirectoryClient
}

// NewUserResolver 创建 UserResolver 适配器
func NewUserResolver(dir *authGrpc.DirectoryClient) *UserResolver {
	return &UserResolver{dir: dir}
}

// ResolveByEmail 实现 UserResolver
func (r *UserResolver) ResolveByEmail(ctx context.Context, email string) (authApp.UserInfo, error) {
	ue, err := r.dir.UserEmail.GetUserEmailByEmail(ctx, email)
	if err != nil {
		return authApp.UserInfo{}, err
	}
	if ue == nil || ue.UserId == "" {
		return authApp.UserInfo{}, fmt.Errorf("user not found by email")
	}
	username, _ := r.usernameByUserID(ctx, ue.UserId)
	return authApp.UserInfo{UserID: ue.UserId, Username: username}, nil
}

// ResolveByPhone 实现 UserResolver
func (r *UserResolver) ResolveByPhone(ctx context.Context, phone string) (authApp.UserInfo, error) {
	up, err := r.dir.UserPhone.GetUserPhoneByPhone(ctx, phone)
	if err != nil {
		return authApp.UserInfo{}, err
	}
	if up == nil || up.UserId == "" {
		return authApp.UserInfo{}, fmt.Errorf("user not found by phone")
	}
	username, _ := r.usernameByUserID(ctx, up.UserId)
	return authApp.UserInfo{UserID: up.UserId, Username: username}, nil
}

func (r *UserResolver) usernameByUserID(ctx context.Context, userID string) (string, error) {
	u, err := r.dir.User.GetUserByID(ctx, userID)
	if err != nil || u == nil {
		return "", err
	}
	return u.Username, nil
}
