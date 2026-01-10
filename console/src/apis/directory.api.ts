// Directory API - 基于 NFX-ID Backend

import type { BaseResponse, DataResponse } from "@/types/api";

import { protectedClient } from "@/apis/clients";
import { URL_PATHS } from "@/apis/ip";

// ========== 用户相关 ==========

// 创建用户
export const CreateUser = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.DIRECTORY.CREATE_USER, params);
  return data.data;
};

// 根据 ID 获取用户
export const GetUser = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.DIRECTORY.GET_USER.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 根据用户名获取用户
export const GetUserByUsername = async (username: string): Promise<unknown> => {
  const url = URL_PATHS.DIRECTORY.GET_USER_BY_USERNAME.replace(":username", username);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新用户状态
export const UpdateUserStatus = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.UPDATE_USER_STATUS.replace(":id", id);
  const { data } = await protectedClient.patch<BaseResponse>(url, params);
  return data;
};

// 更新用户名
export const UpdateUserUsername = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.UPDATE_USER_USERNAME.replace(":id", id);
  const { data } = await protectedClient.patch<BaseResponse>(url, params);
  return data;
};

// 验证用户
export const VerifyUser = async (id: string, params?: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.VERIFY_USER.replace(":id", id);
  const { data } = await protectedClient.patch<BaseResponse>(url, params);
  return data;
};

// 删除用户
export const DeleteUser = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.DELETE_USER.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 徽章相关 ==========

// 创建徽章
export const CreateBadge = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.DIRECTORY.CREATE_BADGE, params);
  return data.data;
};

// 根据 ID 获取徽章
export const GetBadge = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.DIRECTORY.GET_BADGE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 根据名称获取徽章
export const GetBadgeByName = async (name: string): Promise<unknown> => {
  const url = URL_PATHS.DIRECTORY.GET_BADGE_BY_NAME.replace(":name", name);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新徽章
export const UpdateBadge = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.UPDATE_BADGE.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除徽章
export const DeleteBadge = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.DELETE_BADGE.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 用户徽章相关 ==========

// 创建用户徽章
export const CreateUserBadge = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.DIRECTORY.CREATE_USER_BADGE, params);
  return data.data;
};

// 根据 ID 获取用户徽章
export const GetUserBadge = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.DIRECTORY.GET_USER_BADGE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 删除用户徽章
export const DeleteUserBadge = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.DELETE_USER_BADGE.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 用户教育相关 ==========

// 创建用户教育
export const CreateUserEducation = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.DIRECTORY.CREATE_USER_EDUCATION, params);
  return data.data;
};

// 根据 ID 获取用户教育
export const GetUserEducation = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.DIRECTORY.GET_USER_EDUCATION.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新用户教育
export const UpdateUserEducation = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.UPDATE_USER_EDUCATION.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除用户教育
export const DeleteUserEducation = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.DELETE_USER_EDUCATION.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 用户邮箱相关 ==========

// 创建用户邮箱
export const CreateUserEmail = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.DIRECTORY.CREATE_USER_EMAIL, params);
  return data.data;
};

// 根据 ID 获取用户邮箱
export const GetUserEmail = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.DIRECTORY.GET_USER_EMAIL.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新用户邮箱
export const UpdateUserEmail = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.UPDATE_USER_EMAIL.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 设置主邮箱
export const SetPrimaryUserEmail = async (id: string, params?: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.SET_PRIMARY_USER_EMAIL.replace(":id", id);
  const { data } = await protectedClient.patch<BaseResponse>(url, params);
  return data;
};

// 验证用户邮箱
export const VerifyUserEmail = async (id: string, params?: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.VERIFY_USER_EMAIL.replace(":id", id);
  const { data } = await protectedClient.patch<BaseResponse>(url, params);
  return data;
};

// 删除用户邮箱
export const DeleteUserEmail = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.DELETE_USER_EMAIL.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 用户职业相关 ==========

// 创建用户职业
export const CreateUserOccupation = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.DIRECTORY.CREATE_USER_OCCUPATION, params);
  return data.data;
};

// 根据 ID 获取用户职业
export const GetUserOccupation = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.DIRECTORY.GET_USER_OCCUPATION.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新用户职业
export const UpdateUserOccupation = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.UPDATE_USER_OCCUPATION.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除用户职业
export const DeleteUserOccupation = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.DELETE_USER_OCCUPATION.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 用户电话相关 ==========

// 创建用户电话
export const CreateUserPhone = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.DIRECTORY.CREATE_USER_PHONE, params);
  return data.data;
};

// 根据 ID 获取用户电话
export const GetUserPhone = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.DIRECTORY.GET_USER_PHONE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新用户电话
export const UpdateUserPhone = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.UPDATE_USER_PHONE.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 设置主电话
export const SetPrimaryUserPhone = async (id: string, params?: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.SET_PRIMARY_USER_PHONE.replace(":id", id);
  const { data } = await protectedClient.patch<BaseResponse>(url, params);
  return data;
};

// 验证用户电话
export const VerifyUserPhone = async (id: string, params?: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.VERIFY_USER_PHONE.replace(":id", id);
  const { data } = await protectedClient.patch<BaseResponse>(url, params);
  return data;
};

// 删除用户电话
export const DeleteUserPhone = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.DELETE_USER_PHONE.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 用户偏好相关 ==========

// 创建用户偏好
export const CreateUserPreference = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.DIRECTORY.CREATE_USER_PREFERENCE, params);
  return data.data;
};

// 根据 ID 获取用户偏好
export const GetUserPreference = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.DIRECTORY.GET_USER_PREFERENCE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新用户偏好
export const UpdateUserPreference = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.UPDATE_USER_PREFERENCE.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除用户偏好
export const DeleteUserPreference = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.DELETE_USER_PREFERENCE.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 用户资料相关 ==========

// 创建用户资料
export const CreateUserProfile = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.DIRECTORY.CREATE_USER_PROFILE, params);
  return data.data;
};

// 根据 ID 获取用户资料
export const GetUserProfile = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.DIRECTORY.GET_USER_PROFILE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新用户资料
export const UpdateUserProfile = async (id: string, params: unknown): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.UPDATE_USER_PROFILE.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除用户资料
export const DeleteUserProfile = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.DIRECTORY.DELETE_USER_PROFILE.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};
