# API Contracts: Authentication

**Feature**: 001-ai-holiday-booking  
**Date**: 2026-04-20  
**API Pattern**: REST + HATEOAS (HAL format)

## Overview

Authentication contracts for user registration, login, and token management.

---

## 1. Register Request

**POST** `/api/auth/register`

```json
{
  "email": "john.doe@example.com",
  "password": "SecurePass123!",
  "confirmPassword": "SecurePass123!",
  "firstName": "John",
  "lastName": "Doe",
  "acceptedTerms": true
}
```

### Fields

| Field | Type | Required | Constraints |
|-------|------|----------|-------------|
| `email` | string | ✅ | Valid email format, unique |
| `password` | string | ✅ | Min 8 chars, 1 uppercase, 1 number, 1 special char |
| `confirmPassword` | string | ✅ | Must match password |
| `firstName` | string | ❌ | Max 50 chars |
| `lastName` | string | ❌ | Max 50 chars |
| `acceptedTerms` | boolean | ✅ | Must be true |

### Response (201 Created)

```json
{
  "userId": "user-550e8400-e29b-41d4-a716-446655440000",
  "email": "john.doe@example.com",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "tokenExpires": "2026-04-20T09:30:00Z",
  "membershipTier": "Free",
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400-e29b-41d4-a716-446655440000"
    },
    "profile": {
      "href": "/api/users/user-550e8400-e29b-41d4-a716-446655440000/profile"
    },
    "searches": {
      "href": "/api/searches"
    },
    "membership": {
      "href": "/api/memberships"
    }
  }
}
```

### Error Responses

**400 Bad Request** - Validation failure
```json
{
  "status": 400,
  "title": "Validation Failed",
  "errors": {
    "email": ["Email is already registered"],
    "password": ["Password must contain at least one uppercase letter"]
  }
}
```

**409 Conflict** - Email already exists
```json
{
  "status": 409,
  "title": "Email Already Registered",
  "detail": "An account with this email already exists"
}
```

---

## 2. Login Request

**POST** `/api/auth/login`

```json
{
  "email": "john.doe@example.com",
  "password": "SecurePass123!",
  "rememberMe": true
}
```

### Fields

| Field | Type | Required | Constraints |
|-------|------|----------|-------------|
| `email` | string | ✅ | Valid email format |
| `password` | string | ✅ | - |
| `rememberMe` | boolean | ❌ | Default: false (extends token lifetime) |

### Response (200 OK)

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "tokenExpires": "2026-04-20T09:30:00Z",
  "refreshToken": "refresh-token-uuid",
  "user": {
    "id": "user-550e8400-e29b-41d4-a716-446655440000",
    "email": "john.doe@example.com",
    "firstName": "John",
    "lastName": "Doe",
    "membershipTier": "Premium"
  },
  "_links": {
    "self": {
      "href": "/api/auth/login"
    },
    "profile": {
      "href": "/api/users/user-550e8400-e29b-41d4-a716-446655440000/profile"
    },
    "searches": {
      "href": "/api/searches"
    },
    "logout": {
      "href": "/api/auth/logout",
      "method": "POST"
    }
  }
}
```

### Error Responses

**401 Unauthorized** - Invalid credentials
```json
{
  "status": 401,
  "title": "Invalid Credentials",
  "detail": "Email or password is incorrect"
}
```

**429 Too Many Requests** - Rate limit exceeded
```json
{
  "status": 429,
  "title": "Too Many Login Attempts",
  "detail": "Please try again in 15 minutes",
  "retryAfter": 900
}
```

---

## 3. Refresh Token Request

**POST** `/api/auth/refresh`

**Headers**:
```
Authorization: Bearer {expired-jwt-token}
```

```json
{
  "refreshToken": "refresh-token-uuid"
}
```

### Response (200 OK)

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "tokenExpires": "2026-04-20T10:30:00Z",
  "refreshToken": "new-refresh-token-uuid"
}
```

### Error Responses

**401 Unauthorized** - Invalid refresh token
```json
{
  "status": 401,
  "title": "Invalid Refresh Token",
  "detail": "Refresh token is invalid or expired"
}
```

---

## 4. Logout Request

**POST** `/api/auth/logout`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

**Body**: Empty or optional
```json
{
  "refreshToken": "refresh-token-uuid"
}
```

### Response (204 No Content)

Empty body

---

## 5. Password Reset Request

**POST** `/api/auth/password-reset/request`

```json
{
  "email": "john.doe@example.com"
}
```

### Response (200 OK)

```json
{
  "message": "If the email exists, a password reset link has been sent",
  "_links": {
    "login": {
      "href": "/api/auth/login"
    }
  }
}
```

**Note**: Always return 200 even if email doesn't exist (security best practice)

---

## 6. Password Reset Confirm

**POST** `/api/auth/password-reset/confirm`

```json
{
  "email": "john.doe@example.com",
  "resetToken": "reset-token-from-email",
  "newPassword": "NewSecurePass456!",
  "confirmPassword": "NewSecurePass456!"
}
```

### Response (200 OK)

```json
{
  "message": "Password successfully reset",
  "_links": {
    "login": {
      "href": "/api/auth/login"
    }
  }
}
```

### Error Responses

**400 Bad Request** - Invalid or expired token
```json
{
  "status": 400,
  "title": "Invalid Reset Token",
  "detail": "The password reset token is invalid or has expired"
}
```

---

## JWT Token Structure

### Token Claims

```json
{
  "sub": "user-550e8400-e29b-41d4-a716-446655440000",
  "email": "john.doe@example.com",
  "membershipTier": "Premium",
  "features": ["priority-search", "advanced-filters"],
  "iat": 1714466400,
  "exp": 1714470000,
  "iss": "https://holidayfinder.com",
  "aud": "holidayfinder-api"
}
```

### Claim Descriptions

| Claim | Description |
|-------|-------------|
| `sub` | User ID (subject) |
| `email` | User email address |
| `membershipTier` | Current tier (Free/Premium/VIP) |
| `features` | Feature flags for authorization |
| `iat` | Issued at (Unix timestamp) |
| `exp` | Expiration (Unix timestamp) |
| `iss` | Issuer (API domain) |
| `aud` | Audience (API identifier) |

### Token Lifetimes

- **Standard token**: 60 minutes
- **Remember me token**: 30 days
- **Refresh token**: 90 days

---

## Security Notes

1. **HTTPS Required**: All authentication endpoints must use HTTPS in production
2. **Rate Limiting**: Login and registration endpoints limited to 5 requests per minute per IP
3. **Password Storage**: Use ASP.NET Core Identity's PBKDF2 with 10,000 iterations
4. **Token Storage**: Client should store JWT in localStorage or sessionStorage (not cookies for WASM)
5. **CORS**: Configure CORS to allow only known client origins
6. **Validation**: All inputs validated with DataAnnotations before processing
