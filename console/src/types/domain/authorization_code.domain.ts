// Authorization Code Domain Types - 基于 NFX-ID Permission Service

export interface AuthorizationCode {
  id: string;
  code: string;
  maxUses: number;
  usedCount: number;
  createdBy?: string;
  expiresAt?: string;
  isActive: boolean;
  isAvailable: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface CreateAuthorizationCodeParams {
  code: string;
  maxUses: number;
  createdBy?: string;
  expiresAt?: string;
  isActive?: boolean;
}

export interface UseAuthorizationCodeParams {
  code: string;
}
