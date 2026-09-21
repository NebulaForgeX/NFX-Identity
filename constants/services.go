package constants

const (
	ServiceAuth  = "auth"
	ServiceAsset = "asset"
)

func AllServices() []string {
	return []string{ServiceAuth, ServiceAsset}
}
