package members

import (
	"time"

	"github.com/google/uuid"
)

type MemberStatus string

const (
	MemberStatusInvited   MemberStatus = "INVITED"
	MemberStatusActive    MemberStatus = "ACTIVE"
	MemberStatusSuspended MemberStatus = "SUSPENDED"
	MemberStatusRemoved   MemberStatus = "REMOVED"
)

type MemberSource string

const (
	MemberSourceManual  MemberSource = "MANUAL"
	MemberSourceInvite  MemberSource = "INVITE"
	MemberSourceSCIM    MemberSource = "SCIM"
	MemberSourceSSO     MemberSource = "SSO"
	MemberSourceHRSync  MemberSource = "HR_SYNC"
	MemberSourceImport  MemberSource = "IMPORT"
)

type Member struct {
	state MemberState
}

type MemberState struct {
	ID          uuid.UUID
	MemberID    uuid.UUID
	TenantID    uuid.UUID
	UserID      uuid.UUID
	Status      MemberStatus
	Source      MemberSource
	JoinedAt    *time.Time
	LeftAt      *time.Time
	CreatedAt   time.Time
	CreatedBy   *uuid.UUID
	UpdatedAt   time.Time
	ExternalRef *string
	Metadata    map[string]interface{}
}

func (m *Member) ID() uuid.UUID                  { return m.state.ID }
func (m *Member) MemberID() uuid.UUID            { return m.state.MemberID }
func (m *Member) TenantID() uuid.UUID            { return m.state.TenantID }
func (m *Member) UserID() uuid.UUID              { return m.state.UserID }
func (m *Member) Status() MemberStatus           { return m.state.Status }
func (m *Member) Source() MemberSource           { return m.state.Source }
func (m *Member) JoinedAt() *time.Time           { return m.state.JoinedAt }
func (m *Member) LeftAt() *time.Time             { return m.state.LeftAt }
func (m *Member) CreatedAt() time.Time           { return m.state.CreatedAt }
func (m *Member) CreatedBy() *uuid.UUID          { return m.state.CreatedBy }
func (m *Member) UpdatedAt() time.Time           { return m.state.UpdatedAt }
func (m *Member) ExternalRef() *string           { return m.state.ExternalRef }
func (m *Member) Metadata() map[string]interface{} { return m.state.Metadata }
