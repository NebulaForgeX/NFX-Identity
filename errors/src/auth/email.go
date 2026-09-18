package auth

import "nfxidentity/pkgs/errx"

var (
	ErrEmailBindingNotFound          = errx.NotFound("EMAIL_BINDING_NOT_FOUND", "email binding not found")
	ErrEmailAccountIDInvalid         = errx.InvalidArg("EMAIL_ACCOUNT_ID_INVALID", "invalid email account id")
	ErrEmailAddressRequired          = errx.InvalidArg("EMAIL_ADDRESS_REQUIRED", "email address is required")
	ErrEmailAddressFormatInvalid     = errx.InvalidArg("EMAIL_ADDRESS_FORMAT_INVALID", "invalid email address format")
	ErrEmailUnverified               = errx.InvalidArg("EMAIL_UNVERIFIED", "email is not verified")
	ErrPrimaryEmailRequired          = errx.InvalidArg("PRIMARY_EMAIL_REQUIRED", "a verified primary email is required")
	ErrEmailCannotDeleteLastVerified = errx.InvalidArg("EMAIL_CANNOT_DELETE_LAST_VERIFIED", "cannot delete the last verified email")
	ErrEmailCannotModifyPrimary      = errx.InvalidArg("EMAIL_CANNOT_MODIFY_PRIMARY", "cannot modify the primary email address")
	ErrEmailCannotDeletePrimary      = errx.InvalidArg("EMAIL_CANNOT_DELETE_PRIMARY", "cannot delete the primary email")
	ErrEmailCreateFailed             = errx.Internal("EMAIL_CREATE_FAILED", "failed to create email binding")
	ErrEmailDeleteFailed             = errx.Internal("EMAIL_DELETE_FAILED", "failed to delete email binding")
	// ErrEmailRegistrationCheckFailed wraps DB/infrastructure failures when checking whether an email is already registered.
	ErrEmailRegistrationCheckFailed = errx.Internal("EMAIL_REGISTRATION_CHECK_FAILED", "failed to verify email availability")
	// ErrContactUsRecipientNotConfigured is returned when EMAIL_CONTACT_US_RECIPIENT is empty.
	ErrContactUsRecipientNotConfigured = errx.Internal("CONTACT_US_RECIPIENT_NOT_CONFIGURED", "contact us recipient is not configured")
	ErrContactUsMessageRequired        = errx.InvalidArg("CONTACT_US_MESSAGE_REQUIRED", "message is required")
	ErrContactUsProfileRequired        = errx.InvalidArg("CONTACT_US_PROFILE_REQUIRED", "a selected profile is required to send a contact message")
)

/*
!EMAIL_BINDING_NOT_FOUND
*en<Email binding not found>
*zh<邮箱绑定不存在>
*fr<Liaison email introuvable>

!EMAIL_ACCOUNT_ID_INVALID
*en<Invalid email account id>
*zh<邮箱绑定账号 ID 无效>
*fr<Identifiant de compte de l'email invalide>

!EMAIL_ADDRESS_REQUIRED
*en<Email address is required>
*zh<邮箱地址必填>
*fr<Adresse email requise>

!EMAIL_ADDRESS_FORMAT_INVALID
*en<Invalid email address format>
*zh<邮箱地址格式无效>
*fr<Format d'adresse email invalide>

!EMAIL_REGISTRATION_CHECK_FAILED
*en<Failed to verify email availability>
*zh<无法校验邮箱是否可用>
*fr<Impossible de vérifier la disponibilité de l'email>
*p*en<Sorry, we couldn't verify whether this email is available right now. Please try again later.>
*p*zh<抱歉，暂时无法校验该邮箱是否可用，请稍后再试。>
*p*fr<Désolé, nous n'avons pas pu vérifier la disponibilité de cet email. Veuillez réessayer plus tard.>

!EMAIL_UNVERIFIED
*en<Email is not verified>
*zh<邮箱尚未验证>
*fr<Email non vérifié>

!PRIMARY_EMAIL_REQUIRED
*en<A verified primary email is required>
*zh<需要已验证的主邮箱>
*fr<Un email principal vérifié est requis>
*p*en<Please set and verify a primary email before continuing.>
*p*zh<请先设置并验证主邮箱后再继续。>
*p*fr<Veuillez définir et vérifier un email principal avant de continuer.>

!EMAIL_CANNOT_DELETE_LAST_VERIFIED
*en<Cannot delete the last verified email>
*zh<不能删除最后一个已验证邮箱>
*fr<Impossible de supprimer le dernier email vérifié>

!EMAIL_CANNOT_MODIFY_PRIMARY
*en<Cannot modify the primary email address>
*zh<不能修改主邮箱地址>
*fr<Impossible de modifier l'adresse email principale>

!EMAIL_CANNOT_DELETE_PRIMARY
*en<Cannot delete the primary email>
*zh<不能删除主邮箱>
*fr<Impossible de supprimer l'email principal>

!EMAIL_CREATE_FAILED
*en<Failed to create email binding>
*zh<创建邮箱绑定失败>
*fr<Échec de la création de la liaison email>
*p*en<Sorry, we couldn't add this email right now. This is on us—please try again later.>
*p*zh<抱歉，暂时无法绑定该邮箱，这是我们的问题，请稍后再试。>
*p*fr<Désolé, cet email n'a pas pu être lié pour le moment. Veuillez réessayer plus tard.>

!EMAIL_DELETE_FAILED
*en<Failed to delete email binding>
*zh<删除邮箱绑定失败>
*fr<Échec de la suppression de la liaison email>
*p*en<Sorry, we couldn't remove this email right now. Please try again later.>
*p*zh<抱歉，暂时无法解绑该邮箱，请稍后再试。>
*p*fr<Désolé, cet email n'a pas pu être supprimé pour le moment. Veuillez réessayer plus tard.>

!CONTACT_US_RECIPIENT_NOT_CONFIGURED
*en<Contact us recipient is not configured>
*zh<联系表单收件人未配置>
*fr<Destinataire du formulaire de contact non configuré>
*p*en<Sorry, we couldn't send your message right now. This is on us—please try again later.>
*p*zh<抱歉，暂时无法发送留言，这是我们的问题，请稍后再试。>
*p*fr<Désolé, votre message n'a pas pu être envoyé. Veuillez réessayer plus tard.>

!CONTACT_US_MESSAGE_REQUIRED
*en<Message is required>
*zh<留言内容必填>
*fr<Le message est requis>

!CONTACT_US_PROFILE_REQUIRED
*en<A selected profile is required to send a contact message>
*zh<发送留言前请先登录并选择身份>
*fr<Un profil sélectionné est requis pour envoyer un message>
*/
