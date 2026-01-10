package commands

import (
	"github.com/google/uuid"
)

// InitializeSystemCmd 初始化系统命令
type InitializeSystemCmd struct {
	Version string
}

// ResetSystemCmd 重置系统命令
type ResetSystemCmd struct {
	ResetBy uuid.UUID
}

// DeleteSystemStateCmd 删除系统状态命令
type DeleteSystemStateCmd struct {
	SystemStateID uuid.UUID
}
