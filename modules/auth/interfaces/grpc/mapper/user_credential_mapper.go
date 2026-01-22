package mapper

import (
	userCredentialAppResult "nfxid/modules/auth/application/user_credentials/results"
	userCredentialDomain "nfxid/modules/auth/domain/user_credentials"
	usercredentialpb "nfxid/protos/gen/auth/user_credential"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserCredentialROToProto 将 UserCredentialRO 转换为 proto UserCredential 消息
func UserCredentialROToProto(v *userCredentialAppResult.UserCredentialRO) *usercredentialpb.UserCredential {
	if v == nil {
		return nil
	}

	userCredential := &usercredentialpb.UserCredential{
		Id:                 v.ID.String(), // id 直接引用 directory.users.id
		TenantId:           v.TenantID.String(),
		CredentialType:     credentialTypeToProto(v.CredentialType),
		Status:             credentialStatusToProto(v.Status),
		MustChangePassword: v.MustChangePassword,
		Version:            int32(v.Version),
		CreatedAt:          timestamppb.New(v.CreatedAt),
		UpdatedAt:          timestamppb.New(v.UpdatedAt),
	}

	if v.PasswordHash != nil {
		hashVal := *v.PasswordHash
		userCredential.PasswordHash = &hashVal
	}
	if v.HashAlg != nil {
		algVal := *v.HashAlg
		userCredential.HashAlg = &algVal
	}
	if v.HashParams != nil {
		if structVal, err := structpb.NewStruct(v.HashParams); err == nil {
			userCredential.HashParams = structVal
		}
	}
	if v.PasswordUpdatedAt != nil {
		userCredential.PasswordUpdatedAt = timestamppb.New(*v.PasswordUpdatedAt)
	}
	if v.LastSuccessLoginAt != nil {
		userCredential.LastSuccessLoginAt = timestamppb.New(*v.LastSuccessLoginAt)
	}
	if v.DeletedAt != nil {
		userCredential.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return userCredential
}

// UserCredentialListROToProto 批量转换 UserCredentialRO 到 proto UserCredential
func UserCredentialListROToProto(results []userCredentialAppResult.UserCredentialRO) []*usercredentialpb.UserCredential {
	userCredentials := make([]*usercredentialpb.UserCredential, len(results))
	for i, v := range results {
		userCredentials[i] = UserCredentialROToProto(&v)
	}
	return userCredentials
}

func credentialTypeToProto(ct userCredentialDomain.CredentialType) usercredentialpb.AuthCredentialType {
	switch ct {
	case userCredentialDomain.CredentialTypePassword:
		return usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_PASSWORD
	case userCredentialDomain.CredentialTypePasskey:
		return usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_PASSKEY
	case userCredentialDomain.CredentialTypeOauthLink:
		return usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_OAUTH_LINK
	case userCredentialDomain.CredentialTypeSaml:
		return usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_SAML
	case userCredentialDomain.CredentialTypeLdap:
		return usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_LDAP
	default:
		return usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_UNSPECIFIED
	}
}

func credentialStatusToProto(cs userCredentialDomain.CredentialStatus) usercredentialpb.AuthCredentialStatus {
	switch cs {
	case userCredentialDomain.CredentialStatusActive:
		return usercredentialpb.AuthCredentialStatus_AUTH_CREDENTIAL_STATUS_ACTIVE
	case userCredentialDomain.CredentialStatusDisabled:
		return usercredentialpb.AuthCredentialStatus_AUTH_CREDENTIAL_STATUS_DISABLED
	case userCredentialDomain.CredentialStatusExpired:
		return usercredentialpb.AuthCredentialStatus_AUTH_CREDENTIAL_STATUS_EXPIRED
	default:
		return usercredentialpb.AuthCredentialStatus_AUTH_CREDENTIAL_STATUS_UNSPECIFIED
	}
}
