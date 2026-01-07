# Auth Schema 总览 / Auth Schema Overview

## 概述 / Overview

`auth` schema 是 NFX-Identity 平台的身份认证（Authentication）核心模块，负责管理用户登录、会话、MFA、密码重置、账户锁定等认证相关功能。

The `auth` schema is the authentication core module of the NFX-Identity platform, responsible for managing user login, sessions, MFA, password resets, account lockouts, and other authentication-related functions.

## 核心概念 / Core Concepts

### 1. 用户凭据（User Credentials）
- **定义**: 存储用户的各种认证凭据（密码、Passkey、OAuth 等）
- **表**: `auth.user_credentials`
- **特点**: 支持多种认证方式，密码哈希存储

### 2. 会话管理（Sessions）
- **定义**: 管理用户登录会话
- **表**: `auth.sessions`
- **特点**: 支持设备管理、会话撤销、并发会话限制

### 3. 刷新令牌（Refresh Tokens）
- **定义**: 长期会话的刷新令牌
- **表**: `auth.refresh_tokens`
- **特点**: 支持令牌轮换、撤销

### 4. 登录尝试（Login Attempts）
- **定义**: 记录所有登录尝试
- **表**: `auth.login_attempts`
- **特点**: 用于账户锁定、速率限制、风险评分

### 5. 账户锁定（Account Lockouts）
- **定义**: 管理账户锁定状态
- **表**: `auth.account_lockouts`
- **特点**: 支持自动锁定和手动锁定

### 6. 密码重置（Password Resets）
- **定义**: 安全的密码重置流程
- **表**: `auth.password_resets`
- **特点**: 支持审计、滥用防护

### 7. MFA 因子（MFA Factors）
- **定义**: 多因素认证凭据
- **表**: `auth.mfa_factors`
- **特点**: 支持 TOTP、SMS、Email、WebAuthn、备份代码

### 8. 密码历史（Password History）
- **定义**: 防止密码重用
- **表**: `auth.password_history`
- **特点**: 满足企业合规要求

### 9. 信任设备（Trusted Devices）
- **定义**: 设备信任管理
- **表**: `auth.trusted_devices`
- **特点**: 支持"记住此设备"，跳过 MFA

### 10. 凭据事件（Credential Events）
- **定义**: 凭据变更的审计轨迹
- **表**: `auth.credential_events`
- **特点**: 用于安全审计

## 表关系图 / Table Relationships

```
┌─────────────────────┐
│  user_credentials   │
│  (用户凭据)         │
└──────────┬──────────┘
           │
           │ 1:N
           │
    ┌──────▼──────────────┐
    │   login_attempts    │
    │   (登录尝试)        │
    └─────────────────────┘
           │
           │ 触发
           │
    ┌──────▼──────────────┐
    │  account_lockouts   │
    │  (账户锁定)         │
    └─────────────────────┘

┌─────────────────────┐
│    sessions         │
│    (会话)           │
└──────────┬──────────┘
           │
           │ 1:N
           │
    ┌──────▼──────────────┐
    │  refresh_tokens     │
    │  (刷新令牌)         │
    └─────────────────────┘

┌─────────────────────┐
│   mfa_factors       │
│   (MFA 因子)        │
└─────────────────────┘

┌─────────────────────┐
│  password_resets    │
│  (密码重置)         │
└─────────────────────┘

┌─────────────────────┐
│  password_history   │
│  (密码历史)         │
└─────────────────────┘

┌─────────────────────┐
│  trusted_devices    │
│  (信任设备)         │
└─────────────────────┘

┌─────────────────────┐
│ credential_events   │
│ (凭据事件)          │
└─────────────────────┘
```

## 表列表 / Table List

### 1. `auth.user_credentials` - 用户凭据表
- **用途**: 存储用户的各种认证凭据
- **关键字段**: `user_id`、`credential_type`、`password_hash`、`status`
- **详细文档**: [user_credentials.md](./user_credentials.md)

### 2. `auth.sessions` - 会话表
- **用途**: 管理用户登录会话
- **关键字段**: `session_id`、`user_id`、`expires_at`、`revoked_at`
- **详细文档**: [sessions.md](./sessions.md)

### 3. `auth.refresh_tokens` - 刷新令牌表
- **用途**: 存储长期会话的刷新令牌
- **关键字段**: `token_id`、`user_id`、`expires_at`、`rotated_from`
- **详细文档**: [refresh_tokens.md](./refresh_tokens.md)

### 4. `auth.login_attempts` - 登录尝试表
- **用途**: 记录所有登录尝试
- **关键字段**: `identifier`、`user_id`、`success`、`failure_code`
- **详细文档**: [login_attempts.md](./login_attempts.md)

### 5. `auth.account_lockouts` - 账户锁定表
- **用途**: 管理账户锁定状态
- **关键字段**: `user_id`、`locked_until`、`lock_reason`
- **详细文档**: [account_lockouts.md](./account_lockouts.md)

### 6. `auth.password_resets` - 密码重置表
- **用途**: 安全的密码重置流程
- **关键字段**: `reset_id`、`user_id`、`code_hash`、`expires_at`
- **详细文档**: [password_resets.md](./password_resets.md)

### 7. `auth.mfa_factors` - MFA 因子表
- **用途**: 存储多因素认证凭据
- **关键字段**: `factor_id`、`user_id`、`type`、`enabled`
- **详细文档**: [mfa_factors.md](./mfa_factors.md)

### 8. `auth.password_history` - 密码历史表
- **用途**: 防止密码重用
- **关键字段**: `user_id`、`password_hash`、`created_at`
- **详细文档**: [password_history.md](./password_history.md)

### 9. `auth.trusted_devices` - 信任设备表
- **用途**: 设备信任管理
- **关键字段**: `device_id`、`user_id`、`trusted_until`
- **详细文档**: [trusted_devices.md](./trusted_devices.md)

### 10. `auth.credential_events` - 凭据事件表
- **用途**: 凭据变更的审计轨迹
- **关键字段**: `event_id`、`user_id`、`event_type`、`actor_type`
- **详细文档**: [credential_events.md](./credential_events.md)

## 认证流程 / Authentication Flow

### 1. 用户登录流程
```
1. 用户输入用户名/密码
2. 系统查询 user_credentials 验证密码
3. 记录 login_attempts（成功或失败）
4. 如果失败次数过多，创建 account_lockouts
5. 如果成功，检查是否需要 MFA
6. 如果需要 MFA，验证 MFA 因子
7. 检查设备是否在 trusted_devices 中
8. 创建 sessions 和 refresh_tokens
9. 返回访问令牌和刷新令牌
```

### 2. 密码重置流程
```
1. 用户请求密码重置
2. 系统创建 password_resets 记录
3. 发送重置邮件/短信
4. 用户点击重置链接
5. 验证重置令牌
6. 用户设置新密码
7. 检查 password_history 防止重用
8. 更新 user_credentials
9. 记录 password_history
10. 记录 credential_events
11. 撤销所有 refresh_tokens
```

### 3. MFA 启用流程
```
1. 用户选择启用 MFA
2. 根据类型生成凭据（TOTP 密钥、WebAuthn 等）
3. 创建 mfa_factors 记录
4. 用户验证 MFA 因子
5. 启用 mfa_factors（enabled = true）
6. 记录 credential_events
```

## 安全特性 / Security Features

### 1. 密码安全
- **哈希存储**: 永远不存储明文密码
- **强算法**: 使用 argon2id 或 bcrypt
- **密码历史**: 防止重用最近 N 个密码
- **强制更改**: 支持强制用户更改密码

### 2. 账户保护
- **自动锁定**: 多次失败登录后自动锁定
- **速率限制**: 限制登录尝试频率
- **IP 检测**: 检测异常 IP 登录
- **设备指纹**: 检测异常设备

### 3. 会话安全
- **令牌轮换**: 刷新令牌支持轮换
- **会话撤销**: 支持撤销会话
- **并发限制**: 限制并发会话数
- **自动过期**: 会话自动过期

### 4. MFA 支持
- **多种方式**: TOTP、SMS、Email、WebAuthn
- **备份代码**: 支持备份代码
- **信任设备**: 支持"记住此设备"

## 相关文档 / Related Documentation

- [user_credentials.md](./user_credentials.md) - 用户凭据表详细文档
- [sessions.md](./sessions.md) - 会话表详细文档
- [refresh_tokens.md](./refresh_tokens.md) - 刷新令牌表详细文档
- [login_attempts.md](./login_attempts.md) - 登录尝试表详细文档
- [account_lockouts.md](./account_lockouts.md) - 账户锁定表详细文档
- [password_resets.md](./password_resets.md) - 密码重置表详细文档
- [mfa_factors.md](./mfa_factors.md) - MFA 因子表详细文档
- [password_history.md](./password_history.md) - 密码历史表详细文档
- [trusted_devices.md](./trusted_devices.md) - 信任设备表详细文档
- [credential_events.md](./credential_events.md) - 凭据事件表详细文档

