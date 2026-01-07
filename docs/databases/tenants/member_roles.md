# Member Roles Table

## 概述 / Overview

`tenants.member_roles` 表用于租户内的角色分配，表示"公司 A 员工在公司 A 中拥有什么权限级别"。

企业推荐：使用 member_id 而不是 user_id，因为：
- 用户在不同租户中有不同的身份/状态；member_id 更清晰
- 审计："member" 比 "user" 更接近业务语义
- 未来支持"临时角色/过期角色/委托"更容易

The `tenants.member_roles` table represents "what permission level does a Company A employee have in Company A".

## 使用场景 / Use Cases

### 1. 分配角色
**场景**: 管理员为成员分配角色
**例子**:
- 管理员为成员分配 "tenant.admin" 角色
- 操作：`INSERT INTO tenants.member_roles (tenant_id, member_id, role_id, assigned_by) VALUES ('company-a-uuid', 'member-uuid', (SELECT id FROM access.roles WHERE key = 'tenant.admin'), 'admin-uuid')`
- 结果：成员获得租户管理员角色

### 2. 临时角色
**场景**: 为成员分配临时角色
**例子**:
- 为成员分配临时 "tenant.viewer" 角色，30 天后过期
- 操作：`INSERT INTO tenants.member_roles (tenant_id, member_id, role_id, expires_at, assigned_by) VALUES ('company-a-uuid', 'member-uuid', (SELECT id FROM access.roles WHERE key = 'tenant.viewer'), NOW() + INTERVAL '30 days', 'admin-uuid')`
- 结果：成员获得临时角色，30 天后自动失效

### 3. 撤销角色
**场景**: 管理员撤销成员角色
**例子**:
- 管理员撤销成员角色
- 操作：`UPDATE tenants.member_roles SET revoked_at = NOW(), revoked_by = 'admin-uuid', revoke_reason = 'role_change' WHERE member_id = 'member-uuid' AND role_id = 'role-uuid'`
- 结果：角色被撤销

### 4. 查询成员角色
**场景**: 查看成员的所有角色
**例子**:
- 查询成员角色：`SELECT mr.role_id, mr.assigned_at, mr.expires_at FROM tenants.member_roles mr WHERE mr.member_id = 'member-uuid' AND mr.revoked_at IS NULL AND (mr.expires_at IS NULL OR mr.expires_at > NOW())`
- 结果：返回成员的所有有效角色

