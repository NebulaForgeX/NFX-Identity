// System API - 基于 NFX-ID Backend

import type { BaseResponse, DataResponse } from "@/types/api";

import { protectedClient } from "@/apis/clients";
import { URL_PATHS } from "@/apis/ip";

// ========== 系统状态相关 ==========

// 获取最新系统状态
export const GetSystemStateLatest = async (): Promise<unknown> => {
  const { data } = await protectedClient.get<DataResponse<unknown>>(URL_PATHS.SYSTEM.GET_SYSTEM_STATE_LATEST);
  return data.data;
};

// 根据 ID 获取系统状态
export const GetSystemState = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.SYSTEM.GET_SYSTEM_STATE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 初始化系统状态
export const InitializeSystemState = async (params?: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.SYSTEM.INITIALIZE_SYSTEM_STATE, params);
  return data.data;
};

// 重置系统状态
export const ResetSystemState = async (params?: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.SYSTEM.RESET_SYSTEM_STATE, params);
  return data.data;
};

// 删除系统状态
export const DeleteSystemState = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.SYSTEM.DELETE_SYSTEM_STATE.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};
