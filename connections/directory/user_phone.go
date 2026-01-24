package directory

import (
	"context"
	"fmt"

	"nfxid/connections/directory/dto"
	userphonepb "nfxid/protos/gen/directory/user_phone"
)

// UserPhoneClient UserPhone 客户端
type UserPhoneClient struct {
	client userphonepb.UserPhoneServiceClient
}

// NewUserPhoneClient 创建 UserPhone 客户端
func NewUserPhoneClient(client userphonepb.UserPhoneServiceClient) *UserPhoneClient {
	return &UserPhoneClient{client: client}
}

// CreateUserPhone 创建用户手机（完整参数）
func (c *UserPhoneClient) CreateUserPhone(ctx context.Context, createDTO *dto.CreateUserPhoneDTO) (string, error) {
	req := createDTO.ToCreateUserPhoneRequest()

	resp, err := c.client.CreateUserPhone(ctx, req)
	if err != nil {
		return "", fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserPhone.Id, nil
}

// CreateUserPhoneDefault 创建用户手机（默认值，用于系统初始化）
func (c *UserPhoneClient) CreateUserPhoneDefault(ctx context.Context, userID, phone, countryCode string) (string, error) {
	createDTO := &dto.CreateUserPhoneDTO{
		UserID:                userID,
		Phone:                 phone,
		CountryCode:           &countryCode,
		IsPrimary:             true,
		IsVerified:             true,
		VerificationCode:      nil,
		VerificationExpiresAt: nil,
	}

	return c.CreateUserPhone(ctx, createDTO)
}

// GetUserPhoneByID 根据ID获取用户手机
func (c *UserPhoneClient) GetUserPhoneByID(ctx context.Context, id string) (*userphonepb.UserPhone, error) {
	req := &userphonepb.GetUserPhoneByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetUserPhoneByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserPhone, nil
}

// GetUserPhonesByUserID 根据用户ID获取用户手机列表
func (c *UserPhoneClient) GetUserPhonesByUserID(ctx context.Context, userID string) ([]*userphonepb.UserPhone, error) {
	req := &userphonepb.GetUserPhonesByUserIDRequest{
		UserId: userID,
	}

	resp, err := c.client.GetUserPhonesByUserID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserPhones, nil
}

// GetUserPhoneByCountryCodeAndPhone 根据国家代码和手机号获取用户手机
func (c *UserPhoneClient) GetUserPhoneByCountryCodeAndPhone(ctx context.Context, countryCode, phone string) (*userphonepb.UserPhone, error) {
	req := &userphonepb.GetUserPhoneByCountryCodeAndPhoneRequest{
		CountryCode: countryCode,
		Phone:       phone,
	}

	resp, err := c.client.GetUserPhoneByCountryCodeAndPhone(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.UserPhone, nil
}