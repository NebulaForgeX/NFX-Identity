package pipeline

import (
	"nfxidentity/modules/clients/interfaces/pipeline/handler"
)

type Registry struct {
	Clients *handler.ClientsHandler
}
