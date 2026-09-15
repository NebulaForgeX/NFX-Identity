package constants

const (
	ServiceAuth   = "auth"
	ServiceAsset  = "asset"
	ServiceSystem = "system"
)

func AllServices() []string {
	return []string{ServiceAuth, ServiceAsset, ServiceSystem}
}
