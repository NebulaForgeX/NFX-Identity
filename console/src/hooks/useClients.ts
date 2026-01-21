import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";

import {
  CreateAPIKey,
  CreateApp,
  CreateClientCredential,
  CreateClientScope,
  CreateIPAllowlist,
  CreateRateLimit,
  DeleteAPIKeyByKeyID,
  DeleteApp,
  DeleteClientCredentialByClientID,
  DeleteClientScope,
  DeleteIPAllowlistByRuleID,
  DeleteRateLimit,
  GetAPIKey,
  GetApp,
  GetAppByAppID,
  GetClientCredential,
  GetClientScope,
  GetIPAllowlist,
  GetRateLimit,
  UpdateApp,
} from "@/apis/clients.api";
import type {
  APIKey,
  App,
  ClientCredential,
  ClientScope,
  CreateAPIKeyRequest,
  CreateAppRequest,
  CreateClientCredentialRequest,
  CreateClientScopeRequest,
  CreateIPAllowlistRequest,
  CreateRateLimitRequest,
  IPAllowlist,
  RateLimit,
  UpdateAppRequest,
} from "@/types";
import { makeUnifiedQuery } from "@/hooks/core/makeUnifiedQuery";
import { clientsEventEmitter, clientsEvents } from "@/events/clients";
import { showError, showSuccess } from "@/stores/modalStore";
import {
  CLIENTS_APP,
  CLIENTS_API_KEY,
  CLIENTS_CLIENT_CREDENTIAL,
  CLIENTS_CLIENT_SCOPE,
  CLIENTS_IP_ALLOWLIST,
  CLIENTS_RATE_LIMIT,
} from "@/constants";
import type { UnifiedQueryParams } from "./core/type";

// ========== App 相关 ==========

// 根据 ID 获取应用
export const useApp = (params: UnifiedQueryParams<App> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetApp(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(CLIENTS_APP(id), { id }, options);
};

// 根据 App ID 获取应用
export const useAppByAppID = (params: UnifiedQueryParams<App> & { appId: string }) => {
  const { appId, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { appId: string }) => {
      return await GetAppByAppID(params.appId);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(CLIENTS_APP(appId), { appId }, options);
};

// 创建应用
export const useCreateApp = () => {
  return useMutation({
    mutationFn: async (params: CreateAppRequest) => {
      return await CreateApp(params);
    },
    onSuccess: () => {
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_APPS);
      showSuccess("应用创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建应用失败，请稍后重试。" + error.message);
    },
  });
};

// 更新应用
export const useUpdateApp = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateAppRequest }) => {
      return await UpdateApp(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_APPS);
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_APP, variables.id);
      showSuccess("应用更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新应用失败，请稍后重试。" + error.message);
    },
  });
};

// 删除应用
export const useDeleteApp = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteApp(id);
    },
    onSuccess: (_, id) => {
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_APPS);
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_APP, id);
      showSuccess("应用删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除应用失败，请稍后重试。" + error.message);
    },
  });
};

// ========== APIKey 相关 ==========

// 根据 ID 获取 API Key
export const useAPIKey = (params: UnifiedQueryParams<APIKey> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetAPIKey(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(CLIENTS_API_KEY(id), { id }, options);
};

// 创建 API Key
export const useCreateAPIKey = () => {
  return useMutation({
    mutationFn: async (params: CreateAPIKeyRequest) => {
      return await CreateAPIKey(params);
    },
    onSuccess: () => {
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_API_KEYS);
      showSuccess("API Key 创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建 API Key 失败，请稍后重试。" + error.message);
    },
  });
};

// 根据 Key ID 删除 API Key
export const useDeleteAPIKeyByKeyID = () => {
  return useMutation({
    mutationFn: async (keyId: string) => {
      return await DeleteAPIKeyByKeyID(keyId);
    },
    onSuccess: (_, keyId) => {
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_API_KEYS);
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_API_KEY, keyId);
      showSuccess("API Key 删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除 API Key 失败，请稍后重试。" + error.message);
    },
  });
};

// ========== ClientCredential 相关 ==========

// 根据 ID 获取 Client Credential
export const useClientCredential = (params: UnifiedQueryParams<ClientCredential> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetClientCredential(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(CLIENTS_CLIENT_CREDENTIAL(id), { id }, options);
};

// 创建 Client Credential
export const useCreateClientCredential = () => {
  return useMutation({
    mutationFn: async (params: CreateClientCredentialRequest) => {
      return await CreateClientCredential(params);
    },
    onSuccess: () => {
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_CLIENT_CREDENTIALS);
      showSuccess("Client Credential 创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建 Client Credential 失败，请稍后重试。" + error.message);
    },
  });
};

// 根据 Client ID 删除 Client Credential
export const useDeleteClientCredentialByClientID = () => {
  return useMutation({
    mutationFn: async (clientId: string) => {
      return await DeleteClientCredentialByClientID(clientId);
    },
    onSuccess: (_, clientId) => {
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_CLIENT_CREDENTIALS);
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_CLIENT_CREDENTIAL, clientId);
      showSuccess("Client Credential 删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除 Client Credential 失败，请稍后重试。" + error.message);
    },
  });
};

// ========== ClientScope 相关 ==========

// 根据 ID 获取 Client Scope
export const useClientScope = (params: UnifiedQueryParams<ClientScope> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetClientScope(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(CLIENTS_CLIENT_SCOPE(id), { id }, options);
};

// 创建 Client Scope
export const useCreateClientScope = () => {
  return useMutation({
    mutationFn: async (params: CreateClientScopeRequest) => {
      return await CreateClientScope(params);
    },
    onSuccess: () => {
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_CLIENT_SCOPES);
      showSuccess("Client Scope 创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建 Client Scope 失败，请稍后重试。" + error.message);
    },
  });
};

// 删除 Client Scope
export const useDeleteClientScope = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteClientScope(id);
    },
    onSuccess: (_, id) => {
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_CLIENT_SCOPES);
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_CLIENT_SCOPE, id);
      showSuccess("Client Scope 删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除 Client Scope 失败，请稍后重试。" + error.message);
    },
  });
};

// ========== IPAllowlist 相关 ==========

// 根据 ID 获取 IP Allowlist
export const useIPAllowlist = (params: UnifiedQueryParams<IPAllowlist> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetIPAllowlist(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(CLIENTS_IP_ALLOWLIST(id), { id }, options);
};

// 创建 IP Allowlist
export const useCreateIPAllowlist = () => {
  return useMutation({
    mutationFn: async (params: CreateIPAllowlistRequest) => {
      return await CreateIPAllowlist(params);
    },
    onSuccess: () => {
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_IP_ALLOWLISTS);
      showSuccess("IP Allowlist 创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建 IP Allowlist 失败，请稍后重试。" + error.message);
    },
  });
};

// 根据 Rule ID 删除 IP Allowlist
export const useDeleteIPAllowlistByRuleID = () => {
  return useMutation({
    mutationFn: async (ruleId: string) => {
      return await DeleteIPAllowlistByRuleID(ruleId);
    },
    onSuccess: (_, ruleId) => {
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_IP_ALLOWLISTS);
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_IP_ALLOWLIST, ruleId);
      showSuccess("IP Allowlist 删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除 IP Allowlist 失败，请稍后重试。" + error.message);
    },
  });
};

// ========== RateLimit 相关 ==========

// 根据 ID 获取 Rate Limit
export const useRateLimit = (params: UnifiedQueryParams<RateLimit> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetRateLimit(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(CLIENTS_RATE_LIMIT(id), { id }, options);
};

// 创建 Rate Limit
export const useCreateRateLimit = () => {
  return useMutation({
    mutationFn: async (params: CreateRateLimitRequest) => {
      return await CreateRateLimit(params);
    },
    onSuccess: () => {
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_RATE_LIMITS);
      showSuccess("Rate Limit 创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建 Rate Limit 失败，请稍后重试。" + error.message);
    },
  });
};

// 删除 Rate Limit
export const useDeleteRateLimit = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteRateLimit(id);
    },
    onSuccess: (_, id) => {
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_RATE_LIMITS);
      clientsEventEmitter.emit(clientsEvents.INVALIDATE_RATE_LIMIT, id);
      showSuccess("Rate Limit 删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除 Rate Limit 失败，请稍后重试。" + error.message);
    },
  });
};
