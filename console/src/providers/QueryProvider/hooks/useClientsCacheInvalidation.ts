import type { QueryClient } from "@tanstack/react-query";
import { clientsEventEmitter, clientsEvents } from "@/events/clients";
import { CLIENTS_QUERY_KEY_PREFIXES } from "@/constants";

/**
 * Clients 相关的缓存失效事件处理
 */
export const useClientsCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateApps = () => queryClient.invalidateQueries({ queryKey: CLIENTS_QUERY_KEY_PREFIXES.APPS });
  const handleInvalidateApp = (item: string) => queryClient.invalidateQueries({ queryKey: [...CLIENTS_QUERY_KEY_PREFIXES.APP, item] });
  const handleInvalidateAPIKeys = () => queryClient.invalidateQueries({ queryKey: CLIENTS_QUERY_KEY_PREFIXES.API_KEYS });
  const handleInvalidateAPIKey = (item: string) => queryClient.invalidateQueries({ queryKey: [...CLIENTS_QUERY_KEY_PREFIXES.API_KEY, item] });
  const handleInvalidateClientCredentials = () => queryClient.invalidateQueries({ queryKey: CLIENTS_QUERY_KEY_PREFIXES.CLIENT_CREDENTIALS });
  const handleInvalidateClientCredential = (item: string) => queryClient.invalidateQueries({ queryKey: [...CLIENTS_QUERY_KEY_PREFIXES.CLIENT_CREDENTIAL, item] });
  const handleInvalidateClientScopes = () => queryClient.invalidateQueries({ queryKey: CLIENTS_QUERY_KEY_PREFIXES.CLIENT_SCOPES });
  const handleInvalidateClientScope = (item: string) => queryClient.invalidateQueries({ queryKey: [...CLIENTS_QUERY_KEY_PREFIXES.CLIENT_SCOPE, item] });
  const handleInvalidateIPAllowlists = () => queryClient.invalidateQueries({ queryKey: CLIENTS_QUERY_KEY_PREFIXES.IP_ALLOWLISTS });
  const handleInvalidateIPAllowlist = (item: string) => queryClient.invalidateQueries({ queryKey: [...CLIENTS_QUERY_KEY_PREFIXES.IP_ALLOWLIST, item] });
  const handleInvalidateRateLimits = () => queryClient.invalidateQueries({ queryKey: CLIENTS_QUERY_KEY_PREFIXES.RATE_LIMITS });
  const handleInvalidateRateLimit = (item: string) => queryClient.invalidateQueries({ queryKey: [...CLIENTS_QUERY_KEY_PREFIXES.RATE_LIMIT, item] });

  // 注册监听器
  clientsEventEmitter.on(clientsEvents.INVALIDATE_APPS, handleInvalidateApps);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_APP, handleInvalidateApp);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_API_KEYS, handleInvalidateAPIKeys);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_API_KEY, handleInvalidateAPIKey);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_CLIENT_CREDENTIALS, handleInvalidateClientCredentials);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_CLIENT_CREDENTIAL, handleInvalidateClientCredential);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_CLIENT_SCOPES, handleInvalidateClientScopes);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_CLIENT_SCOPE, handleInvalidateClientScope);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_IP_ALLOWLISTS, handleInvalidateIPAllowlists);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_IP_ALLOWLIST, handleInvalidateIPAllowlist);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_RATE_LIMITS, handleInvalidateRateLimits);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_RATE_LIMIT, handleInvalidateRateLimit);

  // 清理监听器
  return () => {
    clientsEventEmitter.off(clientsEvents.INVALIDATE_APPS, handleInvalidateApps);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_APP, handleInvalidateApp);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_API_KEYS, handleInvalidateAPIKeys);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_API_KEY, handleInvalidateAPIKey);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_CLIENT_CREDENTIALS, handleInvalidateClientCredentials);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_CLIENT_CREDENTIAL, handleInvalidateClientCredential);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_CLIENT_SCOPES, handleInvalidateClientScopes);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_CLIENT_SCOPE, handleInvalidateClientScope);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_IP_ALLOWLISTS, handleInvalidateIPAllowlists);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_IP_ALLOWLIST, handleInvalidateIPAllowlist);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_RATE_LIMITS, handleInvalidateRateLimits);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_RATE_LIMIT, handleInvalidateRateLimit);
  };
};
