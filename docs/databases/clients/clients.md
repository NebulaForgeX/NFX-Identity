# Clients Schema 总览 / Clients Schema Overview

## 概述 / Overview

`clients` schema 是 NFX-Identity 平台的客户端/应用管理模块，负责管理集成到平台的系统/应用/后端服务，包括 OAuth 客户端、API Key、服务令牌、速率限制等。

The `clients` schema is the client/application management module of the NFX-Identity platform, responsible for managing systems/applications/backend services integrated with the platform, including OAuth clients, API Keys, service tokens, rate limits, etc.

## 核心概念 / Core Concepts

### 1. 应用（Apps）
- **定义**: 表示可以集成到平台的系统/应用/后端服务
- **表**: `clients.apps`
- **特点**: 支持多环境（生产/测试/开发），多租户隔离

### 2. 客户端凭证（Client Credentials）
- **定义**: OAuth 客户端凭证（client_id 和 secret）
- **表**: `clients.client_credentials`
- **特点**: 支持凭证轮换，哈希存储

### 3. API 密钥（API Keys）
- **定义**: 简化的 API 密钥凭证
- **表**: `clients.api_keys`
- **特点**: 用于脚本/内部系统，需要更强的约束

### 4. 客户端 Scope（Client Scopes）
- **定义**: 应用允许的权限范围（白名单）
- **表**: `clients.client_scopes`
- **特点**: 决定应用可以请求哪些 scope

### 5. IP 白名单（IP Allowlist）
- **定义**: 限制请求来源 IP
- **表**: `clients.ip_allowlist`
- **特点**: 降低凭证泄露风险

### 6. 服务令牌（Service Tokens）
- **定义**: M2M 访问令牌
- **表**: `clients.service_tokens`
- **特点**: 支持令牌撤销、追踪

### 7. 速率限制（Rate Limits）
- **定义**: 客户端级别的速率限制
- **表**: `clients.rate_limits`
- **特点**: 防止滥用

### 8. 令牌使用日志（Token Usage Logs）
- **定义**: 记录每次令牌使用
- **表**: `clients.token_usage_logs`
- **特点**: 用于审计、分析和滥用检测

## 表关系图 / Table Relationships

```
┌─────────────────┐
│      apps       │
│  (应用主表)     │
└────────┬────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────────┐
    │  client_credentials      │
    │  (OAuth 客户端凭证)      │
    └───────────────────────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────────┐
    │      api_keys             │
    │  (API 密钥)               │
    └───────────────────────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────────┐
    │    client_scopes          │
    │  (客户端 Scope)            │
    └───────────────────────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────────┐
    │    ip_allowlist           │
    │  (IP 白名单)               │
    └───────────────────────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────────┐
    │   service_tokens           │
    │  (服务令牌)                │
    └──────┬─────────────────────┘
           │
           │ 1:N
           │
    ┌──────▼─────────────────────┐
    │  token_usage_logs          │
    │  (令牌使用日志)             │
    └─────────────────────────────┘

┌─────────────────┐
│   rate_limits   │
│  (速率限制)     │
│  (独立表)       │
└─────────────────┘
```

## 表列表 / Table List

### 1. `clients.apps` - 应用表
- **用途**: 应用主表，表示可以集成到平台的系统
- **关键字段**: `app_id`、`tenant_id`、`name`、`type`、`status`、`environment`
- **详细文档**: [apps.md](./apps.md)

### 2. `clients.client_credentials` - 客户端凭证表
- **用途**: 存储 OAuth 客户端凭证
- **关键字段**: `client_id`、`secret_hash`、`status`、`rotated_at`
- **详细文档**: [client_credentials.md](./client_credentials.md)

### 3. `clients.api_keys` - API 密钥表
- **用途**: 存储 API 密钥凭证
- **关键字段**: `key_id`、`key_hash`、`status`、`expires_at`
- **详细文档**: [api_keys.md](./api_keys.md)

### 4. `clients.client_scopes` - 客户端 Scope 表
- **用途**: 定义应用允许的权限范围
- **关键字段**: `app_id`、`scope`、`expires_at`
- **详细文档**: [client_scopes.md](./client_scopes.md)

### 5. `clients.ip_allowlist` - IP 白名单表
- **用途**: 限制请求来源 IP
- **关键字段**: `app_id`、`cidr`、`status`
- **详细文档**: [ip_allowlist.md](./ip_allowlist.md)

### 6. `clients.service_tokens` - 服务令牌表
- **用途**: 存储 M2M 访问令牌
- **关键字段**: `token_id`、`app_id`、`scopes`、`expires_at`
- **详细文档**: [service_tokens.md](./service_tokens.md)

### 7. `clients.rate_limits` - 速率限制表
- **用途**: 客户端级别的速率限制
- **关键字段**: `app_id`、`limit_type`、`limit_value`、`window_seconds`
- **详细文档**: [rate_limits.md](./rate_limits.md)

### 8. `clients.token_usage_logs` - 令牌使用日志表
- **用途**: 记录每次令牌使用
- **关键字段**: `token_id`、`app_id`、`endpoint`、`status_code`
- **详细文档**: [token_usage_logs.md](./token_usage_logs.md)

## 应用注册流程 / Application Registration Flow

### 1. 应用注册
```
1. 管理员创建应用（apps 表）
2. 设置应用类型、环境、状态
3. 应用状态为 pending，等待审核
```

### 2. 凭证创建
```
1. 管理员为应用创建 OAuth 凭证（client_credentials）
2. 或创建 API Key（api_keys）
3. 系统返回凭证（明文，只返回一次）
```

### 3. Scope 授权
```
1. 管理员为应用授权 scope（client_scopes）
2. 应用只能请求被授权的 scope
```

### 4. IP 白名单配置
```
1. 管理员配置 IP 白名单（ip_allowlist）
2. 只有来自白名单 IP 的请求才能使用凭证
```

### 5. 速率限制配置
```
1. 管理员配置速率限制（rate_limits）
2. 应用需要遵守速率限制
```

## OAuth 2.0 流程 / OAuth 2.0 Flow

### 1. Client Credentials Grant
```
1. 应用使用 client_id 和 client_secret 请求令牌
2. 系统验证 client_credentials
3. 检查 ip_allowlist
4. 检查 client_scopes（允许的 scope）
5. 签发 service_tokens
6. 记录 token_usage_logs
```

### 2. Scope 验证
```
1. 应用请求 scope: "users.read,users.write"
2. 系统查询 client_scopes，允许的 scope: "users.read,users.count"
3. 最终有效 scope: "users.read"（交集）
```

## 安全特性 / Security Features

### 1. 凭证安全
- **哈希存储**: 永远不存储明文凭证
- **凭证轮换**: 支持定期轮换凭证
- **过期时间**: 支持设置凭证过期时间

### 2. 访问控制
- **IP 白名单**: 限制请求来源 IP
- **Scope 限制**: 限制应用可以请求的 scope
- **速率限制**: 防止滥用

### 3. 审计追踪
- **使用日志**: 记录每次令牌使用
- **状态追踪**: 追踪凭证和令牌状态
- **撤销记录**: 记录撤销原因

## 相关文档 / Related Documentation

- [apps.md](./apps.md) - 应用表详细文档
- [client_credentials.md](./client_credentials.md) - 客户端凭证表详细文档
- [api_keys.md](./api_keys.md) - API 密钥表详细文档
- [client_scopes.md](./client_scopes.md) - 客户端 Scope 表详细文档
- [ip_allowlist.md](./ip_allowlist.md) - IP 白名单表详细文档
- [service_tokens.md](./service_tokens.md) - 服务令牌表详细文档
- [rate_limits.md](./rate_limits.md) - 速率限制表详细文档
- [token_usage_logs.md](./token_usage_logs.md) - 令牌使用日志表详细文档

