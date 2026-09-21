package auth

const (
	LOGIN_SUCCESS                         = "LOGIN_SUCCESS"
	PROFILE_SELECTED                      = "PROFILE_SELECTED"
	TOKENS_REFRESHED                      = "TOKENS_REFRESHED"
	LOGOUT_SUCCESS                        = "LOGOUT_SUCCESS"
	SIGNUP_SUCCESS                        = "SIGNUP_SUCCESS"
	VERIFICATION_CODE_SENT                = "VERIFICATION_CODE_SENT"
	PREFERENCE_UPDATED                    = "PREFERENCE_UPDATED"
	USER_PROFILE_UPDATED                  = "USER_PROFILE_UPDATED"
	USER_PROFILE_SETTINGS_UPDATED         = "USER_PROFILE_SETTINGS_UPDATED"
	USER_PROFILE_AVATAR_UPDATED           = "USER_PROFILE_AVATAR_UPDATED"
	USER_PROFILE_AVATAR_CLEARED           = "USER_PROFILE_AVATAR_CLEARED"
	USER_PROFILE_CREATED                  = "USER_PROFILE_CREATED"
	USER_PROFILE_DELETED                  = "USER_PROFILE_DELETED"
	USER_PROFILE_BACKGROUNDS_UPDATED      = "USER_PROFILE_BACKGROUNDS_UPDATED"
	AUTHORITY_PROFILE_UPDATED             = "AUTHORITY_PROFILE_UPDATED"
	AUTHORITY_PROFILE_SETTINGS_UPDATED    = "AUTHORITY_PROFILE_SETTINGS_UPDATED"
	AUTHORITY_PROFILE_AVATAR_UPDATED      = "AUTHORITY_PROFILE_AVATAR_UPDATED"
	AUTHORITY_PROFILE_AVATAR_CLEARED      = "AUTHORITY_PROFILE_AVATAR_CLEARED"
	AUTHORITY_PROFILE_CREATED             = "AUTHORITY_PROFILE_CREATED"
	AUTHORITY_PROFILE_DELETED             = "AUTHORITY_PROFILE_DELETED"
	AUTHORITY_PROFILE_BACKGROUNDS_UPDATED = "AUTHORITY_PROFILE_BACKGROUNDS_UPDATED"
	AUTHORITY_ROLES_UPDATED               = "AUTHORITY_ROLES_UPDATED"
	EMAIL_ADDED                           = "EMAIL_ADDED"
	EMAIL_VERIFIED                        = "EMAIL_VERIFIED"
	EMAIL_UPDATED                         = "EMAIL_UPDATED"
	PRIMARY_EMAIL_UPDATED                 = "PRIMARY_EMAIL_UPDATED"
	EMAIL_DELETED                         = "EMAIL_DELETED"
	PHONE_ADDED                           = "PHONE_ADDED"
	PHONE_VERIFIED                        = "PHONE_VERIFIED"
	PHONE_UPDATED                         = "PHONE_UPDATED"
	PRIMARY_PHONE_UPDATED                 = "PRIMARY_PHONE_UPDATED"
	PHONE_DELETED                         = "PHONE_DELETED"
	PASSWORD_UPDATED                      = "PASSWORD_UPDATED"
)

/*
!LOGIN_SUCCESS
*title*en<Signed in>
*body*en<You have signed in successfully>
*title*zh<登录成功>
*body*zh<你已成功登录>
*title*fr<Connexion réussie>
*body*fr<Vous êtes connecté>

!PROFILE_SELECTED
*title*en<Profile selected>
*body*en<The profile has been selected>
*title*zh<已选择资料>
*body*zh<已切换到该资料>
*title*fr<Profil sélectionné>
*body*fr<Le profil a été sélectionné>

!TOKENS_REFRESHED
*title*en<Tokens refreshed>
*body*en<Access tokens have been refreshed>
*title*zh<令牌已刷新>
*body*zh<访问令牌已刷新>
*title*fr<Jetons actualisés>
*body*fr<Les jetons d'accès ont été actualisés>

!LOGOUT_SUCCESS
*title*en<Signed out>
*body*en<You have signed out>
*title*zh<已退出>
*body*zh<你已退出登录>
*title*fr<Déconnexion>
*body*fr<Vous êtes déconnecté>

!SIGNUP_SUCCESS
*title*en<Account created>
*body*en<Your account has been created>
*title*zh<注册成功>
*body*zh<账号已创建>
*title*fr<Compte créé>
*body*fr<Votre compte a été créé>

!VERIFICATION_CODE_SENT
*title*en<Code sent>
*body*en<A verification code has been sent>
*title*zh<验证码已发送>
*body*zh<验证码已发送>
*title*fr<Code envoyé>
*body*fr<Un code de vérification a été envoyé>

!PREFERENCE_UPDATED
*title*en<Preference updated>
*body*en<Your preference has been updated>
*title*zh<偏好已更新>
*body*zh<偏好设置已更新>
*title*fr<Préférence mise à jour>
*body*fr<Votre préférence a été mise à jour>

!USER_PROFILE_UPDATED
*title*en<Profile updated>
*body*en<The user profile has been updated>
*title*zh<资料已更新>
*body*zh<用户资料已更新>
*title*fr<Profil mis à jour>
*body*fr<Le profil utilisateur a été mis à jour>

!USER_PROFILE_SETTINGS_UPDATED
*title*en<Settings updated>
*body*en<User profile settings have been updated>
*title*zh<设置已更新>
*body*zh<用户资料设置已更新>
*title*fr<Paramètres mis à jour>
*body*fr<Les paramètres du profil utilisateur ont été mis à jour>

!USER_PROFILE_AVATAR_UPDATED
*title*en<Avatar updated>
*body*en<The user profile avatar has been updated>
*title*zh<头像已更新>
*body*zh<用户资料头像已更新>
*title*fr<Avatar mis à jour>
*body*fr<L'avatar du profil utilisateur a été mis à jour>

!USER_PROFILE_AVATAR_CLEARED
*title*en<Avatar cleared>
*body*en<The user profile avatar has been cleared>
*title*zh<头像已清除>
*body*zh<用户资料头像已清除>
*title*fr<Avatar effacé>
*body*fr<L'avatar du profil utilisateur a été effacé>

!USER_PROFILE_CREATED
*title*en<Profile created>
*body*en<The user profile has been created>
*title*zh<资料已创建>
*body*zh<用户资料已创建>
*title*fr<Profil créé>
*body*fr<Le profil utilisateur a été créé>

!USER_PROFILE_DELETED
*title*en<Profile deleted>
*body*en<The user profile has been deleted>
*title*zh<资料已删除>
*body*zh<用户资料已删除>
*title*fr<Profil supprimé>
*body*fr<Le profil utilisateur a été supprimé>

!USER_PROFILE_BACKGROUNDS_UPDATED
*title*en<Backgrounds updated>
*body*en<User profile backgrounds have been updated>
*title*zh<背景已更新>
*body*zh<用户资料背景已更新>
*title*fr<Arrière-plans mis à jour>
*body*fr<Les arrière-plans du profil utilisateur ont été mis à jour>

!AUTHORITY_PROFILE_UPDATED
*title*en<Profile updated>
*body*en<The authority profile has been updated>
*title*zh<资料已更新>
*body*zh<管理资料已更新>
*title*fr<Profil mis à jour>
*body*fr<Le profil d'autorité a été mis à jour>

!AUTHORITY_PROFILE_SETTINGS_UPDATED
*title*en<Settings updated>
*body*en<Authority profile settings have been updated>
*title*zh<设置已更新>
*body*zh<管理资料设置已更新>
*title*fr<Paramètres mis à jour>
*body*fr<Les paramètres du profil d'autorité ont été mis à jour>

!AUTHORITY_PROFILE_AVATAR_UPDATED
*title*en<Avatar updated>
*body*en<The authority profile avatar has been updated>
*title*zh<头像已更新>
*body*zh<管理资料头像已更新>
*title*fr<Avatar mis à jour>
*body*fr<L'avatar du profil d'autorité a été mis à jour>

!AUTHORITY_PROFILE_AVATAR_CLEARED
*title*en<Avatar cleared>
*body*en<The authority profile avatar has been cleared>
*title*zh<头像已清除>
*body*zh<管理资料头像已清除>
*title*fr<Avatar effacé>
*body*fr<L'avatar du profil d'autorité a été effacé>

!AUTHORITY_PROFILE_CREATED
*title*en<Profile created>
*body*en<The authority profile has been created>
*title*zh<资料已创建>
*body*zh<管理资料已创建>
*title*fr<Profil créé>
*body*fr<Le profil d'autorité a été créé>

!AUTHORITY_PROFILE_DELETED
*title*en<Profile deleted>
*body*en<The authority profile has been deleted>
*title*zh<资料已删除>
*body*zh<管理资料已删除>
*title*fr<Profil supprimé>
*body*fr<Le profil d'autorité a été supprimé>

!AUTHORITY_PROFILE_BACKGROUNDS_UPDATED
*title*en<Backgrounds updated>
*body*en<Authority profile backgrounds have been updated>
*title*zh<背景已更新>
*body*zh<管理资料背景已更新>
*title*fr<Arrière-plans mis à jour>
*body*fr<Les arrière-plans du profil d'autorité ont été mis à jour>

!AUTHORITY_ROLES_UPDATED
*title*en<Roles updated>
*body*en<Authority roles have been updated>
*title*zh<角色已更新>
*body*zh<管理角色已更新>
*title*fr<Rôles mis à jour>
*body*fr<Les rôles d'autorité ont été mis à jour>

!EMAIL_ADDED
*title*en<Email added>
*body*en<The email address has been added>
*title*zh<邮箱已添加>
*body*zh<邮箱地址已添加>
*title*fr<E-mail ajouté>
*body*fr<L'adresse e-mail a été ajoutée>

!EMAIL_VERIFIED
*title*en<Email verified>
*body*en<The email address has been verified>
*title*zh<邮箱已验证>
*body*zh<邮箱地址已验证>
*title*fr<E-mail vérifié>
*body*fr<L'adresse e-mail a été vérifiée>

!EMAIL_UPDATED
*title*en<Email updated>
*body*en<The email address has been updated>
*title*zh<邮箱已更新>
*body*zh<邮箱地址已更新>
*title*fr<E-mail mis à jour>
*body*fr<L'adresse e-mail a été mise à jour>

!PRIMARY_EMAIL_UPDATED
*title*en<Primary email updated>
*body*en<The primary email has been updated>
*title*zh<主邮箱已更新>
*body*zh<主邮箱已更新>
*title*fr<E-mail principal mis à jour>
*body*fr<L'e-mail principal a été mis à jour>

!EMAIL_DELETED
*title*en<Email deleted>
*body*en<The email address has been deleted>
*title*zh<邮箱已删除>
*body*zh<邮箱地址已删除>
*title*fr<E-mail supprimé>
*body*fr<L'adresse e-mail a été supprimée>

!PHONE_ADDED
*title*en<Phone added>
*body*en<The phone number has been added>
*title*zh<手机已添加>
*body*zh<手机号已添加>
*title*fr<Téléphone ajouté>
*body*fr<Le numéro de téléphone a été ajouté>

!PHONE_VERIFIED
*title*en<Phone verified>
*body*en<The phone number has been verified>
*title*zh<手机已验证>
*body*zh<手机号已验证>
*title*fr<Téléphone vérifié>
*body*fr<Le numéro de téléphone a été vérifié>

!PHONE_UPDATED
*title*en<Phone updated>
*body*en<The phone number has been updated>
*title*zh<手机已更新>
*body*zh<手机号已更新>
*title*fr<Téléphone mis à jour>
*body*fr<Le numéro de téléphone a été mis à jour>

!PRIMARY_PHONE_UPDATED
*title*en<Primary phone updated>
*body*en<The primary phone has been updated>
*title*zh<主手机已更新>
*body*zh<主手机号已更新>
*title*fr<Téléphone principal mis à jour>
*body*fr<Le téléphone principal a été mis à jour>

!PHONE_DELETED
*title*en<Phone deleted>
*body*en<The phone number has been deleted>
*title*zh<手机已删除>
*body*zh<手机号已删除>
*title*fr<Téléphone supprimé>
*body*fr<Le numéro de téléphone a été supprimé>

!PASSWORD_UPDATED
*title*en<Password updated>
*body*en<The password has been updated>
*title*zh<密码已更新>
*body*zh<密码已更新>
*title*fr<Mot de passe mis à jour>
*body*fr<Le mot de passe a été mis à jour>
*/
