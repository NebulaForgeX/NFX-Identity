package tenants

import (
	"time"

	"github.com/google/uuid"
)

type TenantStatus string

const (
	TenantStatusActive    TenantStatus = "ACTIVE"
	TenantStatusSuspended TenantStatus = "SUSPENDED"
	TenantStatusClosed    TenantStatus = "CLOSED"
	TenantStatusPending   TenantStatus = "PENDING"
)

type Tenant struct {
	state TenantState
}

type TenantState struct {
	ID           uuid.UUID
	TenantID     string
	Name         string
	DisplayName  *string
	Status       TenantStatus
	PrimaryDomain *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Metadata     map[string]interface{}
}

func (t *Tenant) ID() uuid.UUID                    { return t.state.ID }
func (t *Tenant) TenantID() string                 { return t.state.TenantID }
func (t *Tenant) Name() string                     { return t.state.Name }
func (t *Tenant) DisplayName() *string             { return t.state.DisplayName }
func (t *Tenant) Status() TenantStatus             { return t.state.Status }
func (t *Tenant) PrimaryDomain() *string           { return t.state.PrimaryDomain }
func (t *Tenant) CreatedAt() time.Time             { return t.state.CreatedAt }
func (t *Tenant) UpdatedAt() time.Time             { return t.state.UpdatedAt }
func (t *Tenant) DeletedAt() *time.Time            { return t.state.DeletedAt }
func (t *Tenant) Metadata() map[string]interface{} { return t.state.Metadata }
