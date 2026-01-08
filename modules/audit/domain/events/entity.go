package events

import (
	"time"

	"github.com/google/uuid"
)

type ActorType string

const (
	ActorTypeUser    ActorType = "user"
	ActorTypeService ActorType = "service"
	ActorTypeSystem  ActorType = "system"
	ActorTypeAdmin   ActorType = "admin"
)

type ResultType string

const (
	ResultTypeSuccess ResultType = "success"
	ResultTypeFailure ResultType = "failure"
	ResultTypeDeny    ResultType = "deny"
	ResultTypeError   ResultType = "error"
)

type RiskLevel string

const (
	RiskLevelLow      RiskLevel = "low"
	RiskLevelMedium   RiskLevel = "medium"
	RiskLevelHigh     RiskLevel = "high"
	RiskLevelCritical RiskLevel = "critical"
)

type DataClassification string

const (
	DataClassificationPublic        DataClassification = "public"
	DataClassificationInternal      DataClassification = "internal"
	DataClassificationConfidential  DataClassification = "confidential"
	DataClassificationRestricted    DataClassification = "restricted"
)

type Event struct {
	state EventState
}

type EventState struct {
	ID                   uuid.UUID
	EventID              string
	OccurredAt           time.Time
	ReceivedAt           time.Time
	TenantID             *uuid.UUID
	AppID                *uuid.UUID
	ActorType            ActorType
	ActorID              uuid.UUID
	ActorTenantMemberID  *uuid.UUID
	Action               string
	TargetType           *string
	TargetID             *uuid.UUID
	Result               ResultType
	FailureReasonCode    *string
	HTTPMethod           *string
	HTTPPath             *string
	HTTPStatus           *int
	RequestID            *string
	TraceID              *string
	IP                   *string
	UserAgent            *string
	GeoCountry           *string
	RiskLevel            RiskLevel
	DataClassification   DataClassification
	PrevHash             *string
	EventHash            *string
	Metadata             map[string]interface{}
	CreatedAt            time.Time
}

func (e *Event) ID() uuid.UUID                     { return e.state.ID }
func (e *Event) EventID() string                   { return e.state.EventID }
func (e *Event) OccurredAt() time.Time             { return e.state.OccurredAt }
func (e *Event) ReceivedAt() time.Time             { return e.state.ReceivedAt }
func (e *Event) TenantID() *uuid.UUID              { return e.state.TenantID }
func (e *Event) AppID() *uuid.UUID                 { return e.state.AppID }
func (e *Event) ActorType() ActorType              { return e.state.ActorType }
func (e *Event) ActorID() uuid.UUID                { return e.state.ActorID }
func (e *Event) ActorTenantMemberID() *uuid.UUID   { return e.state.ActorTenantMemberID }
func (e *Event) Action() string                    { return e.state.Action }
func (e *Event) TargetType() *string               { return e.state.TargetType }
func (e *Event) TargetID() *uuid.UUID              { return e.state.TargetID }
func (e *Event) Result() ResultType                { return e.state.Result }
func (e *Event) FailureReasonCode() *string        { return e.state.FailureReasonCode }
func (e *Event) HTTPMethod() *string               { return e.state.HTTPMethod }
func (e *Event) HTTPPath() *string                 { return e.state.HTTPPath }
func (e *Event) HTTPStatus() *int                  { return e.state.HTTPStatus }
func (e *Event) RequestID() *string                { return e.state.RequestID }
func (e *Event) TraceID() *string                  { return e.state.TraceID }
func (e *Event) IP() *string                       { return e.state.IP }
func (e *Event) UserAgent() *string                { return e.state.UserAgent }
func (e *Event) GeoCountry() *string               { return e.state.GeoCountry }
func (e *Event) RiskLevel() RiskLevel              { return e.state.RiskLevel }
func (e *Event) DataClassification() DataClassification { return e.state.DataClassification }
func (e *Event) PrevHash() *string                 { return e.state.PrevHash }
func (e *Event) EventHash() *string                { return e.state.EventHash }
func (e *Event) Metadata() map[string]interface{}  { return e.state.Metadata }
func (e *Event) CreatedAt() time.Time              { return e.state.CreatedAt }
