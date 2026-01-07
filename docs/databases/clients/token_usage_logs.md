# Token Usage Logs Table

## 概述 / Overview

`clients.token_usage_logs` 表用于审计和分析，记录每次令牌使用，用于安全审计、分析和滥用检测。

The `clients.token_usage_logs` table records each token usage for security audit, analytics, and abuse detection.

## 使用场景 / Use Cases

### 1. 记录令牌使用
**场景**: 每次使用服务令牌时记录日志
**例子**:
- 应用使用服务令牌访问 `/api/v1/users` 接口
- 操作：`INSERT INTO clients.token_usage_logs (token_id, app_id, client_id, endpoint, method, ip, status_code, response_time_ms) VALUES ('jwt-jti-claim', 'app-uuid', 'client-id-123', '/api/v1/users', 'GET', '203.0.113.1', 200, 150)`
- 结果：记录令牌使用情况，用于审计和分析

### 2. 分析 API 使用情况
**场景**: 分析应用的 API 使用模式
**例子**:
- 查询应用最常用的 API 端点
- 查询：`SELECT endpoint, COUNT(*) as count, AVG(response_time_ms) as avg_time FROM clients.token_usage_logs WHERE app_id = 'app-uuid' AND created_at >= NOW() - INTERVAL '24 hours' GROUP BY endpoint ORDER BY count DESC`
- 结果：显示各端点的调用次数和平均响应时间

### 3. 检测异常使用
**场景**: 检测异常的令牌使用模式
**例子**:
- 检测某个令牌在短时间内大量调用
- 查询：`SELECT token_id, COUNT(*) as request_count FROM clients.token_usage_logs WHERE created_at >= NOW() - INTERVAL '10 minutes' GROUP BY token_id HAVING COUNT(*) > 1000`
- 结果：发现异常令牌，触发安全警报

### 4. 性能分析
**场景**: 分析 API 性能
**例子**:
- 查询慢接口
- 查询：`SELECT endpoint, AVG(response_time_ms) as avg_time, MAX(response_time_ms) as max_time FROM clients.token_usage_logs WHERE app_id = 'app-uuid' AND created_at >= NOW() - INTERVAL '1 hour' GROUP BY endpoint HAVING AVG(response_time_ms) > 1000`
- 结果：显示平均响应时间超过 1 秒的接口

