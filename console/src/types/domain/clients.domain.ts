// Clients Domain Types - 基于 NFX-ID Backend

import {
  AppType,
  AppStatus,
  Environment,
  ApiKeyStatus,
  ClientCredentialStatus,
  AllowlistStatus,
} from "./enums";

export interface App {
  id: string;
  appId: string;
  tenantId: string;
  name: string;
  description?: string;
  type: AppType;
  status: AppStatus;
  environment: Environment;
  createdAt: string;
  updatedAt: string;
  createdBy?: string;
  updatedBy?: string;
  metadata?: Record<string, unknown>;
  deletedAt?: string;
}

export interface APIKey {
  id: string;
  keyId: string;
  appId: string;
  name?: string;
  status: ApiKeyStatus;
  expiresAt?: string;
  lastUsedAt?: string;
  createdAt: string;
  revokedAt?: string;
}

export interface ClientCredential {
  id: string;
  clientId: string;
  clientSecret: string;
  appId: string;
  status: ClientCredentialStatus;
  createdAt: string;
  revokedAt?: string;
}

export interface ClientScope {
  id: string;
  appId: string;
  scope: string;
  status: string;
  createdAt: string;
}

export interface IPAllowlist {
  id: string;
  ruleId: string;
  appId: string;
  cidr: string;
  description?: string;
  status: AllowlistStatus;
  createdAt: string;
  createdBy?: string;
  updatedAt: string;
  updatedBy?: string;
  revokedAt?: string;
  revokedBy?: string;
  revokeReason?: string;
}

export interface RateLimit {
  id: string;
  appId: string;
  resource: string;
  limit: number;
  window: string;
  description?: string;
  status: string;
  createdAt: string;
  updatedAt: string;
}
