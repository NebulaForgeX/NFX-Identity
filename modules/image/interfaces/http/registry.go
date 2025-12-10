package http

import (
	"nebulaid/modules/image/interfaces/http/handler"
)

type Registry struct {
	Image     *handler.ImageHandler
	ImageType *handler.ImageTypeHandler
}
