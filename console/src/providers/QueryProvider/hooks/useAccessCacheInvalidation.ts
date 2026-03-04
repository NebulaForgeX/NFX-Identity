import type { QueryClient } from "@tanstack/react-query";
import { accessEventEmitter, accessEvents } from "@/events/access";
import {
  ACCESS_ACTION,
  ACCESS_ACTION_LIST,
  ACCESS_ACTION_REQUIREMENT,
  ACCESS_ACTION_REQUIREMENT_LIST,
  ACCESS_GRANT,
  ACCESS_GRANT_LIST,
  ACCESS_PERMISSION,
  ACCESS_PERMISSION_LIST,
  ACCESS_ROLE,
  ACCESS_ROLE_LIST,
  ACCESS_ROLE_PERMISSION,
  ACCESS_ROLE_PERMISSION_LIST,
  ACCESS_SCOPE,
  ACCESS_SCOPE_LIST,
  ACCESS_SCOPE_PERMISSION,
  ACCESS_SCOPE_PERMISSION_LIST,
} from "@/constants";

type EventCb = (...args: unknown[]) => void;

/**
 * Access 相关的缓存失效事件处理
 */
export const useAccessCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateRoles = () => queryClient.invalidateQueries({ queryKey: ACCESS_ROLE_LIST });
  const handleInvalidateRole = (item: string) => queryClient.invalidateQueries({ queryKey: ACCESS_ROLE(item) });
  const handleInvalidatePermissions = () => queryClient.invalidateQueries({ queryKey: ACCESS_PERMISSION_LIST });
  const handleInvalidatePermission = (item: string) => queryClient.invalidateQueries({ queryKey: ACCESS_PERMISSION(item) });
  const handleInvalidateScopes = () => queryClient.invalidateQueries({ queryKey: ACCESS_SCOPE_LIST });
  const handleInvalidateScope = (item: string) => queryClient.invalidateQueries({ queryKey: ACCESS_SCOPE(item) });
  const handleInvalidateGrants = () => queryClient.invalidateQueries({ queryKey: ACCESS_GRANT_LIST });
  const handleInvalidateGrant = (item: string) => queryClient.invalidateQueries({ queryKey: ACCESS_GRANT(item) });
  const handleInvalidateRolePermissions = () => queryClient.invalidateQueries({ queryKey: ACCESS_ROLE_PERMISSION_LIST });
  const handleInvalidateRolePermission = (item: string) => queryClient.invalidateQueries({ queryKey: ACCESS_ROLE_PERMISSION(item) });
  const handleInvalidateScopePermissions = () => queryClient.invalidateQueries({ queryKey: ACCESS_SCOPE_PERMISSION_LIST });
  const handleInvalidateScopePermission = (item: string) => queryClient.invalidateQueries({ queryKey: ACCESS_SCOPE_PERMISSION(item) });

  accessEventEmitter.on(accessEvents.INVALIDATE_ROLES, handleInvalidateRoles as EventCb);
  accessEventEmitter.on(accessEvents.INVALIDATE_ROLE, handleInvalidateRole as EventCb);
  accessEventEmitter.on(accessEvents.INVALIDATE_PERMISSIONS, handleInvalidatePermissions as EventCb);
  accessEventEmitter.on(accessEvents.INVALIDATE_PERMISSION, handleInvalidatePermission as EventCb);
  accessEventEmitter.on(accessEvents.INVALIDATE_SCOPES, handleInvalidateScopes as EventCb);
  accessEventEmitter.on(accessEvents.INVALIDATE_SCOPE, handleInvalidateScope as EventCb);
  accessEventEmitter.on(accessEvents.INVALIDATE_GRANTS, handleInvalidateGrants as EventCb);
  accessEventEmitter.on(accessEvents.INVALIDATE_GRANT, handleInvalidateGrant as EventCb);
  accessEventEmitter.on(accessEvents.INVALIDATE_ROLE_PERMISSIONS, handleInvalidateRolePermissions as EventCb);
  accessEventEmitter.on(accessEvents.INVALIDATE_ROLE_PERMISSION, handleInvalidateRolePermission as EventCb);
  accessEventEmitter.on(accessEvents.INVALIDATE_SCOPE_PERMISSIONS, handleInvalidateScopePermissions as EventCb);
  accessEventEmitter.on(accessEvents.INVALIDATE_SCOPE_PERMISSION, handleInvalidateScopePermission as EventCb);

  return () => {
    accessEventEmitter.off(accessEvents.INVALIDATE_ROLES, handleInvalidateRoles as EventCb);
    accessEventEmitter.off(accessEvents.INVALIDATE_ROLE, handleInvalidateRole as EventCb);
    accessEventEmitter.off(accessEvents.INVALIDATE_PERMISSIONS, handleInvalidatePermissions as EventCb);
    accessEventEmitter.off(accessEvents.INVALIDATE_PERMISSION, handleInvalidatePermission as EventCb);
    accessEventEmitter.off(accessEvents.INVALIDATE_SCOPES, handleInvalidateScopes as EventCb);
    accessEventEmitter.off(accessEvents.INVALIDATE_SCOPE, handleInvalidateScope as EventCb);
    accessEventEmitter.off(accessEvents.INVALIDATE_GRANTS, handleInvalidateGrants as EventCb);
    accessEventEmitter.off(accessEvents.INVALIDATE_GRANT, handleInvalidateGrant as EventCb);
    accessEventEmitter.off(accessEvents.INVALIDATE_ROLE_PERMISSIONS, handleInvalidateRolePermissions as EventCb);
    accessEventEmitter.off(accessEvents.INVALIDATE_ROLE_PERMISSION, handleInvalidateRolePermission as EventCb);
    accessEventEmitter.off(accessEvents.INVALIDATE_SCOPE_PERMISSIONS, handleInvalidateScopePermissions as EventCb);
    accessEventEmitter.off(accessEvents.INVALIDATE_SCOPE_PERMISSION, handleInvalidateScopePermission as EventCb);
  };
};
