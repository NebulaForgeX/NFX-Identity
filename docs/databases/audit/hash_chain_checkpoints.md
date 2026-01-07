# Hash Chain Checkpoints Table

## 概述 / Overview

`audit.hash_chain_checkpoints` 表用于存储防篡改检查点（Tamper-Evident Checkpoints），防止管理员/DBA 直接修改 `audit.events` 表内容而不留痕迹。

这对于合规/安全事件调查非常重要。

**注意**: 这不是绝对防篡改（DB 管理员仍可修改链），但会显著提高篡改成本。真正更强的保护是 WORM 存储/外部审计仓库，但可以先做轻量版。

The `audit.hash_chain_checkpoints` table stores tamper-evident checkpoints, preventing admins/DBAs from directly modifying `audit.events` content without leaving traces.

This is important for compliance/security incident investigations.

**Note**: This is not absolute tamper-proof (DB admin can still modify chain), but significantly increases tampering cost. True stronger protection: WORM storage/external audit warehouse, but lightweight version can be done first.

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 检查点记录的唯一标识符 |
| `checkpoint_id` | VARCHAR(255) | NOT NULL, UNIQUE | 检查点标识符 |
| `tenant_id` | UUID | NULL | 可选的：租户级检查点（NULL 表示全局） |
| `partition_date` | DATE | NOT NULL | 分区日期（用于每日/每周/每月检查点） |
| `checkpoint_hash` | VARCHAR(64) | NOT NULL | 检查点哈希值：该分区/时间段内所有事件的 SHA256 哈希值 |
| `prev_checkpoint_hash` | VARCHAR(64) | NULL | 前一个检查点的哈希值（链式链接） |
| `event_count` | INTEGER | NOT NULL | 该检查点包含的事件数量 |
| `first_event_id` | VARCHAR(255) | NULL | 该检查点中第一个事件的 ID |
| `last_event_id` | VARCHAR(255) | NULL | 该检查点中最后一个事件的 ID |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| `created_by` | VARCHAR(255) | NULL | 创建此检查点的系统/服务 |

## 字段详解 / Field Details

### `id` (UUID)
- **用途**: 检查点记录的主键标识符
- **生成方式**: 自动生成 UUID
- **唯一性**: 全局唯一

### `checkpoint_id` (VARCHAR(255))
- **用途**: 检查点的唯一标识符
- **命名规范**: 建议使用格式如 `tenant-{tenant_id}-{date}` 或 `global-{date}`
- **唯一性**: 全局唯一
- **用途**: 用于标识和引用检查点

### `tenant_id` (UUID)
- **用途**: 租户级检查点（可选）
- **NULL**: 表示全局检查点
- **非 NULL**: 表示租户级检查点
- **引用**: 引用 `tenants.tenants.id`（应用级一致性）

### `partition_date` (DATE)
- **用途**: 分区日期，用于标识检查点覆盖的时间段
- **格式**: DATE 类型，如 `2024-01-15`
- **用途**: 用于按日期查找和验证检查点
- **不可为空**: 是

### `checkpoint_hash` (VARCHAR(64))
- **用途**: 检查点的哈希值
- **计算方式**: SHA256(该分区/时间段内所有事件的哈希值)
- **格式**: SHA256 哈希值（64 个十六进制字符）
- **用途**: 用于验证该时间段内事件的完整性
- **不可为空**: 是

### `prev_checkpoint_hash` (VARCHAR(64))
- **用途**: 前一个检查点的哈希值（链式链接）
- **格式**: SHA256 哈希值（64 个十六进制字符）
- **用途**: 用于构建防篡改链
- **可为空**: 是（链中的第一个检查点为空）

### `event_count` (INTEGER)
- **用途**: 该检查点包含的事件数量
- **用途**: 用于验证和统计
- **不可为空**: 是

### `first_event_id` (VARCHAR(255))
- **用途**: 该检查点中第一个事件的 ID
- **引用**: 引用 `audit.events.event_id`（应用级一致性）
- **可为空**: 是

### `last_event_id` (VARCHAR(255))
- **用途**: 该检查点中最后一个事件的 ID
- **引用**: 引用 `audit.events.event_id`（应用级一致性）
- **可为空**: 是

### `created_at` (TIMESTAMP)
- **用途**: 记录检查点创建的时间
- **自动设置**: 插入时自动设置为当前时间

### `created_by` (VARCHAR(255))
- **用途**: 创建此检查点的系统/服务
- **示例**: `"audit-service"`, `"checkpoint-worker"`
- **可为空**: 是

## 索引 / Indexes

1. **`idx_hash_chain_checkpoints_checkpoint_id`**: 在 `checkpoint_id` 字段上创建索引，用于快速查找检查点
2. **`idx_hash_chain_checkpoints_tenant_id`**: 在 `tenant_id` 字段上创建索引，用于按租户查询
3. **`idx_hash_chain_checkpoints_partition_date`**: 在 `partition_date` 字段上创建索引（DESC），用于按日期查询
4. **`idx_hash_chain_checkpoints_prev_hash`**: 在 `prev_checkpoint_hash` 字段上创建索引，用于链式验证

## 防篡改链机制 / Tamper-Evident Chain Mechanism

### 哈希链构建
1. **事件哈希**: 每个事件计算 `event_hash = SHA256(prev_hash + canonical_json(event_fields))`
2. **检查点哈希**: 每个检查点计算 `checkpoint_hash = SHA256(该时间段内所有事件的 event_hash)`
3. **链式链接**: 每个检查点的 `prev_checkpoint_hash` 指向前一个检查点的 `checkpoint_hash`

### 验证流程
1. **事件验证**: 验证每个事件的 `event_hash` 是否正确
2. **检查点验证**: 验证每个检查点的 `checkpoint_hash` 是否正确
3. **链式验证**: 验证检查点之间的链式链接是否正确

### 篡改检测
- **事件篡改**: 如果事件被修改，其 `event_hash` 会不匹配
- **检查点篡改**: 如果检查点被修改，其 `checkpoint_hash` 会不匹配
- **链式篡改**: 如果链被破坏，`prev_checkpoint_hash` 会不匹配

## 使用场景 / Use Cases

### 1. 合规审计完整性证明
**场景**: 向审计机构证明审计数据未被篡改
**例子**:
- 某公司需要向 SOC 2 审计机构证明：过去 6 个月的审计数据完整且未被篡改
- 审计机构要求提供完整性证明
- 查询：`SELECT checkpoint_id, partition_date, checkpoint_hash, prev_checkpoint_hash, event_count FROM audit.hash_chain_checkpoints WHERE partition_date >= '2023-07-01' AND partition_date < '2024-01-01' ORDER BY partition_date`
- 结果：提供完整的检查点链，证明数据完整性。如果任何数据被篡改，哈希值会不匹配

### 2. 安全事件调查
**场景**: 怀疑审计数据被篡改，需要验证
**例子**:
- 安全团队怀疑某个管理员删除了自己的操作记录
- 验证检查点链：`WITH checkpoints AS (SELECT checkpoint_id, partition_date, checkpoint_hash, prev_checkpoint_hash, LAG(checkpoint_hash) OVER (ORDER BY partition_date) AS expected_prev_hash FROM audit.hash_chain_checkpoints WHERE partition_date >= '2024-01-15' ORDER BY partition_date) SELECT checkpoint_id, partition_date, CASE WHEN prev_checkpoint_hash = expected_prev_hash THEN 'Valid' ELSE 'Invalid' END AS chain_status FROM checkpoints WHERE prev_checkpoint_hash IS NOT NULL`
- 结果：如果链被破坏，会显示 "Invalid"，证明数据被篡改

### 3. 数据归档验证
**场景**: 归档数据前验证完整性
**例子**:
- 系统需要将 2023 年的审计数据归档到冷存储
- 归档前验证：`SELECT checkpoint_id, checkpoint_hash, event_count, first_event_id, last_event_id FROM audit.hash_chain_checkpoints WHERE partition_date >= '2023-01-01' AND partition_date < '2024-01-01' ORDER BY partition_date`
- 结果：获取所有检查点，验证数据完整性后归档。归档后可以定期验证归档数据的完整性

### 4. 定期完整性检查
**场景**: 定期自动检查数据完整性
**例子**:
- 系统每天自动检查检查点链的完整性
- 检查脚本：`SELECT COUNT(*) as broken_links FROM (SELECT checkpoint_id, prev_checkpoint_hash, LAG(checkpoint_hash) OVER (ORDER BY partition_date) AS expected_prev_hash FROM audit.hash_chain_checkpoints WHERE prev_checkpoint_hash IS NOT NULL) t WHERE prev_checkpoint_hash != expected_prev_hash`
- 结果：如果返回 0，说明链完整；如果返回 > 0，说明链被破坏，触发警报

### 5. 外部审计仓库同步
**场景**: 将检查点同步到外部审计仓库（WORM 存储）
**例子**:
- 公司使用 AWS S3 Glacier（WORM 存储）作为外部审计仓库
- 每天将检查点数据同步到 S3：`SELECT checkpoint_id, partition_date, checkpoint_hash, prev_checkpoint_hash, event_count FROM audit.hash_chain_checkpoints WHERE partition_date = CURRENT_DATE - INTERVAL '1 day'`
- 结果：将检查点数据上传到 S3，作为不可篡改的备份

### 6. 数据恢复验证
**场景**: 从备份恢复数据后验证完整性
**例子**:
- 系统从备份恢复 2024 年 1 月的审计数据
- 恢复后验证：`SELECT checkpoint_id, checkpoint_hash FROM audit.hash_chain_checkpoints WHERE partition_date >= '2024-01-01' AND partition_date < '2024-02-01' ORDER BY partition_date`
- 重新计算该时间段内所有事件的哈希值，与检查点哈希值对比
- 结果：如果匹配，说明恢复成功；如果不匹配，说明数据损坏或篡改

## 检查点创建策略 / Checkpoint Creation Strategy

### 创建频率
- **每日检查点**: 每天创建一个检查点（推荐）
- **每周检查点**: 每周创建一个检查点（适合低频率场景）
- **每月检查点**: 每月创建一个检查点（适合归档场景）

### 创建时机
1. **定时任务**: 通过定时任务（如 cron）定期创建检查点
2. **分区切换**: 在切换分区时创建检查点
3. **归档前**: 在归档数据前创建检查点

### 创建流程
1. **收集事件**: 收集该时间段内的所有事件
2. **计算哈希**: 计算所有事件的哈希值
3. **构建检查点**: 创建检查点记录
4. **链式链接**: 链接到前一个检查点

## 查询示例 / Query Examples

```sql
-- 查找特定日期的检查点
SELECT checkpoint_id, checkpoint_hash, event_count, first_event_id, last_event_id
FROM audit.hash_chain_checkpoints
WHERE partition_date = '2024-01-15'
ORDER BY created_at DESC
LIMIT 1;

-- 查找租户的检查点链
SELECT checkpoint_id, partition_date, checkpoint_hash, prev_checkpoint_hash, event_count
FROM audit.hash_chain_checkpoints
WHERE tenant_id = 'tenant-uuid'
ORDER BY partition_date ASC;

-- 验证检查点链的完整性
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

-- 查找最新的检查点
SELECT checkpoint_id, partition_date, checkpoint_hash, event_count
FROM audit.hash_chain_checkpoints
WHERE tenant_id IS NULL -- 全局检查点
ORDER BY partition_date DESC
LIMIT 1;
```

## 验证脚本示例 / Verification Script Example

```sql
-- 验证特定时间段的事件完整性
WITH events_in_period AS (
  SELECT event_id, event_hash, occurred_at
  FROM audit.events
  WHERE occurred_at >= '2024-01-15'
    AND occurred_at < '2024-01-16'
    AND tenant_id = 'tenant-uuid'
),
checkpoint AS (
  SELECT checkpoint_hash, event_count, first_event_id, last_event_id
  FROM audit.hash_chain_checkpoints
  WHERE partition_date = '2024-01-15'
    AND tenant_id = 'tenant-uuid'
  ORDER BY created_at DESC
  LIMIT 1
)
SELECT 
  COUNT(*) AS actual_event_count,
  c.event_count AS expected_event_count,
  CASE 
    WHEN COUNT(*) = c.event_count THEN 'Match'
    ELSE 'Mismatch'
  END AS count_status
FROM events_in_period e
CROSS JOIN checkpoint c;
```

## 注意事项 / Notes

1. **不是绝对防篡改**: DB 管理员仍可修改链，但会显著提高篡改成本
2. **检查点频率**: 根据业务需求合理设置检查点频率
3. **哈希计算**: 需要确保哈希计算的一致性（canonical JSON）
4. **链式维护**: 需要确保链式链接的正确性
5. **性能考虑**: 检查点创建可能影响性能，建议在低峰期创建
6. **存储成本**: 检查点数据量相对较小，但需要长期保留
7. **外部存储**: 对于更强的保护，可以考虑将检查点存储到外部 WORM 存储

