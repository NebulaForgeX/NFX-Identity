// Profile Badge Domain Types - 基于 NFX-ID

import type { Badge } from "./badge.domain";

export interface ProfileBadge {
  id: string;
  profileId: string;
  badgeId: string;
  description?: string;
  level?: number;
  earnedAt: string;
  createdAt: string;
  updatedAt: string;
  badge?: Badge;
}

export interface CreateProfileBadgeParams {
  profileId: string;
  badgeId: string;
  description?: string;
  level?: number;
}

export interface UpdateProfileBadgeParams {
  description?: string;
  level?: number;
}

export interface ProfileBadgeQueryParams {
  profileId?: string;
  badgeId?: string;
}
