import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";

import {
  DeleteSystemState,
  GetSystemState,
  GetSystemStateLatest,
  GetSystemStateLatestPublic,
  InitializeSystemState,
  ResetSystemState,
} from "@/apis/system.api";
import type { InitializeSystemStateRequest, ResetSystemStateRequest, SystemState } from "@/types";
import { makeUnifiedQuery } from "@/hooks/core/makeUnifiedQuery";
import { systemEventEmitter, systemEvents } from "@/events/system";
import { showError, showSuccess } from "@/stores/modalStore";
import { SYSTEM_SYSTEM_STATE, SYSTEM_SYSTEM_STATE_LATEST, SYSTEM_SYSTEM_STATE_INIT } from "@/constants";
import type { UnifiedQueryParams } from "./core/type";

// ========== SystemState 相关 ==========

// 检查系统是否已初始化（公开接口，不需要认证）
export const useSystemInit = (params?: UnifiedQueryParams<boolean>) => {
  const { options, postProcess } = params || {};
  const makeQuery = makeUnifiedQuery(
    async () => {
      try {
        const systemState = await GetSystemStateLatestPublic();
        return systemState.initialized;
      } catch (err) {
        // If API returns 404 or no record, system is not initialized
        // If API returns other error, we'll treat it as not initialized for safety
        console.warn("Failed to check system initialization status:", err);
        return false;
      }
    },
    undefined, // 使用 normal 模式（非 suspense），返回 UseQueryResult
    postProcess,
  );
  const queryResult = makeQuery(SYSTEM_SYSTEM_STATE_INIT, {}, {
    retry: 1, // Only retry once
    staleTime: 5 * 60 * 1000, // Cache for 5 minutes
    ...options,
  });

  return {
    isInitialized: queryResult.data,
    isLoading: queryResult.isLoading,
    error: queryResult.error as Error | null,
  };
};

// 获取最新系统状态
export const useSystemStateLatest = (params?: UnifiedQueryParams<SystemState>) => {
  const { options, postProcess } = params || {};
  const makeQuery = makeUnifiedQuery(
    async () => {
      return await GetSystemStateLatest();
    },
    "suspense",
    postProcess,
  );
  return makeQuery(SYSTEM_SYSTEM_STATE_LATEST, {}, options);
};

// 根据 ID 获取系统状态
export const useSystemState = (params: UnifiedQueryParams<SystemState> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetSystemState(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(SYSTEM_SYSTEM_STATE(id), { id }, options);
};

// 初始化系统状态
export const useInitializeSystemState = () => {
  return useMutation({
    mutationFn: async (params?: InitializeSystemStateRequest) => {
      return await InitializeSystemState(params);
    },
    onSuccess: () => {
      systemEventEmitter.emit(systemEvents.INVALIDATE_SYSTEM_STATES);
      showSuccess("系统状态初始化成功！");
    },
    onError: (error: AxiosError) => {
      showError("初始化系统状态失败，请稍后重试。" + error.message);
    },
  });
};

// 重置系统状态
export const useResetSystemState = () => {
  return useMutation({
    mutationFn: async (params?: ResetSystemStateRequest) => {
      return await ResetSystemState(params);
    },
    onSuccess: () => {
      systemEventEmitter.emit(systemEvents.INVALIDATE_SYSTEM_STATES);
      showSuccess("系统状态重置成功！");
    },
    onError: (error: AxiosError) => {
      showError("重置系统状态失败，请稍后重试。" + error.message);
    },
  });
};

// 删除系统状态
export const useDeleteSystemState = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteSystemState(id);
    },
    onSuccess: (_, id) => {
      systemEventEmitter.emit(systemEvents.INVALIDATE_SYSTEM_STATES);
      systemEventEmitter.emit(systemEvents.INVALIDATE_SYSTEM_STATE, id);
      showSuccess("系统状态删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除系统状态失败，请稍后重试。" + error.message);
    },
  });
};
