package pipeline

import (
	"nfxidentity/modules/image/interfaces/pipeline/handler"
)

type Registry struct {
	ImageHandler *handler.ImageHandler
}
