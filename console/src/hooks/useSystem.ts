import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";

import {
  DeleteSystemState,
  GetSystemState,
  GetSystemStateLatest,
  InitializeSystemState,
  ResetSystemState,
} from "@/apis/system.api";
import type { InitializeSystemStateRequest, ResetSystemStateRequest } from "@/types";
import { makeUnifiedQuery } from "@/hooks/core/makeUnifiedQuery";
import { systemEventEmitter, systemEvents } from "@/events/system";
import { showError, showSuccess } from "@/stores/modalStore";

// ========== SystemState 相关 ==========

// 获取最新系统状态
export const useSystemStateLatest = makeUnifiedQuery(
  async () => {
    return await GetSystemStateLatest();
  },
  "normal",
);

// 根据 ID 获取系统状态
export const useSystemState = makeUnifiedQuery(
  async (params: { id: string }) => {
    return await GetSystemState(params.id);
  },
  "normal",
);

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
