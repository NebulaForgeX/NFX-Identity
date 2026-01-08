package members

import (
	"time"
)

func (m *Member) UpdateStatus(status MemberStatus) error {
	validStatuses := map[MemberStatus]struct{}{
		MemberStatusInvited:   {},
		MemberStatusActive:    {},
		MemberStatusSuspended: {},
		MemberStatusRemoved:   {},
	}
	if _, ok := validStatuses[status]; !ok {
		return ErrInvalidMemberStatus
	}

	m.state.Status = status
	m.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (m *Member) Join() error {
	if m.Status() == MemberStatusActive {
		return nil // already joined
	}

	now := time.Now().UTC()
	m.state.Status = MemberStatusActive
	m.state.JoinedAt = &now
	m.state.LeftAt = nil
	m.state.UpdatedAt = now
	return nil
}

func (m *Member) Leave() error {
	if m.Status() == MemberStatusRemoved {
		return nil // already left
	}

	now := time.Now().UTC()
	m.state.Status = MemberStatusRemoved
	m.state.LeftAt = &now
	m.state.UpdatedAt = now
	return nil
}

func (m *Member) Suspend() error {
	m.state.Status = MemberStatusSuspended
	m.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (m *Member) UpdateExternalRef(externalRef *string) error {
	m.state.ExternalRef = externalRef
	m.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (m *Member) UpdateMetadata(metadata map[string]interface{}) error {
	if metadata == nil {
		return nil
	}
	m.state.Metadata = metadata
	m.state.UpdatedAt = time.Now().UTC()
	return nil
}
