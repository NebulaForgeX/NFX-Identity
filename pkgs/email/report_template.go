package email

import (
	"bytes"
	"fmt"
	"html/template"
)

/**
 ** BuildReportEmailHTML builds HTML content for report emails.
 ** BuildReportEmailHTML 构建举报邮件的 HTML 内容。
 *
 * Parameters:
 *   !- reportType: Type of report (user/message/chatroom) (举报类型)
 *   !- targetID: ID of the reported item (被举报对象的 ID)
 *   !- reportCategory: Category of report (举报类别)
 *   !- details: Report details (举报详情)
 *   !- hasEvidence: Whether evidence files are attached (是否有证据附件)
 *
 * Returns:
 *   !- string: HTML email content (HTML 邮件内容)
 *
 * Examples:
 *
 * 	htmlBody := email.BuildReportEmailHTML("user", "user-123", "harassment", "Details here", true)
 * 	err := emailService.Send(email.EmailMessage{
 * 		To:      []string{"admin@example.com"},
 * 		Subject: "New Report",
 * 		Body:    htmlBody,
 * 		IsHTML:  true,
 * 	})
 */
func BuildReportEmailHTML(reportType, targetID, reportCategory, details string, hasEvidence bool) string {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Report Notification</title>
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; max-width: 600px; margin: 0 auto; padding: 20px;">
    <div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 8px 8px 0 0;">
        <h1 style="color: white; margin: 0;">🚨 New Report Received</h1>
    </div>
    
    <div style="background-color: #f9f9f9; padding: 20px; border: 1px solid #ddd; border-top: none; border-radius: 0 0 8px 8px;">
        <h2 style="color: #667eea; margin-top: 0;">Report Details</h2>
        
        <table style="width: 100%; border-collapse: collapse;">
            <tr>
                <td style="padding: 10px; border-bottom: 1px solid #ddd; font-weight: bold; width: 30%;">Report Type:</td>
                <td style="padding: 10px; border-bottom: 1px solid #ddd;">{{.ReportType}}</td>
            </tr>
            <tr>
                <td style="padding: 10px; border-bottom: 1px solid #ddd; font-weight: bold;">Target ID:</td>
                <td style="padding: 10px; border-bottom: 1px solid #ddd; font-family: monospace;">{{.TargetID}}</td>
            </tr>
            <tr>
                <td style="padding: 10px; border-bottom: 1px solid #ddd; font-weight: bold;">Category:</td>
                <td style="padding: 10px; border-bottom: 1px solid #ddd;">{{.ReportCategory}}</td>
            </tr>
            {{if .HasEvidence}}
            <tr>
                <td style="padding: 10px; border-bottom: 1px solid #ddd; font-weight: bold;">Evidence:</td>
                <td style="padding: 10px; border-bottom: 1px solid #ddd;">✅ Evidence files attached</td>
            </tr>
            {{end}}
        </table>
        
        <h3 style="color: #667eea; margin-top: 20px;">Report Details:</h3>
        <div style="background-color: white; padding: 15px; border-radius: 4px; border-left: 4px solid #667eea;">
            <p style="margin: 0; white-space: pre-wrap;">{{.Details}}</p>
        </div>
        
        <div style="margin-top: 20px; padding: 15px; background-color: #fff3cd; border-left: 4px solid #ffc107; border-radius: 4px;">
            <p style="margin: 0; color: #856404;">
                <strong>⚠️ Action Required:</strong> Please review this report and take appropriate action.
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

	t := template.Must(template.New("report").Parse(tmpl))
	var buf bytes.Buffer
	data := struct {
		ReportType     string
		TargetID       string
		ReportCategory string
		Details        string
		HasEvidence    bool
	}{
		ReportType:     reportType,
		TargetID:       targetID,
		ReportCategory: reportCategory,
		Details:        details,
		HasEvidence:    hasEvidence,
	}

	if err := t.Execute(&buf, data); err != nil {
		return fmt.Sprintf("Error executing template: %v", err)
	}

	return buf.String()
}

