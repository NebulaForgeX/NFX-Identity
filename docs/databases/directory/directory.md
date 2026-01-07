# Directory Schema 总览 / Directory Schema Overview

## 概述 / Overview

`directory` schema 是 NFX-Identity 平台的用户目录模块，负责管理用户基本信息、个人资料、联系方式、教育经历、工作经历等。

The `directory` schema is the user directory module of the NFX-Identity platform, responsible for managing user basic information, profiles, contact information, education history, work history, etc.

## 核心概念 / Core Concepts

### 1. 用户（Users）
- **定义**: 用户主表，存储基本认证信息
- **表**: `directory.users`
- **特点**: 用户名、密码哈希、状态管理

### 2. 用户资料（User Profiles）
- **定义**: 用户详细资料信息
- **表**: `directory.user_profiles`
- **特点**: 姓名、头像、背景、简介、社交链接、技能等

### 3. 用户邮箱（User Emails）
- **定义**: 用户的多个邮箱地址
- **表**: `directory.user_emails`
- **特点**: 支持主邮箱、次要邮箱、邮箱验证

### 4. 用户手机（User Phones）
- **定义**: 用户的多个手机号码
- **表**: `directory.user_phones`
- **特点**: 支持主手机号、次要手机号、手机验证

### 5. 用户偏好（User Preferences）
- **定义**: 用户设置和偏好
- **表**: `directory.user_preferences`
- **特点**: 主题、语言、时区、通知、隐私设置

### 6. 徽章（Badges）
- **定义**: 用户成就和认可
- **表**: `directory.badges`、`directory.user_badges`
- **特点**: 支持多种徽章类型和等级
- **详细文档**: [badges.md](./badges.md)、[user_badges.md](./user_badges.md)

### 7. 教育经历（User Educations）
- **定义**: 用户的教育历史
- **表**: `directory.user_educations`
- **特点**: 支持多个教育记录

### 8. 工作经历（User Occupations）
- **定义**: 用户的工作历史
- **表**: `directory.user_occupations`
- **特点**: 支持多个工作记录

## 表关系图 / Table Relationships

```
┌─────────────────┐
│     users       │
│  (用户主表)     │
└────────┬────────┘
         │
         │ 1:1
         │
    ┌────▼──────────────────┐
    │  user_profiles        │
    │  (用户资料)           │
    └───────────────────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────┐
    │  user_emails          │
    │  (用户邮箱)           │
    └───────────────────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────┐
    │  user_phones          │
    │  (用户手机)           │
    └───────────────────────┘
         │
         │ 1:1
         │
    ┌────▼──────────────────┐
    │  user_preferences     │
    │  (用户偏好)           │
    └───────────────────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────┐
    │  user_educations      │
    │  (教育经历)           │
    └───────────────────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────┐
    │  user_occupations     │
    │  (工作经历)           │
    └───────────────────────┘

┌─────────────────┐
│     badges      │
│  (徽章定义)     │
└────────┬────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────┐
    │  user_badges          │
    │  (用户徽章)           │
    └───────────────────────┘
```

## 表列表 / Table List

### 1. `directory.users` - 用户表
- **用途**: 用户主表，存储基本认证信息
- **关键字段**: `username`、`password_hash`、`status`、`is_verified`
- **详细文档**: [users.md](./users.md)

### 2. `directory.user_profiles` - 用户资料表
- **用途**: 存储用户详细资料
- **关键字段**: `user_id`、`first_name`、`last_name`、`avatar_id`、`social_links`、`skills`
- **详细文档**: [user_profiles.md](./user_profiles.md)

### 3. `directory.user_emails` - 用户邮箱表
- **用途**: 存储用户的多个邮箱地址
- **关键字段**: `user_id`、`email`、`is_primary`、`is_verified`
- **详细文档**: [user_emails.md](./user_emails.md)

### 4. `directory.user_phones` - 用户手机表
- **用途**: 存储用户的多个手机号码
- **关键字段**: `user_id`、`phone`、`is_primary`、`is_verified`
- **详细文档**: [user_phones.md](./user_phones.md)

### 5. `directory.user_preferences` - 用户偏好表
- **用途**: 存储用户设置和偏好
- **关键字段**: `user_id`、`theme`、`language`、`notifications`、`privacy`
- **详细文档**: [user_preferences.md](./user_preferences.md)

### 6. `directory.badges` - 徽章表
- **用途**: 存储徽章定义
- **关键字段**: `name`、`description`、`icon_url`、`category`
- **详细文档**: [badges.md](./badges.md)

### 7. `directory.user_badges` - 用户徽章表
- **用途**: 用户和徽章的多对多关系
- **关键字段**: `user_id`、`badge_id`、`level`、`earned_at`

### 8. `directory.user_educations` - 用户教育表
- **用途**: 存储用户教育经历
- **关键字段**: `user_id`、`school`、`degree`、`major`、`is_current`
- **详细文档**: [user_educations.md](./user_educations.md)

### 9. `directory.user_occupations` - 用户工作表
- **用途**: 存储用户工作经历
- **关键字段**: `user_id`、`company`、`position`、`is_current`
- **详细文档**: [user_occupations.md](./user_occupations.md)

## 用户注册流程 / User Registration Flow

```
1. 用户注册（users 表）
2. 创建用户资料（user_profiles 表）
3. 添加邮箱（user_emails 表）
4. 发送验证邮件
5. 用户验证邮箱
6. 激活用户账户（users.status = 'active'）
```

## 相关文档 / Related Documentation

- [users.md](./users.md) - 用户表详细文档
- [user_profiles.md](./user_profiles.md) - 用户资料表详细文档
- [user_emails.md](./user_emails.md) - 用户邮箱表详细文档
- [user_phones.md](./user_phones.md) - 用户手机表详细文档
- [user_preferences.md](./user_preferences.md) - 用户偏好表详细文档
- [badges.md](./badges.md) - 徽章表详细文档
- [user_badges.md](./user_badges.md) - 用户徽章表详细文档
- [user_educations.md](./user_educations.md) - 用户教育表详细文档
- [user_occupations.md](./user_occupations.md) - 用户工作表详细文档

