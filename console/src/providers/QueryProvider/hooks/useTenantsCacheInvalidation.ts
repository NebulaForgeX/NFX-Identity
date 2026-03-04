import type { QueryClient } from "@tanstack/react-query";
import { tenantsEventEmitter, tenantsEvents } from "@/events/tenants";
import {
  TENANTS_DOMAIN_VERIFICATION,
  TENANTS_DOMAIN_VERIFICATION_LIST,
  TENANTS_GROUP,
  TENANTS_GROUP_LIST,
  TENANTS_INVITATION,
  TENANTS_INVITATION_LIST,
  TENANTS_MEMBER,
  TENANTS_MEMBER_APP_ROLE,
  TENANTS_MEMBER_APP_ROLE_LIST,
  TENANTS_MEMBER_GROUP,
  TENANTS_MEMBER_GROUP_LIST,
  TENANTS_MEMBER_LIST,
  TENANTS_MEMBER_ROLE,
  TENANTS_MEMBER_ROLE_LIST,
  TENANTS_TENANT,
  TENANTS_TENANT_APP,
  TENANTS_TENANT_APP_LIST,
  TENANTS_TENANT_LIST,
  TENANTS_TENANT_SETTING,
  TENANTS_TENANT_SETTING_LIST,
} from "@/constants";

type EventCb = (...args: unknown[]) => void;

/**
 * Tenants 相关的缓存失效事件处理
 */
export const useTenantsCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateTenants = () => queryClient.invalidateQueries({ queryKey: TENANTS_TENANT_LIST });
  const handleInvalidateTenant = (item: string) => queryClient.invalidateQueries({ queryKey: TENANTS_TENANT(item) });
  const handleInvalidateGroups = () => queryClient.invalidateQueries({ queryKey: TENANTS_GROUP_LIST });
  const handleInvalidateGroup = (item: string) => queryClient.invalidateQueries({ queryKey: TENANTS_GROUP(item) });
  const handleInvalidateMembers = () => queryClient.invalidateQueries({ queryKey: TENANTS_MEMBER_LIST });
  const handleInvalidateMember = (item: string) => queryClient.invalidateQueries({ queryKey: TENANTS_MEMBER(item) });
  const handleInvalidateInvitations = () => queryClient.invalidateQueries({ queryKey: TENANTS_INVITATION_LIST });
  const handleInvalidateInvitation = (item: string) => queryClient.invalidateQueries({ queryKey: TENANTS_INVITATION(item) });
  const handleInvalidateTenantApps = () => queryClient.invalidateQueries({ queryKey: TENANTS_TENANT_APP_LIST });
  const handleInvalidateTenantApp = (item: string) => queryClient.invalidateQueries({ queryKey: TENANTS_TENANT_APP(item) });
  const handleInvalidateTenantSettings = () => queryClient.invalidateQueries({ queryKey: TENANTS_TENANT_SETTING_LIST });
  const handleInvalidateTenantSetting = (item: string) => queryClient.invalidateQueries({ queryKey: TENANTS_TENANT_SETTING(item) });
  const handleInvalidateDomainVerifications = () => queryClient.invalidateQueries({ queryKey: TENANTS_DOMAIN_VERIFICATION_LIST });
  const handleInvalidateDomainVerification = (item: string) => queryClient.invalidateQueries({ queryKey: TENANTS_DOMAIN_VERIFICATION(item) });
  const handleInvalidateMemberRoles = () => queryClient.invalidateQueries({ queryKey: TENANTS_MEMBER_ROLE_LIST });
  const handleInvalidateMemberRole = (item: string) => queryClient.invalidateQueries({ queryKey: TENANTS_MEMBER_ROLE(item) });
  const handleInvalidateMemberGroups = () => queryClient.invalidateQueries({ queryKey: TENANTS_MEMBER_GROUP_LIST });
  const handleInvalidateMemberGroup = (item: string) => queryClient.invalidateQueries({ queryKey: TENANTS_MEMBER_GROUP(item) });
  const handleInvalidateMemberAppRoles = () => queryClient.invalidateQueries({ queryKey: TENANTS_MEMBER_APP_ROLE_LIST });
  const handleInvalidateMemberAppRole = (item: string) => queryClient.invalidateQueries({ queryKey: TENANTS_MEMBER_APP_ROLE(item) });

  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_TENANTS, handleInvalidateTenants as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_TENANT, handleInvalidateTenant as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_GROUPS, handleInvalidateGroups as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_GROUP, handleInvalidateGroup as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBERS, handleInvalidateMembers as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBER, handleInvalidateMember as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_INVITATIONS, handleInvalidateInvitations as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_INVITATION, handleInvalidateInvitation as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_TENANT_APPS, handleInvalidateTenantApps as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_TENANT_APP, handleInvalidateTenantApp as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_TENANT_SETTINGS, handleInvalidateTenantSettings as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_TENANT_SETTING, handleInvalidateTenantSetting as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_DOMAIN_VERIFICATIONS, handleInvalidateDomainVerifications as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_DOMAIN_VERIFICATION, handleInvalidateDomainVerification as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBER_ROLES, handleInvalidateMemberRoles as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBER_ROLE, handleInvalidateMemberRole as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBER_GROUPS, handleInvalidateMemberGroups as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBER_GROUP, handleInvalidateMemberGroup as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBER_APP_ROLES, handleInvalidateMemberAppRoles as EventCb);
  tenantsEventEmitter.on(tenantsEvents.INVALIDATE_MEMBER_APP_ROLE, handleInvalidateMemberAppRole as EventCb);

  return () => {
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_TENANTS, handleInvalidateTenants as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_TENANT, handleInvalidateTenant as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_GROUPS, handleInvalidateGroups as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_GROUP, handleInvalidateGroup as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBERS, handleInvalidateMembers as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBER, handleInvalidateMember as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_INVITATIONS, handleInvalidateInvitations as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_INVITATION, handleInvalidateInvitation as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_TENANT_APPS, handleInvalidateTenantApps as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_TENANT_APP, handleInvalidateTenantApp as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_TENANT_SETTINGS, handleInvalidateTenantSettings as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_TENANT_SETTING, handleInvalidateTenantSetting as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_DOMAIN_VERIFICATIONS, handleInvalidateDomainVerifications as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_DOMAIN_VERIFICATION, handleInvalidateDomainVerification as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBER_ROLES, handleInvalidateMemberRoles as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBER_ROLE, handleInvalidateMemberRole as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBER_GROUPS, handleInvalidateMemberGroups as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBER_GROUP, handleInvalidateMemberGroup as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBER_APP_ROLES, handleInvalidateMemberAppRoles as EventCb);
    tenantsEventEmitter.off(tenantsEvents.INVALIDATE_MEMBER_APP_ROLE, handleInvalidateMemberAppRole as EventCb);
  };
};
