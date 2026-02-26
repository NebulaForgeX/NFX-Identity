package http

import (
	"nfxidentity/modules/image/interfaces/http/handler"
)

type Registry struct {
	Image        *handler.ImageHandler
	ImageType    *handler.ImageTypeHandler
	ImageVariant *handler.ImageVariantHandler
	ImageTag     *handler.ImageTagHandler
	Upload       *handler.UploadHandler
}
