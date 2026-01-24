import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import { useTranslation } from "react-i18next";

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

// 获取最新系统状态（公开接口，不需要认证）- suspense 模式
export const useSystemInit = (params?: UnifiedQueryParams<SystemState>) => {
  const { options, postProcess } = params || {};
  const makeQuery = makeUnifiedQuery(
    async () => {
      return await GetSystemStateLatestPublic();
    },
    "suspense",
    postProcess,
  );
  return makeQuery(SYSTEM_SYSTEM_STATE_INIT, {}, options);
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
  const { t } = useTranslation("hooks.system");
  return useMutation({
    mutationFn: async (params?: InitializeSystemStateRequest) => {
      return await InitializeSystemState(params);
    },
    onSuccess: () => {
      systemEventEmitter.emit(systemEvents.INVALIDATE_SYSTEM_STATES);
      showSuccess(t("systemState.initializeSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("systemState.initializeError") + error.message);
    },
  });
};

// 重置系统状态
export const useResetSystemState = () => {
  const { t } = useTranslation("hooks.system");
  return useMutation({
    mutationFn: async (params?: ResetSystemStateRequest) => {
      return await ResetSystemState(params);
    },
    onSuccess: () => {
      systemEventEmitter.emit(systemEvents.INVALIDATE_SYSTEM_STATES);
      showSuccess(t("systemState.resetSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("systemState.resetError") + error.message);
    },
  });
};

// 删除系统状态
export const useDeleteSystemState = () => {
  const { t } = useTranslation("hooks.system");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteSystemState(id);
    },
    onSuccess: (_, id) => {
      systemEventEmitter.emit(systemEvents.INVALIDATE_SYSTEM_STATES);
      systemEventEmitter.emit(systemEvents.INVALIDATE_SYSTEM_STATE, id);
      showSuccess(t("systemState.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("systemState.deleteError") + error.message);
    },
  });
};
