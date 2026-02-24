package members

import (
	tenantsErr "nfxid/errors/src/tenants"
	"time"

	"github.com/google/uuid"
)

type NewMemberParams struct {
	TenantID    uuid.UUID
	UserID      uuid.UUID
	Status      MemberStatus
	Source      MemberSource
	CreatedBy   *uuid.UUID
	ExternalRef *string
	Metadata    map[string]interface{}
}

func NewMember(p NewMemberParams) (*Member, error) {
	if err := validateMemberParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	memberID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	status := p.Status
	if status == "" {
		status = MemberStatusInvited
	}

	source := p.Source
	if source == "" {
		source = MemberSourceManual
	}

	now := time.Now().UTC()
	var joinedAt *time.Time
	if status == MemberStatusActive {
		joinedAt = &now
	}

	return NewMemberFromState(MemberState{
		ID:          id,
		MemberID:    memberID,
		TenantID:    p.TenantID,
		UserID:      p.UserID,
		Status:      status,
		Source:      source,
		JoinedAt:    joinedAt,
		CreatedAt:   now,
		CreatedBy:   p.CreatedBy,
		UpdatedAt:   now,
		ExternalRef: p.ExternalRef,
		Metadata:    p.Metadata,
	}), nil
}

func NewMemberFromState(st MemberState) *Member {
	return &Member{state: st}
}

func validateMemberParams(p NewMemberParams) error {
	if p.TenantID == uuid.Nil {
		return tenantsErr.ErrTenantIDRequired
	}
	if p.UserID == uuid.Nil {
		return tenantsErr.ErrUserIDRequired
	}
	if p.Status != "" {
		validStatuses := map[MemberStatus]struct{}{
			MemberStatusInvited:   {},
			MemberStatusActive:    {},
			MemberStatusSuspended: {},
			MemberStatusRemoved:   {},
		}
		if _, ok := validStatuses[p.Status]; !ok {
			return tenantsErr.ErrInvalidMemberStatus
		}
	}
	if p.Source != "" {
		validSources := map[MemberSource]struct{}{
			MemberSourceManual: {},
			MemberSourceInvite: {},
			MemberSourceSCIM:   {},
			MemberSourceSSO:    {},
			MemberSourceHRSync: {},
			MemberSourceImport: {},
		}
		if _, ok := validSources[p.Source]; !ok {
			return tenantsErr.ErrInvalidMemberSource
		}
	}
	return nil
}
