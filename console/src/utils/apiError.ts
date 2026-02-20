/**
 * 对齐 Rex-Backend 错误响应格式的解析工具
 * 后端错误体：{ status: number, err_code?: string, message: string, details?: unknown, trace_id?: string }
 * axios-case-converter 会将 response.data 转为 camelCase：errCode, traceId
 */

import type { AxiosError } from "axios";

/** Rex 风格 API 错误体（前端 camelCase） */
export interface ApiErrorBody {
  status?: number;
  errCode?: string;
  message?: string;
  details?: unknown;
  traceId?: string;
}

/**
 * 从任意 caught 错误中解析 Rex 格式的 API 错误体
 */
export function getApiError(error: unknown): ApiErrorBody | null {
  if (!error || typeof error !== "object") return null;
  const ax = error as AxiosError<ApiErrorBody>;
  if (ax.response?.data && typeof ax.response.data === "object") {
    const d = ax.response.data;
    return {
      status: typeof d.status === "number" ? d.status : ax.response.status,
      errCode: d.errCode,
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
export function getApiErrorMessage(error: unknown, fallback = "请求失败，请稍后重试。"): string {
  const api = getApiError(error);
  if (api?.message && api.message.trim()) return api.message.trim();
  if (error instanceof Error && error.message) return error.message;
  return fallback;
}
