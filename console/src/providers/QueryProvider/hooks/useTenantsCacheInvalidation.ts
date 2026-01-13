import type { QueryClient } from "@tanstack/react-query";
import { tenantsEventEmitter, tenantsEvents } from "@/events/tenants";
import { TENANTS_QUERY_KEY_PREFIXES } from "@/constants";

/**
 * Tenants 相关的缓存失效事件处理
 */
export const useTenantsCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateTenants = () => queryClient.invalidateQueries({ queryKey: TENANTS_QUERY_KEY_PREFIXES.TENANTS });
  const handleInvalidateTenant = (item: string) => queryClient.invalidateQueries({ queryKey: [...TENANTS_QUERY_KEY_PREFIXES.TENANT, item] });
  const handleInvalidateGroups = () => queryClient.invalidateQueries({ queryKey: TENANTS_QUERY_KEY_PREFIXES.GROUPS });
  const handleInvalidateGroup = (item: string) => queryClient.invalidateQueries({ queryKey: [...TENANTS_QUERY_KEY_PREFIXES.GROUP, item] });
  const handleInvalidateMembers = () => queryClient.invalidateQueries({ queryKey: TENANTS_QUERY_KEY_PREFIXES.MEMBERS });
  const handleInvalidateMember = (item: string) => queryClient.invalidateQueries({ queryKey: [...TENANTS_QUERY_KEY_PREFIXES.MEMBER, item] });
  const handleInvalidateInvitations = () => queryClient.invalidateQueries({ queryKey: TENANTS_QUERY_KEY_PREFIXES.INVITATIONS });
  const handleInvalidateInvitation = (item: string) => queryClient.invalidateQueries({ queryKey: [...TENANTS_QUERY_KEY_PREFIXES.INVITATION, item] });
  const handleInvalidateTenantApps = () => queryClient.invalidateQueries({ queryKey: TENANTS_QUERY_KEY_PREFIXES.TENANT_APPS });
  const handleInvalidateTenantApp = (item: string) => queryClient.invalidateQueries({ queryKey: [...TENANTS_QUERY_KEY_PREFIXES.TENANT_APP, item] });
  const handleInvalidateTenantSettings = () => queryClient.invalidateQueries({ queryKey: TENANTS_QUERY_KEY_PREFIXES.TENANT_SETTINGS });
  const handleInvalidateTenantSetting = (item: string) => queryClient.invalidateQueries({ queryKey: [...TENANTS_QUERY_KEY_PREFIXES.TENANT_SETTING, item] });
  const handleInvalidateDomainVerifications = () => queryClient.invalidateQueries({ queryKey: TENANTS_QUERY_KEY_PREFIXES.DOMAIN_VERIFICATIONS });
  const handleInvalidateDomainVerification = (item: string) => queryClient.invalidateQueries({ queryKey: [...TENANTS_QUERY_KEY_PREFIXES.DOMAIN_VERIFICATION, item] });
  const handleInvalidateMemberRoles = () => queryClient.invalidateQueries({ queryKey: TENANTS_QUERY_KEY_PREFIXES.MEMBER_ROLES });
  const handleInvalidateMemberRole = (item: string) => queryClient.invalidateQueries({ queryKey: [...TENANTS_QUERY_KEY_PREFIXES.MEMBER_ROLE, item] });
  const handleInvalidateMemberGroups = () => queryClient.invalidateQueries({ queryKey: TENANTS_QUERY_KEY_PREFIXES.MEMBER_GROUPS });
  const handleInvalidateMemberGroup = (item: string) => queryClient.invalidateQueries({ queryKey: [...TENANTS_QUERY_KEY_PREFIXES.MEMBER_GROUP, item] });
  const handleInvalidateMemberAppRoles = () => queryClient.invalidateQueries({ queryKey: TENANTS_QUERY_KEY_PREFIXES.MEMBER_APP_ROLES });
  const handleInvalidateMemberAppRole = (item: string) => queryClient.invalidateQueries({ queryKey: [...TENANTS_QUERY_KEY_PREFIXES.MEMBER_APP_ROLE, item] });

  // 注册监听器
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

  // 清理监听器
  return () => {
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
};
