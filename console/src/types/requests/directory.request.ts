// Directory Request Types - 基于 NFX-ID Backend

// ========== 用户相关 ==========

export interface CreateUserRequest {
  tenantId: string;
  username: string;
  status?: string;
  isVerified?: boolean;
}

export interface UpdateUserStatusRequest {
  status: string;
}

export interface UpdateUserUsernameRequest {
  username: string;
}

// ========== 徽章相关 ==========

export interface CreateBadgeRequest {
  name: string;
  description?: string;
  iconUrl?: string;
  color?: string;
  category?: string;
  isSystem?: boolean;
}

export interface UpdateBadgeRequest {
  name: string;
  description?: string;
  iconUrl?: string;
  color?: string;
  category?: string;
}

// ========== 用户徽章相关 ==========

export interface CreateUserBadgeRequest {
  userId: string;
  badgeId: string;
  awardedAt: string;
}

// ========== 用户教育相关 ==========

export interface CreateUserEducationRequest {
  userId: string;
  school: string;
  degree?: string;
  field?: string;
  startDate?: string;
  endDate?: string;
  description?: string;
}

export interface UpdateUserEducationRequest {
  school?: string;
  degree?: string;
  field?: string;
  startDate?: string;
  endDate?: string;
  description?: string;
}

// ========== 用户邮箱相关 ==========

export interface CreateUserEmailRequest {
  userId: string;
  email: string;
  isPrimary?: boolean;
  isVerified?: boolean;
  verificationToken?: string;
}

export interface UpdateUserEmailRequest {
  email?: string;
}

// ========== 用户职业相关 ==========

export interface CreateUserOccupationRequest {
  userId: string;
  company: string;
  position?: string;
  startDate?: string;
  endDate?: string;
  description?: string;
}

export interface UpdateUserOccupationRequest {
  company?: string;
  position?: string;
  startDate?: string;
  endDate?: string;
  description?: string;
}

// ========== 用户电话相关 ==========

export interface CreateUserPhoneRequest {
  userId: string;
  phone: string;
  countryCode?: string;
  isPrimary?: boolean;
  isVerified?: boolean;
  verificationCode?: string;
  verificationExpiresAt?: string;
}

export interface UpdateUserPhoneRequest {
  phone?: string;
}

// ========== 用户偏好相关 ==========

export interface CreateUserPreferenceRequest {
  userId: string;
  key: string;
  value: string;
}

export interface UpdateUserPreferenceRequest {
  value: string;
}

// ========== 用户资料相关 ==========

export interface CreateUserProfileRequest {
  userId: string;
  firstName?: string;
  lastName?: string;
  displayName?: string;
  bio?: string;
  avatarId?: string;
}

export interface UpdateUserProfileRequest {
  firstName?: string;
  lastName?: string;
  displayName?: string;
  bio?: string;
  avatarId?: string;
}
