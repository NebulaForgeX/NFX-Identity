package event_search_index

import (
	"context"
	eventSearchIndexCommands "nfxid/modules/audit/application/event_search_index/commands"
)

// DeleteEventSearchIndex 删除事件搜索索引
func (s *Service) DeleteEventSearchIndex(ctx context.Context, cmd eventSearchIndexCommands.DeleteEventSearchIndexCmd) error {
	// Delete from repository (hard delete)
	return s.eventSearchIndexRepo.Delete.ByID(ctx, cmd.EventSearchIndexID)
}
