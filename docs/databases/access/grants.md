# Grants Table

## 概述 / Overview

`access.grants` 表是核心授权表，定义了"谁在什么范围内被授予了什么"。

企业必须能够表达：
- **subject**: 用户或服务客户端
- **scope**: 在租户内、在应用内，甚至是在特定资源上（资源级）
- **grant**: 授予角色或直接授予权限（两者都必须支持，企业通常需要异常授权）

The `access.grants` table is the core authorization table, defining "who is granted what in what scope."

Enterprise must be able to express:
- **subject**: user or service_client
- **scope**: in tenant, in app, or even on specific resources (resource-level)
- **grant**: grant role or directly grant permission (both must be supported, enterprises often need exception grants)

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 授权的唯一标识符 |
| `subject_type` | `access.subject_type` ENUM | NOT NULL | 授权主体类型：USER（用户）或 CLIENT（客户端） |
| `subject_id` | UUID | NOT NULL | 主体 ID：user_id 或 client_id（引用 directory.users.id 或 clients.apps.id，应用级一致性） |
| `grant_type` | `access.grant_type` ENUM | NOT NULL | 授权类型：ROLE（角色）或 PERMISSION（权限） |
| `grant_ref_id` | UUID | NOT NULL | 授权引用 ID：role_id 或 permission_id（引用 access.roles.id 或 access.permissions.id，应用级一致性） |
| `tenant_id` | UUID | NULL | 租户 ID，NULL 表示全局授权 |
| `app_id` | UUID | NULL | 应用 ID，NULL 表示非应用级授权 |
| `resource_type` | VARCHAR(100) | NULL | 资源类型："user", "tenant", "app", "asset" 等 |
| `resource_id` | UUID | NULL | 特定资源 ID（用于资源级授权） |
| `effect` | `access.grant_effect` ENUM | NOT NULL, DEFAULT 'ALLOW' | 授权效果：ALLOW（允许）或 DENY（拒绝） |
| `expires_at` | TIMESTAMP | NULL | 临时授权的过期时间，NULL 表示永久授权 |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| `created_by` | UUID | NULL | 谁授予了这个授权（管理员 user_id） |
| `revoked_at` | TIMESTAMP | NULL | 撤销时间，NULL 表示未撤销 |
| `revoked_by` | UUID | NULL | 谁撤销了这个授权 |
| `revoke_reason` | TEXT | NULL | 撤销原因 |

## 枚举类型 / Enum Types

### `access.subject_type`
- **`USER`**: 用户主体，授权给用户
- **`CLIENT`**: 客户端主体，授权给服务客户端

### `access.grant_type`
- **`ROLE`**: 授予角色，主体获得角色包含的所有权限
- **`PERMISSION`**: 直接授予权限，主体获得特定权限

### `access.grant_effect`
- **`ALLOW`**: 允许访问（白名单授权）
- **`DENY`**: 拒绝访问（黑名单授权）

## 字段详解 / Field Details

### `id` (UUID)
- **用途**: 授权的主键标识符
- **生成方式**: 自动生成 UUID
- **唯一性**: 全局唯一

### `subject_type` (`access.subject_type` ENUM)
- **用途**: 定义授权主体的类型
- **USER**: 授权给用户（引用 `directory.users.id`）
- **CLIENT**: 授权给服务客户端（引用 `clients.apps.id`）
- **不可为空**: 是

### `subject_id` (UUID)
- **用途**: 授权主体的 ID
- **引用**: 
  - 当 `subject_type='USER'` 时，引用 `directory.users.id`（应用级一致性）
  - 当 `subject_type='CLIENT'` 时，引用 `clients.apps.id`（应用级一致性）
- **不可为空**: 是

### `grant_type` (`access.grant_type` ENUM)
- **用途**: 定义授权的类型
- **ROLE**: 授予角色，主体获得角色包含的所有权限（通过 `role_permissions` 表查找）
- **PERMISSION**: 直接授予权限，主体获得特定权限
- **不可为空**: 是

### `grant_ref_id` (UUID)
- **用途**: 授权引用的 ID
- **引用**: 
  - 当 `grant_type='ROLE'` 时，引用 `access.roles.id`（应用级一致性）
  - 当 `grant_type='PERMISSION'` 时，引用 `access.permissions.id`（应用级一致性）
- **不可为空**: 是

### `tenant_id` (UUID)
- **用途**: 定义授权的租户范围
- **NULL**: 表示全局授权，跨租户生效
- **非 NULL**: 表示租户级授权，仅在特定租户内生效
- **引用**: 引用 `tenants.tenants.id`（应用级一致性）

### `app_id` (UUID)
- **用途**: 定义授权的应用范围
- **NULL**: 表示非应用级授权
- **非 NULL**: 表示应用级授权，仅在特定应用内生效
- **引用**: 引用 `clients.apps.id`（应用级一致性）

### `resource_type` (VARCHAR(100))
- **用途**: 定义资源级授权的资源类型
- **示例**: `"user"`, `"tenant"`, `"app"`, `"asset"`, `"client"`
- **可为空**: 是（NULL 表示非资源级授权）
- **使用场景**: 用于细粒度的资源级授权，如"用户 A 只能管理用户 B"

### `resource_id` (UUID)
- **用途**: 特定资源的 ID（用于资源级授权）
- **可为空**: 是（NULL 表示非资源级授权）
- **使用场景**: 与 `resource_type` 配合使用，实现资源级授权

### `effect` (`access.grant_effect` ENUM)
- **用途**: 定义授权的效果
- **ALLOW**: 允许访问（白名单授权）
- **DENY**: 拒绝访问（黑名单授权，优先级通常高于 ALLOW）
- **默认值**: `'ALLOW'`

### `expires_at` (TIMESTAMP)
- **用途**: 临时授权的过期时间
- **NULL**: 表示永久授权
- **非 NULL**: 表示临时授权，过期后自动失效
- **使用场景**: 临时访问、临时角色分配等

### `created_at` (TIMESTAMP)
- **用途**: 记录授权创建的时间
- **自动设置**: 插入时自动设置为当前时间

### `created_by` (UUID)
- **用途**: 记录是谁授予了这个授权
- **引用**: 通常是管理员用户 ID（引用 `directory.users.id`，应用级一致性）
- **审计用途**: 用于审计追踪
- **可为空**: 是

### `revoked_at` (TIMESTAMP)
- **用途**: 记录授权撤销的时间
- **NULL**: 表示授权未被撤销
- **非 NULL**: 表示授权已被撤销
- **查询**: 查询有效授权时应过滤 `revoked_at IS NULL` 的记录

### `revoked_by` (UUID)
- **用途**: 记录是谁撤销了这个授权
- **引用**: 通常是管理员用户 ID（引用 `directory.users.id`，应用级一致性）
- **审计用途**: 用于审计追踪
- **可为空**: 是（当 `revoked_at` 不为 NULL 时通常应设置）

### `revoke_reason` (TEXT)
- **用途**: 记录撤销授权的原因
- **可为空**: 是
- **使用场景**: 用于审计和问题追踪

## 索引 / Indexes

1. **`idx_grants_subject`**: 在 `(subject_type, subject_id)` 上创建索引，用于快速查找主体的所有授权
2. **`idx_grants_grant_type`**: 在 `(grant_type, grant_ref_id)` 上创建索引，用于快速查找特定角色/权限的所有授权
3. **`idx_grants_tenant_id`**: 在 `tenant_id` 字段上创建索引，用于按租户筛选授权
4. **`idx_grants_app_id`**: 在 `app_id` 字段上创建索引，用于按应用筛选授权
5. **`idx_grants_resource`**: 在 `(resource_type, resource_id)` 上创建索引，用于资源级授权查询
6. **`idx_grants_effect`**: 在 `effect` 字段上创建索引，用于筛选允许/拒绝授权
7. **`idx_grants_expires_at`**: 在 `expires_at` 字段上创建部分索引（WHERE expires_at IS NOT NULL），用于查找即将过期的授权
8. **`idx_grants_revoked_at`**: 在 `revoked_at` 字段上创建部分索引（WHERE revoked_at IS NULL），用于查找有效授权
9. **`idx_grants_subject_tenant`**: 在 `(subject_type, subject_id, tenant_id)` 上创建索引，用于查找主体在特定租户的授权

## 使用场景 / Use Cases

### 1. 新员工入职授权
**场景**: 新员工入职，需要授予租户管理员角色
**例子**:
- 新员工"张三"加入"公司 A"租户，需要授予 `tenant.admin` 角色
- 操作：`INSERT INTO access.grants (subject_type, subject_id, grant_type, grant_ref_id, tenant_id, created_by) VALUES ('USER', 'zhangsan-uuid', 'ROLE', (SELECT id FROM access.roles WHERE key = 'tenant.admin'), 'company-a-uuid', 'hr-admin-uuid')`
- 结果：张三自动获得 `tenant.admin` 角色包含的所有权限（通过 `role_permissions` 表查找）
- 验证：张三可以管理租户成员、查看用户列表等

### 2. 临时访问授权
**场景**: 临时授予用户访问权限，7 天后自动过期
**例子**:
- 外部审计人员"李四"需要临时访问用户数据，7 天后自动撤销
- 操作：`INSERT INTO access.grants (subject_type, subject_id, grant_type, grant_ref_id, tenant_id, expires_at, created_by) VALUES ('USER', 'lisi-uuid', 'PERMISSION', (SELECT id FROM access.permissions WHERE key = 'users.read'), 'company-a-uuid', NOW() + INTERVAL '7 days', 'admin-uuid')`
- 结果：李四获得 7 天的用户读取权限，7 天后自动失效（应用层需要检查 `expires_at`）
- 验证：7 天后，查询 `expires_at < NOW()` 的授权，自动拒绝访问

### 3. 资源级细粒度授权
**场景**: 用户只能管理特定的资源，不能管理其他资源
**例子**:
- 用户"王五"只能管理"项目 A"的资源，不能管理其他项目
- 操作：`INSERT INTO access.grants (subject_type, subject_id, grant_type, grant_ref_id, tenant_id, resource_type, resource_id, created_by) VALUES ('USER', 'wangwu-uuid', 'PERMISSION', (SELECT id FROM access.permissions WHERE key = 'assets.write'), 'company-a-uuid', 'project', 'project-a-uuid', 'admin-uuid')`
- 结果：王五只能对 `project-a-uuid` 执行 `assets.write` 操作
- 验证：当王五尝试修改其他项目时，系统检查 `resource_id`，拒绝访问

### 4. 应用级授权
**场景**: 某个服务客户端只能在特定应用内操作
**例子**:
- 服务客户端"数据分析服务"只能在"应用 B"内读取数据
- 操作：`INSERT INTO access.grants (subject_type, subject_id, grant_type, grant_ref_id, app_id, created_by) VALUES ('CLIENT', 'analytics-service-uuid', 'PERMISSION', (SELECT id FROM access.permissions WHERE key = 'users.read'), 'app-b-uuid', 'admin-uuid')`
- 结果：该服务只能在 `app-b-uuid` 应用内读取用户数据
- 验证：当服务尝试访问其他应用的数据时，系统检查 `app_id`，拒绝访问

### 5. 拒绝授权（黑名单）
**场景**: 明确拒绝某个用户的特定权限
**例子**:
- 用户"赵六"虽然有 `tenant.admin` 角色（包含 `users.export` 权限），但需要明确禁止导出功能
- 操作：`INSERT INTO access.grants (subject_type, subject_id, grant_type, grant_ref_id, tenant_id, effect, created_by) VALUES ('USER', 'zhaoliu-uuid', 'PERMISSION', (SELECT id FROM access.permissions WHERE key = 'users.export'), 'company-a-uuid', 'DENY', 'admin-uuid')`
- 结果：即使赵六有 `tenant.admin` 角色，DENY 授权优先级更高，无法导出用户数据
- 验证：系统检查授权时，DENY 授权优先于 ALLOW 授权

### 6. 跨租户全局授权
**场景**: 系统管理员需要跨所有租户的权限
**例子**:
- 系统管理员"系统管理员"需要全局权限，不受租户限制
- 操作：`INSERT INTO access.grants (subject_type, subject_id, grant_type, grant_ref_id, tenant_id, created_by) VALUES ('USER', 'system-admin-uuid', 'ROLE', (SELECT id FROM access.roles WHERE key = 'service.writer'), NULL, 'root-admin-uuid')`
- 结果：系统管理员获得全局 `service.writer` 角色，可以在所有租户内操作
- 验证：`tenant_id IS NULL` 表示全局授权，不受租户限制

### 7. 权限撤销
**场景**: 员工离职，需要撤销所有权限
**例子**:
- 员工"孙七"离职，需要撤销所有授权
- 操作：`UPDATE access.grants SET revoked_at = NOW(), revoked_by = 'hr-admin-uuid', revoke_reason = '员工离职' WHERE subject_type = 'USER' AND subject_id = 'sunqi-uuid' AND revoked_at IS NULL`
- 结果：孙七的所有授权被标记为已撤销
- 验证：系统检查授权时，过滤 `revoked_at IS NULL` 的记录，孙七无法访问任何资源

### 8. 批量授权
**场景**: 为整个部门批量授予权限
**例子**:
- "IT 部门"的所有成员需要 `tenant.viewer` 角色
- 操作：`INSERT INTO access.grants (subject_type, subject_id, grant_type, grant_ref_id, tenant_id, created_by) SELECT 'USER', user_id, 'ROLE', (SELECT id FROM access.roles WHERE key = 'tenant.viewer'), 'company-a-uuid', 'admin-uuid' FROM department_members WHERE department = 'IT'`
- 结果：IT 部门的所有成员都获得 `tenant.viewer` 角色
- 验证：批量查询所有 IT 部门成员的授权，确认都已授予

## 授权查询流程 / Authorization Query Flow

1. **收集主体信息**: 获取请求的用户或客户端信息
2. **查找授权**: 根据 `subject_type` 和 `subject_id` 查找所有有效授权（`revoked_at IS NULL` 且 `expires_at IS NULL OR expires_at > NOW()`）
3. **解析授权**: 
   - 如果 `grant_type='ROLE'`，通过 `role_permissions` 表查找角色包含的所有权限
   - 如果 `grant_type='PERMISSION'`，直接使用权限
4. **应用范围**: 根据 `tenant_id`、`app_id`、`resource_type`、`resource_id` 过滤授权
5. **处理冲突**: 如果同时有 ALLOW 和 DENY 授权，DENY 通常优先级更高
6. **返回结果**: 返回主体拥有的所有权限

## 查询示例 / Query Examples

```sql
-- 查询用户的所有有效授权
SELECT g.*, 
       CASE 
         WHEN g.grant_type = 'ROLE' THEN r.key
         WHEN g.grant_type = 'PERMISSION' THEN p.key
       END AS grant_key
FROM access.grants g
LEFT JOIN access.roles r ON g.grant_type = 'ROLE' AND g.grant_ref_id = r.id
LEFT JOIN access.permissions p ON g.grant_type = 'PERMISSION' AND g.grant_ref_id = p.id
WHERE g.subject_type = 'USER'
  AND g.subject_id = '...'
  AND g.revoked_at IS NULL
  AND (g.expires_at IS NULL OR g.expires_at > NOW());

-- 查询用户在特定租户的授权
SELECT g.*
FROM access.grants g
WHERE g.subject_type = 'USER'
  AND g.subject_id = '...'
  AND g.tenant_id = '...'
  AND g.revoked_at IS NULL
  AND (g.expires_at IS NULL OR g.expires_at > NOW());

-- 查询角色的所有授权（包括用户和客户端）
SELECT g.*, 
       CASE WHEN g.subject_type = 'USER' THEN 'User' ELSE 'Client' END AS subject_type_name
FROM access.grants g
WHERE g.grant_type = 'ROLE'
  AND g.grant_ref_id = '...'
  AND g.revoked_at IS NULL;

-- 查询即将过期的临时授权
SELECT g.*
FROM access.grants g
WHERE g.expires_at IS NOT NULL
  AND g.expires_at BETWEEN NOW() AND NOW() + INTERVAL '7 days'
  AND g.revoked_at IS NULL;
```

## 示例数据 / Example Data

```sql
-- 授予用户租户管理员角色
INSERT INTO access.grants (subject_type, subject_id, grant_type, grant_ref_id, tenant_id, created_by) VALUES
  (
    'USER',
    'user-uuid',
    'ROLE',
    (SELECT id FROM access.roles WHERE key = 'tenant.admin'),
    'tenant-uuid',
    'admin-user-id'
  );

-- 直接授予用户权限
INSERT INTO access.grants (subject_type, subject_id, grant_type, grant_ref_id, tenant_id, created_by) VALUES
  (
    'USER',
    'user-uuid',
    'PERMISSION',
    (SELECT id FROM access.permissions WHERE key = 'users.export'),
    'tenant-uuid',
    'admin-user-id'
  );

-- 临时授权（7 天后过期）
INSERT INTO access.grants (subject_type, subject_id, grant_type, grant_ref_id, tenant_id, expires_at, created_by) VALUES
  (
    'USER',
    'user-uuid',
    'ROLE',
    (SELECT id FROM access.roles WHERE key = 'tenant.viewer'),
    'tenant-uuid',
    NOW() + INTERVAL '7 days',
    'admin-user-id'
  );

-- 资源级授权（用户只能管理特定资源）
INSERT INTO access.grants (subject_type, subject_id, grant_type, grant_ref_id, tenant_id, resource_type, resource_id, created_by) VALUES
  (
    'USER',
    'user-uuid',
    'PERMISSION',
    (SELECT id FROM access.permissions WHERE key = 'assets.write'),
    'tenant-uuid',
    'asset',
    'asset-uuid',
    'admin-user-id'
  );
```

## 注意事项 / Notes

1. **应用级一致性**: `subject_id`、`grant_ref_id`、`tenant_id`、`app_id` 等字段使用应用级一致性，不建立数据库外键
2. **授权撤销**: 使用 `revoked_at` 字段实现软撤销，保留历史记录用于审计
3. **临时授权**: `expires_at` 字段用于实现临时授权，应用层需要定期清理过期授权
4. **授权冲突**: 当同时有 ALLOW 和 DENY 授权时，DENY 通常优先级更高
5. **角色权限继承**: 当 `grant_type='ROLE'` 时，主体自动获得角色包含的所有权限（通过 `role_permissions` 表）
6. **性能考虑**: 授权查询可能涉及多表关联，建议使用索引优化查询性能
7. **资源级授权**: `resource_type` 和 `resource_id` 用于实现细粒度的资源级授权

