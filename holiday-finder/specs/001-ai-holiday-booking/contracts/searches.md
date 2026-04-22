# API Contracts: Holiday Searches

**Feature**: 001-ai-holiday-booking  
**Date**: 2026-04-20  
**API Pattern**: REST + HATEOAS (HAL format)

## Overview

Contracts for search criteria management and AI-powered holiday search execution.

---

## 1. Execute Search

**POST** `/api/searches`

**Headers**:
```
Authorization: Bearer {jwt-token}
Content-Type: application/json
```

```json
{
  "criteriaId": "criteria-7c9e6679-7425-40de-944b-e07fc1f90ae7",
  "destinations": [
    {
      "code": "PAR",
      "name": "Paris, France",
      "type": "city"
    }
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
  }
}
```

### Fields

| Field | Type | Required | Constraints |
|-------|------|----------|-------------|
| `criteriaId` | string | ❌ | Reference to saved criteria (if reusing) |
| `destinations` | array | ✅ | 1-10 destinations |
| `dateRange` | object | ✅ | Start date >= today, end > start |
| `budget` | object | ✅ | Min < max, both > 0 |
| `preferences` | object | ❌ | Optional preferences |
| `filters` | object | ❌ | Premium/VIP only (validated by tier) |

### Response (200 OK) - Search Results

```json
{
  "searchId": "search-8f3a1b9c-4d2e-4a6b-9c7f-1e5d8a9b0c3d",
  "executedAt": "2026-04-20T08:35:12Z",
  "processingTime": 4.27,
  "resultCount": 8,
  "criteria": {
    "destinations": ["Paris, France"],
    "dateRange": "2026-07-01 to 2026-07-15",
    "budget": "$1,000 - $3,000"
  },
  "packages": [
    {
      "packageId": "pkg-abc123",
      "destination": {
        "code": "PAR",
        "name": "Paris, France",
        "country": "France",
        "imageUrl": "https://cdn.example.com/paris.jpg"
      },
      "dates": {
        "departure": "2026-07-03T10:30:00Z",
        "return": "2026-07-13T18:45:00Z",
        "nights": 10
      },
      "price": {
        "total": 2450,
        "currency": "USD",
        "priceType": "per-person",
        "breakdown": {
          "flight": 850,
          "accommodation": 1400,
          "fees": 200
        }
      },
      "accommodation": {
        "name": "Le Grand Hotel",
        "type": "hotel",
        "stars": 4,
        "address": "123 Rue de Rivoli, Paris",
        "amenities": ["wifi", "breakfast", "pool"]
      },
      "flight": {
        "airline": "Air France",
        "flightNumber": "AF123",
        "class": "economy",
        "stops": 0,
        "duration": "7h 30m"
      },
      "relevanceScore": 0.95,
      "highlights": [
        "Central location near Louvre",
        "Direct flight available",
        "Free cancellation until 7 days before"
      ],
      "_links": {
        "self": {
          "href": "/api/searches/search-8f3a1b9c/packages/pkg-abc123"
        },
        "book": {
          "href": "https://partner.booking.com/book/abc123",
          "external": true
        },
        "details": {
          "href": "/api/searches/search-8f3a1b9c/packages/pkg-abc123/details"
        },
        "save": {
          "href": "/api/users/user-550e8400/favorites",
          "method": "POST"
        }
      }
    }
  ],
  "metadata": {
    "aiModel": "holiday-search-v2",
    "membershipTierUsed": "Premium",
    "searchQuotaRemaining": 7
  },
  "_links": {
    "self": {
      "href": "/api/searches/search-8f3a1b9c-4d2e-4a6b-9c7f-1e5d8a9b0c3d"
    },
    "history": {
      "href": "/api/users/user-550e8400/searches"
    },
    "saveCriteria": {
      "href": "/api/criteria",
      "method": "POST"
    },
    "refineSearch": {
      "href": "/api/searches",
      "method": "POST",
      "templated": true
    }
  }
}
```

### Response (202 Accepted) - Async Processing

If search takes > 10s, return job ID for polling:

```json
{
  "searchId": "search-8f3a1b9c-4d2e-4a6b-9c7f-1e5d8a9b0c3d",
  "status": "processing",
  "estimatedWaitTime": 15,
  "_links": {
    "self": {
      "href": "/api/searches/search-8f3a1b9c/status"
    },
    "poll": {
      "href": "/api/searches/search-8f3a1b9c",
      "pollInterval": 2000
    }
  }
}
```

### Error Responses

**400 Bad Request** - Invalid criteria
```json
{
  "status": 400,
  "title": "Invalid Search Criteria",
  "errors": {
    "dateRange": ["Start date must be in the future"],
    "budget": ["Minimum budget must be less than maximum"]
  }
}
```

**403 Forbidden** - Quota exceeded
```json
{
  "status": 403,
  "title": "Search Quota Exceeded",
  "detail": "You have reached your daily search limit of 3 searches",
  "quotaReset": "2026-04-21T00:00:00Z",
  "_links": {
    "upgrade": {
      "href": "/api/memberships/upgrade"
    }
  }
}
```

**503 Service Unavailable** - AI service down
```json
{
  "status": 503,
  "title": "Search Service Unavailable",
  "detail": "Unable to process search at this time. Please try again.",
  "retryAfter": 60
}
```

---

## 2. Get Search Results (Retrieve)

**GET** `/api/searches/{searchId}`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

### Response (200 OK)

Same structure as execute search response above.

### Error Responses

**404 Not Found** - Search doesn't exist or not owned by user
```json
{
  "status": 404,
  "title": "Search Not Found",
  "detail": "Search with ID 'search-8f3a1b9c' was not found"
}
```

---

## 3. Get Package Details

**GET** `/api/searches/{searchId}/packages/{packageId}/details`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

### Response (200 OK)

```json
{
  "packageId": "pkg-abc123",
  "destination": {
    "code": "PAR",
    "name": "Paris, France",
    "description": "The City of Light offers world-class museums...",
    "timezone": "Europe/Paris",
    "language": "French",
    "currency": "EUR",
    "images": [
      "https://cdn.example.com/paris-1.jpg",
      "https://cdn.example.com/paris-2.jpg"
    ]
  },
  "itinerary": [
    {
      "day": 1,
      "title": "Arrival & Eiffel Tower",
      "activities": [
        "Airport transfer to hotel",
        "Afternoon visit to Eiffel Tower",
        "Seine river cruise (optional)"
      ]
    },
    {
      "day": 2,
      "title": "Louvre Museum & Notre-Dame",
      "activities": [
        "Guided tour of Louvre Museum",
        "Lunch at Café de Flore",
        "Visit Notre-Dame Cathedral"
      ]
    }
  ],
  "accommodation": {
    "name": "Le Grand Hotel",
    "type": "hotel",
    "stars": 4,
    "checkIn": "2026-07-03T15:00:00Z",
    "checkOut": "2026-07-13T11:00:00Z",
    "roomType": "Deluxe Double Room",
    "amenities": ["wifi", "breakfast", "pool", "gym", "spa"],
    "policies": {
      "cancellation": "Free cancellation until 7 days before check-in",
      "payment": "Pay at hotel or online",
      "childPolicy": "Children under 12 stay free"
    },
    "reviews": {
      "rating": 4.5,
      "count": 1247,
      "topComments": [
        "Excellent location",
        "Very clean rooms",
        "Friendly staff"
      ]
    }
  },
  "flights": {
    "outbound": {
      "airline": "Air France",
      "flightNumber": "AF123",
      "departure": {
        "airport": "JFK",
        "city": "New York",
        "time": "2026-07-03T10:30:00Z",
        "terminal": "1"
      },
      "arrival": {
        "airport": "CDG",
        "city": "Paris",
        "time": "2026-07-03T22:00:00Z",
        "terminal": "2E"
      },
      "duration": "7h 30m",
      "aircraft": "Boeing 777",
      "class": "economy",
      "baggage": "1 checked bag included"
    },
    "return": {
      "airline": "Air France",
      "flightNumber": "AF124",
      "departure": {
        "airport": "CDG",
        "city": "Paris",
        "time": "2026-07-13T14:45:00Z",
        "terminal": "2E"
      },
      "arrival": {
        "airport": "JFK",
        "city": "New York",
        "time": "2026-07-13T18:45:00Z",
        "terminal": "1"
      },
      "duration": "8h 0m",
      "aircraft": "Boeing 777",
      "class": "economy",
      "baggage": "1 checked bag included"
    }
  },
  "inclusions": [
    "Round-trip flights",
    "10 nights accommodation with breakfast",
    "Airport transfers",
    "City map and guidebook",
    "24/7 customer support"
  ],
  "exclusions": [
    "Lunch and dinner",
    "Museum entrance fees",
    "Travel insurance",
    "Personal expenses"
  ],
  "termsAndConditions": "https://partner.booking.com/terms/abc123",
  "_links": {
    "self": {
      "href": "/api/searches/search-8f3a1b9c/packages/pkg-abc123/details"
    },
    "book": {
      "href": "https://partner.booking.com/book/abc123",
      "external": true
    },
    "summary": {
      "href": "/api/searches/search-8f3a1b9c/packages/pkg-abc123"
    }
  }
}
```

---

## 4. List Search History

**GET** `/api/users/{userId}/searches`

**Query Parameters**:
- `page` (default: 1)
- `pageSize` (default: 20, max: 50)
- `orderBy` (default: executedAt, options: executedAt, resultCount)

**Headers**:
```
Authorization: Bearer {jwt-token}
```

### Response (200 OK)

```json
{
  "searches": [
    {
      "searchId": "search-8f3a1b9c-4d2e-4a6b-9c7f-1e5d8a9b0c3d",
      "executedAt": "2026-04-20T08:35:12Z",
      "criteria": {
        "destinations": ["Paris, France"],
        "dateRange": "2026-07-01 to 2026-07-15",
        "budget": "$1,000 - $3,000"
      },
      "resultCount": 8,
      "_links": {
        "self": {
          "href": "/api/searches/search-8f3a1b9c-4d2e-4a6b-9c7f-1e5d8a9b0c3d"
        },
        "results": {
          "href": "/api/searches/search-8f3a1b9c-4d2e-4a6b-9c7f-1e5d8a9b0c3d"
        },
        "rerun": {
          "href": "/api/searches",
          "method": "POST"
        }
      }
    }
  ],
  "pagination": {
    "page": 1,
    "pageSize": 20,
    "totalItems": 47,
    "totalPages": 3
  },
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400/searches?page=1&pageSize=20"
    },
    "next": {
      "href": "/api/users/user-550e8400/searches?page=2&pageSize=20"
    },
    "newSearch": {
      "href": "/api/searches",
      "method": "POST"
    }
  }
}
```

---

## 5. Save Search Criteria

**POST** `/api/criteria`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

```json
{
  "name": "Summer Europe Trip",
  "destinations": [
    { "code": "PAR", "name": "Paris, France", "type": "city" }
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
    "accommodation": ["hotel"],
    "activities": ["sightseeing"],
    "travelClass": "economy"
  }
}
```

### Response (201 Created)

```json
{
  "criteriaId": "criteria-7c9e6679-7425-40de-944b-e07fc1f90ae7",
  "name": "Summer Europe Trip",
  "createdAt": "2026-04-20T08:40:00Z",
  "_links": {
    "self": {
      "href": "/api/criteria/criteria-7c9e6679-7425-40de-944b-e07fc1f90ae7"
    },
    "execute": {
      "href": "/api/searches",
      "method": "POST"
    },
    "update": {
      "href": "/api/criteria/criteria-7c9e6679-7425-40de-944b-e07fc1f90ae7",
      "method": "PUT"
    },
    "delete": {
      "href": "/api/criteria/criteria-7c9e6679-7425-40de-944b-e07fc1f90ae7",
      "method": "DELETE"
    }
  }
}
```

---

## 6. Get Saved Criteria

**GET** `/api/criteria`

**Query Parameters**:
- `favorite` (default: false) - Show only favorites

**Headers**:
```
Authorization: Bearer {jwt-token}
```

### Response (200 OK)

```json
{
  "criteria": [
    {
      "criteriaId": "criteria-7c9e6679-7425-40de-944b-e07fc1f90ae7",
      "name": "Summer Europe Trip",
      "destinations": ["Paris, France", "Barcelona, Spain"],
      "dateRange": "2026-07-01 to 2026-07-15",
      "budget": "$1,000 - $3,000",
      "lastUsedAt": "2026-04-20T08:30:00Z",
      "isFavorite": true,
      "_links": {
        "self": {
          "href": "/api/criteria/criteria-7c9e6679-7425-40de-944b-e07fc1f90ae7"
        },
        "execute": {
          "href": "/api/searches",
          "method": "POST"
        }
      }
    }
  ],
  "_links": {
    "self": {
      "href": "/api/criteria"
    },
    "create": {
      "href": "/api/criteria",
      "method": "POST"
    }
  }
}
```

---

## Business Rules

### Search Quotas by Tier

| Tier | Searches/Day | Saved Criteria | AI Timeout |
|------|--------------|----------------|------------|
| Free | 3 | 5 | 10s |
| Premium | 10 | 20 | 15s |
| VIP | Unlimited | 50 | 20s |

### Validation Rules

1. **Destinations**: 1-10 destinations, each must be from valid destination list
2. **Date Range**: Start date >= today, end date > start date, max 90 days duration
3. **Budget**: Min < max, both > 0, max budget <= $50,000
4. **Flexible Dates**: flexibilityDays between 0-7
5. **Advanced Filters**: Only available for Premium/VIP tiers

### Caching Strategy

- Popular destination searches cached for 1 hour
- User-specific searches not cached (personalized results)
- Pagination uses continuation tokens (not cached)

---

## HATEOAS Links Reference

| Link Rel | Description | Method |
|----------|-------------|--------|
| `self` | Current resource | GET |
| `book` | External booking link | GET (redirect) |
| `details` | Package detailed information | GET |
| `save` | Save package to favorites | POST |
| `history` | User's search history | GET |
| `saveCriteria` | Save current criteria | POST |
| `refineSearch` | Execute refined search | POST |
| `execute` | Execute saved criteria | POST |
| `update` | Update saved criteria | PUT |
| `delete` | Delete saved criteria | DELETE |
| `rerun` | Rerun previous search | POST |
| `newSearch` | Create new search | POST |
| `upgrade` | Upgrade membership | GET |
| `poll` | Poll async search status | GET |
