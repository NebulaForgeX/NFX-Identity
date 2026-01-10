// Profile Domain Types - 基于 NFX-ID

import type { Location, Social, ImageInfo } from "./user.domain";

export interface Profile {
  id: string;
  userId: string;
  username: string;
  email: string;
  userPhone?: string;
  userStatus: string;
  isVerified: boolean;
  firstName?: string;
  lastName?: string;
  nickname?: string;
  displayName?: string;
  avatarId?: string;
  backgroundId?: string;
  backgroundIds?: string[];
  images?: ImageInfo[];
  bio?: string;
  phone?: string;
  birthday?: string;
  age?: number;
  gender?: string;
  location?: Location;
  website?: string;
  github?: string;
  social?: Social;
  preferences?: Record<string, unknown>;
  skills?: Record<string, number>;
  privacySettings?: Record<string, unknown>;
  createdAt: string;
  updatedAt: string;
}

export interface CreateProfileParams {
  userId: string;
  firstName?: string;
  lastName?: string;
  nickname?: string;
  displayName?: string;
  avatarId?: string;
  backgroundId?: string;
  backgroundIds?: string[];
  bio?: string;
  phone?: string;
  birthday?: string;
  age?: number;
  gender?: string;
  location?: Location;
  website?: string;
  github?: string;
  social?: Social;
  preferences?: Record<string, unknown>;
  skills?: Record<string, number>;
  privacySettings?: Record<string, unknown>;
}

export interface UpdateProfileParams {
  firstName?: string;
  lastName?: string;
  nickname?: string;
  displayName?: string;
  avatarId?: string;
  backgroundId?: string;
  backgroundIds?: string[];
  bio?: string;
  phone?: string;
  birthday?: string;
  age?: number;
  gender?: string;
  location?: Location;
  website?: string;
  github?: string;
  social?: Social;
  preferences?: Record<string, unknown>;
  skills?: Record<string, number>;
  privacySettings?: Record<string, unknown>;
}

export interface ProfileQueryParams {
  offset?: number;
  limit?: number;
  userId?: string;
  search?: string;
  sort?: string[];
}
