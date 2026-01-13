import { useEffect } from "react";
import type { QueryClient } from "@tanstack/react-query";

import { accessEventEmitter, accessEvents } from "@/events/access";
import { auditEventEmitter, auditEvents } from "@/events/audit";
import { authEventEmitter, authEvents } from "@/events/auth";
import { clientsEventEmitter, clientsEvents } from "@/events/clients";
import { directoryEventEmitter, directoryEvents } from "@/events/directory";
import { imageEventEmitter, imageEvents } from "@/events/image";
import { systemEventEmitter, systemEvents } from "@/events/system";
import { tenantsEventEmitter, tenantsEvents } from "@/events/tenants";

/**
 * Hook for handling all cache invalidation events via event emitters
 * 监听所有缓存失效事件并自动刷新对应的查询
 * @param queryClient - QueryClient 实例
 */
export const useCacheInvalidationEvents = (queryClient: QueryClient) => {

  useEffect(() => {
    // Helper function to create invalidate handlers
    const createInvalidateHandler = (key: string) => () => {
      queryClient.invalidateQueries({ queryKey: [key] });
    };

    const createInvalidateByIdHandler = (key: string) => (id: string) => {
      queryClient.invalidateQueries({ queryKey: [key, id] });
    };

    // ========== Auth Events ==========
    const handleInvalidateSessions = createInvalidateHandler("sessions");
    const handleInvalidateSession = createInvalidateByIdHandler("session");
    const handleInvalidateUserCredentials = createInvalidateHandler("userCredentials");
    const handleInvalidateUserCredential = createInvalidateByIdHandler("userCredential");
    const handleInvalidateMFAFactors = createInvalidateHandler("mfaFactors");
    const handleInvalidateMFAFactor = createInvalidateByIdHandler("mfaFactor");
    const handleInvalidateRefreshTokens = createInvalidateHandler("refreshTokens");
    const handleInvalidateRefreshToken = createInvalidateByIdHandler("refreshToken");
    const handleInvalidatePasswordResets = createInvalidateHandler("passwordResets");
    const handleInvalidatePasswordReset = createInvalidateByIdHandler("passwordReset");
    const handleInvalidatePasswordHistories = createInvalidateHandler("passwordHistories");
    const handleInvalidatePasswordHistory = createInvalidateByIdHandler("passwordHistory");
    const handleInvalidateLoginAttempts = createInvalidateHandler("loginAttempts");
    const handleInvalidateLoginAttempt = createInvalidateByIdHandler("loginAttempt");
    const handleInvalidateAccountLockouts = createInvalidateHandler("accountLockouts");
    const handleInvalidateAccountLockout = createInvalidateByIdHandler("accountLockout");
    const handleInvalidateTrustedDevices = createInvalidateHandler("trustedDevices");
    const handleInvalidateTrustedDevice = createInvalidateByIdHandler("trustedDevice");

    // ========== Access Events ==========
    const handleInvalidateRoles = createInvalidateHandler("roles");
    const handleInvalidateRole = createInvalidateByIdHandler("role");
    const handleInvalidatePermissions = createInvalidateHandler("permissions");
    const handleInvalidatePermission = createInvalidateByIdHandler("permission");
    const handleInvalidateScopes = createInvalidateHandler("scopes");
    const handleInvalidateScope = (scope: string) => {
      queryClient.invalidateQueries({ queryKey: ["scope", scope] });
    };
    const handleInvalidateGrants = createInvalidateHandler("grants");
    const handleInvalidateGrant = createInvalidateByIdHandler("grant");
    const handleInvalidateRolePermissions = createInvalidateHandler("rolePermissions");
    const handleInvalidateRolePermission = createInvalidateByIdHandler("rolePermission");
    const handleInvalidateScopePermissions = createInvalidateHandler("scopePermissions");
    const handleInvalidateScopePermission = createInvalidateByIdHandler("scopePermission");

    // ========== Audit Events ==========
    const handleInvalidateEvents = createInvalidateHandler("events");
    const handleInvalidateEvent = createInvalidateByIdHandler("event");
    const handleInvalidateActorSnapshots = createInvalidateHandler("actorSnapshots");
    const handleInvalidateActorSnapshot = createInvalidateByIdHandler("actorSnapshot");
    const handleInvalidateEventRetentionPolicies = createInvalidateHandler("eventRetentionPolicies");
    const handleInvalidateEventRetentionPolicy = createInvalidateByIdHandler("eventRetentionPolicy");
    const handleInvalidateEventSearchIndices = createInvalidateHandler("eventSearchIndices");
    const handleInvalidateEventSearchIndex = createInvalidateByIdHandler("eventSearchIndex");
    const handleInvalidateHashChainCheckpoints = createInvalidateHandler("hashChainCheckpoints");
    const handleInvalidateHashChainCheckpoint = createInvalidateByIdHandler("hashChainCheckpoint");

    // ========== Clients Events ==========
    const handleInvalidateApps = createInvalidateHandler("apps");
    const handleInvalidateApp = createInvalidateByIdHandler("app");
    const handleInvalidateAPIKeys = createInvalidateHandler("apiKeys");
    const handleInvalidateAPIKey = createInvalidateByIdHandler("apiKey");
    const handleInvalidateClientCredentials = createInvalidateHandler("clientCredentials");
    const handleInvalidateClientCredential = createInvalidateByIdHandler("clientCredential");
    const handleInvalidateClientScopes = createInvalidateHandler("clientScopes");
    const handleInvalidateClientScope = createInvalidateByIdHandler("clientScope");
    const handleInvalidateIPAllowlists = createInvalidateHandler("ipAllowlists");
    const handleInvalidateIPAllowlist = createInvalidateByIdHandler("ipAllowlist");
    const handleInvalidateRateLimits = createInvalidateHandler("rateLimits");
    const handleInvalidateRateLimit = createInvalidateByIdHandler("rateLimit");

    // ========== Directory Events ==========
    const handleInvalidateUsers = createInvalidateHandler("users");
    const handleInvalidateUser = createInvalidateByIdHandler("user");
    const handleInvalidateBadges = createInvalidateHandler("badges");
    const handleInvalidateBadge = createInvalidateByIdHandler("badge");
    const handleInvalidateUserBadges = createInvalidateHandler("userBadges");
    const handleInvalidateUserBadge = createInvalidateByIdHandler("userBadge");
    const handleInvalidateUserEducations = createInvalidateHandler("userEducations");
    const handleInvalidateUserEducation = createInvalidateByIdHandler("userEducation");
    const handleInvalidateUserEmails = createInvalidateHandler("userEmails");
    const handleInvalidateUserEmail = createInvalidateByIdHandler("userEmail");
    const handleInvalidateUserOccupations = createInvalidateHandler("userOccupations");
    const handleInvalidateUserOccupation = createInvalidateByIdHandler("userOccupation");
    const handleInvalidateUserPhones = createInvalidateHandler("userPhones");
    const handleInvalidateUserPhone = createInvalidateByIdHandler("userPhone");
    const handleInvalidateUserPreferences = createInvalidateHandler("userPreferences");
    const handleInvalidateUserPreference = createInvalidateByIdHandler("userPreference");
    const handleInvalidateUserProfiles = createInvalidateHandler("userProfiles");
    const handleInvalidateUserProfile = createInvalidateByIdHandler("userProfile");

    // ========== Image Events ==========
    const handleInvalidateImages = createInvalidateHandler("images");
    const handleInvalidateImage = createInvalidateByIdHandler("image");
    const handleInvalidateImageTypes = createInvalidateHandler("imageTypes");
    const handleInvalidateImageType = createInvalidateByIdHandler("imageType");
    const handleInvalidateImageVariants = createInvalidateHandler("imageVariants");
    const handleInvalidateImageVariant = createInvalidateByIdHandler("imageVariant");
    const handleInvalidateImageTags = createInvalidateHandler("imageTags");
    const handleInvalidateImageTag = createInvalidateByIdHandler("imageTag");

    // ========== System Events ==========
    const handleInvalidateSystemStates = createInvalidateHandler("systemStates");
    const handleInvalidateSystemState = createInvalidateByIdHandler("systemState");

    // ========== Tenants Events ==========
    const handleInvalidateTenants = createInvalidateHandler("tenants");
    const handleInvalidateTenant = createInvalidateByIdHandler("tenant");
    const handleInvalidateGroups = createInvalidateHandler("groups");
    const handleInvalidateGroup = createInvalidateByIdHandler("group");
    const handleInvalidateMembers = createInvalidateHandler("members");
    const handleInvalidateMember = createInvalidateByIdHandler("member");
    const handleInvalidateInvitations = createInvalidateHandler("invitations");
    const handleInvalidateInvitation = createInvalidateByIdHandler("invitation");
    const handleInvalidateTenantApps = createInvalidateHandler("tenantApps");
    const handleInvalidateTenantApp = createInvalidateByIdHandler("tenantApp");
    const handleInvalidateTenantSettings = createInvalidateHandler("tenantSettings");
    const handleInvalidateTenantSetting = createInvalidateByIdHandler("tenantSetting");
    const handleInvalidateDomainVerifications = createInvalidateHandler("domainVerifications");
    const handleInvalidateDomainVerification = createInvalidateByIdHandler("domainVerification");
    const handleInvalidateMemberRoles = createInvalidateHandler("memberRoles");
    const handleInvalidateMemberRole = createInvalidateByIdHandler("memberRole");
    const handleInvalidateMemberGroups = createInvalidateHandler("memberGroups");
    const handleInvalidateMemberGroup = createInvalidateByIdHandler("memberGroup");
    const handleInvalidateMemberAppRoles = createInvalidateHandler("memberAppRoles");
    const handleInvalidateMemberAppRole = createInvalidateByIdHandler("memberAppRole");

    // ========== 注册所有监听器 ==========

    // Auth Events
    authEventEmitter.on(authEvents.INVALIDATE_SESSIONS, handleInvalidateSessions);
    authEventEmitter.on(authEvents.INVALIDATE_SESSION, handleInvalidateSession);
    authEventEmitter.on(authEvents.INVALIDATE_USER_CREDENTIALS, handleInvalidateUserCredentials);
    authEventEmitter.on(authEvents.INVALIDATE_USER_CREDENTIAL, handleInvalidateUserCredential);
    authEventEmitter.on(authEvents.INVALIDATE_MFA_FACTORS, handleInvalidateMFAFactors);
    authEventEmitter.on(authEvents.INVALIDATE_MFA_FACTOR, handleInvalidateMFAFactor);
    authEventEmitter.on(authEvents.INVALIDATE_REFRESH_TOKENS, handleInvalidateRefreshTokens);
    authEventEmitter.on(authEvents.INVALIDATE_REFRESH_TOKEN, handleInvalidateRefreshToken);
    authEventEmitter.on(authEvents.INVALIDATE_PASSWORD_RESETS, handleInvalidatePasswordResets);
    authEventEmitter.on(authEvents.INVALIDATE_PASSWORD_RESET, handleInvalidatePasswordReset);
    authEventEmitter.on(authEvents.INVALIDATE_PASSWORD_HISTORIES, handleInvalidatePasswordHistories);
    authEventEmitter.on(authEvents.INVALIDATE_PASSWORD_HISTORY, handleInvalidatePasswordHistory);
    authEventEmitter.on(authEvents.INVALIDATE_LOGIN_ATTEMPTS, handleInvalidateLoginAttempts);
    authEventEmitter.on(authEvents.INVALIDATE_LOGIN_ATTEMPT, handleInvalidateLoginAttempt);
    authEventEmitter.on(authEvents.INVALIDATE_ACCOUNT_LOCKOUTS, handleInvalidateAccountLockouts);
    authEventEmitter.on(authEvents.INVALIDATE_ACCOUNT_LOCKOUT, handleInvalidateAccountLockout);
    authEventEmitter.on(authEvents.INVALIDATE_TRUSTED_DEVICES, handleInvalidateTrustedDevices);
    authEventEmitter.on(authEvents.INVALIDATE_TRUSTED_DEVICE, handleInvalidateTrustedDevice);

    // Access Events
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

    // Audit Events
    auditEventEmitter.on(auditEvents.INVALIDATE_EVENTS, handleInvalidateEvents);
    auditEventEmitter.on(auditEvents.INVALIDATE_EVENT, handleInvalidateEvent);
    auditEventEmitter.on(auditEvents.INVALIDATE_ACTOR_SNAPSHOTS, handleInvalidateActorSnapshots);
    auditEventEmitter.on(auditEvents.INVALIDATE_ACTOR_SNAPSHOT, handleInvalidateActorSnapshot);
    auditEventEmitter.on(auditEvents.INVALIDATE_EVENT_RETENTION_POLICIES, handleInvalidateEventRetentionPolicies);
    auditEventEmitter.on(auditEvents.INVALIDATE_EVENT_RETENTION_POLICY, handleInvalidateEventRetentionPolicy);
    auditEventEmitter.on(auditEvents.INVALIDATE_EVENT_SEARCH_INDICES, handleInvalidateEventSearchIndices);
    auditEventEmitter.on(auditEvents.INVALIDATE_EVENT_SEARCH_INDEX, handleInvalidateEventSearchIndex);
    auditEventEmitter.on(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINTS, handleInvalidateHashChainCheckpoints);
    auditEventEmitter.on(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINT, handleInvalidateHashChainCheckpoint);

    // Clients Events
    clientsEventEmitter.on(clientsEvents.INVALIDATE_APPS, handleInvalidateApps);
    clientsEventEmitter.on(clientsEvents.INVALIDATE_APP, handleInvalidateApp);
    clientsEventEmitter.on(clientsEvents.INVALIDATE_API_KEYS, handleInvalidateAPIKeys);
    clientsEventEmitter.on(clientsEvents.INVALIDATE_API_KEY, handleInvalidateAPIKey);
    clientsEventEmitter.on(clientsEvents.INVALIDATE_CLIENT_CREDENTIALS, handleInvalidateClientCredentials);
    clientsEventEmitter.on(clientsEvents.INVALIDATE_CLIENT_CREDENTIAL, handleInvalidateClientCredential);
    clientsEventEmitter.on(clientsEvents.INVALIDATE_CLIENT_SCOPES, handleInvalidateClientScopes);
    clientsEventEmitter.on(clientsEvents.INVALIDATE_CLIENT_SCOPE, handleInvalidateClientScope);
    clientsEventEmitter.on(clientsEvents.INVALIDATE_IP_ALLOWLISTS, handleInvalidateIPAllowlists);
    clientsEventEmitter.on(clientsEvents.INVALIDATE_IP_ALLOWLIST, handleInvalidateIPAllowlist);
    clientsEventEmitter.on(clientsEvents.INVALIDATE_RATE_LIMITS, handleInvalidateRateLimits);
    clientsEventEmitter.on(clientsEvents.INVALIDATE_RATE_LIMIT, handleInvalidateRateLimit);

    // Directory Events
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USERS, handleInvalidateUsers);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER, handleInvalidateUser);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_BADGES, handleInvalidateBadges);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_BADGE, handleInvalidateBadge);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_BADGES, handleInvalidateUserBadges);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_BADGE, handleInvalidateUserBadge);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_EDUCATIONS, handleInvalidateUserEducations);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_EDUCATION, handleInvalidateUserEducation);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_EMAILS, handleInvalidateUserEmails);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_EMAIL, handleInvalidateUserEmail);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_OCCUPATIONS, handleInvalidateUserOccupations);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_OCCUPATION, handleInvalidateUserOccupation);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PHONES, handleInvalidateUserPhones);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PHONE, handleInvalidateUserPhone);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PREFERENCES, handleInvalidateUserPreferences);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PREFERENCE, handleInvalidateUserPreference);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PROFILES, handleInvalidateUserProfiles);
    directoryEventEmitter.on(directoryEvents.INVALIDATE_USER_PROFILE, handleInvalidateUserProfile);

    // Image Events
    imageEventEmitter.on(imageEvents.INVALIDATE_IMAGES, handleInvalidateImages);
    imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE, handleInvalidateImage);
    imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_TYPES, handleInvalidateImageTypes);
    imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_TYPE, handleInvalidateImageType);
    imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_VARIANTS, handleInvalidateImageVariants);
    imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_VARIANT, handleInvalidateImageVariant);
    imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_TAGS, handleInvalidateImageTags);
    imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_TAG, handleInvalidateImageTag);

    // System Events
    systemEventEmitter.on(systemEvents.INVALIDATE_SYSTEM_STATES, handleInvalidateSystemStates);
    systemEventEmitter.on(systemEvents.INVALIDATE_SYSTEM_STATE, handleInvalidateSystemState);

    // Tenants Events
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_TENANTS, handleInvalidateTenants);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_TENANT, handleInvalidateTenant);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_GROUPS, handleInvalidateGroups);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_GROUP, handleInvalidateGroup);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBERS, handleInvalidateMembers);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBER, handleInvalidateMember);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_INVITATIONS, handleInvalidateInvitations);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_INVITATION, handleInvalidateInvitation);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_TENANT_APPS, handleInvalidateTenantApps);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_TENANT_APP, handleInvalidateTenantApp);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_TENANT_SETTINGS, handleInvalidateTenantSettings);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_TENANT_SETTING, handleInvalidateTenantSetting);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_DOMAIN_VERIFICATIONS, handleInvalidateDomainVerifications);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_DOMAIN_VERIFICATION, handleInvalidateDomainVerification);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBER_ROLES, handleInvalidateMemberRoles);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBER_ROLE, handleInvalidateMemberRole);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBER_GROUPS, handleInvalidateMemberGroups);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBER_GROUP, handleInvalidateMemberGroup);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBER_APP_ROLES, handleInvalidateMemberAppRoles);
    tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBER_APP_ROLE, handleInvalidateMemberAppRole);

    // ========== 清理所有监听器 ==========
    return () => {
      // Auth Events
      authEventEmitter.off(authEvents.INVALIDATE_SESSIONS, handleInvalidateSessions);
      authEventEmitter.off(authEvents.INVALIDATE_SESSION, handleInvalidateSession);
      authEventEmitter.off(authEvents.INVALIDATE_USER_CREDENTIALS, handleInvalidateUserCredentials);
      authEventEmitter.off(authEvents.INVALIDATE_USER_CREDENTIAL, handleInvalidateUserCredential);
      authEventEmitter.off(authEvents.INVALIDATE_MFA_FACTORS, handleInvalidateMFAFactors);
      authEventEmitter.off(authEvents.INVALIDATE_MFA_FACTOR, handleInvalidateMFAFactor);
      authEventEmitter.off(authEvents.INVALIDATE_REFRESH_TOKENS, handleInvalidateRefreshTokens);
      authEventEmitter.off(authEvents.INVALIDATE_REFRESH_TOKEN, handleInvalidateRefreshToken);
      authEventEmitter.off(authEvents.INVALIDATE_PASSWORD_RESETS, handleInvalidatePasswordResets);
      authEventEmitter.off(authEvents.INVALIDATE_PASSWORD_RESET, handleInvalidatePasswordReset);
      authEventEmitter.off(authEvents.INVALIDATE_PASSWORD_HISTORIES, handleInvalidatePasswordHistories);
      authEventEmitter.off(authEvents.INVALIDATE_PASSWORD_HISTORY, handleInvalidatePasswordHistory);
      authEventEmitter.off(authEvents.INVALIDATE_LOGIN_ATTEMPTS, handleInvalidateLoginAttempts);
      authEventEmitter.off(authEvents.INVALIDATE_LOGIN_ATTEMPT, handleInvalidateLoginAttempt);
      authEventEmitter.off(authEvents.INVALIDATE_ACCOUNT_LOCKOUTS, handleInvalidateAccountLockouts);
      authEventEmitter.off(authEvents.INVALIDATE_ACCOUNT_LOCKOUT, handleInvalidateAccountLockout);
      authEventEmitter.off(authEvents.INVALIDATE_TRUSTED_DEVICES, handleInvalidateTrustedDevices);
      authEventEmitter.off(authEvents.INVALIDATE_TRUSTED_DEVICE, handleInvalidateTrustedDevice);

      // Access Events
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

      // Audit Events
      auditEventEmitter.off(auditEvents.INVALIDATE_EVENTS, handleInvalidateEvents);
      auditEventEmitter.off(auditEvents.INVALIDATE_EVENT, handleInvalidateEvent);
      auditEventEmitter.off(auditEvents.INVALIDATE_ACTOR_SNAPSHOTS, handleInvalidateActorSnapshots);
      auditEventEmitter.off(auditEvents.INVALIDATE_ACTOR_SNAPSHOT, handleInvalidateActorSnapshot);
      auditEventEmitter.off(auditEvents.INVALIDATE_EVENT_RETENTION_POLICIES, handleInvalidateEventRetentionPolicies);
      auditEventEmitter.off(auditEvents.INVALIDATE_EVENT_RETENTION_POLICY, handleInvalidateEventRetentionPolicy);
      auditEventEmitter.off(auditEvents.INVALIDATE_EVENT_SEARCH_INDICES, handleInvalidateEventSearchIndices);
      auditEventEmitter.off(auditEvents.INVALIDATE_EVENT_SEARCH_INDEX, handleInvalidateEventSearchIndex);
      auditEventEmitter.off(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINTS, handleInvalidateHashChainCheckpoints);
      auditEventEmitter.off(auditEvents.INVALIDATE_HASH_CHAIN_CHECKPOINT, handleInvalidateHashChainCheckpoint);

      // Clients Events
      clientsEventEmitter.off(clientsEvents.INVALIDATE_APPS, handleInvalidateApps);
      clientsEventEmitter.off(clientsEvents.INVALIDATE_APP, handleInvalidateApp);
      clientsEventEmitter.off(clientsEvents.INVALIDATE_API_KEYS, handleInvalidateAPIKeys);
      clientsEventEmitter.off(clientsEvents.INVALIDATE_API_KEY, handleInvalidateAPIKey);
      clientsEventEmitter.off(clientsEvents.INVALIDATE_CLIENT_CREDENTIALS, handleInvalidateClientCredentials);
      clientsEventEmitter.off(clientsEvents.INVALIDATE_CLIENT_CREDENTIAL, handleInvalidateClientCredential);
      clientsEventEmitter.off(clientsEvents.INVALIDATE_CLIENT_SCOPES, handleInvalidateClientScopes);
      clientsEventEmitter.off(clientsEvents.INVALIDATE_CLIENT_SCOPE, handleInvalidateClientScope);
      clientsEventEmitter.off(clientsEvents.INVALIDATE_IP_ALLOWLISTS, handleInvalidateIPAllowlists);
      clientsEventEmitter.off(clientsEvents.INVALIDATE_IP_ALLOWLIST, handleInvalidateIPAllowlist);
      clientsEventEmitter.off(clientsEvents.INVALIDATE_RATE_LIMITS, handleInvalidateRateLimits);
      clientsEventEmitter.off(clientsEvents.INVALIDATE_RATE_LIMIT, handleInvalidateRateLimit);

      // Directory Events
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USERS, handleInvalidateUsers);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER, handleInvalidateUser);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_BADGES, handleInvalidateBadges);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_BADGE, handleInvalidateBadge);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_BADGES, handleInvalidateUserBadges);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_BADGE, handleInvalidateUserBadge);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_EDUCATIONS, handleInvalidateUserEducations);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_EDUCATION, handleInvalidateUserEducation);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_EMAILS, handleInvalidateUserEmails);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_EMAIL, handleInvalidateUserEmail);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_OCCUPATIONS, handleInvalidateUserOccupations);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_OCCUPATION, handleInvalidateUserOccupation);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PHONES, handleInvalidateUserPhones);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PHONE, handleInvalidateUserPhone);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PREFERENCES, handleInvalidateUserPreferences);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PREFERENCE, handleInvalidateUserPreference);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PROFILES, handleInvalidateUserProfiles);
      directoryEventEmitter.off(directoryEvents.INVALIDATE_USER_PROFILE, handleInvalidateUserProfile);

      // Image Events
      imageEventEmitter.off(imageEvents.INVALIDATE_IMAGES, handleInvalidateImages);
      imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE, handleInvalidateImage);
      imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_TYPES, handleInvalidateImageTypes);
      imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_TYPE, handleInvalidateImageType);
      imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_VARIANTS, handleInvalidateImageVariants);
      imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_VARIANT, handleInvalidateImageVariant);
      imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_TAGS, handleInvalidateImageTags);
      imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_TAG, handleInvalidateImageTag);

      // System Events
      systemEventEmitter.off(systemEvents.INVALIDATE_SYSTEM_STATES, handleInvalidateSystemStates);
      systemEventEmitter.off(systemEvents.INVALIDATE_SYSTEM_STATE, handleInvalidateSystemState);

      // Tenants Events
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_TENANTS, handleInvalidateTenants);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_TENANT, handleInvalidateTenant);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_GROUPS, handleInvalidateGroups);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_GROUP, handleInvalidateGroup);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBERS, handleInvalidateMembers);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBER, handleInvalidateMember);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_INVITATIONS, handleInvalidateInvitations);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_INVITATION, handleInvalidateInvitation);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_TENANT_APPS, handleInvalidateTenantApps);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_TENANT_APP, handleInvalidateTenantApp);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_TENANT_SETTINGS, handleInvalidateTenantSettings);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_TENANT_SETTING, handleInvalidateTenantSetting);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_DOMAIN_VERIFICATIONS, handleInvalidateDomainVerifications);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_DOMAIN_VERIFICATION, handleInvalidateDomainVerification);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBER_ROLES, handleInvalidateMemberRoles);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBER_ROLE, handleInvalidateMemberRole);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBER_GROUPS, handleInvalidateMemberGroups);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBER_GROUP, handleInvalidateMemberGroup);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBER_APP_ROLES, handleInvalidateMemberAppRoles);
      tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBER_APP_ROLE, handleInvalidateMemberAppRole);
    };
  }, [queryClient]);
};
