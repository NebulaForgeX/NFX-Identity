import type { VirtualizerOptions } from "@tanstack/react-virtual";
import type { ReactNode } from "react";

import { memo, useCallback, useEffect } from "react";
import { useTranslation } from "react-i18next";
import { useWindowVirtualizer } from "@tanstack/react-virtual";

import styles from "./styles.module.css";

interface VirtualWindowListProps<T>
  extends Partial<
    Omit<
      VirtualizerOptions<Window, Element>,
      | "count"
      | "getScrollElement"
      | "observeElementRect"
      | "observeElementOffset"
      | "scrollToFn"
      | "estimateSize"
      | "getItemKey"
    >
  > {
  data: T[];
  hasNextPage?: boolean;
  isFetchingNextPage?: boolean;
  fetchNextPage?: () => void;

  renderItem: (item: T, index: number) => ReactNode;
  estimateSize?: number | ((index: number) => number);
  height?: string | number;
  getItemKey: (item: T, index: number) => string | number;

  emptyState?: ReactNode;
  loadingIndicator?: ReactNode;
  endOfListIndicator?: ReactNode;

  flexClass?: string;
  outerClass?: string;
  innerClass?: string;
}

/* -------------------------------------------
 ✅ MAIN COMPONENT
-------------------------------------------- */
function VirtualWindowListComponent<T>({
  data,
  hasNextPage,
  isFetchingNextPage,
  fetchNextPage,
  renderItem,
  estimateSize = 200,
  overscan = 5,
  height,
  getItemKey,

  emptyState,
  loadingIndicator,
  endOfListIndicator,
  flexClass,
  outerClass,
  innerClass,
  ...virtualizerOptions
}: VirtualWindowListProps<T>) {
  /* ✅ 防御性检查：确保 data 是数组 */
  const safeData = Array.isArray(data) ? data : [];
  const dataLength = safeData.length;

  /* ✅ 真正的 window 虚拟滚动器 */
  const virtualizer = useWindowVirtualizer({
    count: hasNextPage ? dataLength + 1 : dataLength,
    estimateSize: typeof estimateSize === "number" ? () => estimateSize : estimateSize,
    overscan,
    ...virtualizerOptions,
  });

  const items = virtualizer.getVirtualItems();

  /* ✅ 自动加载更多 */
  useEffect(() => {
    const last = items[items.length - 1];
    if (!last) return;

    const isLoaderRow = last.index >= dataLength;

    if (isLoaderRow && hasNextPage && !isFetchingNextPage && fetchNextPage) {
      fetchNextPage();
    }
  }, [items, hasNextPage, isFetchingNextPage, fetchNextPage, dataLength]);

  const { t } = useTranslation("components");

  /* ✅ renderEmptyState — 保持原 VirtualList 一致 */
  const renderEmptyState = useCallback(() => {
    return (
      <div
        className={styles.virtualList}
        style={{
          height: typeof height === "number" ? `${height}px` : (height ?? "auto"),
        }}
      >
        {emptyState ?? (
          <div className={styles.emptyContainer}>
            <span>{t("virtualList.noData")}</span>
          </div>
        )}
      </div>
    );
  }, [emptyState, height, t]);

  /* ✅ renderLoadingIndicator */
  const renderLoadingIndicator = useCallback(() => {
    return (
      loadingIndicator ?? (
        <div className={styles.loadingMore}>
          <span>{t("virtualList.loadingMore")}</span>
        </div>
      )
    );
  }, [loadingIndicator, t]);

  /* ✅ renderEndOfListIndicator */
  const renderEndOfListIndicator = useCallback(() => {
    return (
      endOfListIndicator ?? (
        <div className={styles.endOfList}>
          <span>{t("virtualList.noMoreItems")}</span>
        </div>
      )
    );
  }, [endOfListIndicator, t]);

  /* ✅ 无数据情况 */
  if (dataLength === 0) return renderEmptyState();

  /* ✅ 正常渲染 */
  return (
    <div className={`${styles.virtualList} ${outerClass ?? ""}`}>
      <div
        className={`${styles.virtualListInner} ${innerClass ?? ""}`}
        style={{
          height: virtualizer.getTotalSize(),
        }}
      >
        <div
          className={`${styles.virtualListItems} ${flexClass ?? ""}`}
          style={{
            transform: `translateY(${items[0]?.start ?? 0}px)`,
          }}
        >
          {items.map((row) => {
            const isLoaderRow = row.index >= dataLength;
            const item = safeData[row.index];

            return (
              <div key={row.key} data-index={row.index} ref={virtualizer.measureElement}>
                {isLoaderRow ? (
                  hasNextPage ? (
                    renderLoadingIndicator()
                  ) : (
                    renderEndOfListIndicator()
                  )
                ) : (
                  <div key={getItemKey(item, row.index)}>{renderItem(item, row.index)}</div>
                )}
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );
}

const VirtualWindowList = memo(VirtualWindowListComponent) as typeof VirtualWindowListComponent;

export default VirtualWindowList;
