# Members Table

## 概述 / Overview

`tenants.members` 表是租户成员关系表，表示"谁属于公司 A"（员工/管理员/开发者）。

企业注意：成员不仅仅是 "user_id"，还有成员状态和来源。

The `tenants.members` table represents "who belongs to Company A" (employees/admins/developers).

Enterprise note: Members are not just "user_id", but also have member status and source.

## 使用场景 / Use Cases

### 1. 添加成员
**场景**: 管理员添加成员到租户
**例子**:
- 管理员添加用户"张三"到公司 A
- 操作：`INSERT INTO tenants.members (tenant_id, user_id, status, source, created_by) VALUES ('company-a-uuid', 'zhangsan-uuid', 'ACTIVE', 'MANUAL', 'admin-uuid')`
- 结果：成员添加，状态为 `ACTIVE`

### 2. 邀请成员
**场景**: 管理员邀请新成员加入
**例子**:
- 管理员邀请 "lisi@example.com" 加入公司 A
- 先创建邀请：`INSERT INTO tenants.invitations (invite_id, tenant_id, email, token_hash, expires_at, invited_by) VALUES ('invite-uuid', 'company-a-uuid', 'lisi@example.com', 'hashed-token', NOW() + INTERVAL '7 days', 'admin-uuid')`
- 用户接受邀请后，创建成员：`INSERT INTO tenants.members (tenant_id, user_id, status, source, joined_at, created_by) VALUES ('company-a-uuid', 'lisi-uuid', 'ACTIVE', 'INVITE', NOW(), 'admin-uuid')`
- 结果：成员通过邀请加入

### 3. 移除成员
**场景**: 管理员移除成员
**例子**:
- 管理员移除成员
- 操作：`UPDATE tenants.members SET status = 'REMOVED', left_at = NOW(), updated_at = NOW() WHERE member_id = 'member-uuid'`
- 结果：成员被移除，但记录保留

### 4. 查询租户成员
**场景**: 查看租户的所有成员
**例子**:
- 查询活跃成员：`SELECT m.member_id, m.user_id, m.status, m.joined_at FROM tenants.members m WHERE m.tenant_id = 'company-a-uuid' AND m.status = 'ACTIVE' ORDER BY m.joined_at DESC`
- 结果：返回所有活跃成员

