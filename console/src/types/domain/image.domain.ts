// Image Domain Types - 基于 NFX-ID

export interface Image {
  id: string;
  typeId?: string;
  userId?: string;
  sourceDomain?: string;
  filename: string;
  originalFilename: string;
  mimeType: string;
  size: number;
  width?: number;
  height?: number;
  storagePath: string;
  url?: string;
  isPublic: boolean;
  metadata?: Record<string, unknown>;
  createdAt: string;
  updatedAt: string;
}

export interface CreateImageParams {
  typeId?: string;
  userId?: string;
  sourceDomain?: string;
  filename: string;
  originalFilename: string;
  mimeType: string;
  size: number;
  width?: number;
  height?: number;
  storagePath: string;
  url?: string;
  isPublic?: boolean;
  metadata?: Record<string, unknown>;
}

export interface UpdateImageParams {
  typeId?: string;
  userId?: string;
  sourceDomain?: string;
  filename?: string;
  originalFilename?: string;
  mimeType?: string;
  size?: number;
  width?: number;
  height?: number;
  storagePath?: string;
  url?: string;
  isPublic?: boolean;
  metadata?: Record<string, unknown>;
}

export interface ImageQueryParams {
  page?: number;
  pageSize?: number;
  search?: string;
  typeId?: string;
  userId?: string;
  isPublic?: boolean;
  orderBy?: string;
  order?: string;
}

export interface ImageType {
  id: string;
  key: string;
  description?: string;
  maxWidth?: number;
  maxHeight?: number;
  aspectRatio?: string;
  isSystem: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface CreateImageTypeParams {
  key: string;
  description?: string;
  maxWidth?: number;
  maxHeight?: number;
  aspectRatio?: string;
  isSystem?: boolean;
}

export interface UpdateImageTypeParams {
  key?: string;
  description?: string;
  maxWidth?: number;
  maxHeight?: number;
  aspectRatio?: string;
  isSystem?: boolean;
}

export interface ImageTypeQueryParams {
  page?: number;
  pageSize?: number;
  search?: string;
  isSystem?: boolean;
  orderBy?: string;
  order?: string;
}
