package pipeline

import (
	"nfxid/modules/system/interfaces/pipeline/handler"
)

type Registry struct {
	SystemHandler *handler.SystemHandler
}
