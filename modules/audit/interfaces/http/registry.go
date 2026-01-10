package http

import (
	"nfxid/modules/audit/interfaces/http/handler"
)

type Registry struct {
	Event                *handler.EventHandler
	ActorSnapshot        *handler.ActorSnapshotHandler
	EventRetentionPolicy *handler.EventRetentionPolicyHandler
	EventSearchIndex     *handler.EventSearchIndexHandler
	HashChainCheckpoint  *handler.HashChainCheckpointHandler
}
