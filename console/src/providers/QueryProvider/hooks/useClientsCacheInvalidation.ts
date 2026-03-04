import type { QueryClient } from "@tanstack/react-query";
import { clientsEventEmitter, clientsEvents } from "@/events/clients";
import {
  CLIENTS_API_KEY,
  CLIENTS_API_KEY_LIST,
  CLIENTS_APP,
  CLIENTS_APP_LIST,
  CLIENTS_CLIENT_CREDENTIAL,
  CLIENTS_CLIENT_CREDENTIAL_LIST,
  CLIENTS_CLIENT_SCOPE,
  CLIENTS_CLIENT_SCOPE_LIST,
  CLIENTS_IP_ALLOWLIST,
  CLIENTS_IP_ALLOWLIST_LIST,
  CLIENTS_RATE_LIMIT,
  CLIENTS_RATE_LIMIT_LIST,
} from "@/constants";

type EventCb = (...args: unknown[]) => void;

/**
 * Clients 相关的缓存失效事件处理
 */
export const useClientsCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateApps = () => queryClient.invalidateQueries({ queryKey: CLIENTS_APP_LIST });
  const handleInvalidateApp = (item: string) => queryClient.invalidateQueries({ queryKey: CLIENTS_APP(item) });
  const handleInvalidateAPIKeys = () => queryClient.invalidateQueries({ queryKey: CLIENTS_API_KEY_LIST });
  const handleInvalidateAPIKey = (item: string) => queryClient.invalidateQueries({ queryKey: CLIENTS_API_KEY(item) });
  const handleInvalidateClientCredentials = () => queryClient.invalidateQueries({ queryKey: CLIENTS_CLIENT_CREDENTIAL_LIST });
  const handleInvalidateClientCredential = (item: string) => queryClient.invalidateQueries({ queryKey: CLIENTS_CLIENT_CREDENTIAL(item) });
  const handleInvalidateClientScopes = () => queryClient.invalidateQueries({ queryKey: CLIENTS_CLIENT_SCOPE_LIST });
  const handleInvalidateClientScope = (item: string) => queryClient.invalidateQueries({ queryKey: CLIENTS_CLIENT_SCOPE(item) });
  const handleInvalidateIPAllowlists = () => queryClient.invalidateQueries({ queryKey: CLIENTS_IP_ALLOWLIST_LIST });
  const handleInvalidateIPAllowlist = (item: string) => queryClient.invalidateQueries({ queryKey: CLIENTS_IP_ALLOWLIST(item) });
  const handleInvalidateRateLimits = () => queryClient.invalidateQueries({ queryKey: CLIENTS_RATE_LIMIT_LIST });
  const handleInvalidateRateLimit = (item: string) => queryClient.invalidateQueries({ queryKey: CLIENTS_RATE_LIMIT(item) });

  clientsEventEmitter.on(clientsEvents.INVALIDATE_APPS, handleInvalidateApps as EventCb);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_APP, handleInvalidateApp as EventCb);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_API_KEYS, handleInvalidateAPIKeys as EventCb);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_API_KEY, handleInvalidateAPIKey as EventCb);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_CLIENT_CREDENTIALS, handleInvalidateClientCredentials as EventCb);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_CLIENT_CREDENTIAL, handleInvalidateClientCredential as EventCb);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_CLIENT_SCOPES, handleInvalidateClientScopes as EventCb);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_CLIENT_SCOPE, handleInvalidateClientScope as EventCb);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_IP_ALLOWLISTS, handleInvalidateIPAllowlists as EventCb);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_IP_ALLOWLIST, handleInvalidateIPAllowlist as EventCb);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_RATE_LIMITS, handleInvalidateRateLimits as EventCb);
  clientsEventEmitter.on(clientsEvents.INVALIDATE_RATE_LIMIT, handleInvalidateRateLimit as EventCb);

  return () => {
    clientsEventEmitter.off(clientsEvents.INVALIDATE_APPS, handleInvalidateApps as EventCb);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_APP, handleInvalidateApp as EventCb);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_API_KEYS, handleInvalidateAPIKeys as EventCb);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_API_KEY, handleInvalidateAPIKey as EventCb);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_CLIENT_CREDENTIALS, handleInvalidateClientCredentials as EventCb);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_CLIENT_CREDENTIAL, handleInvalidateClientCredential as EventCb);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_CLIENT_SCOPES, handleInvalidateClientScopes as EventCb);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_CLIENT_SCOPE, handleInvalidateClientScope as EventCb);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_IP_ALLOWLISTS, handleInvalidateIPAllowlists as EventCb);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_IP_ALLOWLIST, handleInvalidateIPAllowlist as EventCb);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_RATE_LIMITS, handleInvalidateRateLimits as EventCb);
    clientsEventEmitter.off(clientsEvents.INVALIDATE_RATE_LIMIT, handleInvalidateRateLimit as EventCb);
  };
};
