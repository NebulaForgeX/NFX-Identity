# User Emails Table

## 概述 / Overview

`directory.user_emails` 表用于存储用户的多个邮箱地址，支持主邮箱、次要邮箱等。

一个用户可以有多个邮箱地址（主邮箱、次要邮箱等）。

The `directory.user_emails` table stores multiple email addresses per user, supporting primary, secondary, etc.

One user can have multiple email addresses (primary, secondary, etc.).

## 使用场景 / Use Cases

### 1. 添加邮箱
**场景**: 用户添加邮箱地址
**例子**:
- 用户添加主邮箱 "zhangsan@example.com"
- 操作：`INSERT INTO directory.user_emails (user_id, email, is_primary, is_verified) VALUES ('zhangsan-uuid', 'zhangsan@example.com', true, false)`
- 结果：邮箱添加，但未验证

### 2. 验证邮箱
**场景**: 用户验证邮箱地址
**例子**:
- 用户点击邮箱验证链接
- 操作：`UPDATE directory.user_emails SET is_verified = true, verified_at = NOW(), verification_token = NULL WHERE email = 'zhangsan@example.com'`
- 同时更新用户状态：`UPDATE directory.users SET is_verified = true WHERE id = 'zhangsan-uuid'`
- 结果：邮箱验证成功

### 3. 设置主邮箱
**场景**: 用户更改主邮箱
**例子**:
- 用户将 "zhangsan@work.com" 设置为主邮箱
- 先取消旧主邮箱：`UPDATE directory.user_emails SET is_primary = false WHERE user_id = 'zhangsan-uuid' AND is_primary = true`
- 设置新主邮箱：`UPDATE directory.user_emails SET is_primary = true WHERE email = 'zhangsan@work.com'`
- 结果：主邮箱更改成功

### 4. 通过邮箱查找用户
**场景**: 用户登录时使用邮箱查找用户
**例子**:
- 用户输入邮箱 "zhangsan@example.com" 登录
- 查询：`SELECT ue.user_id, u.username FROM directory.user_emails ue JOIN directory.users u ON ue.user_id = u.id WHERE ue.email = 'zhangsan@example.com' AND ue.is_verified = true AND u.deleted_at IS NULL`
- 结果：找到用户，继续密码验证

