# Actor Snapshots Table

## 概述 / Overview

`audit.actor_snapshots` 表用于存储操作者（Actor）的快照，用于取证（Forensics）。

当用户/客户端被删除、重命名或租户迁移时，审计记录仍然可读、可取证。避免在每次审计查询时跨服务反查"这个操作者当时是谁"。

**企业要求**: 审计的目标是"可取证"，取证不能依赖外部数据还存在。

The `audit.actor_snapshots` table stores actor snapshots for forensics.

When user/client is deleted, renamed, or tenant migrated, audit records remain readable and forensically valid. Avoids cross-service lookups for "who was this actor at that time" in every audit query.

**Enterprise requirement**: Audit goal is "forensic evidence", evidence cannot depend on external data still existing.

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 快照记录的唯一标识符 |
| `actor_type` | `audit.actor_type` ENUM | NOT NULL | 操作者类型：user、service、system、admin |
| `actor_id` | UUID | NOT NULL | 操作者 ID |
| `display_name` | VARCHAR(255) | NULL | 当时的显示名称 |
| `email` | VARCHAR(255) | NULL | 当时的邮箱（可能脱敏） |
| `client_name` | VARCHAR(255) | NULL | 客户端/应用名称（用于服务操作者） |
| `tenant_id` | UUID | NULL | 当时的租户（引用 tenants.tenants.id，应用级一致性） |
| `snapshot_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 快照拍摄时间 |
| `snapshot_data` | JSONB | DEFAULT '{}'::jsonb | 扩展快照数据（可能脱敏） |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |

## 唯一约束 / Unique Constraints

- **`(actor_type, actor_id, snapshot_at)`**: 确保同一操作者在同一时间点只有一个快照

## 枚举类型 / Enum Types

### `audit.actor_type`
- **`user`**: 用户操作者
- **`service`**: 服务操作者
- **`system`**: 系统操作者
- **`admin`**: 管理员操作者

## 字段详解 / Field Details

### `id` (UUID)
- **用途**: 快照记录的主键标识符
- **生成方式**: 自动生成 UUID
- **唯一性**: 全局唯一

### `actor_type` (`audit.actor_type` ENUM)
- **用途**: 操作者类型
- **不可为空**: 是

### `actor_id` (UUID)
- **用途**: 操作者 ID
- **不可为空**: 是

### `display_name` (VARCHAR(255))
- **用途**: 操作者当时的显示名称
- **示例**: 
  - 用户: "John Doe", "张三"
  - 服务: "API Gateway", "Background Worker"
- **可为空**: 是

### `email` (VARCHAR(255))
- **用途**: 操作者当时的邮箱地址
- **脱敏**: 可能被脱敏（如 "user@example.com" → "u***@example.com"）
- **可为空**: 是

### `client_name` (VARCHAR(255))
- **用途**: 客户端/应用名称（用于服务操作者）
- **示例**: "Mobile App", "Web Dashboard", "Integration Service"
- **可为空**: 是

### `tenant_id` (UUID)
- **用途**: 操作者当时的租户
- **引用**: 引用 `tenants.tenants.id`（应用级一致性）
- **可为空**: 是

### `snapshot_at` (TIMESTAMP)
- **用途**: 快照拍摄时间
- **默认值**: 当前时间
- **不可为空**: 是
- **用途**: 用于查找特定时间点的快照

### `snapshot_data` (JSONB)
- **用途**: 扩展快照数据
- **格式**: JSON 对象
- **示例**:
  ```json
  {
    "username": "johndoe",
    "role": "admin",
    "department": "IT",
    "location": "HQ",
    "last_login": "2024-01-15T10:30:00Z"
  }
  ```
- **脱敏**: 可能包含脱敏数据
- **默认值**: `'{}'::jsonb`

### `created_at` (TIMESTAMP)
- **用途**: 记录创建时间
- **自动设置**: 插入时自动设置为当前时间

## 索引 / Indexes

1. **`idx_actor_snapshots_actor`**: 在 `(actor_type, actor_id, snapshot_at DESC)` 上创建索引，用于查找操作者的快照历史
2. **`idx_actor_snapshots_tenant_id`**: 在 `tenant_id` 字段上创建索引，用于按租户查询
3. **`idx_actor_snapshots_snapshot_at`**: 在 `snapshot_at` 字段上创建索引，用于按时间查询

## 使用场景 / Use Cases

### 1. 已删除用户的审计查询
**场景**: 用户被删除后，审计记录仍然需要可读
**例子**:
- 某公司员工张三在 2024 年 1 月离职，账户被删除
- 3 个月后，合规审计需要查看张三在 2023 年 12 月的所有操作
- 查询：`SELECT e.event_id, e.action, e.occurred_at, a.display_name, a.email FROM audit.events e LEFT JOIN audit.actor_snapshots a ON e.actor_type = a.actor_type AND e.actor_id = a.actor_id AND a.snapshot_at <= e.occurred_at WHERE e.actor_id = 'zhangsan-uuid' AND e.occurred_at >= '2023-12-01'`
- 结果：即使张三的账户已删除，仍然可以显示"张三 (zhangsan@company.com)"，而不是只显示 UUID

### 2. 用户重命名后的历史记录
**场景**: 用户改名后，历史审计记录仍然显示当时的名字
**例子**:
- 用户"李四"在 2024 年 1 月改名为"李五"
- 查询 2023 年的审计记录时，应该显示"李四"（当时的名字）
- 系统在改名时创建快照：`INSERT INTO audit.actor_snapshots (actor_type, actor_id, display_name, email, snapshot_at) VALUES ('user', 'lisi-uuid', '李四', 'lisi@company.com', '2023-12-31 23:59:59')`
- 查询时自动匹配到改名前的快照，显示正确的历史名字

### 3. 租户迁移后的审计追踪
**场景**: 用户从一个租户迁移到另一个租户后，需要追踪历史操作
**例子**:
- 用户"王五"从"租户 A"迁移到"租户 B"
- 迁移时创建快照：`INSERT INTO audit.actor_snapshots (actor_type, actor_id, display_name, tenant_id, snapshot_at) VALUES ('user', 'wangwu-uuid', '王五', 'tenant-a-uuid', '2024-01-15 10:00:00')`
- 查询迁移前的操作时，可以显示"王五 (租户 A)"，而不是只显示 UUID

### 4. 合规报告生成
**场景**: 生成合规报告时，需要显示操作者的可读信息
**例子**:
- 合规团队需要生成"过去 6 个月所有管理员操作的报告"
- 查询：`SELECT e.action, e.occurred_at, a.display_name, a.email FROM audit.events e LEFT JOIN audit.actor_snapshots a ON e.actor_type = a.actor_type AND e.actor_id = a.actor_id AND a.snapshot_at <= e.occurred_at WHERE e.actor_type = 'admin' AND e.occurred_at >= NOW() - INTERVAL '6 months'`
- 结果：报告显示"管理员：张三 (zhangsan@company.com)"，而不是只显示 UUID，提高报告可读性

### 5. 安全事件调查
**场景**: 调查安全事件时，需要了解操作者的完整信息
**例子**:
- 安全团队发现某个账户在异常时间执行了敏感操作
- 查询：`SELECT a.display_name, a.email, a.snapshot_data FROM audit.actor_snapshots a WHERE a.actor_id = 'suspicious-user-uuid' AND a.snapshot_at <= '2024-01-15 02:00:00' ORDER BY a.snapshot_at DESC LIMIT 1`
- 结果：获取该用户在当时的信息（姓名、邮箱、部门等），用于进一步调查

### 6. 客户端应用变更追踪
**场景**: 服务客户端改名或删除后，历史记录仍然可读
**例子**:
- 某个服务客户端"旧版 API Gateway"被重命名为"新版 API Gateway"
- 查询历史操作时，应该显示当时的名字"旧版 API Gateway"
- 系统在重命名时创建快照：`INSERT INTO audit.actor_snapshots (actor_type, actor_id, client_name, snapshot_at) VALUES ('service', 'api-gateway-uuid', '旧版 API Gateway', '2024-01-01 00:00:00')`
- 查询时自动匹配到重命名前的快照，显示正确的历史名字

## 快照策略 / Snapshot Strategy

### 何时创建快照
1. **事件发生时**: 在写入审计事件时，如果操作者信息发生变化，创建快照
2. **定期快照**: 定期（如每天）为活跃操作者创建快照
3. **变更时快照**: 当操作者信息发生重要变更时（如重命名、邮箱变更）创建快照
4. **删除前快照**: 在删除操作者之前创建最终快照

### 快照频率
- **高频操作者**: 每天或每周创建快照
- **低频操作者**: 每月或按需创建快照
- **变更时**: 立即创建快照

## 查询示例 / Query Examples

```sql
-- 查找操作者在特定时间点的快照
SELECT display_name, email, client_name, snapshot_data
FROM audit.actor_snapshots
WHERE actor_type = 'user'
  AND actor_id = 'user-uuid'
  AND snapshot_at <= '2024-01-15 10:30:00'
ORDER BY snapshot_at DESC
LIMIT 1;

-- 查找操作者的所有快照历史
SELECT snapshot_at, display_name, email, snapshot_data
FROM audit.actor_snapshots
WHERE actor_type = 'user'
  AND actor_id = 'user-uuid'
ORDER BY snapshot_at DESC;

-- 查找租户内所有操作者的快照
SELECT actor_type, actor_id, display_name, snapshot_at
FROM audit.actor_snapshots
WHERE tenant_id = 'tenant-uuid'
  AND snapshot_at >= NOW() - INTERVAL '30 days'
ORDER BY snapshot_at DESC;

-- 在审计查询中使用快照
SELECT 
  e.event_id,
  e.action,
  e.occurred_at,
  COALESCE(a.display_name, 'Unknown') AS actor_name,
  a.email AS actor_email
FROM audit.events e
LEFT JOIN LATERAL (
  SELECT display_name, email
  FROM audit.actor_snapshots
  WHERE actor_type = e.actor_type
    AND actor_id = e.actor_id
    AND snapshot_at <= e.occurred_at
  ORDER BY snapshot_at DESC
  LIMIT 1
) a ON true
WHERE e.tenant_id = 'tenant-uuid'
  AND e.occurred_at >= NOW() - INTERVAL '7 days'
ORDER BY e.occurred_at DESC;
```

## 与 events 表的关系 / Relationship with events Table

### 数据关联
- **通过 actor_type 和 actor_id**: 关联 `audit.events` 表中的操作者
- **时间匹配**: 使用 `snapshot_at <= occurred_at` 查找事件发生时的快照

### 查询优化
- **LATERAL JOIN**: 使用 LATERAL JOIN 在查询事件时同时获取快照
- **索引优化**: 通过索引快速查找相关快照

## 数据脱敏 / Data Masking

### 脱敏策略
1. **邮箱脱敏**: `user@example.com` → `u***@example.com`
2. **姓名脱敏**: `John Doe` → `J*** D***`
3. **敏感信息**: 在 `snapshot_data` 中标记敏感字段

### 脱敏时机
- **创建快照时**: 在创建快照时进行脱敏
- **查询时**: 在查询时根据权限进行脱敏

## 注意事项 / Notes

1. **数据保留**: 快照数据应长期保留，即使操作者被删除
2. **唯一性约束**: 同一操作者在同一时间点只能有一个快照
3. **时间匹配**: 查询时使用 `snapshot_at <= occurred_at` 查找事件发生时的快照
4. **性能优化**: 使用索引优化快照查询性能
5. **数据脱敏**: 根据合规要求对敏感信息进行脱敏
6. **快照频率**: 根据业务需求合理设置快照频率，平衡存储成本和查询需求

