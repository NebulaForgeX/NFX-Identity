import { useEffect } from "react";
import { useQueryClient } from "@tanstack/react-query";

import { cacheEventEmitter, cacheEvents } from "@/events/cache";

/**
 * Hook for handling cache invalidation via event emitter
 * 监听缓存失效事件并自动刷新对应的查询
 */
export const useCacheInvalidation = () => {
  const queryClient = useQueryClient();

  useEffect(() => {
    // Category 相关
    const handleInvalidateCategories = () => {
      queryClient.invalidateQueries({ queryKey: ["categories"] });
      queryClient.invalidateQueries({ queryKey: ["stats"] });
    };

    const handleInvalidateCategory = (id: string) => {
      queryClient.invalidateQueries({ queryKey: ["category", id] });
    };

    // Subcategory 相关
    const handleInvalidateSubcategories = () => {
      queryClient.invalidateQueries({ queryKey: ["subcategories"] });
      queryClient.invalidateQueries({ queryKey: ["stats"] });
    };

    const handleInvalidateSubcategory = (id: string) => {
      queryClient.invalidateQueries({ queryKey: ["subcategory", id] });
    };

    // 注册监听器
    cacheEventEmitter.on(cacheEvents.INVALIDATE_CATEGORIES, handleInvalidateCategories);
    cacheEventEmitter.on(cacheEvents.INVALIDATE_CATEGORY, handleInvalidateCategory);
    cacheEventEmitter.on(cacheEvents.INVALIDATE_SUBCATEGORIES, handleInvalidateSubcategories);
    cacheEventEmitter.on(cacheEvents.INVALIDATE_SUBCATEGORY, handleInvalidateSubcategory);

    // 清理函数
    return () => {
      cacheEventEmitter.off(cacheEvents.INVALIDATE_CATEGORIES, handleInvalidateCategories);
      cacheEventEmitter.off(cacheEvents.INVALIDATE_CATEGORY, handleInvalidateCategory);
      cacheEventEmitter.off(cacheEvents.INVALIDATE_SUBCATEGORIES, handleInvalidateSubcategories);
      cacheEventEmitter.off(cacheEvents.INVALIDATE_SUBCATEGORY, handleInvalidateSubcategory);
    };
  }, [queryClient]);
};
