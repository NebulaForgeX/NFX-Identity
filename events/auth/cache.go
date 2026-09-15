package auth

import "nfxidentity/events"

type RefreshTokensInvalidateCacheEvent struct {
	events.AuthTopic
	ID        string `json:"id"`
	Prefix    string `json:"prefix"`
	Namespace string `json:"namespace"`
}
