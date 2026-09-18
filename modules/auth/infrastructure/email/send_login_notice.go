package authmail

import (
	"context"
	"strings"

	authErr "nfxidentity/errors/src/auth"
	pkgemail "nfxidentity/pkgs/email"
)

func SendLoginNoticeEmail(ctx context.Context, mail *pkgemail.EmailService, to, lang, accountID, loginAt, loginEmail, provider, subject string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if mail == nil {
		return authErr.ErrLoginNoticeSendFailed
	}
	to = strings.ToLower(strings.TrimSpace(to))
	if to == "" {
		return authErr.ErrInvalidEmail
	}
	copy := loginNoticeCopy(lang)
	html := pkgemail.BuildLoginNoticeEmailHTML(pkgemail.LoginNoticeData{
		Headline:         copy.headline,
		Intro:            copy.intro,
		LabelTime:        copy.labelTime,
		LabelEmail:       copy.labelEmail,
		LabelProvider:    copy.labelProvider,
		LabelSubject:     copy.labelSubject,
		IfNotYou:         copy.ifNotYou,
		LoginAt:          loginAt,
		LoginEmail:       loginEmail,
		IdentityProvider: provider,
		ProviderSubject:  subject,
		AccountID:        accountID,
	})
	if err := mail.Send(pkgemail.EmailMessage{
		To:      []string{to},
		Subject: copy.subject,
		Body:    html,
		IsHTML:  true,
	}); err != nil {
		return authErr.ErrLoginNoticeSendFailed.WithCause(err)
	}
	return nil
}

type loginNoticeText struct {
	subject, headline, intro, labelTime, labelEmail, labelProvider, labelSubject, ifNotYou string
}

func loginNoticeCopy(lang string) loginNoticeText {
	switch strings.ToLower(lang) {
	case "zh", "zh-cn", "zh-tw":
		return loginNoticeText{
			subject:       "你的 Identity 账号有新的登录",
			headline:      "登录通知",
			intro:         "我们检测到一次新的登录。如果这是你本人操作，可以忽略此邮件。",
			labelTime:     "时间：",
			labelEmail:    "邮箱：",
			labelProvider: "登录方式：",
			labelSubject:  "标识：",
			ifNotYou:      "如果这不是你的操作，请立即修改密码并检查已绑定的登录方式。",
		}
	case "fr":
		return loginNoticeText{
			subject:       "Nouvelle connexion à votre compte Identity",
			headline:      "Avis de connexion",
			intro:         "Une nouvelle connexion a été détectée. Si c'était vous, ignorez cet e-mail.",
			labelTime:     "Heure : ",
			labelEmail:    "E-mail : ",
			labelProvider: "Méthode : ",
			labelSubject:  "Identifiant : ",
			ifNotYou:      "Si ce n'était pas vous, changez votre mot de passe et vérifiez vos identités liées.",
		}
	default:
		return loginNoticeText{
			subject:       "New sign-in to your Identity account",
			headline:      "Sign-in notice",
			intro:         "A new sign-in was recorded. If this was you, you can ignore this email.",
			labelTime:     "Time: ",
			labelEmail:    "Email: ",
			labelProvider: "Method: ",
			labelSubject:  "Identifier: ",
			ifNotYou:      "If this was not you, change your password and review linked sign-in methods.",
		}
	}
}
