package constants

// 所有微服务名称常量
const (
	ServiceAccess    = "access"
	ServiceAudit     = "audit"
	ServiceAuth      = "auth"
	ServiceClients   = "clients"
	ServiceDirectory = "directory"
	ServiceImage     = "image"
	ServiceSystem    = "system"
	ServiceTenants   = "tenants"
)

// AllServices 返回所有服务名称列表
func AllServices() []string {
	return []string{
		ServiceAccess,
		ServiceAudit,
		ServiceAuth,
		ServiceClients,
		ServiceDirectory,
		ServiceImage,
		ServiceSystem,
		ServiceTenants,
	}
}
