package mapper

import (
	emailDomain "nfxidentity/modules/auth/domain/email"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/pkgs/ptrx"
)

func EmailDomainToModel(e *emailDomain.Email) *rdbmodels.Email {
	if e == nil {
		return nil
	}
	return &rdbmodels.Email{
		ID:         e.ID(),
		AccountID:  e.AccountID(),
		Email:      e.Email(),
		IsPrimary:  e.IsPrimary(),
		VerifiedAt: e.VerifiedAt(),
		CreatedAt:  e.CreatedAt(),
		UpdatedAt:  e.UpdatedAt(),
		DeletedAt:  ptrx.TimePtrToDeletedAt(e.DeletedAt()),
	}
}

func EmailModelToDomain(m *rdbmodels.Email) *emailDomain.Email {
	if m == nil {
		return nil
	}
	return emailDomain.NewEmailFromState(emailDomain.EmailState{
		ID:         m.ID,
		AccountID:  m.AccountID,
		Email:      m.Email,
		IsPrimary:  m.IsPrimary,
		VerifiedAt: m.VerifiedAt,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		DeletedAt:  ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}

func EmailModelsToDomains(ms []rdbmodels.Email) []*emailDomain.Email {
	out := make([]*emailDomain.Email, 0, len(ms))
	for i := range ms {
		out = append(out, EmailModelToDomain(&ms[i]))
	}
	return out
}

func EmailDomainToUpdates(e *emailDomain.Email) map[string]any {
	m := EmailDomainToModel(e)
	return map[string]any{
		rdbmodels.EmailCols.AccountID:  m.AccountID,
		rdbmodels.EmailCols.Email:      m.Email,
		rdbmodels.EmailCols.IsPrimary:  m.IsPrimary,
		rdbmodels.EmailCols.VerifiedAt: m.VerifiedAt,
		rdbmodels.EmailCols.UpdatedAt:  m.UpdatedAt,
		rdbmodels.EmailCols.DeletedAt:  m.DeletedAt,
	}
}
