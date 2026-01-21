import type {
  QueryKey,
  UseQueryResult,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";
import type { AxiosError } from "axios";

import { useMemo } from "react";
import { useQuery, useSuspenseQuery } from "@tanstack/react-query";
import type { UnifiedQueryMode, suspenseUnifiedQueryOptions, SuspenseUnifiedQueryOptions } from "./type";


// ------------------ Overload Signatures ------------------
export function makeUnifiedQuery<T, F extends object = Record<string, unknown>>(
  fetchRemote: (params: F) => Promise<T>,
  mode: "suspense",
  postProcess?: (data: T) => void,
): (
  queryKey: QueryKey,
  filter?: F,
  options?: SuspenseUnifiedQueryOptions<T>,
) => UseSuspenseQueryResult<T, AxiosError>;

export function makeUnifiedQuery<T, F extends object = Record<string, unknown>>(
  fetchRemote: (params: F) => Promise<T>,
  mode: "normal",
  postProcess?: (data: T) => void,
): (
  queryKey: QueryKey,
  filter?: F,
  options?: suspenseUnifiedQueryOptions<T>,
) => UseQueryResult<T, AxiosError>;

export function makeUnifiedQuery<T, F extends object = Record<string, unknown>>(
  fetchRemote: (params: F) => Promise<T>,
  mode?: "suspense",
  postProcess?: (data: T) => void,
): (
  queryKey: QueryKey,
  filter?: F,
  options?: suspenseUnifiedQueryOptions<T>,
) => UseQueryResult<T, AxiosError>;

// ------------------ Implementation ------------------
export function makeUnifiedQuery<T, F extends object = Record<string, unknown>>(
  fetchRemote: (params: F) => Promise<T>,
  mode: UnifiedQueryMode = "suspense",
  postProcess?: (data: T) => void,
) {

  const fetchFunction = async (filter: F): Promise<T> => {
    const data = await fetchRemote(filter);
    postProcess?.(data);
    return data;
  };

  const buildCommonOptions = (queryKey: QueryKey, filter?: F, options?: suspenseUnifiedQueryOptions<T> | SuspenseUnifiedQueryOptions<T>) => {
    return {
      queryKey: filter !== undefined ? [...queryKey, filter] : queryKey,
      queryFn: () => fetchFunction(filter || ({} as F)),
      select: (data: T) => data,
      retry: (failureCount: number, error: AxiosError) => {
        const status = error?.status ?? error?.response?.status;
        const transient = (typeof status === "number" && status >= 500) || error?.code === "NETWORK_ERROR";
        const retryMax = typeof options?.retry === "number" ? options.retry : 3;
        return transient && failureCount < retryMax;
      },
      ...(options as object),
    };
  };

  // suspense Query
  function useQuerysuspense(queryKey: QueryKey, filter?: F, options?: suspenseUnifiedQueryOptions<T>): UseQueryResult<T, AxiosError> {
    const common = useMemo(() => buildCommonOptions(queryKey, filter, options), [queryKey, filter, options]);
    return useQuery(common);
  }

  // Suspense Query
  function useQuerySuspense(
    queryKey: QueryKey,
    filter?: F,
    options?: SuspenseUnifiedQueryOptions<T>,
  ): UseSuspenseQueryResult<T, AxiosError> {
    const common = useMemo(() => buildCommonOptions(queryKey, filter, options), [queryKey, filter, options]);
    return useSuspenseQuery(common);
  }

  if (mode === "suspense") {
    return (queryKey: QueryKey, filter?: F, options?: SuspenseUnifiedQueryOptions<T>) =>
      useQuerySuspense(queryKey, filter, options);
  }
  return (queryKey: QueryKey, filter?: F, options?: suspenseUnifiedQueryOptions<T>) => useQuerysuspense(queryKey, filter, options);
}
