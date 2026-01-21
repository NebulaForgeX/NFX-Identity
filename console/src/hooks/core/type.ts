import type { UseInfiniteQueryOptions, UseQueryOptions, UseSuspenseInfiniteQueryOptions, UseSuspenseQueryOptions } from "@tanstack/react-query";
import type { AxiosError } from "axios";

//! ================ Base Type ================

//* Number Cursor Fetch Result
export type NumberPagePayload<T> = {
  data: T[];
  nextCursor?: number;
};

export type FetchNumberListParams<F extends object = Record<string, unknown>> = F & {
  offset: number;
  limit: number;
};

export type ListNumberCursorFetchResult<T> = {
  items: T[];
  total: number;
};
//*

//* String Cursor Fetch Result
export type StringPagePayload<T> = {
  data: T[];
  nextCursor: string;
};

export type FetchStringListParams<F extends object = Record<string, unknown>> = F & {
  offset: string;
  limit: number;
};

export type ListStringCursorFetchResult<T> = {
  items: T[];
  nextCursor: string;
};
//*

export type MutationCtx = {
  prev?: unknown;
};

export type Id = string;

//! ================ High Order Type ================

export type InfiniteQueryMode = "suspense" | "suspense";
export type InfiniteQueryOptions<T> = Omit<
  UseInfiniteQueryOptions<NumberPagePayload<T>, AxiosError, T[]>,
  "queryKey" | "queryFn" | "getNextPageParam" | "initialPageParam"
>;
export type SuspenseInfiniteQueryOptions<T> = Omit<
  UseSuspenseInfiniteQueryOptions<NumberPagePayload<T>, AxiosError, T[]>,
  "queryKey" | "queryFn" | "getNextPageParam" | "initialPageParam"
>;


export type StringInfiniteQueryOptions<T> = Omit<
  UseInfiniteQueryOptions<StringPagePayload<T>, AxiosError, T[]>,
  "queryKey" | "queryFn" | "getNextPageParam" | "initialPageParam"
>;
export type SuspenseStringInfiniteQueryOptions<T> = Omit<
  UseSuspenseInfiniteQueryOptions<StringPagePayload<T>, AxiosError, T[]>,
  "queryKey" | "queryFn" | "getNextPageParam" | "initialPageParam"
>;


export type InfiniteNumberQueryOptions<T> = InfiniteQueryOptions<T> | SuspenseInfiniteQueryOptions<T>;
export type InfiniteStringQueryOptions<T> = StringInfiniteQueryOptions<T> | SuspenseStringInfiniteQueryOptions<T>;

export type UnifiedQueryMode = "suspense" | "normal";
export type suspenseUnifiedQueryOptions<T> =  Omit<UseQueryOptions<T, AxiosError, T>, "queryKey" | "queryFn">;
export type SuspenseUnifiedQueryOptions<T> =  Omit<UseSuspenseQueryOptions<T, AxiosError, T>, "queryKey" | "queryFn">;

// 固定包含 options 和 postProcess，同时允许添加任意其他属性
export type UnifiedQueryParams<T> = {
  options?: SuspenseUnifiedQueryOptions<T>;
  postProcess?: (data: T) => void;
} & Record<string, any>;