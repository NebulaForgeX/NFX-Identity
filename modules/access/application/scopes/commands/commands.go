package commands

// CreateScopeCmd 创建作用域命令
type CreateScopeCmd struct {
	Scope       string
	Description *string
	IsSystem    bool
}

// UpdateScopeCmd 更新作用域命令
type UpdateScopeCmd struct {
	Scope       string
	Description *string
}

// DeleteScopeCmd 删除作用域命令
type DeleteScopeCmd struct {
	Scope string
}
