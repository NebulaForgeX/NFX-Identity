package pipeline

import (
	"nfxid/modules/image/interfaces/pipeline/handler"
)

type Registry struct {
	ImageHandler *handler.ImageHandler
}
