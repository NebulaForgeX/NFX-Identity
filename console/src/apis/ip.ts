/*
 * API path tree: base paths + dynamic segments via path().
 * path("/a", { b: "/b", c: path("/c", { d: "/d" }) }) => b="/a/b", c="/a/c", c.d="/a/c/d"
 */

const [BASE, CHILDREN] = [Symbol(), Symbol()];

// Recursive type: each key maps to string (leaf), function (dynamic path), or nested PathNode
type PathNode<T> = string & {
  [K in keyof T]: T[K] extends (...args: infer A) => infer R
    ? R extends object
      ? (...args: A) => PathNode<R> // function returning path
      : (...args: A) => string // function returning string
    : T[K] extends object
      ? PathNode<T[K]>
      : string;
};

const path = <T extends Record<string, unknown>>(base: string, children: T): PathNode<T> =>
  Object.assign(
    Object.defineProperties(new String(base), {
      [BASE]: { value: base },
      [CHILDREN]: { value: children },
      toString: { value: () => base },
      valueOf: { value: () => base },
      [Symbol.toPrimitive]: { value: () => base },
    }),
    Object.fromEntries(
      Object.entries(children).map(([k, v]) => [
        k,
        typeof v === "function"
          ? (...args: unknown[]) => {
              const result = (v as (...args: unknown[]) => unknown)(...args);
              return BASE in Object(result)
                ? path(
                    `${base}${(result as Record<symbol, string>)[BASE]}`,
                    (result as Record<symbol, Record<string, unknown>>)[CHILDREN] ?? {},
                  )
                : `${base}${result}`;
            }
          : BASE in Object(v)
            ? path(
                `${base}${(v as Record<symbol, string>)[BASE]}`,
                (v as Record<symbol, Record<string, unknown>>)[CHILDREN] ?? {},
              )
            : `${base}${v}`,
      ]),
    ),
  ) as PathNode<T>;

// 从环境变量获取配置
const HTTP_BASE_URL = import.meta.env.VITE_API_URL || "http://localhost:10166";
const WS_BASE_URL = import.meta.env.VITE_WS_URL || "ws://localhost:10166";

export const URL_PATHS = {
  ACCESS: path("/access/auth", {
    roles: path("/roles", { byId: (id: string) => `/${id}`, byKey: (key: string) => `/key/${key}` }),
    permissions: path("/permissions", {
      byId: (id: string) => `/${id}`,
      byKey: (key: string) => `/key/${key}`,
    }),
    scopes: path("/scopes", { byScope: (scope: string) => `/${scope}` }),
    grants: path("/grants", { byId: (id: string) => `/${id}` }),
    rolePermissions: path("/role-permissions", {
      byRole: (roleId: string) => `/role/${roleId}`,
      byId: (id: string) => `/${id}`,
    }),
    scopePermissions: path("/scope-permissions", { byId: (id: string) => `/${id}` }),
    actions: path("/actions", { byId: (id: string) => `/${id}`, byKey: (key: string) => `/key/${key}` }),
    actionRequirements: path("/action-requirements", {
      byId: (id: string) => `/${id}`,
      byPermission: (permissionId: string) => `/permission/${permissionId}`,
    }),
  }),

  AUDIT: path("/audit/auth", {
    events: path("/events", { byId: (id: string) => `/${id}` }),
    actorSnapshots: path("/actor-snapshots", { byId: (id: string) => `/${id}` }),
    eventRetentionPolicies: path("/event-retention-policies", { byId: (id: string) => `/${id}` }),
    eventSearchIndex: path("/event-search-index", { byId: (id: string) => `/${id}` }),
    hashChainCheckpoints: path("/hash-chain-checkpoints", { byId: (id: string) => `/${id}` }),
  }),

  AUTH: {
    login: path("/auth/login", { email: "/email", phone: "/phone" }),
    refresh: "/auth/refresh",
    sendVerificationCode: "/auth/send-verification-code",
    signup: "/auth/signup",
    ...(path("/auth/auth", {
      sessions: path("/sessions", {
        byId: (id: string) => `/${id}`,
        revoke: (sessionId: string) => `/${sessionId}/revoke`,
      }),
      userCredentials: path("/user-credentials", { byId: (id: string) => `/${id}` }),
      mfaFactors: path("/mfa-factors", { byId: (id: string) => `/${id}` }),
      refreshTokens: path("/refresh-tokens", { byId: (id: string) => `/${id}` }),
      passwordResets: path("/password-resets", { byId: (id: string) => `/${id}` }),
      passwordHistory: path("/password-history", { byId: (id: string) => `/${id}` }),
      loginAttempts: path("/login-attempts", { byId: (id: string) => `/${id}` }),
      accountLockouts: path("/account-lockouts", { byId: (id: string) => `/${id}` }),
      trustedDevices: path("/trusted-devices", { byId: (id: string) => `/${id}` }),
    }) as unknown as Record<string, unknown>),
  },

  CLIENTS: path("/clients/auth", {
    apps: path("/apps", {
      byId: (id: string) => `/${id}`,
      byAppId: (appId: string) => `/app-id/${appId}`,
    }),
    apiKeys: path("/api-keys", {
      byId: (id: string) => `/${id}`,
      byKeyId: (keyId: string) => `/key-id/${keyId}`,
    }),
    clientCredentials: path("/client-credentials", {
      byId: (id: string) => `/${id}`,
      byClientId: (clientId: string) => `/client-id/${clientId}`,
    }),
    clientScopes: path("/client-scopes", { byId: (id: string) => `/${id}` }),
    ipAllowlist: path("/ip-allowlist", {
      byId: (id: string) => `/${id}`,
      byRuleId: (ruleId: string) => `/rule-id/${ruleId}`,
    }),
    rateLimits: path("/rate-limits", { byId: (id: string) => `/${id}` }),
  }),

  DIRECTORY: path("/directory/auth", {
    users: path("/users", {
      byId: (userId: string) =>
        path(`/${userId}`, {
          status: "/status",
          username: "/username",
          verify: "/verify",
          userEmails: "/user-emails",
          userPhones: "/user-phones",
          userEducations: "/user-educations",
          userOccupations: "/user-occupations",
          userImages: "/user-images",
          currentImage: "/user-images/current",
          imagesDisplayOrder: "/user-images/display-order",
        }),
      byUsername: (username: string) => `/username/${username}`,
    }),
    badges: path("/badges", { byId: (id: string) => `/${id}`, byName: (name: string) => `/name/${name}` }),
    userBadges: path("/user-badges", { byId: (id: string) => `/${id}` }),
    userEducations: path("/user-educations", { byId: (id: string) => `/${id}` }),
    userEmails: path("/user-emails", {
      byId: (id: string) => `/${id}`,
      setPrimary: (id: string) => `/${id}/set-primary`,
      verify: (id: string) => `/${id}/verify`,
    }),
    userOccupations: path("/user-occupations", { byId: (id: string) => `/${id}` }),
    userPhones: path("/user-phones", {
      byId: (id: string) => `/${id}`,
      setPrimary: (id: string) => `/${id}/set-primary`,
      verify: (id: string) => `/${id}/verify`,
    }),
    userPreferences: path("/user-preferences", { byId: (id: string) => `/${id}` }),
    userProfiles: path("/user-profiles", { byId: (id: string) => `/${id}` }),
    userAvatars: path("/user-avatars", { byUserId: (userId: string) => `/user/${userId}` }),
    userImages: path("/user-images", {
      byId: (id: string) => `/${id}`,
      setPrimary: (id: string) => `/${id}/set-primary`,
      displayOrder: (id: string) => `/${id}/display-order`,
    }),
  }),

  IMAGE: path("/image/auth", {
    upload: "/upload",
    images: path("/images", {
      byId: (id: string) => `/${id}`,
      move: (id: string) => `/${id}/move`,
    }),
    imageTypes: path("/image-types", { byId: (id: string) => `/${id}` }),
    imageVariants: path("/image-variants", { byId: (id: string) => `/${id}` }),
    imageTags: path("/image-tags", { byId: (id: string) => `/${id}` }),
  }),

  SYSTEM: {
    systemState: path("/system/system-state", {
      latest: "/latest",
      initialize: "/initialize",
      byId: (id: string) => `/${id}`,
    }),
    systemStateAuth: path("/system/auth/system-state", {
      byId: (id: string) => `/${id}`,
      initialize: "/initialize",
      reset: "/reset",
    }),
    i18nErrors: path("/system/i18n/errors", { byLang: (lang: string) => `/${lang}` }),
  },

  TENANTS: path("/tenants/auth", {
    tenants: path("/", {
      byId: (id: string) => `/${id}`,
      byTenantId: (tenantId: string) => `/tenant-id/${tenantId}`,
      status: (id: string) => `/${id}/status`,
    }),
    groups: path("/groups", { byId: (id: string) => `/${id}` }),
    members: path("/members", { byId: (id: string) => `/${id}` }),
    invitations: path("/invitations", {
      byId: (id: string) => `/${id}`,
      byInviteId: (inviteId: string) =>
        path(`/invite-id/${inviteId}`, {
          accept: "/accept",
          revoke: "/revoke",
        }),
    }),
    tenantApps: path("/tenant-apps", { byId: (id: string) => `/${id}` }),
    tenantSettings: path("/tenant-settings", { byId: (id: string) => `/${id}` }),
    domainVerifications: path("/domain-verifications", { byId: (id: string) => `/${id}` }),
    memberRoles: path("/member-roles", { byId: (id: string) => `/${id}`, revoke: (id: string) => `/${id}/revoke` }),
    memberGroups: path("/member-groups", { byId: (id: string) => `/${id}`, revoke: (id: string) => `/${id}/revoke` }),
    memberAppRoles: path("/member-app-roles", {
      byId: (id: string) => `/${id}`,
      revoke: (id: string) => `/${id}/revoke`,
    }),
  }),
} as const;

export const API_ENDPOINTS = {
  PURE: HTTP_BASE_URL,
  WS: WS_BASE_URL,
} as const;

export type URL_PATHS_TYPE = typeof URL_PATHS;
export type API_ENDPOINTS_TYPE = typeof API_ENDPOINTS;
