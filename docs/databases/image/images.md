# Images Table

## 概述 / Overview

`image.images` 表用于存储所有上传图片的元数据和文件信息。

这是所有服务中所有上传图片的主表。

注意：user_id 没有外键约束，以保持服务独立性（微服务架构）。

The `image.images` table stores original image metadata and file information.

This is the main table for all uploaded images across all services.

## 使用场景 / Use Cases

### 1. 上传图片
**场景**: 用户上传头像
**例子**:
- 用户上传头像，文件大小 500KB，尺寸 500x500
- 操作：`INSERT INTO image.images (type_id, user_id, tenant_id, filename, original_filename, mime_type, size, width, height, storage_path, url, is_public) VALUES ((SELECT id FROM image.image_types WHERE key = 'avatar'), 'zhangsan-uuid', 'tenant-uuid', 'avatar-500x500.jpg', 'my-avatar.jpg', 'image/jpeg', 512000, 500, 500, '/storage/avatars/avatar-uuid.jpg', 'https://cdn.example.com/avatars/avatar-uuid.jpg', false)`
- 结果：图片元数据保存，文件存储到存储后端

### 2. 生成图片变体
**场景**: 系统自动生成缩略图
**例子**:
- 上传原图后，系统生成 150x150 缩略图
- 操作：`INSERT INTO image.image_variants (image_id, variant_key, width, height, size, mime_type, storage_path, url) VALUES ('image-uuid', 'thumbnail', 150, 150, 15000, 'image/jpeg', '/storage/avatars/avatar-uuid-thumb.jpg', 'https://cdn.example.com/avatars/avatar-uuid-thumb.jpg')`
- 结果：缩略图变体创建，用于快速加载

### 3. 查询用户图片
**场景**: 查询用户的所有图片
**例子**:
- 查询用户的所有公开图片：`SELECT id, filename, url, width, height, created_at FROM image.images WHERE user_id = 'zhangsan-uuid' AND is_public = true AND deleted_at IS NULL ORDER BY created_at DESC`
- 结果：返回用户的所有公开图片

### 4. 按类型查询图片
**场景**: 查询特定类型的图片
**例子**:
- 查询所有头像：`SELECT i.id, i.url, i.user_id FROM image.images i JOIN image.image_types it ON i.type_id = it.id WHERE it.key = 'avatar' AND i.deleted_at IS NULL`
- 结果：返回所有头像图片

### 5. 软删除图片
**场景**: 删除图片（保留元数据）
**例子**:
- 用户删除图片
- 操作：`UPDATE image.images SET deleted_at = NOW(), updated_at = NOW() WHERE id = 'image-uuid'`
- 结果：图片被软删除，元数据保留但不可访问

