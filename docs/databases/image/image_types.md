# Image Types Table

## 概述 / Overview

`image.image_types` 表用于定义图片使用类别和规格（头像、背景、产品封面等）。

此表允许不同业务域有不同的图片要求。

The `image.image_types` table defines image usage categories and specifications (avatar, background, product_cover, etc.).

## 使用场景 / Use Cases

### 1. 创建图片类型
**场景**: 系统创建新的图片类型
**例子**:
- 创建"头像"类型，要求 1:1 比例，最大 1000x1000
- 操作：`INSERT INTO image.image_types (key, description, max_width, max_height, aspect_ratio, is_system) VALUES ('avatar', '用户头像', 1000, 1000, '1:1', true)`
- 结果：图片类型创建，上传时进行验证

### 2. 验证图片规格
**场景**: 上传图片时验证是否符合类型要求
**例子**:
- 用户上传头像，尺寸 500x500
- 查询类型要求：`SELECT max_width, max_height, aspect_ratio FROM image.image_types WHERE key = 'avatar'`
- 验证：宽度 <= 1000，高度 <= 1000，比例 = 1:1
- 如果符合，允许上传；否则拒绝

