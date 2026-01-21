# Project Structure Documentation

**NFX-Identity = NebulaForgeX Identity Platform**

This document provides a comprehensive overview of the NFX-Identity project structure, architecture layers, module organization, and API routing.

---

## 📂 Root Directory Structure

```
Identity-Backend/
├── assets/              # Static assets (templates, etc.)
├── databases/               # Database schema & migrations (Atlas)
├── certs/               # SSL certificates (Let's Encrypt)
├── data/                # User-uploaded data (avatars, backgrounds)
├── enums/               # Generated database enum types
├── events/              # Event definitions for event-driven architecture
├── inputs/              # Service entry points (API, Connection, Pipeline)
├── modules/             # Business modules (Auth, Image, Permission)
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

NFX-Identity follows **Clean Architecture** with clear separation of concerns and **CQRS** pattern:

### Layer Hierarchy

```
┌─────────────────────────────────────────────────────────────┐
│  Interfaces Layer (Presentation)                            │
│  - HTTP handlers (Fiber)                                    │
│  - gRPC handlers                                            │
│  - Event handlers (Kafka consumers)                         │
│  - DTOs (Data Transfer Objects)                             │
│  - Middleware (Auth, CORS, Recovery)                        │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  Application Layer (Use Cases)                              │
│  - Commands (CQRS Write Side)                              │
│  - Queries (CQRS Read Side)                                 │
│  - Application services                                     │
│  - View models (Application Views)                          │
│  - Business logic orchestration                             │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  Domain Layer (Business Logic)                              │
│  - Entities                                                 │
│  - Value Objects                                            │
│  - Domain behaviors                                         │
│  - Domain events                                            │
│  - Query interfaces (CQRS Read Side)                        │
│  - Repository interfaces (CQRS Write Side)                  │
│  - Domain errors                                            │
│  - Domain views                                             │
│  - Factories                                                │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  Infrastructure Layer (Technical Details)                   │
│  - Query implementations (single/list packages)              │
│  - Repository implementations (create/get/check/update/delete)│
│  - Database models (GORM)                                   │
│  - Cache implementations (Redis)                            │
│  - Event bus (Kafka)                                        │
│  - External service clients (gRPC)                          │
│  - Mappers (Entity ↔ Model)                                 │
└─────────────────────────────────────────────────────────────┘
```

---

## 📦 Module Structure

Each business module (e.g., `auth`, `image`, `permission`) follows the same layered structure with **CQRS** pattern:

### Module Directory Layout

```
modules/{module}/
├── application/          # Application layer (use cases)
│   ├── {entity}/         # Per-entity application logic
│   │   ├── create.go     # Create use case
│   │   ├── update.go     # Update use case
│   │   ├── delete.go     # Delete use case
│   │   ├── get.go        # Get use case
│   │   ├── list.go       # List use case
│   │   └── service.go    # Application service (orchestrates domain)
│   └── ...
├── domain/               # Domain layer (business logic)
│   ├── {entity}/         # Per-entity domain logic
│   │   ├── entity.go     # Domain entity (core business object)
│   │   ├── behavior.go   # Domain behaviors (business rules)
│   │   ├── factory.go    # Entity factory (ensures valid creation)
│   │   ├── query.go      # Query interface (CQRS Read Side)
│   │   │                  #   - Query struct with Single/List/Count
│   │   │                  #   - Single interface (returns *View)
│   │   │                  #   - List interface (returns []*View)
│   │   │                  #   - ListQuery struct (pagination/sorting)
│   │   ├── repo.go        # Repository interface (CQRS Write Side)
│   │   │                  #   - Repo struct with sub-interfaces:
│   │   │                  #     * Create (New)
│   │   │                  #     * Get (ByID, ByUsername, etc.)
│   │   │                  #     * Check (existence checks)
│   │   │                  #     * Update (Generic, Password, etc.)
│   │   │                  #     * Delete (ByID)
│   │   ├── validation.go # Domain validation rules
│   │   ├── errors/        # Domain-specific errors
│   │   └── views/         # Domain view models (read models)
│   └── ...
├── infrastructure/       # Infrastructure layer
│   ├── query/            # Query implementations (CQRS Read Side)
│   │   ├── {entity}/     # Per-entity query implementations
│   │   │   ├── query.go  # Query factory (creates *domain.Query)
│   │   │   ├── single/   # Single query handlers
│   │   │   │   ├── query.go
│   │   │   │   ├── by_id.go
│   │   │   │   ├── by_username.go
│   │   │   │   └── ...
│   │   │   ├── list/      # List query handlers
│   │   │   │   ├── query.go
│   │   │   │   ├── generic.go
│   │   │   │   └── ...
│   │   │   └── mapper/   # Query mappers (Model → View)
│   │   └── ...
│   ├── repository/       # Repository implementations (CQRS Write Side)
│   │   ├── {entity}/     # Per-entity repository implementations
│   │   │   ├── repo.go   # Repository factory (creates *domain.Repo)
│   │   │   ├── create/   # Create handlers
│   │   │   │   ├── repo.go
│   │   │   │   └── new.go
│   │   │   ├── get/      # Get handlers
│   │   │   │   ├── repo.go
│   │   │   │   ├── by_id.go
│   │   │   │   ├── by_username.go
│   │   │   │   └── ...
│   │   │   ├── check/    # Check handlers
│   │   │   │   ├── repo.go
│   │   │   │   ├── by_id.go
│   │   │   │   └── ...
│   │   │   ├── update/   # Update handlers
│   │   │   │   ├── repo.go
│   │   │   │   ├── generic.go
│   │   │   │   ├── password.go
│   │   │   │   └── ...
│   │   │   └── delete/   # Delete handlers
│   │   │       ├── repo.go
│   │   │       └── by_id.go
│   │   └── mapper/       # Entity mappers (Entity ↔ Model)
│   ├── rdb/              # Database models & views
│   │   ├── models/       # GORM models (write models)
│   │   └── views/        # Database views (read models)
│   └── grpcclient/       # gRPC clients to other services
├── interfaces/           # Interface layer
│   ├── http/             # HTTP handlers
│   │   ├── handler/      # Request handlers (per entity)
│   │   ├── dto/          # Data Transfer Objects (request/response)
│   │   ├── middleware/   # HTTP middleware (auth, CORS, etc.)
│   │   ├── router.go     # Route definitions
│   │   └── server.go     # HTTP server setup (Fiber app)
│   ├── grpc/             # gRPC handlers
│   │   ├── handler/      # gRPC handlers (per service method)
│   │   ├── mapper/       # gRPC mappers (Domain ↔ Protobuf)
│   │   └── server.go     # gRPC server setup
│   └── eventbus/         # Event handlers
│       ├── handler/      # Event handlers (per event type)
│       ├── registry.go   # Event handler registry
│       └── server.go     # Event bus server (Kafka consumer)
├── config/               # Module configuration
│   ├── config.go         # Config loader
│   └── types.go          # Config types
└── server/               # Server wiring
    ├── server.go         # Server initialization (HTTP/gRPC/Pipeline)
    └── wiring.go         # Dependency injection (DI container)
```

---

## 🎯 CQRS Pattern Implementation

### Query Layer (Read Side)

**Domain Interface** (`domain/{entity}/query.go`):

```go
type Query struct {
    Single Single
    List   List
    Count  Count  // Optional
}

type Single interface {
    ByID(ctx context.Context, id uuid.UUID) (*views.EntityView, error)
    ByUsername(ctx context.Context, username string) (*views.EntityView, error)
    // Other single-return methods
}

type List interface {
    Generic(ctx context.Context, q ListQuery) ([]*views.EntityView, int64, error)
    // Other list-return methods
}

type Count interface {
    All(ctx context.Context) (int64, error)
}
```

**Infrastructure Implementation** (`infrastructure/query/{entity}/`):

- `query.go` - Factory that creates `*domain.Query` with all implementations
- `single/` - Handlers for `Single` interface (return `*View`)
  - `by_id.go`, `by_username.go`, etc.
- `list/` - Handlers for `List` interface (return `[]*View`)
  - `generic.go` - Generic list with pagination/sorting/filtering
- `mapper/` - Mappers that convert database models to domain views

### Repository Layer (Write Side)

**Domain Interface** (`domain/{entity}/repo.go`):

```go
type Repo struct {
    Create Create
    Get    Get
    Check  Check
    Update Update
    Delete Delete
}

type Create interface {
    New(ctx context.Context, e *Entity) error
}

type Get interface {
    ByID(ctx context.Context, id uuid.UUID) (*Entity, error)
    ByUsername(ctx context.Context, username string) (*Entity, error)
    // Other get methods
}

type Check interface {
    ByID(ctx context.Context, id uuid.UUID) (bool, error)
    ByUsername(ctx context.Context, username string) (bool, error)
    // Other check methods
}

type Update interface {
    Generic(ctx context.Context, e *Entity) error
    Password(ctx context.Context, userID uuid.UUID, hashedPassword string) error
    // Other update methods
}

type Delete interface {
    ByID(ctx context.Context, id uuid.UUID) error
}
```

**Infrastructure Implementation** (`infrastructure/repository/{entity}/`):

- `repo.go` - Factory that creates `*domain.Repo` with all implementations
- `create/` - Handlers for `Create` interface
  - `new.go` - Creates new entity
- `get/` - Handlers for `Get` interface
  - `by_id.go`, `by_username.go`, etc.
- `check/` - Handlers for `Check` interface
  - `by_id.go`, `by_username.go`, etc.
- `update/` - Handlers for `Update` interface
  - `generic.go` - Generic update
  - `password.go` - Password-specific update
- `delete/` - Handlers for `Delete` interface
  - `by_id.go` - Delete by ID
- `mapper/` - Mappers that convert between domain entities and database models

---

## 🔧 Modules Overview

### 1. Auth Module (`modules/auth/`)

**Purpose**: Authentication, authorization, user management, and profile management.

**Domain Entities**:
- `user` - User accounts, authentication
- `profile` - User profiles with rich metadata
- `role` - User roles
- `user_role` - User-role associations
- `badge` - Achievement badges
- `profile_badge` - User badge associations
- `profile_education` - Education history
- `profile_occupation` - Occupation history

**Services**:
- **API Service** (`inputs/auth/api/`) - HTTP REST API
- **Connection Service** (`inputs/auth/connection/`) - gRPC service
- **Pipeline Service** (`inputs/auth/pipeline/`) - Kafka event consumers

**HTTP API Routes** (`/auth/*`):

**Public Routes** (no authentication required):
- `POST /auth/login` - User login
- `POST /auth/refresh` - Refresh access token
- `POST /auth/verification-code` - Send verification code

**Protected Routes** (requires JWT token):
- **Users**:
  - `POST /auth/users` - Create user
  - `GET /auth/users` - List users
  - `GET /auth/users/:id` - Get user by ID
  - `PUT /auth/users/:id` - Update user
  - `DELETE /auth/users/:id` - Delete user
  - `DELETE /auth/users/:id/account` - Delete user account
- **Profiles**:
  - `POST /auth/profiles` - Create profile
  - `GET /auth/profiles` - List profiles
  - `GET /auth/profiles/:id` - Get profile by ID
  - `GET /auth/profiles/user/:user_id` - Get profile by user ID
  - `PUT /auth/profiles/:id` - Update profile
  - `DELETE /auth/profiles/:id` - Delete profile
- **Roles**:
  - `POST /auth/roles` - Create role
  - `GET /auth/roles` - List roles
  - `GET /auth/roles/:id` - Get role by ID
  - `GET /auth/roles/name/:name` - Get role by name
  - `PUT /auth/roles/:id` - Update role
  - `DELETE /auth/roles/:id` - Delete role
- **Badges**:
  - `POST /auth/badges` - Create badge
  - `GET /auth/badges` - List badges
  - `GET /auth/badges/:id` - Get badge by ID
  - `GET /auth/badges/name/:name` - Get badge by name
  - `PUT /auth/badges/:id` - Update badge
  - `DELETE /auth/badges/:id` - Delete badge
- **Educations**:
  - `POST /auth/educations` - Create education
  - `GET /auth/educations` - List educations
  - `GET /auth/educations/:id` - Get education by ID
  - `GET /auth/educations/profile/:profile_id` - Get educations by profile ID
  - `PUT /auth/educations/:id` - Update education
  - `DELETE /auth/educations/:id` - Delete education
- **Occupations**:
  - `POST /auth/occupations` - Create occupation
  - `GET /auth/occupations` - List occupations
  - `GET /auth/occupations/:id` - Get occupation by ID
  - `GET /auth/occupations/profile/:profile_id` - Get occupations by profile ID
  - `PUT /auth/occupations/:id` - Update occupation
  - `DELETE /auth/occupations/:id` - Delete occupation
- **Profile Badges**:
  - `POST /auth/profile-badges` - Create profile-badge association
  - `GET /auth/profile-badges/:id` - Get profile-badge by ID
  - `GET /auth/profile-badges/profile/:profile_id` - Get badges by profile ID
  - `GET /auth/profile-badges/badge/:badge_id` - Get profiles by badge ID
  - `PUT /auth/profile-badges/:id` - Update profile-badge
  - `DELETE /auth/profile-badges/:id` - Delete profile-badge
  - `DELETE /auth/profile-badges/profile/:profile_id/badge/:badge_id` - Delete by profile and badge

**Key Features**:
- User registration and login
- JWT token management (access + refresh tokens)
- Profile CRUD operations
- Role-based access control
- Badge and achievement system
- Education and occupation history

### 2. Image Module (`modules/image/`)

**Purpose**: Image storage, metadata management, and image type definitions.

**Domain Entities**:
- `image` - Image metadata and storage references
- `image_type` - Image type definitions (avatar, background, etc.)

**Services**:
- **API Service** (`inputs/image/api/`) - HTTP REST API
- **Connection Service** (`inputs/image/connection/`) - gRPC service
- **Pipeline Service** (`inputs/image/pipeline/`) - Kafka event consumers

**HTTP API Routes** (`/image/*`):

**All Routes** (no authentication required by default):
- **Images**:
  - `POST /image/images` - Create/upload image
  - `GET /image/images` - List images
  - `GET /image/images/:id` - Get image by ID
  - `PUT /image/images/:id` - Update image
  - `DELETE /image/images/:id` - Delete image
- **Image Types**:
  - `POST /image/image-types` - Create image type
  - `GET /image/image-types` - List image types
  - `GET /image/image-types/:id` - Get image type by ID
  - `GET /image/image-types/key/:key` - Get image type by key
  - `PUT /image/image-types/:id` - Update image type
  - `DELETE /image/image-types/:id` - Delete image type

**Key Features**:
- Image upload and storage
- Image metadata management
- Image type categorization
- Image deletion and cleanup

### 3. Permission Module (`modules/permission/`)

**Purpose**: Permission management, user-permission associations, and admin authentication.

**Domain Entities**:
- `permission` - Permission definitions
- `user_permission` - User-permission associations
- `authorization_code` - Authorization codes for employee registration

**Services**:
- **API Service** (`inputs/permission/api/`) - HTTP REST API
- **Connection Service** (`inputs/permission/connection/`) - gRPC service
- **Pipeline Service** (`inputs/permission/pipeline/`) - Kafka event consumers

**HTTP API Routes** (`/permission/*`):

**Public Routes** (no authentication required):
- `POST /permission/login` - Admin login (returns permissions)
- `POST /permission/register` - User registration with authorization code

**Protected Routes** (requires JWT token):
- **Permissions**:
  - `POST /permission/permissions` - Create permission
  - `PUT /permission/permissions/:id` - Update permission
  - `DELETE /permission/permissions/:id` - Delete permission
  - `GET /permission/permissions/:id` - Get permission by ID
  - `GET /permission/permissions/tag/:tag` - Get permission by tag
  - `GET /permission/permissions` - List permissions
- **User Permissions**:
  - `POST /permission/user-permissions` - Assign permission to user
  - `DELETE /permission/user-permissions` - Revoke permission from user
  - `GET /permission/users/:user_id/permissions` - Get permissions by user ID
  - `GET /permission/users/:user_id/permission-tags` - Get permission tags by user ID
  - `POST /permission/user-permissions/check` - Check if user has permission
- **Authorization Codes**:
  - `POST /permission/authorization-codes` - Create authorization code
  - `GET /permission/authorization-codes/:id` - Get authorization code by ID
  - `GET /permission/authorization-codes/code/:code` - Get authorization code by code
  - `POST /permission/authorization-codes/use` - Use authorization code
  - `DELETE /permission/authorization-codes/:id` - Delete authorization code
  - `POST /permission/authorization-codes/:id/activate` - Activate authorization code
  - `POST /permission/authorization-codes/:id/deactivate` - Deactivate authorization code

**Key Features**:
- Permission CRUD operations
- User permission assignment/revocation
- Permission checking
- Permission categorization (tags)
- Authorization code system for employee registration
- Admin login with permission information

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
  - `DomainPagination` - Pagination support
  - `DomainSorts` - Sorting support
  - `ExecuteQuery` - Query execution helper

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
- **Port**: Configurable (default: 8080 for auth, varies by module)
- **Features**:
  - RESTful endpoints
  - JWT authentication middleware (for protected routes)
  - Request validation
  - Error handling
  - CORS support
  - Panic recovery middleware
  - Request logging

**Example**: `inputs/auth/api/main.go`

### Connection Service

**Location**: `inputs/{module}/connection/`

- **Purpose**: gRPC service for inter-service communication
- **Framework**: Google gRPC
- **Port**: Configurable (default: 10012 for auth, 10013 for image, 10014 for permission)
- **Features**:
  - gRPC service definitions (from protobuf)
  - Server-to-server authentication
  - OpenTelemetry instrumentation
  - Interceptor support

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
  - Event handler registry

**Example**: `inputs/auth/pipeline/main.go`

---

## 🗄️ Database Schema (`databases/`)

Atlas is used for database schema management and migrations.

### Directory Structure

```
databases/
├── src/                   # Source SQL schemas
│   ├── main.sql           # Main schema entry
│   ├── schemas/           # Schema definitions
│   │   ├── auth/          # Auth module schemas
│   │   ├── image/         # Image module schemas
│   │   └── permission/    # Permission module schemas
│   └── extensions/        # PostgreSQL extensions
├── migrations/            # Generated migrations
│   ├── development/       # Dev migrations
│   └── production/        # Prod migrations
├── templates/              # Code generation templates
└── scripts/                # Generation scripts
```

### Schema Organization

- **`schemas/auth/`** - Auth module database schemas
  - User tables
  - Profile tables
  - Role and permission tables
  - Badge and achievement tables
  - Education and occupation tables
- **`schemas/image/`** - Image module database schemas
  - Image metadata tables
  - Image type tables
- **`schemas/permission/`** - Permission module database schemas
  - Permission tables
  - User permission tables
  - Authorization code tables

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
│   ├── image/             # Image service definitions
│   └── permission/        # Permission service definitions
└── gen/                   # Generated Go code
    ├── auth/              # Generated auth code
    ├── image/              # Generated image code
    └── permission/         # Generated permission code
```

---

## 🎯 Events (`events/`)

Event-driven architecture definitions.

### Files

- **`events.go`** - Event type constants
- **`topics.go`** - Kafka topic definitions
- **`auth.go`** - Auth module events
- **`image.go`** - Image module events
- **`permission.go`** - Permission module events

### Event Types

**Auth Events**:
- `auth_to_auth.user.created` - User creation
- `auth_to_auth.profile.updated` - Profile update
- `auth_to_auth.user.invalidate_cache` - Cache invalidation
- `auth_to_image.image_delete` - Image deletion requests

**Image Events**:
- `image_to_auth.image_success` - Image operation success
- `image_to_auth.image_delete` - Image deletion notifications

**Permission Events**:
- `permission_to_auth.permission.assigned` - Permission assigned
- `permission_to_auth.permission.revoked` - Permission revoked

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
8. **permission-api** - Permission HTTP API service
9. **permission-connection** - Permission gRPC service
10. **permission-pipeline** - Permission Kafka consumer

---

## 🔄 Data Flow Examples

### User Registration Flow (CQRS Write)

```
HTTP Request → API Service → Application Layer (Command)
                                          ↓
                                    Domain Layer
                                          ↓
                                    Repository (Create)
                                          ↓
                                    PostgreSQL
                                          ↓
                                    Event Publisher (Kafka)
                                          ↓
                                    Pipeline Service (Consumer)
```

### User Query Flow (CQRS Read)

```
HTTP Request → API Service → Application Layer (Query)
                                          ↓
                                    Domain Layer (Query Interface)
                                          ↓
                                    Query Handler (Single/List)
                                          ↓
                                    PostgreSQL (Read)
                                          ↓
                                    Cache (Redis) - Optional
```

### Inter-Service Communication

```
Service A → gRPC Client → Connection Service → Application Layer
                                                      ↓
                                              Domain Layer
                                                      ↓
                                              Repository/Query
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

1. **Protocol Buffers** - `task proto:gen` or `task gen-proto`
2. **Database Models** - `task atlas:gen:models` (via Atlas pipeline)
3. **Database Enums** - `task atlas:gen:enums` (via Atlas pipeline)
4. **Database Views** - `task atlas:gen:views` (via Atlas pipeline)

Generated code is placed in:
- `protos/gen/` - Generated protobuf code
- `databases/gen/` - Generated database code
- `enums/` - Generated enum types

---

## 🎨 Design Patterns

### CQRS (Command Query Responsibility Segregation)

- **Commands (Write Side)**: Repository interfaces and implementations
  - `Create`, `Update`, `Delete` operations
  - Return domain entities or errors
  - Write to database models (GORM)
- **Queries (Read Side)**: Query interfaces and implementations
  - `Single` interface - Returns single objects (`*View`)
  - `List` interface - Returns arrays (`[]*View`)
  - `ListQuery` - Pagination, sorting, and filtering support
  - Read from database views or models
- Separate handlers and models for commands and queries

### Repository Pattern

- Domain layer defines repository interfaces as structured `Repo` with sub-interfaces
- Infrastructure layer implements repositories in separate packages
- Abstraction over data access
- Clear separation of concerns (Create, Get, Check, Update, Delete)

### Factory Pattern

- Domain entities use factories for creation
- Ensures valid entity construction
- Query and Repository factories in infrastructure layer

### Event Sourcing (Partial)

- Domain events for important state changes
- Event-driven communication between services
- Kafka-based event bus

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

### Domain Layer

- `modules/{module}/domain/{entity}/query.go` - Query interface (CQRS Read)
- `modules/{module}/domain/{entity}/repo.go` - Repository interface (CQRS Write)
- `modules/{module}/domain/{entity}/entity.go` - Domain entity
- `modules/{module}/domain/{entity}/factory.go` - Entity factory

### Infrastructure Layer

- `modules/{module}/infrastructure/query/{entity}/query.go` - Query factory
- `modules/{module}/infrastructure/repository/{entity}/repo.go` - Repository factory

### HTTP Layer

- `modules/{module}/interfaces/http/router.go` - Route definitions
- `modules/{module}/interfaces/http/server.go` - HTTP server setup
- `modules/{module}/interfaces/http/handler/` - Request handlers

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

This structure ensures maintainability, testability, and scalability while following industry best practices for microservice architecture with CQRS pattern.
