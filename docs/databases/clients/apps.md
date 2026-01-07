# Apps Table

## 概述 / Overview

`clients.apps` 表是应用/客户端主表，表示可以集成到平台的系统/应用/后端服务。

公司 A 可能有多个系统，每个系统应该有一个 app。

The `clients.apps` table is the main table for applications/clients, representing systems/applications/backend services that can integrate with the platform.

Company A may have multiple systems, each should have one app.

## 使用场景 / Use Cases

### 1. 注册新应用
**场景**: 公司注册新应用，用于集成平台 API
**例子**:
- 公司 A 需要集成"数据分析服务"
- 操作：`INSERT INTO clients.apps (app_id, tenant_id, name, description, type, status, environment, created_by) VALUES ('analytics-service', 'company-a-uuid', '数据分析服务', '用于数据分析的后端服务', 'service', 'pending', 'production', 'admin-uuid')`
- 结果：创建新应用，状态为 `pending`，等待审核激活

### 2. 环境隔离
**场景**: 同一应用在不同环境（生产/测试/开发）使用不同的凭证
**例子**:
- 公司 A 的"数据分析服务"需要生产环境和测试环境
- 创建生产应用：`INSERT INTO clients.apps (app_id, tenant_id, name, environment, status) VALUES ('analytics-service-prod', 'company-a-uuid', '数据分析服务-生产', 'production', 'active')`
- 创建测试应用：`INSERT INTO clients.apps (app_id, tenant_id, name, environment, status) VALUES ('analytics-service-test', 'company-a-uuid', '数据分析服务-测试', 'test', 'active')`
- 结果：两个环境完全隔离，使用不同的凭证和配置

### 3. 应用状态管理
**场景**: 管理员管理应用状态（激活/禁用/暂停）
**例子**:
- 管理员发现应用异常，暂时禁用
- 操作：`UPDATE clients.apps SET status = 'suspended', updated_by = 'admin-uuid' WHERE app_id = 'analytics-service'`
- 结果：应用被暂停，无法使用 API

### 4. 查询租户的应用
**场景**: 查看某个租户的所有应用
**例子**:
- 管理员查看公司 A 的所有应用
- 查询：`SELECT app_id, name, type, status, environment FROM clients.apps WHERE tenant_id = 'company-a-uuid' AND deleted_at IS NULL ORDER BY created_at DESC`
- 结果：显示所有应用及其状态

