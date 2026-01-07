# Domain Verifications Table

## 概述 / Overview

`tenants.domain_verifications` 表用于租户域名验证。

用于 SSO 和邮箱域名限制。

The `tenants.domain_verifications` table manages domain verification for tenants.

Used for SSO and email domain restrictions.

## 使用场景 / Use Cases

### 1. 验证域名
**场景**: 租户验证域名所有权
**例子**:
- 公司 A 需要验证域名 "company-a.com"
- 创建验证记录：`INSERT INTO tenants.domain_verifications (tenant_id, domain, verification_method, verification_token) VALUES ('company-a-uuid', 'company-a.com', 'dns', 'verification-token')`
- 要求租户在 DNS 中添加 TXT 记录
- 验证后更新：`UPDATE tenants.domain_verifications SET verification_status = 'VERIFIED', verified_at = NOW() WHERE domain = 'company-a.com'`
- 结果：域名验证成功，可以用于 SSO 和邮箱限制

### 2. 启用 SSO
**场景**: 租户验证域名后启用 SSO
**例子**:
- 域名验证成功后，启用 SSO
- 更新租户设置：`UPDATE tenants.tenant_settings SET login_policy = '{"allowed_methods": ["password", "sso"], "sso_provider": "company-a.com"}'::jsonb WHERE tenant_id = 'company-a-uuid'`
- 结果：租户可以使用 SSO 登录

