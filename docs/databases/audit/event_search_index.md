# Event Search Index Table

## 概述 / Overview

`audit.event_search_index` 表是"面向查询"的冗余索引表，用于在审计管理页面进行快速过滤：按操作者、操作、目标、结果、时间范围等。

解决 `audit.events` 表中 JSONB metadata 过滤慢的问题。允许你将热字段抽出来，甚至做轻量倒排（如果你不接 ELK）。

**注意**: 这是一个"查询加速表"，可以由 pipeline 异步写入。允许比 `events` 表稍微延迟几秒/几分钟（可接受）。

The `audit.event_search_index` table is a "query-optimized" redundant index table for fast filtering in audit admin pages: by actor, action, target, result, time range, etc.

Solves slow filtering on `audit.events` JSONB metadata. Allows extracting hot fields and even lightweight inverted index (if not using ELK).

**Note**: This is a "query acceleration table", can be written asynchronously by pipeline. Allows slight delay (seconds/minutes) compared to events table (acceptable).

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 索引记录的唯一标识符 |
| `event_id` | VARCHAR(255) | NOT NULL, UNIQUE | 事件 ID（引用 audit.events.event_id，应用级一致性） |
| `tenant_id` | UUID | NULL | 多租户隔离 |
| `app_id` | UUID | NULL | 应用隔离 |
| `actor_type` | `audit.actor_type` ENUM | NOT NULL | 操作者类型：user、service、system、admin |
| `actor_id` | UUID | NOT NULL | 操作者 ID |
| `action` | VARCHAR(255) | NOT NULL | 操作枚举字符串 |
| `target_type` | VARCHAR(100) | NULL | 目标资源类型 |
| `target_id` | UUID | NULL | 目标资源 ID |
| `result` | `audit.result_type` ENUM | NOT NULL | 结果类型：success、failure、deny、error |
| `occurred_at` | TIMESTAMP | NOT NULL | 事件发生时间 |
| `ip` | INET | NULL | IP 地址 |
| `risk_level` | `audit.risk_level` ENUM | NOT NULL, DEFAULT 'low' | 风险等级：low、medium、high、critical |
| `data_classification` | `audit.data_classification` ENUM | NOT NULL, DEFAULT 'internal' | 数据分类：public、internal、confidential、restricted |
| `tags` | TEXT[] | NULL | 可选的标签数组，例如：["security", "admin", "export", "sensitive", ...] |
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
- **用途**: 索引记录的主键标识符
- **生成方式**: 自动生成 UUID
- **唯一性**: 全局唯一

### `event_id` (VARCHAR(255))
- **用途**: 关联的事件 ID
- **引用**: 引用 `audit.events.event_id`（应用级一致性）
- **唯一性**: 全局唯一
- **用途**: 用于关联回 `audit.events` 表获取完整事件信息

### `tenant_id` (UUID)
- **用途**: 多租户隔离
- **引用**: 引用 `tenants.tenants.id`（应用级一致性）
- **可为空**: 是

### `app_id` (UUID)
- **用途**: 应用隔离
- **引用**: 引用 `clients.apps.id`（应用级一致性）
- **可为空**: 是

### `actor_type` (`audit.actor_type` ENUM)
- **用途**: 操作者类型
- **不可为空**: 是

### `actor_id` (UUID)
- **用途**: 操作者 ID
- **不可为空**: 是

### `action` (VARCHAR(255))
- **用途**: 操作枚举字符串
- **不可为空**: 是

### `target_type` (VARCHAR(100))
- **用途**: 目标资源类型
- **可为空**: 是

### `target_id` (UUID)
- **用途**: 目标资源 ID
- **可为空**: 是

### `result` (`audit.result_type` ENUM)
- **用途**: 结果类型
- **不可为空**: 是

### `occurred_at` (TIMESTAMP)
- **用途**: 事件发生时间
- **不可为空**: 是

### `ip` (INET)
- **用途**: IP 地址
- **可为空**: 是

### `risk_level` (`audit.risk_level` ENUM)
- **用途**: 风险等级
- **默认值**: `'low'`
- **不可为空**: 是

### `data_classification` (`audit.data_classification` ENUM)
- **用途**: 数据分类
- **默认值**: `'internal'`
- **不可为空**: 是

### `tags` (TEXT[])
- **用途**: 可选的标签数组
- **示例**: `["security", "admin", "export", "sensitive", "compliance"]`
- **用途**: 用于快速分类和过滤事件
- **可为空**: 是

### `created_at` (TIMESTAMP)
- **用途**: 记录创建时间
- **自动设置**: 插入时自动设置为当前时间

## 索引 / Indexes

### 核心索引
1. **`idx_event_search_index_event_id`**: 在 `event_id` 字段上创建索引，用于快速查找事件
2. **`idx_event_search_index_tenant_occurred`**: 在 `(tenant_id, occurred_at)` 上创建索引，用于租户时间范围查询
3. **`idx_event_search_index_actor`**: 在 `(actor_type, actor_id, occurred_at)` 上创建索引，用于操作者时间序列查询
4. **`idx_event_search_index_action`**: 在 `(action, occurred_at)` 上创建索引，用于操作类型时间序列查询
5. **`idx_event_search_index_target`**: 在 `(target_type, target_id)` 上创建索引，用于按目标资源查询
6. **`idx_event_search_index_result`**: 在 `(result, occurred_at)` 上创建索引，用于按结果类型查询
7. **`idx_event_search_index_risk_level`**: 在 `(risk_level, occurred_at)` 上创建索引，用于按风险等级查询
8. **`idx_event_search_index_tags`**: 在 `tags` 字段上创建 GIN 索引，用于数组搜索

## 使用场景 / Use Cases

### 1. 审计管理页面快速查询
**场景**: 管理员在审计管理页面需要快速查看和过滤事件
**例子**:
- 管理员打开审计页面，需要查看"过去 7 天内所有失败的用户登录"
- 使用 `event_search_index` 表查询：`SELECT event_id, action, occurred_at FROM audit.event_search_index WHERE action = 'user.login' AND result = 'failure' AND occurred_at >= NOW() - INTERVAL '7 days' ORDER BY occurred_at DESC`
- 查询速度比直接查询 `events` 表快 10 倍（因为索引表字段更少，索引更优化）

### 2. 标签分类搜索
**场景**: 安全团队需要快速查找所有"安全相关"或"管理员操作"的事件
**例子**:
- 安全团队想查看所有标记为 `["security", "admin"]` 的事件
- 查询：`SELECT event_id, action, occurred_at FROM audit.event_search_index WHERE tags @> ARRAY['security', 'admin'] AND occurred_at >= NOW() - INTERVAL '24 hours'`
- 结果：快速返回所有安全相关的管理员操作，用于安全审查

### 3. 多条件组合查询
**场景**: 需要同时按多个条件过滤事件
**例子**:
- 合规团队需要查看"租户 A 在过去 30 天内所有高风险的数据导出操作"
- 查询：`SELECT event_id, action, actor_id, occurred_at FROM audit.event_search_index WHERE tenant_id = 'tenant-a-uuid' AND action = 'users.export' AND risk_level = 'high' AND occurred_at >= NOW() - INTERVAL '30 days'`
- 结果：快速返回符合条件的 50 条记录，用于合规报告

### 4. 实时监控大屏
**场景**: 运维团队需要实时监控系统状态，显示最近的操作统计
**例子**:
- 监控大屏需要显示"过去 1 小时内各操作类型的统计"
- 查询：`SELECT action, result, COUNT(*) as count FROM audit.event_search_index WHERE occurred_at >= NOW() - INTERVAL '1 hour' GROUP BY action, result ORDER BY count DESC`
- 结果：快速返回统计结果，更新频率高（每 30 秒刷新一次），使用索引表可以避免对主表造成压力

### 5. 用户行为分析
**场景**: 产品团队需要分析用户的操作模式
**例子**:
- 产品团队想了解"哪些用户在过去一周内执行了超过 100 次操作"
- 查询：`SELECT actor_id, COUNT(*) as operation_count FROM audit.event_search_index WHERE actor_type = 'user' AND occurred_at >= NOW() - INTERVAL '7 days' GROUP BY actor_id HAVING COUNT(*) > 100 ORDER BY operation_count DESC`
- 结果：识别出活跃用户，用于产品优化

### 6. 异常检测
**场景**: 自动检测异常操作模式
**例子**:
- 系统需要检测"某个 IP 在短时间内执行了大量操作"
- 查询：`SELECT ip, COUNT(*) as request_count FROM audit.event_search_index WHERE occurred_at >= NOW() - INTERVAL '10 minutes' GROUP BY ip HAVING COUNT(*) > 100`
- 结果：发现异常 IP，触发安全警报

### 7. 性能优化场景
**场景**: 避免在主表上进行复杂查询
**例子**:
- 原本需要在 `events` 表上查询，但 `events` 表有大量字段和 JSONB 数据，查询慢
- 使用 `event_search_index` 表先快速过滤，然后通过 `event_id` 关联回 `events` 表获取完整信息
- 查询性能提升：从 5 秒降低到 0.5 秒

## 数据同步策略 / Data Synchronization Strategy

### 异步写入
- **方式**: 通过消息队列或后台任务异步写入
- **延迟**: 允许几秒到几分钟的延迟（可接受）
- **优势**: 不影响主事件表的写入性能

### 同步写入
- **方式**: 在写入 `audit.events` 时同步写入
- **优势**: 数据一致性更好
- **劣势**: 可能影响写入性能

### 推荐方案
- **实时性要求高**: 同步写入
- **性能要求高**: 异步写入（通过消息队列）

## 查询示例 / Query Examples

```sql
-- 快速查询租户在特定时间范围内的事件
SELECT event_id, action, result, occurred_at
FROM audit.event_search_index
WHERE tenant_id = 'tenant-uuid'
  AND occurred_at BETWEEN '2024-01-01' AND '2024-01-31'
ORDER BY occurred_at DESC
LIMIT 100;

-- 查询特定用户的所有失败操作
SELECT event_id, action, occurred_at
FROM audit.event_search_index
WHERE actor_type = 'user'
  AND actor_id = 'user-uuid'
  AND result = 'failure'
ORDER BY occurred_at DESC;

-- 查询高风险事件
SELECT event_id, action, actor_id, occurred_at
FROM audit.event_search_index
WHERE risk_level IN ('high', 'critical')
  AND occurred_at >= NOW() - INTERVAL '24 hours'
ORDER BY occurred_at DESC;

-- 通过标签搜索
SELECT event_id, action, occurred_at
FROM audit.event_search_index
WHERE tags @> ARRAY['security', 'admin']
ORDER BY occurred_at DESC;

-- 组合查询：特定操作者在特定时间范围内的特定操作
SELECT event_id, result, occurred_at
FROM audit.event_search_index
WHERE actor_type = 'user'
  AND actor_id = 'user-uuid'
  AND action = 'users.export'
  AND occurred_at >= NOW() - INTERVAL '30 days'
ORDER BY occurred_at DESC;
```

## 与 events 表的关系 / Relationship with events Table

### 数据一致性
- **一对一关系**: 每个 `event_id` 对应 `audit.events` 表中的一条记录
- **数据同步**: 需要确保索引表与事件表的数据一致性

### 查询策略
1. **快速过滤**: 使用 `event_search_index` 表进行快速过滤
2. **详细信息**: 通过 `event_id` 关联回 `audit.events` 表获取完整信息

### 示例查询
```sql
-- 先通过索引表快速过滤
WITH filtered_events AS (
  SELECT event_id
  FROM audit.event_search_index
  WHERE tenant_id = 'tenant-uuid'
    AND action = 'users.export'
    AND occurred_at >= NOW() - INTERVAL '7 days'
)
-- 再获取完整事件信息
SELECT e.*
FROM audit.events e
JOIN filtered_events fe ON e.event_id = fe.event_id
ORDER BY e.occurred_at DESC;
```

## 注意事项 / Notes

1. **数据同步**: 需要确保索引表与事件表的数据同步（同步或异步）
2. **延迟容忍**: 允许几秒到几分钟的延迟，但需要监控延迟情况
3. **标签管理**: `tags` 字段需要应用层维护，确保标签的一致性和准确性
4. **索引维护**: GIN 索引对数组搜索很有效，但会增加写入开销
5. **数据清理**: 当 `audit.events` 表的数据被清理时，需要同步清理索引表
6. **性能优化**: 索引表的主要目的是优化查询性能，应定期监控查询性能

