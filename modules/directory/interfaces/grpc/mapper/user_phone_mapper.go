package mapper

import (
	userPhoneAppResult "nfxid/modules/directory/application/user_phones/results"
	userphonepb "nfxid/protos/gen/directory/user_phone"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserPhoneROToProto 将 UserPhoneRO 转换为 proto UserPhone 消息
func UserPhoneROToProto(v *userPhoneAppResult.UserPhoneRO) *userphonepb.UserPhone {
	if v == nil {
		return nil
	}

	userPhone := &userphonepb.UserPhone{
		Id:        v.ID.String(),
		UserId:    v.UserID.String(),
		Phone:     v.Phone,
		IsPrimary: v.IsPrimary,
		IsVerified: v.IsVerified,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.CountryCode != nil {
		userPhone.CountryCode = v.CountryCode
	}

	if v.VerifiedAt != nil {
		userPhone.VerifiedAt = timestamppb.New(*v.VerifiedAt)
	}

	if v.VerificationCode != nil {
		userPhone.VerificationCode = v.VerificationCode
	}

	if v.VerificationExpiresAt != nil {
		userPhone.VerificationExpiresAt = timestamppb.New(*v.VerificationExpiresAt)
	}

	if v.DeletedAt != nil {
		userPhone.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return userPhone
}

// UserPhoneListROToProto 批量转换 UserPhoneRO 到 proto UserPhone
func UserPhoneListROToProto(results []userPhoneAppResult.UserPhoneRO) []*userphonepb.UserPhone {
	userPhones := make([]*userphonepb.UserPhone, len(results))
	for i, v := range results {
		userPhones[i] = UserPhoneROToProto(&v)
	}
	return userPhones
}
