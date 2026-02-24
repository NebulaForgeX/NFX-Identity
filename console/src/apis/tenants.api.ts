// Tenants API - 基于 NFX-ID Backend

import type {
  AcceptInvitationRequest,
  BaseResponse,
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
  DataResponse,
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

import { protectedClient } from "./clients";
import { URL_PATHS } from "./ip";

// ========== 租户相关 ==========

// 创建租户
export const CreateTenant = async (params: CreateTenantRequest): Promise<Tenant> => {
  const { data } = await protectedClient.post<DataResponse<Tenant>>(
    URL_PATHS.TENANTS.tenants,
    params,
  );
  return data.data;
};

// 根据 ID 获取租户
export const GetTenant = async (id: string): Promise<Tenant> => {
  const { data } = await protectedClient.get<DataResponse<Tenant>>(
    URL_PATHS.TENANTS.tenants.byId(id),
  );
  return data.data;
};

// 根据 Tenant ID 获取租户
export const GetTenantByTenantID = async (tenantId: string): Promise<Tenant> => {
  const { data } = await protectedClient.get<DataResponse<Tenant>>(
    URL_PATHS.TENANTS.tenants.byTenantId(tenantId),
  );
  return data.data;
};

// 更新租户
export const UpdateTenant = async (id: string, params: UpdateTenantRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(
    URL_PATHS.TENANTS.tenants.byId(id),
    params,
  );
  return data;
};

// 更新租户状态
export const UpdateTenantStatus = async (id: string, params: UpdateTenantStatusRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(
    URL_PATHS.TENANTS.tenants.status(id),
    params,
  );
  return data;
};

// 删除租户
export const DeleteTenant = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.TENANTS.tenants.byId(id));
  return data;
};

// ========== 组相关 ==========

// 创建组
export const CreateGroup = async (params: CreateGroupRequest): Promise<Group> => {
  const { data } = await protectedClient.post<DataResponse<Group>>(
    URL_PATHS.TENANTS.groups,
    params,
  );
  return data.data;
};

// 根据 ID 获取组
export const GetGroup = async (id: string): Promise<Group> => {
  const { data } = await protectedClient.get<DataResponse<Group>>(
    URL_PATHS.TENANTS.groups.byId(id),
  );
  return data.data;
};

// 更新组
export const UpdateGroup = async (id: string, params: UpdateGroupRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(
    URL_PATHS.TENANTS.groups.byId(id),
    params,
  );
  return data;
};

// 删除组
export const DeleteGroup = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.TENANTS.groups.byId(id));
  return data;
};

// ========== 成员相关 ==========

// 创建成员
export const CreateMember = async (params: CreateMemberRequest): Promise<Member> => {
  const { data } = await protectedClient.post<DataResponse<Member>>(
    URL_PATHS.TENANTS.members,
    params,
  );
  return data.data;
};

// 根据 ID 获取成员
export const GetMember = async (id: string): Promise<Member> => {
  const { data } = await protectedClient.get<DataResponse<Member>>(
    URL_PATHS.TENANTS.members.byId(id),
  );
  return data.data;
};

// 更新成员
export const UpdateMember = async (id: string, params: UpdateMemberStatusRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(
    URL_PATHS.TENANTS.members.byId(id),
    params,
  );
  return data;
};

// 删除成员
export const DeleteMember = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.TENANTS.members.byId(id));
  return data;
};

// ========== 邀请相关 ==========

// 创建邀请
export const CreateInvitation = async (params: CreateInvitationRequest): Promise<Invitation> => {
  const { data } = await protectedClient.post<DataResponse<Invitation>>(
    URL_PATHS.TENANTS.invitations,
    params,
  );
  return data.data;
};

// 根据 ID 获取邀请
export const GetInvitation = async (id: string): Promise<Invitation> => {
  const { data } = await protectedClient.get<DataResponse<Invitation>>(
    URL_PATHS.TENANTS.invitations.byId(id),
  );
  return data.data;
};

// 根据 Invite ID 获取邀请
export const GetInvitationByInviteID = async (inviteId: string): Promise<Invitation> => {
  const { data } = await protectedClient.get<DataResponse<Invitation>>(
    URL_PATHS.TENANTS.invitations.byInviteId(inviteId),
  );
  return data.data;
};

// 接受邀请
export const AcceptInvitation = async (inviteId: string, params: AcceptInvitationRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(
    URL_PATHS.TENANTS.invitations.byInviteId(inviteId).accept,
    params,
  );
  return data;
};

// 撤销邀请
export const RevokeInvitation = async (inviteId: string, params: RevokeInvitationRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(
    URL_PATHS.TENANTS.invitations.byInviteId(inviteId).revoke,
    params,
  );
  return data;
};

// 删除邀请
export const DeleteInvitation = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.TENANTS.invitations.byId(id),
  );
  return data;
};

// ========== 租户应用相关 ==========

// 创建租户应用
export const CreateTenantApp = async (params: CreateTenantAppRequest): Promise<TenantApp> => {
  const { data } = await protectedClient.post<DataResponse<TenantApp>>(
    URL_PATHS.TENANTS.tenantApps,
    params,
  );
  return data.data;
};

// 根据 ID 获取租户应用
export const GetTenantApp = async (id: string): Promise<TenantApp> => {
  const { data } = await protectedClient.get<DataResponse<TenantApp>>(
    URL_PATHS.TENANTS.tenantApps.byId(id),
  );
  return data.data;
};

// 更新租户应用
export const UpdateTenantApp = async (id: string, params: UpdateTenantAppRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(
    URL_PATHS.TENANTS.tenantApps.byId(id),
    params,
  );
  return data;
};

// 删除租户应用
export const DeleteTenantApp = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.TENANTS.tenantApps.byId(id),
  );
  return data;
};

// ========== 租户设置相关 ==========

// 创建租户设置
export const CreateTenantSetting = async (params: CreateTenantSettingRequest): Promise<TenantSetting> => {
  const { data } = await protectedClient.post<DataResponse<TenantSetting>>(
    URL_PATHS.TENANTS.tenantSettings,
    params,
  );
  return data.data;
};

// 根据 ID 获取租户设置
export const GetTenantSetting = async (id: string): Promise<TenantSetting> => {
  const { data } = await protectedClient.get<DataResponse<TenantSetting>>(
    URL_PATHS.TENANTS.tenantSettings.byId(id),
  );
  return data.data;
};

// 更新租户设置
export const UpdateTenantSetting = async (id: string, params: UpdateTenantSettingRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(
    URL_PATHS.TENANTS.tenantSettings.byId(id),
    params,
  );
  return data;
};

// 删除租户设置
export const DeleteTenantSetting = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.TENANTS.tenantSettings.byId(id),
  );
  return data;
};

// ========== 域名验证相关 ==========

// 创建域名验证
export const CreateDomainVerification = async (
  params: CreateDomainVerificationRequest,
): Promise<DomainVerification> => {
  const { data } = await protectedClient.post<DataResponse<DomainVerification>>(
    URL_PATHS.TENANTS.domainVerifications,
    params,
  );
  return data.data;
};

// 根据 ID 获取域名验证
export const GetDomainVerification = async (id: string): Promise<DomainVerification> => {
  const { data } = await protectedClient.get<DataResponse<DomainVerification>>(
    URL_PATHS.TENANTS.domainVerifications.byId(id),
  );
  return data.data;
};

// 更新域名验证（验证或失败）
export const UpdateDomainVerification = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(
    URL_PATHS.TENANTS.domainVerifications.byId(id),
  );
  return data;
};

// 删除域名验证
export const DeleteDomainVerification = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.TENANTS.domainVerifications.byId(id),
  );
  return data;
};

// ========== 成员角色相关 ==========

// 创建成员角色
export const CreateMemberRole = async (params: CreateMemberRoleRequest): Promise<MemberRole> => {
  const { data } = await protectedClient.post<DataResponse<MemberRole>>(
    URL_PATHS.TENANTS.memberRoles,
    params,
  );
  return data.data;
};

// 根据 ID 获取成员角色
export const GetMemberRole = async (id: string): Promise<MemberRole> => {
  const { data } = await protectedClient.get<DataResponse<MemberRole>>(
    URL_PATHS.TENANTS.memberRoles.byId(id),
  );
  return data.data;
};

// 撤销成员角色
export const RevokeMemberRole = async (id: string, params: RevokeMemberRoleRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(
    URL_PATHS.TENANTS.memberRoles.revoke(id),
    params,
  );
  return data;
};

// 删除成员角色
export const DeleteMemberRole = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.TENANTS.memberRoles.byId(id),
  );
  return data;
};

// ========== 成员组相关 ==========

// 创建成员组
export const CreateMemberGroup = async (params: CreateMemberGroupRequest): Promise<MemberGroup> => {
  const { data } = await protectedClient.post<DataResponse<MemberGroup>>(
    URL_PATHS.TENANTS.memberGroups,
    params,
  );
  return data.data;
};

// 根据 ID 获取成员组
export const GetMemberGroup = async (id: string): Promise<MemberGroup> => {
  const { data } = await protectedClient.get<DataResponse<MemberGroup>>(
    URL_PATHS.TENANTS.memberGroups.byId(id),
  );
  return data.data;
};

// 撤销成员组
export const RevokeMemberGroup = async (id: string, params: RevokeMemberGroupRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(
    URL_PATHS.TENANTS.memberGroups.revoke(id),
    params,
  );
  return data;
};

// 删除成员组
export const DeleteMemberGroup = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.TENANTS.memberGroups.byId(id),
  );
  return data;
};

// ========== 成员应用角色相关 ==========

// 创建成员应用角色
export const CreateMemberAppRole = async (params: CreateMemberAppRoleRequest): Promise<MemberAppRole> => {
  const { data } = await protectedClient.post<DataResponse<MemberAppRole>>(
    URL_PATHS.TENANTS.memberAppRoles,
    params,
  );
  return data.data;
};

// 根据 ID 获取成员应用角色
export const GetMemberAppRole = async (id: string): Promise<MemberAppRole> => {
  const { data } = await protectedClient.get<DataResponse<MemberAppRole>>(
    URL_PATHS.TENANTS.memberAppRoles.byId(id),
  );
  return data.data;
};

// 撤销成员应用角色
export const RevokeMemberAppRole = async (id: string, params: RevokeMemberAppRoleRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.patch<BaseResponse>(
    URL_PATHS.TENANTS.memberAppRoles.revoke(id),
    params,
  );
  return data;
};

// 删除成员应用角色
export const DeleteMemberAppRole = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.TENANTS.memberAppRoles.byId(id),
  );
  return data;
};
