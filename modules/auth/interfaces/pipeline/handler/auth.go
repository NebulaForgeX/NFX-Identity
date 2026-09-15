package handler

import (
	"context"

	"nfxidentity/events"
	"nfxidentity/pkgs/logx"

	"github.com/ThreeDotsLabs/watermill/message"
)

type AuthEvent struct {
	events.AuthTopic
	Type      string `json:"type"`
	AccountID string `json:"account_id"`
}

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler { return &AuthHandler{} }

func (h *AuthHandler) OnAuthEvent(ctx context.Context, evt AuthEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Auth Pipeline] event type=%s account_id=%s", evt.Type, evt.AccountID)
	return nil
}
