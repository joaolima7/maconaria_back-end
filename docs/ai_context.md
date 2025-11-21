# AI Context - Maçonaria Back-End

## 📌 Project Overview

**Project Name**: Maçonaria Back-End  
**Language**: Go 1.23.2  
**Architecture**: Clean Architecture (Hexagonal)  
**Database**: MySQL 8.0  
**Framework**: Chi Router v5  
**Authentication**: JWT (golang-jwt/jwt v5)

---

## 🎯 Purpose

RESTful API backend for a Masonic organization management system. Handles users, posts, workers, timelines, acacias (tributes), and library resources with role-based access control.

---

## 🏗️ Architecture Pattern

### Clean Architecture Layers

```
internal/
├── domain/          # Business logic & entities (innermost)
│   ├── entity/      # Domain models (User, Post, Library, etc)
│   ├── repositories/# Repository interfaces
│   └── usecases/    # Business use cases
├── data/            # Data access layer
│   └── repositories/# Repository implementations
└── infra/           # Infrastructure (outermost)
    ├── database/    # DB connection & SQLC generated code
    ├── web/         # HTTP handlers, routes, middleware
    └── di/          # Dependency injection (Wire)
```

### Dependency Rule
- **Inner layers** NEVER depend on outer layers
- **Outer layers** depend on inner layers via interfaces
- Domain layer is pure business logic (no frameworks)

---

## 🔑 Key Design Patterns

### 1. Repository Pattern
- **Interface** in `domain/repositories/`
- **Implementation** in `data/repositories/`
- Abstracts data access from business logic

```go
// Domain layer - interface
type CreateUserRepository interface {
    CreateUser(user *entity.User) (*entity.User, error)
}

// Data layer - implementation
type CreateUserRepositoryImpl struct {
    queries *db.Queries
}
```

### 2. Use Case Pattern
- Each business operation = 1 use case
- Input/Output DTOs for data transfer
- Orchestrates repository calls

```go
type CreateUserUseCase struct {
    Repository user.CreateUserRepository
}

func (uc *CreateUserUseCase) Execute(input CreateUserInputDTO) (*CreateUserOutputDTO, error)
```

### 3. Dependency Injection (Wire)
- Google Wire for compile-time DI
- All dependencies wired in `internal/infra/di/wire.go`
- Run `wire` to generate `wire_gen.go`

---

## 📊 Database Strategy

### Migration Tool
- **golang-migrate/migrate v4**
- Versioned migrations in `sql/migrations/`
- Format: `000001_description.up.sql` / `000001_description.down.sql`

### Query Generation
- **SQLC** for type-safe SQL
- SQL queries in `sql/queries/`
- Generated Go code in `internal/infra/database/db/`

### Naming Convention
- Tables: plural lowercase (`users`, `libraries`)
- Columns: snake_case (`created_at`, `small_description`)
- Enums: singular (`UsersDegree`, `LibrariesDegree`)

---

## 🔐 Authentication & Authorization

### JWT Strategy
```go
// Token contains: user_id, email, is_admin
type Claims struct {
    UserID  string `json:"user_id"`
    Email   string `json:"email"`
    IsAdmin bool   `json:"is_admin"`
}
```

### Middleware
- `AuthMiddleware.Authenticate` validates JWT
- Injects claims into request context
- Admin-only routes check `is_admin` claim

---

## 📁 Domain Entities

### User
```go
type User struct {
    ID       string
    Name     string
    Email    string
    Password string     // bcrypt hashed
    CIM      string     // Masonic ID
    Degree   UserDegree // apprentice|companion|master
    IsActive bool
    IsAdmin  bool
}
```

### Library
```go
type Library struct {
    ID               string
    Title            string
    SmallDescription string
    Degree           UserDegree
    FileData         []byte // PDF/DOCX (LONGBLOB)
    CoverData        []byte // Image (LONGBLOB)
    Link             string // External URL
}
```

### Shared Enum
```go
type UserDegree string
const (
    DegreeApprentice UserDegree = "apprentice"
    DegreeCompanion  UserDegree = "companion"
    DegreeMaster     UserDegree = "master"
)
```

---

## 🚀 Request Flow

```
HTTP Request
    ↓
[Router] (Chi)
    ↓
[Middleware] (Auth, CORS, Logger)
    ↓
[Handler] (Decode JSON → Validate)
    ↓
[UseCase] (Business Logic)
    ↓
[Repository] (Database Access)
    ↓
[SQLC Generated Code]
    ↓
MySQL Database
    ↓
[Response] (JSON with standard format)
```

### Standard Response Format
```json
{
  "success": true,
  "message": "Operação realizada com sucesso!",
  "data": { ... }
}
```

---

## 🛠️ Error Handling

### Custom Error Types
```go
// Domain layer errors (apperrors package)
- ValidationError    // 400 Bad Request
- NotFoundError      // 404 Not Found
- DuplicateError     // 409 Conflict
- UnauthorizedError  // 401 Unauthorized
- ForbiddenError     // 403 Forbidden
- InternalError      // 500 Internal Server Error
```

### Error Propagation
- Repositories return domain errors
- UseCases validate and return domain errors
- Handlers convert to HTTP responses via `response` package

---

## 📦 Binary Data Handling

### File Storage Strategy
- BLOBs stored in MySQL (LONGBLOB type)
- API uses **Base64 encoding** for transfer
- Optional fields: `file_data`, `cover_data`

```go
// UseCase converts Base64 → []byte
fileData, _ := base64.StdEncoding.DecodeString(input.FileData)

// Response converts []byte → Base64
output.FileData = base64.StdEncoding.EncodeToString(library.FileData)
```

---

## 🔄 Common Operations

### Creating a New Entity (Example: Library)

1. **Migration** (`sql/migrations/000008_library.up.sql`)
2. **Queries** (`sql/queries/library_queries.sql`)
3. **Run SQLC** → generates `internal/infra/database/db/library_queries.sql.go`
4. **Entity** (`internal/domain/entity/library.go`)
5. **Repository Interface** (`internal/domain/repositories/library/`)
6. **Repository Implementation** (`internal/data/repositories/library/`)
7. **Use Cases** (`internal/domain/usecases/library_usecase/`)
8. **Handler** (`internal/infra/web/handlers/library_handler.go`)
9. **Routes** (add to `internal/infra/web/routes/routes.go`)
10. **Wire DI** (add to `internal/infra/di/wire.go`)
11. **Regenerate Wire** (`cd internal/infra/di && wire`)

---

## 🧪 Development Commands

```bash
# Database migrations
make migrate-up      # Apply migrations
make migrate-down    # Rollback last migration
make migrate-create  # Create new migration

# Code generation
make sqlc-generate   # Generate SQLC code
make wire-generate   # Regenerate Wire DI

# Build & Run
make build          # Build binary
make run            # Run application
make docker-up      # Start Docker services
```

---

## 🌐 API Endpoints

### Authentication
- `POST /api/auth/login` - Login (returns JWT)

### Users (Authenticated)
- `POST /api/users` - Create user
- `GET /api/users` - List all users
- `GET /api/users/{id}` - Get user by ID
- `PUT /api/users/{id}` - Update user
- `PATCH /api/users/{id}/password` - Update password

### Libraries (Authenticated)
- `POST /api/libraries` - Create library
- `GET /api/libraries` - List all
- `GET /api/libraries/{id}` - Get by ID
- `GET /api/libraries/degree/{degree}` - Filter by degree
- `PUT /api/libraries/{id}` - Update
- `DELETE /api/libraries/{id}` - Delete

*(Similar patterns for Posts, Workers, Timelines, Acacias)*

---

## 📝 Validation Strategy

### Entity-level Validation
```go
func (u *User) ValidateEmail() error {
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
    if !emailRegex.MatchString(u.Email) {
        return apperrors.NewValidationError("e-mail", "Formato inválido!")
    }
    return nil
}
```

### UseCase Validation
- DTOs have struct tags: `json:"field" validate:"required,email"`
- Business rules validated before repository calls

---

## 🔒 Security Practices

1. **Password Hashing**: bcrypt (cost 10)
2. **JWT Secret**: From environment variable
3. **CORS**: Configured in router
4. **SQL Injection**: Prevented by SQLC (parameterized queries)
5. **No sensitive data** in logs/responses

---

## 📚 Technology Stack

| Layer | Technology | Purpose |
|-------|-----------|---------|
| Router | Chi v5 | HTTP routing |
| ORM | SQLC | Type-safe SQL |
| Migration | golang-migrate | Version control DB |
| DI | Google Wire | Dependency injection |
| Auth | jwt-go v5 | JWT tokens |
| Config | Viper | Environment config |
| Database | MySQL 8.0 | Data persistence |
| Crypto | bcrypt | Password hashing |

---

## 🎨 Code Conventions

### Naming
- **Entities**: PascalCase (`User`, `Library`)
- **Interfaces**: PascalCase + suffix (`CreateUserRepository`)
- **Implementations**: Interface name + `Impl` suffix
- **UseCases**: Operation + `UseCase` (`CreateLibraryUseCase`)
- **DTOs**: Operation + `InputDTO`/`OutputDTO`

### Package Structure
```
library_usecase/
├── create_library_usecase.go
├── get_library_by_id_usecase.go
└── update_library_by_id_usecase.go
```

### File Naming
- Snake_case for files
- One struct per file (main implementation)

---

## 🔍 Important Notes for AI

1. **NEVER modify `internal/infra/database/db/*.go`** - SQLC generated
2. **NEVER modify `internal/infra/di/wire_gen.go`** - Wire generated
3. **Always follow layer separation** - Domain doesn't import infra
4. **Migrations are immutable** - Create new migration for changes
5. **Enum values in English** - Error messages in Portuguese
6. **Binary data as Base64** in API, `[]byte` in domain
7. **Repository methods are atomic** - One operation per method

---

## 🚧 Current Migration Version

Last migration: `000008_library.up.sql`  
Next migration number: `000009`

---

## 📞 Context for Common Tasks

### "Add new CRUD entity"
→ Follow 11-step process in **Common Operations** section

### "Change database schema"
→ Create new migration, run `make migrate-up`, then `make sqlc-generate`

### "Add new endpoint"
→ Create UseCase → Handler method → Add route in `routes.go`

### "Fix dependency injection"
→ Edit `wire.go` → Run `wire` in `internal/infra/di/`

### "Handle binary files"
→ Use Base64 in DTOs, `[]byte` in entity, LONGBLOB in DB

---

## 🎓 Learning Resources

- Clean Architecture: Robert C. Martin
- SQLC: https://docs.sqlc.dev/
- Wire: https://github.com/google/wire
- Chi Router: https://go-chi.io/