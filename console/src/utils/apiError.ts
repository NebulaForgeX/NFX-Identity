/**
 * 与约定一致：HTTP 错误一律经 fiberx.Error/ErrorFromErrx → httpx.BuildErrorResp，
 * 响应体固定为 { status, err_code, message, details?, trace_id? }（axios 转 camelCase）。
 * 前端只认「有 response.data」= 后端错误体，直接取 errCode/message；否则用 fallback。
 */

import type { AxiosError } from "axios";

import i18n from "@/assets/languages/i18n";
import type { ApiErrCode, ApiErrorBody } from "@/types/apiError";

export type { ApiErrorBody, ApiErrCode };

/** 仅当为后端错误（有 response.data）时解析，否则返回 null。不对外部 error 做多余判断。 */
export function getApiError(error: unknown): ApiErrorBody | null {
  const d = (error as AxiosError<ApiErrorBody>)?.response?.data;
  if (!d || typeof d !== "object") return null;
  return {
    status: d.status,
    errCode: d.errCode as ApiErrCode | undefined,
    message: d.message,
    details: d.details,
    traceId: d.traceId,
  };
}

/**
 * UI 展示用文案。fallback 建议 [函数名] error，便于发现未返回 err_code/message 的接口。
 */
export function getApiErrorMessage(error: unknown, fallback: string): string {
  const api = getApiError(error);
  if (api?.errCode) {
    const out = i18n.t(`errors:${api.errCode}`);
    if (out && out !== `errors:${api.errCode}`) return out;
  }
  if (api?.message) return api.message;
  return fallback;
}
