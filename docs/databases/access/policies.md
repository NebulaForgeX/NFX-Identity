# Policies Table

## 概述 / Overview

`access.policies` 表定义了 ABAC（基于属性的访问控制）条件授权策略，用于回答"在什么条件下允许/拒绝访问？"这个问题。

并非所有场景都可以通过简单的 RBAC 覆盖。典型的 ABAC 条件包括：
- "只能读取同一租户的用户"（通常通过 tenant_id claim + 查询条件实现，可能不需要 policy）
- "支持人员只能读取他们负责的用户票据"
- "只能在工作时间访问导出接口"
- "高风险权限需要二次验证/MFA"（通常与验证/风险系统关联）

The `access.policies` table defines ABAC (Attribute-Based Access Control) conditional authorization policies, answering the question "Under what conditions is access allowed/denied?"

Not all scenarios can be covered by simple RBAC. Typical ABAC conditions include:
- "Can only read users in the same tenant" (often implemented via tenant_id claim + query conditions, may not need policy)
- "Support can only read tickets for users they are responsible for"
- "Can only access export interface during working hours"
- "High-risk permissions require secondary verification/MFA" (usually linked with verification/risk)

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 策略的唯一标识符 |
| `name` | VARCHAR(255) | NOT NULL | 策略名称 |
| `description` | TEXT | NULL | 策略的详细描述 |
| `effect` | `access.policy_effect` ENUM | NOT NULL, DEFAULT 'ALLOW' | 策略效果：ALLOW（允许）或 DENY（拒绝） |
| `priority` | INTEGER | NOT NULL, DEFAULT 100 | 冲突时的优先级（数字越小优先级越高） |
| `condition` | JSONB | DEFAULT '{}'::jsonb | 条件表达式/规则，例如：{"tenant_id": "...", "time_range": {"start": "09:00", "end": "18:00"}, "mfa_required": true, ...} |
| `resource_type` | `access.resource_type` ENUM | NULL | 策略适用的资源类型 |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| `updated_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 更新时间 |
| `created_by` | UUID | NULL | 谁创建了这个策略（通常是管理员用户 ID） |
| `deleted_at` | TIMESTAMP | NULL | 软删除时间戳，NULL 表示未删除 |

## 枚举类型 / Enum Types

### `access.policy_effect`
- **`ALLOW`**: 允许访问（白名单策略）
- **`DENY`**: 拒绝访问（黑名单策略）

### `access.resource_type`
- **`user`**: 用户资源
- **`tenant`**: 租户资源
- **`app`**: 应用资源
- **`asset`**: 资产资源
- **`client`**: 客户端资源
- **`other`**: 其他资源类型

## 字段详解 / Field Details

### `id` (UUID)
- **用途**: 策略的主键标识符
- **生成方式**: 自动生成 UUID
- **唯一性**: 全局唯一

### `name` (VARCHAR(255))
- **用途**: 策略的名称，用于管理界面展示和识别
- **示例**: "工作时间访问策略", "高风险操作 MFA 策略", "租户隔离策略"
- **不可为空**: 是

### `description` (TEXT)
- **用途**: 策略的详细描述，说明策略的作用、适用场景和条件
- **可为空**: 是
- **示例**: "限制导出功能只能在工作日 9:00-18:00 使用"

### `effect` (`access.policy_effect` ENUM)
- **用途**: 定义策略的效果
- **ALLOW**: 允许访问，满足条件时允许操作
- **DENY**: 拒绝访问，满足条件时拒绝操作
- **默认值**: `'ALLOW'`
- **使用场景**: 
  - ALLOW: 用于白名单策略，如"工作时间允许导出"
  - DENY: 用于黑名单策略，如"禁止在非工作时间访问敏感操作"

### `priority` (INTEGER)
- **用途**: 当多个策略冲突时的优先级
- **规则**: 数字越小，优先级越高
- **默认值**: `100`
- **使用场景**: 
  - 当同时有 ALLOW 和 DENY 策略时，优先级高的策略生效
  - 例如：优先级 10 的 DENY 策略会覆盖优先级 100 的 ALLOW 策略

### `condition` (JSONB)
- **用途**: 存储策略的条件表达式和规则
- **格式**: JSON 对象，包含各种条件判断
- **常见条件类型**:
  ```json
  {
    "tenant_id": "uuid",                    // 租户 ID
    "time_range": {                         // 时间范围
      "start": "09:00",
      "end": "18:00",
      "weekdays": [1, 2, 3, 4, 5]          // 工作日
    },
    "mfa_required": true,                   // 需要 MFA
    "ip_whitelist": ["192.168.1.0/24"],    // IP 白名单
    "risk_level": ["high", "critical"],     // 风险等级
    "resource_owner": "user_id",            // 资源所有者
    "custom_attributes": {                 // 自定义属性
      "department": "IT",
      "location": "HQ"
    }
  }
  ```
- **默认值**: `'{}'::jsonb`
- **评估**: 应用层需要实现条件评估引擎

### `resource_type` (`access.resource_type` ENUM)
- **用途**: 定义策略适用的资源类型
- **可选值**: `user`, `tenant`, `app`, `asset`, `client`, `other`
- **可为空**: 是（NULL 表示适用于所有资源类型）
- **使用场景**: 用于筛选和匹配策略

### `created_at` (TIMESTAMP)
- **用途**: 记录策略创建的时间
- **自动设置**: 插入时自动设置为当前时间

### `updated_at` (TIMESTAMP)
- **用途**: 记录策略最后更新的时间
- **自动设置**: 更新时自动设置为当前时间

### `created_by` (UUID)
- **用途**: 记录是谁创建了这个策略
- **引用**: 通常是管理员用户 ID（引用 `directory.users.id`，应用级一致性）
- **审计用途**: 用于审计追踪
- **可为空**: 是

### `deleted_at` (TIMESTAMP)
- **用途**: 软删除标记
- **NULL**: 表示策略未被删除
- **非 NULL**: 表示策略已被删除（软删除）
- **查询**: 查询时应过滤 `deleted_at IS NULL` 的记录

## 索引 / Indexes

1. **`idx_policies_name`**: 在 `name` 字段上创建索引，用于快速查找策略
2. **`idx_policies_effect`**: 在 `effect` 字段上创建索引，用于筛选允许/拒绝策略
3. **`idx_policies_priority`**: 在 `priority` 字段上创建索引，用于按优先级排序
4. **`idx_policies_resource_type`**: 在 `resource_type` 字段上创建索引，用于按资源类型筛选
5. **`idx_policies_deleted_at`**: 在 `deleted_at` 字段上创建索引，用于软删除查询优化

## 使用场景 / Use Cases

### 1. 工作时间限制
**场景**: 限制敏感操作只能在工作时间执行
**例子**:
- 公司要求：数据导出功能只能在工作日 9:00-18:00 执行
- 创建策略：`INSERT INTO access.policies (name, description, effect, priority, condition, resource_type, created_by) VALUES ('工作时间导出策略', '限制导出功能只能在工作日 9:00-18:00 使用', 'ALLOW', 50, '{"time_range": {"start": "09:00", "end": "18:00", "weekdays": [1, 2, 3, 4, 5]}}'::jsonb, 'user', 'admin-uuid')`
- 结果：用户在非工作时间尝试导出时，系统检查策略，拒绝访问
- 验证：用户在周六尝试导出，系统返回"此操作只能在工作日 9:00-18:00 执行"

### 2. IP 白名单限制
**场景**: 限制管理操作只能从内网 IP 执行
**例子**:
- 公司要求：所有管理员操作只能从内网 IP（192.168.0.0/16）执行
- 创建策略：`INSERT INTO access.policies (name, description, effect, priority, condition, created_by) VALUES ('内网访问策略', '只允许内网 IP 访问管理接口', 'ALLOW', 30, '{"ip_whitelist": ["192.168.0.0/16", "10.0.0.0/8"]}'::jsonb, 'admin-uuid')`
- 结果：管理员从外网 IP 尝试操作时，系统检查策略，拒绝访问
- 验证：管理员从 203.0.113.1 尝试操作，系统返回"此操作只能从内网 IP 执行"

### 3. MFA 二次验证要求
**场景**: 高风险操作必须进行二次验证
**例子**:
- 公司要求：所有高风险操作（如权限变更、数据导出）必须进行 MFA 验证
- 创建策略：`INSERT INTO access.policies (name, description, effect, priority, condition, created_by) VALUES ('高风险操作 MFA 策略', '高风险操作必须进行二次验证', 'ALLOW', 10, '{"mfa_required": true, "risk_level": ["high", "critical"]}'::jsonb, 'admin-uuid')`
- 结果：用户尝试执行高风险操作时，系统检查策略，要求进行 MFA 验证
- 验证：用户尝试导出数据，系统提示"此操作需要二次验证"，要求输入 MFA 代码

### 4. 资源所有者限制
**场景**: 用户只能访问自己创建的资源
**例子**:
- 公司要求：用户只能修改自己创建的项目
- 创建策略：`INSERT INTO access.policies (name, description, effect, priority, condition, resource_type, created_by) VALUES ('资源所有者策略', '用户只能修改自己创建的资源', 'ALLOW', 20, '{"resource_owner": "{{actor_id}}"}'::jsonb, 'project', 'admin-uuid')`
- 结果：用户尝试修改项目时，系统检查项目的 `owner_id` 是否等于 `actor_id`
- 验证：用户 A 尝试修改用户 B 创建的项目，系统检查策略，拒绝访问

### 5. 部门限制
**场景**: 不同部门只能访问自己部门的数据
**例子**:
- 公司要求：IT 部门只能访问 IT 相关的资源
- 创建策略：`INSERT INTO access.policies (name, description, effect, priority, condition, resource_type, created_by) VALUES ('部门限制策略', 'IT 部门只能访问 IT 相关资源', 'ALLOW', 40, '{"department": "IT", "resource_department": "IT"}'::jsonb, 'resource', 'admin-uuid')`
- 结果：IT 部门用户尝试访问其他部门的资源时，系统检查策略，拒绝访问
- 验证：IT 部门用户尝试访问财务部门的资源，系统返回"您只能访问 IT 部门的资源"

### 6. 地理位置限制
**场景**: 限制某些操作只能从特定地理位置执行
**例子**:
- 公司要求：敏感数据导出只能从公司总部（中国）执行
- 创建策略：`INSERT INTO access.policies (name, description, effect, priority, condition, created_by) VALUES ('地理位置限制策略', '敏感数据导出只能从中国执行', 'ALLOW', 25, '{"geo_country": "CN"}'::jsonb, 'admin-uuid')`
- 结果：用户从其他国家尝试导出敏感数据时，系统检查策略，拒绝访问
- 验证：用户从美国尝试导出数据，系统返回"此操作只能从中国执行"

### 7. 组合条件策略
**场景**: 多个条件组合，必须同时满足
**例子**:
- 公司要求：数据导出必须同时满足：工作日、内网 IP、MFA 验证
- 创建策略：`INSERT INTO access.policies (name, description, effect, priority, condition, created_by) VALUES ('组合条件策略', '数据导出必须同时满足多个条件', 'ALLOW', 5, '{"time_range": {"start": "09:00", "end": "18:00", "weekdays": [1, 2, 3, 4, 5]}, "ip_whitelist": ["192.168.0.0/16"], "mfa_required": true}'::jsonb, 'admin-uuid')`
- 结果：用户尝试导出时，系统检查所有条件，只有全部满足才允许
- 验证：用户在工作日、内网 IP、但未进行 MFA 验证，系统拒绝访问

## 策略评估流程 / Policy Evaluation Flow

1. **收集上下文**: 收集请求的上下文信息（用户、资源、时间、IP 等）
2. **匹配策略**: 根据 `resource_type` 和其他条件匹配相关策略
3. **评估条件**: 对每个策略的 `condition` 字段进行评估
4. **应用优先级**: 按 `priority` 排序，优先级高的策略优先应用
5. **决定结果**: 根据 `effect` 字段决定允许或拒绝

## 查询示例 / Query Examples

```sql
-- 查询所有活跃的策略
SELECT id, name, effect, priority, resource_type
FROM access.policies
WHERE deleted_at IS NULL
ORDER BY priority ASC;

-- 查询适用于特定资源类型的策略
SELECT id, name, effect, condition
FROM access.policies
WHERE resource_type = 'user' AND deleted_at IS NULL
ORDER BY priority ASC;

-- 查询拒绝类策略
SELECT id, name, description, condition
FROM access.policies
WHERE effect = 'DENY' AND deleted_at IS NULL
ORDER BY priority ASC;
```

## 示例数据 / Example Data

```sql
-- 工作时间访问策略
INSERT INTO access.policies (name, description, effect, priority, condition, resource_type, created_by) VALUES
  (
    '工作时间导出策略',
    '限制导出功能只能在工作日 9:00-18:00 使用',
    'ALLOW',
    50,
    '{"time_range": {"start": "09:00", "end": "18:00", "weekdays": [1, 2, 3, 4, 5]}}'::jsonb,
    'user',
    'admin-user-id'
  );

-- MFA 要求策略
INSERT INTO access.policies (name, description, effect, priority, condition, resource_type, created_by) VALUES
  (
    '高风险操作 MFA 策略',
    '高风险操作必须进行二次验证',
    'ALLOW',
    10,
    '{"mfa_required": true, "risk_level": ["high", "critical"]}'::jsonb,
    'other',
    'admin-user-id'
  );

-- IP 白名单策略
INSERT INTO access.policies (name, description, effect, priority, condition, resource_type, created_by) VALUES
  (
    '内网访问策略',
    '只允许内网 IP 访问管理接口',
    'ALLOW',
    30,
    '{"ip_whitelist": ["192.168.0.0/16", "10.0.0.0/8"]}'::jsonb,
    'other',
    'admin-user-id'
  );
```

## 注意事项 / Notes

1. **条件评估**: `condition` 字段的 JSONB 结构需要应用层实现评估引擎
2. **优先级管理**: `priority` 字段的值需要合理规划，避免冲突
3. **性能考虑**: 策略评估可能影响性能，建议缓存常用策略
4. **软删除**: 使用软删除机制，保留历史数据用于审计
5. **策略冲突**: 当多个策略冲突时，按优先级和 effect 决定最终结果
6. **条件复杂度**: `condition` 字段支持复杂的 JSON 结构，但评估逻辑需要在应用层实现

