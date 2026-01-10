// Image Request Types - 基于 NFX-ID Backend

// ========== 图片相关 ==========

export interface CreateImageRequest {
  url: string;
  typeId?: string;
  variantId?: string;
  width?: number;
  height?: number;
  size?: number;
  mimeType?: string;
  isPublic: boolean;
}

export interface UpdateImageRequest {
  url?: string;
  typeId?: string;
  variantId?: string;
  width?: number;
  height?: number;
  size?: number;
  mimeType?: string;
  isPublic?: boolean;
}

// ========== 图片类型相关 ==========

export interface CreateImageTypeRequest {
  key: string;
  name: string;
  description?: string;
}

export interface UpdateImageTypeRequest {
  name?: string;
  description?: string;
}

// ========== 图片变体相关 ==========

export interface CreateImageVariantRequest {
  imageId: string;
  variantType: string;
  url: string;
  width?: number;
  height?: number;
  size?: number;
}

export interface UpdateImageVariantRequest {
  variantType?: string;
  url?: string;
  width?: number;
  height?: number;
  size?: number;
}

// ========== 图片标签相关 ==========

export interface CreateImageTagRequest {
  imageId: string;
  tag: string;
}

export interface UpdateImageTagRequest {
  tag: string;
}
