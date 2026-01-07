# User Phones Table

## 概述 / Overview

`directory.user_phones` 表用于存储用户的多个手机号码，支持主手机号、次要手机号等。

一个用户可以有多个手机号码（主手机号、次要手机号等）。

The `directory.user_phones` table stores multiple phone numbers per user, supporting primary, secondary, etc.

One user can have multiple phone numbers (primary, secondary, etc.).

## 使用场景 / Use Cases

### 1. 添加手机号
**场景**: 用户添加手机号码
**例子**:
- 用户添加手机号 "+86 13800138000"
- 操作：`INSERT INTO directory.user_phones (user_id, phone, country_code, is_primary, is_verified) VALUES ('zhangsan-uuid', '13800138000', '+86', true, false)`
- 结果：手机号添加，但未验证

### 2. 发送验证码
**场景**: 用户验证手机号时发送验证码
**例子**:
- 系统生成验证码 "123456"，5 分钟有效
- 操作：`UPDATE directory.user_phones SET verification_code = '123456', verification_expires_at = NOW() + INTERVAL '5 minutes' WHERE phone = '13800138000'`
- 发送 SMS 验证码
- 结果：验证码已发送

### 3. 验证手机号
**场景**: 用户输入验证码验证手机号
**例子**:
- 用户输入验证码 "123456"
- 查询：`SELECT * FROM directory.user_phones WHERE phone = '13800138000' AND verification_code = '123456' AND verification_expires_at > NOW()`
- 如果匹配，更新：`UPDATE directory.user_phones SET is_verified = true, verified_at = NOW(), verification_code = NULL, verification_expires_at = NULL WHERE phone = '13800138000'`
- 结果：手机号验证成功

