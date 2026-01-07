# MFA Factors Table

## 概述 / Overview

`auth.mfa_factors` 表用于存储多因素认证（MFA）凭据，包括 TOTP、SMS、Email、WebAuthn、备份代码等。

The `auth.mfa_factors` table stores multi-factor authentication (MFA) credentials, including TOTP, SMS, Email, WebAuthn, backup codes.

## 使用场景 / Use Cases

### 1. 启用 TOTP 双因素认证
**场景**: 用户启用 TOTP（如 Google Authenticator）
**例子**:
- 用户在设置中启用 TOTP
- 系统生成 TOTP 密钥
- 操作：`INSERT INTO auth.mfa_factors (factor_id, tenant_id, user_id, type, secret_encrypted, name, enabled) VALUES ('factor-uuid', 'tenant-uuid', 'zhangsan-uuid', 'totp', 'encrypted-secret', 'My iPhone', true)`
- 结果：用户可以使用 TOTP 应用进行二次验证

### 2. 验证 MFA 代码
**场景**: 用户登录时输入 MFA 代码
**例子**:
- 用户输入 TOTP 代码 "123456"
- 查询：`SELECT secret_encrypted, type FROM auth.mfa_factors WHERE user_id = 'zhangsan-uuid' AND type = 'totp' AND enabled = true AND deleted_at IS NULL`
- 解密密钥，验证 TOTP 代码
- 如果验证成功，更新 `last_used_at`：`UPDATE auth.mfa_factors SET last_used_at = NOW() WHERE factor_id = 'factor-uuid'`

### 3. 添加 WebAuthn Passkey
**场景**: 用户添加 Passkey（WebAuthn）作为 MFA
**例子**:
- 用户选择"添加 Passkey"
- 操作：`INSERT INTO auth.mfa_factors (factor_id, tenant_id, user_id, type, secret_encrypted, name, enabled) VALUES ('passkey-uuid', 'tenant-uuid', 'zhangsan-uuid', 'webauthn', 'encrypted-public-key', 'My MacBook', true)`
- 结果：用户可以使用 Passkey 进行二次验证

### 4. 禁用 MFA 因子
**场景**: 用户丢失设备，禁用 MFA 因子
**例子**:
- 用户报告设备丢失
- 操作：`UPDATE auth.mfa_factors SET enabled = false, deleted_at = NOW() WHERE factor_id = 'lost-factor-uuid'`
- 结果：该 MFA 因子被禁用，不能用于验证

### 5. 生成备份代码
**场景**: 用户启用 MFA 时生成备份代码
**例子**:
- 用户启用 MFA 后，系统生成 10 个备份代码
- 操作：`UPDATE auth.mfa_factors SET recovery_codes_hash = 'hashed-backup-codes' WHERE factor_id = 'factor-uuid'`
- 结果：用户可以使用备份代码在无法使用主 MFA 因子时进行验证

