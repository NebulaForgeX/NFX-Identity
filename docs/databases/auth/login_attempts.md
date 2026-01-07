# Login Attempts Table

## 概述 / Overview

`auth.login_attempts` 表用于记录所有认证尝试事件，用于账户锁定、速率限制、审计和风险评分。

The `auth.login_attempts` table records every authentication attempt for locking, rate limiting, audit, and risk scoring.

## 使用场景 / Use Cases

### 1. 记录登录尝试
**场景**: 每次登录尝试都记录到表中
**例子**:
- 用户"张三"尝试登录，输入错误密码
- 操作：`INSERT INTO auth.login_attempts (tenant_id, identifier, user_id, ip, success, failure_code) VALUES ('tenant-uuid', 'zhangsan@example.com', NULL, '203.0.113.1', false, 'bad_password')`
- 结果：记录失败尝试，用于后续的账户锁定和风险分析

### 2. 账户锁定检测
**场景**: 检测连续失败登录，自动锁定账户
**例子**:
- 系统检测用户在过去 15 分钟内失败 5 次
- 查询：`SELECT COUNT(*) FROM auth.login_attempts WHERE identifier = 'zhangsan@example.com' AND success = false AND created_at >= NOW() - INTERVAL '15 minutes'`
- 如果 >= 5，锁定账户：`INSERT INTO auth.account_lockouts (user_id, tenant_id, locked_until, lock_reason, locked_by) VALUES ('zhangsan-uuid', 'tenant-uuid', NOW() + INTERVAL '30 minutes', 'too_many_attempts', 'system')`

### 3. IP 速率限制
**场景**: 限制同一 IP 的登录尝试频率
**例子**:
- 系统检测到某个 IP 在 1 分钟内尝试登录 10 次
- 查询：`SELECT COUNT(*) FROM auth.login_attempts WHERE ip = '203.0.113.1' AND created_at >= NOW() - INTERVAL '1 minute'`
- 如果 >= 10，拒绝后续请求，返回 `failure_code = 'rate_limited'`

### 4. 风险评分
**场景**: 基于登录尝试历史计算风险评分
**例子**:
- 系统分析用户的登录模式
- 查询：`SELECT COUNT(*) as fail_count, COUNT(DISTINCT ip) as ip_count, COUNT(DISTINCT device_fingerprint) as device_count FROM auth.login_attempts WHERE identifier = 'zhangsan@example.com' AND created_at >= NOW() - INTERVAL '24 hours'`
- 如果失败次数多、IP 多、设备多，风险评分高，要求 MFA 验证

