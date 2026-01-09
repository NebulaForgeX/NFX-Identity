package mapper

import (
	"nfxid/enums"
	"nfxid/modules/audit/domain/event_search_index"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"strings"
)

// EventSearchIndexDomainToModel 将 Domain EventSearchIndex 转换为 Model EventSearchIndex
func EventSearchIndexDomainToModel(esi *event_search_index.EventSearchIndex) *models.EventSearchIndex {
	if esi == nil {
		return nil
	}

	var tags *string
	if esi.Tags() != nil && len(esi.Tags()) > 0 {
		tagsStr := "{" + strings.Join(esi.Tags(), ",") + "}"
		tags = &tagsStr
	}

	return &models.EventSearchIndex{
		ID:                 esi.ID(),
		EventID:            esi.EventID(),
		TenantID:           esi.TenantID(),
		AppID:              esi.AppID(),
		ActorType:          enums.AuditActorType(esi.ActorType()),
		ActorID:            esi.ActorID(),
		Action:             esi.Action(),
		TargetType:         esi.TargetType(),
		TargetID:           esi.TargetID(),
		Result:             enums.AuditResultType(esi.Result()),
		OccurredAt:         esi.OccurredAt(),
		IP:                 esi.IP(),
		RiskLevel:          enums.AuditRiskLevel(esi.RiskLevel()),
		DataClassification: enums.AuditDataClassification(esi.DataClassification()),
		Tags:               tags,
		CreatedAt:          esi.CreatedAt(),
	}
}

// EventSearchIndexModelToDomain 将 Model EventSearchIndex 转换为 Domain EventSearchIndex
func EventSearchIndexModelToDomain(m *models.EventSearchIndex) *event_search_index.EventSearchIndex {
	if m == nil {
		return nil
	}

	var tags []string
	if m.Tags != nil && *m.Tags != "" {
		// Remove { and } and split by comma
		tagsStr := strings.Trim(*m.Tags, "{}")
		if tagsStr != "" {
			tags = strings.Split(tagsStr, ",")
		}
	}

	state := event_search_index.EventSearchIndexState{
		ID:                 m.ID,
		EventID:            m.EventID,
		TenantID:           m.TenantID,
		AppID:              m.AppID,
		ActorType:          event_search_index.ActorType(m.ActorType),
		ActorID:            m.ActorID,
		Action:             m.Action,
		TargetType:         m.TargetType,
		TargetID:           m.TargetID,
		Result:             event_search_index.ResultType(m.Result),
		OccurredAt:         m.OccurredAt,
		IP:                 m.IP,
		RiskLevel:          event_search_index.RiskLevel(m.RiskLevel),
		DataClassification: event_search_index.DataClassification(m.DataClassification),
		Tags:               tags,
		CreatedAt:          m.CreatedAt,
	}

	return event_search_index.NewEventSearchIndexFromState(state)
}

// EventSearchIndexModelToUpdates 将 Model EventSearchIndex 转换为更新字段映射
func EventSearchIndexModelToUpdates(m *models.EventSearchIndex) map[string]any {
	var tags any
	if m.Tags != nil {
		tags = m.Tags
	}

	return map[string]any{
		models.EventSearchIndexCols.EventID:            m.EventID,
		models.EventSearchIndexCols.TenantID:           m.TenantID,
		models.EventSearchIndexCols.AppID:              m.AppID,
		models.EventSearchIndexCols.ActorType:          m.ActorType,
		models.EventSearchIndexCols.ActorID:            m.ActorID,
		models.EventSearchIndexCols.Action:             m.Action,
		models.EventSearchIndexCols.TargetType:         m.TargetType,
		models.EventSearchIndexCols.TargetID:           m.TargetID,
		models.EventSearchIndexCols.Result:             m.Result,
		models.EventSearchIndexCols.OccurredAt:         m.OccurredAt,
		models.EventSearchIndexCols.IP:                 m.IP,
		models.EventSearchIndexCols.RiskLevel:          m.RiskLevel,
		models.EventSearchIndexCols.DataClassification: m.DataClassification,
		models.EventSearchIndexCols.Tags:               tags,
	}
}
