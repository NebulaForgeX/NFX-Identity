import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";

import {
  AcceptInvitation,
  CreateDomainVerification,
  CreateGroup,
  CreateInvitation,
  CreateMember,
  CreateMemberAppRole,
  CreateMemberGroup,
  CreateMemberRole,
  CreateTenant,
  CreateTenantApp,
  CreateTenantSetting,
  DeleteDomainVerification,
  DeleteGroup,
  DeleteInvitation,
  DeleteMember,
  DeleteMemberAppRole,
  DeleteMemberGroup,
  DeleteMemberRole,
  DeleteTenant,
  DeleteTenantApp,
  DeleteTenantSetting,
  GetDomainVerification,
  GetGroup,
  GetInvitation,
  GetInvitationByInviteID,
  GetMember,
  GetMemberAppRole,
  GetMemberGroup,
  GetMemberRole,
  GetTenant,
  GetTenantApp,
  GetTenantByTenantID,
  GetTenantSetting,
  RevokeInvitation,
  RevokeMemberAppRole,
  RevokeMemberGroup,
  RevokeMemberRole,
  UpdateDomainVerification,
  UpdateGroup,
  UpdateMember,
  UpdateTenant,
  UpdateTenantApp,
  UpdateTenantSetting,
  UpdateTenantStatus,
} from "@/apis/tenants.api";
import type {
  AcceptInvitationRequest,
  CreateDomainVerificationRequest,
  CreateGroupRequest,
  CreateInvitationRequest,
  CreateMemberAppRoleRequest,
  CreateMemberGroupRequest,
  CreateMemberRequest,
  CreateMemberRoleRequest,
  CreateTenantAppRequest,
  CreateTenantRequest,
  CreateTenantSettingRequest,
  DomainVerification,
  Group,
  Invitation,
  Member,
  MemberAppRole,
  MemberGroup,
  MemberRole,
  RevokeInvitationRequest,
  RevokeMemberAppRoleRequest,
  RevokeMemberGroupRequest,
  RevokeMemberRoleRequest,
  Tenant,
  TenantApp,
  TenantSetting,
  UpdateGroupRequest,
  UpdateMemberStatusRequest,
  UpdateTenantAppRequest,
  UpdateTenantRequest,
  UpdateTenantSettingRequest,
  UpdateTenantStatusRequest,
} from "@/types";
import { makeUnifiedQuery } from "@/hooks/core/makeUnifiedQuery";
import { tenantsEventEmitter, tenantsEvents } from "@/events/tenants";
import { showError, showSuccess } from "@/stores/modalStore";
import {
  TENANTS_TENANT,
  TENANTS_GROUP,
  TENANTS_MEMBER,
  TENANTS_INVITATION,
  TENANTS_TENANT_APP,
  TENANTS_TENANT_SETTING,
  TENANTS_DOMAIN_VERIFICATION,
  TENANTS_MEMBER_ROLE,
  TENANTS_MEMBER_GROUP,
  TENANTS_MEMBER_APP_ROLE,
} from "@/constants";
import type { UnifiedQueryParams } from "./core/type";

// ========== Tenant 相关 ==========

// 根据 ID 获取租户
export const useTenant = (params: UnifiedQueryParams<Tenant> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetTenant(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(TENANTS_TENANT(id), { id }, options);
};

// 根据 Tenant ID 获取租户
export const useTenantByTenantID = (params: UnifiedQueryParams<Tenant> & { tenantId: string }) => {
  const { tenantId, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { tenantId: string }) => {
      return await GetTenantByTenantID(params.tenantId);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(TENANTS_TENANT(tenantId), { tenantId }, options);
};

// 创建租户
export const useCreateTenant = () => {
  return useMutation({
    mutationFn: async (params: CreateTenantRequest) => {
      return await CreateTenant(params);
    },
    onSuccess: () => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANTS);
      showSuccess("租户创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建租户失败，请稍后重试。" + error.message);
    },
  });
};

// 更新租户
export const useUpdateTenant = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateTenantRequest }) => {
      return await UpdateTenant(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANTS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT, variables.id);
      showSuccess("租户更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新租户失败，请稍后重试。" + error.message);
    },
  });
};

// 更新租户状态
export const useUpdateTenantStatus = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateTenantStatusRequest }) => {
      return await UpdateTenantStatus(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANTS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT, variables.id);
      showSuccess("租户状态更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新租户状态失败，请稍后重试。" + error.message);
    },
  });
};

// 删除租户
export const useDeleteTenant = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteTenant(id);
    },
    onSuccess: (_, id) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANTS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT, id);
      showSuccess("租户删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除租户失败，请稍后重试。" + error.message);
    },
  });
};

// ========== Group 相关 ==========

// 根据 ID 获取组
export const useGroup = (params: UnifiedQueryParams<Group> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetGroup(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(TENANTS_GROUP(id), { id }, options);
};

// 创建组
export const useCreateGroup = () => {
  return useMutation({
    mutationFn: async (params: CreateGroupRequest) => {
      return await CreateGroup(params);
    },
    onSuccess: () => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_GROUPS);
      showSuccess("组创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建组失败，请稍后重试。" + error.message);
    },
  });
};

// 更新组
export const useUpdateGroup = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateGroupRequest }) => {
      return await UpdateGroup(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_GROUPS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_GROUP, variables.id);
      showSuccess("组更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新组失败，请稍后重试。" + error.message);
    },
  });
};

// 删除组
export const useDeleteGroup = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteGroup(id);
    },
    onSuccess: (_, id) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_GROUPS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_GROUP, id);
      showSuccess("组删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除组失败，请稍后重试。" + error.message);
    },
  });
};

// ========== Member 相关 ==========

// 根据 ID 获取成员
export const useMember = (params: UnifiedQueryParams<Member> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetMember(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(TENANTS_MEMBER(id), { id }, options);
};

// 创建成员
export const useCreateMember = () => {
  return useMutation({
    mutationFn: async (params: CreateMemberRequest) => {
      return await CreateMember(params);
    },
    onSuccess: () => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBERS);
      showSuccess("成员创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建成员失败，请稍后重试。" + error.message);
    },
  });
};

// 更新成员
export const useUpdateMember = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateMemberStatusRequest }) => {
      return await UpdateMember(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBERS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER, variables.id);
      showSuccess("成员更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新成员失败，请稍后重试。" + error.message);
    },
  });
};

// 删除成员
export const useDeleteMember = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteMember(id);
    },
    onSuccess: (_, id) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBERS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER, id);
      showSuccess("成员删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除成员失败，请稍后重试。" + error.message);
    },
  });
};

// ========== Invitation 相关 ==========

// 根据 ID 获取邀请
export const useInvitation = (params: UnifiedQueryParams<Invitation> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetInvitation(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(TENANTS_INVITATION(id), { id }, options);
};

// 根据 Invite ID 获取邀请
export const useInvitationByInviteID = (params: UnifiedQueryParams<Invitation> & { inviteId: string }) => {
  const { inviteId, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { inviteId: string }) => {
      return await GetInvitationByInviteID(params.inviteId);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(TENANTS_INVITATION(inviteId), { inviteId }, options);
};

// 创建邀请
export const useCreateInvitation = () => {
  return useMutation({
    mutationFn: async (params: CreateInvitationRequest) => {
      return await CreateInvitation(params);
    },
    onSuccess: () => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_INVITATIONS);
      showSuccess("邀请创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建邀请失败，请稍后重试。" + error.message);
    },
  });
};

// 接受邀请
export const useAcceptInvitation = () => {
  return useMutation({
    mutationFn: async (params: { inviteId: string; data: AcceptInvitationRequest }) => {
      return await AcceptInvitation(params.inviteId, params.data);
    },
    onSuccess: (_, variables) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_INVITATIONS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_INVITATION, variables.inviteId);
      showSuccess("邀请接受成功！");
    },
    onError: (error: AxiosError) => {
      showError("接受邀请失败，请稍后重试。" + error.message);
    },
  });
};

// 撤销邀请
export const useRevokeInvitation = () => {
  return useMutation({
    mutationFn: async (params: { inviteId: string; data: RevokeInvitationRequest }) => {
      return await RevokeInvitation(params.inviteId, params.data);
    },
    onSuccess: (_, variables) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_INVITATIONS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_INVITATION, variables.inviteId);
      showSuccess("邀请撤销成功！");
    },
    onError: (error: AxiosError) => {
      showError("撤销邀请失败，请稍后重试。" + error.message);
    },
  });
};

// 删除邀请
export const useDeleteInvitation = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteInvitation(id);
    },
    onSuccess: (_, id) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_INVITATIONS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_INVITATION, id);
      showSuccess("邀请删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除邀请失败，请稍后重试。" + error.message);
    },
  });
};

// ========== TenantApp 相关 ==========

// 根据 ID 获取租户应用
export const useTenantApp = (params: UnifiedQueryParams<TenantApp> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetTenantApp(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(TENANTS_TENANT_APP(id), { id }, options);
};

// 创建租户应用
export const useCreateTenantApp = () => {
  return useMutation({
    mutationFn: async (params: CreateTenantAppRequest) => {
      return await CreateTenantApp(params);
    },
    onSuccess: () => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT_APPS);
      showSuccess("租户应用创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建租户应用失败，请稍后重试。" + error.message);
    },
  });
};

// 更新租户应用
export const useUpdateTenantApp = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateTenantAppRequest }) => {
      return await UpdateTenantApp(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT_APPS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT_APP, variables.id);
      showSuccess("租户应用更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新租户应用失败，请稍后重试。" + error.message);
    },
  });
};

// 删除租户应用
export const useDeleteTenantApp = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteTenantApp(id);
    },
    onSuccess: (_, id) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT_APPS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT_APP, id);
      showSuccess("租户应用删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除租户应用失败，请稍后重试。" + error.message);
    },
  });
};

// ========== TenantSetting 相关 ==========

// 根据 ID 获取租户设置
export const useTenantSetting = (params: UnifiedQueryParams<TenantSetting> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetTenantSetting(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(TENANTS_TENANT_SETTING(id), { id }, options);
};

// 创建租户设置
export const useCreateTenantSetting = () => {
  return useMutation({
    mutationFn: async (params: CreateTenantSettingRequest) => {
      return await CreateTenantSetting(params);
    },
    onSuccess: () => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT_SETTINGS);
      showSuccess("租户设置创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建租户设置失败，请稍后重试。" + error.message);
    },
  });
};

// 更新租户设置
export const useUpdateTenantSetting = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateTenantSettingRequest }) => {
      return await UpdateTenantSetting(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT_SETTINGS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT_SETTING, variables.id);
      showSuccess("租户设置更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新租户设置失败，请稍后重试。" + error.message);
    },
  });
};

// 删除租户设置
export const useDeleteTenantSetting = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteTenantSetting(id);
    },
    onSuccess: (_, id) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT_SETTINGS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_TENANT_SETTING, id);
      showSuccess("租户设置删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除租户设置失败，请稍后重试。" + error.message);
    },
  });
};

// ========== DomainVerification 相关 ==========

// 根据 ID 获取域名验证
export const useDomainVerification = (params: UnifiedQueryParams<DomainVerification> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetDomainVerification(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(TENANTS_DOMAIN_VERIFICATION(id), { id }, options);
};

// 创建域名验证
export const useCreateDomainVerification = () => {
  return useMutation({
    mutationFn: async (params: CreateDomainVerificationRequest) => {
      return await CreateDomainVerification(params);
    },
    onSuccess: () => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_DOMAIN_VERIFICATIONS);
      showSuccess("域名验证创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建域名验证失败，请稍后重试。" + error.message);
    },
  });
};

// 更新域名验证
export const useUpdateDomainVerification = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await UpdateDomainVerification(id);
    },
    onSuccess: (_, id) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_DOMAIN_VERIFICATIONS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_DOMAIN_VERIFICATION, id);
      showSuccess("域名验证更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新域名验证失败，请稍后重试。" + error.message);
    },
  });
};

// 删除域名验证
export const useDeleteDomainVerification = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteDomainVerification(id);
    },
    onSuccess: (_, id) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_DOMAIN_VERIFICATIONS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_DOMAIN_VERIFICATION, id);
      showSuccess("域名验证删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除域名验证失败，请稍后重试。" + error.message);
    },
  });
};

// ========== MemberRole 相关 ==========

// 根据 ID 获取成员角色
export const useMemberRole = (params: UnifiedQueryParams<MemberRole> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetMemberRole(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(TENANTS_MEMBER_ROLE(id), { id }, options);
};

// 创建成员角色
export const useCreateMemberRole = () => {
  return useMutation({
    mutationFn: async (params: CreateMemberRoleRequest) => {
      return await CreateMemberRole(params);
    },
    onSuccess: () => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_ROLES);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBERS);
      showSuccess("成员角色创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建成员角色失败，请稍后重试。" + error.message);
    },
  });
};

// 撤销成员角色
export const useRevokeMemberRole = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: RevokeMemberRoleRequest }) => {
      return await RevokeMemberRole(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_ROLES);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_ROLE, variables.id);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBERS);
      showSuccess("成员角色撤销成功！");
    },
    onError: (error: AxiosError) => {
      showError("撤销成员角色失败，请稍后重试。" + error.message);
    },
  });
};

// 删除成员角色
export const useDeleteMemberRole = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteMemberRole(id);
    },
    onSuccess: (_, id) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_ROLES);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_ROLE, id);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBERS);
      showSuccess("成员角色删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除成员角色失败，请稍后重试。" + error.message);
    },
  });
};

// ========== MemberGroup 相关 ==========

// 根据 ID 获取成员组
export const useMemberGroup = (params: UnifiedQueryParams<MemberGroup> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetMemberGroup(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(TENANTS_MEMBER_GROUP(id), { id }, options);
};

// 创建成员组
export const useCreateMemberGroup = () => {
  return useMutation({
    mutationFn: async (params: CreateMemberGroupRequest) => {
      return await CreateMemberGroup(params);
    },
    onSuccess: () => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_GROUPS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBERS);
      showSuccess("成员组创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建成员组失败，请稍后重试。" + error.message);
    },
  });
};

// 撤销成员组
export const useRevokeMemberGroup = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: RevokeMemberGroupRequest }) => {
      return await RevokeMemberGroup(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_GROUPS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_GROUP, variables.id);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBERS);
      showSuccess("成员组撤销成功！");
    },
    onError: (error: AxiosError) => {
      showError("撤销成员组失败，请稍后重试。" + error.message);
    },
  });
};

// 删除成员组
export const useDeleteMemberGroup = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteMemberGroup(id);
    },
    onSuccess: (_, id) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_GROUPS);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_GROUP, id);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBERS);
      showSuccess("成员组删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除成员组失败，请稍后重试。" + error.message);
    },
  });
};

// ========== MemberAppRole 相关 ==========

// 根据 ID 获取成员应用角色
export const useMemberAppRole = (params: UnifiedQueryParams<MemberAppRole> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetMemberAppRole(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(TENANTS_MEMBER_APP_ROLE(id), { id }, options);
};

// 创建成员应用角色
export const useCreateMemberAppRole = () => {
  return useMutation({
    mutationFn: async (params: CreateMemberAppRoleRequest) => {
      return await CreateMemberAppRole(params);
    },
    onSuccess: () => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_APP_ROLES);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBERS);
      showSuccess("成员应用角色创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建成员应用角色失败，请稍后重试。" + error.message);
    },
  });
};

// 撤销成员应用角色
export const useRevokeMemberAppRole = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: RevokeMemberAppRoleRequest }) => {
      return await RevokeMemberAppRole(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_APP_ROLES);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_APP_ROLE, variables.id);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBERS);
      showSuccess("成员应用角色撤销成功！");
    },
    onError: (error: AxiosError) => {
      showError("撤销成员应用角色失败，请稍后重试。" + error.message);
    },
  });
};

// 删除成员应用角色
export const useDeleteMemberAppRole = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteMemberAppRole(id);
    },
    onSuccess: (_, id) => {
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_APP_ROLES);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBER_APP_ROLE, id);
      tenantsEventEmitter.emit(tenantsEvents.INVALIDATE_MEMBERS);
      showSuccess("成员应用角色删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除成员应用角色失败，请稍后重试。" + error.message);
    },
  });
};
