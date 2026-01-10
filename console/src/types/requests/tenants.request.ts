// Tenants Request Types - 基于 NFX-ID Backend

// ========== 租户相关 ==========

export interface CreateTenantRequest {
  tenantId: string;
  name: string;
  displayName?: string;
  status?: string;
  primaryDomain?: string;
  metadata?: Record<string, unknown>;
}

export interface UpdateTenantRequest {
  name: string;
  displayName?: string;
  primaryDomain?: string;
  metadata?: Record<string, unknown>;
}

export interface UpdateTenantStatusRequest {
  status: string;
}

// ========== 组相关 ==========

export interface CreateGroupRequest {
  groupId: string;
  tenantId: string;
  name: string;
  type: string;
  parentGroupId?: string;
  description?: string;
  createdBy?: string;
  metadata?: Record<string, unknown>;
}

export interface UpdateGroupRequest {
  name: string;
  type: string;
  parentGroupId?: string;
  description?: string;
  metadata?: Record<string, unknown>;
}

// ========== 成员相关 ==========

export interface CreateMemberRequest {
  tenantId: string;
  userId: string;
  status?: string;
  source?: string;
  createdBy?: string;
  externalRef?: string;
  metadata?: Record<string, unknown>;
}

export interface UpdateMemberStatusRequest {
  status: string;
}

// ========== 邀请相关 ==========

export interface CreateInvitationRequest {
  inviteId: string;
  tenantId: string;
  email: string;
  tokenHash: string;
  expiresAt: string;
  status?: string;
  invitedBy: string;
  roleIds?: string[];
  metadata?: Record<string, unknown>;
}

export interface AcceptInvitationRequest {
  userId: string;
}

export interface RevokeInvitationRequest {
  revokedBy: string;
  revokeReason?: string;
}

// ========== 租户应用相关 ==========

export interface CreateTenantAppRequest {
  tenantId: string;
  appId: string;
  status?: string;
  createdBy?: string;
  settings?: Record<string, unknown>;
}

export interface UpdateTenantAppRequest {
  status?: string;
  settings?: Record<string, unknown>;
}

// ========== 租户设置相关 ==========

export interface CreateTenantSettingRequest {
  tenantId: string;
  enforceMfa: boolean;
  allowedEmailDomains?: string[];
  sessionTtlMinutes?: number;
  passwordPolicy?: Record<string, unknown>;
  loginPolicy?: Record<string, unknown>;
  mfaPolicy?: Record<string, unknown>;
}

export interface UpdateTenantSettingRequest {
  enforceMfa?: boolean;
  allowedEmailDomains?: string[];
  sessionTtlMinutes?: number;
  passwordPolicy?: Record<string, unknown>;
  loginPolicy?: Record<string, unknown>;
  mfaPolicy?: Record<string, unknown>;
  updatedBy?: string;
}

// ========== 域名验证相关 ==========

export interface CreateDomainVerificationRequest {
  tenantId: string;
  domain: string;
  verificationMethod: string;
  verificationToken?: string;
  status?: string;
  expiresAt?: string;
  createdBy?: string;
  metadata?: Record<string, unknown>;
}

// ========== 成员角色相关 ==========

export interface CreateMemberRoleRequest {
  tenantId: string;
  memberId: string;
  roleId: string;
  assignedBy?: string;
  expiresAt?: string;
  scope?: string;
}

export interface RevokeMemberRoleRequest {
  revokedBy: string;
  revokeReason: string;
}

// ========== 成员组相关 ==========

export interface CreateMemberGroupRequest {
  memberId: string;
  groupId: string;
  assignedBy?: string;
}

export interface RevokeMemberGroupRequest {
  revokedBy: string;
}

// ========== 成员应用角色相关 ==========

export interface CreateMemberAppRoleRequest {
  memberId: string;
  appId: string;
  roleId: string;
  assignedBy?: string;
  expiresAt?: string;
}

export interface RevokeMemberAppRoleRequest {
  revokedBy: string;
  revokeReason: string;
}
