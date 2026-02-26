package pipeline

import (
	"nfxidentity/modules/directory/interfaces/pipeline/handler"
)

type Registry struct {
	DirectoryHandler *handler.DirectoryHandler
}
