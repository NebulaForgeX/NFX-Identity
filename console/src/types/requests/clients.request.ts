// Clients Request Types - 基于 NFX-ID Backend

// ========== 应用相关 ==========

export interface CreateAppRequest {
  appId: string;
  tenantId: string;
  name: string;
  description?: string;
  type: string;
  status?: string;
  environment: string;
  createdBy?: string;
  metadata?: Record<string, unknown>;
}

export interface UpdateAppRequest {
  name: string;
  description?: string;
  type: string;
  environment: string;
  updatedBy?: string;
  metadata?: Record<string, unknown>;
}

// ========== API Key 相关 ==========

export interface CreateAPIKeyRequest {
  keyId: string;
  appId: string;
  keyHash: string;
  hashAlg: string;
  name: string;
  expiresAt?: string;
  createdBy?: string;
  metadata?: Record<string, unknown>;
}

// ========== Client Credential 相关 ==========

export interface CreateClientCredentialRequest {
  appId: string;
  clientId: string;
  secretHash: string;
  hashAlg: string;
  expiresAt?: string;
  createdBy?: string;
}

// ========== Client Scope 相关 ==========

export interface CreateClientScopeRequest {
  appId: string;
  scope: string;
  grantedBy?: string;
  expiresAt?: string;
}

// ========== IP Allowlist 相关 ==========

export interface CreateIPAllowlistRequest {
  ruleId: string;
  appId: string;
  cidr: string;
  description?: string;
  status?: string;
  createdBy?: string;
}

// ========== Rate Limit 相关 ==========

export interface CreateRateLimitRequest {
  appId: string;
  limitType: string;
  limitValue: number;
  windowSeconds: number;
  description?: string;
  status?: string;
  createdBy?: string;
}
