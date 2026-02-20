// API 响应类型定义 - 对齐 Rex-Backend httpx.HTTPResp 规范

/** 成功/错误统一信封：status 为 HTTP 状态码数字，错误时含 err_code */
export interface BaseResponse {
  status: number; // HTTP status code (e.g. 200, 400, 404)
  errCode?: string; // 错误码（仅错误响应时有）
  message: string;
  details?: unknown; // 结构化错误详情
  meta?: Record<string, unknown>;
  traceId?: string; // 请求追踪 ID（通常仅错误时返回）
}

/** 成功响应：含 data 负载 */
export interface DataResponse<T> extends BaseResponse {
  data: T;
}

/** 列表分页负载：与 Rex-Backend httpx.Page 一致，放在 data 中 */
export interface ListDTO<T> {
  items: T[];
  total: number;
}

/** 兼容：列表接口的 data 可能是 ListDTO 或直接数组（按后端实际返回） */
export type ListData<T> = ListDTO<T> | T[];

/** 旧分页结构（如后端未统一为 items/total 时可保留引用） */
export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  page: number;
  pageSize: number;
}
