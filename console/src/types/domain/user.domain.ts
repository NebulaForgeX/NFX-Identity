// User Domain Types - 基于 NFX-ID

export interface Role {
  id: string;
  name: string;
  description?: string;
  permissions: string[];
  isSystem: boolean;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface User {
  id: string;
  username: string;
  email: string;
  phone: string;
  roleId?: string;
  status: string;
  isVerified: boolean;
  lastLoginAt?: string;
  createdAt: string;
  updatedAt: string;
  role?: Role;
  profile?: Profile;
}

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

export interface Location {
  country?: string;
  province?: string;
  city?: string;
  timezone?: string;
}

export interface Social {
  twitter?: string;
  linkedin?: string;
  instagram?: string;
  youtube?: string;
}

export interface ImageInfo {
  id: string;
  typeId?: string;
  url?: string;
  isPublic: boolean;
}

// Login
export interface LoginParams {
  identifier: string; // username, email 或 phone
  password: string;
}

export interface RefreshTokenParams {
  refreshToken: string;
}

export interface LoginResponse {
  accessToken: string;
  refreshToken: string;
  user?: User;
}

export interface RefreshResponse {
  accessToken: string;
  refreshToken: string;
}

// User Create
export interface CreateUserParams {
  username: string;
  email: string;
  phone: string;
  password: string;
  roleId?: string;
  status?: string;
}

// User Update
export interface UpdateUserParams {
  username?: string;
  email?: string;
  phone?: string;
  roleId?: string;
  status?: string;
}

// User Query
export interface UserQueryParams {
  offset?: number;
  limit?: number;
  search?: string;
  status?: string;
  roleId?: string;
  sort?: string[];
}
