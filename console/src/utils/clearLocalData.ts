/**
 * 清除本地数据：authStore + localStorage 持久化 + 发送各模块 query invalidate 事件
 * 用于后端数据库重新初始化后，前端仍持有旧 userId 导致 404 时，让用户一键清除本地并回到登录页
 */

import { AuthStore } from "@/stores/authStore";
import { accessEventEmitter, accessEvents } from "@/events/access";
import { auditEventEmitter, auditEvents } from "@/events/audit";
import { authEventEmitter, authEvents } from "@/events/auth";
import { clientsEventEmitter, clientsEvents } from "@/events/clients";
import { directoryEventEmitter, directoryEvents } from "@/events/directory";
import { imageEventEmitter, imageEvents } from "@/events/image";
import { routerEventEmitter, routerEvents } from "@/events/router";
import { systemEventEmitter, systemEvents } from "@/events/system";
import { tenantsEventEmitter, tenantsEvents } from "@/events/tenants";

const AUTH_STORAGE_KEY = "auth-storage";

/**
 * 1. 清空 authStore 状态
 * 2. 移除 localStorage 中的 auth 持久化数据
 * 3. 发送各模块的 query invalidate 事件（列表级），使 React Query 缓存失效
 * 4. 可选：跳转到登录页
 */
export function clearLocalData(options?: { navigateToLogin?: boolean }) {
  // 1. 清空 auth store（内存状态）
  AuthStore.getState().clearAuth();

  // 2. 清除本地存储中的 auth 持久化（zustand persist 默认 key）
  try {
    localStorage.removeItem(AUTH_STORAGE_KEY);
    sessionStorage.removeItem(AUTH_STORAGE_KEY);
  } catch {
    // ignore
  }

  // 3. 发送各模块的 invalidate 事件（列表级），触发 queryClient.invalidateQueries
  directoryEventEmitter.emit(directoryEvents.INVALIDATE_USERS);
  directoryEventEmitter.emit(directoryEvents.INVALIDATE_BADGES);
  directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_BADGES);
  directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EDUCATIONS);
  directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_EMAILS);
  directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_OCCUPATIONS);
  directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PHONES);
  directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PREFERENCES);
  directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_PROFILES);
  directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_AVATARS);
  directoryEventEmitter.emit(directoryEvents.INVALIDATE_USER_IMAGES);

  accessEventEmitter.emit(accessEvents.INVALIDATE_ROLES);
  accessEventEmitter.emit(accessEvents.INVALIDATE_PERMISSIONS);
  accessEventEmitter.emit(accessEvents.INVALIDATE_SCOPES);
  accessEventEmitter.emit(accessEvents.INVALIDATE_GRANTS);
  accessEventEmitter.emit(accessEvents.INVALIDATE_ROLE_PERMISSIONS);
  accessEventEmitter.emit(accessEvents.INVALIDATE_SCOPE_PERMISSIONS);
  accessEventEmitter.emit(accessEvents.INVALIDATE_ACTIONS);
  accessEventEmitter.emit(accessEvents.INVALIDATE_ACTION_REQUIREMENTS);

  authEventEmitter.emit(authEvents.INVALIDATE_SESSIONS);
  authEventEmitter.emit(authEvents.INVALIDATE_USER_CREDENTIALS);
  authEventEmitter.emit(authEvents.INVALIDATE_MFA_FACTORS);
  authEventEmitter.emit(authEvents.INVALIDATE_REFRESH_TOKENS);
  authEventEmitter.emit(authEvents.INVALIDATE_PASSWORD_RESETS);
  authEventEmitter.emit(authEvents.INVALIDATE_PASSWORD_HISTORIES);
  authEventEmitter.emit(authEvents.INVALIDATE_LOGIN_ATTEMPTS);
  authEventEmitter.emit(authEvents.INVALIDATE_ACCOUNT_LOCKOUTS);
  authEventEmitter.emit(authEvents.INVALIDATE_TRUSTED_DEVICES);

  systemEventEmitter.emit(systemEvents.INVALIDATE_SYSTEM_STATES);

  tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANTS);
  tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_GROUPS);
  tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBERS);
  tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_INVITATIONS);
  tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT_APPS);
  tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT_SETTINGS);
  tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_DOMAIN_VERIFICATIONS);
  tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_ROLES);
  tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_GROUPS);
  tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_APP_ROLES);

  imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGES);
  imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TYPES);
  imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_VARIANTS);
  imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TAGS);

  clientsEventEmitter.emit(clientsEvents.INVALIDATE_APPS);
  clientsEventEmitter.emit(clientsEvents.INVALIDATE_API_KEYS);
  clientsEventEmitter.emit(clientsEvents.INVALIDATE_CLIENT_CREDENTIALS);
  clientsEventEmitter.emit(clientsEvents.INVALIDATE_CLIENT_SCOPES);
  clientsEventEmitter.emit(clientsEvents.INVALIDATE_IP_ALLOWLISTS);
  clientsEventEmitter.emit(clientsEvents.INVALIDATE_RATE_LIMITS);

  auditEventEmitter.emit(auditEvents.INVALIDATE_EVENTS);
  auditEventEmitter.emit(auditEvents.INVALIDATE_ACTOR_SNAPSHOTS);
  auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_RETENTION_POLICIES);
  auditEventEmitter.emit(auditEvents.INVALIDATE_EVENT_SEARCH_INDICES);
  auditEventEmitter.emit(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINTS);

  // 4. 跳转到登录页
  if (options?.navigateToLogin !== false) {
    routerEventEmitter.emit(routerEvents.NAVIGATE_TO_LOGIN);
  }
}
