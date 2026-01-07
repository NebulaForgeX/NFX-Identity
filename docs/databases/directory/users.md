# Users Table

## 概述 / Overview

`directory.users` 表是用户主表，用于身份认证和用户管理。

The `directory.users` table is the main user table for authentication and user management.

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 用户的唯一标识符 |
| `username` | VARCHAR(50) | NOT NULL, UNIQUE | 用户名（唯一） |
| `password_hash` | VARCHAR(255) | NOT NULL | 密码哈希值（永远不存储明文） |
| `status` | `directory.user_status` ENUM | NOT NULL, DEFAULT 'pending' | 用户状态：pending（待激活）、active（活跃）、deactive（停用） |
| `is_verified` | BOOLEAN | NOT NULL, DEFAULT false | 是否已验证（邮箱/手机验证） |
| `last_login_at` | TIMESTAMP | NULL | 最后登录时间 |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| `updated_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 更新时间 |
| `deleted_at` | TIMESTAMP | NULL | 软删除时间戳 |

## 使用场景 / Use Cases

### 1. 用户注册
**场景**: 新用户注册账户
**例子**:
- 用户"张三"注册，用户名 "zhangsan"，密码 "MyPassword123!"
- 操作：`INSERT INTO directory.users (username, password_hash, status) VALUES ('zhangsan', '$argon2id$v=19$m=65536,t=3,p=4$...', 'pending')`
- 结果：创建用户，状态为 `pending`，等待邮箱验证激活

### 2. 用户登录验证
**场景**: 用户登录时验证用户名和密码
**例子**:
- 用户输入用户名 "zhangsan" 和密码
- 查询：`SELECT id, password_hash, status FROM directory.users WHERE username = 'zhangsan' AND deleted_at IS NULL`
- 验证密码哈希
- 如果成功，更新 `last_login_at`：`UPDATE directory.users SET last_login_at = NOW(), updated_at = NOW() WHERE id = 'user-uuid'`

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

