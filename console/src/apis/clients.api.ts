// Clients API - 基于 NFX-ID Backend

import type { BaseResponse, DataResponse } from "nfx-ui/types";
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

import { protectedClient } from "./clients";
import { URL_PATHS } from "./ip";

// ========== 应用相关 ==========

// 创建应用
export const CreateApp = async (params: CreateAppRequest): Promise<App> => {
  const { data } = await protectedClient.post<DataResponse<App>>(
    URL_PATHS.CLIENTS.apps,
    params,
  );
  return data.data;
};

// 根据 ID 获取应用
export const GetApp = async (id: string): Promise<App> => {
  const { data } = await protectedClient.get<DataResponse<App>>(
    URL_PATHS.CLIENTS.apps.byId(id),
  );
  return data.data;
};

// 根据 App ID 获取应用
export const GetAppByAppID = async (appId: string): Promise<App> => {
  const { data } = await protectedClient.get<DataResponse<App>>(
    URL_PATHS.CLIENTS.apps.byAppId(appId),
  );
  return data.data;
};

// 更新应用
export const UpdateApp = async (id: string, params: UpdateAppRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(
    URL_PATHS.CLIENTS.apps.byId(id),
    params,
  );
  return data;
};

// 删除应用
export const DeleteApp = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.CLIENTS.apps.byId(id),
  );
  return data;
};

// ========== API Key 相关 ==========

// 创建 API Key
export const CreateAPIKey = async (params: CreateAPIKeyRequest): Promise<APIKey> => {
  const { data } = await protectedClient.post<DataResponse<APIKey>>(
    URL_PATHS.CLIENTS.apiKeys,
    params,
  );
  return data.data;
};

// 根据 ID 获取 API Key
export const GetAPIKey = async (id: string): Promise<APIKey> => {
  const { data } = await protectedClient.get<DataResponse<APIKey>>(
    URL_PATHS.CLIENTS.apiKeys.byId(id),
  );
  return data.data;
};

// 根据 Key ID 删除 API Key
export const DeleteAPIKeyByKeyID = async (keyId: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.CLIENTS.apiKeys.byKeyId(keyId),
  );
  return data;
};

// ========== Client Credential 相关 ==========

// 创建 Client Credential
export const CreateClientCredential = async (params: CreateClientCredentialRequest): Promise<ClientCredential> => {
  const { data } = await protectedClient.post<DataResponse<ClientCredential>>(
    URL_PATHS.CLIENTS.clientCredentials,
    params,
  );
  return data.data;
};

// 根据 ID 获取 Client Credential
export const GetClientCredential = async (id: string): Promise<ClientCredential> => {
  const { data } = await protectedClient.get<DataResponse<ClientCredential>>(
    URL_PATHS.CLIENTS.clientCredentials.byId(id),
  );
  return data.data;
};

// 根据 Client ID 删除 Client Credential
export const DeleteClientCredentialByClientID = async (clientId: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.CLIENTS.clientCredentials.byClientId(clientId),
  );
  return data;
};

// ========== Client Scope 相关 ==========

// 创建 Client Scope
export const CreateClientScope = async (params: CreateClientScopeRequest): Promise<ClientScope> => {
  const { data } = await protectedClient.post<DataResponse<ClientScope>>(
    URL_PATHS.CLIENTS.clientScopes,
    params,
  );
  return data.data;
};

// 根据 ID 获取 Client Scope
export const GetClientScope = async (id: string): Promise<ClientScope> => {
  const { data } = await protectedClient.get<DataResponse<ClientScope>>(
    URL_PATHS.CLIENTS.clientScopes.byId(id),
  );
  return data.data;
};

// 删除 Client Scope
export const DeleteClientScope = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.CLIENTS.clientScopes.byId(id),
  );
  return data;
};

// ========== IP Allowlist 相关 ==========

// 创建 IP Allowlist
export const CreateIPAllowlist = async (params: CreateIPAllowlistRequest): Promise<IPAllowlist> => {
  const { data } = await protectedClient.post<DataResponse<IPAllowlist>>(
    URL_PATHS.CLIENTS.ipAllowlist,
    params,
  );
  return data.data;
};

// 根据 ID 获取 IP Allowlist
export const GetIPAllowlist = async (id: string): Promise<IPAllowlist> => {
  const { data } = await protectedClient.get<DataResponse<IPAllowlist>>(
    URL_PATHS.CLIENTS.ipAllowlist.byId(id),
  );
  return data.data;
};

// 根据 Rule ID 删除 IP Allowlist
export const DeleteIPAllowlistByRuleID = async (ruleId: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.CLIENTS.ipAllowlist.byRuleId(ruleId),
  );
  return data;
};

// ========== Rate Limit 相关 ==========

// 创建 Rate Limit
export const CreateRateLimit = async (params: CreateRateLimitRequest): Promise<RateLimit> => {
  const { data } = await protectedClient.post<DataResponse<RateLimit>>(
    URL_PATHS.CLIENTS.rateLimits,
    params,
  );
  return data.data;
};

// 根据 ID 获取 Rate Limit
export const GetRateLimit = async (id: string): Promise<RateLimit> => {
  const { data } = await protectedClient.get<DataResponse<RateLimit>>(
    URL_PATHS.CLIENTS.rateLimits.byId(id),
  );
  return data.data;
};

// 删除 Rate Limit
export const DeleteRateLimit = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.CLIENTS.rateLimits.byId(id),
  );
  return data;
};
