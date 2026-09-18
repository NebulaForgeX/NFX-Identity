package pipeline

import (
	"nfxidentity/modules/auth/interface/pipeline/handler"
)

type Registry struct {
	Email *handler.EmailHandler
}
