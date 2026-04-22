# Technology Research: AI-Powered Holiday Booking Platform

**Feature**: 001-ai-holiday-booking  
**Date**: 2026-04-20  
**Purpose**: Document technology decisions and best practices for implementation

## 1. HATEOAS Implementation in ASP.NET Core

### Decision: Use HAL (Hypertext Application Language) format

**Rationale**:
- HAL is a widely-adopted standard for HATEOAS REST APIs
- Provides clear structure for links and embedded resources
- Good library support in .NET ecosystem

**Recommended Approach**:
- Use `HAL.AspNetCore.OData` or custom link generation
- Include `_links` in all response DTOs for navigation
- Key links to include:
  - `self`: Current resource URI
  - `next`/`prev`: Pagination for collections
  - Action links: `submit`, `update`, `delete` based on permissions
  - Navigation links: Related resources (e.g., user → searches)

**Example Response Structure**:
```json
{
  "id": "123",
  "email": "user@example.com",
  "membershipTier": "Premium",
  "_links": {
    "self": { "href": "/api/users/123" },
    "searches": { "href": "/api/users/123/searches" },
    "membership": { "href": "/api/users/123/membership" },
    "upgrade": { "href": "/api/memberships/upgrade", "method": "POST" }
  }
}
```

**Implementation Notes**:
- Generate links dynamically based on user permissions and membership tier
- Premium users see additional action links (e.g., advanced search filters)
- Include only valid actions in `_links` (principle of least surprise)

**Alternatives Considered**:
- JSON-LD: More complex, overkill for this use case
- Custom link format: Reinventing the wheel, HAL is established

---

## 2. Azure Cosmos DB Design & Best Practices

### Decision: Use NoSQL API with partition strategy optimized for multi-tenant access

**Partition Key Strategy**:

**Primary Choice: `/userId`**
- **Rationale**: Aligns with query patterns (users accessing their own data)
- **High cardinality**: Each user is a unique partition
- **Avoids hot partitions**: Writes distributed across users
- **Query optimization**: Most queries are user-scoped

**For large-scale deployments with >20GB per user**, consider Hierarchical Partition Keys (HPK):
- `/userId/year` - Distributes data if single users exceed 20GB partition limit

**Container Design**:

1. **Users Container**
   - Partition Key: `/userId` (use `id` field for single-item partition)
   - Stores: User profile, authentication info, membership tier
   - Access Pattern: Point reads by userId

2. **SearchCriteria Container**
   - Partition Key: `/userId`
   - Stores: User's saved search criteria
   - Access Pattern: Query by userId, ordered by timestamp

3. **SearchHistory Container**
   - Partition Key: `/userId`
   - Stores: Search records with results snapshot
   - Access Pattern: Query by userId, paginated results
   - Consider HPK `/userId/year` if history grows large

4. **MembershipTiers Container**
   - Partition Key: `/id` (low cardinality acceptable for reference data)
   - Stores: Tier definitions (Free, Premium, VIP)
   - Access Pattern: Read-heavy, infrequent updates
   - Use change feed for cache invalidation

**Data Modeling Best Practices** (from cosmosdb-best-practices skill):

- **Embed related data**: Include membership tier details in User document for fast access
- **Denormalize for reads**: Store user's current tier name/features in User doc (avoid joins)
- **Reference large data**: Holiday packages should be referenced, not embedded (too large)
- **Keep under 2MB**: Each document must be < 2MB
- **Type discriminators**: Add `"type": "User"` field for polymorphic queries

**SDK Best Practices**:

- **Singleton client**: Create one `CosmosClient` instance, register in DI container
- **Async APIs**: Use `await` for all Cosmos operations
- **Connection mode**: Use Direct mode for production (better latency)
- **Handle 429s**: Implement retry logic with exponential backoff for rate limiting
- **Diagnostics**: Log diagnostic strings when latency > 100ms

**Indexing Strategy**:

- **Users**: Default indexing (small docs, infrequent queries)
- **SearchCriteria**: Composite index on `[/userId, /timestamp]` for sorted queries
- **SearchHistory**: Exclude `/results/*` from indexing (large, never queried)

**Query Optimization**:

- **Avoid cross-partition queries**: All user queries include `WHERE userId = @userId`
- **Use projections**: `SELECT c.id, c.email, c.membershipTier` (not `SELECT *`)
- **Pagination**: Use continuation tokens for large result sets
- **Parameterize**: Always use parameterized queries to enable query plan caching

**Throughput Strategy**:

- **Autoscale RU/s**: Start with 1000-4000 max RU/s, adjust based on metrics
- **Container-level throughput**: Better cost efficiency for multi-tenant scenarios
- **Monitor**: Track RU consumption per request, optimize high-RU operations

**Alternatives Considered**:
- SQL Database: Cosmos DB chosen for global distribution, scalability, NoSQL flexibility
- Partition by `/searchId`: Rejected - queries are user-centric, not search-centric
- Single container with type field: Rejected - separate containers provide better throughput isolation

---

## 3. Blazor WebAssembly Mobile-Responsive Design

### Decision: Build Progressive Web App (PWA) with responsive Blazor WASM

**Approach**:

1. **Blazor WASM** (Client-side):
   - Single-page application runs entirely in browser
   - Installable as PWA on mobile devices
   - Offline capability with service workers
   - Smaller initial download than Blazor Server

2. **Responsive Design**:
   - Use CSS Grid and Flexbox for fluid layouts
   - Mobile-first approach with breakpoints:
     - Mobile: < 640px
     - Tablet: 640px - 1024px
     - Desktop: > 1024px
   - Touch-friendly UI: 44px minimum touch target size
   - Consider Mudblazor or Blazorise component library for pre-built responsive components

3. **PWA Configuration**:
   - Add `manifest.json` for installability
   - Service worker for offline support and caching
   - Cache API responses for search results (with expiration)
   - Background sync for failed requests

4. **Mobile Optimization**:
   - Lazy loading for Blazor components/pages
   - Image optimization (WebP format, responsive images)
   - Minimize bundle size (remove unused CSS/JS)
   - Virtual scrolling for long result lists

**Implementation Pattern**:
```csharp
// Use Blazor's NavigationManager for SPA routing
@inject NavigationManager Navigation
@inject HttpClient Http

// Responsive layout component
<div class="container">
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
        @foreach (var package in Packages)
        {
            <HolidayPackageCard Package="@package" />
        }
    </div>
</div>
```

**State Management**:
- Use Blazor's cascading parameters for global state (user session)
- Local storage for persisting auth tokens
- In-memory cache for frequently accessed data (membership tiers)

**Alternatives Considered**:
- Blazor Server: Rejected - requires persistent connection, poor mobile experience with spotty networks
- Separate mobile app: Deferred to v2 - PWA provides 80% of native experience at lower cost
- React/Vue: Rejected - staying in .NET ecosystem for consistency

---

## 4. AI Agent Integration Patterns

### Decision: External AI Service via HTTP with async processing

**Architecture**:

1. **Synchronous Search** (for initial MVP):
   - User submits criteria → API calls AI service → Returns results
   - Timeout: 10 seconds (per SC-002 success criteria)
   - If timeout, return partial results or cached suggestions

2. **For v2 - Async Processing** (if AI takes > 10s):
   - User submits criteria → API queues request → Returns job ID
   - Client polls `/api/searches/{jobId}/status` for completion
   - SignalR for real-time updates (push results when ready)

**AI Service Integration**:

```csharp
public interface IHolidaySearchService
{
    Task<IEnumerable<HolidayPackage>> SearchAsync(
        SearchCriteria criteria, 
        CancellationToken ct);
}

// Implementation
public class AzureOpenAISearchService : IHolidaySearchService
{
    private readonly HttpClient _httpClient;
    private readonly ILogger<AzureOpenAISearchService> _logger;
    
    public async Task<IEnumerable<HolidayPackage>> SearchAsync(
        SearchCriteria criteria, 
        CancellationToken ct)
    {
        // Call Azure OpenAI or custom AI service
        // Parse and rank results
        // Return top packages
    }
}
```

**Error Handling**:
- AI service unavailable: Return cached popular destinations
- Timeout: Return partial results with warning
- Invalid criteria: Validate before calling AI service
- Rate limiting: Queue requests if AI service has rate limits

**Caching Strategy**:
- Cache popular destination searches (e.g., "Paris, budget €1000") for 1 hour
- User-specific searches not cached (personalized)
- Use distributed cache (Azure Cache for Redis) if multi-instance deployment

**Membership Tier Handling**:
- Free tier: Standard AI processing queue
- Premium: Priority queue, faster response
- VIP: Advanced filters, multiple AI models for comparison

**Alternatives Considered**:
- Azure Functions for AI processing: Good option for v2 async architecture
- Embedded ML model: Too complex for MVP, external service provides flexibility
- No AI (manual search): Defeats core value proposition

---

## 5. Authentication with Blazor WASM + Web API

### Decision: JWT Bearer tokens with ASP.NET Core Identity

**Architecture**:

1. **API Side** (ASP.NET Core Web API):
   - Use ASP.NET Core Identity for user management
   - Store users in Cosmos DB (custom Identity provider)
   - Issue JWT tokens on successful login
   - Validate JWT on protected endpoints using `[Authorize]` attribute

2. **Client Side** (Blazor WASM):
   - Store JWT in browser's localStorage (or sessionStorage for better security)
   - Include JWT in Authorization header for all API requests
   - Use Blazor's `AuthenticationStateProvider` for auth state
   - Redirect to login page if token expired/invalid

**Implementation**:

```csharp
// API - Startup.cs / Program.cs
services.AddAuthentication(JwtBearerDefaults.AuthenticationScheme)
    .AddJwtBearer(options =>
    {
        options.TokenValidationParameters = new TokenValidationParameters
        {
            ValidateIssuer = true,
            ValidateAudience = true,
            ValidateLifetime = true,
            ValidateIssuerSigningKey = true,
            ValidIssuer = Configuration["Jwt:Issuer"],
            ValidAudience = Configuration["Jwt:Audience"],
            IssuerSigningKey = new SymmetricSecurityKey(
                Encoding.UTF8.GetBytes(Configuration["Jwt:Key"]))
        };
    });

// Client - AuthStateProvider
public class JwtAuthenticationStateProvider : AuthenticationStateProvider
{
    private readonly HttpClient _httpClient;
    private readonly ILocalStorageService _localStorage;

    public override async Task<AuthenticationState> GetAuthenticationStateAsync()
    {
        var token = await _localStorage.GetItemAsync<string>("authToken");
        
        if (string.IsNullOrEmpty(token))
            return new AuthenticationState(new ClaimsPrincipal(new ClaimsIdentity()));

        _httpClient.DefaultRequestHeaders.Authorization = 
            new AuthenticationHeaderValue("Bearer", token);

        return new AuthenticationState(
            new ClaimsPrincipal(new ClaimsIdentity(ParseClaimsFromJwt(token), "jwt")));
    }
}
```

**Token Management**:
- Token expiration: 60 minutes (configurable)
- Refresh tokens: Implement for seamless re-authentication (store securely)
- Logout: Clear token from localStorage and redirect

**Security Considerations**:
- HTTPS only (enforce in production)
- HttpOnly cookies alternative: More secure but complex for WASM
- CORS configuration: Allow only known origins
- Password hashing: Use ASP.NET Core Identity's default (PBKDF2)
- Rate limiting on login endpoint: Prevent brute force

**Membership Tier Access**:
- Include `membershipTier` claim in JWT
- Use `[Authorize(Policy = "PremiumOnly")]` for premium features
- Client-side: Show/hide UI elements based on tier claim

**Password Reset Flow**:
1. User requests reset → API sends email with reset token
2. User clicks link → Client shows reset form
3. User submits new password + token → API validates and updates

**Alternatives Considered**:
- Azure AD B2C: Overkill for MVP, adds external dependency
- Cookie-based auth: Doesn't work well with WASM cross-origin
- OAuth2/OpenID Connect: Consider for v2 if social login needed

---

## Technology Stack Summary

| Component | Technology | Version |
|-----------|------------|---------|
| Language | C# | .NET 8 |
| API Framework | ASP.NET Core Web API | 8.0+ |
| Client Framework | Blazor WebAssembly | 8.0+ |
| Database | Azure Cosmos DB | NoSQL API |
| Authentication | JWT + ASP.NET Identity | Core 8.0+ |
| Testing | xUnit + WebApplicationFactory | Latest |
| API Pattern | REST + HATEOAS (HAL) | - |
| State Management | Blazor Cascading Parameters | - |
| HTTP Client | HttpClient (DI registered) | - |
| Logging | ILogger + Application Insights | - |

## Development Tools

- **IDE**: Visual Studio 2022 or VS Code with C# Dev Kit
- **Database**: Azure Cosmos DB Emulator for local development
- **API Testing**: Swagger/OpenAPI (built-in with ASP.NET Core)
- **Client Testing**: bUnit for Blazor component tests
- **Package Manager**: NuGet
- **Version Control**: Git (already assumed)

## Next Steps (Phase 1)

1. Generate data-model.md with Cosmos DB schema definitions
2. Create API contracts/ with DTO definitions and HATEOAS link structures
3. Generate quickstart.md with setup instructions
4. Update agent context in .github/copilot-instructions.md

## References

- [Azure Cosmos DB Best Practices](https://learn.microsoft.com/azure/cosmos-db/nosql/best-practice-dotnet)
- [Blazor WASM Authentication](https://learn.microsoft.com/aspnet/core/blazor/security/webassembly/)
- [HAL Specification](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal)
- [ASP.NET Core REST API Best Practices](https://learn.microsoft.com/aspnet/core/web-api/)
