package pipeline

import (
	"nfxidentity/modules/auth/interfaces/pipeline/handler"
)

type Registry struct {
	Auth *handler.AuthHandler
}
