# User Preferences Table

## 概述 / Overview

`directory.user_preferences` 表用于存储用户的设置和偏好，包括主题、语言、时区、通知设置、隐私设置等。

The `directory.user_preferences` table stores user-specific preferences and settings.

## 使用场景 / Use Cases

### 1. 设置主题
**场景**: 用户选择主题（浅色/深色/自动）
**例子**:
- 用户选择深色主题
- 操作：`UPDATE directory.user_preferences SET theme = 'dark', updated_at = NOW() WHERE user_id = 'zhangsan-uuid'`
- 结果：用户界面切换为深色主题

### 2. 设置语言
**场景**: 用户选择界面语言
**例子**:
- 用户选择中文
- 操作：`UPDATE directory.user_preferences SET language = 'zh', updated_at = NOW() WHERE user_id = 'zhangsan-uuid'`
- 结果：界面语言切换为中文

### 3. 配置通知设置
**场景**: 用户配置通知偏好
**例子**:
- 用户只启用邮件通知，禁用 SMS 和推送
- 操作：`UPDATE directory.user_preferences SET notifications = '{"email": true, "sms": false, "push": false}'::jsonb WHERE user_id = 'zhangsan-uuid'`
- 结果：通知设置保存

### 4. 配置隐私设置
**场景**: 用户配置隐私选项
**例子**:
- 用户设置资料可见性为"仅好友"，邮箱可见性为"私有"
- 操作：`UPDATE directory.user_preferences SET privacy = '{"profile_visibility": "friends_only", "email_visibility": "private"}'::jsonb WHERE user_id = 'zhangsan-uuid'`
- 结果：隐私设置保存

