package members

import "github.com/google/uuid"

func (m *Member) Validate() error {
	if m.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	if m.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	validStatuses := map[MemberStatus]struct{}{
		MemberStatusInvited:   {},
		MemberStatusActive:    {},
		MemberStatusSuspended: {},
		MemberStatusRemoved:   {},
	}
	if _, ok := validStatuses[m.Status()]; !ok {
		return ErrInvalidMemberStatus
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
		return ErrInvalidMemberSource
	}
	return nil
}
