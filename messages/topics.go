package messages

const (
	MKAuth   = "auth"
	MKAsset  = "asset"
	MKSystem = "system"
)

type AuthTopic struct{}

func (AuthTopic) RoutingKey() string { return MKAuth }

type AssetTopic struct{}

func (AssetTopic) RoutingKey() string { return MKAsset }

type SystemTopic struct{}

func (SystemTopic) RoutingKey() string { return MKSystem }
