package commands

// BootstrapInitCmd 系统初始化命令
// 用于初始化系统基础数据，包括：
// 1. 创建第一个系统管理员用户
// 2. 创建初始角色和权限
// 3. 创建用户凭证
// 4. 初始化其他服务的基础数据
type BootstrapInitCmd struct {
	// 初始化版本号
	Version string
	// 管理员用户名
	AdminUsername string
	// 管理员密码（用于创建用户凭证）
	AdminPassword string
	// 管理员邮箱（可选）
	AdminEmail *string
	// 管理员手机号（可选）
	AdminPhone *string
	// 管理员手机国家代码（可选）
	AdminCountryCode *string
}
