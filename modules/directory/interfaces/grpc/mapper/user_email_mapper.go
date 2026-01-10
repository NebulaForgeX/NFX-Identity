package mapper

import (
	userEmailAppResult "nfxid/modules/directory/application/user_emails/results"
	useremailpb "nfxid/protos/gen/directory/user_email"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserEmailROToProto 将 UserEmailRO 转换为 proto UserEmail 消息
func UserEmailROToProto(v *userEmailAppResult.UserEmailRO) *useremailpb.UserEmail {
	if v == nil {
		return nil
	}

	userEmail := &useremailpb.UserEmail{
		Id:        v.ID.String(),
		UserId:    v.UserID.String(),
		Email:     v.Email,
		IsPrimary: v.IsPrimary,
		IsVerified: v.IsVerified,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.VerifiedAt != nil {
		userEmail.VerifiedAt = timestamppb.New(*v.VerifiedAt)
	}

	if v.VerificationToken != nil {
		userEmail.VerificationToken = v.VerificationToken
	}

	if v.DeletedAt != nil {
		userEmail.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return userEmail
}

// UserEmailListROToProto 批量转换 UserEmailRO 到 proto UserEmail
func UserEmailListROToProto(results []userEmailAppResult.UserEmailRO) []*useremailpb.UserEmail {
	userEmails := make([]*useremailpb.UserEmail, len(results))
	for i, v := range results {
		userEmails[i] = UserEmailROToProto(&v)
	}
	return userEmails
}
