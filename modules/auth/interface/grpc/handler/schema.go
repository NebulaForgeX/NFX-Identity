package handler

import (
	"context"

	schemapb "nfxidentity/protos/gen/common/schema"
)

type SchemaClearFn func(ctx context.Context) (tablesCleared int32, err error)

type SchemaHandler struct {
	schemapb.UnimplementedSchemaServiceServer
	clear SchemaClearFn
}

func NewSchemaHandler(clear SchemaClearFn) *SchemaHandler {
	return &SchemaHandler{clear: clear}
}

func (h *SchemaHandler) ClearSchema(ctx context.Context, req *schemapb.ClearSchemaRequest) (*schemapb.ClearSchemaResponse, error) {
	tablesCleared, err := h.clear(ctx)
	if err != nil {
		errMsg := err.Error()
		return &schemapb.ClearSchemaResponse{
			Success:       false,
			ErrorMessage:  &errMsg,
			TablesCleared: 0,
		}, nil
	}
	return &schemapb.ClearSchemaResponse{
		Success:       true,
		TablesCleared: tablesCleared,
	}, nil
}
