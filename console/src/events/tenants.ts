export const tenantsEvents = {
  // Tenant 相关
  INVALIDATE_TENANT: "TENANTS:INVALIDATE_TENANT",
  INVALIDATE_TENANTS: "TENANTS:INVALIDATE_TENANTS",

  // Group 相关
  INVALIDATE_GROUP: "TENANTS:INVALIDATE_GROUP",
  INVALIDATE_GROUPS: "TENANTS:INVALIDATE_GROUPS",

  // Member 相关
  INVALIDATE_MEMBER: "TENANTS:INVALIDATE_MEMBER",
  INVALIDATE_MEMBERS: "TENANTS:INVALIDATE_MEMBERS",

  // Invitation 相关
  INVALIDATE_INVITATION: "TENANTS:INVALIDATE_INVITATION",
  INVALIDATE_INVITATIONS: "TENANTS:INVALIDATE_INVITATIONS",

  // TenantApp 相关
  INVALIDATE_TENANT_APP: "TENANTS:INVALIDATE_TENANT_APP",
  INVALIDATE_TENANT_APPS: "TENANTS:INVALIDATE_TENANT_APPS",

  // TenantSetting 相关
  INVALIDATE_TENANT_SETTING: "TENANTS:INVALIDATE_TENANT_SETTING",
  INVALIDATE_TENANT_SETTINGS: "TENANTS:INVALIDATE_TENANT_SETTINGS",

  // DomainVerification 相关
  INVALIDATE_DOMAIN_VERIFICATION: "TENANTS:INVALIDATE_DOMAIN_VERIFICATION",
  INVALIDATE_DOMAIN_VERIFICATIONS: "TENANTS:INVALIDATE_DOMAIN_VERIFICATIONS",

  // MemberRole 相关
  INVALIDATE_MEMBER_ROLE: "TENANTS:INVALIDATE_MEMBER_ROLE",
  INVALIDATE_MEMBER_ROLES: "TENANTS:INVALIDATE_MEMBER_ROLES",

  // MemberGroup 相关
  INVALIDATE_MEMBER_GROUP: "TENANTS:INVALIDATE_MEMBER_GROUP",
  INVALIDATE_MEMBER_GROUPS: "TENANTS:INVALIDATE_MEMBER_GROUPS",

  // MemberAppRole 相关
  INVALIDATE_MEMBER_APP_ROLE: "TENANTS:INVALIDATE_MEMBER_APP_ROLE",
  INVALIDATE_MEMBER_APP_ROLES: "TENANTS:INVALIDATE_MEMBER_APP_ROLES",
} as const;

type TenantsEvent = (typeof tenantsEvents)[keyof typeof tenantsEvents];

class TenantsEventEmitter {
  private listeners: Record<TenantsEvent, Set<Function>> = {
    [tenantsEvents.INVALIDATE_TENANT]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_TENANTS]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_GROUP]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_GROUPS]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_MEMBER]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_MEMBERS]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_INVITATION]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_INVITATIONS]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_TENANT_APP]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_TENANT_APPS]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_TENANT_SETTING]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_TENANT_SETTINGS]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_DOMAIN_VERIFICATION]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_DOMAIN_VERIFICATIONS]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_MEMBER_ROLE]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_MEMBER_ROLES]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_MEMBER_GROUP]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_MEMBER_GROUPS]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_MEMBER_APP_ROLE]: new Set<Function>(),
    [tenantsEvents.INVALIDATE_MEMBER_APP_ROLES]: new Set<Function>(),
  };

  on(event: TenantsEvent, callback: Function) {
    this.listeners[event].add(callback);
  }

  off(event: TenantsEvent, callback: Function) {
    this.listeners[event].delete(callback);
  }

  emit(event: TenantsEvent, ...args: unknown[]) {
    this.listeners[event].forEach((callback) => {
      callback(...args);
    });
  }
}

export const tenantsEventEmitter = new TenantsEventEmitter();
