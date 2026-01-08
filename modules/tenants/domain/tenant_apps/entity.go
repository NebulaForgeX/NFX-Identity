package tenant_apps

import (
	"time"

	"github.com/google/uuid"
)

type TenantAppStatus string

const (
	TenantAppStatusActive    TenantAppStatus = "ACTIVE"
	TenantAppStatusDisabled  TenantAppStatus = "DISABLED"
	TenantAppStatusSuspended TenantAppStatus = "SUSPENDED"
)

type TenantApp struct {
	state TenantAppState
}

type TenantAppState struct {
	ID        uuid.UUID
	TenantID  uuid.UUID
	AppID     uuid.UUID
	Status    TenantAppStatus
	CreatedAt time.Time
	CreatedBy *uuid.UUID
	UpdatedAt time.Time
	Settings  map[string]interface{}
}

func (ta *TenantApp) ID() uuid.UUID                    { return ta.state.ID }
func (ta *TenantApp) TenantID() uuid.UUID               { return ta.state.TenantID }
func (ta *TenantApp) AppID() uuid.UUID                  { return ta.state.AppID }
func (ta *TenantApp) Status() TenantAppStatus           { return ta.state.Status }
func (ta *TenantApp) CreatedAt() time.Time              { return ta.state.CreatedAt }
func (ta *TenantApp) CreatedBy() *uuid.UUID             { return ta.state.CreatedBy }
func (ta *TenantApp) UpdatedAt() time.Time              { return ta.state.UpdatedAt }
func (ta *TenantApp) Settings() map[string]interface{}  { return ta.state.Settings }
