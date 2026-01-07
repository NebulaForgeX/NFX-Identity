# Member Groups Table

## 概述 / Overview

`tenants.member_groups` 表用于成员和组的多对多关系。

将成员添加到组织架构中的组（部门、团队等）。

The `tenants.member_groups` table represents the many-to-many relationship between members and groups.

## 使用场景 / Use Cases

### 1. 添加成员到组
**场景**: 将成员添加到组
**例子**:
- 将成员添加到"技术部"
- 操作：`INSERT INTO tenants.member_groups (member_id, group_id) VALUES ('member-uuid', 'tech-dept-uuid')`
- 结果：成员加入组，可以在组内进行权限管理

### 2. 查询组成员
**场景**: 查看组的所有成员
**例子**:
- 查询"技术部"的所有成员：`SELECT m.member_id, m.user_id FROM tenants.members m JOIN tenants.member_groups mg ON m.member_id = mg.member_id WHERE mg.group_id = 'tech-dept-uuid' AND m.status = 'ACTIVE'`
- 结果：返回组的所有活跃成员

### 3. 查询成员所属的组
**场景**: 查看成员属于哪些组
**例子**:
- 查询成员所属的组：`SELECT g.name, g.type FROM tenants.groups g JOIN tenants.member_groups mg ON g.id = mg.group_id WHERE mg.member_id = 'member-uuid'`
- 结果：返回成员所属的所有组

