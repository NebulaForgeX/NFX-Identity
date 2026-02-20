package mapper

import (
	"encoding/json"
	"nfxid/enums"
	"nfxid/modules/audit/domain/events"
	"nfxid/modules/audit/infrastructure/rdb/models"

	"gorm.io/datatypes"
)

// EventDomainToModel 将 Domain Event 转换为 Model Event
func EventDomainToModel(e *events.Event) *models.Event {
	if e == nil {
		return nil
	}

	var metadata *datatypes.JSON
	if e.Metadata() != nil && len(e.Metadata()) > 0 {
		metaBytes, _ := json.Marshal(e.Metadata())
		jsonData := datatypes.JSON(metaBytes)
		metadata = &jsonData
	}

	return &models.Event{
		ID:                  e.ID(),
		EventID:             e.EventID(),
		OccurredAt:          e.OccurredAt(),
		ReceivedAt:          e.ReceivedAt(),
		TenantID:            e.TenantID(),
		ApplicationID:       e.AppID(),
		ActorType:           enums.AuditActorType(e.ActorType()),
		ActorID:             e.ActorID(),
		ActorTenantMemberID: e.ActorTenantMemberID(),
		Action:              e.Action(),
		TargetType:          e.TargetType(),
		TargetID:            e.TargetID(),
		Result:              enums.AuditResultType(e.Result()),
		FailureReasonCode:   e.FailureReasonCode(),
		HttpMethod:          e.HTTPMethod(),
		HttpPath:            e.HTTPPath(),
		HttpStatus:          e.HTTPStatus(),
		RequestID:           e.RequestID(),
		TraceID:             e.TraceID(),
		IP:                  e.IP(),
		UserAgent:           e.UserAgent(),
		GeoCountry:          e.GeoCountry(),
		RiskLevel:           enums.AuditRiskLevel(e.RiskLevel()),
		DataClassification:  enums.AuditDataClassification(e.DataClassification()),
		PrevHash:            e.PrevHash(),
		EventHash:           e.EventHash(),
		Metadata:            metadata,
		CreatedAt:           e.CreatedAt(),
	}
}

// EventModelToDomain 将 Model Event 转换为 Domain Event
func EventModelToDomain(m *models.Event) *events.Event {
	if m == nil {
		return nil
	}

	var metadata map[string]interface{}
	if m.Metadata != nil {
		json.Unmarshal(*m.Metadata, &metadata)
	}

	state := events.EventState{
		ID:                   m.ID,
		EventID:              m.EventID,
		OccurredAt:           m.OccurredAt,
		ReceivedAt:           m.ReceivedAt,
		TenantID:             m.TenantID,
		AppID:                m.ApplicationID,
		ActorType:            events.ActorType(m.ActorType),
		ActorID:              m.ActorID,
		ActorTenantMemberID:  m.ActorTenantMemberID,
		Action:               m.Action,
		TargetType:           m.TargetType,
		TargetID:             m.TargetID,
		Result:               events.ResultType(m.Result),
		FailureReasonCode:    m.FailureReasonCode,
		HTTPMethod:           m.HttpMethod,
		HTTPPath:             m.HttpPath,
		HTTPStatus:           m.HttpStatus,
		RequestID:            m.RequestID,
		TraceID:              m.TraceID,
		IP:                   m.IP,
		UserAgent:            m.UserAgent,
		GeoCountry:           m.GeoCountry,
		RiskLevel:            events.RiskLevel(m.RiskLevel),
		DataClassification:   events.DataClassification(m.DataClassification),
		PrevHash:             m.PrevHash,
		EventHash:            m.EventHash,
		Metadata:             metadata,
		CreatedAt:            m.CreatedAt,
	}

	return events.NewEventFromState(state)
}

// EventModelToUpdates 将 Model Event 转换为更新字段映射
func EventModelToUpdates(m *models.Event) map[string]any {
	var metadata any
	if m.Metadata != nil {
		metadata = m.Metadata
	}

	return map[string]any{
		models.EventCols.EventID:             m.EventID,
		models.EventCols.OccurredAt:          m.OccurredAt,
		models.EventCols.ReceivedAt:          m.ReceivedAt,
		models.EventCols.TenantID:            m.TenantID,
		models.EventCols.ApplicationID:        m.ApplicationID,
		models.EventCols.ActorType:           m.ActorType,
		models.EventCols.ActorID:             m.ActorID,
		models.EventCols.ActorTenantMemberID: m.ActorTenantMemberID,
		models.EventCols.Action:              m.Action,
		models.EventCols.TargetType:          m.TargetType,
		models.EventCols.TargetID:           m.TargetID,
		models.EventCols.Result:             m.Result,
		models.EventCols.FailureReasonCode:  m.FailureReasonCode,
		models.EventCols.HttpMethod:         m.HttpMethod,
		models.EventCols.HttpPath:           m.HttpPath,
		models.EventCols.HttpStatus:         m.HttpStatus,
		models.EventCols.RequestID:          m.RequestID,
		models.EventCols.TraceID:            m.TraceID,
		models.EventCols.IP:                 m.IP,
		models.EventCols.UserAgent:          m.UserAgent,
		models.EventCols.GeoCountry:         m.GeoCountry,
		models.EventCols.RiskLevel:          m.RiskLevel,
		models.EventCols.DataClassification: m.DataClassification,
		models.EventCols.PrevHash:           m.PrevHash,
		models.EventCols.EventHash:          m.EventHash,
		models.EventCols.Metadata:           metadata,
	}
}
