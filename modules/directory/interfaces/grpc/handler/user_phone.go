package handler

import (
	"context"
	"time"

	userPhoneApp "nfxid/modules/directory/application/user_phones"
	userPhoneAppCommands "nfxid/modules/directory/application/user_phones/commands"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	userphonepb "nfxid/protos/gen/directory/user_phone"
	"nfxid/pkgs/errx"

	"github.com/google/uuid"
)

type UserPhoneHandler struct {
	userphonepb.UnimplementedUserPhoneServiceServer
	userPhoneAppSvc *userPhoneApp.Service
}

func NewUserPhoneHandler(userPhoneAppSvc *userPhoneApp.Service) *UserPhoneHandler {
	return &UserPhoneHandler{
		userPhoneAppSvc: userPhoneAppSvc,
	}
}

// CreateUserPhone 创建用户手机
func (h *UserPhoneHandler) CreateUserPhone(ctx context.Context, req *userphonepb.CreateUserPhoneRequest) (*userphonepb.CreateUserPhoneResponse, error) {
	// 解析用户ID
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	// 处理验证码过期时间
	var verificationExpiresAt *string
	if req.VerificationExpiresAt != nil {
		expiresAt := req.VerificationExpiresAt.AsTime().Format(time.RFC3339)
		verificationExpiresAt = &expiresAt
	}

	// 创建命令
	cmd := userPhoneAppCommands.CreateUserPhoneCmd{
		UserID:                userID,
		Phone:                 req.Phone,
		CountryCode:           req.CountryCode,
		IsPrimary:             req.IsPrimary,
		IsVerified:            req.IsVerified,
		VerificationCode:      req.VerificationCode,
		VerificationExpiresAt: verificationExpiresAt,
	}

	// 调用应用服务创建用户手机
	userPhoneID, err := h.userPhoneAppSvc.CreateUserPhone(ctx, cmd)
	if err != nil {
		return nil, err
	}

	// 获取创建的用户手机
	userPhoneView, err := h.userPhoneAppSvc.GetUserPhone(ctx, userPhoneID)
	if err != nil {
		return nil, err
	}

	// 转换为 protobuf 响应
	userPhone := mapper.UserPhoneROToProto(&userPhoneView)
	return &userphonepb.CreateUserPhoneResponse{UserPhone: userPhone}, nil
}

// GetUserPhoneByID 根据ID获取用户手机
func (h *UserPhoneHandler) GetUserPhoneByID(ctx context.Context, req *userphonepb.GetUserPhoneByIDRequest) (*userphonepb.GetUserPhoneByIDResponse, error) {
	userPhoneID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userPhoneView, err := h.userPhoneAppSvc.GetUserPhone(ctx, userPhoneID)
	if err != nil {
		return nil, err
	}

	userPhone := mapper.UserPhoneROToProto(&userPhoneView)
	return &userphonepb.GetUserPhoneByIDResponse{UserPhone: userPhone}, nil
}

// GetUserPhoneByCountryCodeAndPhone 根据国家代码和手机号获取用户手机
func (h *UserPhoneHandler) GetUserPhoneByCountryCodeAndPhone(ctx context.Context, req *userphonepb.GetUserPhoneByCountryCodeAndPhoneRequest) (*userphonepb.GetUserPhoneByCountryCodeAndPhoneResponse, error) {
	userPhoneView, err := h.userPhoneAppSvc.GetUserPhoneByCountryCodeAndPhone(ctx, req.CountryCode, req.Phone)
	if err != nil {
		return nil, err
	}

	userPhone := mapper.UserPhoneROToProto(&userPhoneView)
	return &userphonepb.GetUserPhoneByCountryCodeAndPhoneResponse{UserPhone: userPhone}, nil
}

// GetUserPhonesByUserID 根据用户ID获取用户手机列表
func (h *UserPhoneHandler) GetUserPhonesByUserID(ctx context.Context, req *userphonepb.GetUserPhonesByUserIDRequest) (*userphonepb.GetUserPhonesByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	userPhoneViews, err := h.userPhoneAppSvc.GetUserPhonesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	userPhones := mapper.UserPhoneListROToProto(userPhoneViews)
	return &userphonepb.GetUserPhonesByUserIDResponse{UserPhones: userPhones}, nil
}
