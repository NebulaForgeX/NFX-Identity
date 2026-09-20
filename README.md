# NFX-Identity

[English](README.en.md)

NebulaForgeX 登录与资料中心。Go 模块：**auth / asset / system**。其它产品没有本地账号表。

部署、HTTP 全表、schema：[NFX-Documentation 第六章](https://github.com/NebulaForgeX/NFX-Documentation/blob/main/books/zh/chapter-06-nfx-identity-deployment.md)。先 Stack 与 Edge。主机 gRPC **10200–10202**，console **10203**。`nfx-ui` **0.29.0**。

```bash
cp .example.env .env
task proto:gen
task errors:gen-langs
task atlas:pipeline:run
task console:i
sudo docker compose -f docker-compose.dev.yml up --build
```
