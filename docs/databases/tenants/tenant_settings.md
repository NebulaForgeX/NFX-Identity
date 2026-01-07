# Tenant Settings Table

## 概述 / Overview

`tenants.tenant_settings` 表用于租户级别的策略和配置。

企业策略：强制 MFA、密码最小长度、会话时长、允许的登录方法、邮箱域名限制等。

The `tenants.tenant_settings` table stores tenant-level policies and configurations.

## 使用场景 / Use Cases

### 1. 配置密码策略
**场景**: 租户设置密码策略
**例子**:
- 公司 A 要求密码最小 12 位，必须包含大小写字母、数字和符号
- 操作：`INSERT INTO tenants.tenant_settings (tenant_id, password_policy) VALUES ('company-a-uuid', '{"min_length": 12, "require_uppercase": true, "require_lowercase": true, "require_numbers": true, "require_symbols": true, "max_age_days": 90}')`
- 结果：密码策略保存，用户更改密码时进行验证

### 2. 强制 MFA
**场景**: 租户要求所有用户启用 MFA
**例子**:
- 公司 A 要求所有用户启用 MFA
- 操作：`UPDATE tenants.tenant_settings SET enforce_mfa = true, mfa_policy = '{"required_for_admin": true, "allowed_methods": ["totp", "sms"]}'::jsonb WHERE tenant_id = 'company-a-uuid'`
- 结果：所有用户必须启用 MFA 才能登录

### 3. 邮箱域名限制
**场景**: 租户限制只能使用公司邮箱注册
**例子**:
- 公司 A 只允许使用 @company-a.com 邮箱注册
- 操作：`UPDATE tenants.tenant_settings SET allowed_email_domains = ARRAY['company-a.com'] WHERE tenant_id = 'company-a-uuid'`
- 结果：只有公司邮箱可以注册和登录

### 4. 会话时长配置
**场景**: 租户设置会话过期时间
**例子**:
- 公司 A 设置会话 8 小时过期
- 操作：`UPDATE tenants.tenant_settings SET session_ttl_minutes = 480 WHERE tenant_id = 'company-a-uuid'`
- 结果：用户会话 8 小时后自动过期

