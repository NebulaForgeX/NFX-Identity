package constants

// 系统初始化（bootstrap）相关常量：初始管理员角色与权限

// InitAdminRoleKey 初始系统管理员角色 key
const InitAdminRoleKey = "system.admin"

// InitAdminRoleName 初始系统管理员角色显示名
const InitAdminRoleName = "System Administrator"

// InitAdminRoleDesc 初始系统管理员角色描述
const InitAdminRoleDesc = "System Administrator Role, with all permissions"

// InitAdminRoleScope 初始系统管理员角色作用域（global）
const InitAdminRoleScope = "global"

// InitPermission 初始化时创建的一条权限（access.permissions 表）
type InitPermission struct {
	Key         string // permission_key，如 directory.users.read、directory.*
	Name        string // 显示名
	Description string // 描述
}

// InitPermissions 通配权限，仅用于赋给 system.admin（bootstrap 时创建并赋给该角色）
// 细粒度权限（permissions 表其他行）、action_requirements（action 需要哪些 permission）均由前端/管理后台后续配置
var InitPermissions = []InitPermission{
	{Key: "system.*", Name: "System", Description: "All system-level permissions"},
	{Key: "directory.*", Name: "Directory", Description: "All directory permissions"},
	{Key: "tenants.*", Name: "Tenants", Description: "All tenants permissions"},
	{Key: "access.*", Name: "Access", Description: "All access permissions"},
	{Key: "auth.*", Name: "Auth", Description: "All auth permissions"},
	{Key: "clients.*", Name: "Clients", Description: "All clients permissions"},
	{Key: "audit.*", Name: "Audit", Description: "All audit permissions"},
	{Key: "image.*", Name: "Image", Description: "All image permissions"},
}

// InitAction 初始化时注册的一条 action（access.actions 表）
// action_key 格式：service:resource.verb，如 directory:user.get、directory:users.count
// 请求带来的 action_key 必须在 actions 表中存在且 status=active，否则 ACTION_NOT_CONFIGURED / 拼错
type InitAction struct {
	ActionKey string // 如 directory:user.get
	Service   string // 如 directory（与 action_key 前缀一致）
	Name      string // 显示名，如 "User Get"
}

// InitActions 系统注册的所有 action；CheckAccess 前需在 actions 表查到该 key
var InitActions = []InitAction{
	// directory - user
	{ActionKey: "directory:user.get", Service: "directory", Name: "User Get"},
	{ActionKey: "directory:user.create", Service: "directory", Name: "User Create"},
	{ActionKey: "directory:user.update", Service: "directory", Name: "User Update"},
	{ActionKey: "directory:user.delete", Service: "directory", Name: "User Delete"},
	{ActionKey: "directory:users.count", Service: "directory", Name: "Users Count"},
	// directory - user_email
	{ActionKey: "directory:user_email.get", Service: "directory", Name: "User Email Get"},
	{ActionKey: "directory:user_email.create", Service: "directory", Name: "User Email Create"},
	{ActionKey: "directory:user_email.update", Service: "directory", Name: "User Email Update"},
	{ActionKey: "directory:user_email.delete", Service: "directory", Name: "User Email Delete"},
	// directory - user_phone
	{ActionKey: "directory:user_phone.get", Service: "directory", Name: "User Phone Get"},
	{ActionKey: "directory:user_phone.create", Service: "directory", Name: "User Phone Create"},
	{ActionKey: "directory:user_phone.update", Service: "directory", Name: "User Phone Update"},
	{ActionKey: "directory:user_phone.delete", Service: "directory", Name: "User Phone Delete"},
	// directory - user_education / user_occupation / user_preference / user_profile / user_avatar / user_image
	{ActionKey: "directory:user_education.get", Service: "directory", Name: "User Education Get"},
	{ActionKey: "directory:user_education.create", Service: "directory", Name: "User Education Create"},
	{ActionKey: "directory:user_education.update", Service: "directory", Name: "User Education Update"},
	{ActionKey: "directory:user_education.delete", Service: "directory", Name: "User Education Delete"},
	{ActionKey: "directory:user_occupation.get", Service: "directory", Name: "User Occupation Get"},
	{ActionKey: "directory:user_occupation.create", Service: "directory", Name: "User Occupation Create"},
	{ActionKey: "directory:user_occupation.update", Service: "directory", Name: "User Occupation Update"},
	{ActionKey: "directory:user_occupation.delete", Service: "directory", Name: "User Occupation Delete"},
	{ActionKey: "directory:user_preference.get", Service: "directory", Name: "User Preference Get"},
	{ActionKey: "directory:user_preference.create", Service: "directory", Name: "User Preference Create"},
	{ActionKey: "directory:user_preference.update", Service: "directory", Name: "User Preference Update"},
	{ActionKey: "directory:user_preference.delete", Service: "directory", Name: "User Preference Delete"},
	{ActionKey: "directory:user_profile.get", Service: "directory", Name: "User Profile Get"},
	{ActionKey: "directory:user_profile.create", Service: "directory", Name: "User Profile Create"},
	{ActionKey: "directory:user_profile.update", Service: "directory", Name: "User Profile Update"},
	{ActionKey: "directory:user_profile.delete", Service: "directory", Name: "User Profile Delete"},
	{ActionKey: "directory:user_avatar.get", Service: "directory", Name: "User Avatar Get"},
	{ActionKey: "directory:user_avatar.create", Service: "directory", Name: "User Avatar Create"},
	{ActionKey: "directory:user_avatar.update", Service: "directory", Name: "User Avatar Update"},
	{ActionKey: "directory:user_avatar.delete", Service: "directory", Name: "User Avatar Delete"},
	{ActionKey: "directory:user_image.get", Service: "directory", Name: "User Image Get"},
	{ActionKey: "directory:user_image.create", Service: "directory", Name: "User Image Create"},
	{ActionKey: "directory:user_image.update", Service: "directory", Name: "User Image Update"},
	{ActionKey: "directory:user_image.delete", Service: "directory", Name: "User Image Delete"},
	// access
	{ActionKey: "access:grant.get", Service: "access", Name: "Grant Get"},
	{ActionKey: "access:role.get", Service: "access", Name: "Role Get"},
	{ActionKey: "access:permission.get", Service: "access", Name: "Permission Get"},
	{ActionKey: "access:role_permission.get", Service: "access", Name: "Role Permission Get"},
	// image
	{ActionKey: "image:upload.create", Service: "image", Name: "Upload Create"},
}

// 细粒度权限（permissions 表除上述通配外的行）、action_requirements（每个 action 需要哪些 permission）均由前端/管理后台配置，不在此 init 中写死。
// Action 由路由注册时通过 withAction(actionKey) 绑定，中间件从 context 取 action_key 再调 CheckAccess。
