# NFX-Identity

[English](README.en.md)

NebulaForgeX 登录与资料中心。Go 模块：**auth / asset**。其它产品没有本地账号表。

部署、HTTP 全表、schema：[NFX-Documentation 第六章](https://github.com/NebulaForgeX/NFX-Documentation/blob/main/books/zh/chapter-06-nfx-identity-deployment.md)。先 Stack 与 Edge。dev：auth HTTP/gRPC **10030/10031**，console **10034**。secure 从 **10035** 起。`nfx-ui` **0.33.0**。

```bash
cp .example.env .env
task proto:gen
task errors:gen-langs
task atlas:pipeline:run
task console:i
sudo docker compose -f docker-compose.dev.yml up --build
```
