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
  major?: string;
  fieldOfStudy?: string; // 后端期望 field_of_study，前端使用 fieldOfStudy (会自动转换)
  startDate?: string;
  endDate?: string;
  isCurrent?: boolean;
  description?: string;
  grade?: string;
  activities?: string;
  achievements?: string;
}

export interface UpdateUserEducationRequest {
  school?: string;
  degree?: string;
  major?: string;
  fieldOfStudy?: string; // 后端期望 field_of_study，前端使用 fieldOfStudy (会自动转换)
  startDate?: string;
  endDate?: string;
  isCurrent?: boolean;
  description?: string;
  grade?: string;
  activities?: string;
  achievements?: string;
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
  position: string;
  department?: string;
  industry?: string;
  location?: string;
  employmentType?: string;
  startDate?: string;
  endDate?: string;
  isCurrent?: boolean;
  description?: string;
  responsibilities?: string;
  achievements?: string;
  skillsUsed?: string[];
}

export interface UpdateUserOccupationRequest {
  company?: string;
  position?: string;
  department?: string;
  industry?: string;
  location?: string;
  employmentType?: string;
  startDate?: string;
  endDate?: string;
  isCurrent?: boolean;
  description?: string;
  responsibilities?: string;
  achievements?: string;
  skillsUsed?: string[];
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
  theme?: string;
  language?: string;
  timezone?: string;
  dashboardBackground?: string;
  notifications?: Record<string, unknown>;
  privacy?: Record<string, unknown>;
  display?: Record<string, unknown>;
  other?: Record<string, unknown>;
}

export interface UpdateUserPreferenceRequest {
  theme?: string;
  language?: string;
  timezone?: string;
  dashboardBackground?: string;
  notifications?: Record<string, unknown>;
  privacy?: Record<string, unknown>;
  display?: Record<string, unknown>;
  other?: Record<string, unknown>;
}

// ========== 用户资料相关 ==========

export interface CreateUserProfileRequest {
  userId: string;
  role?: string;
  firstName?: string;
  lastName?: string;
  nickname?: string;
  displayName?: string;
  bio?: string;
  birthday?: string;
  age?: number;
  gender?: string;
  location?: string;
  website?: string;
  github?: string;
  socialLinks?: Record<string, unknown>;
  skills?: Record<string, unknown>;
}

export interface UpdateUserProfileRequest {
  role?: string;
  firstName?: string;
  lastName?: string;
  nickname?: string;
  displayName?: string;
  bio?: string;
  birthday?: string;
  age?: number;
  gender?: string;
  location?: string;
  website?: string;
  github?: string;
  socialLinks?: Record<string, unknown>;
  skills?: Record<string, unknown>;
}

// ========== 用户头像相关 ==========

export interface CreateOrUpdateUserAvatarRequest {
  userId: string;
  imageId: string;
}

export interface UpdateUserAvatarImageIDRequest {
  imageId: string;
}

// ========== 用户图片相关 ==========

export interface CreateUserImageRequest {
  userId: string;
  imageId: string;
  displayOrder?: number;
}

export interface UpdateUserImageDisplayOrderRequest {
  displayOrder: number;
}

export interface UpdateUserImageImageIDRequest {
  imageId: string;
}
