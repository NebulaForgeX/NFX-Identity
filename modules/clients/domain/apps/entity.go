package apps

import (
	"time"

	"github.com/google/uuid"
)

type AppType string

const (
	AppTypeServer    AppType = "server"
	AppTypeService   AppType = "service"
	AppTypeInternal  AppType = "internal"
	AppTypePartner   AppType = "partner"
	AppTypeThirdParty AppType = "third_party"
)

type AppStatus string

const (
	AppStatusActive    AppStatus = "active"
	AppStatusDisabled  AppStatus = "disabled"
	AppStatusSuspended AppStatus = "suspended"
	AppStatusPending   AppStatus = "pending"
)

type Environment string

const (
	EnvironmentProduction  Environment = "production"
	EnvironmentStaging     Environment = "staging"
	EnvironmentDevelopment Environment = "development"
	EnvironmentTest        Environment = "test"
)

type App struct {
	state AppState
}

type AppState struct {
	ID          uuid.UUID
	AppID       string
	TenantID    uuid.UUID
	Name        string
	Description *string
	Type        AppType
	Status      AppStatus
	Environment Environment
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   *uuid.UUID
	UpdatedBy   *uuid.UUID
	Metadata    map[string]interface{}
	DeletedAt   *time.Time
}

func (a *App) ID() uuid.UUID                    { return a.state.ID }
func (a *App) AppID() string                    { return a.state.AppID }
func (a *App) TenantID() uuid.UUID              { return a.state.TenantID }
func (a *App) Name() string                     { return a.state.Name }
func (a *App) Description() *string             { return a.state.Description }
func (a *App) Type() AppType                    { return a.state.Type }
func (a *App) Status() AppStatus                { return a.state.Status }
func (a *App) Environment() Environment         { return a.state.Environment }
func (a *App) CreatedAt() time.Time             { return a.state.CreatedAt }
func (a *App) UpdatedAt() time.Time             { return a.state.UpdatedAt }
func (a *App) CreatedBy() *uuid.UUID            { return a.state.CreatedBy }
func (a *App) UpdatedBy() *uuid.UUID            { return a.state.UpdatedBy }
func (a *App) Metadata() map[string]interface{} { return a.state.Metadata }
func (a *App) DeletedAt() *time.Time            { return a.state.DeletedAt }
