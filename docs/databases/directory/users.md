# Users Table

## 概述 / Overview

`directory.users` 表是用户主表，用于用户目录管理。注意：认证凭据存储在 `auth.user_credentials` 表中。

The `directory.users` table is the main user table for user directory management. Note: Authentication credentials are stored in `auth.user_credentials` table.

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 用户的唯一标识符 |
| `tenant_id` | UUID | NOT NULL | 租户ID（多租户隔离） |
| `username` | VARCHAR(50) | NOT NULL, UNIQUE | 用户名（唯一） |
| `status` | `directory.user_status` ENUM | NOT NULL, DEFAULT 'pending' | 用户状态：pending（待激活）、active（活跃）、deactive（停用） |
| `is_verified` | BOOLEAN | NOT NULL, DEFAULT false | 是否已验证（邮箱/手机验证） |
| `last_login_at` | TIMESTAMP | NULL | 最后登录时间 |
| `primary_email_id` | UUID | NULL, REFERENCES user_emails.id | 主邮箱ID（外键指向 user_emails 表） |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| `updated_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 更新时间 |
| `deleted_at` | TIMESTAMP | NULL | 软删除时间戳 |

## 使用场景 / Use Cases

### 1. 用户注册
**场景**: 新用户注册账户
**例子**:
- 用户"张三"注册，用户名 "zhangsan"
- 操作：`INSERT INTO directory.users (tenant_id, username, status) VALUES ('tenant-uuid', 'zhangsan', 'pending')`
- 同时在 `auth.user_credentials` 中创建密码凭据
- 结果：创建用户，状态为 `pending`，等待邮箱验证激活

### 2. 查询用户
**场景**: 根据用户名查找用户
**例子**:
- 用户输入用户名 "zhangsan"
- 查询：`SELECT id, username, status, is_verified, primary_email_id FROM directory.users WHERE username = 'zhangsan' AND deleted_at IS NULL`
- 然后从 `auth.user_credentials` 中查询密码哈希进行验证

### 2.1. 设置主邮箱
**场景**: 用户设置主邮箱
**例子**:
- 用户有多个邮箱，设置其中一个为主邮箱
- 操作：`UPDATE directory.users SET primary_email_id = 'email-uuid', updated_at = NOW() WHERE id = 'user-uuid'`
- 结果：主邮箱设置成功，可以通过主邮箱快速查找用户

### 2.2. 更新最后登录时间
**场景**: 用户登录成功后更新最后登录时间
**例子**:
- 用户登录成功
- 操作：`UPDATE directory.users SET last_login_at = NOW(), updated_at = NOW() WHERE id = 'user-uuid'`
- 结果：最后登录时间更新，可用于统计和审计

### 3. 激活用户账户
**场景**: 用户验证邮箱后激活账户
**例子**:
- 用户点击邮箱验证链接
- 操作：`UPDATE directory.users SET status = 'active', is_verified = true, updated_at = NOW() WHERE id = 'user-uuid'`
- 结果：用户账户激活，可以正常使用

### 4. 停用用户账户
**场景**: 管理员停用用户账户
**例子**:
- 管理员停用违规用户
- 操作：`UPDATE directory.users SET status = 'deactive', updated_at = NOW() WHERE id = 'user-uuid'`
- 结果：用户无法登录，但数据保留

### 5. 软删除用户
**场景**: 删除用户账户（保留数据用于审计）
**例子**:
- 管理员删除用户
- 操作：`UPDATE directory.users SET deleted_at = NOW(), updated_at = NOW() WHERE id = 'user-uuid'`
- 结果：用户被软删除，数据保留但不可访问

