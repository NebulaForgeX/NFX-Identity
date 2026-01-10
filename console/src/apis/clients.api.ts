// Clients API - 基于 NFX-ID Backend

import type {
  APIKey,
  App,
  BaseResponse,
  ClientCredential,
  ClientScope,
  CreateAPIKeyRequest,
  CreateAppRequest,
  CreateClientCredentialRequest,
  CreateClientScopeRequest,
  CreateIPAllowlistRequest,
  CreateRateLimitRequest,
  DataResponse,
  IPAllowlist,
  RateLimit,
  UpdateAppRequest,
} from "@/types";

import { protectedClient } from "./clients";
import { URL_PATHS } from "./ip";

// ========== 应用相关 ==========

// 创建应用
export const CreateApp = async (params: CreateAppRequest): Promise<App> => {
  const { data } = await protectedClient.post<DataResponse<App>>(URL_PATHS.CLIENTS.CREATE_APP, params);
  return data.data;
};

// 根据 ID 获取应用
export const GetApp = async (id: string): Promise<App> => {
  const url = URL_PATHS.CLIENTS.GET_APP.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<App>>(url);
  return data.data;
};

// 根据 App ID 获取应用
export const GetAppByAppID = async (appId: string): Promise<App> => {
  const url = URL_PATHS.CLIENTS.GET_APP_BY_APP_ID.replace(":app_id", appId);
  const { data } = await protectedClient.get<DataResponse<App>>(url);
  return data.data;
};

// 更新应用
export const UpdateApp = async (id: string, params: UpdateAppRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.CLIENTS.UPDATE_APP.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除应用
export const DeleteApp = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.CLIENTS.DELETE_APP.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== API Key 相关 ==========

// 创建 API Key
export const CreateAPIKey = async (params: CreateAPIKeyRequest): Promise<APIKey> => {
  const { data } = await protectedClient.post<DataResponse<APIKey>>(URL_PATHS.CLIENTS.CREATE_API_KEY, params);
  return data.data;
};

// 根据 ID 获取 API Key
export const GetAPIKey = async (id: string): Promise<APIKey> => {
  const url = URL_PATHS.CLIENTS.GET_API_KEY.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<APIKey>>(url);
  return data.data;
};

// 根据 Key ID 删除 API Key
export const DeleteAPIKeyByKeyID = async (keyId: string): Promise<BaseResponse> => {
  const url = URL_PATHS.CLIENTS.DELETE_API_KEY_BY_KEY_ID.replace(":key_id", keyId);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== Client Credential 相关 ==========

// 创建 Client Credential
export const CreateClientCredential = async (params: CreateClientCredentialRequest): Promise<ClientCredential> => {
  const { data } = await protectedClient.post<DataResponse<ClientCredential>>(
    URL_PATHS.CLIENTS.CREATE_CLIENT_CREDENTIAL,
    params,
  );
  return data.data;
};

// 根据 ID 获取 Client Credential
export const GetClientCredential = async (id: string): Promise<ClientCredential> => {
  const url = URL_PATHS.CLIENTS.GET_CLIENT_CREDENTIAL.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<ClientCredential>>(url);
  return data.data;
};

// 根据 Client ID 删除 Client Credential
export const DeleteClientCredentialByClientID = async (clientId: string): Promise<BaseResponse> => {
  const url = URL_PATHS.CLIENTS.DELETE_CLIENT_CREDENTIAL_BY_CLIENT_ID.replace(":client_id", clientId);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== Client Scope 相关 ==========

// 创建 Client Scope
export const CreateClientScope = async (params: CreateClientScopeRequest): Promise<ClientScope> => {
  const { data } = await protectedClient.post<DataResponse<ClientScope>>(URL_PATHS.CLIENTS.CREATE_CLIENT_SCOPE, params);
  return data.data;
};

// 根据 ID 获取 Client Scope
export const GetClientScope = async (id: string): Promise<ClientScope> => {
  const url = URL_PATHS.CLIENTS.GET_CLIENT_SCOPE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<ClientScope>>(url);
  return data.data;
};

// 删除 Client Scope
export const DeleteClientScope = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.CLIENTS.DELETE_CLIENT_SCOPE.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== IP Allowlist 相关 ==========

// 创建 IP Allowlist
export const CreateIPAllowlist = async (params: CreateIPAllowlistRequest): Promise<IPAllowlist> => {
  const { data } = await protectedClient.post<DataResponse<IPAllowlist>>(URL_PATHS.CLIENTS.CREATE_IP_ALLOWLIST, params);
  return data.data;
};

// 根据 ID 获取 IP Allowlist
export const GetIPAllowlist = async (id: string): Promise<IPAllowlist> => {
  const url = URL_PATHS.CLIENTS.GET_IP_ALLOWLIST.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<IPAllowlist>>(url);
  return data.data;
};

// 根据 Rule ID 删除 IP Allowlist
export const DeleteIPAllowlistByRuleID = async (ruleId: string): Promise<BaseResponse> => {
  const url = URL_PATHS.CLIENTS.DELETE_IP_ALLOWLIST_BY_RULE_ID.replace(":rule_id", ruleId);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== Rate Limit 相关 ==========

// 创建 Rate Limit
export const CreateRateLimit = async (params: CreateRateLimitRequest): Promise<RateLimit> => {
  const { data } = await protectedClient.post<DataResponse<RateLimit>>(URL_PATHS.CLIENTS.CREATE_RATE_LIMIT, params);
  return data.data;
};

// 根据 ID 获取 Rate Limit
export const GetRateLimit = async (id: string): Promise<RateLimit> => {
  const url = URL_PATHS.CLIENTS.GET_RATE_LIMIT.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<RateLimit>>(url);
  return data.data;
};

// 删除 Rate Limit
export const DeleteRateLimit = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.CLIENTS.DELETE_RATE_LIMIT.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};
