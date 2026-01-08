package event_search_index

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Repo 是 EventSearchIndex 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, esi *EventSearchIndex) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*EventSearchIndex, error)
	ByEventID(ctx context.Context, eventID string) (*EventSearchIndex, error)
	ByTenantID(ctx context.Context, tenantID uuid.UUID, startTime, endTime *time.Time) ([]*EventSearchIndex, error)
	ByActor(ctx context.Context, actorType ActorType, actorID uuid.UUID, startTime, endTime *time.Time) ([]*EventSearchIndex, error)
	ByAction(ctx context.Context, action string, startTime, endTime *time.Time) ([]*EventSearchIndex, error)
	ByTags(ctx context.Context, tags []string) ([]*EventSearchIndex, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByEventID(ctx context.Context, eventID string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, esi *EventSearchIndex) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByEventID(ctx context.Context, eventID string) error
}
