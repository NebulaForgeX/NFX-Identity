import type { QueryKey } from "@tanstack/react-query";
import { DOMAIN_SYSTEM, DOMAIN_SYSTEM_STATE } from "./domain.key";
import { CACHE_ITEM, CACHE_LIST } from "./cache.key";

/**
 * System 相关的 Query Key 常量
 */

// ========== SystemState 相关 ==========
export const SYSTEM_SYSTEM_STATE_LIST = [DOMAIN_SYSTEM, CACHE_LIST, DOMAIN_SYSTEM_STATE];
export const SYSTEM_SYSTEM_STATE = (item: string): QueryKey => [DOMAIN_SYSTEM, CACHE_ITEM, DOMAIN_SYSTEM_STATE, item];

/**
 * 用于 useCacheInvalidationEvents.ts 的 query key 前缀
 */
export const SYSTEM_QUERY_KEY_PREFIXES = {
  SYSTEM_STATES: [DOMAIN_SYSTEM, CACHE_LIST, DOMAIN_SYSTEM_STATE],
  SYSTEM_STATE: [DOMAIN_SYSTEM, CACHE_ITEM, DOMAIN_SYSTEM_STATE],
};
