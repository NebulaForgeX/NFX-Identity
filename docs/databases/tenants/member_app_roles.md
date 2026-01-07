# Member App Roles Table

## 概述 / Overview

`tenants.member_app_roles` 表用于成员在特定应用中的角色。

支持应用级别的权限管理。

The `tenants.member_app_roles` table stores roles for members in specific applications.

## 使用场景 / Use Cases

### 1. 分配应用角色
**场景**: 为成员分配在特定应用中的角色
**例子**:
- 为成员在"数据分析服务"应用中分配 "app.operator" 角色
- 操作：`INSERT INTO tenants.member_app_roles (member_id, app_id, role_id) VALUES ('member-uuid', 'analytics-service-uuid', (SELECT id FROM access.roles WHERE key = 'app.operator'))`
- 结果：成员在该应用中获得操作员角色

### 2. 查询成员的应用角色
**场景**: 查看成员在哪些应用中有角色
**例子**:
- 查询成员的应用角色：`SELECT mar.app_id, mar.role_id FROM tenants.member_app_roles mar WHERE mar.member_id = 'member-uuid'`
- 结果：返回成员的所有应用角色

