package events

import "nfxidentity/pkgs/kafkax/eventbus"

const (
	TKAuth    eventbus.TopicKey = "auth"
	TKAuthDLQ eventbus.TopicKey = "auth_poison"

	TKAsset    eventbus.TopicKey = "asset"
	TKAssetDLQ eventbus.TopicKey = "asset_poison"
)

type AuthTopic struct{}

func (AuthTopic) TopicKey() eventbus.TopicKey { return TKAuth }

type AssetTopic struct{}

func (AssetTopic) TopicKey() eventbus.TopicKey { return TKAsset }
