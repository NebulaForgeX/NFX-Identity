# NFX-Identity

[中文](README.md)

<div align="center">
  <img src="console/public/logo.png" alt="NFX-Identity" width="200">
</div>

The NebulaForgeX login and profile hub. Two Go modules: **auth / asset**. Edge, News, and Storages have **no** local account table. They verify JWTs with the same `TOKEN_SECRET_KEY` and call gRPC `EnsureOwnedProfile` / `HasForgerRole`. Do not bring back the deleted tenants / access / directory / clients tables. The only schemas are `auth` and `asset`. The first owner is written by `scripts/init.sql` and auth gRPC `BootstrapOwner`.

[NFX-Stack](https://github.com/NebulaForgeX/NFX-Stack) (Postgres **10004**, Redis **10006**, MinIO **10012**) and [NFX-Edge](https://github.com/NebulaForgeX/NFX-Edge) come first. This repo has no Traefik labels and joins no shared network. Compose service names are `auth-base` and `asset-base`.

## kind and the URL are different words

`POST /auth/me/select-profile` `kind`, which is the JWT `profile_scope`, is **`community` | `authority`** (lowercase). `forger` is the `forger_role` on a community profile, not a kind. Sending `kind: "forger"` fails.

The console path tree still uses `/forger/*` for community and `/authority/*` for authority. `ScopeRoute` sends the wrong tree back to `/forger/desk` or `/authority/desk`. Pages use nfx-ui hooks only. Do not call a repository from a page. Pin **nfx-ui 0.36.0**.

| Stage | Paths |
|-------|--------|
| Guest | `/auth/login`, `/auth/signup` |
| Signed in, no profile yet | `/auth/select-profile` |
| community (URL prefix forger) | `/forger/desk`, `/forger/profile/overview\|edit\|identity\|security`, `/forger/assets`, `/forger/settings` |
| authority | the same under `/authority/*`, plus `/authority/directory` |

## Ports

Container HTTP is `8080`. Container gRPC is auth `50071` and asset `50072`. Do not set `GRPC_EXT_*` to `50071` on the host.

| | dev | secure |
|--|-----|--------|
| auth HTTP / gRPC | **10030 / 10031** | **10035 / 10036** |
| asset HTTP / gRPC | **10032 / 10033** | **10037 / 10038** |
| console | **10034** | **10039** |

Vite is `5173`. Edge prefix: secure `/nfx-identity`, dev `/dev/nfx-identity`. After StripPrefix, Fiber still serves `/auth` and `/asset`. The domain stays on HTTPS. LAN `Host(NAS1_IP)` returns 302 to the console port (dev **10034**, secure **10039**), and the console nginx then proxies `/nfx-identity/auth` and `/asset`. Browser: dev `VITE_API_URL=/dev/nfx-identity`, `VITE_BASE=/dev/console/nfx-identity/`; secure drops `/dev`.

Other products dial this repo with `GRPC_HOST_AUTH` set to this NAS IP and `GRPC_EXT_PORT_AUTH` (dev **10031**, secure **10036**). Their `.env` files have no `GRPC_PORT_AUTH`.

## Tokens

Edge, News, and Storages must copy this set. Do not put a real secret in Git.

```
TOKEN_SECRET_KEY=<long-random>
TOKEN_ISSUER=nfxidentity
TOKEN_ACCESS_TTL=24h
TOKEN_REFRESH_TTL=168h
TOKEN_ALGORITHM=HS256
```

## HTTP shape

Implemented in `modules/auth/interface/http/router.go` and `modules/asset/interface/http/router.go`.

Public: `POST /auth/login/with-email`, `/auth/login/with-phone`, `/auth/signup/send-code`, `/auth/signup/with-email`, `/auth/refresh`, `/auth/logout`, plus `/auth/health`, `/auth/locales/:lang`, `/auth/messages/:lang`.

Bearer `/auth/me`: `POST /select-profile`; forger and authority each have profile, settings, avatar, backgrounds, and preference; emails, phones, and password; list, create, search, delete, and public card for `/profiles` and `/authority-profiles`.

Bearer `/auth/owner`: `GET /forger-profiles`, `GET /authority-profiles`, `PATCH /authority-profiles/:profileId/roles`.

asset `kind` is `images | files | videos | audios`. List and upload need a token. `GET /asset/{kind}/:id/file` is public at the router. Bytes live in Stack MinIO (path-style), not Storages.

## Database

`nfxidentity_dev` / `nfxidentity` / shadow `nfxidentity_diff`. Tables: `Accounts`, `Identities`, `Emails`, `Phones`, `ForgerProfiles`, `AuthorityProfiles`, `RefreshTokens`, plus avatar, background, and settings link tables. asset has `Images`, `Files`, `Audios`, `Videos`.

Enums: `profile_scope` is `community | authority`; `forger_role` is only `forger` today; `authority_role` is `auditor | administrator | owner` (default `{auditor}`); `account_status` is `active | suspended | deleted`; `signup_platform` defaults to `nfxidentity` and also has `nfxnews`, `nfxstorages`, `nfxedge`; `profile_language` is `en | zh | fr`; `identity_provider` is only `password` today. A role is membership in an array, not a rank.

```bash
cp .example.env .env
task proto:gen
task errors:gen-langs
task atlas:pipeline:run
task console:i
sudo docker compose -f docker-compose.dev.yml up --build
```

`ENV=secure` uses the secure ports. Change `databases/src/**.sql` first, then run Atlas. Do not hand-edit `protos/gen` or `*_dbgen.go`.

Full detail: [NFX-Documentation chapter 6](https://github.com/NebulaForgeX/NFX-Documentation/blob/main/books/en/chapter-06-nfx-identity-deployment.md).
