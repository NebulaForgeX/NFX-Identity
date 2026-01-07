# Tenant Apps Table

## 概述 / Overview

`tenants.tenant_apps` 表用于租户和应用之间的关系，决定"公司 A 可以管理/使用哪些应用，以及这些应用配置属于哪个租户"。

企业常见用法：
- 一个租户有多个应用（Web、Mobile、API、Backend 等）
- 应用属于租户，但应用权限/scopes、回调、安全策略必须与租户强绑定

The `tenants.tenant_apps` table represents the relationship between tenant and applications (client/app).

## 使用场景 / Use Cases

### 1. 关联应用到租户
**场景**: 管理员将应用关联到租户
**例子**:
- 管理员将"数据分析服务"应用关联到公司 A
- 操作：`INSERT INTO tenants.tenant_apps (tenant_id, app_id, status, created_by) VALUES ('company-a-uuid', 'analytics-service-uuid', 'ACTIVE', 'admin-uuid')`
- 结果：应用关联到租户，租户可以使用该应用

### 2. 禁用租户应用
**场景**: 管理员禁用租户的某个应用
**例子**:
- 管理员禁用租户的某个应用
- 操作：`UPDATE tenants.tenant_apps SET status = 'DISABLED', updated_at = NOW() WHERE tenant_id = 'company-a-uuid' AND app_id = 'app-uuid'`
- 结果：租户无法使用该应用

### 3. 查询租户的应用
**场景**: 查看租户的所有应用
**例子**:
- 查询租户的应用：`SELECT ta.app_id, ta.status, ta.settings FROM tenants.tenant_apps ta WHERE ta.tenant_id = 'company-a-uuid' AND ta.status = 'ACTIVE'`
- 结果：返回租户的所有活跃应用

