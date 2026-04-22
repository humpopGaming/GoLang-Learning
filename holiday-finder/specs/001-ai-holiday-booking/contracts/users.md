# API Contracts: Users

**Feature**: 001-ai-holiday-booking  
**Date**: 2026-04-20  
**API Pattern**: REST + HATEOAS (HAL format)

## Overview

Contracts for user profile management and account operations.

---

## 1. Get User Profile

**GET** `/api/users/{userId}/profile`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

### Response (200 OK)

```json
{
  "userId": "user-550e8400-e29b-41d4-a716-446655440000",
  "email": "john.doe@example.com",
  "firstName": "John",
  "lastName": "Doe",
  "membershipTier": "Premium",
  "membershipExpires": "2027-03-20T10:00:00Z",
  "preferences": {
    "currency": "USD",
    "language": "en-US",
    "notifications": {
      "email": true,
      "priceDropAlerts": true,
      "newsletter": false
    },
    "defaultSearchSettings": {
      "travelClass": "economy",
      "accommodation": ["hotel"]
    }
  },
  "accountInfo": {
    "createdAt": "2026-01-15T10:30:00Z",
    "lastLoginAt": "2026-04-20T08:15:00Z",
    "emailVerified": false,
    "totalSearches": 47,
    "totalSavedCriteria": 8
  },
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400/profile"
    },
    "update": {
      "href": "/api/users/user-550e8400/profile",
      "method": "PUT"
    },
    "changePassword": {
      "href": "/api/users/user-550e8400/password",
      "method": "PUT"
    },
    "membership": {
      "href": "/api/users/user-550e8400/membership"
    },
    "searches": {
      "href": "/api/users/user-550e8400/searches"
    },
    "criteria": {
      "href": "/api/criteria"
    },
    "deleteAccount": {
      "href": "/api/users/user-550e8400",
      "method": "DELETE"
    }
  }
}
```

---

## 2. Update User Profile

**PUT** `/api/users/{userId}/profile`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

```json
{
  "firstName": "John",
  "lastName": "Doe",
  "preferences": {
    "currency": "EUR",
    "language": "en-US",
    "notifications": {
      "email": true,
      "priceDropAlerts": true,
      "newsletter": false
    },
    "defaultSearchSettings": {
      "travelClass": "business",
      "accommodation": ["hotel", "apartment"]
    }
  }
}
```

### Fields

| Field | Type | Required | Constraints |
|-------|------|----------|-------------|
| `firstName` | string | ❌ | Max 50 chars |
| `lastName` | string | ❌ | Max 50 chars |
| `preferences` | object | ❌ | Preferences object |

### Response (200 OK)

```json
{
  "userId": "user-550e8400-e29b-41d4-a716-446655440000",
  "updatedAt": "2026-04-20T09:30:00Z",
  "changes": ["firstName", "preferences.currency", "preferences.defaultSearchSettings"],
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400/profile"
    },
    "profile": {
      "href": "/api/users/user-550e8400/profile"
    }
  }
}
```

### Error Responses

**400 Bad Request** - Validation error
```json
{
  "status": 400,
  "title": "Validation Failed",
  "errors": {
    "firstName": ["First name cannot exceed 50 characters"],
    "preferences.currency": ["Invalid currency code. Use ISO 4217 codes."]
  }
}
```

---

## 3. Change Password

**PUT** `/api/users/{userId}/password`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

```json
{
  "currentPassword": "SecurePass123!",
  "newPassword": "NewSecurePass456!",
  "confirmNewPassword": "NewSecurePass456!"
}
```

### Fields

| Field | Type | Required | Constraints |
|-------|------|----------|-------------|
| `currentPassword` | string | ✅ | Must match existing password |
| `newPassword` | string | ✅ | Min 8 chars, 1 uppercase, 1 number, 1 special |
| `confirmNewPassword` | string | ✅ | Must match newPassword |

### Response (200 OK)

```json
{
  "changedAt": "2026-04-20T09:35:00Z",
  "message": "Password changed successfully",
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400/password"
    },
    "profile": {
      "href": "/api/users/user-550e8400/profile"
    }
  }
}
```

### Error Responses

**401 Unauthorized** - Wrong current password
```json
{
  "status": 401,
  "title": "Invalid Current Password",
  "detail": "The current password you provided is incorrect"
}
```

**400 Bad Request** - Password validation failed
```json
{
  "status": 400,
  "title": "Invalid New Password",
  "errors": {
    "newPassword": [
      "Password must be at least 8 characters long",
      "Password must contain at least one uppercase letter"
    ]
  }
}
```

---

## 4. Request Email Verification

**POST** `/api/users/{userId}/verify-email/request`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

**Body**: Empty

### Response (200 OK)

```json
{
  "sentAt": "2026-04-20T09:40:00Z",
  "email": "john.doe@example.com",
  "message": "Verification email sent successfully",
  "expiresIn": 86400,
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400/verify-email/request"
    },
    "profile": {
      "href": "/api/users/user-550e8400/profile"
    }
  }
}
```

---

## 5. Confirm Email Verification

**POST** `/api/users/{userId}/verify-email/confirm`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

```json
{
  "verificationToken": "token-from-email"
}
```

### Response (200 OK)

```json
{
  "verifiedAt": "2026-04-20T09:45:00Z",
  "email": "john.doe@example.com",
  "message": "Email verified successfully",
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400/verify-email/confirm"
    },
    "profile": {
      "href": "/api/users/user-550e8400/profile"
    }
  }
}
```

### Error Responses

**400 Bad Request** - Invalid/expired token
```json
{
  "status": 400,
  "title": "Invalid Verification Token",
  "detail": "The verification token is invalid or has expired",
  "_links": {
    "resend": {
      "href": "/api/users/user-550e8400/verify-email/request",
      "method": "POST"
    }
  }
}
```

---

## 6. Delete Account

**DELETE** `/api/users/{userId}`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

```json
{
  "password": "SecurePass123!",
  "reason": "No longer need the service",
  "feedback": "Service was great but no longer traveling"
}
```

### Fields

| Field | Type | Required | Constraints |
|-------|------|----------|-------------|
| `password` | string | ✅ | Must match current password |
| `reason` | string | ❌ | Deletion reason |
| `feedback` | string | ❌ | User feedback |

### Response (200 OK)

```json
{
  "deletedAt": "2026-04-20T09:50:00Z",
  "userId": "user-550e8400-e29b-41d4-a716-446655440000",
  "message": "Account scheduled for deletion in 30 days",
  "dataRetentionPolicy": {
    "immediatelyDeleted": [
      "Personal information",
      "Payment methods"
    ],
    "retainedFor30Days": [
      "Search history (anonymized)",
      "Usage statistics"
    ]
  },
  "refund": {
    "eligible": true,
    "amount": 47.99,
    "currency": "USD",
    "processedWithin": "5-7 business days"
  },
  "_links": {
    "cancelDeletion": {
      "href": "/api/users/user-550e8400/cancel-deletion",
      "method": "POST",
      "description": "Cancel deletion within 30 days"
    }
  }
}
```

### Error Responses

**401 Unauthorized** - Wrong password
```json
{
  "status": 401,
  "title": "Invalid Password",
  "detail": "Password is incorrect. Account deletion requires password confirmation."
}
```

---

## 7. Cancel Account Deletion

**POST** `/api/users/{userId}/cancel-deletion`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

**Body**: Empty or optional
```json
{
  "reason": "Changed my mind, still traveling"
}
```

### Response (200 OK)

```json
{
  "cancelledAt": "2026-04-22T10:00:00Z",
  "userId": "user-550e8400-e29b-41d4-a716-446655440000",
  "message": "Account deletion cancelled successfully",
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400/profile"
    },
    "searches": {
      "href": "/api/searches",
      "method": "POST"
    }
  }
}
```

---

## 8. Get Usage Statistics

**GET** `/api/users/{userId}/stats`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

### Response (200 OK)

```json
{
  "userId": "user-550e8400-e29b-41d4-a716-446655440000",
  "period": "all-time",
  "stats": {
    "totalSearches": 47,
    "searchesThisMonth": 12,
    "searchesToday": 3,
    "quotaRemaining": 7,
    "quotaResetsAt": "2026-04-21T00:00:00Z",
    "savedCriteria": 8,
    "savedCriteriaLimit": 20,
    "averageSearchTime": 5.3,
    "topDestinations": [
      { "name": "Paris, France", "searches": 8 },
      { "name": "Barcelona, Spain", "searches": 5 },
      { "name": "Tokyo, Japan", "searches": 4 }
    ]
  },
  "membershipInfo": {
    "tier": "Premium",
    "since": "2026-03-20T10:00:00Z",
    "expiresAt": "2027-03-20T10:00:00Z",
    "daysRemaining": 335
  },
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400/stats"
    },
    "searches": {
      "href": "/api/users/user-550e8400/searches"
    },
    "upgrade": {
      "href": "/api/memberships"
    }
  }
}
```

---

## Business Rules

### Email Verification (v2 Feature)

- Verification token expires in 24 hours
- User can request new token if expired
- Email changes require re-verification
- Certain features may require verified email (future enhancement)

### Account Deletion

- **Grace Period**: 30 days before permanent deletion
- **Data Retention**: 
  - Personal info deleted immediately
  - Anonymized usage stats retained for analytics
  - Search history anonymized but retained for 30 days
- **Refunds**: Pro-rated refund for active subscriptions
- **Cancellation**: User can cancel deletion within 30 days

### Password Requirements

- Minimum 8 characters
- At least one uppercase letter
- At least one lowercase letter
- At least one number
- At least one special character (!@#$%^&*)
- Cannot be same as email
- Cannot be a common password (checked against list)

### Profile Updates

- Email changes require password confirmation
- Email must be unique (checked before update)
- Currency must be valid ISO 4217 code
- Language must be supported locale

### Privacy & Security

- User can export all their data (GDPR compliance) - v2 feature
- User can download search history as JSON/CSV - v2 feature
- Two-factor authentication - v2 feature
- Login history and active sessions - v2 feature

---

## HATEOAS Links Reference

| Link Rel | Description | Method |
|----------|-------------|--------|
| `self` | Current resource | GET |
| `profile` | User profile | GET |
| `update` | Update profile | PUT |
| `changePassword` | Change password | PUT |
| `membership` | Membership details | GET |
| `searches` | User's search history | GET |
| `criteria` | Saved search criteria | GET |
| `deleteAccount` | Delete user account | DELETE |
| `cancelDeletion` | Cancel scheduled deletion | POST |
| `resend` | Resend verification email | POST |
| `upgrade` | View membership tiers | GET |
| `stats` | Usage statistics | GET |
