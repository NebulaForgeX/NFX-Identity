package pipeline

import (
	"nfxid/modules/auth/interfaces/pipeline/handler"
)

type Registry struct {
	Auth *handler.AuthHandler
}
