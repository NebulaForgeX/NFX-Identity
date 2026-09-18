package auth

import "nfxidentity/pkgs/errx"

var (
	ErrForgerProfileNotFound            = errx.NotFound("FORGER_PROFILE_NOT_FOUND", "forger profile not found")
	ErrForgerProfileScopeRequired       = errx.Forbidden("FORGER_PROFILE_SCOPE_REQUIRED", "forger profile session required")
	ErrForgerProfileAccountIDInvalid    = errx.InvalidArg("FORGER_PROFILE_ACCOUNT_ID_INVALID", "invalid forger profile account id")
	ErrForgerProfileLanguageInvalid     = errx.InvalidArg("FORGER_PROFILE_LANGUAGE_INVALID", "invalid forger profile language")
	ErrForgerProfileForgerRoleInvalid   = errx.InvalidArg("FORGER_PROFILE_FORGER_ROLE_INVALID", "invalid forger profile forger role")
	ErrForgerProfilePatchEmpty          = errx.InvalidArg("FORGER_PROFILE_PATCH_EMPTY", "no forger profile fields to update")
	ErrForgerProfileFieldInvalid        = errx.InvalidArg("FORGER_PROFILE_FIELD_INVALID", "invalid forger profile field")
	ErrForgerProfileUpdateFailed        = errx.Internal("FORGER_PROFILE_UPDATE_FAILED", "failed to update forger profile")
	ErrForgerProfileCreateFailed        = errx.Internal("FORGER_PROFILE_CREATE_FAILED", "failed to create forger profile")
	ErrForgerProfileDeleteFailed        = errx.Internal("FORGER_PROFILE_DELETE_FAILED", "failed to delete forger profile")
	ErrForgerProfileCountFailed         = errx.Internal("FORGER_PROFILE_COUNT_FAILED", "failed to count forger profiles")
	ErrForgerProfileLimitReached        = errx.InvalidArg("FORGER_PROFILE_LIMIT_REACHED", "forger profile limit reached")
	ErrForgerProfileCannotDeleteLast    = errx.InvalidArg("FORGER_PROFILE_CANNOT_DELETE_LAST", "cannot delete the last forger profile")
	ErrForgerProfileCannotDeleteCurrent = errx.InvalidArg("FORGER_PROFILE_CANNOT_DELETE_CURRENT", "cannot delete the currently active forger profile")

	ErrForgerProfileAvatarAccountIDInvalid = errx.InvalidArg("FORGER_PROFILE_AVATAR_ACCOUNT_ID_INVALID", "invalid forger profile avatar account id")
	ErrForgerProfileAvatarProfileIDInvalid = errx.InvalidArg("FORGER_PROFILE_AVATAR_PROFILE_ID_INVALID", "invalid forger profile avatar profile id")
	ErrForgerProfileAvatarImageIDInvalid   = errx.InvalidArg("FORGER_PROFILE_AVATAR_IMAGE_ID_INVALID", "invalid forger profile avatar image id")

	ErrForgerProfileBackgroundAccountIDInvalid = errx.InvalidArg("FORGER_PROFILE_BACKGROUND_ACCOUNT_ID_INVALID", "invalid forger profile background account id")
	ErrForgerProfileBackgroundProfileIDInvalid = errx.InvalidArg("FORGER_PROFILE_BACKGROUND_PROFILE_ID_INVALID", "invalid forger profile background profile id")
	ErrForgerProfileBackgroundImageIDInvalid   = errx.InvalidArg("FORGER_PROFILE_BACKGROUND_IMAGE_ID_INVALID", "invalid forger profile background image id")
	ErrForgerProfileBackgroundSortOrderInvalid = errx.InvalidArg("FORGER_PROFILE_BACKGROUND_SORT_ORDER_INVALID", "invalid forger profile background sort order")
	ErrForgerProfileBackgroundConfirmFailed    = errx.Internal("FORGER_PROFILE_BACKGROUND_CONFIRM_FAILED", "failed to confirm forger profile backgrounds")
	ErrForgerProfileBackgroundDeleteFailed     = errx.Internal("FORGER_PROFILE_BACKGROUND_DELETE_FAILED", "failed to delete forger profile background image")

	ErrForgerProfileSearchFailed   = errx.Internal("FORGER_PROFILE_SEARCH_FAILED", "failed to search forger profiles")
	ErrForgerProfileBatchGetFailed = errx.Internal("FORGER_PROFILE_BATCH_GET_FAILED", "failed to batch get forger profiles")

	ErrForgerProfileSettingsNotFound         = errx.NotFound("FORGER_PROFILE_SETTINGS_NOT_FOUND", "forger profile settings not found")
	ErrForgerProfileSettingsProfileIDInvalid = errx.InvalidArg("FORGER_PROFILE_SETTINGS_PROFILE_ID_INVALID", "invalid forger profile settings profile id")
	ErrForgerProfileSettingsPatchEmpty       = errx.InvalidArg("FORGER_PROFILE_SETTINGS_PATCH_EMPTY", "no forger profile settings fields to update")
	ErrForgerProfileSettingsUpdateFailed     = errx.Internal("FORGER_PROFILE_SETTINGS_UPDATE_FAILED", "failed to update forger profile settings")
	ErrForgerProfileSettingsCreateFailed     = errx.Internal("FORGER_PROFILE_SETTINGS_CREATE_FAILED", "failed to create forger profile settings")
	ErrForgerProfileSettingsGetFailed        = errx.Internal("FORGER_PROFILE_SETTINGS_GET_FAILED", "failed to get forger profile settings")
)

/*
!FORGER_PROFILE_NOT_FOUND
*en<Forger profile not found>
*zh<创作者资料不存在>
*fr<Profil forger introuvable>

!FORGER_PROFILE_ACCOUNT_ID_INVALID
*en<Invalid forger profile account id>
*zh<创作者资料账号 ID 无效>
*fr<Identifiant de compte du profil invalide>

!FORGER_PROFILE_LANGUAGE_INVALID
*en<Invalid forger profile language>
*zh<创作者资料语言无效>
*fr<Langue du profil invalide>

!FORGER_PROFILE_FORGER_ROLE_INVALID
*en<Invalid forger profile forger role>
*zh<社区资料角色无效>
*fr<Rôle communautaire du profil invalide>

!FORGER_PROFILE_PATCH_EMPTY
*en<No forger profile fields to update>
*zh<未提供要更新的资料字段>
*fr<Aucun champ de profil à mettre à jour>

!FORGER_PROFILE_FIELD_INVALID
*en<Invalid forger profile field>
*zh<创作者资料字段无效>
*fr<Champ de profil invalide>

!FORGER_PROFILE_UPDATE_FAILED
*en<Failed to update forger profile>
*zh<更新创作者资料失败>
*fr<Échec de la mise à jour du profil>
*p*en<Sorry, we couldn't update your profile right now. Please try again later.>
*p*zh<抱歉，暂时无法更新资料，请稍后再试。>
*p*fr<Désolé, le profil n'a pas pu être mis à jour. Veuillez réessayer plus tard.>

!FORGER_PROFILE_CREATE_FAILED
*en<Failed to create forger profile>
*zh<创建创作者资料失败>
*fr<Échec de la création du profil>
*p*en<Sorry, we couldn't create your profile right now. This is on us—please try again later.>
*p*zh<抱歉，暂时无法创建资料，这是我们的问题，请稍后再试。>
*p*fr<Désolé, le profil n'a pas pu être créé. Veuillez réessayer plus tard.>

!FORGER_PROFILE_DELETE_FAILED
*en<Failed to delete forger profile>
*zh<删除创作者资料失败>
*fr<Échec de la suppression du profil>
*p*en<Sorry, we couldn't delete this profile right now. Please try again later.>
*p*zh<抱歉，暂时无法删除资料，请稍后再试。>
*p*fr<Désolé, le profil n'a pas pu être supprimé. Veuillez réessayer plus tard.>

!FORGER_PROFILE_COUNT_FAILED
*en<Failed to count forger profiles>
*zh<统计创作者资料数量失败>
*fr<Échec du comptage des profils>
*p*en<Sorry, we couldn't complete this count right now. Please try again later.>
*p*zh<抱歉，暂时无法完成统计，请稍后再试。>
*p*fr<Désolé, le décompte n'a pas pu aboutir. Veuillez réessayer plus tard.>

!FORGER_PROFILE_LIMIT_REACHED
*en<Forger profile limit reached>
*zh<已达到资料数量上限>
*fr<Limite de profils atteinte>

!FORGER_PROFILE_CANNOT_DELETE_LAST
*en<Cannot delete the last forger profile>
*zh<无法删除最后一个资料>
*fr<Impossible de supprimer le dernier profil>

!FORGER_PROFILE_CANNOT_DELETE_CURRENT
*en<Cannot delete the currently active forger profile>
*zh<无法删除当前正在使用的资料>
*fr<Impossible de supprimer le profil actif>

!FORGER_PROFILE_AVATAR_ACCOUNT_ID_INVALID
*en<Invalid forger profile avatar account id>
*zh<用户头像关联账号 ID 无效>
*fr<Identifiant de compte de l'avatar invalide>

!FORGER_PROFILE_AVATAR_PROFILE_ID_INVALID
*en<Invalid forger profile avatar profile id>
*zh<用户头像关联资料 ID 无效>
*fr<Identifiant de profil de l'avatar invalide>

!FORGER_PROFILE_AVATAR_IMAGE_ID_INVALID
*en<Invalid forger profile avatar image id>
*zh<用户头像图片 ID 无效>
*fr<Identifiant d'image de l'avatar invalide>

!FORGER_PROFILE_BACKGROUND_ACCOUNT_ID_INVALID
*en<Invalid forger profile background account id>
*zh<用户背景图关联账号 ID 无效>
*fr<Identifiant de compte du fond invalide>

!FORGER_PROFILE_BACKGROUND_PROFILE_ID_INVALID
*en<Invalid forger profile background profile id>
*zh<用户背景图关联资料 ID 无效>
*fr<Identifiant de profil du fond invalide>

!FORGER_PROFILE_BACKGROUND_IMAGE_ID_INVALID
*en<Invalid forger profile background image id>
*zh<用户背景图图片 ID 无效>
*fr<Identifiant d'image du fond invalide>

!FORGER_PROFILE_BACKGROUND_SORT_ORDER_INVALID
*en<Invalid forger profile background sort order>
*zh<用户背景图排序无效>
*fr<Ordre de tri du fond invalide>

!FORGER_PROFILE_BACKGROUND_CONFIRM_FAILED
*en<Failed to confirm forger profile backgrounds>
*zh<确认用户背景图失败>
*fr<Échec de la confirmation des fonds du profil>
*p*en<Sorry, we couldn't save your background images right now. Please try again later.>
*p*zh<抱歉，暂时无法保存背景图，请稍后再试。>
*p*fr<Désolé, les images de fond n'ont pas pu être enregistrées. Veuillez réessayer plus tard.>

!FORGER_PROFILE_BACKGROUND_DELETE_FAILED
*en<Failed to delete forger profile background image>
*zh<删除用户背景图失败>
*fr<Échec de la suppression de l'image de fond>
*p*en<Sorry, we couldn't remove this background image right now. Please try again later.>
*p*zh<抱歉，暂时无法删除背景图，请稍后再试。>
*p*fr<Désolé, l'image de fond n'a pas pu être supprimée. Veuillez réessayer plus tard.>

!FORGER_PROFILE_SEARCH_FAILED
*en<Failed to search forger profiles>
*zh<搜索创作者资料失败>
*fr<Échec de la recherche de profils>
*p*en<Sorry, search isn't working right now. Please try again later.>
*p*zh<抱歉，暂时无法搜索资料，请稍后再试。>
*p*fr<Désolé, la recherche de profils a échoué. Veuillez réessayer plus tard.>

!FORGER_PROFILE_BATCH_GET_FAILED
*en<Failed to batch get forger profiles>
*zh<批量获取创作者资料失败>
*fr<Échec de la récupération groupée des profils>
*p*en<Sorry, we couldn't load these profiles right now. Please try again later.>
*p*zh<抱歉，暂时无法加载资料，请稍后再试。>
*p*fr<Désolé, les profils n'ont pas pu être chargés. Veuillez réessayer plus tard.>

!FORGER_PROFILE_SETTINGS_NOT_FOUND
*en<Forger profile settings not found>
*zh<创作者资料设置不存在>
*fr<Paramètres du profil introuvables>

!FORGER_PROFILE_SETTINGS_PROFILE_ID_INVALID
*en<Invalid forger profile settings profile id>
*zh<创作者资料设置的资料 ID 无效>
*fr<Identifiant de profil des paramètres invalide>

!FORGER_PROFILE_SETTINGS_PATCH_EMPTY
*en<No forger profile settings fields to update>
*zh<未提供要更新的资料设置字段>
*fr<Aucun paramètre de profil à mettre à jour>

!FORGER_PROFILE_SETTINGS_UPDATE_FAILED
*en<Failed to update forger profile settings>
*zh<更新创作者资料设置失败>
*fr<Échec de la mise à jour des paramètres du profil>
*p*en<Sorry, we couldn't update your settings right now. Please try again later.>
*p*zh<抱歉，暂时无法更新设置，请稍后再试。>
*p*fr<Désolé, les paramètres n'ont pas pu être mis à jour. Veuillez réessayer plus tard.>

!FORGER_PROFILE_SETTINGS_CREATE_FAILED
*en<Failed to create forger profile settings>
*zh<创建创作者资料设置失败>
*fr<Échec de la création des paramètres du profil>
*p*en<Sorry, we couldn't create your settings right now. This is on us—please try again later.>
*p*zh<抱歉，暂时无法创建设置，这是我们的问题，请稍后再试。>
*p*fr<Désolé, les paramètres n'ont pas pu être créés. Veuillez réessayer plus tard.>

!FORGER_PROFILE_SETTINGS_GET_FAILED
*en<Failed to get forger profile settings>
*zh<获取创作者资料设置失败>
*fr<Échec de la récupération des paramètres du profil>
*p*en<Sorry, we couldn't load your settings right now. Please try again later.>
*p*zh<抱歉，暂时无法加载设置，请稍后再试。>
*p*fr<Désolé, les paramètres n'ont pas pu être chargés. Veuillez réessayer plus tard.>
*/
