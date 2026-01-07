package mapper

import (
	"nfxid/enums"
	"nfxid/modules/auth/domain/user"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

func UserDomainToModel(u *user.User) *models.User {
	if u == nil {
		return nil
	}

	editable := u.Editable()
	return &models.User{
		ID:           u.ID(),
		Username:     editable.Username,
		Email:        editable.Email,
		Phone:        editable.Phone,
		PasswordHash: editable.Password,
		Status:       enums.UserStatus(u.Status()),
		IsVerified:   u.IsVerified(),
		LastLoginAt:  u.LastLoginAt(),
		CreatedAt:    u.CreatedAt(),
		UpdatedAt:    u.UpdatedAt(),
		DeletedAt:    timex.TimeToGormDeletedAt(u.DeletedAt()),
	}
}

func UserModelToDomain(m *models.User) *user.User {
	if m == nil {
		return nil
	}

	editable := user.UserEditable{
		Username: m.Username,
		Email:    m.Email,
		Phone:    m.Phone,
		Password: m.PasswordHash,
	}

	state := user.UserState{
		ID:          m.ID,
		Editable:    editable,
		Status:      string(m.Status),
		IsVerified:  m.IsVerified,
		LastLoginAt: m.LastLoginAt,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return user.NewUserFromState(state)
}

func UserModelsToUpdates(m *models.User) map[string]any {
	return map[string]any{
		models.UserCols.Username:    m.Username,
		models.UserCols.Email:       m.Email,
		models.UserCols.Phone:       m.Phone,
		models.UserCols.Status:      m.Status,
		models.UserCols.IsVerified:  m.IsVerified,
		models.UserCols.LastLoginAt: m.LastLoginAt,
		models.UserCols.DeletedAt:   m.DeletedAt,
	}
}
