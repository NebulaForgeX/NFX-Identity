# System State Table

## 概述 / Overview

`system.system_state` 表用于记录系统初始化状态和引导状态。

用于在服务启动时检查系统是否已初始化。

The `system.system_state` table records system initialization and bootstrap state.

Used to check if system has been initialized on service startup.

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 记录的唯一标识符 |
| `initialized` | BOOLEAN | NOT NULL, DEFAULT false | 系统是否已初始化 |
| `initialized_at` | TIMESTAMP | NULL | 系统初始化时间 |
| `initialization_version` | VARCHAR(50) | NULL | 初始化版本号（用于迁移追踪） |
| `last_reset_at` | TIMESTAMP | NULL | 最后重置时间 |
| `last_reset_by` | UUID | NULL | 重置系统的用户ID（引用 directory.users.id，应用级一致性）。即使重置后数据库清空，可通过日志文件追溯） |
| `reset_count` | INTEGER | NOT NULL, DEFAULT 0 | 系统重置次数 |
| `metadata` | JSONB | DEFAULT '{}'::jsonb | 扩展字段：{"bootstrap_token": "...", "services_initialized": [...], ...} |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 记录创建时间（用于确定最新状态） |
| `updated_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 记录更新时间 |

## 查询逻辑 / Query Logic

系统启动时检查初始化状态：

```sql
-- 查询最新记录
SELECT initialized FROM system.system_state ORDER BY created_at DESC LIMIT 1;
```

判断逻辑：
- **如果没有记录** → 系统未初始化
- **如果 `initialized = false`** → 系统未初始化
- **如果 `initialized = true`** → 系统已初始化

## 使用场景 / Use Cases

### 1. 服务启动时检查初始化状态
**场景**: 服务启动时检查系统是否已初始化
**例子**:
- 服务启动时执行检查
- 查询：`SELECT initialized FROM system.system_state ORDER BY created_at DESC LIMIT 1`
- 如果没有记录或返回 `false`，系统未初始化，需要引导初始化流程
- 如果返回 `true`，系统已初始化，正常启动

### 2. 系统初始化
**场景**: 通过 `/bootstrap/initialize` gRPC 接口初始化系统
**例子**:
- 调用初始化接口，创建第一个系统管理员
- 操作：`INSERT INTO system.system_state (initialized, initialized_at, initialization_version) VALUES (true, NOW(), '1.0.0')`
- 结果：系统标记为已初始化，创建时间记录在 `created_at`

### 3. 系统重置
**场景**: 重置系统，允许重新初始化
**例子**:
- 管理员执行系统重置（重置前会生成日志文件记录执行者，用于最终追责）
- 选项1：删除所有记录
  - 操作：`DELETE FROM system.system_state`
  - 结果：表为空，下次启动时检测为未初始化
  - 注意：即使数据库清空，可通过日志文件追溯重置操作
- 选项2：创建新记录标记为未初始化
  - 操作：`INSERT INTO system.system_state (initialized, last_reset_at, last_reset_by, reset_count) VALUES (false, NOW(), 'admin-user-uuid', (SELECT COALESCE(MAX(reset_count), 0) + 1 FROM system.system_state))`
  - 结果：创建新记录，`initialized = false`，重置次数递增，记录重置者

### 4. 查询初始化历史
**场景**: 查看系统初始化历史记录
**例子**:
- 查询所有初始化记录：`SELECT initialized, initialized_at, last_reset_at, last_reset_by, created_at FROM system.system_state ORDER BY created_at DESC`
- 结果：返回所有初始化/重置记录，按时间倒序

## 索引 / Indexes

- **`idx_system_state_initialized`**: 在 `initialized` 字段上创建索引，用于快速筛选
- **`idx_system_state_created_at`**: 在 `created_at` 字段上创建降序索引，用于查询最新记录

## 注意事项 / Notes

1. **不插入初始值**: 表创建时不插入任何记录，通过查询结果判断状态
2. **支持多次重置**: 允许删除记录或创建新记录来重置系统状态
3. **最新记录优先**: 始终通过 `ORDER BY created_at DESC LIMIT 1` 查询最新状态
4. **重置追责**: `last_reset_by` 记录重置操作者。即使重置后数据库清空，可通过重置前生成的日志文件追溯，用于最终追责
5. **应用级一致性**: `last_reset_by` 引用 `directory.users.id`，但使用应用级一致性（无外键约束）
