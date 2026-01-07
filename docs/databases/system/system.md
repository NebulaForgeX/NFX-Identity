# System Schema 总览 / System Schema Overview

## 概述 / Overview

`system` schema 是 NFX-Identity 平台的系统级状态和管理模块，负责记录系统初始化状态、引导流程等。

The `system` schema is the system-level state and administration module of the NFX-Identity platform, responsible for recording system initialization state, bootstrap processes, etc.

## 核心概念 / Core Concepts

### 1. 系统状态（System State）
- **定义**: 记录系统初始化状态
- **表**: `system.system_state`
- **特点**: 用于服务启动时检查系统是否已初始化

## 表关系图 / Table Relationships

```
┌─────────────────┐
│  system_state   │
│  (系统状态)     │
└─────────────────┘
```

## 表列表 / Table List

### 1. `system.system_state` - 系统状态表
- **用途**: 记录系统初始化状态和引导状态
- **关键字段**: `initialized`、`initialized_at`、`initialized_by`、`created_at`
- **详细文档**: [system_state.md](./system_state.md)

## 初始化流程 / Initialization Flow

### 1. 服务启动检查
```
1. 服务启动时查询 system_state 表
2. SELECT initialized FROM system_state ORDER BY created_at DESC LIMIT 1
3. 如果没有记录或 initialized = false → 系统未初始化
4. 如果 initialized = true → 系统已初始化，正常启动
```

### 2. 系统初始化（/bootstrap/initialize）
```
1. 调用 /bootstrap/initialize gRPC 接口
2. 通过 gRPC 与各个微服务通信
3. 创建第一个系统管理员（不可删除）
4. 创建 system_state 记录，设置 initialized = true
5. 记录初始化时间、初始化者、版本号
```

### 3. 系统重置
```
1. 管理员执行系统重置（重置前生成日志文件记录执行者，用于最终追责）
2. 删除所有 system_state 记录（或创建新记录标记为未初始化）
3. 记录 last_reset_by（即使重置后数据库清空，可通过日志文件追溯）
4. 系统状态恢复为未初始化
5. 允许重新执行初始化流程
```

## 查询示例 / Query Examples

### 检查系统是否已初始化
```sql
-- 查询最新状态
SELECT initialized 
FROM system.system_state 
ORDER BY created_at DESC 
LIMIT 1;

-- 如果没有返回结果，系统未初始化
-- 如果返回 false，系统未初始化
-- 如果返回 true，系统已初始化
```

### 初始化系统
```sql
INSERT INTO system.system_state (
  initialized, 
  initialized_at, 
  initialized_by, 
  initialization_version
) VALUES (
  true, 
  NOW(), 
  'admin-user-uuid', 
  '1.0.0'
);
```

### 重置系统
```sql
-- 方式1: 删除所有记录（重置前会生成日志文件记录执行者）
DELETE FROM system.system_state;

-- 方式2: 创建新记录标记为未初始化（记录重置者）
INSERT INTO system.system_state (
  initialized, 
  last_reset_at, 
  last_reset_by, 
  reset_count
) VALUES (
  false, 
  NOW(), 
  'admin-user-uuid', 
  (SELECT COALESCE(MAX(reset_count), 0) + 1 FROM system.system_state)
);
-- 注意：即使重置后数据库清空，可通过日志文件追溯重置操作
```

## 相关文档 / Related Documentation

- [system_state.md](./system_state.md) - 系统状态表详细文档
