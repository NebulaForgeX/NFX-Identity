package auth

import "nfxidentity/pkgs/errx"

var (
	ErrAccountNotFound              = errx.NotFound("ACCOUNT_NOT_FOUND", "account not found")
	ErrAccountStatusInvalid         = errx.InvalidArg("ACCOUNT_STATUS_INVALID", "invalid account status")
	ErrAccountSignupPlatformInvalid = errx.InvalidArg("ACCOUNT_SIGNUP_PLATFORM_INVALID", "invalid signup platform")
	ErrAccountEmailRequired         = errx.InvalidArg("ACCOUNT_EMAIL_REQUIRED", "account email is required")
	ErrAccountDeviceIDRequired      = errx.InvalidArg("ACCOUNT_DEVICE_ID_REQUIRED", "device id is required")
	ErrAccountPhoneRequired         = errx.InvalidArg("ACCOUNT_PHONE_REQUIRED", "account phone is required")
	ErrAccountFirstNameRequired     = errx.InvalidArg("ACCOUNT_FIRST_NAME_REQUIRED", "first name is required")
	ErrAccountLastNameRequired      = errx.InvalidArg("ACCOUNT_LAST_NAME_REQUIRED", "last name is required")
	ErrAccountPasswordRequired      = errx.InvalidArg("ACCOUNT_PASSWORD_REQUIRED", "password is required")
	ErrInvalidCredentials           = errx.Unauthorized("INVALID_CREDENTIALS", "invalid credentials")
	ErrEmailOrPasswordIncorrect     = errx.Unauthorized("EMAIL_OR_PASSWORD_INCORRECT", "email or password is incorrect")
	ErrInvalidRefreshToken          = errx.Unauthorized("INVALID_REFRESH_TOKEN", "invalid or expired refresh token")
	ErrAccountAlreadyExists         = errx.Conflict("ACCOUNT_ALREADY_EXISTS", "account already exists")
	ErrEmailAlreadyExists           = errx.Conflict("EMAIL_ALREADY_EXISTS", "email already exists or already in use")
	ErrPhoneAlreadyExists           = errx.Conflict("PHONE_ALREADY_EXISTS", "phone already exists or already in use")
	ErrAccountInactive              = errx.Forbidden("ACCOUNT_INACTIVE", "account is inactive")
	ErrAccountHasNoEmail            = errx.InvalidArg("ACCOUNT_HAS_NO_EMAIL", "account has no email")
	ErrInvalidAccountPermission     = errx.InvalidArg("INVALID_ACCOUNT_PERMISSION", "invalid account permission")
	ErrProfileScopeInvalid          = errx.InvalidArg("PROFILE_SCOPE_INVALID", "invalid profile scope")
	// 注册/登录/邮件
	ErrEmailSendTimeout     = errx.Internal("EMAIL_SEND_TIMEOUT", "email service connection timed out")
	ErrEmailAuthFailed      = errx.Unauthorized("EMAIL_AUTH_FAILED", "email server authentication failed")
	ErrEmailSendFailed      = errx.Internal("EMAIL_SEND_FAILED", "failed to send email")
	ErrSignupFailed         = errx.InvalidArg("SIGNUP_FAILED", "signup failed")
	ErrLoginFailed          = errx.Internal("LOGIN_FAILED", "login failed")
	ErrUpdateEmailFailed    = errx.InvalidArg("UPDATE_EMAIL_FAILED", "failed to update email")
	ErrUpdatePasswordFailed = errx.InvalidArg("UPDATE_PASSWORD_FAILED", "failed to update password")
	// 操作失败
	ErrAccountCreateFailed = errx.Internal("ACCOUNT_CREATE_FAILED", "failed to create account")
	ErrAccountUpdateFailed = errx.Internal("ACCOUNT_UPDATE_FAILED", "failed to update account")
	ErrAccountDeleteFailed = errx.Internal("ACCOUNT_DELETE_FAILED", "failed to delete account")

	// Query (read) failures
	ErrAccountBaseGetFailed               = errx.Internal("ACCOUNT_BASE_GET_FAILED", "failed to get account base record")
	ErrAccountEmailsGetFailed             = errx.Internal("ACCOUNT_EMAILS_GET_FAILED", "failed to get account emails")
	ErrAccountPhonesGetFailed             = errx.Internal("ACCOUNT_PHONES_GET_FAILED", "failed to get account phones")
	ErrAccountProfileGetFailed            = errx.Internal("ACCOUNT_PROFILE_GET_FAILED", "failed to get account profile")
	ErrAccountProfileAvatarsGetFailed     = errx.Internal("ACCOUNT_PROFILE_AVATARS_GET_FAILED", "failed to get account profile avatars")
	ErrAccountProfileBackgroundsGetFailed = errx.Internal("ACCOUNT_PROFILE_BACKGROUNDS_GET_FAILED", "failed to get account profile backgrounds")
	ErrAccountProfileSettingsGetFailed    = errx.Internal("ACCOUNT_PROFILE_SETTINGS_GET_FAILED", "failed to get account profile settings")
	ErrAccountByEmailLookupFailed         = errx.Internal("ACCOUNT_BY_EMAIL_LOOKUP_FAILED", "failed to look up account by email")
	ErrAccountByPhoneLookupFailed         = errx.Internal("ACCOUNT_BY_PHONE_LOOKUP_FAILED", "failed to look up account by phone")
	ErrAccountListQueryFailed             = errx.Internal("ACCOUNT_LIST_QUERY_FAILED", "failed to query account list")
	ErrAccountListScopeResolveFailed      = errx.Internal("ACCOUNT_LIST_SCOPE_RESOLVE_FAILED", "failed to resolve account list scope")
	ErrAccountListProfilesGetFailed       = errx.Internal("ACCOUNT_LIST_PROFILES_GET_FAILED", "failed to get profiles for account list")
	ErrAccountCountFailed                 = errx.Internal("ACCOUNT_COUNT_FAILED", "failed to count accounts")

	ErrAccountFullInformationCacheInvalidationFailed = errx.Internal(
		"ACCOUNT_FULL_INFORMATION_CACHE_INVALIDATION_FAILED",
		"data updated but cache may be delayed",
	)
)

/*
!ACCOUNT_NOT_FOUND
*en<Account not found>
*zh<账号不存在>
*fr<Compte introuvable>

!ACCOUNT_STATUS_INVALID
*en<Invalid account status>
*zh<无效的账号状态>
*fr<Statut de compte invalide>

!ACCOUNT_SIGNUP_PLATFORM_INVALID
*en<Invalid signup platform>
*zh<无效的注册平台>
*fr<Plateforme d'inscription invalide>

!ACCOUNT_EMAIL_REQUIRED
*en<Account email is required>
*zh<账号邮箱必填>
*fr<Email du compte requis>

!ACCOUNT_DEVICE_ID_REQUIRED
*en<Device id is required>
*zh<设备标识必填>
*fr<Identifiant de l'appareil requis>

!ACCOUNT_PHONE_REQUIRED
*en<Account phone is required>
*zh<账号手机号必填>
*fr<Téléphone du compte requis>

!ACCOUNT_FIRST_NAME_REQUIRED
*en<First name is required>
*zh<名字必填>
*fr<Prénom requis>

!ACCOUNT_LAST_NAME_REQUIRED
*en<Last name is required>
*zh<姓氏必填>
*fr<Nom requis>

!ACCOUNT_PASSWORD_REQUIRED
*en<Password is required>
*zh<密码必填>
*fr<Mot de passe requis>

!INVALID_CREDENTIALS
*en<Invalid credentials>
*zh<账号或密码错误>
*fr<Identifiants invalides>

!EMAIL_OR_PASSWORD_INCORRECT
*en<Email or password is incorrect>
*zh<邮箱或密码不正确>
*fr<Email ou mot de passe incorrect>

!INVALID_REFRESH_TOKEN
*en<Invalid or expired refresh token>
*zh<刷新令牌无效或已过期>
*fr<Jeton de rafraîchissement invalide ou expiré>

!ACCOUNT_ALREADY_EXISTS
*en<Account already exists>
*zh<账号已存在>
*fr<Compte déjà existant>

!EMAIL_ALREADY_EXISTS
*en<Email already exists or already in use>
*zh<邮箱已被使用>
*fr<Email déjà utilisé>

!PHONE_ALREADY_EXISTS
*en<Phone already exists or already in use>
*zh<手机号已被使用>
*fr<Téléphone déjà utilisé>

!ACCOUNT_INACTIVE
*en<Account is inactive>
*zh<账号已停用>
*fr<Compte inactif>

!ACCOUNT_HAS_NO_EMAIL
*en<Account has no email>
*zh<账号未绑定邮箱>
*fr<Compte sans email>

!INVALID_ACCOUNT_PERMISSION
*en<Invalid account permission>
*zh<无效的账号权限>
*fr<Permission de compte invalide>

!EMAIL_SEND_TIMEOUT
*en<Email service connection timed out>
*zh<邮件服务连接超时>
*fr<Délai de connexion au service email dépassé>
*p*en<Sorry, sending email timed out. This is on us—please try again later.>
*p*zh<抱歉，邮件发送超时，这是我们的问题，请稍后再试。>
*p*fr<Désolé, l'envoi de l'email a expiré. Veuillez réessayer plus tard.>

!EMAIL_AUTH_FAILED
*en<Email server authentication failed>
*zh<邮件服务器认证失败>
*fr<Échec d'authentification du serveur email>

!EMAIL_SEND_FAILED
*en<Failed to send email>
*zh<邮件发送失败>
*fr<Échec de l'envoi de l'email>
*p*en<Sorry, we couldn't send the email right now. This is on us—we'll fix it as soon as we can. Please try again later.>
*p*zh<抱歉，邮件暂时发送失败，这是我们的问题，我们会尽快处理。请稍后再试。>
*p*fr<Désolé, l'email n'a pas pu être envoyé. Le problème vient de notre côté ; veuillez réessayer plus tard.>

!SIGNUP_FAILED
*en<Signup failed>
*zh<注册失败>
*fr<Inscription échouée>
*p*en<Sorry, we couldn't complete your signup right now. This is on us—please try again later.>
*p*zh<抱歉，注册暂时无法完成，这是我们的问题，请稍后再试。>
*p*fr<Désolé, l'inscription n'a pas pu aboutir. Veuillez réessayer plus tard.>

!LOGIN_FAILED
*en<Login failed>
*zh<登录失败>
*fr<Échec de la connexion>
*p*en<Sorry, login didn't work right now. This is on us—please try again later.>
*p*zh<抱歉，登录暂时失败，这是我们的问题，请稍后再试。>
*p*fr<Désolé, la connexion a échoué. Le problème vient de notre côté ; veuillez réessayer plus tard.>

!UPDATE_EMAIL_FAILED
*en<Failed to update email>
*zh<更新邮箱失败>
*fr<Échec de la mise à jour de l'email>
*p*en<Sorry, we couldn't update your email right now. Please try again later.>
*p*zh<抱歉，暂时无法更新邮箱，请稍后再试。>
*p*fr<Désolé, l'email n'a pas pu être mis à jour. Veuillez réessayer plus tard.>

!UPDATE_PASSWORD_FAILED
*en<Failed to update password>
*zh<修改密码失败>
*fr<Échec de la mise à jour du mot de passe>
*p*en<Sorry, we couldn't change your password right now. Please try again later.>
*p*zh<抱歉，暂时无法修改密码，请稍后再试。>
*p*fr<Désolé, le mot de passe n'a pas pu être modifié. Veuillez réessayer plus tard.>

!ACCOUNT_CREATE_FAILED
*en<Failed to create account>
*zh<创建账号失败>
*fr<Échec de la création du compte>
*p*en<Sorry, we couldn't create your account right now. This is on us—please try again later.>
*p*zh<抱歉，暂时无法创建账号，这是我们的问题，请稍后再试。>
*p*fr<Désolé, le compte n'a pas pu être créé. Veuillez réessayer plus tard.>

!ACCOUNT_UPDATE_FAILED
*en<Failed to update account>
*zh<更新账号失败>
*fr<Échec de la mise à jour du compte>
*p*en<Sorry, we couldn't update your account right now. Please try again later.>
*p*zh<抱歉，暂时无法更新账号信息，请稍后再试。>
*p*fr<Désolé, le compte n'a pas pu être mis à jour. Veuillez réessayer plus tard.>

!ACCOUNT_DELETE_FAILED
*en<Failed to delete account>
*zh<删除账号失败>
*fr<Échec de la suppression du compte>
*p*en<Sorry, we couldn't delete your account right now. Please try again later.>
*p*zh<抱歉，暂时无法删除账号，请稍后再试。>
*p*fr<Désolé, le compte n'a pas pu être supprimé. Veuillez réessayer plus tard.>

!ACCOUNT_BASE_GET_FAILED
*en<Failed to get account base record>
*zh<获取账号基础信息失败>
*fr<Échec de la récupération de l'enregistrement de base du compte>
*p*en<Sorry, we couldn't load your account information right now. Please try again later.>
*p*zh<抱歉，暂时无法加载账号信息，请稍后再试。>
*p*fr<Désolé, les informations du compte n'ont pas pu être chargées. Veuillez réessayer plus tard.>

!ACCOUNT_EMAILS_GET_FAILED
*en<Failed to get account emails>
*zh<获取账号邮箱列表失败>
*fr<Échec de la récupération des emails du compte>
*p*en<Sorry, we couldn't load your email addresses right now. Please try again later.>
*p*zh<抱歉，暂时无法加载邮箱信息，请稍后再试。>
*p*fr<Désolé, les adresses email n'ont pas pu être chargées. Veuillez réessayer plus tard.>

!ACCOUNT_PHONES_GET_FAILED
*en<Failed to get account phones>
*zh<获取账号手机号列表失败>
*fr<Échec de la récupération des téléphones du compte>
*p*en<Sorry, we couldn't load your phone numbers right now. Please try again later.>
*p*zh<抱歉，暂时无法加载手机号信息，请稍后再试。>
*p*fr<Désolé, les numéros de téléphone n'ont pas pu être chargés. Veuillez réessayer plus tard.>

!ACCOUNT_PROFILE_GET_FAILED
*en<Failed to get account profile>
*zh<获取账号资料失败>
*fr<Échec de la récupération du profil du compte>
*p*en<Sorry, we couldn't load your profile right now. Please try again later.>
*p*zh<抱歉，暂时无法加载资料信息，请稍后再试。>
*p*fr<Désolé, le profil n'a pas pu être chargé. Veuillez réessayer plus tard.>

!ACCOUNT_PROFILE_AVATARS_GET_FAILED
*en<Failed to get account profile avatars>
*zh<获取账号头像列表失败>
*fr<Échec de la récupération des avatars du profil>
*p*en<Sorry, we couldn't load your avatars right now. Please try again later.>
*p*zh<抱歉，暂时无法加载头像，请稍后再试。>
*p*fr<Désolé, les avatars n'ont pas pu être chargés. Veuillez réessayer plus tard.>

!ACCOUNT_PROFILE_BACKGROUNDS_GET_FAILED
*en<Failed to get account profile backgrounds>
*zh<获取账号背景图列表失败>
*fr<Échec de la récupération des fonds du profil>
*p*en<Sorry, we couldn't load your background images right now. Please try again later.>
*p*zh<抱歉，暂时无法加载背景图，请稍后再试。>
*p*fr<Désolé, les images de fond n'ont pas pu être chargées. Veuillez réessayer plus tard.>

!ACCOUNT_PROFILE_SETTINGS_GET_FAILED
*en<Failed to get account profile settings>
*zh<获取账号资料设置失败>
*fr<Échec de la récupération des paramètres du profil>
*p*en<Sorry, we couldn't load your settings right now. Please try again later.>
*p*zh<抱歉，暂时无法加载设置，请稍后再试。>
*p*fr<Désolé, les paramètres n'ont pas pu être chargés. Veuillez réessayer plus tard.>

!ACCOUNT_BY_EMAIL_LOOKUP_FAILED
*en<Failed to look up account by email>
*zh<按邮箱查找账号失败>
*fr<Échec de la recherche du compte par email>
*p*en<Sorry, we couldn't look up that account right now. Please try again later.>
*p*zh<抱歉，暂时无法完成查询，请稍后再试。>
*p*fr<Désolé, la recherche du compte a échoué. Veuillez réessayer plus tard.>

!ACCOUNT_BY_PHONE_LOOKUP_FAILED
*en<Failed to look up account by phone>
*zh<按手机号查找账号失败>
*fr<Échec de la recherche du compte par téléphone>
*p*en<Sorry, we couldn't look up that account right now. Please try again later.>
*p*zh<抱歉，暂时无法完成查询，请稍后再试。>
*p*fr<Désolé, la recherche du compte a échoué. Veuillez réessayer plus tard.>

!ACCOUNT_LIST_QUERY_FAILED
*en<Failed to query account list>
*zh<查询账号列表失败>
*fr<Échec de la requête de la liste des comptes>
*p*en<Sorry, we couldn't load the account list right now. Please try again later.>
*p*zh<抱歉，暂时无法加载账号列表，请稍后再试。>
*p*fr<Désolé, la liste des comptes n'a pas pu être chargée. Veuillez réessayer plus tard.>

!ACCOUNT_LIST_SCOPE_RESOLVE_FAILED
*en<Failed to resolve account list scope>
*zh<解析账号列表范围失败>
*fr<Échec de la résolution de la portée de la liste des comptes>
*p*en<Sorry, we couldn't load this list right now. Please try again later.>
*p*zh<抱歉，暂时无法加载列表，请稍后再试。>
*p*fr<Désolé, la liste n'a pas pu être chargée. Veuillez réessayer plus tard.>

!ACCOUNT_LIST_PROFILES_GET_FAILED
*en<Failed to get profiles for account list>
*zh<获取账号列表关联资料失败>
*fr<Échec de la récupération des profils pour la liste des comptes>
*p*en<Sorry, we couldn't load related profiles right now. Please try again later.>
*p*zh<抱歉，暂时无法加载相关资料，请稍后再试。>
*p*fr<Désolé, les profils associés n'ont pas pu être chargés. Veuillez réessayer plus tard.>

!ACCOUNT_COUNT_FAILED
*en<Failed to count accounts>
*zh<统计账号数量失败>
*fr<Échec du comptage des comptes>
*p*en<Sorry, we couldn't complete this count right now. Please try again later.>
*p*zh<抱歉，暂时无法完成统计，请稍后再试。>
*p*fr<Désolé, le décompte n'a pas pu aboutir. Veuillez réessayer plus tard.>

!ACCOUNT_FULL_INFORMATION_CACHE_INVALIDATION_FAILED
*en<Data updated but cache may be delayed>
*zh<数据已更新，缓存可能延迟>
*fr<Données mises à jour, le cache peut être retardé>
*/
