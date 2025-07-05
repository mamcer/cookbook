# Internal Package Structure

This directory contains the internal packages for the cookbook application, organized using clean architecture principles.

## Package Structure

### `config/`
Configuration management package that handles:
- Loading configuration from JSON files
- Environment variable overrides
- Configuration validation
- Default values

### `models/`
Data models and DTOs (Data Transfer Objects):
- `recipe.go`: Main recipe models and request/response structures
- `dto.go`: Internal DTOs for database operations

### `database/`
Database layer with repository pattern:
- `connection.go`: Database connection management with connection pooling
- `repository.go`: Repository interfaces defining database operations
- `mysql_repository.go`: MySQL implementation of repository interfaces

### `handlers/`
HTTP handlers and business logic:
- `service.go`: Business logic layer (service layer)
- `recipe_handlers.go`: HTTP handlers for recipe endpoints

### `middleware/`
HTTP middleware:
- `cors.go`: CORS handling middleware
- `logging.go`: Request/response logging middleware

## Architecture Benefits

1. **Separation of Concerns**: Each package has a single responsibility
2. **Dependency Injection**: Services depend on interfaces, not concrete implementations
3. **Testability**: Easy to mock dependencies for unit testing
4. **Maintainability**: Clear structure makes code easier to understand and modify
5. **Scalability**: Easy to add new features or change implementations

## Usage

The main application (`cmd/api/main_new.go`) demonstrates how to wire all these components together:

1. Load configuration
2. Initialize database connection
3. Create repositories
4. Create services
5. Create handlers
6. Set up HTTP routes with middleware

## Migration from Old Structure

The new structure maintains backward compatibility by keeping the old routes while adding new `/api/v1` routes. This allows for a gradual migration. 