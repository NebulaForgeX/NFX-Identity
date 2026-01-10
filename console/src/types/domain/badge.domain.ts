// Badge Domain Types - 基于 NFX-ID

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
}

export interface CreateBadgeParams {
  name: string;
  description?: string;
  iconUrl?: string;
  color?: string;
  category?: string;
  isSystem?: boolean;
}

export interface UpdateBadgeParams {
  name?: string;
  description?: string;
  iconUrl?: string;
  color?: string;
  category?: string;
  isSystem?: boolean;
}

export interface BadgeQueryParams {
  offset?: number;
  limit?: number;
  search?: string;
  category?: string;
  isSystem?: boolean;
  sort?: string[];
}
