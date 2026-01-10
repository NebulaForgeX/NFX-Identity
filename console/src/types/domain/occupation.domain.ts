// Occupation Domain Types - 基于 NFX-ID

export interface Occupation {
  id: string;
  profileId: string;
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
  skillsUsed?: string;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface CreateOccupationParams {
  profileId: string;
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

export interface UpdateOccupationParams {
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

export interface OccupationQueryParams {
  offset?: number;
  limit?: number;
  profileId?: string;
  sort?: string[];
}
