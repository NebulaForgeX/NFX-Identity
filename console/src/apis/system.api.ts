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

// ========== 错误码翻译（公开，供 i18n 按语言加载） ==========

/** 后端错误码翻译：code -> 当前语言的文案 */
export type ErrorTranslations = Record<string, string>;

/** 获取指定语言的错误码翻译 JSON（从后端挂载的 errors/langs 读取，外部更新即生效） */
export const getErrorTranslations = async (lang: string): Promise<ErrorTranslations> => {
  const url = URL_PATHS.SYSTEM.i18nErrors.byLang(lang);
  const { data } = await publicClient.get<ErrorTranslations>(url);
  return data ?? {};
};

// ========== 系统状态相关 ==========

// 获取最新系统状态（公开接口，不需要认证）
export const GetSystemStateLatestPublic = async (): Promise<SystemState> => {
  const { data } = await publicClient.get<DataResponse<SystemState>>(
    URL_PATHS.SYSTEM.systemState.latest,
  );
  return data.data;
};

// 获取最新系统状态（需要认证）
export const GetSystemStateLatest = async (): Promise<SystemState> => {
  const { data } = await protectedClient.get<DataResponse<SystemState>>(
    URL_PATHS.SYSTEM.systemState.latest,
  );
  return data.data;
};

// 根据 ID 获取系统状态
export const GetSystemState = async (id: string): Promise<SystemState> => {
  const { data } = await protectedClient.get<DataResponse<SystemState>>(
    URL_PATHS.SYSTEM.systemState.byId(id),
  );
  return data.data;
};

// 初始化系统状态（公开接口，不需要认证；超时 2 分钟，初始化耗时长）
const INITIALIZE_TIMEOUT_MS = 240_000;

export const InitializeSystemState = async (params?: InitializeSystemStateRequest): Promise<SystemState> => {
  const { data } = await publicClient.post<DataResponse<SystemState>>(
    URL_PATHS.SYSTEM.systemState.initialize,
    params,
    { timeout: INITIALIZE_TIMEOUT_MS },
  );
  return data.data;
};

// 重置系统状态
export const ResetSystemState = async (params?: ResetSystemStateRequest): Promise<SystemState> => {
  const { data } = await protectedClient.post<DataResponse<SystemState>>(
    URL_PATHS.SYSTEM.systemStateAuth.reset,
    params,
  );
  return data.data;
};

// 删除系统状态
export const DeleteSystemState = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.SYSTEM.systemStateAuth.byId(id),
  );
  return data;
};
