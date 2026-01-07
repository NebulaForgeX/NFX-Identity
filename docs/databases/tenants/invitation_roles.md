# Invitation Roles Table

## 概述 / Overview

`tenants.invitation_roles` 表用于邀请时预分配的角色。

当用户接受邀请时，自动获得这些角色。

The `tenants.invitation_roles` table stores roles pre-assigned to invitations.

When user accepts invitation, they automatically get these roles.

## 使用场景 / Use Cases

### 1. 邀请时预分配角色
**场景**: 管理员邀请用户时预分配角色
**例子**:
- 管理员邀请用户，预分配 "tenant.viewer" 角色
- 先创建邀请：`INSERT INTO tenants.invitations (invite_id, tenant_id, email, token_hash, expires_at, invited_by) VALUES ('invite-uuid', 'company-a-uuid', 'lisi@example.com', 'hashed-token', NOW() + INTERVAL '7 days', 'admin-uuid')`
- 预分配角色：`INSERT INTO tenants.invitation_roles (invite_id, role_id) VALUES ('invite-uuid', (SELECT id FROM access.roles WHERE key = 'tenant.viewer'))`
- 结果：用户接受邀请后自动获得该角色

### 2. 接受邀请时应用角色
**场景**: 用户接受邀请时，自动分配预定义的角色
**例子**:
- 用户接受邀请
- 查询预分配角色：`SELECT role_id FROM tenants.invitation_roles WHERE invite_id = 'invite-uuid'`
- 为成员分配角色：`INSERT INTO tenants.member_roles (tenant_id, member_id, role_id) SELECT 'company-a-uuid', 'member-uuid', role_id FROM tenants.invitation_roles WHERE invite_id = 'invite-uuid'`
- 结果：成员自动获得预分配的角色

