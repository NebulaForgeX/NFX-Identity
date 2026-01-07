# Event Retention Policies Table

## 概述 / Overview

`audit.event_retention_policies` 表用于定义不同事件类型/分类的保留策略。

支持合规要求：180 天 / 1 年 / 7 年保留。

The `audit.event_retention_policies` table defines retention periods for different event types/classifications.

Supports compliance requirements: 180 days / 1 year / 7 years retention.

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 策略记录的唯一标识符 |
| `policy_name` | VARCHAR(255) | NOT NULL, UNIQUE | 策略名称 |
| `tenant_id` | UUID | NULL | NULL 表示全局策略 |
| `action_pattern` | VARCHAR(255) | NULL | 操作模式："auth.*", "directory.users.export", "*" 表示所有 |
| `data_classification` | `audit.data_classification` ENUM | NULL | 数据分类过滤器 |
| `risk_level` | `audit.risk_level` ENUM | NULL | 风险等级过滤器 |
| `retention_days` | INTEGER | NOT NULL | 保留天数 |
| `retention_action` | `audit.retention_action` ENUM | NOT NULL, DEFAULT 'archive' | 保留操作：archive（归档）、delete（删除）、export（导出） |
| `archive_location` | TEXT | NULL | 归档存储位置（S3 路径等） |
| `status` | VARCHAR(50) | NOT NULL, DEFAULT 'active' | 状态：'active'（活跃）、'disabled'（禁用） |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| `updated_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 更新时间 |
| `created_by` | UUID | NULL | 谁创建了这个策略 |

## 枚举类型 / Enum Types

### `audit.retention_action`
- **`archive`**: 归档到外部存储（如 S3）
- **`delete`**: 删除数据
- **`export`**: 导出数据后删除

### `audit.data_classification`
- **`public`**: 公开数据
- **`internal`**: 内部数据
- **`confidential`**: 机密数据
- **`restricted`**: 受限数据

### `audit.risk_level`
- **`low`**: 低风险
- **`medium`**: 中风险
- **`high`**: 高风险
- **`critical`**: 严重风险

## 字段详解 / Field Details

### `id` (UUID)
- **用途**: 策略记录的主键标识符
- **生成方式**: 自动生成 UUID
- **唯一性**: 全局唯一

### `policy_name` (VARCHAR(255))
- **用途**: 策略的名称
- **示例**: "标准保留策略", "高风险事件保留策略", "合规保留策略"
- **唯一性**: 全局唯一
- **不可为空**: 是

### `tenant_id` (UUID)
- **用途**: 租户级策略（可选）
- **NULL**: 表示全局策略
- **非 NULL**: 表示租户级策略
- **引用**: 引用 `tenants.tenants.id`（应用级一致性）

### `action_pattern` (VARCHAR(255))
- **用途**: 操作模式匹配
- **格式**: 支持通配符模式
  - `"auth.*"`: 匹配所有以 "auth." 开头的操作
  - `"directory.users.export"`: 匹配特定操作
  - `"*"`: 匹配所有操作
- **可为空**: 是（NULL 表示匹配所有操作）

### `data_classification` (`audit.data_classification` ENUM)
- **用途**: 数据分类过滤器
- **可为空**: 是（NULL 表示不按数据分类过滤）

### `risk_level` (`audit.risk_level` ENUM)
- **用途**: 风险等级过滤器
- **可为空**: 是（NULL 表示不按风险等级过滤）

### `retention_days` (INTEGER)
- **用途**: 保留天数
- **示例**: 
  - `180`: 保留 180 天
  - `365`: 保留 1 年
  - `2555`: 保留 7 年
- **不可为空**: 是

### `retention_action` (`audit.retention_action` ENUM)
- **用途**: 保留操作类型
- **archive**: 归档到外部存储（如 S3），然后从主表删除
- **delete**: 直接删除数据
- **export**: 导出数据后删除
- **默认值**: `'archive'`
- **不可为空**: 是

### `archive_location` (TEXT)
- **用途**: 归档存储位置
- **示例**: 
  - `"s3://audit-archive/2024/01/"`
  - `"gs://audit-bucket/archive/"`
- **可为空**: 是（当 `retention_action='delete'` 时通常为空）

### `status` (VARCHAR(50))
- **用途**: 策略状态
- **active**: 策略活跃，会定期执行
- **disabled**: 策略禁用，不会执行
- **默认值**: `'active'`
- **不可为空**: 是

### `created_at` (TIMESTAMP)
- **用途**: 记录策略创建的时间
- **自动设置**: 插入时自动设置为当前时间

### `updated_at` (TIMESTAMP)
- **用途**: 记录策略最后更新的时间
- **自动设置**: 更新时自动设置为当前时间

### `created_by` (UUID)
- **用途**: 记录是谁创建了这个策略
- **引用**: 通常是管理员用户 ID（引用 `directory.users.id`，应用级一致性）
- **可为空**: 是

## 索引 / Indexes

1. **`idx_event_retention_policies_policy_name`**: 在 `policy_name` 字段上创建索引，用于快速查找策略
2. **`idx_event_retention_policies_tenant_id`**: 在 `tenant_id` 字段上创建索引，用于按租户查询
3. **`idx_event_retention_policies_status`**: 在 `status` 字段上创建索引，用于筛选活跃策略

## 策略匹配规则 / Policy Matching Rules

### 匹配优先级
1. **租户级策略**: 优先匹配租户级策略
2. **全局策略**: 如果没有租户级策略，匹配全局策略
3. **最具体策略**: 如果有多个策略匹配，选择最具体的策略（action_pattern 最具体）

### 匹配条件
- **action_pattern**: 使用通配符匹配（如 `"auth.*"` 匹配所有以 "auth." 开头的操作）
- **data_classification**: 精确匹配
- **risk_level**: 精确匹配

## 使用场景 / Use Cases

### 1. GDPR 合规要求
**场景**: 根据 GDPR 要求，用户数据相关事件需要保留特定时间
**例子**:
- GDPR 要求：用户数据导出操作需要保留 1 年
- 创建策略：`INSERT INTO audit.event_retention_policies (policy_name, action_pattern, retention_days, retention_action, archive_location, status) VALUES ('GDPR 用户导出保留', 'users.export', 365, 'archive', 's3://audit-archive/gdpr/', 'active')`
- 系统每天执行：查找 `action = 'users.export'` 且 `occurred_at < NOW() - INTERVAL '365 days'` 的事件，归档到 S3，然后从主表删除
- 结果：主表数据量减少，但合规数据已归档，可以随时恢复

### 2. SOC 2 合规要求
**场景**: SOC 2 要求所有管理员操作保留 7 年
**例子**:
- SOC 2 要求：所有管理员操作需要保留 7 年（2555 天）
- 创建策略：`INSERT INTO audit.event_retention_policies (policy_name, action_pattern, data_classification, retention_days, retention_action, archive_location, status) VALUES ('SOC 2 管理员操作保留', '*', 'confidential', 2555, 'archive', 's3://audit-archive/soc2/', 'active')`
- 系统每月执行：查找所有 `data_classification = 'confidential'` 且 `occurred_at < NOW() - INTERVAL '2555 days'` 的事件，归档到 S3
- 结果：满足 SOC 2 合规要求，数据长期保留

### 3. 存储成本优化
**场景**: 减少主表数据量，降低存储成本
**例子**:
- 主表有 1 亿条记录，存储成本高
- 创建策略：`INSERT INTO audit.event_retention_policies (policy_name, action_pattern, retention_days, retention_action, status) VALUES ('标准事件保留', '*', 180, 'archive', 'active')`
- 系统每天执行：查找所有 `occurred_at < NOW() - INTERVAL '180 days'` 的事件，归档到 S3，然后删除
- 结果：主表数据量从 1 亿减少到 5000 万，存储成本降低 50%

### 4. 高风险事件长期保留
**场景**: 高风险事件需要长期保留，用于安全分析
**例子**:
- 安全团队需要长期分析高风险事件
- 创建策略：`INSERT INTO audit.event_retention_policies (policy_name, action_pattern, risk_level, retention_days, retention_action, archive_location, status) VALUES ('高风险事件保留', '*', 'high', 1095, 'archive', 's3://audit-archive/high-risk/', 'active')`
- 系统每月执行：查找所有 `risk_level = 'high'` 且 `occurred_at < NOW() - INTERVAL '1095 days'` 的事件，归档到 S3
- 结果：高风险事件保留 3 年，用于安全分析和威胁情报

### 5. 低价值事件快速清理
**场景**: 低风险、低价值的查询操作可以快速删除
**例子**:
- 普通查询操作（如 `users.list`）价值低，不需要长期保留
- 创建策略：`INSERT INTO audit.event_retention_policies (policy_name, action_pattern, risk_level, retention_days, retention_action, status) VALUES ('低风险查询保留', 'users.list', 'low', 30, 'delete', 'active')`
- 系统每天执行：查找所有 `action = 'users.list'` 且 `risk_level = 'low'` 且 `occurred_at < NOW() - INTERVAL '30 days'` 的事件，直接删除
- 结果：快速清理低价值数据，减少存储压力

### 6. 租户自定义保留策略
**场景**: 不同租户有不同的合规要求
**例子**:
- 租户 A（金融行业）需要保留 7 年
- 租户 B（普通企业）只需要保留 180 天
- 创建策略：
  - `INSERT INTO audit.event_retention_policies (policy_name, tenant_id, action_pattern, retention_days, retention_action, status) VALUES ('租户A保留策略', 'tenant-a-uuid', '*', 2555, 'archive', 'active')`
  - `INSERT INTO audit.event_retention_policies (policy_name, tenant_id, action_pattern, retention_days, retention_action, status) VALUES ('租户B保留策略', 'tenant-b-uuid', '*', 180, 'archive', 'active')`
- 系统执行时根据 `tenant_id` 匹配对应的策略
- 结果：不同租户有不同的保留期限，满足各自的合规要求

### 7. 数据导出后删除
**场景**: 导出数据用于外部分析后，从主表删除
**例子**:
- 数据分析团队需要导出数据用于 BI 分析
- 创建策略：`INSERT INTO audit.event_retention_policies (policy_name, action_pattern, retention_days, retention_action, archive_location, status) VALUES ('数据分析导出', '*', 90, 'export', 's3://audit-export/bi-analysis/', 'active')`
- 系统执行：查找所有 `occurred_at < NOW() - INTERVAL '90 days'` 的事件，导出为 CSV 格式到 S3，然后从主表删除
- 结果：数据已导出用于分析，主表数据量减少

## 策略执行流程 / Policy Execution Flow

### 1. 策略匹配
```sql
-- 查找匹配的策略
SELECT *
FROM audit.event_retention_policies
WHERE status = 'active'
  AND (tenant_id = 'tenant-uuid' OR tenant_id IS NULL)
  AND (
    action_pattern = '*' OR
    action LIKE action_pattern OR
    action_pattern IS NULL
  )
  AND (data_classification = 'internal' OR data_classification IS NULL)
  AND (risk_level = 'high' OR risk_level IS NULL)
ORDER BY 
  CASE WHEN tenant_id IS NOT NULL THEN 1 ELSE 2 END, -- 租户级优先
  LENGTH(action_pattern) DESC -- 最具体优先
LIMIT 1;
```

### 2. 查找过期事件
```sql
-- 查找需要处理的事件
SELECT event_id, occurred_at, action, data_classification, risk_level
FROM audit.events
WHERE occurred_at < NOW() - INTERVAL '180 days'
  AND action LIKE 'auth.%'
  AND data_classification = 'internal'
  AND risk_level = 'low';
```

### 3. 执行保留操作
- **archive**: 导出到外部存储，然后删除
- **delete**: 直接删除
- **export**: 导出数据，然后删除

## 查询示例 / Query Examples

```sql
-- 查询所有活跃的保留策略
SELECT policy_name, action_pattern, retention_days, retention_action, status
FROM audit.event_retention_policies
WHERE status = 'active'
ORDER BY retention_days DESC;

-- 查询租户的保留策略
SELECT policy_name, action_pattern, retention_days, retention_action
FROM audit.event_retention_policies
WHERE tenant_id = 'tenant-uuid'
  AND status = 'active';

-- 查找需要归档的事件
WITH policy AS (
  SELECT retention_days, retention_action, archive_location
  FROM audit.event_retention_policies
  WHERE policy_name = '标准保留策略'
    AND status = 'active'
  LIMIT 1
)
SELECT e.event_id, e.occurred_at, e.action
FROM audit.events e
CROSS JOIN policy p
WHERE e.occurred_at < NOW() - (p.retention_days || ' days')::INTERVAL
  AND p.retention_action = 'archive';
```

## 策略示例 / Policy Examples

```sql
-- 标准保留策略：所有事件保留 180 天
INSERT INTO audit.event_retention_policies (
  policy_name, action_pattern, retention_days, retention_action, status, created_by
) VALUES (
  '标准保留策略',
  '*',
  180,
  'archive',
  'active',
  'admin-user-id'
);

-- 高风险事件保留策略：高风险事件保留 1 年
INSERT INTO audit.event_retention_policies (
  policy_name, action_pattern, risk_level, retention_days, retention_action, archive_location, status, created_by
) VALUES (
  '高风险事件保留策略',
  '*',
  'high',
  365,
  'archive',
  's3://audit-archive/high-risk/',
  'active',
  'admin-user-id'
);

-- 合规保留策略：机密数据保留 7 年
INSERT INTO audit.event_retention_policies (
  policy_name, action_pattern, data_classification, retention_days, retention_action, archive_location, status, created_by
) VALUES (
  '合规保留策略',
  '*',
  'confidential',
  2555,
  'archive',
  's3://audit-archive/compliance/',
  'active',
  'admin-user-id'
);

-- 认证事件保留策略：认证相关事件保留 90 天
INSERT INTO audit.event_retention_policies (
  policy_name, action_pattern, retention_days, retention_action, status, created_by
) VALUES (
  '认证事件保留策略',
  'auth.*',
  90,
  'delete',
  'active',
  'admin-user-id'
);
```

## 注意事项 / Notes

1. **策略优先级**: 需要明确策略匹配的优先级规则
2. **执行频率**: 建议定期（如每天）执行保留策略
3. **归档验证**: 归档后需要验证数据完整性
4. **策略更新**: 更新策略时需要谨慎，避免影响已有数据
5. **合规要求**: 确保保留策略符合合规要求
6. **存储成本**: 归档到外部存储需要考虑存储成本
7. **数据恢复**: 归档的数据需要能够恢复（如果需要）

