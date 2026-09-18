package authmail

import (
	"context"
	"strings"

	authErr "nfxidentity/errors/src/auth"
	pkgemail "nfxidentity/pkgs/email"
)

func SendSignupWelcomeEmail(ctx context.Context, mail *pkgemail.EmailService, to, lang, accountID string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if mail == nil {
		return authErr.ErrSignupWelcomeSendFailed
	}
	to = strings.ToLower(strings.TrimSpace(to))
	if to == "" {
		return authErr.ErrInvalidEmail
	}
	copy := signupWelcomeCopy(lang)
	html := pkgemail.BuildSignupWelcomeEmailHTML(pkgemail.SignupWelcomeData{
		Headline:  copy.headline,
		Intro:     copy.intro,
		Highlight: copy.highlight,
		NextStep:  copy.nextStep,
		AccountID: accountID,
	})
	if err := mail.Send(pkgemail.EmailMessage{
		To:      []string{to},
		Subject: copy.subject,
		Body:    html,
		IsHTML:  true,
	}); err != nil {
		return authErr.ErrSignupWelcomeSendFailed.WithCause(err)
	}
	return nil
}

type signupWelcomeText struct {
	subject, headline, intro, highlight, nextStep string
}

func signupWelcomeCopy(lang string) signupWelcomeText {
	switch strings.ToLower(lang) {
	case "zh", "zh-cn", "zh-tw":
		return signupWelcomeText{
			subject:   "欢迎使用 NebulaForgeX Identity",
			headline:  "欢迎",
			intro:     "你的登录中心账号已创建。",
			highlight: "接下来请选择或创建资料，即可进入控制台。",
			nextStep:  "你可以在资料页绑定邮箱、GitHub，并打开登录通知。",
		}
	case "fr":
		return signupWelcomeText{
			subject:   "Bienvenue sur NebulaForgeX Identity",
			headline:  "Bienvenue",
			intro:     "Votre compte du centre de connexion a été créé.",
			highlight: "Choisissez ou créez un profil pour ouvrir la console.",
			nextStep:  "Vous pouvez lier un e-mail ou GitHub et activer l'avis de connexion.",
		}
	default:
		return signupWelcomeText{
			subject:   "Welcome to NebulaForgeX Identity",
			headline:  "Welcome",
			intro:     "Your login-center account is ready.",
			highlight: "Select or create a profile to open the console.",
			nextStep:  "You can link email or GitHub and turn on sign-in notices in your profile.",
		}
	}
}
