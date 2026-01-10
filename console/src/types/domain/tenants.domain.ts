// Tenants Domain Types - 基于 NFX-ID Backend

export interface Tenant {
  id: string;
  tenantId: string;
  name: string;
  displayName?: string;
  status: string;
  primaryDomain?: string;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
  metadata?: Record<string, unknown>;
}

export interface Group {
  id: string;
  groupId: string;
  tenantId: string;
  name: string;
  type: string;
  parentGroupId?: string;
  description?: string;
  createdBy?: string;
  metadata?: Record<string, unknown>;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface Member {
  id: string;
  tenantId: string;
  userId: string;
  role?: string;
  status: string;
  joinedAt: string;
  createdAt: string;
  updatedAt: string;
}

export interface Invitation {
  id: string;
  inviteId: string;
  tenantId: string;
  email: string;
  role?: string;
  status: string;
  expiresAt?: string;
  acceptedAt?: string;
  revokedAt?: string;
  createdAt: string;
}

export interface TenantApp {
  id: string;
  tenantId: string;
  appId: string;
  status: string;
  createdAt: string;
  updatedAt: string;
}

export interface TenantSetting {
  id: string;
  tenantId: string;
  key: string;
  value: string;
  createdAt: string;
  updatedAt: string;
}

export interface DomainVerification {
  id: string;
  tenantId: string;
  domain: string;
  verificationToken: string;
  status: string;
  verifiedAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface MemberRole {
  id: string;
  memberId: string;
  roleId: string;
  status: string;
  grantedAt: string;
  revokedAt?: string;
  createdAt: string;
}

export interface MemberGroup {
  id: string;
  memberId: string;
  groupId: string;
  status: string;
  grantedAt: string;
  revokedAt?: string;
  createdAt: string;
}

export interface MemberAppRole {
  id: string;
  memberId: string;
  appId: string;
  roleId: string;
  status: string;
  grantedAt: string;
  revokedAt?: string;
  createdAt: string;
}
