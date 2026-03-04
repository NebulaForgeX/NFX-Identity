import type { QueryClient } from "@tanstack/react-query";
import { systemEventEmitter, systemEvents } from "@/events/system";
import { SYSTEM_SYSTEM_STATE, SYSTEM_SYSTEM_STATE_LIST } from "@/constants";

type EventCb = (...args: unknown[]) => void;

/**
 * System 相关的缓存失效事件处理
 */
export const useSystemCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateSystemStates = () => queryClient.invalidateQueries({ queryKey: SYSTEM_SYSTEM_STATE_LIST });
  const handleInvalidateSystemState = (item: string) => queryClient.invalidateQueries({ queryKey: SYSTEM_SYSTEM_STATE(item) });

  systemEventEmitter.on(systemEvents.INVALIDATE_SYSTEM_STATES, handleInvalidateSystemStates as EventCb);
  systemEventEmitter.on(systemEvents.INVALIDATE_SYSTEM_STATE, handleInvalidateSystemState as EventCb);

  return () => {
    systemEventEmitter.off(systemEvents.INVALIDATE_SYSTEM_STATES, handleInvalidateSystemStates as EventCb);
    systemEventEmitter.off(systemEvents.INVALIDATE_SYSTEM_STATE, handleInvalidateSystemState as EventCb);
  };
};
