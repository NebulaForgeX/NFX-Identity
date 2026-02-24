/**
 * 对齐 NFX-Identity 后端 errx + fiberx 错误响应格式的解析工具
 * 后端错误体：{ status, err_code, message, details?, trace_id? }
 * axios-case-converter 将 response.data 转为 camelCase：errCode, traceId
 */

import type { AxiosError } from "axios";

import type { ApiErrorBody, ApiErrCode } from "@/types/apiError";

export type { ApiErrorBody, ApiErrCode };

/**
 * 从任意 caught 错误中解析 API 错误体
 */
export function getApiError(error: unknown): ApiErrorBody | null {
  if (!error || typeof error !== "object") return null;
  const ax = error as AxiosError<ApiErrorBody>;
  if (ax.response?.data && typeof ax.response.data === "object") {
    const d = ax.response.data;
    return {
      status: typeof d.status === "number" ? d.status : ax.response.status,
      errCode: d.errCode as ApiErrCode | undefined,
      message: typeof d.message === "string" ? d.message : undefined,
      details: d.details,
      traceId: d.traceId,
    };
  }
  return null;
}

/**
 * 获取用于 UI 展示的错误文案：优先使用 API 返回的 message，否则回退到 error.message 或默认文案
 */
export function getApiErrorMessage(error: unknown, fallback: string): string {
  const api = getApiError(error);
  if (api?.message && api.message.trim()) return api.message.trim();
  if (error instanceof Error && error.message) return error.message;
  return fallback;
}
