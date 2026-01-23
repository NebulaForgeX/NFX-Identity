// System Request Types - 基于 NFX-ID Backend

// ========== 系统状态相关 ==========

export interface InitializeSystemStateRequest {
  version?: string;
  admin_username: string;
  admin_password: string;
  admin_email?: string;
  admin_phone?: string;
  admin_country_code?: string;
  metadata?: Record<string, unknown>;
}

export interface ResetSystemStateRequest {
  metadata?: Record<string, unknown>;
}
