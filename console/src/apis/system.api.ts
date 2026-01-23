// System API - 基于 NFX-ID Backend

import type {
  BaseResponse,
  DataResponse,
  InitializeSystemStateRequest,
  ResetSystemStateRequest,
  SystemState,
} from "@/types";

import { protectedClient, publicClient } from "./clients";
import { URL_PATHS } from "./ip";

// ========== 系统状态相关 ==========

// 获取最新系统状态（公开接口，不需要认证）
export const GetSystemStateLatestPublic = async (): Promise<SystemState> => {
  const { data } = await publicClient.get<DataResponse<SystemState>>(URL_PATHS.SYSTEM.GET_SYSTEM_STATE_LATEST);
  return data.data;
};

// 获取最新系统状态（需要认证）
export const GetSystemStateLatest = async (): Promise<SystemState> => {
  const { data } = await protectedClient.get<DataResponse<SystemState>>(URL_PATHS.SYSTEM.GET_SYSTEM_STATE_LATEST);
  return data.data;
};

// 根据 ID 获取系统状态
export const GetSystemState = async (id: string): Promise<SystemState> => {
  const url = URL_PATHS.SYSTEM.GET_SYSTEM_STATE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<SystemState>>(url);
  return data.data;
};

// 初始化系统状态（公开接口，不需要认证；超时 2 分钟，初始化耗时长）
const INITIALIZE_TIMEOUT_MS = 240_000;

export const InitializeSystemState = async (params?: InitializeSystemStateRequest): Promise<SystemState> => {
  const { data } = await publicClient.post<DataResponse<SystemState>>(
    URL_PATHS.SYSTEM.INITIALIZE_SYSTEM_STATE,
    params,
    { timeout: INITIALIZE_TIMEOUT_MS },
  );
  return data.data;
};

// 重置系统状态
export const ResetSystemState = async (params?: ResetSystemStateRequest): Promise<SystemState> => {
  const { data } = await protectedClient.post<DataResponse<SystemState>>(URL_PATHS.SYSTEM.RESET_SYSTEM_STATE, params);
  return data.data;
};

// 删除系统状态
export const DeleteSystemState = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.SYSTEM.DELETE_SYSTEM_STATE.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};
