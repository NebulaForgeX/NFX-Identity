import { useEffect } from "react";
import type { QueryClient } from "@tanstack/react-query";

import { useAccessCacheInvalidation } from "./useAccessCacheInvalidation";
import { useAuthCacheInvalidation } from "./useAuthCacheInvalidation";
import { useAuditCacheInvalidation } from "./useAuditCacheInvalidation";
import { useClientsCacheInvalidation } from "./useClientsCacheInvalidation";
import { useDirectoryCacheInvalidation } from "./useDirectoryCacheInvalidation";
import { useImageCacheInvalidation } from "./useImageCacheInvalidation";
import { useLogoutCleanup } from "./useLogoutCleanup";
import { useSystemCacheInvalidation } from "./useSystemCacheInvalidation";
import { useTenantsCacheInvalidation } from "./useTenantsCacheInvalidation";

/**
 * Hook for handling all cache invalidation events via event emitters
 * 监听所有缓存失效事件并自动刷新对应的查询
 * @param queryClient - QueryClient 实例
 */
export const useCacheInvalidationEvents = (queryClient: QueryClient) => {
  useEffect(() => {
    const cleanupAccess = useAccessCacheInvalidation(queryClient);
    const cleanupAuth = useAuthCacheInvalidation(queryClient);
    const cleanupAudit = useAuditCacheInvalidation(queryClient);
    const cleanupClients = useClientsCacheInvalidation(queryClient);
    const cleanupDirectory = useDirectoryCacheInvalidation(queryClient);
    const cleanupImage = useImageCacheInvalidation(queryClient);
    const cleanupSystem = useSystemCacheInvalidation(queryClient);
    const cleanupTenants = useTenantsCacheInvalidation(queryClient);

    return () => {
      cleanupAccess();
      cleanupAuth();
      cleanupAudit();
      cleanupClients();
      cleanupDirectory();
      cleanupImage();
      cleanupSystem();
      cleanupTenants();
    };
  }, [queryClient]);
};
