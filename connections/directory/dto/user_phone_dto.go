package dto

import (
	"time"

	userphonepb "nfxid/protos/gen/directory/user_phone"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateUserPhoneDTO 创建用户手机的 DTO
type CreateUserPhoneDTO struct {
	UserID                string
	Phone                 string
	CountryCode           *string
	IsPrimary             bool
	IsVerified             bool
	VerificationCode      *string
	VerificationExpiresAt *time.Time
}

// ToCreateUserPhoneRequest 转换为 protobuf 请求
func (d *CreateUserPhoneDTO) ToCreateUserPhoneRequest() *userphonepb.CreateUserPhoneRequest {
	req := &userphonepb.CreateUserPhoneRequest{
		UserId:          d.UserID,
		Phone:           d.Phone,
		CountryCode:     d.CountryCode,
		IsPrimary:       d.IsPrimary,
		IsVerified:      d.IsVerified,
		VerificationCode: d.VerificationCode,
	}

	if d.VerificationExpiresAt != nil {
		req.VerificationExpiresAt = timestamppb.New(*d.VerificationExpiresAt)
	}

	return req
}
