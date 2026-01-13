import type { QueryClient } from "@tanstack/react-query";
import { systemEventEmitter, systemEvents } from "@/events/system";
import { SYSTEM_QUERY_KEY_PREFIXES } from "@/constants";

/**
 * System 相关的缓存失效事件处理
 */
export const useSystemCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateSystemStates = () => queryClient.invalidateQueries({ queryKey: SYSTEM_QUERY_KEY_PREFIXES.SYSTEM_STATES });
  const handleInvalidateSystemState = (item: string) => queryClient.invalidateQueries({ queryKey: [...SYSTEM_QUERY_KEY_PREFIXES.SYSTEM_STATE, item] });

  // 注册监听器
  systemEventEmitter.on(systemEvents.INVALIDATE_SYSTEM_STATES, handleInvalidateSystemStates);
  systemEventEmitter.on(systemEvents.INVALIDATE_SYSTEM_STATE, handleInvalidateSystemState);

  // 清理监听器
  return () => {
    systemEventEmitter.off(systemEvents.INVALIDATE_SYSTEM_STATES, handleInvalidateSystemStates);
    systemEventEmitter.off(systemEvents.INVALIDATE_SYSTEM_STATE, handleInvalidateSystemState);
  };
};
