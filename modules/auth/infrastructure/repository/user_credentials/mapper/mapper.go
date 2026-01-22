package mapper

import (
	"encoding/json"
	"nfxid/enums"
	"nfxid/modules/auth/domain/user_credentials"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"

	"gorm.io/datatypes"
)

// UserCredentialDomainToModel 将 Domain UserCredential 转换为 Model UserCredential
func UserCredentialDomainToModel(uc *user_credentials.UserCredential) *models.UserCredential {
	if uc == nil {
		return nil
	}

	var hashParams *datatypes.JSON
	if uc.HashParams() != nil && len(uc.HashParams()) > 0 {
		paramsBytes, _ := json.Marshal(uc.HashParams())
		jsonData := datatypes.JSON(paramsBytes)
		hashParams = &jsonData
	}

	return &models.UserCredential{
		ID:                 uc.ID(), // id 直接引用 directory.users.id
		TenantID:           uc.TenantID(),
		CredentialType:     credentialTypeDomainToEnum(uc.CredentialType()),
		PasswordHash:       uc.PasswordHash(),
		HashAlg:            uc.HashAlg(),
		HashParams:         hashParams,
		PasswordUpdatedAt:  uc.PasswordUpdatedAt(),
		LastSuccessLoginAt: uc.LastSuccessLoginAt(),
		Status:             credentialStatusDomainToEnum(uc.Status()),
		MustChangePassword: uc.MustChangePassword(),
		Version:            uc.Version(),
		CreatedAt:          uc.CreatedAt(),
		UpdatedAt:          uc.UpdatedAt(),
		DeletedAt:          timex.TimeToGormDeletedAt(uc.DeletedAt()),
	}
}

// UserCredentialModelToDomain 将 Model UserCredential 转换为 Domain UserCredential
func UserCredentialModelToDomain(m *models.UserCredential) *user_credentials.UserCredential {
	if m == nil {
		return nil
	}

	var hashParams map[string]interface{}
	if m.HashParams != nil {
		json.Unmarshal(*m.HashParams, &hashParams)
	}

	state := user_credentials.UserCredentialState{
		ID:                 m.ID, // id 直接引用 directory.users.id
		UserID:             m.ID, // UserID 从 ID 获取（一对一关系）
		TenantID:           m.TenantID,
		CredentialType:     credentialTypeEnumToDomain(m.CredentialType),
		PasswordHash:       m.PasswordHash,
		HashAlg:            m.HashAlg,
		HashParams:         hashParams,
		PasswordUpdatedAt:  m.PasswordUpdatedAt,
		LastSuccessLoginAt: m.LastSuccessLoginAt,
		Status:             credentialStatusEnumToDomain(m.Status),
		MustChangePassword: m.MustChangePassword,
		Version:            m.Version,
		CreatedAt:          m.CreatedAt,
		UpdatedAt:          m.UpdatedAt,
		DeletedAt:          timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return user_credentials.NewUserCredentialFromState(state)
}

// UserCredentialModelToUpdates 将 Model UserCredential 转换为更新字段映射
func UserCredentialModelToUpdates(m *models.UserCredential) map[string]any {
	var hashParams any
	if m.HashParams != nil {
		hashParams = m.HashParams
	}

	return map[string]any{
		models.UserCredentialCols.TenantID:           m.TenantID,
		models.UserCredentialCols.CredentialType:     m.CredentialType,
		models.UserCredentialCols.PasswordHash:       m.PasswordHash,
		models.UserCredentialCols.HashAlg:            m.HashAlg,
		models.UserCredentialCols.HashParams:         hashParams,
		models.UserCredentialCols.PasswordUpdatedAt:  m.PasswordUpdatedAt,
		models.UserCredentialCols.LastSuccessLoginAt: m.LastSuccessLoginAt,
		models.UserCredentialCols.Status:             m.Status,
		models.UserCredentialCols.MustChangePassword: m.MustChangePassword,
		models.UserCredentialCols.Version:            m.Version,
		models.UserCredentialCols.UpdatedAt:          m.UpdatedAt,
		models.UserCredentialCols.DeletedAt:          m.DeletedAt,
	}
}

// 枚举转换辅助函数

func credentialTypeDomainToEnum(ct user_credentials.CredentialType) enums.AuthCredentialType {
	switch ct {
	case user_credentials.CredentialTypePassword:
		return enums.AuthCredentialTypePassword
	case user_credentials.CredentialTypePasskey:
		return enums.AuthCredentialTypePasskey
	case user_credentials.CredentialTypeOauthLink:
		return enums.AuthCredentialTypeOauthLink
	case user_credentials.CredentialTypeSaml:
		return enums.AuthCredentialTypeSaml
	case user_credentials.CredentialTypeLdap:
		return enums.AuthCredentialTypeLdap
	default:
		return enums.AuthCredentialTypePassword
	}
}

func credentialTypeEnumToDomain(ct enums.AuthCredentialType) user_credentials.CredentialType {
	switch ct {
	case enums.AuthCredentialTypePassword:
		return user_credentials.CredentialTypePassword
	case enums.AuthCredentialTypePasskey:
		return user_credentials.CredentialTypePasskey
	case enums.AuthCredentialTypeOauthLink:
		return user_credentials.CredentialTypeOauthLink
	case enums.AuthCredentialTypeSaml:
		return user_credentials.CredentialTypeSaml
	case enums.AuthCredentialTypeLdap:
		return user_credentials.CredentialTypeLdap
	default:
		return user_credentials.CredentialTypePassword
	}
}

// CredentialStatusDomainToEnum 将 Domain CredentialStatus 转换为 Enum CredentialStatus
func CredentialStatusDomainToEnum(cs user_credentials.CredentialStatus) enums.AuthCredentialStatus {
	return credentialStatusDomainToEnum(cs)
}

func credentialStatusDomainToEnum(cs user_credentials.CredentialStatus) enums.AuthCredentialStatus {
	switch cs {
	case user_credentials.CredentialStatusActive:
		return enums.AuthCredentialStatusActive
	case user_credentials.CredentialStatusDisabled:
		return enums.AuthCredentialStatusDisabled
	case user_credentials.CredentialStatusExpired:
		return enums.AuthCredentialStatusExpired
	default:
		return enums.AuthCredentialStatusActive
	}
}

func credentialStatusEnumToDomain(cs enums.AuthCredentialStatus) user_credentials.CredentialStatus {
	switch cs {
	case enums.AuthCredentialStatusActive:
		return user_credentials.CredentialStatusActive
	case enums.AuthCredentialStatusDisabled:
		return user_credentials.CredentialStatusDisabled
	case enums.AuthCredentialStatusExpired:
		return user_credentials.CredentialStatusExpired
	default:
		return user_credentials.CredentialStatusActive
	}
}
