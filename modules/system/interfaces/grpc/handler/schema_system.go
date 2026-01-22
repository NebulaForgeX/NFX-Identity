package handler

import (
	"context"
	"fmt"

	"nfxid/pkgs/postgresqlx"
	schemapb "nfxid/protos/gen/common/schema"

	"gorm.io/gorm"
)

// SystemSchemaHandler System 服务的 schema 清空处理器
// 特殊处理：确保 system_state 表只有一条记录
type SystemSchemaHandler struct {
	schemapb.UnimplementedSchemaServiceServer
	db         *gorm.DB
	schemaName string
}

// NewSystemSchemaHandler 创建 System schema 处理器
func NewSystemSchemaHandler(db *gorm.DB, schemaName string) *SystemSchemaHandler {
	return &SystemSchemaHandler{
		db:         db,
		schemaName: schemaName,
	}
}

// ClearSchema 清空 schema 中所有表的数据（不删除表）
// 对于 system_state 表，确保只有一条记录
func (h *SystemSchemaHandler) ClearSchema(ctx context.Context, req *schemapb.ClearSchemaRequest) (*schemapb.ClearSchemaResponse, error) {
	// 特殊处理 system_state 表：排除它，单独处理
	systemStateTableName := fmt.Sprintf(`"%s"."system_state"`, h.schemaName)
	excludeTables := []string{systemStateTableName}

	// 清空其他表
	tablesCleared, err := postgresqlx.ClearSchema(ctx, h.db, h.schemaName, excludeTables)
	if err != nil {
		errMsg := err.Error()
		return &schemapb.ClearSchemaResponse{
			Success:       false,
			ErrorMessage:  &errMsg,
			TablesCleared: 0,
		}, nil
	}

	// 清空 system_state 表（删除所有记录，确保只有一条）
	if err := postgresqlx.DeleteAllFromTable(ctx, h.db, h.schemaName, "system_state"); err != nil {
		errMsg := fmt.Sprintf("failed to clear system_state table: %v", err)
		return &schemapb.ClearSchemaResponse{
			Success:       false,
			ErrorMessage:  &errMsg,
			TablesCleared: int32(tablesCleared),
		}, nil
	}

	totalCleared := tablesCleared + 1 // +1 for system_state

	return &schemapb.ClearSchemaResponse{
		Success:       true,
		TablesCleared: int32(totalCleared),
	}, nil
}
