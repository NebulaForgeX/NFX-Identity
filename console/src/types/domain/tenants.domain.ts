// Tenants Domain Types - 基于 NFX-ID Backend

import {
  TenantStatus,
  MemberStatus,
  MemberSource,
  TenantAppStatus,
  VerificationStatus,
  VerificationMethod,
  InvitationStatus,
  GroupType,
} from "./enums";

export interface Tenant {
  id: string;
  tenantId: string;
  name: string;
  displayName?: string;
  status: TenantStatus;
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
  type: GroupType;
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
  status: MemberStatus;
  source?: MemberSource;
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
  status: InvitationStatus;
  expiresAt?: string;
  acceptedAt?: string;
  revokedAt?: string;
  createdAt: string;
}

export interface TenantApp {
  id: string;
  tenantId: string;
  appId: string;
  status: TenantAppStatus;
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
  verificationMethod?: VerificationMethod;
  verificationToken: string;
  status: VerificationStatus;
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
