# NFX-Identity — Unified Authentication & Identity Service

**NFX-Identity = NebulaForgeX Identity Platform**

<div align="center">
  <img src="image.png" alt="NFX-Identity Logo" width="200">
</div>

**NFX-Identity** is the centralized identity and user-profile platform of the NebulaForgeX ecosystem.

It provides authentication, authorization, user management, permission management, and enriched profile capabilities for all internal applications and microservices.

Built with **Go**, designed with **Clean Architecture**, **Domain-Driven Design (DDD)**, and **CQRS** principles, NFX-Identity is engineered for scalability, API-first integration, and cross-service interoperability.

---

## 🎯 Service Overview

NFX-Identity consists of three main services, each serving different purposes:

### 1. **Auth Service** - External User API Service

**Purpose**: Provides authentication and user management APIs for **external backend services**.

**Target Users**: 
- Other backend services that need to integrate with NFX-Identity
- Third-party applications requiring user authentication
- Client applications (mobile apps, web apps) that need user management

**Key Features**:
- **User Registration & Login** - Standard authentication flow for external users
- **JWT Token Management** - Issue and validate JWT tokens for external services
- **User Profile Management** - Full CRUD operations for user profiles
- **Role Management** - User roles and role assignment
- **Badge & Achievement System** - User badges and achievements
- **Image Management** - Avatar and profile image handling

**API Endpoints**: `/auth/*`

**Login Flow**: 
- Standard user authentication (username/email/phone + password)
- Returns JWT tokens for client applications
- **Designed for external backend services** to authenticate their users

### 2. **Permission Service** - Internal Admin Console Service

**Purpose**: Provides permission management and admin authentication for **NFX Console (Identity-Admin)** control panel.

**Target Users**: 
- Internal administrators managing the platform
- Company employees using authorization codes to register and access the admin panel
- Backend developers who need to manage permissions

**Key Features**:
- **Admin Login** - Special login flow that includes permission information
- **Permission Management** - Create, update, delete permissions
- **User Permission Assignment** - Assign/revoke permissions to users
- **Authorization Code System** - Generate codes for employee registration
- **Permission Checking** - Verify user permissions

**API Endpoints**: `/permission/*`

**Login Flow**:
- Admin authentication (username/email/phone + password)
- Returns JWT tokens **plus user permissions and permission tags**
- **Designed specifically for NFX Console (Identity-Admin)** frontend panel
- Includes authorization code support for employee self-registration

**Important Note**: The login in Permission service is **different from Auth service login** - it's specifically designed for the admin control panel and includes permission information in the response.

### 3. **Image Service** - Image Storage Service

**Purpose**: Centralized image storage and management for the entire platform.

**Target Users**:
- All services that need image storage (Auth, Permission, etc.)
- User profiles requiring avatars and background images

**Key Features**:
- **Image Upload & Storage** - Store and manage images
- **Image Metadata** - Track image information and references
- **Image Types** - Categorize images (avatar, background, etc.)
- **Image Variants** - Multiple sizes and formats
- **Image Tags** - Flexible tagging system

**API Endpoints**: `/image/*`

---

## 🔑 Key Differences: Auth vs Permission Login

### Auth Service Login (`/auth/login`)
- **Target**: External backend services and client applications
- **Returns**: JWT tokens (access token + refresh token), basic user info
- **Use Case**: Standard user authentication for external services
- **No Permission Info**: Does not include permission details in response

### Permission Service Login (`/permission/login`)
- **Target**: NFX Console (Identity-Admin) control panel
- **Returns**: JWT tokens **plus full permission list and permission tags**
- **Use Case**: Admin panel authentication with permission-aware access control
- **Permission-Aware**: Includes all user permissions in the login response
- **Authorization Code Support**: Supports registration via authorization codes for company employees

---

## 🏗️ Architecture

NFX-Identity follows **Clean Architecture**, **Domain-Driven Design (DDD)**, and **CQRS** principles:

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
│  (Entities, Value Objects, Domain Logic,                │
│   Query Interfaces, Repository Interfaces)              │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│              Infrastructure Layer                        │
│  (Database, Cache, Event Bus, External Services)        │
└─────────────────────────────────────────────────────────┘
```

### CQRS Pattern

The project implements **CQRS (Command Query Responsibility Segregation)**:

- **Query Layer** (`domain/{entity}/query.go`):
  - `Single` interface - Returns single objects (`*View`)
  - `List` interface - Returns arrays (`[]*View`)
  - `ListQuery` - Pagination, sorting, and filtering support

- **Repository Layer** (`domain/{entity}/repo.go`):
  - Structured as `Repo` struct with sub-interfaces:
    - `Create` - Creation operations
    - `Get` - Retrieval operations
    - `Check` - Existence checking
    - `Update` - Update operations
    - `Delete` - Deletion operations

### Service Architecture

Each module (Auth, Image, Permission) is split into three services:

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

- **Go 1.24.4+** (see [Important Notes](#-important-notes))
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
cp inputs/permission/config/dev.toml.example inputs/permission/config/dev.toml

# Edit config files with your database credentials
```

5. **Run database migrations**

```bash
# Run Atlas pipeline (interactive - choose dev or prod)
task atlas:pipeline:run:sh
# Or PowerShell:
task atlas:pipeline:run:ps
```

6. **Generate code**

```bash
# Generate protobuf code
task gen-proto

# Database models are generated automatically during migration pipeline
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

- `modules/` - Business modules (auth, image, permission)
- `pkgs/` - Shared packages and utilities
- `protos/` - Protocol Buffer definitions
- `atlas/` - Database schema and migrations
- `inputs/` - Service entry points (API, Connection, Pipeline)
- `events/` - Event definitions for event-driven architecture

---

## 🔄 Event-Driven Architecture

NFX-Identity uses Kafka for asynchronous event processing:

### Event Topics

- `auth` - Auth module events
- `image` - Image module events
- `permission` - Permission module events
- `auth_poison` / `image_poison` / `permission_poison` - Dead letter queues

### Key Events

**Auth Events:**
- `auth_to_auth.user.created`
- `auth_to_auth.profile.updated`
- `auth_to_auth.user.invalidate_cache`
- `auth_to_image.image_delete`

**Image Events:**
- `image_to_auth.image_success`
- `image_to_auth.image_delete`

**Permission Events:**
- `permission_to_auth.permission.assigned`
- `permission_to_auth.permission.revoked`

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
task gen-proto

# Generate database code (via Atlas pipeline)
task atlas:pipeline:run:sh
```

### Database Migrations

Use the Atlas pipeline for database migrations:

```bash
# Interactive migration pipeline (recommended)
task atlas:pipeline:run:sh
# Or PowerShell:
task atlas:pipeline:run:ps

# The pipeline will:
# 1. Generate migrations from schema changes
# 2. Lint migrations
# 3. Apply migrations
# 4. Generate Go models/views/enums
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
- **Permission-based access control**

---

## 📊 Monitoring & Observability

- **Structured logging** with Zap
- **OpenTelemetry** integration for distributed tracing
- **Health check endpoints** (`/health`)
- **Metrics** (via OpenTelemetry)

---

## 🤝 Integration Guide

### For External Backend Services

Use the **Auth Service** APIs:

1. **User Authentication**: Use `/auth/login` for standard user login
2. **User Management**: Use `/auth/*` endpoints for user operations
3. **JWT Validation**: Validate tokens issued by Auth service
4. **gRPC Client**: Connect to Connection services for inter-service communication

### For NFX Console (Identity-Admin)

Use the **Permission Service** APIs:

1. **Admin Login**: Use `/permission/login` for admin authentication (returns permissions)
2. **Permission Management**: Use `/permission/permissions/*` for permission CRUD
3. **User Permissions**: Use `/permission/user-permissions/*` for permission assignment
4. **Authorization Codes**: Use `/permission/authorization-codes/*` for code management

### Example: External Service Login (Auth Service)

```go
// POST /auth/login
{
  "identifier": "user@example.com",  // or username, or phone
  "password": "password123"
}

// Response:
{
  "access_token": "...",
  "refresh_token": "...",
  "user": { ... }
}
```

### Example: Admin Panel Login (Permission Service)

```go
// POST /permission/login
{
  "identifier": "admin@company.com",
  "password": "password123"
}

// Response:
{
  "access_token": "...",
  "refresh_token": "...",
  "user_id": "...",
  "permissions": [ ... ],  // Full permission list
  "permission_tags": [ "admin.access", "user.manage", ... ]  // Permission tags
}
```

### Example: Validating JWT

```go
import "nfxid/pkgs/security/token/usertoken"

verifier := usertoken.NewVerifier(cfg)
claims, err := verifier.Verify(token)
```

### Example: Checking Permissions (gRPC)

```go
// Via gRPC
client := permissionGRPCClient
hasPermission, err := client.CheckPermission(ctx, &pb.CheckPermissionRequest{
    UserID: userID,
    Tag:    "admin.access",
})
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
- [ ] Permission caching optimization

---

## 📖 Service Summary Table

| Service | Purpose | Target Users | Login Endpoint | Login Response Includes |
|---------|---------|--------------|----------------|------------------------|
| **Auth** | External user API | External backends, client apps | `/auth/login` | JWT tokens, basic user info |
| **Permission** | Admin console API | NFX Console (Identity-Admin) | `/permission/login` | JWT tokens, user info, **permissions**, **permission tags** |
| **Image** | Image storage | All services | N/A (no login) | N/A |

---

## 🔗 Related Projects

- **Identity-Admin** (NFX Console) - Admin control panel frontend that uses Permission service
- **Resources** - Shared Docker Compose services (PostgreSQL, Redis, Kafka, etc.)

