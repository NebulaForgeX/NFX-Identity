# Password History Table

## 概述 / Overview

`auth.password_history` 表用于防止密码重用，满足企业合规要求：防止使用最近 N 个密码。

The `auth.password_history` table prevents password reuse, meeting enterprise compliance requirements: prevent using last N passwords.

## 使用场景 / Use Cases

### 1. 记录密码历史
**场景**: 用户更改密码时，记录旧密码到历史表
**例子**:
- 用户更改密码从 "OldPassword123!" 到 "NewPassword456!"
- 操作：`INSERT INTO auth.password_history (user_id, tenant_id, password_hash, hash_alg) VALUES ('zhangsan-uuid', 'tenant-uuid', 'old-password-hash', 'argon2id')`
- 结果：旧密码被记录到历史表

### 2. 检查密码重用
**场景**: 用户尝试更改密码时，检查是否与历史密码重复
**例子**:
- 用户尝试更改密码为 "OldPassword123!"（最近使用过）
- 查询：`SELECT password_hash FROM auth.password_history WHERE user_id = 'zhangsan-uuid' ORDER BY created_at DESC LIMIT 5`
- 对新密码进行哈希，与历史密码比较
- 如果匹配，拒绝更改，提示"不能使用最近 5 个密码"

### 3. 清理旧密码历史
**场景**: 定期清理超过保留期限的密码历史
**例子**:
- 系统只保留最近 12 个月的密码历史
- 查询：`SELECT id FROM auth.password_history WHERE created_at < NOW() - INTERVAL '12 months'`
- 操作：`DELETE FROM auth.password_history WHERE created_at < NOW() - INTERVAL '12 months'`
- 结果：超过 12 个月的密码历史被删除

