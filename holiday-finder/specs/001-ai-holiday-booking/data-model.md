# Data Model: AI-Powered Holiday Booking Platform

**Feature**: 001-ai-holiday-booking  
**Date**: 2026-04-20  
**Database**: Azure Cosmos DB (NoSQL API)

## Overview

This document defines the data model for the Holiday Finder platform, optimized for Azure Cosmos DB with partition strategies that align with query patterns and constitutional requirements.

## Container Architecture

### Design Principles

1. **Partition by user** for user-scoped data (90% of queries)
2. **Embed frequently-accessed data** to minimize round trips
3. **Keep documents under 2MB** limit
4. **Denormalize for read performance** per Cosmos DB best practices
5. **Use type discriminators** for polymorphic queries
6. **Reference large/external data** to avoid document bloat

### Container Summary

| Container | Partition Key | Purpose | Access Pattern |
|-----------|---------------|---------|----------------|
| Users | `/id` | User profiles and authentication | Point reads by userId |
| SearchCriteria | `/userId` | Saved search criteria | Query by userId + timestamp |
| SearchHistory | `/userId` | Search results archive | Query by userId, paginated |
| MembershipTiers | `/id` | Tier definitions (reference) | Read-heavy, cached |

---

## 1. Users Container

**Partition Key**: `/id` (userId is the id field for single-document partition)

**Purpose**: Store user accounts, authentication info, and membership details

### Document Schema

```json
{
  "id": "user-550e8400-e29b-41d4-a716-446655440000",
  "type": "User",
  "email": "john.doe@example.com",
  "passwordHash": "$2a$12$KIXxLVQHvN...",
  "firstName": "John",
  "lastName": "Doe",
  "membershipTier": {
    "id": "tier-premium",
    "name": "Premium",
    "features": ["priority-search", "advanced-filters", "5-searches-per-day"],
    "expiresAt": "2027-04-20T00:00:00Z"
  },
  "preferences": {
    "currency": "USD",
    "language": "en-US",
    "notifications": true
  },
  "createdAt": "2026-01-15T10:30:00Z",
  "lastLoginAt": "2026-04-20T08:15:00Z",
  "emailVerified": false,
  "_etag": "\"0000d315-0000-0800-0000-63f8a5c00000\""
}
```

### Field Definitions

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | ✅ | Unique user ID (GUID format with "user-" prefix) |
| `type` | string | ✅ | Discriminator field, always "User" |
| `email` | string | ✅ | User's email address (unique, lowercase) |
| `passwordHash` | string | ✅ | Bcrypt hashed password (ASP.NET Identity format) |
| `firstName` | string | ❌ | User's first name |
| `lastName` | string | ❌ | User's last name |
| `membershipTier` | object | ✅ | Embedded tier details for fast access |
| `membershipTier.id` | string | ✅ | Reference to MembershipTiers container |
| `membershipTier.name` | string | ✅ | Tier name (Free/Premium/VIP) |
| `membershipTier.features` | string[] | ✅ | Feature flags for authorization |
| `membershipTier.expiresAt` | datetime | ❌ | Expiration for paid tiers (null = permanent) |
| `preferences` | object | ❌ | User preferences |
| `createdAt` | datetime | ✅ | Account creation timestamp |
| `lastLoginAt` | datetime | ❌ | Last successful login |
| `emailVerified` | boolean | ✅ | Email verification status (v2 feature) |
| `_etag` | string | auto | Optimistic concurrency control |

### Indexing Strategy

- **Default indexing** on all fields (document size is small)
- **Unique constraint**: Email uniqueness enforced at application layer (Cosmos DB doesn't support unique indexes)

### Query Patterns

```sql
-- Point read by user ID (most common)
SELECT * FROM c WHERE c.id = 'user-550e8400-e29b-41d4-a716-446655440000'

-- Find user by email (login)
SELECT * FROM c WHERE c.email = 'john.doe@example.com' AND c.type = 'User'

-- Find users by membership tier (admin query)
SELECT c.id, c.email, c.membershipTier FROM c 
WHERE c.type = 'User' AND c.membershipTier.name = 'Premium'
```

### Data Integrity

- Email uniqueness validated before insert (check if email exists)
- Password complexity enforced at API layer (ASP.NET Identity)
- Membership tier expiration checked on each request

---

## 2. SearchCriteria Container

**Partition Key**: `/userId`

**Purpose**: Store user's saved search criteria for quick re-execution

### Document Schema

```json
{
  "id": "criteria-7c9e6679-7425-40de-944b-e07fc1f90ae7",
  "userId": "user-550e8400-e29b-41d4-a716-446655440000",
  "type": "SearchCriteria",
  "name": "Summer Europe Trip",
  "destinations": [
    { "code": "PAR", "name": "Paris, France", "type": "city" },
    { "code": "BCN", "name": "Barcelona, Spain", "type": "city" }
  ],
  "dateRange": {
    "startDate": "2026-07-01",
    "endDate": "2026-07-15",
    "flexible": true,
    "flexibilityDays": 3
  },
  "budget": {
    "min": 1000,
    "max": 3000,
    "currency": "USD",
    "priceType": "per-person"
  },
  "preferences": {
    "accommodation": ["hotel", "apartment"],
    "activities": ["sightseeing", "dining"],
    "travelClass": "economy"
  },
  "filters": {
    "directFlightsOnly": false,
    "refundable": true,
    "includeBaggage": true
  },
  "createdAt": "2026-04-15T14:20:00Z",
  "updatedAt": "2026-04-18T09:45:00Z",
  "lastUsedAt": "2026-04-20T08:30:00Z",
  "isFavorite": true
}
```

### Field Definitions

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | ✅ | Unique criteria ID (GUID with "criteria-" prefix) |
| `userId` | string | ✅ | Owner user ID (partition key) |
| `type` | string | ✅ | Discriminator, always "SearchCriteria" |
| `name` | string | ❌ | User-defined name for saved search |
| `destinations` | array | ✅ | List of destination objects (1-10 destinations) |
| `dateRange` | object | ✅ | Travel date specifications |
| `budget` | object | ✅ | Budget constraints |
| `preferences` | object | ❌ | User preferences for this search |
| `filters` | object | ❌ | Advanced filters (Premium/VIP feature) |
| `createdAt` | datetime | ✅ | When criteria was created |
| `updatedAt` | datetime | ✅ | Last modification timestamp |
| `lastUsedAt` | datetime | ❌ | Last time used for search |
| `isFavorite` | boolean | ✅ | Quick access flag |

### Indexing Strategy

- **Composite index**: `[/userId, /lastUsedAt DESC]` for recent searches
- **Exclude**: `/preferences/*` and `/filters/*` from indexing (rarely queried)

### Query Patterns

```sql
-- Get user's recent criteria (dashboard)
SELECT * FROM c 
WHERE c.userId = 'user-550e8400-e29b-41d4-a716-446655440000' 
ORDER BY c.lastUsedAt DESC
OFFSET 0 LIMIT 10

-- Get favorite criteria
SELECT * FROM c 
WHERE c.userId = 'user-550e8400-e29b-41d4-a716-446655440000' 
  AND c.isFavorite = true

-- Point read for execution
SELECT * FROM c WHERE c.id = 'criteria-7c9e6679-7425-40de-944b-e07fc1f90ae7'
  AND c.userId = 'user-550e8400-e29b-41d4-a716-446655440000'
```

---

## 3. SearchHistory Container

**Partition Key**: `/userId`

**Purpose**: Archive search executions with results snapshot

### Document Schema

```json
{
  "id": "search-8f3a1b9c-4d2e-4a6b-9c7f-1e5d8a9b0c3d",
  "userId": "user-550e8400-e29b-41d4-a716-446655440000",
  "type": "SearchHistory",
  "criteriaId": "criteria-7c9e6679-7425-40de-944b-e07fc1f90ae7",
  "criteriaSnapshot": {
    "name": "Summer Europe Trip",
    "destinations": ["PAR", "BCN"],
    "dateRange": "2026-07-01 to 2026-07-15",
    "budget": "1000-3000 USD"
  },
  "executedAt": "2026-04-20T08:35:12Z",
  "processingTime": 4.27,
  "resultCount": 8,
  "results": [
    {
      "packageId": "pkg-abc123",
      "destination": "Paris, France",
      "price": 2450,
      "currency": "USD",
      "dates": "2026-07-03 to 2026-07-13",
      "relevanceScore": 0.95,
      "accommodation": "4-star Hotel",
      "flightDetails": "Direct from JFK",
      "bookingUrl": "https://partner.site/book/abc123"
    }
  ],
  "metadata": {
    "aiModel": "holiday-search-v2",
    "membershipTierUsed": "Premium",
    "wasFromCache": false
  }
}
```

### Field Definitions

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | ✅ | Unique search ID (GUID with "search-" prefix) |
| `userId` | string | ✅ | User who executed search (partition key) |
| `type` | string | ✅ | Discriminator, always "SearchHistory" |
| `criteriaId` | string | ❌ | Reference to SearchCriteria (if from saved) |
| `criteriaSnapshot` | object | ✅ | Simplified criteria for display |
| `executedAt` | datetime | ✅ | When search was executed |
| `processingTime` | number | ✅ | Seconds taken (for monitoring) |
| `resultCount` | integer | ✅ | Number of packages returned |
| `results` | array | ✅ | Holiday packages (limited to top 20) |
| `metadata` | object | ✅ | Search execution metadata |

### Indexing Strategy

- **Composite index**: `[/userId, /executedAt DESC]` for search history listing
- **Exclude**: `/results/*` from indexing (large arrays, never queried)
- **Include**: `/resultCount` and `/metadata/membershipTierUsed` for analytics

### Query Patterns

```sql
-- Get user's search history (paginated)
SELECT c.id, c.criteriaSnapshot, c.executedAt, c.resultCount 
FROM c 
WHERE c.userId = 'user-550e8400-e29b-41d4-a716-446655440000'
ORDER BY c.executedAt DESC
OFFSET 0 LIMIT 20

-- Point read for full results
SELECT * FROM c WHERE c.id = 'search-8f3a1b9c-4d2e-4a6b-9c7f-1e5d8a9b0c3d'
  AND c.userId = 'user-550e8400-e29b-41d4-a716-446655440000'
```

### Data Lifecycle

- **Retention**: Keep all history (until user deletes account)
- **For large-scale**: Consider HPK `/userId/year` if user exceeds 20GB
- **Cleanup**: Implement user-initiated deletion of old searches

---

## 4. MembershipTiers Container

**Partition Key**: `/id`

**Purpose**: Reference data for tier definitions (low cardinality acceptable)

### Document Schema

```json
{
  "id": "tier-premium",
  "type": "MembershipTier",
  "name": "Premium",
  "displayName": "Premium Explorer",
  "description": "Enhanced search with priority processing",
  "pricing": {
    "amount": 9.99,
    "currency": "USD",
    "billingCycle": "monthly",
    "annualDiscount": 0.20
  },
  "features": [
    "priority-search",
    "advanced-filters",
    "10-searches-per-day",
    "email-alerts",
    "price-drop-notifications"
  ],
  "limits": {
    "searchesPerDay": 10,
    "savedCriteria": 20,
    "aiProcessingTimeout": 15
  },
  "order": 2,
  "isActive": true,
  "createdAt": "2026-01-01T00:00:00Z",
  "updatedAt": "2026-03-15T12:00:00Z"
}
```

### Field Definitions

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | ✅ | Tier ID (tier-free, tier-premium, tier-vip) |
| `type` | string | ✅ | Discriminator, always "MembershipTier" |
| `name` | string | ✅ | Internal tier name |
| `displayName` | string | ✅ | User-facing tier name |
| `description` | string | ✅ | Tier description for marketing |
| `pricing` | object | ✅ | Pricing information |
| `features` | string[] | ✅ | Feature flags enabled for this tier |
| `limits` | object | ✅ | Usage limits |
| `order` | integer | ✅ | Display order (1=Free, 2=Premium, 3=VIP) |
| `isActive` | boolean | ✅ | Whether tier is available for new sign-ups |

### Indexing Strategy

- **Default indexing** (only 3-5 documents, read-heavy)

### Query Patterns

```sql
-- Get all active tiers (cached in API)
SELECT * FROM c WHERE c.type = 'MembershipTier' AND c.isActive = true
ORDER BY c.order

-- Point read for tier details
SELECT * FROM c WHERE c.id = 'tier-premium'
```

### Caching Strategy

- **Cache**: Load all tiers into memory on API startup
- **Invalidation**: Use Cosmos DB change feed to detect updates
- **TTL**: Refresh cache every 1 hour as fallback

---

## Relationships

```
User (1) ─────── (*) SearchCriteria
  │
  │
  └────────────── (*) SearchHistory
  
MembershipTier (1) ─────── (*) User (embedded)
```

### Relationship Implementation

- **User → SearchCriteria**: Query by `userId` partition
- **User → SearchHistory**: Query by `userId` partition
- **User → MembershipTier**: Embedded in User document (denormalized)
- **SearchCriteria → SearchHistory**: Reference via `criteriaId` field

---

## Data Validation Rules

### User
- Email must be unique (application-level check before insert)
- Password must meet complexity requirements (min 8 chars, 1 uppercase, 1 number)
- Membership tier expiration checked on each authenticated request

### SearchCriteria
- At least 1 destination required, max 10
- Start date must be >= today
- End date must be > start date
- Budget min must be < budget max
- Advanced filters only for Premium/VIP tiers

### SearchHistory
- Results array limited to 20 packages (for document size)
- Processing time logged for performance monitoring

### MembershipTiers
- Feature flags validated against known features list
- Pricing amounts must be >= 0

---

## Cosmos DB Configuration

### Throughput Strategy

**Initial Allocation**:
- **Users**: 400 RU/s provisioned (autoscale to 4000)
- **SearchCriteria**: Shared database throughput 400 RU/s (autoscale to 4000)
- **SearchHistory**: Shared database throughput (same pool)
- **MembershipTiers**: Shared database throughput (minimal usage)

### Performance Targets

| Operation | Target Latency | Target RU Cost |
|-----------|----------------|----------------|
| User login (point read) | < 10ms | 1 RU |
| Load search criteria | < 20ms | 3-5 RU |
| Save search history | < 30ms | 5-10 RU |
| List history (20 items) | < 50ms | 10-20 RU |

### Monitoring

- Track RU consumption per operation type
- Alert if p99 latency > 100ms
- Monitor partition hot spots (should be evenly distributed)

---

## Migration Considerations

### Future Enhancements

1. **Hierarchical Partition Keys**: If a user's SearchHistory exceeds 20GB:
   - Change partition key to `/userId/year`
   - Enables better distribution for power users

2. **Global Distribution**: 
   - Add read regions near user populations
   - Use session consistency for cost optimization

3. **Time-to-Live (TTL)**:
   - Consider TTL for SearchHistory older than 1 year (optional)
   - User can opt-in to keep history longer

4. **Change Feed**:
   - Implement change feed processor for analytics
   - Track search patterns, popular destinations

---

## Appendix: Sample Data

### Sample User (Free Tier)
```json
{
  "id": "user-123",
  "type": "User",
  "email": "jane@example.com",
  "passwordHash": "$2a$12$...",
  "membershipTier": {
    "id": "tier-free",
    "name": "Free",
    "features": ["basic-search"],
    "expiresAt": null
  },
  "createdAt": "2026-04-20T10:00:00Z"
}
```

### Sample MembershipTier (Free)
```json
{
  "id": "tier-free",
  "type": "MembershipTier",
  "name": "Free",
  "displayName": "Free Explorer",
  "description": "Basic holiday search",
  "pricing": { "amount": 0, "currency": "USD", "billingCycle": "none" },
  "features": ["basic-search"],
  "limits": { "searchesPerDay": 3, "savedCriteria": 5, "aiProcessingTimeout": 10 },
  "order": 1,
  "isActive": true
}
```
