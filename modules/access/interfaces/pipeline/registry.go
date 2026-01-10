package pipeline

import (
	"nfxid/modules/access/interfaces/pipeline/handler"
)

type Registry struct {
	Access *handler.AccessHandler
}
