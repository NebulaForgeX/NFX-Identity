# NFX-Identity

[中文](README.md)

NebulaForgeX login and profile hub. Go modules: **auth / asset**. Other products have no local account table.

Deploy, full HTTP tables, schemas: [NFX-Documentation chapter 6](https://github.com/NebulaForgeX/NFX-Documentation/blob/main/books/en/chapter-06-nfx-identity-deployment.md). Stack and Edge first. Host gRPC **10200–10201**, console **10203**. Pin **nfx-ui 0.33.0**.

```bash
cp .example.env .env
task proto:gen
task errors:gen-langs
task atlas:pipeline:run
task console:i
sudo docker compose -f docker-compose.dev.yml up --build
```
