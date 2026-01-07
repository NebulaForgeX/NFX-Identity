# Events Table

## 概述 / Overview

`audit.events` 表是审计事件事实表（核心必备），存储每一条审计事件的"事实记录"（不可丢失）。

这是所有审计查询、合规导出、事后取证的源数据。必须能回答：
- 谁做的（actor）
- 对什么做的（target）
- 做了什么（action）
- 在哪个租户/应用下（tenant_id/app_id）
- 结果如何（success/failure + reason）
- 从哪里来的（ip、user_agent、request_id、trace_id）
- 风险与敏感度（risk_level、data_classification）
- 发生时间（occurred_at）

The `audit.events` table is the core audit event fact table, storing every audit event record (cannot be lost).

This is the source data for all audit queries, compliance exports, and post-incident forensics. Must answer:
- Who did it (actor)
- What was done to (target)
- What was done (action)
- Under which tenant/app (tenant_id/app_id)
- What was the result (success/failure + reason)
- Where did it come from (ip, user_agent, request_id, trace_id)
- Risk and sensitivity (risk_level, data_classification)
- When did it occur (occurred_at)

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 事件记录的唯一标识符 |
| `event_id` | VARCHAR(255) | NOT NULL, UNIQUE | 事件标识符（建议使用 ULID，便于排序） |
| `occurred_at` | TIMESTAMP | NOT NULL | 事件发生时间（不是写入时间） |
| `received_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 写入审计数据库的时间（用于排查延迟） |
| `tenant_id` | UUID | NULL | 多租户隔离（引用 tenants.tenants.id，应用级一致性） |
| `app_id` | UUID | NULL | 应用隔离（引用 clients.apps.id，应用级一致性） |
| `actor_type` | `audit.actor_type` ENUM | NOT NULL | 操作者类型：user（用户）、service（服务）、system（系统）、admin（管理员） |
| `actor_id` | UUID | NOT NULL | 操作者 ID：user_id 或 client_id（引用 directory.users.id 或 clients.apps.id，应用级一致性） |
| `actor_tenant_member_id` | UUID | NULL | 可选的：后端成员身份（如果需要区分成员） |
| `action` | VARCHAR(255) | NOT NULL | 操作枚举字符串，例如："user.login", "users.list", "grants.update", "clients.secret.rotate" 等 |
| `target_type` | VARCHAR(100) | NULL | 目标资源类型："user", "tenant", "client", "role", "asset", "token", "export_job" 等 |
| `target_id` | UUID | NULL | 目标资源 ID（可为空，例如"列表查询"） |
| `result` | `audit.result_type` ENUM | NOT NULL | 结果类型：success（成功）、failure（失败）、deny（拒绝）、error（错误） |
| `failure_reason_code` | VARCHAR(100) | NULL | 失败原因代码，例如："INVALID_PASSWORD", "INSUFFICIENT_SCOPE", "RATE_LIMITED" 等 |
| `http_method` | VARCHAR(10) | NULL | HTTP 方法："GET", "POST", "PUT", "DELETE" 等 |
| `http_path` | VARCHAR(500) | NULL | HTTP 路径："/api/v1/users", "/api/v1/users/count" 等 |
| `http_status` | INTEGER | NULL | HTTP 状态码：200, 401, 403, 500 等 |
| `request_id` | VARCHAR(255) | NULL | 请求 ID，用于请求追踪（非常重要） |
| `trace_id` | VARCHAR(255) | NULL | 追踪 ID，用于分布式追踪 |
| `ip` | INET | NULL | IP 地址 |
| `user_agent` | TEXT | NULL | 用户代理字符串 |
| `geo_country` | VARCHAR(10) | NULL | 可选的：国家代码（来自 IP 地理位置） |
| `risk_level` | `audit.risk_level` ENUM | NOT NULL, DEFAULT 'low' | 风险等级：low（低）、medium（中）、high（高）、critical（严重） |
| `data_classification` | `audit.data_classification` ENUM | NOT NULL, DEFAULT 'internal' | 数据分类：public（公开）、internal（内部）、confidential（机密）、restricted（受限） |
| `prev_hash` | VARCHAR(64) | NULL | 前一个事件的哈希值，用于防篡改链 |
| `event_hash` | VARCHAR(64) | NULL | 当前事件的哈希值：SHA256(prev_hash + canonical_json(event_fields)) |
| `metadata` | JSONB | DEFAULT '{}'::jsonb | 扩展字段（不要把关键字段都放在 metadata 中） |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |

## 枚举类型 / Enum Types

### `audit.actor_type`
- **`user`**: 用户操作者
- **`service`**: 服务操作者
- **`system`**: 系统操作者
- **`admin`**: 管理员操作者

### `audit.result_type`
- **`success`**: 操作成功
- **`failure`**: 操作失败
- **`deny`**: 操作被拒绝
- **`error`**: 操作错误

### `audit.risk_level`
- **`low`**: 低风险
- **`medium`**: 中风险
- **`high`**: 高风险
- **`critical`**: 严重风险

### `audit.data_classification`
- **`public`**: 公开数据
- **`internal`**: 内部数据
- **`confidential`**: 机密数据
- **`restricted`**: 受限数据

## 字段详解 / Field Details

### `id` (UUID)
- **用途**: 事件记录的主键标识符
- **生成方式**: 自动生成 UUID
- **唯一性**: 全局唯一

### `event_id` (VARCHAR(255))
- **用途**: 事件的唯一标识符
- **推荐格式**: ULID（便于排序和查询）
- **唯一性**: 全局唯一
- **用途**: 用于事件追踪和去重

### `occurred_at` (TIMESTAMP)
- **用途**: 事件实际发生的时间
- **重要**: 不是写入数据库的时间，而是事件发生的时间
- **用途**: 用于时间序列分析和合规查询
- **不可为空**: 是

### `received_at` (TIMESTAMP)
- **用途**: 事件写入审计数据库的时间
- **用途**: 用于排查延迟问题（occurred_at 与 received_at 的差值）
- **自动设置**: 插入时自动设置为当前时间

### `tenant_id` (UUID)
- **用途**: 多租户隔离
- **引用**: 引用 `tenants.tenants.id`（应用级一致性）
- **可为空**: 是（系统级事件可能没有租户）

### `app_id` (UUID)
- **用途**: 应用隔离
- **引用**: 引用 `clients.apps.id`（应用级一致性）
- **可为空**: 是（非应用级事件可能没有应用）

### `actor_type` (`audit.actor_type` ENUM)
- **用途**: 定义操作者的类型
- **user**: 用户操作者（引用 `directory.users.id`）
- **service**: 服务操作者（引用 `clients.apps.id`）
- **system**: 系统操作者（系统自动操作）
- **admin**: 管理员操作者（管理员用户）
- **不可为空**: 是

### `actor_id` (UUID)
- **用途**: 操作者的 ID
- **引用**: 
  - 当 `actor_type='user'` 时，引用 `directory.users.id`（应用级一致性）
  - 当 `actor_type='service'` 时，引用 `clients.apps.id`（应用级一致性）
- **不可为空**: 是

### `actor_tenant_member_id` (UUID)
- **用途**: 可选的：后端成员身份（如果需要区分成员）
- **引用**: 引用 `tenants.members.member_id`（应用级一致性）
- **可为空**: 是

### `action` (VARCHAR(255))
- **用途**: 操作的枚举字符串
- **命名规范**: 建议使用点分隔的层级结构，如 `resource.action` 或 `resource.action.subaction`
- **示例**: 
  - `user.login`, `user.logout`, `user.password.change`
  - `users.list`, `users.create`, `users.update`, `users.delete`, `users.export`
  - `grants.create`, `grants.update`, `grants.revoke`
  - `clients.secret.rotate`, `clients.api_key.create`
  - `tenants.member.invite`, `tenants.member.remove`
- **不可为空**: 是

### `target_type` (VARCHAR(100))
- **用途**: 目标资源类型
- **示例**: `"user"`, `"tenant"`, `"client"`, `"role"`, `"asset"`, `"token"`, `"export_job"`
- **可为空**: 是（某些操作可能没有目标，如"列表查询"）

### `target_id` (UUID)
- **用途**: 目标资源的 ID
- **可为空**: 是（某些操作可能没有目标，如"列表查询"）
- **使用场景**: 与 `target_type` 配合使用，标识被操作的具体资源

### `result` (`audit.result_type` ENUM)
- **用途**: 定义操作的结果
- **success**: 操作成功
- **failure**: 操作失败（如密码错误）
- **deny**: 操作被拒绝（如权限不足）
- **error**: 操作错误（如系统错误）
- **不可为空**: 是

### `failure_reason_code` (VARCHAR(100))
- **用途**: 失败原因代码
- **示例**: 
  - `"INVALID_PASSWORD"`: 密码错误
  - `"INSUFFICIENT_SCOPE"`: 权限不足
  - `"RATE_LIMITED"`: 速率限制
  - `"ACCOUNT_LOCKED"`: 账户锁定
- **可为空**: 是（成功操作时为空）

### `http_method` (VARCHAR(10))
- **用途**: HTTP 方法
- **示例**: `"GET"`, `"POST"`, `"PUT"`, `"DELETE"`, `"PATCH"`
- **可为空**: 是（非 HTTP 操作时为空）

### `http_path` (VARCHAR(500))
- **用途**: HTTP 路径
- **示例**: `"/api/v1/users"`, `"/api/v1/users/count"`, `"/api/v1/grants"`
- **可为空**: 是（非 HTTP 操作时为空）

### `http_status` (INTEGER)
- **用途**: HTTP 状态码
- **示例**: `200`（成功）、`401`（未授权）、`403`（禁止）、`500`（服务器错误）
- **可为空**: 是（非 HTTP 操作时为空）

### `request_id` (VARCHAR(255))
- **用途**: 请求 ID，用于请求追踪
- **重要性**: 非常重要，用于关联同一请求的所有事件
- **可为空**: 是

### `trace_id` (VARCHAR(255))
- **用途**: 追踪 ID，用于分布式追踪
- **用途**: 用于跨服务的请求追踪
- **可为空**: 是

### `ip` (INET)
- **用途**: IP 地址
- **类型**: PostgreSQL INET 类型，支持 IPv4 和 IPv6
- **用途**: 用于安全分析和地理位置分析
- **可为空**: 是

### `user_agent` (TEXT)
- **用途**: 用户代理字符串
- **用途**: 用于设备类型和浏览器分析
- **可为空**: 是

### `geo_country` (VARCHAR(10))
- **用途**: 国家代码（来自 IP 地理位置）
- **格式**: ISO 3166-1 alpha-2 代码（如 "US", "CN"）
- **可为空**: 是

### `risk_level` (`audit.risk_level` ENUM)
- **用途**: 风险等级
- **low**: 低风险（如普通查询）
- **medium**: 中风险（如修改操作）
- **high**: 高风险（如删除操作、导出数据）
- **critical**: 严重风险（如权限变更、密钥轮换）
- **默认值**: `'low'`
- **不可为空**: 是

### `data_classification` (`audit.data_classification` ENUM)
- **用途**: 数据分类
- **public**: 公开数据
- **internal**: 内部数据
- **confidential**: 机密数据
- **restricted**: 受限数据
- **默认值**: `'internal'`
- **不可为空**: 是

### `prev_hash` (VARCHAR(64))
- **用途**: 前一个事件的哈希值，用于防篡改链
- **格式**: SHA256 哈希值（64 个十六进制字符）
- **用途**: 用于构建防篡改的事件链
- **可为空**: 是（链中的第一个事件为空）

### `event_hash` (VARCHAR(64))
- **用途**: 当前事件的哈希值
- **计算方式**: SHA256(prev_hash + canonical_json(event_fields))
- **格式**: SHA256 哈希值（64 个十六进制字符）
- **用途**: 用于验证事件完整性和防篡改
- **可为空**: 是（如果未启用防篡改链）

### `metadata` (JSONB)
- **用途**: 扩展字段
- **格式**: JSON 对象
- **注意**: 不要把关键字段都放在 metadata 中，关键字段应作为独立列
- **默认值**: `'{}'::jsonb`
- **用途**: 存储额外的上下文信息，如请求参数、响应数据等

### `created_at` (TIMESTAMP)
- **用途**: 记录创建时间
- **自动设置**: 插入时自动设置为当前时间

## 索引 / Indexes

### 核心索引
1. **`idx_events_event_id`**: 在 `event_id` 字段上创建索引，用于快速查找事件
2. **`idx_events_occurred_at`**: 在 `occurred_at` 字段上创建索引，用于时间范围查询
3. **`idx_events_tenant_id`**: 在 `tenant_id` 字段上创建索引，用于租户隔离查询
4. **`idx_events_app_id`**: 在 `app_id` 字段上创建索引，用于应用隔离查询
5. **`idx_events_actor`**: 在 `(actor_type, actor_id)` 上创建索引，用于按操作者查询
6. **`idx_events_action`**: 在 `action` 字段上创建索引，用于按操作类型查询
7. **`idx_events_target`**: 在 `(target_type, target_id)` 上创建索引，用于按目标资源查询
8. **`idx_events_result`**: 在 `result` 字段上创建索引，用于按结果类型查询
9. **`idx_events_risk_level`**: 在 `risk_level` 字段上创建索引，用于按风险等级查询
10. **`idx_events_data_classification`**: 在 `data_classification` 字段上创建索引，用于按数据分类查询

### 复合索引
11. **`idx_events_tenant_occurred`**: 在 `(tenant_id, occurred_at)` 上创建索引，用于租户时间范围查询
12. **`idx_events_actor_occurred`**: 在 `(actor_type, actor_id, occurred_at)` 上创建索引，用于操作者时间序列查询
13. **`idx_events_action_occurred`**: 在 `(action, occurred_at)` 上创建索引，用于操作类型时间序列查询
14. **`idx_events_tenant_action`**: 在 `(tenant_id, action, occurred_at)` 上创建索引，用于租户操作类型查询
15. **`idx_events_request_id`**: 在 `request_id` 字段上创建索引，用于请求追踪
16. **`idx_events_trace_id`**: 在 `trace_id` 字段上创建索引，用于分布式追踪

## 分区策略 / Partitioning Strategy

**重要**: 此表应该按 `occurred_at` 进行分区（按月或按周）。

### 分区示例
```sql
-- 按月分区
CREATE TABLE audit.events (
  -- ... 字段定义 ...
) PARTITION BY RANGE (occurred_at);

-- 创建分区
CREATE TABLE audit.events_2024_01 PARTITION OF audit.events
  FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');
CREATE TABLE audit.events_2024_02 PARTITION OF audit.events
  FOR VALUES FROM ('2024-02-01') TO ('2024-03-01');
-- ... 更多分区 ...
```

### 分区管理
- **归档**: 到期分区可以 detach 并归档到冷存储
- **删除**: 根据保留策略删除过期分区
- **查询优化**: 查询时 PostgreSQL 会自动进行分区裁剪

## 使用场景 / Use Cases

### 1. 合规审计
**场景**: 满足 GDPR、SOC 2、ISO 27001 等合规要求
**例子**: 
- 某公司需要向审计机构证明：在过去 6 个月内，所有用户数据导出操作都有完整记录
- 查询：`SELECT * FROM audit.events WHERE action = 'users.export' AND occurred_at >= NOW() - INTERVAL '6 months'`
- 结果：返回所有导出操作，包括操作者、时间、IP、结果等完整信息

### 2. 安全事件调查
**场景**: 发现异常登录或可疑操作，需要追踪完整事件链
**例子**:
- 安全团队发现某个用户账户在凌晨 2 点从异常 IP（如国外 IP）登录
- 查询：`SELECT * FROM audit.events WHERE actor_id = 'user-uuid' AND action = 'user.login' AND occurred_at >= '2024-01-15 02:00:00'`
- 进一步查询该用户之后的所有操作：`SELECT * FROM audit.events WHERE actor_id = 'user-uuid' AND occurred_at >= '2024-01-15 02:00:00' ORDER BY occurred_at`
- 发现该用户登录后立即导出了大量用户数据，触发安全警报

### 3. 问题排查
**场景**: 用户报告"我的权限被错误撤销了"，需要追踪是谁、什么时候、为什么撤销的
**例子**:
- 用户投诉：昨天还能访问某个资源，今天就不能了
- 查询：`SELECT * FROM audit.events WHERE target_id = 'resource-uuid' AND action LIKE 'grants.%' AND occurred_at >= NOW() - INTERVAL '2 days'`
- 发现：管理员 A 在昨天下午 3 点撤销了该用户的权限，原因是"权限调整"
- 通过 `request_id` 可以追踪到完整的操作链路

### 4. 性能分析
**场景**: 分析系统负载，找出高频操作和慢操作
**例子**:
- 运维团队发现系统响应变慢，需要分析哪些操作最频繁
- 查询：`SELECT action, COUNT(*) as count, AVG(EXTRACT(EPOCH FROM (received_at - occurred_at))) as avg_delay FROM audit.events WHERE occurred_at >= NOW() - INTERVAL '1 hour' GROUP BY action ORDER BY count DESC`
- 发现：`users.list` 操作在 1 小时内执行了 10,000 次，平均延迟 500ms，需要优化

### 5. 用户行为分析
**场景**: 分析用户使用模式，优化产品体验
**例子**:
- 产品团队想了解用户最常用的功能
- 查询：`SELECT action, COUNT(*) as usage_count FROM audit.events WHERE actor_type = 'user' AND occurred_at >= NOW() - INTERVAL '30 days' GROUP BY action ORDER BY usage_count DESC LIMIT 10`
- 发现：用户最常用的操作是 `users.list` 和 `users.detail`，可以考虑优化这些接口

### 6. 异常检测
**场景**: 自动检测异常行为模式
**例子**:
- 系统检测到某个客户端在短时间内大量调用 API
- 查询：`SELECT actor_id, COUNT(*) as request_count FROM audit.events WHERE actor_type = 'service' AND occurred_at >= NOW() - INTERVAL '1 hour' GROUP BY actor_id HAVING COUNT(*) > 1000`
- 发现：某个客户端在 1 小时内发起了 5000 次请求，触发速率限制警报

### 7. 数据泄露调查
**场景**: 发现数据泄露，需要追踪谁访问了敏感数据
**例子**:
- 安全团队发现某个敏感文件被泄露，需要找出所有访问过该文件的用户
- 查询：`SELECT actor_id, occurred_at, ip, user_agent FROM audit.events WHERE target_type = 'asset' AND target_id = 'file-uuid' AND action = 'assets.read' ORDER BY occurred_at`
- 结果：列出所有访问过该文件的用户、时间、IP 等信息，用于进一步调查

### 8. 合规报告生成
**场景**: 定期生成合规报告，证明系统符合安全标准
**例子**:
- 每季度需要生成合规报告，证明所有管理员操作都有记录
- 查询：`SELECT action, COUNT(*) as count, COUNT(DISTINCT actor_id) as unique_admins FROM audit.events WHERE actor_type = 'admin' AND occurred_at >= '2024-01-01' AND occurred_at < '2024-04-01' GROUP BY action`
- 结果：生成报告，显示所有管理员操作的类型、次数、涉及的管理员数量等

## 查询示例 / Query Examples

```sql
-- 查询租户在特定时间范围内的事件
SELECT event_id, action, result, occurred_at, actor_id
FROM audit.events
WHERE tenant_id = 'tenant-uuid'
  AND occurred_at BETWEEN '2024-01-01' AND '2024-01-31'
ORDER BY occurred_at DESC;

-- 查询特定用户的所有失败操作
SELECT event_id, action, failure_reason_code, occurred_at, ip
FROM audit.events
WHERE actor_type = 'user'
  AND actor_id = 'user-uuid'
  AND result = 'failure'
ORDER BY occurred_at DESC;

-- 查询高风险事件
SELECT event_id, action, actor_id, target_type, target_id, occurred_at
FROM audit.events
WHERE risk_level IN ('high', 'critical')
  AND occurred_at >= NOW() - INTERVAL '24 hours'
ORDER BY occurred_at DESC;

-- 通过 request_id 追踪请求
SELECT event_id, action, result, occurred_at, http_method, http_path, http_status
FROM audit.events
WHERE request_id = 'request-uuid'
ORDER BY occurred_at ASC;
```

## 注意事项 / Notes

1. **分区策略**: 此表必须进行分区，否则会严重影响性能
2. **防篡改链**: `prev_hash` 和 `event_hash` 用于构建防篡改链，需要应用层实现
3. **事件 ID**: 建议使用 ULID 作为 `event_id`，便于排序和查询
4. **时间字段**: `occurred_at` 是事件发生时间，`received_at` 是写入时间，两者可能不同
5. **应用级一致性**: `tenant_id`、`app_id`、`actor_id` 等字段使用应用级一致性，不建立数据库外键
6. **元数据字段**: `metadata` 字段用于扩展信息，但关键字段应作为独立列

