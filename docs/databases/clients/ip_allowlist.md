# IP Allowlist Table

## 概述 / Overview

`clients.ip_allowlist` 表用于源 IP 约束，限制使用该客户端/API Key 的请求必须来自指定的 IP/CIDR。

降低凭证泄露的风险。企业常见：生产环境只允许固定的出口 IP。

The `clients.ip_allowlist` table constrains requests using this client/API key must come from specified IPs/CIDRs.

Reduces risk of credential leakage. Enterprise common: production only allows fixed egress IPs.

## 使用场景 / Use Cases

### 1. 配置 IP 白名单
**场景**: 为应用配置允许的 IP 范围
**例子**:
- 为"数据分析服务"配置生产服务器 IP 白名单
- 操作：`INSERT INTO clients.ip_allowlist (rule_id, app_id, cidr, description, status, created_by) VALUES ('rule-uuid', 'app-uuid', '203.0.113.0/24', '生产服务器 IP 段', 'active', 'admin-uuid')`
- 结果：只有来自该 IP 段的请求才能使用该应用的凭证

### 2. IP 验证
**场景**: 应用请求时验证 IP 是否在白名单中
**例子**:
- 应用从 IP `203.0.113.50` 发送请求
- 查询：`SELECT COUNT(*) FROM clients.ip_allowlist WHERE app_id = 'app-uuid' AND status = 'active' AND revoked_at IS NULL AND '203.0.113.50'::inet <<= cidr`
- 如果返回 0，拒绝请求；如果 > 0，允许请求

### 3. 多 IP 段配置
**场景**: 为应用配置多个 IP 段
**例子**:
- 应用需要从办公室网络和云服务器访问
- 操作：
  - `INSERT INTO clients.ip_allowlist (rule_id, app_id, cidr, description) VALUES ('rule-1', 'app-uuid', '192.168.1.0/24', '办公室网络')`
  - `INSERT INTO clients.ip_allowlist (rule_id, app_id, cidr, description) VALUES ('rule-2', 'app-uuid', '203.0.113.0/24', '云服务器')`
- 结果：两个 IP 段都可以访问

### 4. 撤销 IP 规则
**场景**: 撤销某个 IP 规则
**例子**:
- 管理员撤销某个 IP 规则
- 操作：`UPDATE clients.ip_allowlist SET status = 'revoked', revoked_at = NOW(), revoked_by = 'admin-uuid', revoke_reason = 'ip_changed' WHERE rule_id = 'rule-uuid'`
- 结果：该 IP 段不再允许访问

