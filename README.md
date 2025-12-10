# NebulaID — Unified Authentication & Identity Service

<div align="center">
  <img src="image.png" alt="NebulaStack Logo" width="200">
</div>

**NebulaID** is the centralized identity and user-profile platform of the NebulaForgeX ecosystem.

It provides authentication, authorization, user management, and enriched profile capabilities for all internal applications and microservices.

Built with **Go**, designed with **clean architecture** and **domain-driven principles**, NebulaID is engineered for scalability, API-first integration, and cross-service interoperability.

---

## 🎯 Core Responsibilities

### 1. Authentication & Security

- **Email / password login** - Traditional authentication flow
- **OAuth support** (future: Google, GitHub) - Social login integration
- **JWT issuing & validation** - Secure token-based authentication
- **Refresh token lifecycle** - Long-lived session management
- **Session management** - Active session tracking and control

### 2. Unified User Identity

Acts as the **single source of truth** for:

- User accounts
- Verification state
- Roles & permissions
- Account status (active, suspended, deleted)

### 3. Advanced Profile System

Fully customizable profile domain with:

- **Basic Info**: Nickname, display name
- **Media**: Avatar & background image references
- **Personalization**: Skills, preferences, social links
- **Professional**: Educations / occupations
- **Achievements**: Badges & achievements

Perfect for **Netup、ReX、TrendRadar** 等产品共享的用户体系。

### 4. Service-to-Service Integration (Microservices Ready)

NebulaID 提供：

- **Standardized JWT** for service authentication
- **Internal service tokens** (machine-to-machine)
- **REST + GraphQL** ready
- **Event-driven hooks** (via Kafka):
  - `user.created`
  - `profile.updated`
  - `user.verified`
  - And more...

---

## 🏗️ Architecture

NebulaID follows **Clean Architecture** and **Domain-Driven Design (DDD)** principles:

```
┌─────────────────────────────────────────────────────────┐
│                    Interfaces Layer                      │
│  (HTTP/gRPC Handlers, Event Handlers, DTOs)              │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│                 Application Layer                        │
│  (Use Cases, Commands, Queries, Services)               │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│                   Domain Layer                           │
│  (Entities, Value Objects, Domain Logic, Repositories)   │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│              Infrastructure Layer                        │
│  (Database, Cache, Event Bus, External Services)         │
└─────────────────────────────────────────────────────────┘
```

### Service Architecture

Each module (Auth, Image) is split into three services:

1. **API Service** - HTTP REST API (Fiber)
2. **Connection Service** - gRPC service for inter-service communication
3. **Pipeline Service** - Kafka event consumers for async processing

---

## 🛠️ Tech Stack

### Core
- **Language**: Go 1.24.4+
- **Web Framework**: [Fiber v2](https://github.com/gofiber/fiber)
- **gRPC**: Google gRPC with middleware
- **Event Bus**: Kafka (via Watermill)

### Data Layer
- **Primary DB**: PostgreSQL (via GORM + pgx)
- **Cache**: Redis
- **Document Store**: MongoDB (optional)
- **Migration**: Atlas

### Infrastructure
- **Containerization**: Docker & Docker Compose
- **Reverse Proxy**: Traefik v3
- **Logging**: Zap (structured logging)
- **Configuration**: Koanf (TOML, ENV)
- **Observability**: OpenTelemetry

### Security
- **JWT**: golang-jwt/jwt/v5
- **Password Hashing**: bcrypt (via golang.org/x/crypto)
- **Rate Limiting**: Redis-based
- **Circuit Breaker**: Sony gobreaker

---

## 🚀 Quick Start

### Prerequisites

- **Go 1.24.4+** (see [Memory Note](#-important-notes))
- **Docker & Docker Compose**
- **PostgreSQL** (or use Docker Compose from Resources)
- **Redis** (or use Docker Compose from Resources)
- **Kafka** (or use Docker Compose from Resources)

### Installation

1. **Clone the repository**

```bash
git clone <repository-url>
cd Identity-Backend
```

2. **Install Go dependencies**

```bash
go mod download
```

3. **Install development tools**

```bash
# Install Taskfile
go install github.com/go-task/task/v3/cmd/task@latest

# Install all dev tools
task install
```

4. **Set up environment**

```bash
# Copy example config (if exists)
cp inputs/auth/config/dev.toml.example inputs/auth/config/dev.toml
cp inputs/image/config/dev.toml.example inputs/image/config/dev.toml

# Edit config files with your database credentials
```

5. **Run database migrations**

```bash
# Generate migrations from Atlas schemas
task atlas:gen:migrations

# Apply migrations
task atlas:apply:dev
```

6. **Generate code**

```bash
# Generate protobuf code
task proto:gen

# Generate database models/enums
task atlas:gen:models
task atlas:gen:enums
```

7. **Start services**

```bash
# Development mode (with hot reload)
task run:dev

# Or start individual services
cd inputs/auth/api && go run main.go -env=dev
cd inputs/auth/connection && go run main.go -env=dev
cd inputs/auth/pipeline && go run main.go -env=dev
```

### Docker Deployment

```bash
# Build and start all services
docker-compose up -d

# View logs
docker-compose logs -f auth-api

# Stop services
docker-compose down
```

---

## 📁 Project Structure

See [STRUCTURE.md](./STRUCTURE.md) for detailed project structure documentation.

### Key Directories

- `modules/` - Business modules (auth, image)
- `pkgs/` - Shared packages and utilities
- `protos/` - Protocol Buffer definitions
- `atlas/` - Database schema and migrations
- `inputs/` - Service entry points (API, Connection, Pipeline)
- `events/` - Event definitions for event-driven architecture

---

## 🔄 Event-Driven Architecture

NebulaID uses Kafka for asynchronous event processing:

### Event Topics

- `auth` - Auth module events
- `image` - Image module events
- `auth_poison` / `image_poison` - Dead letter queues

### Key Events

**Auth Events:**
- `auth_to_auth.user.created`
- `auth_to_auth.profile.updated`
- `auth_to_auth.user.invalidate_cache`
- `auth_to_image.image_delete`

**Image Events:**
- `image_to_auth.image_success`
- `image_to_auth.image_delete`

See `events/` directory for complete event definitions.

---

## 🧪 Development

### Running Tests

```bash
go test ./...
```

### Code Generation

```bash
# Generate protobuf
task proto:gen

# Generate database code
task atlas:gen:models
task atlas:gen:enums
task atlas:gen:views
```

### Database Migrations

```bash
# Create new migration
task atlas:migrate:create -- <migration_name>

# Apply migrations (dev)
task atlas:apply:dev

# Apply migrations (prod)
task atlas:apply:prod
```

### Linting

```bash
task lint
```

---

## 📦 Deployment

### Production Build

```bash
# Build Docker images
docker-compose build

# Start services
docker-compose up -d
```

### Environment Configuration

- Development: `inputs/*/config/dev.toml`
- Production: `inputs/*/config/prod.toml`

Configuration is loaded via environment variable `ENV=dev|prod`.

---

## 🔐 Security

- **JWT tokens** with configurable expiration
- **Refresh tokens** for long-lived sessions
- **Password hashing** with bcrypt
- **Rate limiting** on authentication endpoints
- **CORS** configuration
- **Input validation** via protobuf validate

---

## 📊 Monitoring & Observability

- **Structured logging** with Zap
- **OpenTelemetry** integration for distributed tracing
- **Health check endpoints** (`/health`)
- **Metrics** (via OpenTelemetry)

---

## 🤝 Integration Guide

### For Other Services

1. **JWT Authentication**: Use tokens issued by NebulaID
2. **gRPC Client**: Connect to Connection services
3. **Event Subscription**: Subscribe to Kafka topics
4. **HTTP Client**: Call REST APIs

### Example: Validating JWT

```go
import "identity/pkgs/security/token/usertoken"

verifier := usertoken.NewVerifier(cfg)
claims, err := verifier.Verify(token)
```

---

## 📝 Important Notes

⚠️ **Go Version Requirement**: This project requires **Go 1.24.4+**. Before compiling, ensure you set the correct Go environment:

```bash
export GOROOT=/opt/bin/go
export PATH="$GOROOT/bin:$PATH"
export PATH="$PATH:$HOME/go/bin"
```

The system default Go 1.21.6 does not meet dependency requirements.

---

## 🗺️ Roadmap

- [ ] OAuth providers (Google, GitHub)
- [ ] GraphQL API
- [ ] Multi-factor authentication (MFA)
- [ ] Advanced role-based access control (RBAC)
- [ ] Audit logging
- [ ] Webhook support

---

## 📄 License

[Your License Here]

---

## 👥 Contributing

[Contributing Guidelines]

---

## 📧 Contact

For questions or support, please contact: [Your Contact Info]

