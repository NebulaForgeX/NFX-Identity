// System Request Types - 基于 NFX-ID Backend

// ========== 系统状态相关 ==========

export interface InitializeSystemStateRequest {
  version?: string;
  metadata?: Record<string, unknown>;
}

export interface ResetSystemStateRequest {
  metadata?: Record<string, unknown>;
}
