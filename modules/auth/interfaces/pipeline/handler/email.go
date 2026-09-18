package handler

import (
	"context"
	"strings"
	"time"

	"nfxidentity/events"
	authmail "nfxidentity/modules/auth/infrastructure/email"
	repofactory "nfxidentity/modules/auth/infrastructure/repository/factory"
	pkgemail "nfxidentity/pkgs/email"
	"nfxidentity/pkgs/transaction"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

type EmailHandler struct {
	mail        *pkgemail.EmailService
	repoFactory *repofactory.TxRepoFactory
}

func NewEmailHandler(mail *pkgemail.EmailService, factory *repofactory.TxRepoFactory) *EmailHandler {
	return &EmailHandler{mail: mail, repoFactory: factory}
}

func none() transaction.UoW { return transaction.UoW{} }

func (h *EmailHandler) SignupSuccess(ctx context.Context, evt events.SignupSuccessEvent, _ *message.Message) error {
	if h.mail == nil {
		return nil
	}
	return authmail.SendSignupWelcomeEmail(ctx, h.mail, evt.Email, evt.Lang, evt.AccountID.String())
}

func (h *EmailHandler) LoginSuccess(ctx context.Context, evt events.LoginSuccessEvent, _ *message.Message) error {
	if h.mail == nil || h.repoFactory == nil {
		return nil
	}
	kind := strings.TrimSpace(evt.ProfileKind)
	if kind == "" {
		kind = "forger"
	}
	st, err := h.repoFactory.Settings(none()).Get.ByID(ctx, kind, evt.ProfileID)
	if err != nil || !st.LoginNotification() {
		return nil
	}
	to := strings.TrimSpace(evt.LoginEmail)
	if to == "" {
		to = h.primaryEmail(ctx, evt.AccountID)
	}
	if to == "" {
		return nil
	}
	lang := h.profileLang(ctx, kind, evt.ProfileID)
	loginAt := evt.LoginAt.UTC()
	if loginAt.IsZero() {
		loginAt = time.Now().UTC()
	}
	return authmail.SendLoginNoticeEmail(
		ctx,
		h.mail,
		to,
		lang,
		evt.AccountID.String(),
		loginAt.Format(time.RFC3339),
		to,
		evt.IdentityProvider,
		evt.ProviderSubject,
	)
}

func (h *EmailHandler) primaryEmail(ctx context.Context, accountID uuid.UUID) string {
	rows, err := h.repoFactory.Email(none()).Get.ByAccountID(ctx, accountID)
	if err != nil {
		return ""
	}
	for _, row := range rows {
		if row.IsPrimary() {
			return row.Address()
		}
	}
	if len(rows) > 0 {
		return rows[0].Address()
	}
	return ""
}

func (h *EmailHandler) profileLang(ctx context.Context, kind string, profileID uuid.UUID) string {
	if kind == "authority" {
		p, err := h.repoFactory.Authority(none()).Get.ByID(ctx, profileID)
		if err == nil {
			return p.ProfileLanguage()
		}
		return "en"
	}
	p, err := h.repoFactory.Forger(none()).Get.ByID(ctx, profileID)
	if err == nil {
		return p.ProfileLanguage()
	}
	return "en"
}
