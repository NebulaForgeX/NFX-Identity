package eventbus

import (
	"nfxid/modules/auth/interfaces/eventbus/handler"
)

type Registry struct {
	Auth  *handler.AuthHandler
	Image *handler.ImageHandler
}
