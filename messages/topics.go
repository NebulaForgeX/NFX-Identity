package messages

const (
	MKAuth  = "auth"
	MKAsset = "asset"
)

type AuthTopic struct{}

func (AuthTopic) RoutingKey() string { return MKAuth }

type AssetTopic struct{}

func (AssetTopic) RoutingKey() string { return MKAsset }
