# Groups Table

## 概述 / Overview

`tenants.groups` 表用于租户内的组织架构管理，支持层级结构（部门、团队等）。

The `tenants.groups` table manages organizational structure within tenants, supporting hierarchical structure (departments, teams, etc.).

## 使用场景 / Use Cases

### 1. 创建组
**场景**: 管理员创建组织组（如部门、团队）
**例子**:
- 创建"技术部"组
- 操作：`INSERT INTO tenants.groups (tenant_id, name, description, group_type) VALUES ('company-a-uuid', '技术部', '技术部门', 'department')`
- 结果：组创建，可以添加成员

### 2. 创建子组
**场景**: 创建层级结构（部门下的团队）
**例子**:
- 创建"技术部"下的"前端团队"
- 先查询父组：`SELECT id FROM tenants.groups WHERE tenant_id = 'company-a-uuid' AND name = '技术部'`
- 操作：`INSERT INTO tenants.groups (tenant_id, name, description, parent_group_id, group_type) VALUES ('company-a-uuid', '前端团队', '前端开发团队', 'tech-dept-uuid', 'team')`
- 结果：子组创建，形成层级结构

### 3. 添加成员到组
**场景**: 将成员添加到组
**例子**:
- 将成员添加到"前端团队"
- 操作：`INSERT INTO tenants.member_groups (member_id, group_id) VALUES ('member-uuid', 'frontend-team-uuid')`
- 结果：成员加入组

