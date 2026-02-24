package audit

import "nfxid/pkgs/errx"

const (
	CodeEventNotFound = "EVENT_NOT_FOUND"
)

var (
	ErrEventNotFound = errx.NotFound(CodeEventNotFound, "event not found")
)

/*
!EVENT_NOT_FOUND
*en<event not found>
*zh<事件不存在>
*fr<événement introuvable>

*/
