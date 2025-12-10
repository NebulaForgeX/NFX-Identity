package email

import (
	"bytes"
	"fmt"
	"html/template"
)

/**
 ** BuildVerificationEmailHTML builds a beautiful HTML email for verification code.
 ** BuildVerificationEmailHTML 构建验证码邮件的精美 HTML 内容。
 *
 * Parameters:
 *   !- code: Verification code (验证码)
 *
 * Returns:
 *   !- string: HTML email content (HTML 邮件内容)
 *
 * Examples:
 *
 * 	htmlBody := email.BuildVerificationEmailHTML("123456")
 * 	err := emailService.Send(email.EmailMessage{
 * 		To:      []string{"user@example.com"},
 * 		Subject: "Your Verification Code",
 * 		Body:    htmlBody,
 * 		IsHTML:  true,
 * 	})
 */
func BuildVerificationEmailHTML(code string) string {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Email Verification</title>
</head>
<body style="margin: 0; padding: 0; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif; background-color: #f5f5f5;">
    <table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color: #f5f5f5; padding: 40px 0;">
        <tr>
            <td align="center">
                <table width="600" cellpadding="0" cellspacing="0" border="0" style="background-color: #ffffff; border-radius: 12px; box-shadow: 0 2px 8px rgba(0,0,0,0.1);">
                    <!-- Header -->
                    <tr>
                        <td style="padding: 40px 40px 20px 40px; text-align: center; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); border-radius: 12px 12px 0 0;">
                            <h1 style="margin: 0; color: #ffffff; font-size: 28px; font-weight: 600;">Email Verification</h1>
                        </td>
                    </tr>
                    
                    <!-- Content -->
                    <tr>
                        <td style="padding: 40px;">
                            <p style="margin: 0 0 20px 0; font-size: 16px; line-height: 24px; color: #333333;">
                                Hello! You've requested to verify your email address. Please use the verification code below:
                            </p>
                            
                            <!-- Verification Code Box -->
                            <table width="100%" cellpadding="0" cellspacing="0" border="0" style="margin: 30px 0;">
                                <tr>
                                    <td align="center" style="background-color: #f8f9fa; border-radius: 8px; padding: 30px;">
                                        <div style="font-size: 36px; font-weight: 700; letter-spacing: 8px; color: #667eea; font-family: 'Courier New', monospace;">
                                            {{.Code}}
                                        </div>
                                    </td>
                                </tr>
                            </table>
                            
                            <p style="margin: 20px 0 0 0; font-size: 14px; line-height: 20px; color: #666666;">
                                ⏱️ This code will expire in <strong>5 minutes</strong>
                            </p>
                            
                            <p style="margin: 30px 0 0 0; font-size: 14px; line-height: 20px; color: #666666;">
                                If you didn't request this verification code, please ignore this email.
                            </p>
                        </td>
                    </tr>
                    
                    <!-- Footer -->
                    <tr>
                        <td style="padding: 20px 40px 40px 40px; text-align: center; border-top: 1px solid #eeeeee;">
                            <p style="margin: 0; font-size: 12px; color: #999999;">
                                This is an automated email, please do not reply.
                            </p>
                            <p style="margin: 10px 0 0 0; font-size: 12px; color: #999999;">
                                © 2025 Rex App. All rights reserved.
                            </p>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>
`

	t := template.Must(template.New("verification").Parse(tmpl))
	var buf bytes.Buffer
	data := struct {
		Code string
	}{
		Code: code,
	}

	if err := t.Execute(&buf, data); err != nil {
		return fmt.Sprintf("Your verification code is: %s\n\nThis code will expire in 5 minutes.", code)
	}

	return buf.String()
}
