package mapper

import (
	phoneDomain "nfxidentity/modules/auth/domain/phone"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/pkgs/ptrx"
)

func PhoneDomainToModel(p *phoneDomain.Phone) *rdbmodels.Phone {
	if p == nil {
		return nil
	}
	return &rdbmodels.Phone{
		ID:         p.ID(),
		AccountID:  p.AccountID(),
		Phone:      p.Phone(),
		IsPrimary:  p.IsPrimary(),
		VerifiedAt: p.VerifiedAt(),
		CreatedAt:  p.CreatedAt(),
		UpdatedAt:  p.UpdatedAt(),
		DeletedAt:  ptrx.TimePtrToDeletedAt(p.DeletedAt()),
	}
}

func PhoneModelToDomain(m *rdbmodels.Phone) *phoneDomain.Phone {
	if m == nil {
		return nil
	}
	return phoneDomain.NewPhoneFromState(phoneDomain.PhoneState{
		ID:         m.ID,
		AccountID:  m.AccountID,
		Phone:      m.Phone,
		IsPrimary:  m.IsPrimary,
		VerifiedAt: m.VerifiedAt,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		DeletedAt:  ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}

func PhoneDomainToUpdates(p *phoneDomain.Phone) map[string]any {
	m := PhoneDomainToModel(p)
	return map[string]any{
		rdbmodels.PhoneCols.AccountID:  m.AccountID,
		rdbmodels.PhoneCols.Phone:      m.Phone,
		rdbmodels.PhoneCols.IsPrimary:  m.IsPrimary,
		rdbmodels.PhoneCols.VerifiedAt: m.VerifiedAt,
		rdbmodels.PhoneCols.UpdatedAt:  m.UpdatedAt,
		rdbmodels.PhoneCols.DeletedAt:  m.DeletedAt,
	}
}
