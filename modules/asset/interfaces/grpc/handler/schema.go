package handler

import (
	"context"

	"nfxidentity/pkgs/postgresqlx"
	schemapb "nfxidentity/protos/gen/common/schema"

	"gorm.io/gorm"
)

type SchemaHandler struct {
	schemapb.UnimplementedSchemaServiceServer
	db         *gorm.DB
	schemaName string
}

func NewSchemaHandler(db *gorm.DB, schemaName string) *SchemaHandler {
	return &SchemaHandler{db: db, schemaName: schemaName}
}

func (h *SchemaHandler) ClearSchema(ctx context.Context, req *schemapb.ClearSchemaRequest) (*schemapb.ClearSchemaResponse, error) {
	tablesCleared, err := postgresqlx.ClearSchema(ctx, h.db, h.schemaName, nil)
	if err != nil {
		errMsg := err.Error()
		return &schemapb.ClearSchemaResponse{Success: false, ErrorMessage: &errMsg, TablesCleared: 0}, nil
	}
	return &schemapb.ClearSchemaResponse{Success: true, TablesCleared: int32(tablesCleared)}, nil
}
