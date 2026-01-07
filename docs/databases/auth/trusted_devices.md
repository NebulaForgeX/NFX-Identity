# Trusted Devices Table

## 概述 / Overview

`auth.trusted_devices` 表用于设备信任管理，支持"记住此设备 30 天，跳过 MFA"和风险控制白名单功能。

The `auth.trusted_devices` table manages device trust, supporting "Remember this device for 30 days, skip MFA" and risk control whitelist.

## 使用场景 / Use Cases

### 1. 标记设备为信任设备
**场景**: 用户选择"记住此设备"，30 天内跳过 MFA
**例子**:
- 用户登录时选择"记住此设备"
- 操作：`INSERT INTO auth.trusted_devices (device_id, user_id, tenant_id, device_fingerprint_hash, device_name, trusted_until, ip, ua_hash) VALUES ('device-uuid', 'zhangsan-uuid', 'tenant-uuid', 'fingerprint-hash', 'My MacBook', NOW() + INTERVAL '30 days', '203.0.113.1', 'ua-hash')`
- 结果：该设备在 30 天内登录时跳过 MFA 验证

### 2. 检查设备是否信任
**场景**: 用户登录时检查设备是否在信任列表中
**例子**:
- 用户从某个设备登录
- 查询：`SELECT trusted_until FROM auth.trusted_devices WHERE user_id = 'zhangsan-uuid' AND device_fingerprint_hash = 'fingerprint-hash' AND trusted_until > NOW()`
- 如果存在且未过期，跳过 MFA；否则要求 MFA 验证

### 3. 更新设备最后使用时间
**场景**: 用户使用信任设备时，更新最后使用时间
**例子**:
- 用户使用信任设备登录
- 操作：`UPDATE auth.trusted_devices SET last_used_at = NOW(), updated_at = NOW() WHERE device_id = 'device-uuid'`
- 结果：保持设备信任状态

### 4. 撤销设备信任
**场景**: 用户或管理员撤销设备信任
**例子**:
- 用户发现设备丢失，撤销信任
- 操作：`DELETE FROM auth.trusted_devices WHERE device_id = 'lost-device-uuid' AND user_id = 'zhangsan-uuid'`
- 结果：该设备不再被信任，需要 MFA 验证

