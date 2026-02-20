package mapper

import (
	"nfxid/modules/access/domain/super_admins"
	"nfxid/modules/access/infrastructure/rdb/models"
)

func SuperAdminDomainToModel(s *super_admins.SuperAdmin) *models.SuperAdmin {
	if s == nil {
		return nil
	}
	return &models.SuperAdmin{UserID: s.UserID(), CreatedAt: s.CreatedAt()}
}

func SuperAdminModelToDomain(m *models.SuperAdmin) *super_admins.SuperAdmin {
	if m == nil {
		return nil
	}
	return super_admins.NewSuperAdminFromState(super_admins.SuperAdminState{UserID: m.UserID, CreatedAt: m.CreatedAt})
}
