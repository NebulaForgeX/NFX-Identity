# Image Schema 总览 / Image Schema Overview

## 概述 / Overview

`image` schema 是 NFX-Identity 平台的图片管理模块，负责管理所有上传图片的元数据、变体、类型和标签。

The `image` schema is the image management module of the NFX-Identity platform, responsible for managing metadata, variants, types, and tags for all uploaded images.

## 核心概念 / Core Concepts

### 1. 图片（Images）
- **定义**: 图片主表，存储所有上传图片的元数据
- **表**: `image.images`
- **特点**: 支持多租户、多应用、多用户

### 2. 图片类型（Image Types）
- **定义**: 图片使用类别和规格
- **表**: `image.image_types`
- **特点**: 定义不同类型的图片要求（尺寸、比例等）

### 3. 图片变体（Image Variants）
- **定义**: 原始图片的多个版本/衍生版本
- **表**: `image.image_variants`
- **特点**: 支持缩略图、调整大小、格式转换

### 4. 图片标签（Image Tags）
- **定义**: 图片标签，用于分类和搜索
- **表**: `image.image_tags`
- **特点**: 支持多个标签，支持 AI 自动标记

## 表关系图 / Table Relationships

```
┌─────────────────┐
│  image_types    │
│  (图片类型)     │
└────────┬────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────┐
    │     images            │
    │  (图片主表)           │
    └────┬──────────────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────┐
    │  image_variants       │
    │  (图片变体)           │
    └───────────────────────┘

┌─────────────────┐
│  image_tags     │
│  (图片标签)     │
│  (独立表)       │
└─────────────────┘
```

## 表列表 / Table List

### 1. `image.images` - 图片表
- **用途**: 图片主表，存储所有上传图片的元数据
- **关键字段**: `id`、`type_id`、`user_id`、`tenant_id`、`app_id`、`filename`、`url`、`is_public`
- **详细文档**: [images.md](./images.md)

### 2. `image.image_types` - 图片类型表
- **用途**: 定义图片使用类别和规格
- **关键字段**: `key`、`description`、`max_width`、`max_height`、`aspect_ratio`
- **详细文档**: [image_types.md](./image_types.md)

### 3. `image.image_variants` - 图片变体表
- **用途**: 存储原始图片的多个版本
- **关键字段**: `image_id`、`variant_key`、`width`、`height`、`url`
- **详细文档**: [image_variants.md](./image_variants.md)

### 4. `image.image_tags` - 图片标签表
- **用途**: 图片标签，用于分类和搜索
- **关键字段**: `image_id`、`tag`、`confidence`
- **详细文档**: [image_tags.md](./image_tags.md)

## 图片上传流程 / Image Upload Flow

```
1. 用户上传图片
2. 验证图片类型要求（尺寸、比例等）
3. 保存原图到存储后端
4. 创建图片记录（images 表）
5. 生成图片变体（缩略图、不同尺寸等）
6. 保存变体记录（image_variants 表）
7. 返回图片 URL 和变体 URL
```

## 多租户支持 / Multi-Tenant Support

- **租户隔离**: 所有图片都关联到 `tenant_id`
- **应用隔离**: 图片可以关联到 `app_id`
- **用户隔离**: 图片可以关联到 `user_id`
- **隐私控制**: `is_public` 字段控制图片可见性

## 相关文档 / Related Documentation

- [images.md](./images.md) - 图片表详细文档
- [image_types.md](./image_types.md) - 图片类型表详细文档
- [image_variants.md](./image_variants.md) - 图片变体表详细文档
- [image_tags.md](./image_tags.md) - 图片标签表详细文档

