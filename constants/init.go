package constants

// 系统初始化（bootstrap）相关常量：初始管理员角色与权限

// InitAdminRoleKey 初始系统管理员角色 key
const InitAdminRoleKey = "system.admin"

// InitAdminRoleName 初始系统管理员角色显示名
const InitAdminRoleName = "System Administrator"

// InitAdminRoleDesc 初始系统管理员角色描述
const InitAdminRoleDesc = "System Administrator Role, with all permissions"

// InitAdminRoleScope 初始系统管理员角色作用域（global）
const InitAdminRoleScope = "global"

// InitPermission 初始化时创建的一条权限
type InitPermission struct {
	Key         string // 权限 key，如 directory.*
	Name        string // 显示名
	Description string // 描述
}

// InitPermissions 初始化时创建的基础权限（按服务/域使用 domain.* 表示该域下所有权限）
var InitPermissions = []InitPermission{
	{Key: "system.*", Name: "System", Description: "All system-level permissions"},
	{Key: "directory.*", Name: "Directory", Description: "All directory service permissions (users, profiles, etc.)"},
	{Key: "tenants.*", Name: "Tenants", Description: "All tenants service permissions"},
	{Key: "access.*", Name: "Access", Description: "All access service permissions (roles, permissions, grants)"},
	{Key: "auth.*", Name: "Auth", Description: "All auth service permissions"},
	{Key: "clients.*", Name: "Clients", Description: "All clients service permissions (apps, credentials)"},
	{Key: "audit.*", Name: "Audit", Description: "All audit service permissions"},
	{Key: "image.*", Name: "Image", Description: "All image service permissions"},
}
