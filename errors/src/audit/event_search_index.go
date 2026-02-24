package audit

import "nfxid/pkgs/errx"

const (
	CodeEventSearchIndexNotFound = "EVENT_SEARCH_INDEX_NOT_FOUND"
)

var (
	ErrEventSearchIndexNotFound = errx.NotFound(CodeEventSearchIndexNotFound, "event search index not found")
)
