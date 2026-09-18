package auth

import "nfxidentity/pkgs/errx"

var (
	// SMTP / transport or pre-send failures (caller may attach cause).
	ErrSendVerificationEmailFailed = errx.Internal("SEND_VERIFICATION_EMAIL_FAILED", "failed to send verification code")
	// Missing [email] templates_html_dir (or equivalent) when rendering verification HTML.
	ErrVerificationEmailTemplatesDirRequired = errx.InvalidArg("VERIFICATION_EMAIL_TEMPLATES_HTML_DIR_REQUIRED", "email HTML templates directory is required")
	// Disk/template parse or execute failure before send.
	ErrVerificationEmailTemplateFailed = errx.Internal("VERIFICATION_EMAIL_TEMPLATE_FAILED", "failed to render verification email")
	// Persisting OTP after send (application / cache layer).
	ErrVerificationCodeSaveFailed = errx.Internal("VERIFICATION_CODE_SAVE_FAILED", "failed to save verification code")
	// User-submitted code does not match stored value.
	ErrVerificationCodeWrong = errx.InvalidArg("VERIFICATION_CODE_WRONG", "incorrect verification code")
	// Stored OTP missing or past TTL.
	ErrVerificationCodeExpired = errx.Expired("VERIFICATION_CODE_EXPIRED", "verification code has expired")
)

/*
!SEND_VERIFICATION_EMAIL_FAILED
*en<Failed to send verification code>
*zh<发送验证码失败>
*fr<Échec de l'envoi du code de vérification>
*p*en<Sorry, we couldn't send the verification code right now. This is on us—we'll fix it as soon as we can. Please try again later.>
*p*zh<抱歉，验证码暂时发送失败，这是我们的问题，我们会尽快处理。请稍后再试。>
*p*fr<Désolé, l'envoi du code de vérification a échoué temporairement. Le problème vient de notre côté ; veuillez réessayer plus tard.>

!VERIFICATION_EMAIL_TEMPLATES_HTML_DIR_REQUIRED
*en<Email HTML templates directory is required>
*zh<未配置邮件 HTML 模板目录>
*fr<Répertoire des modèles HTML email requis>

!VERIFICATION_EMAIL_TEMPLATE_FAILED
*en<Failed to render verification email>
*zh<验证码邮件渲染失败>
*fr<Échec du rendu de l'email de vérification>
*p*en<Sorry, we couldn't prepare the verification email right now. This is on us—please try again later.>
*p*zh<抱歉，验证码邮件暂时无法生成，这是我们的问题，请稍后再试。>
*p*fr<Désolé, l'email de vérification n'a pas pu être préparé. Veuillez réessayer plus tard.>

!VERIFICATION_CODE_SAVE_FAILED
*en<Failed to save verification code>
*zh<验证码存储失败>
*fr<Échec de l'enregistrement du code de vérification>
*p*en<Sorry, we couldn't save your verification code right now. Please try sending it again.>
*p*zh<抱歉，验证码暂时无法保存，请重新获取后再试。>
*p*fr<Désolé, le code de vérification n'a pas pu être enregistré. Veuillez le renvoyer et réessayer.>

!VERIFICATION_CODE_WRONG
*en<Incorrect verification code>
*zh<验证码错误>
*fr<Code de vérification incorrect>
*p*en<Oops, that verification code isn't correct. Please check and try again.>
*p*zh<糟糕，验证码不正确，请重新输入。>
*p*fr<Oups, le code de vérification est incorrect. Veuillez réessayer.>

!VERIFICATION_CODE_EXPIRED
*en<Verification code has expired>
*zh<验证码已过期>
*fr<Code de vérification expiré>
*p*en<Oops, this verification code has expired. Please request a new one.>
*p*zh<糟糕，验证码已过期，请重新获取。>
*p*fr<Oups, ce code de vérification a expiré. Veuillez en demander un nouveau.>
*/
