# User Profiles Table

## 概述 / Overview

`directory.user_profiles` 表用于存储用户的详细资料信息，包括姓名、头像、背景、简介、社交链接等。

The `directory.user_profiles` table stores detailed user profile information, including name, avatar, background, bio, social links, etc.

## 使用场景 / Use Cases

### 1. 创建用户资料
**场景**: 用户注册后创建个人资料
**例子**:
- 用户"张三"完善个人资料
- 操作：`INSERT INTO directory.user_profiles (user_id, first_name, last_name, display_name, bio, location) VALUES ('zhangsan-uuid', '三', '张', '张三', '软件工程师', '北京, 中国')`
- 结果：用户资料创建，可以在个人页面显示

### 2. 上传头像
**场景**: 用户上传头像
**例子**:
- 用户上传头像，获得 `avatar_id = 'image-uuid'`
- 操作：`UPDATE directory.user_profiles SET avatar_id = 'image-uuid', updated_at = NOW() WHERE user_id = 'zhangsan-uuid'`
- 结果：用户头像更新，个人页面显示新头像

### 3. 设置社交链接
**场景**: 用户添加社交链接
**例子**:
- 用户添加 Twitter 和 LinkedIn 链接
- 操作：`UPDATE directory.user_profiles SET social_links = '{"twitter": "https://twitter.com/zhangsan", "linkedin": "https://linkedin.com/in/zhangsan"}'::jsonb WHERE user_id = 'zhangsan-uuid'`
- 结果：社交链接保存，可以在个人页面显示

### 4. 设置技能
**场景**: 用户添加技能标签
**例子**:
- 用户添加技能：Golang 10级，Python 8级
- 操作：`UPDATE directory.user_profiles SET skills = '{"golang": 10, "python": 8, "javascript": 7}'::jsonb WHERE user_id = 'zhangsan-uuid'`
- 结果：技能保存，可以用于搜索和匹配

### 5. 查询用户资料
**场景**: 查看用户完整资料
**例子**:
- 查询用户资料：`SELECT first_name, last_name, display_name, bio, avatar_id, social_links, skills FROM directory.user_profiles WHERE user_id = 'zhangsan-uuid' AND deleted_at IS NULL`
- 结果：返回用户完整资料信息

