// Image Domain Types - 基于 NFX-ID Backend

export interface Image {
  id: string;
  url: string;
  typeId?: string;
  variantId?: string;
  width?: number;
  height?: number;
  size?: number;
  mimeType?: string;
  isPublic: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface ImageType {
  id: string;
  key: string;
  name: string;
  description?: string;
  createdAt: string;
  updatedAt: string;
}

export interface ImageVariant {
  id: string;
  imageId: string;
  variantType: string;
  url: string;
  width?: number;
  height?: number;
  size?: number;
  createdAt: string;
  updatedAt: string;
}

export interface ImageTag {
  id: string;
  imageId: string;
  tag: string;
  createdAt: string;
  updatedAt: string;
}
