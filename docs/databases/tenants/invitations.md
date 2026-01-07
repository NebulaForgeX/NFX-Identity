# Invitations Table

## 概述 / Overview

`tenants.invitations` 表用于邀请机制（企业必需）。

"公司 A 管理员邀请同事加入后端"依赖此表。

企业邀请表必须考虑安全性和可追溯性。

The `tenants.invitations` table manages invitation mechanism (enterprise essential).

"Company A admin invites colleagues to join backend" relies on this.

## 使用场景 / Use Cases

### 1. 发送邀请
**场景**: 管理员邀请新成员加入租户
**例子**:
- 管理员邀请 "lisi@example.com" 加入公司 A
- 生成邀请令牌并哈希
- 操作：`INSERT INTO tenants.invitations (invite_id, tenant_id, email, token_hash, expires_at, invited_by) VALUES ('invite-uuid', 'company-a-uuid', 'lisi@example.com', 'hashed-token', NOW() + INTERVAL '7 days', 'admin-uuid')`
- 发送邀请邮件（包含邀请链接）
- 结果：邀请发送，7 天内有效

### 2. 接受邀请
**场景**: 用户接受邀请
**例子**:
- 用户点击邀请链接，验证令牌
- 查询：`SELECT * FROM tenants.invitations WHERE invite_id = 'invite-uuid' AND status = 'PENDING' AND expires_at > NOW()`
- 如果有效，更新邀请：`UPDATE tenants.invitations SET status = 'ACCEPTED', accepted_by_user_id = 'lisi-uuid', accepted_at = NOW() WHERE invite_id = 'invite-uuid'`
- 创建成员：`INSERT INTO tenants.members (tenant_id, user_id, status, source, joined_at) VALUES ('company-a-uuid', 'lisi-uuid', 'ACTIVE', 'INVITE', NOW())`
- 结果：用户接受邀请，成为租户成员

### 3. 撤销邀请
**场景**: 管理员撤销邀请
**例子**:
- 管理员撤销邀请
- 操作：`UPDATE tenants.invitations SET status = 'REVOKED', revoked_by = 'admin-uuid', revoked_at = NOW(), revoke_reason = 'no_longer_needed' WHERE invite_id = 'invite-uuid'`
- 结果：邀请失效，用户无法接受

### 4. 过期邀请清理
**场景**: 定期清理过期邀请
**例子**:
- 定时任务清理过期邀请
- 查询：`SELECT invite_id FROM tenants.invitations WHERE expires_at < NOW() AND status = 'PENDING'`
- 操作：`UPDATE tenants.invitations SET status = 'EXPIRED' WHERE expires_at < NOW() AND status = 'PENDING'`
- 结果：过期邀请被标记为已过期

