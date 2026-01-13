export const accessEvents = {
  // Role 相关
  INVALIDATE_ROLE: "ACCESS:INVALIDATE_ROLE",
  INVALIDATE_ROLES: "ACCESS:INVALIDATE_ROLES",

  // Permission 相关
  INVALIDATE_PERMISSION: "ACCESS:INVALIDATE_PERMISSION",
  INVALIDATE_PERMISSIONS: "ACCESS:INVALIDATE_PERMISSIONS",

  // Scope 相关
  INVALIDATE_SCOPE: "ACCESS:INVALIDATE_SCOPE",
  INVALIDATE_SCOPES: "ACCESS:INVALIDATE_SCOPES",

  // Grant 相关
  INVALIDATE_GRANT: "ACCESS:INVALIDATE_GRANT",
  INVALIDATE_GRANTS: "ACCESS:INVALIDATE_GRANTS",

  // RolePermission 相关
  INVALIDATE_ROLE_PERMISSION: "ACCESS:INVALIDATE_ROLE_PERMISSION",
  INVALIDATE_ROLE_PERMISSIONS: "ACCESS:INVALIDATE_ROLE_PERMISSIONS",

  // ScopePermission 相关
  INVALIDATE_SCOPE_PERMISSION: "ACCESS:INVALIDATE_SCOPE_PERMISSION",
  INVALIDATE_SCOPE_PERMISSIONS: "ACCESS:INVALIDATE_SCOPE_PERMISSIONS",
} as const;

type AccessEvent = (typeof accessEvents)[keyof typeof accessEvents];

class AccessEventEmitter {
  private listeners: Record<AccessEvent, Set<Function>> = {
    [accessEvents.INVALIDATE_ROLE]: new Set<Function>(),
    [accessEvents.INVALIDATE_ROLES]: new Set<Function>(),
    [accessEvents.INVALIDATE_PERMISSION]: new Set<Function>(),
    [accessEvents.INVALIDATE_PERMISSIONS]: new Set<Function>(),
    [accessEvents.INVALIDATE_SCOPE]: new Set<Function>(),
    [accessEvents.INVALIDATE_SCOPES]: new Set<Function>(),
    [accessEvents.INVALIDATE_GRANT]: new Set<Function>(),
    [accessEvents.INVALIDATE_GRANTS]: new Set<Function>(),
    [accessEvents.INVALIDATE_ROLE_PERMISSION]: new Set<Function>(),
    [accessEvents.INVALIDATE_ROLE_PERMISSIONS]: new Set<Function>(),
    [accessEvents.INVALIDATE_SCOPE_PERMISSION]: new Set<Function>(),
    [accessEvents.INVALIDATE_SCOPE_PERMISSIONS]: new Set<Function>(),
  };

  on(event: AccessEvent, callback: Function) {
    this.listeners[event].add(callback);
  }

  off(event: AccessEvent, callback: Function) {
    this.listeners[event].delete(callback);
  }

  emit(event: AccessEvent, ...args: unknown[]) {
    this.listeners[event].forEach((callback) => {
      callback(...args);
    });
  }
}

export const accessEventEmitter = new AccessEventEmitter();
