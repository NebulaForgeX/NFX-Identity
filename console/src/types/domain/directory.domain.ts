// Directory Domain Types - 基于 NFX-ID Backend

export interface User {
  id: string;
  tenantId: string;
  username: string;
  status: string;
  isVerified: boolean;
  lastLoginAt?: string;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface Badge {
  id: string;
  name: string;
  description?: string;
  iconUrl?: string;
  color?: string;
  category?: string;
  isSystem: boolean;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface UserBadge {
  id: string;
  userId: string;
  badgeId: string;
  awardedAt: string;
  createdAt: string;
}

export interface UserEducation {
  id: string;
  userId: string;
  school: string;
  degree?: string;
  field?: string;
  startDate?: string;
  endDate?: string;
  description?: string;
  createdAt: string;
  updatedAt: string;
}

export interface UserEmail {
  id: string;
  userId: string;
  email: string;
  isPrimary: boolean;
  isVerified: boolean;
  verifiedAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface UserOccupation {
  id: string;
  userId: string;
  company: string;
  position?: string;
  startDate?: string;
  endDate?: string;
  description?: string;
  createdAt: string;
  updatedAt: string;
}

export interface UserPhone {
  id: string;
  userId: string;
  phone: string;
  isPrimary: boolean;
  isVerified: boolean;
  verifiedAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface UserPreference {
  id: string;
  userId: string;
  key: string;
  value: string;
  createdAt: string;
  updatedAt: string;
}

export interface UserProfile {
  id: string;
  userId: string;
  firstName?: string;
  lastName?: string;
  displayName?: string;
  bio?: string;
  avatarId?: string;
  createdAt: string;
  updatedAt: string;
}
