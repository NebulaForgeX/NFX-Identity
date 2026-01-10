package mapper

import (
	auditAppResult "nfxid/modules/audit/application/events/results"
	auditDomain "nfxid/modules/audit/domain/events"
	eventpb "nfxid/protos/gen/audit/event"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// EventROToProto 将 EventRO 转换为 proto Event 消息
func EventROToProto(v *auditAppResult.EventRO) *eventpb.Event {
	if v == nil {
		return nil
	}

	event := &eventpb.Event{
		Id:          v.ID.String(),
		EventId:     v.EventID,
		OccurredAt:  timestamppb.New(v.OccurredAt),
		ReceivedAt:  timestamppb.New(v.ReceivedAt),
		ActorType:   actorTypeToProto(v.ActorType),
		ActorId:     v.ActorID.String(),
		Action:      v.Action,
		Result:      resultTypeToProto(v.Result),
		RiskLevel:   riskLevelToProto(v.RiskLevel),
		DataClassification: dataClassificationToProto(v.DataClassification),
		CreatedAt:   timestamppb.New(v.CreatedAt),
	}

	if v.TenantID != nil {
		tenantIDStr := v.TenantID.String()
		event.TenantId = &tenantIDStr
	}
	if v.AppID != nil {
		appIDStr := v.AppID.String()
		event.AppId = &appIDStr
	}
	if v.ActorTenantMemberID != nil {
		actorTenantMemberIDStr := v.ActorTenantMemberID.String()
		event.ActorTenantMemberId = &actorTenantMemberIDStr
	}
	if v.TargetType != nil {
		event.TargetType = v.TargetType
	}
	if v.TargetID != nil {
		targetIDStr := v.TargetID.String()
		event.TargetId = &targetIDStr
	}
	if v.FailureReasonCode != nil {
		event.FailureReasonCode = v.FailureReasonCode
	}
	if v.HTTPMethod != nil {
		event.HttpMethod = v.HTTPMethod
	}
	if v.HTTPPath != nil {
		event.HttpPath = v.HTTPPath
	}
	if v.HTTPStatus != nil {
		httpStatus := int32(*v.HTTPStatus)
		event.HttpStatus = &httpStatus
	}
	if v.RequestID != nil {
		event.RequestId = v.RequestID
	}
	if v.TraceID != nil {
		event.TraceId = v.TraceID
	}
	if v.IP != nil {
		event.Ip = v.IP
	}
	if v.UserAgent != nil {
		event.UserAgent = v.UserAgent
	}
	if v.GeoCountry != nil {
		event.GeoCountry = v.GeoCountry
	}
	if v.PrevHash != nil {
		event.PrevHash = v.PrevHash
	}
	if v.EventHash != nil {
		event.EventHash = v.EventHash
	}
	if v.Metadata != nil {
		if metadata, err := structpb.NewStruct(v.Metadata); err == nil {
			event.Metadata = metadata
		}
	}

	return event
}

// EventListROToProto 批量转换 EventRO 到 proto Event
func EventListROToProto(results []auditAppResult.EventRO) []*eventpb.Event {
	events := make([]*eventpb.Event, len(results))
	for i, v := range results {
		events[i] = EventROToProto(&v)
	}
	return events
}

func actorTypeToProto(t auditDomain.ActorType) eventpb.AuditActorType {
	switch t {
	case auditDomain.ActorTypeUser:
		return eventpb.AuditActorType_AUDIT_ACTOR_TYPE_USER
	case auditDomain.ActorTypeService:
		return eventpb.AuditActorType_AUDIT_ACTOR_TYPE_SERVICE
	case auditDomain.ActorTypeSystem:
		return eventpb.AuditActorType_AUDIT_ACTOR_TYPE_SYSTEM
	case auditDomain.ActorTypeAdmin:
		return eventpb.AuditActorType_AUDIT_ACTOR_TYPE_ADMIN
	default:
		return eventpb.AuditActorType_AUDIT_ACTOR_TYPE_UNSPECIFIED
	}
}

func resultTypeToProto(r auditDomain.ResultType) eventpb.AuditResultType {
	switch r {
	case auditDomain.ResultTypeSuccess:
		return eventpb.AuditResultType_AUDIT_RESULT_TYPE_SUCCESS
	case auditDomain.ResultTypeFailure:
		return eventpb.AuditResultType_AUDIT_RESULT_TYPE_FAILURE
	case auditDomain.ResultTypeDeny:
		return eventpb.AuditResultType_AUDIT_RESULT_TYPE_DENY
	case auditDomain.ResultTypeError:
		return eventpb.AuditResultType_AUDIT_RESULT_TYPE_ERROR
	default:
		return eventpb.AuditResultType_AUDIT_RESULT_TYPE_UNSPECIFIED
	}
}

func riskLevelToProto(r auditDomain.RiskLevel) eventpb.AuditRiskLevel {
	switch r {
	case auditDomain.RiskLevelLow:
		return eventpb.AuditRiskLevel_AUDIT_RISK_LEVEL_LOW
	case auditDomain.RiskLevelMedium:
		return eventpb.AuditRiskLevel_AUDIT_RISK_LEVEL_MEDIUM
	case auditDomain.RiskLevelHigh:
		return eventpb.AuditRiskLevel_AUDIT_RISK_LEVEL_HIGH
	case auditDomain.RiskLevelCritical:
		return eventpb.AuditRiskLevel_AUDIT_RISK_LEVEL_CRITICAL
	default:
		return eventpb.AuditRiskLevel_AUDIT_RISK_LEVEL_UNSPECIFIED
	}
}

func dataClassificationToProto(d auditDomain.DataClassification) eventpb.AuditDataClassification {
	switch d {
	case auditDomain.DataClassificationPublic:
		return eventpb.AuditDataClassification_AUDIT_DATA_CLASSIFICATION_PUBLIC
	case auditDomain.DataClassificationInternal:
		return eventpb.AuditDataClassification_AUDIT_DATA_CLASSIFICATION_INTERNAL
	case auditDomain.DataClassificationConfidential:
		return eventpb.AuditDataClassification_AUDIT_DATA_CLASSIFICATION_CONFIDENTIAL
	case auditDomain.DataClassificationRestricted:
		return eventpb.AuditDataClassification_AUDIT_DATA_CLASSIFICATION_RESTRICTED
	default:
		return eventpb.AuditDataClassification_AUDIT_DATA_CLASSIFICATION_UNSPECIFIED
	}
}

// ActorTypeFromProto 将 proto ActorType 转换为 domain ActorType
func ActorTypeFromProto(t eventpb.AuditActorType) auditDomain.ActorType {
	switch t {
	case eventpb.AuditActorType_AUDIT_ACTOR_TYPE_USER:
		return auditDomain.ActorTypeUser
	case eventpb.AuditActorType_AUDIT_ACTOR_TYPE_SERVICE:
		return auditDomain.ActorTypeService
	case eventpb.AuditActorType_AUDIT_ACTOR_TYPE_SYSTEM:
		return auditDomain.ActorTypeSystem
	case eventpb.AuditActorType_AUDIT_ACTOR_TYPE_ADMIN:
		return auditDomain.ActorTypeAdmin
	default:
		return auditDomain.ActorTypeUser // Default to user
	}
}
