package dto

import (
	useremailpb "nfxid/protos/gen/directory/user_email"
)

// CreateUserEmailDTO 创建用户邮箱的 DTO
type CreateUserEmailDTO struct {
	UserID            string
	Email             string
	IsPrimary         bool
	IsVerified        bool
	VerificationToken *string
}

// ToCreateUserEmailRequest 转换为 protobuf 请求
func (d *CreateUserEmailDTO) ToCreateUserEmailRequest() *useremailpb.CreateUserEmailRequest {
	return &useremailpb.CreateUserEmailRequest{
		UserId:            d.UserID,
		Email:             d.Email,
		IsPrimary:         d.IsPrimary,
		IsVerified:         d.IsVerified,
		VerificationToken: d.VerificationToken,
	}
}
