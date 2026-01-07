# Image Tags Table

## 概述 / Overview

`image.image_tags` 表用于存储图片标签（AI 生成的标签和用户定义的标签）。

支持基于内容的搜索和发现（例如，查找所有标记为 'cat'、'shoes'、'selfie' 的图片）。

支持手动用户标记和自动 AI 标记。

The `image.image_tags` table stores tags for images (AI-generated labels and user-defined tags).

Enables content-based search and discovery.

## 使用场景 / Use Cases

### 1. 用户添加标签
**场景**: 用户为图片添加标签
**例子**:
- 用户为图片添加标签 "nature"、"landscape"
- 操作：
  - `INSERT INTO image.image_tags (image_id, tag) VALUES ('image-uuid', 'nature')`
  - `INSERT INTO image.image_tags (image_id, tag) VALUES ('image-uuid', 'landscape')`
- 结果：标签添加，可以用于搜索

### 2. AI 自动标记
**场景**: 系统使用 AI 自动为图片生成标签
**例子**:
- AI 分析图片，生成标签 "cat"（置信度 0.95）、"indoor"（置信度 0.87）
- 操作：
  - `INSERT INTO image.image_tags (image_id, tag, confidence) VALUES ('image-uuid', 'cat', 0.95)`
  - `INSERT INTO image.image_tags (image_id, tag, confidence) VALUES ('image-uuid', 'indoor', 0.87)`
- 结果：AI 标签添加，可以用于搜索和分类

### 3. 按标签搜索图片
**场景**: 用户按标签搜索图片
**例子**:
- 用户搜索标签 "nature" 的图片
- 查询：`SELECT DISTINCT i.id, i.url FROM image.images i JOIN image.image_tags it ON i.id = it.image_id WHERE it.tag = 'nature' AND i.deleted_at IS NULL ORDER BY it.confidence DESC NULLS LAST`
- 结果：返回所有标记为 "nature" 的图片，按置信度排序

### 4. 过滤低置信度标签
**场景**: 只显示高置信度的 AI 标签
**例子**:
- 查询置信度 >= 0.8 的标签：`SELECT image_id, tag, confidence FROM image.image_tags WHERE confidence >= 0.8 ORDER BY confidence DESC`
- 结果：只返回高置信度标签，用于展示

