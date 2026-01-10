// Education Domain Types - 基于 NFX-ID

export interface Education {
  id: string;
  profileId: string;
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

export interface CreateEducationParams {
  profileId: string;
  school: string;
  degree?: string;
  major?: string;
  fieldOfStudy?: string;
  startDate?: string;
  endDate?: string;
  isCurrent?: boolean;
  description?: string;
  grade?: string;
  activities?: string;
  achievements?: string;
}

export interface UpdateEducationParams {
  school?: string;
  degree?: string;
  major?: string;
  fieldOfStudy?: string;
  startDate?: string;
  endDate?: string;
  isCurrent?: boolean;
  description?: string;
  grade?: string;
  activities?: string;
  achievements?: string;
}

export interface EducationQueryParams {
  offset?: number;
  limit?: number;
  profileId?: string;
  sort?: string[];
}
