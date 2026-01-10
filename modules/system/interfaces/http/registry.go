package http

import (
	"nfxid/modules/system/interfaces/http/handler"
)

type Registry struct {
	SystemState *handler.SystemStateHandler
}
