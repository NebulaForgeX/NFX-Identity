# Rate Limits Table

## 概述 / Overview

`clients.rate_limits` 表用于客户端级别的速率限制，控制每个应用/客户端的请求速率，防止滥用。

The `clients.rate_limits` table controls request rate per app/client to prevent abuse.

## 使用场景 / Use Cases

### 1. 配置速率限制
**场景**: 为应用配置请求速率限制
**例子**:
- 为"数据分析服务"配置每分钟最多 100 次请求
- 操作：`INSERT INTO clients.rate_limits (app_id, limit_type, limit_value, window_seconds, description, created_by) VALUES ('app-uuid', 'requests_per_minute', 100, 60, '数据分析服务速率限制', 'admin-uuid')`
- 结果：应用每分钟最多只能发送 100 次请求

### 2. 速率限制检查
**场景**: 应用请求时检查是否超过速率限制
**例子**:
- 应用发送请求
- 查询速率限制：`SELECT limit_value, window_seconds FROM clients.rate_limits WHERE app_id = 'app-uuid' AND limit_type = 'requests_per_minute' AND status = 'active'`
- 查询该应用在时间窗口内的请求数（使用 Redis 或数据库计数）
- 如果超过限制，拒绝请求，返回 429 Too Many Requests

### 3. 不同时间窗口的速率限制
**场景**: 为应用配置多个时间窗口的速率限制
**例子**:
- 配置每秒 10 次、每分钟 100 次、每小时 1000 次
- 操作：
  - `INSERT INTO clients.rate_limits (app_id, limit_type, limit_value, window_seconds) VALUES ('app-uuid', 'requests_per_second', 10, 1)`
  - `INSERT INTO clients.rate_limits (app_id, limit_type, limit_value, window_seconds) VALUES ('app-uuid', 'requests_per_minute', 100, 60)`
  - `INSERT INTO clients.rate_limits (app_id, limit_type, limit_value, window_seconds) VALUES ('app-uuid', 'requests_per_hour', 1000, 3600)`
- 结果：应用需要同时满足所有速率限制

