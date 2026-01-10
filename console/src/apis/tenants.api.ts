// Tenants API - 基于 NFX-ID Backend

import type { BaseResponse, DataResponse } from "@/types/api";

import { protectedClient } from "@/apis/clients";
import { URL_PATHS } from "@/apis/ip";

// ========== 租户相关 ==========

// 创建租户
export const CreateTenant = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.TENANTS.CREATE_TENANT, params);
  return data.data;
};

// 根据 ID 获取租户
export const GetTenant = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.TENANTS.GET_TENANT.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 根据 Tenant ID 获取租户
export const GetTenantByTenantID = async (tenantId: string): Promise<unknown> => {
  const url = URL_PATHS.TENANTS.GET_TENANT_BY_TENANT_ID.replace(":tenant_id", tenantId);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新租户
export const UpdateTenant = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.UPDATE_TENANT.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 更新租户状态
export const UpdateTenantStatus = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.UPDATE_TENANT_STATUS.replace(":id", id);
  const { data } = await protectedClient.patch<BaseResponse>(url, params);
  return data;
};

// 删除租户
export const DeleteTenant = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.DELETE_TENANT.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 组相关 ==========

// 创建组
export const CreateGroup = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.TENANTS.CREATE_GROUP, params);
  return data.data;
};

// 根据 ID 获取组
export const GetGroup = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.TENANTS.GET_GROUP.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新组
export const UpdateGroup = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.UPDATE_GROUP.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除组
export const DeleteGroup = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.DELETE_GROUP.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 成员相关 ==========

// 创建成员
export const CreateMember = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.TENANTS.CREATE_MEMBER, params);
  return data.data;
};

// 根据 ID 获取成员
export const GetMember = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.TENANTS.GET_MEMBER.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新成员
export const UpdateMember = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.UPDATE_MEMBER.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除成员
export const DeleteMember = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.DELETE_MEMBER.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 邀请相关 ==========

// 创建邀请
export const CreateInvitation = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.TENANTS.CREATE_INVITATION, params);
  return data.data;
};

// 根据 ID 获取邀请
export const GetInvitation = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.TENANTS.GET_INVITATION.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 根据 Invite ID 获取邀请
export const GetInvitationByInviteID = async (inviteId: string): Promise<unknown> => {
  const url = URL_PATHS.TENANTS.GET_INVITATION_BY_INVITE_ID.replace(":invite_id", inviteId);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 接受邀请
export const AcceptInvitation = async (inviteId: string, params?: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.ACCEPT_INVITATION.replace(":invite_id", inviteId);
  const { data } = await protectedClient.patch<BaseResponse>(url, params);
  return data;
};

// 撤销邀请
export const RevokeInvitation = async (inviteId: string, params?: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.REVOKE_INVITATION.replace(":invite_id", inviteId);
  const { data } = await protectedClient.patch<BaseResponse>(url, params);
  return data;
};

// 删除邀请
export const DeleteInvitation = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.DELETE_INVITATION.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 租户应用相关 ==========

// 创建租户应用
export const CreateTenantApp = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.TENANTS.CREATE_TENANT_APP, params);
  return data.data;
};

// 根据 ID 获取租户应用
export const GetTenantApp = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.TENANTS.GET_TENANT_APP.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新租户应用
export const UpdateTenantApp = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.UPDATE_TENANT_APP.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除租户应用
export const DeleteTenantApp = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.DELETE_TENANT_APP.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 租户设置相关 ==========

// 创建租户设置
export const CreateTenantSetting = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.TENANTS.CREATE_TENANT_SETTING, params);
  return data.data;
};

// 根据 ID 获取租户设置
export const GetTenantSetting = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.TENANTS.GET_TENANT_SETTING.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新租户设置
export const UpdateTenantSetting = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.UPDATE_TENANT_SETTING.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除租户设置
export const DeleteTenantSetting = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.DELETE_TENANT_SETTING.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 域名验证相关 ==========

// 创建域名验证
export const CreateDomainVerification = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.TENANTS.CREATE_DOMAIN_VERIFICATION, params);
  return data.data;
};

// 根据 ID 获取域名验证
export const GetDomainVerification = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.TENANTS.GET_DOMAIN_VERIFICATION.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新域名验证
export const UpdateDomainVerification = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.UPDATE_DOMAIN_VERIFICATION.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除域名验证
export const DeleteDomainVerification = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.DELETE_DOMAIN_VERIFICATION.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 成员角色相关 ==========

// 创建成员角色
export const CreateMemberRole = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.TENANTS.CREATE_MEMBER_ROLE, params);
  return data.data;
};

// 根据 ID 获取成员角色
export const GetMemberRole = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.TENANTS.GET_MEMBER_ROLE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 撤销成员角色
export const RevokeMemberRole = async (id: string, params?: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.REVOKE_MEMBER_ROLE.replace(":id", id);
  const { data } = await protectedClient.patch<BaseResponse>(url, params);
  return data;
};

// 删除成员角色
export const DeleteMemberRole = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.DELETE_MEMBER_ROLE.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 成员组相关 ==========

// 创建成员组
export const CreateMemberGroup = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.TENANTS.CREATE_MEMBER_GROUP, params);
  return data.data;
};

// 根据 ID 获取成员组
export const GetMemberGroup = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.TENANTS.GET_MEMBER_GROUP.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 撤销成员组
export const RevokeMemberGroup = async (id: string, params?: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.REVOKE_MEMBER_GROUP.replace(":id", id);
  const { data } = await protectedClient.patch<BaseResponse>(url, params);
  return data;
};

// 删除成员组
export const DeleteMemberGroup = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.DELETE_MEMBER_GROUP.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 成员应用角色相关 ==========

// 创建成员应用角色
export const CreateMemberAppRole = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.TENANTS.CREATE_MEMBER_APP_ROLE, params);
  return data.data;
};

// 根据 ID 获取成员应用角色
export const GetMemberAppRole = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.TENANTS.GET_MEMBER_APP_ROLE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 撤销成员应用角色
export const RevokeMemberAppRole = async (id: string, params?: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.REVOKE_MEMBER_APP_ROLE.replace(":id", id);
  const { data } = await protectedClient.patch<BaseResponse>(url, params);
  return data;
};

// 删除成员应用角色
export const DeleteMemberAppRole = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.TENANTS.DELETE_MEMBER_APP_ROLE.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};
