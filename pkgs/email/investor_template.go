package email

import (
	"bytes"
	"fmt"
	"html/template"
)

/**
 ** BuildInvestorContactEmailHTML builds HTML content for investor contact emails.
 ** BuildInvestorContactEmailHTML 构建投资人联系邮件的 HTML 内容。
 *
 * Parameters:
 *   !- name: Investor name (投资人姓名)
 *   !- email: Investor email (投资人邮箱)
 *   !- subject: Message subject (消息标题)
 *   !- message: Message content (消息内容)
 *
 * Returns:
 *   !- string: HTML email content (HTML 邮件内容)
 *
 * Examples:
 *
 * 	htmlBody := email.BuildInvestorContactEmailHTML("John Investor", "john@investment.com", "Investment Inquiry", "I am interested...")
 * 	err := emailService.Send(email.EmailMessage{
 * 		To:      []string{"admin@example.com"},
 * 		Subject: "Investor Contact",
 * 		Body:    htmlBody,
 * 		IsHTML:  true,
 * 	})
 */
func BuildInvestorContactEmailHTML(name, email, subject, message string) string {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Investor Contact</title>
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; max-width: 600px; margin: 0 auto; padding: 20px;">
    <div style="background: linear-gradient(135deg, #2ecc71 0%, #27ae60 100%); padding: 20px; border-radius: 8px 8px 0 0;">
        <h1 style="color: white; margin: 0;">💼 Investor Contact</h1>
    </div>
    
    <div style="background-color: #f9f9f9; padding: 20px; border: 1px solid #ddd; border-top: none; border-radius: 0 0 8px 8px;">
        <h2 style="color: #27ae60; margin-top: 0;">Contact Information</h2>
        
        <table style="width: 100%; border-collapse: collapse;">
            <tr>
                <td style="padding: 10px; border-bottom: 1px solid #ddd; font-weight: bold; width: 30%;">Name:</td>
                <td style="padding: 10px; border-bottom: 1px solid #ddd;">{{.Name}}</td>
            </tr>
            <tr>
                <td style="padding: 10px; border-bottom: 1px solid #ddd; font-weight: bold;">Email:</td>
                <td style="padding: 10px; border-bottom: 1px solid #ddd;">
                    <a href="mailto:{{.Email}}" style="color: #27ae60; text-decoration: none;">{{.Email}}</a>
                </td>
            </tr>
            <tr>
                <td style="padding: 10px; border-bottom: 1px solid #ddd; font-weight: bold;">Subject:</td>
                <td style="padding: 10px; border-bottom: 1px solid #ddd;">{{.Subject}}</td>
            </tr>
        </table>
        
        <h3 style="color: #27ae60; margin-top: 20px;">Message:</h3>
        <div style="background-color: white; padding: 15px; border-radius: 4px; border-left: 4px solid #27ae60;">
            <p style="margin: 0; white-space: pre-wrap;">{{.Message}}</p>
        </div>
        
        <div style="margin-top: 20px; padding: 15px; background-color: #d4edda; border-left: 4px solid #28a745; border-radius: 4px;">
            <p style="margin: 0; color: #155724;">
                <strong>✉️ Reply Required:</strong> Please respond to this investor inquiry at your earliest convenience.
            </p>
        </div>
    </div>
    
    <div style="text-align: center; margin-top: 20px; padding: 20px; color: #666; font-size: 12px;">
        <p>This is an automated notification from Rex App</p>
        <p>© 2025 Rex App. All rights reserved.</p>
    </div>
</body>
</html>
`

	t := template.Must(template.New("investor").Parse(tmpl))
	var buf bytes.Buffer
	data := struct {
		Name    string
		Email   string
		Subject string
		Message string
	}{
		Name:    name,
		Email:   email,
		Subject: subject,
		Message: message,
	}

	if err := t.Execute(&buf, data); err != nil {
		return fmt.Sprintf("Error executing template: %v", err)
	}

	return buf.String()
}

