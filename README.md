# NFX-Identity

[English](README.en.md)

<div align="center">
  <img src="console/public/logo.png" alt="NFX-Identity" width="200">
</div>

NebulaForgeX 的登录与资料中心。两个 Go 模块：**auth / asset**。Edge、News、Storages **没有**本地账号表。它们用同一把 `TOKEN_SECRET_KEY` 验 JWT，并用 gRPC 调 `EnsureOwnedProfile` / `HasForgerRole`。已删除的 tenants / access / directory / clients 表不要再写回来。权威 schema 只有 `auth`、`asset`。第一个 owner 由 `scripts/init.sql` 和 auth gRPC `BootstrapOwner` 写入。

先有 [NFX-Stack](https://github.com/NebulaForgeX/NFX-Stack)（Postgres **10004**、Redis **10006**、MinIO **10012**）和 [NFX-Edge](https://github.com/NebulaForgeX/NFX-Edge)。本仓没有 Traefik labels，也不加入共享网络。compose 服务名是 `auth-base` / `asset-base`。

## kind 和 URL 不是同一个词

`POST /auth/me/select-profile` 的 `kind`，也就是 JWT 的 `profile_scope`，是 **`community` | `authority`**（小写）。`forger` 是 community 档上的角色 `forger_role`，不是 kind。发 `kind: "forger"` 会失败。

Console 的路径树仍然用 `/forger/*` 表示 community、`/authority/*` 表示 authority。走错树时 `ScopeRoute` 打回 `/forger/desk` 或 `/authority/desk`。页面只走 nfx-ui hooks，不要在页面里直调 repository。`nfx-ui` 钉 **0.36.0**。

| 阶段 | 路径 |
|------|------|
| 访客 | `/auth/login`、`/auth/signup` |
| 已登录未选资料 | `/auth/select-profile` |
| community（URL 前缀 forger） | `/forger/desk`、`/forger/profile/overview\|edit\|identity\|security`、`/forger/assets`、`/forger/settings` |
| authority | 同样的 `/authority/*`，另加 `/authority/directory` |

## 端口

容器 HTTP `8080`。容器 gRPC：auth `50071`、asset `50072`。主机不要把 `GRPC_EXT_*` 设成 `50071`。

| | dev | secure |
|--|-----|--------|
| auth HTTP / gRPC | **10030 / 10031** | **10035 / 10036** |
| asset HTTP / gRPC | **10032 / 10033** | **10037 / 10038** |
| console | **10034** | **10039** |

Vite `5173`。Edge 前缀：secure `/nfx-identity`，dev `/dev/nfx-identity`，StripPrefix 之后 Fiber 仍是 `/auth`、`/asset`。域名走 HTTPS。局域网 `Host(NAS1_IP)` 会 302 到控制台端口（dev **10034**，secure **10039**），之后由 console nginx 转发 `/nfx-identity/auth` 和 `/asset`。浏览器：dev `VITE_API_URL=/dev/nfx-identity`、`VITE_BASE=/dev/console/nfx-identity/`；secure 去掉 `/dev`。

其它产品拨本仓：`GRPC_HOST_AUTH` 填这台 NAS 的 IP，端口是 `GRPC_EXT_PORT_AUTH`（dev **10031**，secure **10036**）。它们的 `.env` 里没有 `GRPC_PORT_AUTH`。

## Token

Edge / News / Storages 必须复制同一组。不要把真实密钥写进 Git。

```
TOKEN_SECRET_KEY=<long-random>
TOKEN_ISSUER=nfxidentity
TOKEN_ACCESS_TTL=24h
TOKEN_REFRESH_TTL=168h
TOKEN_ALGORITHM=HS256
```

## HTTP 形状

实现在 `modules/auth/interface/http/router.go` 和 `modules/asset/interface/http/router.go`。

公开：`POST /auth/login/with-email`、`/auth/login/with-phone`、`/auth/signup/send-code`、`/auth/signup/with-email`、`/auth/refresh`、`/auth/logout`，以及 `/auth/health`、`/auth/locales/:lang`、`/auth/messages/:lang`。

Bearer `/auth/me`：`POST /select-profile`；forger 与 authority 各有一份资料、设置、头像、背景、preference；邮箱、手机、密码；`/profiles` 与 `/authority-profiles` 的列表、创建、搜索、删除、公开卡片。

Bearer `/auth/owner`：`GET /forger-profiles`、`GET /authority-profiles`、`PATCH /authority-profiles/:profileId/roles`。

asset 的 `kind` 是 `images | files | videos | audios`。列表和上传要 Token。`GET /asset/{kind}/:id/file` 在路由上公开。对象字节在 Stack MinIO（path-style），不是 Storages。

## 库

`nfxidentity_dev` / `nfxidentity` / shadow `nfxidentity_diff`。表：`Accounts`、`Identities`、`Emails`、`Phones`、`ForgerProfiles`、`AuthorityProfiles`、`RefreshTokens`，以及头像、背景、设置的链接表。asset 四张表：`Images`、`Files`、`Audios`、`Videos`。

枚举：`profile_scope` 是 `community | authority`；`forger_role` 目前只有 `forger`；`authority_role` 是 `auditor | administrator | owner`（默认 `{auditor}`）；`account_status` 是 `active | suspended | deleted`；`signup_platform` 默认 `nfxidentity`，另有 `nfxnews`、`nfxstorages`、`nfxedge`；`profile_language` 是 `en | zh | fr`；`identity_provider` 目前只有 `password`。角色是数组里的成员，不是层级。

```bash
cp .example.env .env
task proto:gen
task errors:gen-langs
task atlas:pipeline:run
task console:i
sudo docker compose -f docker-compose.dev.yml up --build
```

`ENV=secure` 时走 secure 端口。改表先改 `databases/src/**.sql`，再跑 Atlas。不要手改 `protos/gen` 和 `*_dbgen.go`。

详细信息见 [NFX-Documentation 第六章](https://github.com/NebulaForgeX/NFX-Documentation/blob/main/books/zh/chapter-06-nfx-identity-deployment.md)。
