package email

import (
	"bytes"
	"html/template"
)

type ContactUsEmailData struct {
	Name    string
	Email   string
	Phone   *string
	Wechat  *string
	Message string
}

func BuildContactUsAdminEmailHTML(data ContactUsEmailData) string {
	tmpl := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>新的联系表单留言</title>
</head>
<body style="margin:0;padding:0;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,'Helvetica Neue',Arial,sans-serif;background-color:#f5f5f5;">
	<table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:#f5f5f5;padding:32px 0;">
		<tr>
			<td align="center">
				<table width="600" cellpadding="0" cellspacing="0" border="0" style="background-color:#ffffff;border-radius:12px;box-shadow:0 2px 12px rgba(0,0,0,0.08);">
					<tr>
						<td style="padding:32px 32px 24px;background:linear-gradient(135deg,#0ea5e9 0%,#2563eb 100%);border-radius:12px 12px 0 0;color:#ffffff;">
							<h2 style="margin:0;font-size:24px;">🍵 有新的用户留言</h2>
							<p style="margin:12px 0 0;font-size:14px;opacity:0.85;">来自 {{.Name}} ({{.Email}})</p>
						</td>
					</tr>
					<tr>
						<td style="padding:32px;">
							<table width="100%" cellpadding="0" cellspacing="0" border="0" style="border-collapse:collapse;">
								<tr>
									<td style="width:120px;padding:12px 0;color:#6b7280;font-size:14px;">姓名</td>
									<td style="padding:12px 0;font-size:15px;color:#111827;">{{.Name}}</td>
								</tr>
								<tr>
									<td style="width:120px;padding:12px 0;color:#6b7280;font-size:14px;">邮箱</td>
									<td style="padding:12px 0;font-size:15px;color:#2563eb;">{{.Email}}</td>
								</tr>
								{{if .Phone}}
								<tr>
									<td style="width:120px;padding:12px 0;color:#6b7280;font-size:14px;">电话</td>
									<td style="padding:12px 0;font-size:15px;color:#111827;">{{.Phone}}</td>
								</tr>
								{{end}}
								{{if .Wechat}}
								<tr>
									<td style="width:120px;padding:12px 0;color:#6b7280;font-size:14px;">微信</td>
									<td style="padding:12px 0;font-size:15px;color:#111827;">{{.Wechat}}</td>
								</tr>
								{{end}}
								<tr>
									<td style="width:120px;padding:12px 0;color:#6b7280;font-size:14px;vertical-align:top;">留言内容</td>
									<td style="padding:12px 0;font-size:15px;color:#111827;line-height:1.6;">
										<div style="padding:16px;background-color:#f9fafb;border-radius:8px;border:1px solid #e5e7eb;white-space:pre-wrap;">{{.Message}}</div>
									</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding:20px 32px 32px;border-top:1px solid #e5e7eb;text-align:center;color:#9ca3af;font-size:12px;">
							<p style="margin:0;">这是一封来自官网联系表单的自动通知邮件，请及时跟进。</p>
							<p style="margin:8px 0 0;">© 2025 nfxid</p>
						</td>
					</tr>
				</table>
			</td>
		</tr>
	</table>
</body>
</html>
`

	t := template.Must(template.New("contact-us").Parse(tmpl))
	var buf bytes.Buffer

	if err := t.Execute(&buf, data); err != nil {
		return "New contact message received.\n\n" +
			"Name: " + data.Name + "\n" +
			"Email: " + data.Email + "\n" +
			maybeLine("Phone", data.Phone) +
			maybeLine("Wechat", data.Wechat) +
			"Message:\n" + data.Message
	}

	return buf.String()
}

func maybeLine(label string, value *string) string {
	if value == nil || *value == "" {
		return ""
	}
	return label + ": " + *value + "\n"
}

type ContactUsConfirmationData struct {
	Name       string
	AdminEmail string
}

func BuildContactUsConfirmationEmailHTML(data ContactUsConfirmationData) string {
	tmpl := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>我们已收到您的留言</title>
</head>
<body style="margin:0;padding:0;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,'Helvetica Neue',Arial,sans-serif;background-color:#f9fafb;">
	<table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:#f9fafb;padding:32px 0;">
		<tr>
			<td align="center">
				<table width="600" cellpadding="0" cellspacing="0" border="0" style="background-color:#ffffff;border-radius:12px;box-shadow:0 2px 12px rgba(0,0,0,0.06);">
					<tr>
						<td style="padding:32px 32px 24px;background:linear-gradient(135deg,#22d3ee 0%,#0ea5e9 100%);border-radius:12px 12px 0 0;color:#ffffff;">
							<h2 style="margin:0;font-size:22px;">感谢您的联系，{{.Name}}</h2>
						</td>
					</tr>
					<tr>
						<td style="padding:32px;color:#1f2937;font-size:15px;line-height:1.7;">
							<p style="margin:0 0 16px;">我们已经成功收到您的留言，团队会尽快与您取得联系。</p>
							<p style="margin:0 0 16px;">请留意来自 <strong>{{.AdminEmail}}</strong> 的邮件（这是我们的官方客服邮箱），以免遗漏后续沟通。</p>
							<p style="margin:0;">若需立即沟通，也可直接回复此邮件，我们会尽快响应。</p>
						</td>
					</tr>
					<tr>
						<td style="padding:20px 32px 32px;border-top:1px solid #e5e7eb;text-align:center;color:#9ca3af;font-size:12px;">
							<p style="margin:0;">这是一封自动通知邮件，请勿直接回复。</p>
							<p style="margin:8px 0 0;">© 2025 nfxid</p>
						</td>
					</tr>
				</table>
			</td>
		</tr>
	</table>
</body>
</html>
`

	t := template.Must(template.New("contact-us-confirmation").Parse(tmpl))
	var buf bytes.Buffer

	if err := t.Execute(&buf, data); err != nil {
		return "我们已经收到您的留言，会尽快联系您。\n\nidentity 团队"
	}

	return buf.String()
}
