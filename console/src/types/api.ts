// API 响应类型定义 - 基于 NFX-ID Backend

export interface BaseResponse {
  status: string; // "success" | "error"
  code: number;
  message: string;
  traceId?: string; // trace_id (axios-case-converter 会自动转换)
}

export interface DataResponse<T> extends BaseResponse {
  data: T;
  meta?: Record<string, unknown>; // 用于存储额外的元数据，如 total 等
}

export interface ListDTO<T> {
  items: T[];
  total: number;
}

export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  page: number;
  pageSize: number;
}

