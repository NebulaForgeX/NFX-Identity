package members

import (
	tenantsErr "nfxid/errors/src/tenants"

	"github.com/google/uuid"
)

func (m *Member) Validate() error {
	if m.TenantID() == uuid.Nil {
		return tenantsErr.ErrTenantIDRequired
	}
	if m.UserID() == uuid.Nil {
		return tenantsErr.ErrUserIDRequired
	}
	validStatuses := map[MemberStatus]struct{}{
		MemberStatusInvited:   {},
		MemberStatusActive:    {},
		MemberStatusSuspended: {},
		MemberStatusRemoved:   {},
	}
	if _, ok := validStatuses[m.Status()]; !ok {
		return tenantsErr.ErrInvalidMemberStatus
	}
	validSources := map[MemberSource]struct{}{
		MemberSourceManual: {},
		MemberSourceInvite: {},
		MemberSourceSCIM:   {},
		MemberSourceSSO:    {},
		MemberSourceHRSync: {},
		MemberSourceImport: {},
	}
	if _, ok := validSources[m.Source()]; !ok {
		return tenantsErr.ErrInvalidMemberSource
	}
	return nil
}
