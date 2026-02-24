package audit

import "nfxid/pkgs/errx"

const (
	CodeEventNotFound = "EVENT_NOT_FOUND"
)

var (
	ErrEventNotFound = errx.NotFound(CodeEventNotFound, "event not found")
)
