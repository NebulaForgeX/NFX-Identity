package authmail

import (
	"context"
	"strings"

	authErr "nfxidentity/errors/src/auth"
	pkgemail "nfxidentity/pkgs/email"
)

func SendVerificationEmail(ctx context.Context, mail *pkgemail.EmailService, to, code, lang string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if mail == nil {
		return authErr.ErrSendVerificationEmailFailed
	}
	to = strings.ToLower(strings.TrimSpace(to))
	if to == "" {
		return authErr.ErrInvalidEmail
	}
	subject := "Your verification code"
	switch strings.ToLower(lang) {
	case "zh", "zh-cn", "zh-tw":
		subject = "你的验证码"
	case "fr":
		subject = "Votre code de vérification"
	}
	html := pkgemail.BuildVerificationEmailHTML(code)
	if err := mail.Send(pkgemail.EmailMessage{
		To:      []string{to},
		Subject: subject,
		Body:    html,
		IsHTML:  true,
	}); err != nil {
		return authErr.ErrSendVerificationEmailFailed.WithCause(err)
	}
	return nil
}
