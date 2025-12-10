package events

import "nebulaid/pkgs/eventbus"

type TopicKey = eventbus.TopicKey

const (
	// =============== Auth ===============
	TKAuth    TopicKey = "auth"
	TKAuthDLQ TopicKey = "auth_poison"

	// =============== Image ===============
	TKImage    TopicKey = "image"
	TKImageDLQ TopicKey = "image_poison"
)
