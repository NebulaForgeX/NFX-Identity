package handler

import (
	"context"

	"nfxid/events/audit"
	"nfxid/pkgs/logx"

	"github.com/ThreeDotsLabs/watermill/message"
)

type AuditHandler struct {
}

func NewAuditHandler() *AuditHandler {
	return &AuditHandler{}
}

func (h *AuditHandler) OnActorSnapshotsInvalidateCache(ctx context.Context, evt audit.ActorSnapshotsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Audit Pipeline] 收到 ActorSnapshots 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

func (h *AuditHandler) OnEventRetentionPoliciesInvalidateCache(ctx context.Context, evt audit.EventRetentionPoliciesInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Audit Pipeline] 收到 EventRetentionPolicies 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

func (h *AuditHandler) OnEventSearchIndexInvalidateCache(ctx context.Context, evt audit.EventSearchIndexInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Audit Pipeline] 收到 EventSearchIndex 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

func (h *AuditHandler) OnEventsInvalidateCache(ctx context.Context, evt audit.EventsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Audit Pipeline] 收到 Events 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}

func (h *AuditHandler) OnHashChainCheckpointsInvalidateCache(ctx context.Context, evt audit.HashChainCheckpointsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Audit Pipeline] 收到 HashChainCheckpoints 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	return nil
}
