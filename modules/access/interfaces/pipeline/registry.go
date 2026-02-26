package pipeline

import (
	"nfxidentity/modules/access/interfaces/pipeline/handler"
)

type Registry struct {
	Access *handler.AccessHandler
}
