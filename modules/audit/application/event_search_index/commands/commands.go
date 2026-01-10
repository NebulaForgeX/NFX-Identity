package commands

import (
	"nfxid/modules/audit/domain/event_search_index"

	"github.com/google/uuid"
)

// CreateEventSearchIndexCmd 创建事件搜索索引命令
type CreateEventSearchIndexCmd struct {
	EventID            string
	TenantID           *uuid.UUID
	AppID              *uuid.UUID
	ActorType          event_search_index.ActorType
	ActorID            uuid.UUID
	Action             string
	TargetType         *string
	TargetID           *uuid.UUID
	Result             event_search_index.ResultType
	OccurredAt         string
	IP                 *string
	RiskLevel          event_search_index.RiskLevel
	DataClassification event_search_index.DataClassification
	Tags               []string
}

// DeleteEventSearchIndexCmd 删除事件搜索索引命令
type DeleteEventSearchIndexCmd struct {
	EventSearchIndexID uuid.UUID
}
