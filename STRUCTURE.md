# Project Structure Documentation

This document provides a comprehensive overview of the NebulaID project structure, architecture layers, and module organization.

---

## 📂 Root Directory Structure

```
Identity-Backend/
├── assets/              # Static assets (templates, etc.)
├── atlas/               # Database schema & migrations (Atlas)
├── certs/               # SSL certificates (Let's Encrypt)
├── data/                # User-uploaded data (avatars, backgrounds)
├── enums/               # Generated database enum types
├── events/              # Event definitions for event-driven architecture
├── inputs/              # Service entry points (API, Connection, Pipeline)
├── modules/             # Business modules (Auth, Image)
├── pkgs/                # Shared packages and utilities
├── protos/              # Protocol Buffer definitions
├── static/              # Static files served by HTTP
├── tmp/                 # Temporary build artifacts
├── docker-compose.yml   # Production Docker Compose
├── docker-compose.dev.yml # Development Docker Compose
├── go.mod               # Go module definition
├── go.sum               # Go dependencies checksum
├── Taskfile.yml         # Task runner configuration
└── README.md            # Project documentation
```

---

## 🏛️ Architecture Layers

NebulaID follows **Clean Architecture** with clear separation of concerns:

### Layer Hierarchy

```
┌─────────────────────────────────────────────────────────────┐
│  Interfaces Layer (Presentation)                            │
│  - HTTP handlers (Fiber)                                    │
│  - gRPC handlers                                            │
│  - Event handlers (Kafka consumers)                         │
│  - DTOs (Data Transfer Objects)                             │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  Application Layer (Use Cases)                              │
│  - Commands (CQRS)                                          │
│  - Queries (CQRS)                                           │
│  - Application services                                     │
│  - View models                                              │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  Domain Layer (Business Logic)                              │
│  - Entities                                                 │
│  - Value Objects                                            │
│  - Domain behaviors                                         │
│  - Domain events                                            │
│  - Repository interfaces                                    │
│  - Domain errors                                            │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  Infrastructure Layer (Technical Details)                   │
│  - Database repositories (PostgreSQL)                       │
│  - Cache implementations (Redis)                            │
│  - Event bus (Kafka)                                        │
│  - External service clients (gRPC)                          │
│  - Query builders                                           │
│  - Database models (GORM)                                   │
└─────────────────────────────────────────────────────────────┘
```

---

## 📦 Module Structure

Each business module (e.g., `auth`, `image`) follows the same layered structure:

### Module Directory Layout

```
modules/{module}/
├── application/          # Application layer (use cases)
│   ├── {entity}/         # Per-entity application logic
│   │   ├── commands/     # Command definitions (CQRS)
│   │   ├── queries/      # Query definitions (CQRS)
│   │   ├── views/        # View models
│   │   ├── create.go     # Create use case
│   │   ├── update.go     # Update use case
│   │   ├── delete.go     # Delete use case
│   │   ├── get.go        # Get use case
│   │   └── service.go    # Application service
│   └── ...
├── domain/               # Domain layer (business logic)
│   ├── {entity}/         # Per-entity domain logic
│   │   ├── entity.go     # Domain entity
│   │   ├── behavior.go   # Domain behaviors
│   │   ├── factory.go    # Entity factory
│   │   ├── repo.go       # Repository interface
│   │   ├── validation.go # Domain validation
│   │   ├── errors/       # Domain errors
│   │   └── views/        # Domain view models
│   └── ...
├── infrastructure/       # Infrastructure layer
│   ├── repository/       # Repository implementations
│   │   ├── {entity}_pg_repo.go
│   │   └── mapper/      # Entity mappers
│   ├── query/            # Query implementations (CQRS)
│   │   ├── {entity}_pg_query.go
│   │   └── mapper/      # Query mappers
│   ├── rdb/              # Database models & views
│   │   ├── models/       # GORM models
│   │   └── views/        # Database views
│   └── grpcclient/       # gRPC clients to other services
├── interfaces/           # Interface layer
│   ├── http/             # HTTP handlers
│   │   ├── handler/      # Request handlers
│   │   ├── dto/          # Data Transfer Objects
│   │   ├── middleware/   # HTTP middleware
│   │   ├── router.go     # Route definitions
│   │   └── server.go     # HTTP server setup
│   ├── grpc/             # gRPC handlers
│   │   ├── handler/      # gRPC handlers
│   │   ├── mapper/       # gRPC mappers
│   │   └── server.go     # gRPC server setup
│   └── eventbus/         # Event handlers
│       ├── handler/      # Event handlers
│       ├── registry.go   # Event handler registry
│       └── server.go     # Event bus server
├── config/               # Module configuration
│   ├── config.go         # Config loader
│   └── types.go          # Config types
└── server/               # Server wiring
    ├── server.go         # Server initialization
    └── wiring.go         # Dependency injection
```

---

## 🔧 Modules Overview

### 1. Auth Module (`modules/auth/`)

**Purpose**: Authentication, authorization, user management, and profile management.

**Domain Entities**:
- `user` - User accounts, authentication
- `profile` - User profiles with rich metadata
- `role` - User roles and permissions
- `badge` - Achievement badges
- `profile_badge` - User badge associations
- `education` - Education history
- `occupation` - Occupation history

**Services**:
- **API Service** (`inputs/auth/api/`) - HTTP REST API
- **Connection Service** (`inputs/auth/connection/`) - gRPC service
- **Pipeline Service** (`inputs/auth/pipeline/`) - Kafka event consumers

**Key Features**:
- User registration and login
- JWT token management
- Profile CRUD operations
- Role-based access control
- Badge and achievement system

### 2. Image Module (`modules/image/`)

**Purpose**: Image storage, metadata management, and image type definitions.

**Domain Entities**:
- `image` - Image metadata and storage references
- `image_type` - Image type definitions (avatar, background, etc.)

**Services**:
- **API Service** (`inputs/image/api/`) - HTTP REST API
- **Connection Service** (`inputs/image/connection/`) - gRPC service
- **Pipeline Service** (`inputs/image/pipeline/`) - Kafka event consumers

**Key Features**:
- Image upload and storage
- Image metadata management
- Image type categorization
- Image deletion and cleanup

---

## 📚 Shared Packages (`pkgs/`)

Reusable packages used across all modules:

### Core Infrastructure

- **`cache/`** - Redis caching layer with entity cache, list cache, counter cache
- **`configx/`** - Configuration loader (Koanf wrapper)
- **`env/`** - Environment variable utilities
- **`logx/`** - Structured logging (Zap wrapper)
- **`health/`** - Health check manager

### Data Access

- **`postgresqlx/`** - PostgreSQL connection and health check
- **`mysqlx/`** - MySQL connection (legacy/optional)
- **`mongodbx/`** - MongoDB connection (optional)
- **`query/`** - Query builder utilities (GORM helpers)

### Communication

- **`grpcx/`** - gRPC client/server configuration
- **`eventbus/`** - Event bus abstraction (Kafka via Watermill)
- **`kafkax/`** - Kafka publisher/subscriber (Sarama)

### Security

- **`security/token/`** - JWT token management
  - `usertoken/` - User token verification
  - `servertoken/` - Server-to-server token
- **`security/ratelimit/`** - Rate limiting middleware
- **`tokenx/`** - Token generation and verification utilities

### Utilities

- **`utils/`** - Common utilities
  - `id/` - ID conversion utilities
  - `timex/` - Time utilities
  - `ptr/` - Pointer utilities
  - `slice/` - Slice utilities
  - `mapx/` - Map utilities
  - `typeutil/` - Type utilities
  - `file/` - File utilities
  - `filter/` - Filter utilities
  - `contextx/` - Context utilities
  - `cleanup/` - Cleanup utilities
- **`patch/`** - Field patching utilities
- **`circuitbreaker/`** - Circuit breaker pattern
- **`retry/`** - Retry utilities
- **`recover/`** - Panic recovery middleware
- **`safeexec/`** - Safe goroutine execution
- **`cleanup/`** - Resource cleanup utilities
- **`email/`** - Email templates and SMTP

### Network

- **`netx/httpresp/`** - HTTP response utilities
- **`netx/ssh/`** - SSH tunnel utilities

---

## 🔌 Service Entry Points (`inputs/`)

Each module has three service entry points:

### API Service

**Location**: `inputs/{module}/api/`

- **Purpose**: HTTP REST API server
- **Framework**: Fiber v2
- **Port**: 8080 (configurable)
- **Features**:
  - RESTful endpoints
  - JWT authentication middleware
  - Request validation
  - Error handling
  - CORS support

**Example**: `inputs/auth/api/main.go`

### Connection Service

**Location**: `inputs/{module}/connection/`

- **Purpose**: gRPC service for inter-service communication
- **Framework**: Google gRPC
- **Port**: 10012 (auth), 10013 (image)
- **Features**:
  - gRPC service definitions
  - Server-to-server authentication
  - OpenTelemetry instrumentation

**Example**: `inputs/auth/connection/main.go`

### Pipeline Service

**Location**: `inputs/{module}/pipeline/`

- **Purpose**: Kafka event consumers for async processing
- **Framework**: Watermill + Sarama
- **Features**:
  - Event subscription
  - Event processing
  - Dead letter queue handling
  - Retry logic

**Example**: `inputs/auth/pipeline/main.go`

---

## 🗄️ Database Schema (`atlas/`)

Atlas is used for database schema management and migrations.

### Directory Structure

```
atlas/
├── atlas.hcl              # Atlas configuration
├── src/                   # Source SQL schemas
│   ├── main.sql           # Main schema entry
│   ├── schemas/           # Schema definitions
│   │   ├── auth/          # Auth module schemas
│   │   └── image/         # Image module schemas
│   └── extensions/        # PostgreSQL extensions
├── migrations/            # Generated migrations
│   ├── development/       # Dev migrations
│   └── production/        # Prod migrations
├── gen/                   # Generated code
│   ├── models/            # GORM models
│   ├── enums/             # Enum types
│   └── views/             # Database views
├── scripts/               # Generation scripts
│   ├── gen_models.sh      # Generate models
│   ├── gen_enums.sh       # Generate enums
│   └── gen_views.sh       # Generate views
└── templates/             # Code generation templates
```

### Schema Organization

- **`schemas/auth/`** - Auth module database schemas
  - User tables
  - Profile tables
  - Role and permission tables
  - Badge and achievement tables
- **`schemas/image/`** - Image module database schemas
  - Image metadata tables
  - Image type tables

---

## 📡 Protocol Buffers (`protos/`)

gRPC service definitions and generated code.

### Structure

```
protos/
├── buf.yaml               # Buf configuration
├── buf.gen.yaml           # Code generation config
├── buf.lock               # Dependency lock file
├── src/                   # Source .proto files
│   ├── auth/              # Auth service definitions
│   └── image/             # Image service definitions
└── gen/                   # Generated Go code
    ├── auth/              # Generated auth code
    └── image/             # Generated image code
```

---

## 🎯 Events (`events/`)

Event-driven architecture definitions.

### Files

- **`events.go`** - Event type constants
- **`topics.go`** - Kafka topic definitions
- **`auth.go`** - Auth module events
- **`image.go`** - Image module events

### Event Types

**Auth Events**:
- `auth_to_auth.success` - Internal success events
- `auth_to_auth.user.invalidate_cache` - Cache invalidation
- `auth_to_image.image_delete` - Image deletion requests

**Image Events**:
- `image_to_auth.image_success` - Image operation success
- `image_to_auth.image_delete` - Image deletion notifications

---

## 🐳 Docker Configuration

### Files

- **`docker-compose.yml`** - Production deployment
- **`docker-compose.dev.yml`** - Development deployment

### Services

1. **reverse-proxy** - Traefik reverse proxy
2. **auth-api** - Auth HTTP API service
3. **auth-connection** - Auth gRPC service
4. **auth-pipeline** - Auth Kafka consumer
5. **image-api** - Image HTTP API service
6. **image-connection** - Image gRPC service
7. **image-pipeline** - Image Kafka consumer

---

## 🔄 Data Flow Examples

### User Registration Flow

```
HTTP Request → API Service → Application Layer → Domain Layer
                                                      ↓
                                              Repository (PostgreSQL)
                                                      ↓
                                              Event Publisher (Kafka)
                                                      ↓
                                              Pipeline Service (Consumer)
```

### Inter-Service Communication

```
Service A → gRPC Client → Connection Service → Application Layer
                                                      ↓
                                              Domain Layer
                                                      ↓
                                              Repository
```

### Event-Driven Flow

```
Domain Event → Event Publisher → Kafka Topic
                                         ↓
                              Pipeline Service (Consumer)
                                         ↓
                              Event Handler → Application Layer
```

---

## 📝 Code Generation

The project uses code generation for:

1. **Protocol Buffers** - `task proto:gen`
2. **Database Models** - `task atlas:gen:models`
3. **Database Enums** - `task atlas:gen:enums`
4. **Database Views** - `task atlas:gen:views`

Generated code is placed in:
- `protos/gen/` - Generated protobuf code
- `atlas/gen/` - Generated database code
- `enums/` - Generated enum types

---

## 🎨 Design Patterns

### CQRS (Command Query Responsibility Segregation)

- **Commands**: Write operations (create, update, delete)
- **Queries**: Read operations (get, list, search)
- Separate handlers and models for commands and queries

### Repository Pattern

- Domain layer defines repository interfaces
- Infrastructure layer implements repositories
- Abstraction over data access

### Factory Pattern

- Domain entities use factories for creation
- Ensures valid entity construction

### Event Sourcing (Partial)

- Domain events for important state changes
- Event-driven communication between services

---

## 🔍 Key Files Reference

### Configuration

- `inputs/{module}/config/dev.toml` - Development config
- `inputs/{module}/config/prod.toml` - Production config

### Entry Points

- `inputs/{module}/api/main.go` - API service entry
- `inputs/{module}/connection/main.go` - gRPC service entry
- `inputs/{module}/pipeline/main.go` - Pipeline service entry

### Server Setup

- `modules/{module}/server/server.go` - Server initialization
- `modules/{module}/server/wiring.go` - Dependency injection

### Task Runner

- `Taskfile.yml` - Task definitions for common operations

---

## 📖 Further Reading

- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Domain-Driven Design](https://martinfowler.com/bliki/DomainDrivenDesign.html)
- [CQRS Pattern](https://martinfowler.com/bliki/CQRS.html)
- [Atlas Documentation](https://atlasgo.io/)
- [Fiber Documentation](https://docs.gofiber.io/)
- [Watermill Documentation](https://watermill.io/)

---

This structure ensures maintainability, testability, and scalability while following industry best practices for microservice architecture.

