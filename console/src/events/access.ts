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

  // Action 相关
  INVALIDATE_ACTION: "ACCESS:INVALIDATE_ACTION",
  INVALIDATE_ACTIONS: "ACCESS:INVALIDATE_ACTIONS",

  // ActionRequirement 相关
  INVALIDATE_ACTION_REQUIREMENT: "ACCESS:INVALIDATE_ACTION_REQUIREMENT",
  INVALIDATE_ACTION_REQUIREMENTS: "ACCESS:INVALIDATE_ACTION_REQUIREMENTS",
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
    [accessEvents.INVALIDATE_ACTION]: new Set<Function>(),
    [accessEvents.INVALIDATE_ACTIONS]: new Set<Function>(),
    [accessEvents.INVALIDATE_ACTION_REQUIREMENT]: new Set<Function>(),
    [accessEvents.INVALIDATE_ACTION_REQUIREMENTS]: new Set<Function>(),
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
