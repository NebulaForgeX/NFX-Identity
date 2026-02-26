package audit

import "nfxidentity/pkgs/errx"

const (
	CodeEventSearchIndexNotFound = "EVENT_SEARCH_INDEX_NOT_FOUND"
)

var (
	ErrEventSearchIndexNotFound = errx.NotFound(CodeEventSearchIndexNotFound, "event search index not found")
)

/*
!EVENT_SEARCH_INDEX_NOT_FOUND
*en<event search index not found>
*zh<事件搜索索引不存在>
*fr<index de recherche d'événement introuvable>

*/
