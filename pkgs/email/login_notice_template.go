package email

import "fmt"

type LoginNoticeData struct {
	Headline         string
	Intro            string
	LabelTime        string
	LabelEmail       string
	LabelProvider    string
	LabelSubject     string
	IfNotYou         string
	LoginAt          string
	LoginEmail       string
	IdentityProvider string
	ProviderSubject  string
	AccountID        string
}

func BuildLoginNoticeEmailHTML(data LoginNoticeData) string {
	tmpl, err := loadTemplate("login-notice", loginNoticeTemplateContent)
	if err != nil {
		return fmt.Sprintf("A sign-in was recorded for %s at %s.", data.LoginEmail, data.LoginAt)
	}
	html, err := executeTemplate(tmpl, data)
	if err != nil {
		return fmt.Sprintf("A sign-in was recorded for %s at %s.", data.LoginEmail, data.LoginAt)
	}
	return html
}
