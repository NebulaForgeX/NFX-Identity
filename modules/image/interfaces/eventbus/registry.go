package eventbus

import (
	"nebulaid/modules/image/interfaces/eventbus/handler"
)

type Registry struct {
	Auth *handler.AuthHandler
}
