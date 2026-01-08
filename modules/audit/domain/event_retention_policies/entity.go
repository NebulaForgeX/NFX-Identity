package event_retention_policies

import (
	"time"

	"github.com/google/uuid"
)

type RetentionAction string

const (
	RetentionActionArchive RetentionAction = "archive"
	RetentionActionDelete  RetentionAction = "delete"
	RetentionActionExport  RetentionAction = "export"
)

type DataClassification string

const (
	DataClassificationPublic        DataClassification = "public"
	DataClassificationInternal      DataClassification = "internal"
	DataClassificationConfidential  DataClassification = "confidential"
	DataClassificationRestricted    DataClassification = "restricted"
)

type RiskLevel string

const (
	RiskLevelLow      RiskLevel = "low"
	RiskLevelMedium   RiskLevel = "medium"
	RiskLevelHigh     RiskLevel = "high"
	RiskLevelCritical RiskLevel = "critical"
)

type EventRetentionPolicy struct {
	state EventRetentionPolicyState
}

type EventRetentionPolicyState struct {
	ID                uuid.UUID
	PolicyName        string
	TenantID          *uuid.UUID
	ActionPattern     *string
	DataClassification *DataClassification
	RiskLevel         *RiskLevel
	RetentionDays     int
	RetentionAction   RetentionAction
	ArchiveLocation   *string
	Status            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CreatedBy         *uuid.UUID
}

func (erp *EventRetentionPolicy) ID() uuid.UUID                    { return erp.state.ID }
func (erp *EventRetentionPolicy) PolicyName() string               { return erp.state.PolicyName }
func (erp *EventRetentionPolicy) TenantID() *uuid.UUID             { return erp.state.TenantID }
func (erp *EventRetentionPolicy) ActionPattern() *string           { return erp.state.ActionPattern }
func (erp *EventRetentionPolicy) DataClassification() *DataClassification { return erp.state.DataClassification }
func (erp *EventRetentionPolicy) RiskLevel() *RiskLevel            { return erp.state.RiskLevel }
func (erp *EventRetentionPolicy) RetentionDays() int               { return erp.state.RetentionDays }
func (erp *EventRetentionPolicy) RetentionAction() RetentionAction { return erp.state.RetentionAction }
func (erp *EventRetentionPolicy) ArchiveLocation() *string         { return erp.state.ArchiveLocation }
func (erp *EventRetentionPolicy) Status() string                   { return erp.state.Status }
func (erp *EventRetentionPolicy) CreatedAt() time.Time             { return erp.state.CreatedAt }
func (erp *EventRetentionPolicy) UpdatedAt() time.Time             { return erp.state.UpdatedAt }
func (erp *EventRetentionPolicy) CreatedBy() *uuid.UUID            { return erp.state.CreatedBy }
