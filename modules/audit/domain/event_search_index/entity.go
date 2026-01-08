package event_search_index

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

type EventSearchIndex struct {
	state EventSearchIndexState
}

type EventSearchIndexState struct {
	ID                uuid.UUID
	EventID           string
	TenantID          *uuid.UUID
	AppID             *uuid.UUID
	ActorType         ActorType
	ActorID           uuid.UUID
	Action            string
	TargetType        *string
	TargetID          *uuid.UUID
	Result            ResultType
	OccurredAt        time.Time
	IP                *string
	RiskLevel         RiskLevel
	DataClassification DataClassification
	Tags              []string
	CreatedAt         time.Time
}

func (esi *EventSearchIndex) ID() uuid.UUID                   { return esi.state.ID }
func (esi *EventSearchIndex) EventID() string                 { return esi.state.EventID }
func (esi *EventSearchIndex) TenantID() *uuid.UUID            { return esi.state.TenantID }
func (esi *EventSearchIndex) AppID() *uuid.UUID               { return esi.state.AppID }
func (esi *EventSearchIndex) ActorType() ActorType            { return esi.state.ActorType }
func (esi *EventSearchIndex) ActorID() uuid.UUID              { return esi.state.ActorID }
func (esi *EventSearchIndex) Action() string                  { return esi.state.Action }
func (esi *EventSearchIndex) TargetType() *string             { return esi.state.TargetType }
func (esi *EventSearchIndex) TargetID() *uuid.UUID            { return esi.state.TargetID }
func (esi *EventSearchIndex) Result() ResultType              { return esi.state.Result }
func (esi *EventSearchIndex) OccurredAt() time.Time           { return esi.state.OccurredAt }
func (esi *EventSearchIndex) IP() *string                     { return esi.state.IP }
func (esi *EventSearchIndex) RiskLevel() RiskLevel            { return esi.state.RiskLevel }
func (esi *EventSearchIndex) DataClassification() DataClassification { return esi.state.DataClassification }
func (esi *EventSearchIndex) Tags() []string                  { return esi.state.Tags }
func (esi *EventSearchIndex) CreatedAt() time.Time            { return esi.state.CreatedAt }
