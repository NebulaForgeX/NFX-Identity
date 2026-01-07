# User Credentials Table

## 概述 / Overview

`auth.user_credentials` 表用于存储用户认证凭据，包括密码、Passkey、OAuth 链接、SAML、LDAP 等。

存储如何验证用户身份的信息（密码、passkey、oauth 链接等）。

The `auth.user_credentials` table stores user authentication credentials, including passwords, passkeys, OAuth links, SAML, LDAP, etc.

Stores how to verify a user (password, passkey, oauth links, etc.).

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 凭据记录的唯一标识符 |
| `user_id` | UUID | NOT NULL, UNIQUE | 用户 ID（引用 directory.users.id，应用级一致性） |
| `tenant_id` | UUID | NOT NULL | 多租户隔离（引用 tenants.tenants.id，应用级一致性） |
| `credential_type` | `auth.credential_type` ENUM | NOT NULL, DEFAULT 'password' | 凭据类型：password（密码）、passkey（Passkey）、oauth_link（OAuth 链接）、saml（SAML）、ldap（LDAP） |
| `password_hash` | VARCHAR(255) | NULL | 密码哈希值（用于 password 类型） |
| `hash_alg` | VARCHAR(50) | NULL | 哈希算法，例如：'bcrypt', 'argon2id', 'scrypt' |
| `hash_params` | JSONB | DEFAULT '{}'::jsonb | 算法参数：{"cost": 10, "salt": "..."} |
| `password_updated_at` | TIMESTAMP | NULL | 密码最后更新时间 |
| `last_success_login_at` | TIMESTAMP | NULL | 最后成功登录时间戳 |
| `status` | `auth.credential_status` ENUM | NOT NULL, DEFAULT 'active' | 凭据状态：active（活跃）、disabled（禁用）、expired（过期） |
| `must_change_password` | BOOLEAN | NOT NULL, DEFAULT false | 是否强制在下一次登录时更改密码 |
| `version` | INTEGER | NOT NULL, DEFAULT 1 | 乐观锁版本号，用于并发更新 |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| `updated_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 更新时间 |
| `deleted_at` | TIMESTAMP | NULL | 软删除时间戳 |

## 枚举类型 / Enum Types

### `auth.credential_type`
- **`password`**: 密码凭据
- **`passkey`**: Passkey（WebAuthn）凭据
- **`oauth_link`**: OAuth 链接凭据
- **`saml`**: SAML 凭据
- **`ldap`**: LDAP 凭据

### `auth.credential_status`
- **`active`**: 凭据活跃，可以使用
- **`disabled`**: 凭据被禁用，不能使用
- **`expired`**: 凭据已过期，不能使用

## 字段详解 / Field Details

### `user_id` (UUID)
- **用途**: 关联的用户 ID
- **引用**: 引用 `directory.users.id`（应用级一致性）
- **唯一性**: 每个用户只能有一条活跃凭据记录
- **不可为空**: 是

### `credential_type` (`auth.credential_type` ENUM)
- **用途**: 定义凭据类型
- **password**: 传统密码认证
- **passkey**: WebAuthn/Passkey 认证
- **oauth_link**: OAuth 第三方登录
- **saml**: SAML SSO 认证
- **ldap**: LDAP 目录服务认证
- **默认值**: `'password'`

### `password_hash` (VARCHAR(255))
- **用途**: 存储密码哈希值（仅用于 password 类型）
- **安全**: 永远不存储明文密码
- **算法**: 使用 `hash_alg` 字段指定的算法
- **可为空**: 是（非 password 类型时为空）

### `hash_alg` (VARCHAR(50))
- **用途**: 指定哈希算法
- **示例**: `'bcrypt'`, `'argon2id'`, `'scrypt'`
- **推荐**: 使用 `argon2id` 或 `bcrypt`
- **可为空**: 是

### `hash_params` (JSONB)
- **用途**: 存储算法参数
- **示例**: `{"cost": 10, "salt": "...", "memory": 65536, "iterations": 3}`
- **默认值**: `'{}'::jsonb`

### `must_change_password` (BOOLEAN)
- **用途**: 强制用户在下一次登录时更改密码
- **使用场景**: 
  - 管理员重置用户密码后
  - 检测到密码泄露后
  - 首次登录时
- **默认值**: `false`

### `version` (INTEGER)
- **用途**: 乐观锁版本号，用于并发更新控制
- **机制**: 更新时检查版本号，防止并发冲突
- **默认值**: `1`

## 使用场景 / Use Cases

### 1. 用户注册创建凭据
**场景**: 新用户注册时创建密码凭据
**例子**:
- 用户"张三"注册账户，设置密码 "MyPassword123!"
- 操作：`INSERT INTO auth.user_credentials (user_id, tenant_id, credential_type, password_hash, hash_alg, hash_params, password_updated_at) VALUES ('zhangsan-uuid', 'tenant-uuid', 'password', '$argon2id$v=19$m=65536,t=3,p=4$...', 'argon2id', '{"memory": 65536, "iterations": 3}', NOW())`
- 结果：用户可以使用该密码登录

### 2. 密码验证
**场景**: 用户登录时验证密码
**例子**:
- 用户输入密码 "MyPassword123!"
- 系统查询：`SELECT password_hash, hash_alg, hash_params FROM auth.user_credentials WHERE user_id = 'zhangsan-uuid' AND credential_type = 'password' AND status = 'active'`
- 使用相同的算法和参数对输入密码进行哈希，与存储的哈希值比较
- 如果匹配，更新 `last_success_login_at` 字段

### 3. 强制密码更改
**场景**: 管理员重置密码后，要求用户首次登录时更改
**例子**:
- 管理员重置用户密码
- 操作：`UPDATE auth.user_credentials SET password_hash = '...', hash_alg = 'argon2id', must_change_password = true, password_updated_at = NOW(), version = version + 1 WHERE user_id = 'zhangsan-uuid'`
- 结果：用户下次登录时，系统检查 `must_change_password = true`，强制用户更改密码

### 4. 禁用用户凭据
**场景**: 用户账户被禁用，禁用所有凭据
**例子**:
- 管理员禁用用户账户
- 操作：`UPDATE auth.user_credentials SET status = 'disabled', updated_at = NOW(), version = version + 1 WHERE user_id = 'zhangsan-uuid'`
- 结果：用户无法使用任何凭据登录

### 5. 添加 Passkey 凭据
**场景**: 用户启用 Passkey（WebAuthn）认证
**例子**:
- 用户在设置中启用 Passkey
- 操作：`INSERT INTO auth.user_credentials (user_id, tenant_id, credential_type, hash_params) VALUES ('zhangsan-uuid', 'tenant-uuid', 'passkey', '{"public_key": "...", "credential_id": "..."}')`
- 结果：用户可以使用 Passkey 登录

### 6. 链接 OAuth 账户
**场景**: 用户绑定第三方 OAuth 账户（如 Google、GitHub）
**例子**:
- 用户选择"使用 Google 登录"
- 操作：`INSERT INTO auth.user_credentials (user_id, tenant_id, credential_type, hash_params) VALUES ('zhangsan-uuid', 'tenant-uuid', 'oauth_link', '{"provider": "google", "provider_user_id": "...", "access_token_hash": "..."}')`
- 结果：用户可以使用 Google 账户登录

### 7. 密码历史检查
**场景**: 防止用户重复使用最近使用过的密码
**例子**:
- 用户尝试更改密码为 "OldPassword123!"（最近使用过）
- 系统查询 `auth.password_history` 表，检查新密码是否在历史记录中
- 如果在历史记录中，拒绝更改，提示"不能使用最近使用过的密码"

### 8. 并发更新保护
**场景**: 多个请求同时更新用户凭据时，使用乐观锁防止冲突
**例子**:
- 用户同时从两个设备更改密码
- 第一个请求：`UPDATE auth.user_credentials SET password_hash = 'hash1', version = version + 1 WHERE user_id = 'zhangsan-uuid' AND version = 1`
- 第二个请求：`UPDATE auth.user_credentials SET password_hash = 'hash2', version = version + 1 WHERE user_id = 'zhangsan-uuid' AND version = 1`
- 结果：只有第一个请求成功（version 变为 2），第二个请求失败（version 不匹配），需要重试

## 查询示例 / Query Examples

```sql
-- 查询用户的活跃凭据
SELECT credential_type, status, password_updated_at, last_success_login_at
FROM auth.user_credentials
WHERE user_id = 'user-uuid'
  AND status = 'active'
  AND deleted_at IS NULL;

-- 查询需要强制更改密码的用户
SELECT user_id, credential_type, password_updated_at
FROM auth.user_credentials
WHERE must_change_password = true
  AND status = 'active'
  AND deleted_at IS NULL;

-- 更新最后成功登录时间
UPDATE auth.user_credentials
SET last_success_login_at = NOW(), updated_at = NOW()
WHERE user_id = 'user-uuid'
  AND credential_type = 'password';
```

## 注意事项 / Notes

1. **密码安全**: 永远不存储明文密码，只存储哈希值
2. **算法选择**: 推荐使用 `argon2id` 或 `bcrypt`，避免使用 MD5、SHA1 等弱算法
3. **乐观锁**: 使用 `version` 字段实现乐观锁，防止并发更新冲突
4. **软删除**: 使用 `deleted_at` 字段实现软删除，保留历史记录
5. **多凭据支持**: 一个用户可以有多种类型的凭据（密码 + Passkey + OAuth）
6. **强制更改密码**: `must_change_password` 字段用于强制用户更改密码的场景

