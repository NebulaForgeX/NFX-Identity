import type { QueryClient } from "@tanstack/react-query";
import { accessEventEmitter, accessEvents } from "@/events/access";
import { ACCESS_QUERY_KEY_PREFIXES } from "@/constants";

/**
 * Access 相关的缓存失效事件处理
 */
export const useAccessCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateRoles = () => queryClient.invalidateQueries({ queryKey: ACCESS_QUERY_KEY_PREFIXES.ROLES });
  const handleInvalidateRole = (item: string) => queryClient.invalidateQueries({ queryKey: [...ACCESS_QUERY_KEY_PREFIXES.ROLE, item] });
  const handleInvalidatePermissions = () => queryClient.invalidateQueries({ queryKey: ACCESS_QUERY_KEY_PREFIXES.PERMISSIONS });
  const handleInvalidatePermission = (item: string) => queryClient.invalidateQueries({ queryKey: [...ACCESS_QUERY_KEY_PREFIXES.PERMISSION, item] });
  const handleInvalidateScopes = () => queryClient.invalidateQueries({ queryKey: ACCESS_QUERY_KEY_PREFIXES.SCOPES });
  const handleInvalidateScope = (item: string) => queryClient.invalidateQueries({ queryKey: [...ACCESS_QUERY_KEY_PREFIXES.SCOPE, item] });
  const handleInvalidateGrants = () => queryClient.invalidateQueries({ queryKey: ACCESS_QUERY_KEY_PREFIXES.GRANTS });
  const handleInvalidateGrant = (item: string) => queryClient.invalidateQueries({ queryKey: [...ACCESS_QUERY_KEY_PREFIXES.GRANT, item] });
  const handleInvalidateRolePermissions = () => queryClient.invalidateQueries({ queryKey: ACCESS_QUERY_KEY_PREFIXES.ROLE_PERMISSIONS });
  const handleInvalidateRolePermission = (item: string) => queryClient.invalidateQueries({ queryKey: [...ACCESS_QUERY_KEY_PREFIXES.ROLE_PERMISSION, item] });
  const handleInvalidateScopePermissions = () => queryClient.invalidateQueries({ queryKey: ACCESS_QUERY_KEY_PREFIXES.SCOPE_PERMISSIONS });
  const handleInvalidateScopePermission = (item: string) => queryClient.invalidateQueries({ queryKey: [...ACCESS_QUERY_KEY_PREFIXES.SCOPE_PERMISSION, item] });

  // 注册监听器
  accessEventEmitter.on(accessEvents.INVALIDATE_ROLES, handleInvalidateRoles);
  accessEventEmitter.on(accessEvents.INVALIDATE_ROLE, handleInvalidateRole);
  accessEventEmitter.on(accessEvents.INVALIDATE_PERMISSIONS, handleInvalidatePermissions);
  accessEventEmitter.on(accessEvents.INVALIDATE_PERMISSION, handleInvalidatePermission);
  accessEventEmitter.on(accessEvents.INVALIDATE_SCOPES, handleInvalidateScopes);
  accessEventEmitter.on(accessEvents.INVALIDATE_SCOPE, handleInvalidateScope);
  accessEventEmitter.on(accessEvents.INVALIDATE_GRANTS, handleInvalidateGrants);
  accessEventEmitter.on(accessEvents.INVALIDATE_GRANT, handleInvalidateGrant);
  accessEventEmitter.on(accessEvents.INVALIDATE_ROLE_PERMISSIONS, handleInvalidateRolePermissions);
  accessEventEmitter.on(accessEvents.INVALIDATE_ROLE_PERMISSION, handleInvalidateRolePermission);
  accessEventEmitter.on(accessEvents.INVALIDATE_SCOPE_PERMISSIONS, handleInvalidateScopePermissions);
  accessEventEmitter.on(accessEvents.INVALIDATE_SCOPE_PERMISSION, handleInvalidateScopePermission);

  // 清理监听器
  return () => {
    accessEventEmitter.off(accessEvents.INVALIDATE_ROLES, handleInvalidateRoles);
    accessEventEmitter.off(accessEvents.INVALIDATE_ROLE, handleInvalidateRole);
    accessEventEmitter.off(accessEvents.INVALIDATE_PERMISSIONS, handleInvalidatePermissions);
    accessEventEmitter.off(accessEvents.INVALIDATE_PERMISSION, handleInvalidatePermission);
    accessEventEmitter.off(accessEvents.INVALIDATE_SCOPES, handleInvalidateScopes);
    accessEventEmitter.off(accessEvents.INVALIDATE_SCOPE, handleInvalidateScope);
    accessEventEmitter.off(accessEvents.INVALIDATE_GRANTS, handleInvalidateGrants);
    accessEventEmitter.off(accessEvents.INVALIDATE_GRANT, handleInvalidateGrant);
    accessEventEmitter.off(accessEvents.INVALIDATE_ROLE_PERMISSIONS, handleInvalidateRolePermissions);
    accessEventEmitter.off(accessEvents.INVALIDATE_ROLE_PERMISSION, handleInvalidateRolePermission);
    accessEventEmitter.off(accessEvents.INVALIDATE_SCOPE_PERMISSIONS, handleInvalidateScopePermissions);
    accessEventEmitter.off(accessEvents.INVALIDATE_SCOPE_PERMISSION, handleInvalidateScopePermission);
  };
};
