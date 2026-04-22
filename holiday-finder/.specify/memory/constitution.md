# Holiday Finder Constitution

## Core Principles

### I. API-First Design
- Web API defines all business logic and data access
- Blazor client consumes only the public API surface
- No direct database access from client
- API controllers use dependency injection for services

### II. Separation of Concerns
- API project contains controllers, services, models, and data access
- Client project contains Blazor components and HTTP client services
- Shared project for DTOs and contracts (if needed)
- Clear boundaries between presentation and business layers

### III. Testing Requirements
- Unit tests for API controllers and services using xUnit
- Integration tests for API endpoints using WebApplicationFactory
- Component tests for Blazor components using bUnit (optional for MVP)
- Minimum 70% code coverage for API business logic

### IV. Security Standards
- HTTPS enforced in production
- CORS configured explicitly for known origins
- Input validation on all API endpoints using DataAnnotations
- Exception handling middleware - no sensitive data in error responses

### V. Error Handling
- Consistent error response format (status code, message, details)
- Structured logging using ILogger
- Global exception handler in API middleware
- User-friendly error messages in Blazor UI

## Technical Standards

### Technology Stack
- .NET 8 or later
- ASP.NET Core Web API
- Blazor WebAssembly or Blazor Server
- Entity Framework Core (if using database)
- xUnit for testing

### Code Quality
- Follow C# naming conventions (PascalCase for public, camelCase for private)
- Use async/await for all I/O operations
- Dependency injection for all services
- Keep controllers thin - business logic in services

### API Standards
- RESTful conventions (GET, POST, PUT, DELETE)
- Route attribute-based routing
- Return appropriate HTTP status codes
- API versioning for breaking changes (when needed)

## Governance

- All new features must include unit tests
- API changes require testing before client implementation
- Code reviews verify security and separation of concerns
- Constitution updates require team consensus

**Version**: 1.0.0 | **Ratified**: 2026-04-20 | **Last Amended**: 2026-04-20
