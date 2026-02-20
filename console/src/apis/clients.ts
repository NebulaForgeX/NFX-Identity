// api/clients.ts
import type { AxiosRequestTransformer, InternalAxiosRequestConfig } from "axios";

import axios, { AxiosError } from "axios";
import applyCaseMiddleware from "axios-case-converter";

import { API_ENDPOINTS } from "@/apis/ip";
import AuthStore from "@/stores/authStore";
import { onceAsync } from "@/utils/promise";

// 让 config._retry 有类型
declare module "axios" {
  export interface AxiosRequestConfig {
    _retry?: boolean;
  }
}

// 1) 先创建实例并套 case 中间件
export const protectedClient = applyCaseMiddleware(
  axios.create({
    baseURL: API_ENDPOINTS.PURE,
    timeout: 8000,
  }),
);

export const publicClient = applyCaseMiddleware(
  axios.create({
    baseURL: API_ENDPOINTS.PURE,
    timeout: 8000,
  }),
);

// 2) 请求拦截器：加 token（这里看到的是转换前的 camelCase）
protectedClient.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const accessToken = AuthStore.getState().accessToken;
    if (accessToken) config.headers.Authorization = `Bearer ${accessToken}`;
    // 这里的 data/params 还是 camelCase（转换尚未发生）
    // console.log("🧩 Before transform (camelCase) - data:", config.data);
    // console.log("🧩 Before transform (camelCase) - params:", config.params);
    return config;
  },
  (error) => Promise.reject(error),
);

// 3) 在 transformRequest 队列“末尾”追加一个调试 transformer
//    这里的 data 一定已经被 axios-case-converter 转成 snake_case 了
function asArray<T>(v: T | T[] | undefined): T[] {
  return v ? (Array.isArray(v) ? v : [v]) : [];
}

protectedClient.defaults.transformRequest = [
  ...asArray<AxiosRequestTransformer>(protectedClient.defaults.transformRequest),
  (data: unknown, _headers) => {
    let out: unknown = data;
    try {
      if (typeof out === "string") out = JSON.parse(out) as unknown;
    } catch {
      // 忽略解析错误，继续处理
    }
    // console.log("🐍 After transformRequest (snake_case) - data:", out);
    return data; // 不要改动 data
  },
];

// Rex 错误体类型（与后端 httpx.HTTPResp 错误时一致，snake_case 转 camelCase）
type RexErrorData = {
  message?: string;
  errCode?: string;
  status?: number;
  details?: unknown;
  traceId?: string;
};

/** 按 Rex 规范统一错误：把 response.data.message 写入 error.message，便于 UI 直接展示 */
function normalizeRexApiError(error: AxiosError<RexErrorData>): void {
  const errorData = error.response?.data;
  const msg = errorData?.message;
  if (msg && typeof msg === "string") {
    try {
      Object.defineProperty(error, "message", { value: msg, configurable: true });
    } catch {
      (error as unknown as Record<string, string>).message = msg;
    }
  }
}

/** 响应错误时打日志（Rex 字段：message, errCode, status） */
function logRexApiError(error: AxiosError<RexErrorData>): void {
  const errorData = error.response?.data as RexErrorData | undefined;
  const msg = errorData?.message;
  if (msg) {
    console.log("❌ API Error:", {
      message: msg,
      errCode: errorData?.errCode,
      status: error.response?.status ?? errorData?.status,
      url: error.config?.url,
      method: error.config?.method,
    });
  } else if (import.meta.env.DEV && error.response?.status) {
    console.log("❌ HTTP Error:", {
      status: error.response.status,
      url: error.config?.url,
      method: error.config?.method,
    });
  }
}

// 4) 响应拦截器：按 Rex-Backend 错误体统一处理（status, err_code, message, details, trace_id）
protectedClient.interceptors.response.use(
  (response) => response,
  async (error: unknown) => {
    if (!(error instanceof AxiosError)) {
      return Promise.reject(error);
    }
    normalizeRexApiError(error);
    logRexApiError(error);

    if (error.response?.status === 401 && error.config && !error.config._retry) {
      error.config._retry = true;
      
      try {
        // 尝试刷新token
        await refreshTokens();
        
        // 刷新成功，更新请求头并重试原请求
        const newAccessToken = AuthStore.getState().accessToken;
        if (newAccessToken && error.config.headers) {
          error.config.headers.Authorization = `Bearer ${newAccessToken}`;
        }
        
        // 重试原请求
        return protectedClient.request(error.config);
      } catch (refreshError) {
        // 刷新失败，清除认证信息并跳转到登录页
        AuthStore.getState().clearAuth();
        // 触发路由跳转到登录页（如果需要）
        if (window.location.pathname !== "/login") {
          window.location.href = "/login";
        }
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  },
);

// 4b) publicClient 同样按 Rex 错误体统一 error.message（登录/注册/验证码等）
publicClient.interceptors.response.use(
  (response) => response,
  (error: unknown) => {
    if (error instanceof AxiosError) {
      normalizeRexApiError(error);
      logRexApiError(error);
    }
    return Promise.reject(error);
  },
);

// 5) 刷新 token（防重入）- 企业级最佳实践实现
export const refreshTokens = onceAsync(async () => {
  try {
    const { refreshToken } = AuthStore.getState();
    if (!refreshToken) {
      throw new Error("Refresh token not found");
    }

    // 导入登录API（避免循环依赖）
    const { RefreshAccessToken } = await import("./auth.api");
    
    // 调用刷新token API
    const response = await RefreshAccessToken({ refreshToken });
    
    // 更新tokens
    AuthStore.getState().setTokens({
      accessToken: response.accessToken,
      refreshToken: response.refreshToken,
    });
    
    // 设置认证状态
    AuthStore.getState().setIsAuthValid(true);
    
    return response;
  } catch (error) {
    // 刷新失败，清除所有认证信息
    AuthStore.getState().clearAuth();
    throw error;
  }
});
