package eventbus

import (
	"nfxid/modules/image/interfaces/eventbus/handler"
)

type Registry struct {
	Auth *handler.AuthHandler
}
