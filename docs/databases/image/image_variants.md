# Image Variants Table

## 概述 / Overview

`image.image_variants` 表用于存储原始图片的多个版本/衍生版本。

支持：缩略图、调整大小版本、格式转换（webp/jpg/png）、裁剪版本。

这可以高效地为不同用例提供适当大小的图片。

The `image.image_variants` table stores multiple versions/derivatives of an original image.

Supports: thumbnails, resized versions, format conversions (webp/jpg/png), cropped versions.

## 使用场景 / Use Cases

### 1. 生成缩略图
**场景**: 上传图片后自动生成缩略图
**例子**:
- 用户上传 2000x2000 的原图
- 系统生成 150x150 缩略图
- 操作：`INSERT INTO image.image_variants (image_id, variant_key, width, height, size, mime_type, storage_path, url) VALUES ('image-uuid', 'thumbnail', 150, 150, 15000, 'image/jpeg', '/storage/thumbnails/thumb-uuid.jpg', 'https://cdn.example.com/thumbnails/thumb-uuid.jpg')`
- 结果：缩略图创建，用于列表展示

### 2. 生成不同尺寸
**场景**: 为不同场景生成不同尺寸
**例子**:
- 生成小图（300px）、中图（720px）、大图（1080px）
- 操作：
  - `INSERT INTO image.image_variants (image_id, variant_key, width, height, storage_path, url) VALUES ('image-uuid', 'small', 300, 300, '/storage/small/small-uuid.jpg', 'https://cdn.example.com/small/small-uuid.jpg')`
  - `INSERT INTO image.image_variants (image_id, variant_key, width, height, storage_path, url) VALUES ('image-uuid', 'medium', 720, 720, '/storage/medium/medium-uuid.jpg', 'https://cdn.example.com/medium/medium-uuid.jpg')`
  - `INSERT INTO image.image_variants (image_id, variant_key, width, height, storage_path, url) VALUES ('image-uuid', 'large', 1080, 1080, '/storage/large/large-uuid.jpg', 'https://cdn.example.com/large/large-uuid.jpg')`
- 结果：多个尺寸创建，根据场景选择合适尺寸

### 3. 格式转换
**场景**: 将图片转换为 WebP 格式
**例子**:
- 将原图转换为 WebP 格式
- 操作：`INSERT INTO image.image_variants (image_id, variant_key, width, height, mime_type, storage_path, url) VALUES ('image-uuid', 'webp', 2000, 2000, 'image/webp', '/storage/webp/webp-uuid.webp', 'https://cdn.example.com/webp/webp-uuid.webp')`
- 结果：WebP 版本创建，用于现代浏览器

### 4. 查询图片变体
**场景**: 获取图片的所有变体
**例子**:
- 查询图片的所有变体：`SELECT variant_key, width, height, url FROM image.image_variants WHERE image_id = 'image-uuid' ORDER BY width`
- 结果：返回所有变体，用于前端选择合适尺寸

