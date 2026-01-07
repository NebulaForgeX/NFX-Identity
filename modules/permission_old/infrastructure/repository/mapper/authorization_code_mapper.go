package mapper

import (
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

func AuthorizationCodeDomainToModel(ac *authorizationCodeDomain.AuthorizationCode) *models.AuthorizationCode {
	if ac == nil {
		return nil
	}

	editable := ac.Editable()

	return &models.AuthorizationCode{
		ID:        ac.ID(),
		Code:      editable.Code,
		MaxUses:   editable.MaxUses,
		UsedCount: editable.UsedCount,
		CreatedBy: ac.CreatedBy(),
		ExpiresAt: ac.ExpiresAt(),
		IsActive:  ac.IsActive(),
		CreatedAt: ac.CreatedAt(),
		UpdatedAt: ac.UpdatedAt(),
		DeletedAt: timex.TimeToGormDeletedAt(ac.DeletedAt()),
	}
}

func AuthorizationCodeModelToDomain(m *models.AuthorizationCode) *authorizationCodeDomain.AuthorizationCode {
	if m == nil {
		return nil
	}

	editable := authorizationCodeDomain.AuthorizationCodeEditable{
		Code:      m.Code,
		MaxUses:   m.MaxUses,
		UsedCount: m.UsedCount,
	}

	state := authorizationCodeDomain.AuthorizationCodeState{
		ID:        m.ID,
		Editable:  editable,
		CreatedBy: m.CreatedBy,
		ExpiresAt: m.ExpiresAt,
		IsActive:  m.IsActive,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return authorizationCodeDomain.NewAuthorizationCodeFromState(state)
}

func AuthorizationCodeModelsToUpdates(m *models.AuthorizationCode) map[string]any {
	return map[string]any{
		models.AuthorizationCodeCols.Code:      m.Code,
		models.AuthorizationCodeCols.MaxUses:   m.MaxUses,
		models.AuthorizationCodeCols.UsedCount: m.UsedCount,
		models.AuthorizationCodeCols.IsActive:  m.IsActive,
		models.AuthorizationCodeCols.ExpiresAt: m.ExpiresAt,
		models.AuthorizationCodeCols.DeletedAt: m.DeletedAt,
	}
}
