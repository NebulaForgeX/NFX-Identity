/*
 * prettier-ignore
 * This file contains constants for the backend API endpoints and WebSocket URLs.
 * API endpoints are based on NFX-ID routes.
 */

// 从环境变量获取配置
// 通过 Traefik 反向代理访问：10187 是 Traefik 的 HTTP 端口
const HTTP_BASE_URL = import.meta.env.VITE_API_URL || "http://localhost:10187";
const WS_BASE_URL = import.meta.env.VITE_WS_URL || "ws://localhost:10187";

// 图片服务 URL（独立配置）
// 通过 Traefik 反向代理访问
const IMAGE_BASE_URL =
  import.meta.env.VITE_IMAGE_URL ||
  (import.meta.env.VITE_BUILD_ENV === "prod" ? "https://localhost:10187" : "http://localhost:10187");

// API路由定义 - 通过 Traefik 反向代理，路径前缀为 /auth, /image, /permission
export const URL_PATHS = {
  // Auth 模块 - /auth
  AUTH: {
    // 公开路由
    LOGIN: "/auth/login",
    REFRESH_TOKEN: "/auth/refresh",
    SEND_VERIFICATION_CODE: "/auth/verification-code", // 发送验证码
    // 用户相关 - 需要认证
    CREATE_USER: "/auth/users",
    GET_USERS: "/auth/users",
    GET_USER: "/auth/users/:id",
    UPDATE_USER: "/auth/users/:id",
    DELETE_USER: "/auth/users/:id",
    DELETE_USER_ACCOUNT: "/auth/users/:id/account",
    // 资料相关 - 需要认证
    CREATE_PROFILE: "/auth/profiles",
    GET_PROFILES: "/auth/profiles",
    GET_PROFILE: "/auth/profiles/:id",
    GET_PROFILE_BY_USER_ID: "/auth/profiles/user/:user_id",
    UPDATE_PROFILE: "/auth/profiles/:id",
    DELETE_PROFILE: "/auth/profiles/:id",
    // 角色相关 - 需要认证
    CREATE_ROLE: "/auth/roles",
    GET_ROLES: "/auth/roles",
    GET_ROLE: "/auth/roles/:id",
    GET_ROLE_BY_NAME: "/auth/roles/name/:name",
    UPDATE_ROLE: "/auth/roles/:id",
    DELETE_ROLE: "/auth/roles/:id",
    // 徽章相关 - 需要认证
    CREATE_BADGE: "/auth/badges",
    GET_BADGES: "/auth/badges",
    GET_BADGE: "/auth/badges/:id",
    GET_BADGE_BY_NAME: "/auth/badges/name/:name",
    UPDATE_BADGE: "/auth/badges/:id",
    DELETE_BADGE: "/auth/badges/:id",
    // 教育经历相关 - 需要认证
    CREATE_EDUCATION: "/auth/educations",
    GET_EDUCATIONS: "/auth/educations",
    GET_EDUCATION: "/auth/educations/:id",
    GET_EDUCATIONS_BY_PROFILE_ID: "/auth/educations/profile/:profile_id",
    UPDATE_EDUCATION: "/auth/educations/:id",
    DELETE_EDUCATION: "/auth/educations/:id",
    // 职业信息相关 - 需要认证
    CREATE_OCCUPATION: "/auth/occupations",
    GET_OCCUPATIONS: "/auth/occupations",
    GET_OCCUPATION: "/auth/occupations/:id",
    GET_OCCUPATIONS_BY_PROFILE_ID: "/auth/occupations/profile/:profile_id",
    UPDATE_OCCUPATION: "/auth/occupations/:id",
    DELETE_OCCUPATION: "/auth/occupations/:id",
    // 用户徽章关联相关 - 需要认证
    CREATE_PROFILE_BADGE: "/auth/profile-badges",
    GET_PROFILE_BADGE: "/auth/profile-badges/:id",
    GET_PROFILE_BADGES_BY_PROFILE_ID: "/auth/profile-badges/profile/:profile_id",
    GET_PROFILE_BADGES_BY_BADGE_ID: "/auth/profile-badges/badge/:badge_id",
    UPDATE_PROFILE_BADGE: "/auth/profile-badges/:id",
    DELETE_PROFILE_BADGE: "/auth/profile-badges/:id",
    DELETE_PROFILE_BADGE_BY_PROFILE_AND_BADGE: "/auth/profile-badges/profile/:profile_id/badge/:badge_id",
  },
  // Image 模块 - /image
  IMAGE: {
    // 图片相关
    CREATE_IMAGE: "/image/images",
    GET_IMAGES: "/image/images",
    GET_IMAGE: "/image/images/:id",
    UPDATE_IMAGE: "/image/images/:id",
    DELETE_IMAGE: "/image/images/:id",
    // 图片类型相关
    CREATE_IMAGE_TYPE: "/image/image-types",
    GET_IMAGE_TYPES: "/image/image-types",
    GET_IMAGE_TYPE: "/image/image-types/:id",
    GET_IMAGE_TYPE_BY_KEY: "/image/image-types/key/:key",
    UPDATE_IMAGE_TYPE: "/image/image-types/:id",
    DELETE_IMAGE_TYPE: "/image/image-types/:id",
  },
  // Permission 模块 - /permission
  PERMISSION: {
    // 公开路由
    LOGIN: "/permission/login",
    REGISTER: "/permission/register", // 注册
    // 权限管理 - 需要认证
    CREATE_PERMISSION: "/permission/permissions",
    GET_PERMISSIONS: "/permission/permissions",
    GET_PERMISSION: "/permission/permissions/:id",
    GET_PERMISSION_BY_TAG: "/permission/permissions/tag/:tag",
    UPDATE_PERMISSION: "/permission/permissions/:id",
    DELETE_PERMISSION: "/permission/permissions/:id",
    // 用户权限管理 - 需要认证
    ASSIGN_USER_PERMISSION: "/permission/user-permissions",
    REVOKE_USER_PERMISSION: "/permission/user-permissions",
    GET_USER_PERMISSIONS: "/permission/users/:user_id/permissions",
    GET_USER_PERMISSION_TAGS: "/permission/users/:user_id/permission-tags",
    CHECK_USER_PERMISSION: "/permission/user-permissions/check",
    // 授权码管理 - 需要认证
    CREATE_AUTHORIZATION_CODE: "/permission/authorization-codes",
    GET_AUTHORIZATION_CODE: "/permission/authorization-codes/:id",
    GET_AUTHORIZATION_CODE_BY_CODE: "/permission/authorization-codes/code/:code",
    USE_AUTHORIZATION_CODE: "/permission/authorization-codes/use",
    DELETE_AUTHORIZATION_CODE: "/permission/authorization-codes/:id",
    ACTIVATE_AUTHORIZATION_CODE: "/permission/authorization-codes/:id/activate",
    DEACTIVATE_AUTHORIZATION_CODE: "/permission/authorization-codes/:id/deactivate",
  },
} as const;

export const API_ENDPOINTS = {
  PURE: HTTP_BASE_URL,
  WS: WS_BASE_URL,
  IMAGE: IMAGE_BASE_URL, // 图片服务基础 URL
} as const;

// 类型定义
export type URL_PATHS_TYPE = typeof URL_PATHS;
export type API_ENDPOINTS_TYPE = typeof API_ENDPOINTS;

