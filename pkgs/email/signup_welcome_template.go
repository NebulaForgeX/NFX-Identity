package email

type SignupWelcomeData struct {
	Headline  string
	Intro     string
	Highlight string
	NextStep  string
	AccountID string
}

func BuildSignupWelcomeEmailHTML(data SignupWelcomeData) string {
	tmpl, err := loadTemplate("signup-welcome", signupWelcomeTemplateContent)
	if err != nil {
		return "Welcome to NebulaForgeX Identity."
	}
	html, err := executeTemplate(tmpl, data)
	if err != nil {
		return "Welcome to NebulaForgeX Identity."
	}
	return html
}
