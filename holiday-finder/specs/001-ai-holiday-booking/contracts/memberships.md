# API Contracts: Memberships

**Feature**: 001-ai-holiday-booking  
**Date**: 2026-04-20  
**API Pattern**: REST + HATEOAS (HAL format)

## Overview

Contracts for membership tier management, upgrades, and feature access.

---

## 1. Get All Membership Tiers

**GET** `/api/memberships`

**Headers**: Optional (public endpoint, but customized if authenticated)
```
Authorization: Bearer {jwt-token}
```

### Response (200 OK)

```json
{
  "tiers": [
    {
      "id": "tier-free",
      "name": "Free",
      "displayName": "Free Explorer",
      "description": "Perfect for occasional travelers exploring options",
      "pricing": {
        "amount": 0,
        "currency": "USD",
        "billingCycle": "none",
        "annualOption": null
      },
      "features": [
        {
          "id": "basic-search",
          "name": "Basic AI Search",
          "description": "Search up to 3 destinations per query",
          "icon": "search"
        },
        {
          "id": "search-quota-3",
          "name": "3 Searches Per Day",
          "description": "Run up to 3 searches daily",
          "icon": "calendar"
        },
        {
          "id": "saved-criteria-5",
          "name": "5 Saved Searches",
          "description": "Save up to 5 favorite search criteria",
          "icon": "bookmark"
        }
      ],
      "limits": {
        "searchesPerDay": 3,
        "savedCriteria": 5,
        "destinations": 3,
        "aiProcessingTimeout": 10
      },
      "isCurrent": false,
      "_links": {
        "self": {
          "href": "/api/memberships/tier-free"
        }
      }
    },
    {
      "id": "tier-premium",
      "name": "Premium",
      "displayName": "Premium Explorer",
      "description": "Enhanced search with priority processing for frequent travelers",
      "pricing": {
        "amount": 9.99,
        "currency": "USD",
        "billingCycle": "monthly",
        "annualOption": {
          "amount": 95.99,
          "savings": 24
        }
      },
      "features": [
        {
          "id": "priority-search",
          "name": "Priority AI Processing",
          "description": "Get results 2x faster with priority queue",
          "icon": "zap",
          "badge": "New"
        },
        {
          "id": "advanced-filters",
          "name": "Advanced Filters",
          "description": "Filter by flight preferences, refundability, and more",
          "icon": "filter"
        },
        {
          "id": "search-quota-10",
          "name": "10 Searches Per Day",
          "description": "Run up to 10 searches daily",
          "icon": "calendar"
        },
        {
          "id": "saved-criteria-20",
          "name": "20 Saved Searches",
          "description": "Save up to 20 favorite search criteria",
          "icon": "bookmark"
        },
        {
          "id": "email-alerts",
          "name": "Email Alerts",
          "description": "Receive alerts when prices drop for saved searches",
          "icon": "bell"
        }
      ],
      "limits": {
        "searchesPerDay": 10,
        "savedCriteria": 20,
        "destinations": 10,
        "aiProcessingTimeout": 15
      },
      "popularBadge": true,
      "isCurrent": true,
      "_links": {
        "self": {
          "href": "/api/memberships/tier-premium"
        },
        "manage": {
          "href": "/api/users/user-550e8400/membership"
        },
        "cancel": {
          "href": "/api/users/user-550e8400/membership/cancel",
          "method": "POST"
        }
      }
    },
    {
      "id": "tier-vip",
      "name": "VIP",
      "displayName": "VIP Concierge",
      "description": "Ultimate experience with unlimited searches and dedicated support",
      "pricing": {
        "amount": 29.99,
        "currency": "USD",
        "billingCycle": "monthly",
        "annualOption": {
          "amount": 287.99,
          "savings": 72
        }
      },
      "features": [
        {
          "id": "unlimited-searches",
          "name": "Unlimited Searches",
          "description": "No daily search limits",
          "icon": "infinity"
        },
        {
          "id": "multi-model-comparison",
          "name": "Multi-Model AI Comparison",
          "description": "Compare results from multiple AI models",
          "icon": "layers",
          "badge": "Exclusive"
        },
        {
          "id": "priority-support",
          "name": "Priority Support",
          "description": "24/7 dedicated customer support",
          "icon": "headset"
        },
        {
          "id": "concierge-service",
          "name": "Concierge Booking",
          "description": "Personal assistance with bookings",
          "icon": "user-check"
        },
        {
          "id": "all-premium-features",
          "name": "All Premium Features",
          "description": "Includes all Premium tier benefits",
          "icon": "star"
        }
      ],
      "limits": {
        "searchesPerDay": -1,
        "savedCriteria": 50,
        "destinations": 10,
        "aiProcessingTimeout": 20
      },
      "isCurrent": false,
      "_links": {
        "self": {
          "href": "/api/memberships/tier-vip"
        },
        "upgrade": {
          "href": "/api/users/user-550e8400/membership/upgrade",
          "method": "POST"
        }
      }
    }
  ],
  "_links": {
    "self": {
      "href": "/api/memberships"
    },
    "myMembership": {
      "href": "/api/users/user-550e8400/membership"
    }
  }
}
```

**Note**: `isCurrent` is only included if user is authenticated

---

## 2. Get User's Current Membership

**GET** `/api/users/{userId}/membership`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

### Response (200 OK)

```json
{
  "userId": "user-550e8400-e29b-41d4-a716-446655440000",
  "tier": {
    "id": "tier-premium",
    "name": "Premium",
    "displayName": "Premium Explorer"
  },
  "subscription": {
    "status": "active",
    "startedAt": "2026-03-20T10:00:00Z",
    "expiresAt": "2027-03-20T10:00:00Z",
    "billingCycle": "annual",
    "renewsAutomatically": true,
    "nextBillingDate": "2027-03-20T00:00:00Z",
    "nextBillingAmount": 95.99,
    "paymentMethod": {
      "type": "card",
      "last4": "4242",
      "brand": "Visa",
      "expiryMonth": 12,
      "expiryYear": 2027
    }
  },
  "usage": {
    "searchesToday": 3,
    "searchQuotaRemaining": 7,
    "savedCriteria": 8,
    "quotaResetsAt": "2026-04-21T00:00:00Z"
  },
  "features": [
    "priority-search",
    "advanced-filters",
    "search-quota-10",
    "saved-criteria-20",
    "email-alerts"
  ],
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400/membership"
    },
    "tiers": {
      "href": "/api/memberships"
    },
    "upgrade": {
      "href": "/api/users/user-550e8400/membership/upgrade",
      "method": "POST"
    },
    "downgrade": {
      "href": "/api/users/user-550e8400/membership/downgrade",
      "method": "POST"
    },
    "cancel": {
      "href": "/api/users/user-550e8400/membership/cancel",
      "method": "POST"
    },
    "updatePayment": {
      "href": "/api/users/user-550e8400/membership/payment",
      "method": "PUT"
    },
    "invoices": {
      "href": "/api/users/user-550e8400/membership/invoices"
    }
  }
}
```

---

## 3. Upgrade Membership

**POST** `/api/users/{userId}/membership/upgrade`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

```json
{
  "targetTierId": "tier-premium",
  "billingCycle": "annual",
  "paymentMethod": {
    "type": "card",
    "cardToken": "tok_visa_4242424242424242",
    "saveCard": true
  },
  "promoCode": "SUMMER2026"
}
```

### Fields

| Field | Type | Required | Constraints |
|-------|------|----------|-------------|
| `targetTierId` | string | ✅ | Must be higher tier than current |
| `billingCycle` | string | ✅ | "monthly" or "annual" |
| `paymentMethod` | object | ✅ | Payment details (Stripe token) |
| `promoCode` | string | ❌ | Optional discount code |

### Response (200 OK)

```json
{
  "upgradedAt": "2026-04-20T09:00:00Z",
  "previousTier": "Free",
  "newTier": "Premium",
  "subscription": {
    "status": "active",
    "expiresAt": "2027-04-20T09:00:00Z",
    "billingCycle": "annual",
    "amountCharged": 95.99,
    "currency": "USD",
    "nextBillingDate": "2027-04-20T00:00:00Z"
  },
  "discount": {
    "code": "SUMMER2026",
    "amount": 10.00,
    "description": "10% off annual subscription"
  },
  "receipt": {
    "invoiceId": "inv_abc123",
    "invoiceUrl": "https://api.holidayfinder.com/invoices/inv_abc123.pdf"
  },
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400/membership"
    },
    "invoice": {
      "href": "/api/users/user-550e8400/membership/invoices/inv_abc123"
    },
    "startSearching": {
      "href": "/api/searches",
      "method": "POST"
    }
  }
}
```

### Error Responses

**400 Bad Request** - Cannot upgrade
```json
{
  "status": 400,
  "title": "Invalid Upgrade Request",
  "detail": "Cannot upgrade to a lower tier. Use downgrade endpoint instead.",
  "_links": {
    "downgrade": {
      "href": "/api/users/user-550e8400/membership/downgrade"
    }
  }
}
```

**402 Payment Required** - Payment failed
```json
{
  "status": 402,
  "title": "Payment Failed",
  "detail": "Your card was declined. Please use a different payment method.",
  "paymentError": {
    "code": "card_declined",
    "message": "Your card has insufficient funds"
  },
  "_links": {
    "retry": {
      "href": "/api/users/user-550e8400/membership/upgrade",
      "method": "POST"
    }
  }
}
```

---

## 4. Downgrade Membership

**POST** `/api/users/{userId}/membership/downgrade`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

```json
{
  "targetTierId": "tier-free",
  "effectiveDate": "end-of-billing-cycle",
  "reason": "Not using enough searches",
  "feedback": "Great service but too expensive for my usage"
}
```

### Fields

| Field | Type | Required | Constraints |
|-------|------|----------|-------------|
| `targetTierId` | string | ✅ | Must be lower tier than current |
| `effectiveDate` | string | ✅ | "immediate" or "end-of-billing-cycle" |
| `reason` | string | ❌ | Cancellation reason |
| `feedback` | string | ❌ | User feedback |

### Response (200 OK)

```json
{
  "scheduledAt": "2026-04-20T09:15:00Z",
  "currentTier": "Premium",
  "targetTier": "Free",
  "effectiveDate": "2027-03-20T10:00:00Z",
  "accessRetainedUntil": "2027-03-20T10:00:00Z",
  "refund": {
    "eligible": false,
    "reason": "More than 30 days since last payment"
  },
  "message": "Your membership will downgrade to Free on 2027-03-20. You'll continue to have Premium access until then.",
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400/membership"
    },
    "cancel": {
      "href": "/api/users/user-550e8400/membership/downgrade/cancel",
      "method": "POST",
      "description": "Cancel the scheduled downgrade"
    }
  }
}
```

---

## 5. Cancel Subscription

**POST** `/api/users/{userId}/membership/cancel`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

```json
{
  "effectiveDate": "immediate",
  "reason": "Found alternative service",
  "feedback": "Pricing is too high"
}
```

### Response (200 OK)

```json
{
  "cancelledAt": "2026-04-20T09:20:00Z",
  "tier": "Premium",
  "accessUntil": "2027-03-20T10:00:00Z",
  "refund": {
    "eligible": true,
    "amount": 47.99,
    "currency": "USD",
    "description": "Pro-rated refund for unused period",
    "processedWithin": "5-7 business days"
  },
  "message": "Your subscription has been cancelled. You'll continue to have Premium access until 2027-03-20.",
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400/membership"
    },
    "reactivate": {
      "href": "/api/users/user-550e8400/membership/reactivate",
      "method": "POST"
    }
  }
}
```

---

## 6. Update Payment Method

**PUT** `/api/users/{userId}/membership/payment`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

```json
{
  "paymentMethod": {
    "type": "card",
    "cardToken": "tok_visa_5555555555554444",
    "saveCard": true
  }
}
```

### Response (200 OK)

```json
{
  "updatedAt": "2026-04-20T09:25:00Z",
  "paymentMethod": {
    "type": "card",
    "last4": "4444",
    "brand": "Visa",
    "expiryMonth": 8,
    "expiryYear": 2028
  },
  "message": "Payment method updated successfully",
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400/membership/payment"
    },
    "membership": {
      "href": "/api/users/user-550e8400/membership"
    }
  }
}
```

---

## 7. Get Invoices

**GET** `/api/users/{userId}/membership/invoices`

**Headers**:
```
Authorization: Bearer {jwt-token}
```

### Response (200 OK)

```json
{
  "invoices": [
    {
      "invoiceId": "inv_2026_04",
      "date": "2026-04-01T00:00:00Z",
      "amount": 95.99,
      "currency": "USD",
      "status": "paid",
      "description": "Premium Explorer - Annual Subscription",
      "period": {
        "start": "2026-04-01",
        "end": "2027-03-31"
      },
      "_links": {
        "self": {
          "href": "/api/users/user-550e8400/membership/invoices/inv_2026_04"
        },
        "download": {
          "href": "/api/users/user-550e8400/membership/invoices/inv_2026_04.pdf"
        }
      }
    }
  ],
  "pagination": {
    "page": 1,
    "pageSize": 20,
    "totalItems": 3,
    "totalPages": 1
  },
  "_links": {
    "self": {
      "href": "/api/users/user-550e8400/membership/invoices"
    }
  }
}
```

---

## Business Rules

### Tier Transition Rules

| From | To | Effect | Refund |
|------|-----|--------|--------|
| Free | Premium/VIP | Immediate | N/A |
| Premium | VIP | Immediate, pro-rated | N/A |
| Premium | Free | End of cycle | No |
| VIP | Premium | End of cycle | Pro-rated |
| VIP | Free | End of cycle | Pro-rated |

### Feature Access

Features are checked via JWT token claims (`features` array).

**Authorization Example**:
```csharp
[Authorize(Policy = "RequireFeature:advanced-filters")]
public IActionResult AdvancedSearch() { }
```

### Payment Integration

- **Provider**: Stripe (or similar)
- **Flow**: Client generates token → API processes payment
- **Security**: PCI-compliant, no card details stored
- **Webhooks**: Handle payment events (success, failure, refund)

### Promo Codes

- Validated at upgrade time
- Can apply to first payment or recurring
- Discount types: Percentage or fixed amount
- Expiration dates enforced

---

## HATEOAS Links Reference

| Link Rel | Description | Method |
|----------|-------------|--------|
| `self` | Current resource | GET |
| `tiers` | All available tiers | GET |
| `myMembership` | User's current membership | GET |
| `manage` | Manage subscription | GET |
| `upgrade` | Upgrade to higher tier | POST |
| `downgrade` | Downgrade to lower tier | POST |
| `cancel` | Cancel subscription | POST |
| `reactivate` | Reactivate cancelled subscription | POST |
| `updatePayment` | Update payment method | PUT |
| `invoices` | Billing history | GET |
| `invoice` | Specific invoice | GET |
| `download` | Download invoice PDF | GET |
| `startSearching` | Navigate to search | GET/POST |
| `retry` | Retry failed payment | POST |
