package grpcclient

import (
	"context"
	"fmt"
	"nfxid/pkgs/grpcx"
	"nfxid/pkgs/security/token/servertoken"
	authpb "nfxid/protos/gen/auth/auth"

	"google.golang.org/grpc"
	"github.com/google/uuid"
)

type AuthGRPCClient struct {
	conn      *grpc.ClientConn
	AuthStub  authpb.AuthServiceClient
}

func NewAuthGRPCClient(addr string, provider servertoken.TokenProvider) (*AuthGRPCClient, error) {
	conn, err := grpc.NewClient(addr, grpcx.DefaultClientOptions(provider)...)
	if err != nil {
		return nil, err
	}

	return &AuthGRPCClient{
		conn:     conn,
		AuthStub: authpb.NewAuthServiceClient(conn),
	}, nil
}

func (c *AuthGRPCClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// GetAuthByEmail 通过邮箱获取认证信息
func (c *AuthGRPCClient) GetAuthByEmail(ctx context.Context, email string, includeRole, includeProfile bool) (*authpb.Auth, error) {
	resp, err := c.AuthStub.GetAuthByEmail(ctx, &authpb.GetAuthByEmailRequest{
		Email:         email,
		IncludeRole:   &includeRole,
		IncludeProfile: &includeProfile,
	})
	if err != nil {
		return nil, err
	}
	return resp.Auth, nil
}

// GetAuthByPhone 通过手机号获取认证信息
func (c *AuthGRPCClient) GetAuthByPhone(ctx context.Context, phone string, includeRole, includeProfile bool) (*authpb.Auth, error) {
	resp, err := c.AuthStub.GetAuthByPhone(ctx, &authpb.GetAuthByPhoneRequest{
		Phone:         phone,
		IncludeRole:   &includeRole,
		IncludeProfile: &includeProfile,
	})
	if err != nil {
		return nil, err
	}
	return resp.Auth, nil
}

// GetAuthByUsername 通过用户名获取认证信息
func (c *AuthGRPCClient) GetAuthByUsername(ctx context.Context, username string, includeRole, includeProfile bool) (*authpb.Auth, error) {
	resp, err := c.AuthStub.GetAuthByUsername(ctx, &authpb.GetAuthByUsernameRequest{
		Username:      username,
		IncludeRole:   &includeRole,
		IncludeProfile: &includeProfile,
	})
	if err != nil {
		return nil, err
	}
	return resp.Auth, nil
}

// GetAuthByUserID 通过用户ID获取认证信息
func (c *AuthGRPCClient) GetAuthByUserID(ctx context.Context, userID uuid.UUID, includeRoles, includeProfile bool) (*authpb.Auth, error) {
	resp, err := c.AuthStub.GetAuthByUserID(ctx, &authpb.GetAuthByUserIDRequest{
		UserId:        userID.String(),
		IncludeRoles:  &includeRoles,
		IncludeProfile: &includeProfile,
	})
	if err != nil {
		return nil, err
	}
	return resp.Auth, nil
}

// VerifyPassword 验证密码
func (c *AuthGRPCClient) VerifyPassword(ctx context.Context, userID, password string) (bool, error) {
	resp, err := c.AuthStub.VerifyPassword(ctx, &authpb.VerifyPasswordRequest{
		UserId:   userID,
		Password: password,
	})
	if err != nil {
		return false, err
	}
	return resp.Valid, nil
}

// VerifyUserExists 验证用户是否存在
func (c *AuthGRPCClient) VerifyUserExists(ctx context.Context, username, email, phone *string) (bool, string, error) {
	req := &authpb.VerifyUserExistsRequest{}
	if username != nil {
		req.Username = username
	}
	if email != nil {
		req.Email = email
	}
	if phone != nil {
		req.Phone = phone
	}

	resp, err := c.AuthStub.VerifyUserExists(ctx, req)
	if err != nil {
		return false, "", err
	}
	return resp.Exists, resp.Field, nil
}

// SendVerificationCode 发送验证码（发送邮件，存储到 Redis）
func (c *AuthGRPCClient) SendVerificationCode(ctx context.Context, email string) error {
	_, err := c.AuthStub.SendVerificationCode(ctx, &authpb.SendVerificationCodeRequest{
		Email: email,
	})
	return err
}

// CheckUserAndVerificationCode 检查用户是否存在并验证验证码
func (c *AuthGRPCClient) CheckUserAndVerificationCode(ctx context.Context, email, code string) (userExists bool, codeValid bool, errorMessage string, err error) {
	resp, err := c.AuthStub.CheckUserAndVerificationCode(ctx, &authpb.CheckUserAndVerificationCodeRequest{
		Email:           email,
		VerificationCode: code,
	})
	if err != nil {
		return false, false, "", err
	}
	errorMsg := ""
	if resp.ErrorMessage != nil {
		errorMsg = *resp.ErrorMessage
	}
	return resp.UserExists, resp.CodeValid, errorMsg, nil
}

// CreateUserWithProfile 创建用户和 Profile
func (c *AuthGRPCClient) CreateUserWithProfile(ctx context.Context, email, password, username string, phone *string) (userID, usernameResp, emailResp string, err error) {
	req := &authpb.CreateUserWithProfileRequest{
		Email:    email,
		Password: password,
		Username: username,
		Phone:    phone,
	}
	resp, err := c.AuthStub.CreateUserWithProfile(ctx, req)
	if err != nil {
		return "", "", "", err
	}
	if resp.ErrorMessage != nil && *resp.ErrorMessage != "" {
		return "", "", "", fmt.Errorf(*resp.ErrorMessage)
	}
	return resp.UserId, resp.Username, resp.Email, nil
}

