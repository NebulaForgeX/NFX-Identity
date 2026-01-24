// Directory Domain Types - 基于 NFX-ID Backend

import type { DashboardBackgroundType } from "@/types";

import { UserStatus } from "./enums";

export interface User {
  id: string;
  tenantId: string;
  username: string;
  status: UserStatus;
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
  major?: string;
  fieldOfStudy?: string;
  startDate?: string;
  endDate?: string;
  isCurrent: boolean;
  description?: string;
  grade?: string;
  activities?: string;
  achievements?: string;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
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
  position: string;
  department?: string;
  industry?: string;
  location?: string;
  employmentType?: string;
  startDate?: string;
  endDate?: string;
  isCurrent: boolean;
  description?: string;
  responsibilities?: string;
  achievements?: string;
  skillsUsed?: string[];
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface UserPhone {
  id: string;
  userId: string;
  phone: string;
  countryCode?: string;
  isPrimary: boolean;
  isVerified: boolean;
  verifiedAt?: string;
  verificationCode?: string;
  verificationExpiresAt?: string;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface UserPreference {
  id: string;
  userId: string;
  theme?: string;
  language?: string;
  timezone?: string;
  dashboardBackground?: DashboardBackgroundType;
  notifications?: Record<string, unknown>;
  privacy?: Record<string, unknown>;
  display?: Record<string, unknown>;
  other?: Record<string, unknown>;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface UserProfile {
  id: string;
  userId: string;
  role?: string;
  firstName?: string;
  lastName?: string;
  nickname?: string;
  displayName?: string;
  avatarId?: string;
  backgroundId?: string;
  backgroundIds?: string[];
  bio?: string;
  birthday?: string;
  age?: number;
  gender?: string;
  location?: string;
  website?: string;
  github?: string;
  socialLinks?: Record<string, unknown>;
  skills?: Record<string, unknown>;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}
