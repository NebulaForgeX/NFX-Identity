# Audit Schema 总览 / Audit Schema Overview

## 概述 / Overview

`audit` schema 是 NFX-Identity 平台的企业级审计（Audit）模块，负责记录所有关键操作、提供合规证明、支持安全调查和问题排查。

企业级审计的核心目标是：
- **可追责（Accountability）**: 谁在什么时间，对什么资源，做了什么动作，结果如何
- **可取证（Forensics）**: 事后能重建事件链：登录、发 token、查用户、改权限、导出数据……全过程
- **可合规（Compliance）**: 能按租户/应用隔离、可导出、可留存（例如 180 天/1 年/7 年）
- **可运营（Operational）**: 能快速回答：过去 24h 哪个租户失败最多？哪个 client 在刷接口？某个管理员是否导出过用户？

The `audit` schema is the enterprise-grade audit module of the NFX-Identity platform, responsible for recording all critical operations, providing compliance proof, supporting security investigations, and troubleshooting.

Enterprise audit core goals:
- **Accountability**: Who did what, to what, when, where, result
- **Forensics**: Reconstruct event chains post-incident: login, token issuance, user queries, permission changes, data exports, etc.
- **Compliance**: Tenant/app isolation, exportable, retainable (e.g., 180 days/1 year/7 years)
- **Operational**: Quickly answer: Which tenant had the most failures in the past 24h? Which client is hitting the API? Did an admin export users?

## 核心概念 / Core Concepts

### 1. 审计事件（Events）
- **定义**: 所有关键操作的记录
- **表**: `audit.events`
- **特点**: 不可丢失的事实记录，所有审计查询的源数据

### 2. 事件搜索索引（Event Search Index）
- **定义**: 面向查询的冗余索引表
- **表**: `audit.event_search_index`
- **特点**: 快速过滤和查询，支持标签搜索

### 3. 操作者快照（Actor Snapshots）
- **定义**: 操作者的历史快照
- **表**: `audit.actor_snapshots`
- **特点**: 即使操作者被删除，审计记录仍然可读

### 4. 哈希链检查点（Hash Chain Checkpoints）
- **定义**: 防篡改检查点
- **表**: `audit.hash_chain_checkpoints`
- **特点**: 防止管理员/DBA 直接修改审计数据而不留痕迹

### 5. 保留策略（Retention Policies）
- **定义**: 不同事件类型的保留策略
- **表**: `audit.event_retention_policies`
- **特点**: 支持合规要求：180 天/1 年/7 年保留

## 表关系图 / Table Relationships

```
┌─────────────────┐
│     events      │
│  (审计事件)     │
└────────┬────────┘
         │
         │ 1:1
         │
    ┌────▼──────────────────────┐
    │  event_search_index       │
    │  (搜索索引)                │
    └───────────────────────────┘

┌─────────────────┐
│  actor_snapshots│
│  (操作者快照)   │
│  (独立表)       │
└─────────────────┘

┌─────────────────┐
│hash_chain_      │
│checkpoints      │
│(防篡改检查点)   │
│(独立表)         │
└─────────────────┘

┌─────────────────┐
│event_retention_ │
│policies         │
│(保留策略)       │
│(独立表)         │
└─────────────────┘
```

## 表列表 / Table List

### 1. `audit.events` - 审计事件表
- **用途**: 存储所有审计事件的"事实记录"
- **关键字段**: `event_id`（事件 ID）、`actor_type`/`actor_id`（操作者）、`action`（操作）、`result`（结果）、`risk_level`（风险等级）
- **特点**: 必须进行分区（按 `occurred_at` 按月或按周）
- **详细文档**: [events.md](./events.md)

### 2. `audit.event_search_index` - 事件搜索索引表
- **用途**: 面向查询的冗余索引表，用于快速过滤
- **关键字段**: `event_id`、`actor_type`/`actor_id`、`action`、`result`、`tags`（标签数组）
- **特点**: 可以异步写入，允许几秒到几分钟的延迟
- **详细文档**: [event_search_index.md](./event_search_index.md)

### 3. `audit.actor_snapshots` - 操作者快照表
- **用途**: 存储操作者的历史快照，用于取证
- **关键字段**: `actor_type`/`actor_id`、`display_name`、`email`、`snapshot_at`、`snapshot_data`
- **特点**: 即使操作者被删除，审计记录仍然可读
- **详细文档**: [actor_snapshots.md](./actor_snapshots.md)

### 4. `audit.hash_chain_checkpoints` - 哈希链检查点表
- **用途**: 存储防篡改检查点，防止管理员/DBA 直接修改审计数据
- **关键字段**: `checkpoint_id`、`partition_date`、`checkpoint_hash`、`prev_checkpoint_hash`、`event_count`
- **特点**: 构建防篡改链，提高篡改成本
- **详细文档**: [hash_chain_checkpoints.md](./hash_chain_checkpoints.md)

### 5. `audit.event_retention_policies` - 事件保留策略表
- **用途**: 定义不同事件类型的保留策略
- **关键字段**: `policy_name`、`action_pattern`、`retention_days`、`retention_action`、`archive_location`
- **特点**: 支持合规要求：180 天/1 年/7 年保留
- **详细文档**: [event_retention_policies.md](./event_retention_policies.md)

## 数据流 / Data Flow

### 1. 事件记录流程
```
1. 操作发生（如用户登录、权限变更）
2. 系统记录事件到 audit.events
   - 设置 occurred_at（事件发生时间）
   - 设置 received_at（写入时间）
   - 计算 event_hash（如果启用防篡改链）
3. 异步写入 event_search_index（可选）
4. 创建/更新 actor_snapshots（如果需要）
```

### 2. 查询流程
```
1. 用户查询审计事件
2. 通过 event_search_index 快速过滤
3. 通过 event_id 关联回 events 表获取完整信息
4. 通过 actor_snapshots 获取操作者信息
```

### 3. 保留策略执行流程
```
1. 定时任务查找匹配的保留策略
2. 查找过期事件（occurred_at < NOW() - retention_days）
3. 执行保留操作：
   - archive: 归档到外部存储，然后删除
   - delete: 直接删除
   - export: 导出数据，然后删除
```

### 4. 防篡改链构建流程
```
1. 每个事件计算 event_hash = SHA256(prev_hash + canonical_json(event_fields))
2. 定期创建检查点（如每天）
3. 计算检查点哈希 = SHA256(该时间段内所有事件的 event_hash)
4. 链接到前一个检查点（prev_checkpoint_hash）
```

## 审计覆盖范围 / Audit Coverage

### 身份安全
- 登录、登出
- 密码变更、重置
- 账户锁定、解锁
- MFA 启用、禁用
- 密码重置请求

### 访问授权
- 授权变更（roles/permissions/grants）
- 角色分配、撤销
- 权限授予、撤销
- 策略创建、更新、删除

### 令牌管理
- Token 签发、撤销
- 异常刷新
- API Key 创建、删除
- 客户端凭证轮换

### 数据访问
- 用户查询、列表、统计
- 数据导出（尤其敏感）
- 敏感信息访问

### 集成方管理
- Client/App 创建、更新、删除
- 密钥轮换
- API Key 使用

### 管理后台
- Tenant 成员增删
- 邀请、接受
- 角色分配
- 设置变更

## 查询示例 / Query Examples

### 查询租户的审计事件
```sql
SELECT 
  e.event_id,
  e.action,
  e.result,
  e.occurred_at,
  COALESCE(a.display_name, 'Unknown') AS actor_name
FROM audit.events e
LEFT JOIN LATERAL (
  SELECT display_name
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

### 查询高风险事件
```sql
SELECT 
  event_id,
  action,
  actor_id,
  target_type,
  target_id,
  occurred_at,
  ip,
  user_agent
FROM audit.events
WHERE risk_level IN ('high', 'critical')
  AND occurred_at >= NOW() - INTERVAL '24 hours'
ORDER BY occurred_at DESC;
```

### 查询失败操作统计
```sql
SELECT 
  action,
  COUNT(*) AS failure_count,
  COUNT(DISTINCT actor_id) AS unique_actors
FROM audit.events
WHERE result = 'failure'
  AND occurred_at >= NOW() - INTERVAL '24 hours'
GROUP BY action
ORDER BY failure_count DESC;
```

### 验证检查点链完整性
```sql
WITH checkpoints AS (
  SELECT 
    checkpoint_id,
    partition_date,
    checkpoint_hash,
    prev_checkpoint_hash,
    LAG(checkpoint_hash) OVER (ORDER BY partition_date) AS expected_prev_hash
  FROM audit.hash_chain_checkpoints
  WHERE tenant_id = 'tenant-uuid'
  ORDER BY partition_date
)
SELECT 
  checkpoint_id,
  partition_date,
  CASE 
    WHEN prev_checkpoint_hash = expected_prev_hash THEN 'Valid'
    ELSE 'Invalid'
  END AS chain_status
FROM checkpoints
WHERE prev_checkpoint_hash IS NOT NULL;
```

## 最佳实践 / Best Practices

1. **分区策略**: `audit.events` 表必须进行分区（按 `occurred_at` 按月或按周）
2. **索引优化**: 为常用查询字段创建合适的索引
3. **异步写入**: `event_search_index` 可以异步写入，提高性能
4. **快照策略**: 定期为活跃操作者创建快照
5. **检查点频率**: 根据业务需求合理设置检查点频率（推荐每天）
6. **保留策略**: 根据合规要求设置合理的保留策略
7. **数据脱敏**: 对敏感信息进行脱敏处理
8. **性能监控**: 定期监控查询性能，优化慢查询

## 合规要求 / Compliance Requirements

### 数据保留
- **标准事件**: 180 天
- **高风险事件**: 1 年
- **机密数据**: 7 年

### 数据完整性
- **防篡改链**: 通过哈希链检查点确保数据完整性
- **快照保留**: 即使操作者被删除，审计记录仍然可读

### 数据导出
- **合规导出**: 支持按租户、时间范围、事件类型导出
- **归档存储**: 支持归档到外部存储（如 S3）

## 相关文档 / Related Documentation

- [events.md](./events.md) - 审计事件表详细文档
- [event_search_index.md](./event_search_index.md) - 事件搜索索引表详细文档
- [actor_snapshots.md](./actor_snapshots.md) - 操作者快照表详细文档
- [hash_chain_checkpoints.md](./hash_chain_checkpoints.md) - 哈希链检查点表详细文档
- [event_retention_policies.md](./event_retention_policies.md) - 事件保留策略表详细文档

