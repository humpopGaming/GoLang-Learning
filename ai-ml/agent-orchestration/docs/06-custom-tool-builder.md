# Challenge 06: Custom Tool Builder

## Objective

Design and implement a **production-grade custom tool** with comprehensive validation, error handling, documentation, and testing — learning best practices for building reliable tools that agents can use safely and effectively.

## What You'll Learn

- Designing tool interfaces that agents can use correctly
- Input validation and sanitization for security
- Comprehensive error handling and user-friendly error messages
- Writing tool documentation that guides both agents and developers
- Type hints and schema generation
- Testing tools thoroughly before agent integration
- Logging and observability for debugging
- Best practices for production tools

## Prerequisites

- **Challenges 01-05 completed** — You understand function tools, MCP, and Code Interpreter
- Solid Python programming skills
- Understanding of validation and error handling concepts
- Familiarity with type hints and documentation

## Background

In previous challenges, you created simple tools like `get_weather()` with basic implementations. Production tools need much more:

**❌ Simple Tool (Challenge 02):**
```python
def get_weather(location: str) -> str:
    return weather_data.get(location, "Unknown")
```

**✅ Production Tool (This Challenge):**
```python
def get_weather(location: str, unit: str = "celsius") -> WeatherResponse:
    """
    Get current weather with comprehensive error handling.
    
    Args:
        location: City name, properly validated
        unit: Temperature unit with allowed values
        
    Returns:
        WeatherResponse with structured data
        
    Raises:
        ValueError: If location invalid or unit not supported
        APIError: If weather service unavailable
    """
    # Validate inputs
    # Handle API errors
    # Log the call
    # Return structured response
```

This challenge teaches you to build **robust, reliable tools** that work correctly even with:
- Invalid inputs
- External API failures
- Edge cases
- Concurrent usage
- Production load

## What to Build

Build a complete **Database Lookup Tool** that:
1. Connects to a mock database (or real SQLite database)
2. Allows querying customer information safely (no SQL injection!)
3. Has comprehensive input validation
4. Provides detailed error messages
5. Includes logging and observability
6. Has complete test coverage
7. Works reliably with an agent

**Mock Database Schema:**
```sql
CREATE TABLE customers (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    status TEXT CHECK(status IN ('active', 'inactive', 'suspended')),
    created_date TEXT,
    total_orders INTEGER DEFAULT 0
);
```

## File to Create

```
ai-ml/agent-orchestration/solutions/challenge06/
├── main.py                 # Agent implementation using the tool
├── database_tool.py        # Your production-grade tool
├── test_database_tool.py   # Comprehensive test suite
├── database.db             # SQLite database (generated)
├── setup_database.py       # Database initialization script
├── .env
└── requirements.txt
```

## Requirements

### 1. Database Setup Script

Create `setup_database.py`:

```python
import sqlite3
from datetime import datetime, timedelta
import random

def setup_database(db_path: str = "database.db"):
    """Initialize the database with sample data."""
    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()
    
    # Create table
    cursor.execute("""
        CREATE TABLE IF NOT EXISTS customers (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL,
            status TEXT CHECK(status IN ('active', 'inactive', 'suspended')),
            created_date TEXT,
            total_orders INTEGER DEFAULT 0
        )
    """)
    
    # Sample data
    sample_customers = [
        ("Alice Johnson", "alice@example.com", "active", 45),
        ("Bob Smith", "bob@example.com", "active", 32),
        ("Charlie Brown", "charlie@example.com", "inactive", 5),
        ("Diana Prince", "diana@example.com", "active", 78),
        ("Eve Martinez", "eve@example.com", "suspended", 12),
    ]
    
    for name, email, status, orders in sample_customers:
        created = datetime.now() - timedelta(days=random.randint(30, 365))
        cursor.execute("""
            INSERT OR IGNORE INTO customers (name, email, status, created_date, total_orders)
            VALUES (?, ?, ?, ?, ?)
        """, (name, email, status, created.isoformat(), orders))
    
    conn.commit()
    conn.close()
    print(f"Database initialized at {db_path}")

if __name__ == "__main__":
    setup_database()
```

### 2. Production-Grade Tool Implementation

Create `database_tool.py`:

```python
import sqlite3
import logging
from typing import Optional, Dict, Any, List
from dataclasses import dataclass
from datetime import datetime
import json

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

@dataclass
class CustomerRecord:
    """Structured customer data."""
    id: int
    name: str
    email: str
    status: str
    created_date: str
    total_orders: int
    
    def to_dict(self) -> Dict[str, Any]:
        return {
            "id": self.id,
            "name": self.name,
            "email": self.email,
            "status": self.status,
            "created_date": self.created_date,
            "total_orders": self.total_orders
        }

class DatabaseToolError(Exception):
    """Base exception for database tool errors."""
    pass

class ValidationError(DatabaseToolError):
    """Raised when input validation fails."""
    pass

class DatabaseError(DatabaseToolError):
    """Raised when database operations fail."""
    pass

class DatabaseTool:
    """Production-grade database lookup tool with comprehensive error handling."""
    
    VALID_STATUSES = ['active', 'inactive', 'suspended']
    MAX_RESULTS = 100  # Prevent excessive result sets
    
    def __init__(self, db_path: str = "database.db"):
        """
        Initialize database tool.
        
        Args:
            db_path: Path to SQLite database file
            
        Raises:
            DatabaseError: If database cannot be accessed
        """
        self.db_path = db_path
        self._validate_database()
        logger.info(f"DatabaseTool initialized with database: {db_path}")
    
    def _validate_database(self):
        """Ensure database exists and is accessible."""
        try:
            conn = sqlite3.connect(self.db_path)
            cursor = conn.cursor()
            cursor.execute("SELECT name FROM sqlite_master WHERE type='table' AND name='customers'")
            if not cursor.fetchone():
                raise DatabaseError("Customers table not found. Run setup_database.py first.")
            conn.close()
        except sqlite3.Error as e:
            raise DatabaseError(f"Database validation failed: {e}")
    
    def lookup_customer_by_email(self, email: str) -> str:
        """
        Look up customer by email address.
        
        This function is exposed to AI agents as a tool.
        
        Args:
            email: Customer email address (must be valid email format)
        
        Returns:
            JSON string with customer data or error message
            
        Examples:
            >>> tool.lookup_customer_by_email("alice@example.com")
            '{"id": 1, "name": "Alice Johnson", "email": "alice@example.com", ...}'
            
            >>> tool.lookup_customer_by_email("invalid")
            '{"error": "Invalid email format"}'
        """
        try:
            # Validate input
            if not email or not isinstance(email, str):
                raise ValidationError("Email must be a non-empty string")
            
            if "@" not in email or "." not in email.split("@")[-1]:
                raise ValidationError("Invalid email format")
            
            email = email.strip().lower()
            
            # Log the lookup
            logger.info(f"Looking up customer by email: {email}")
            
            # Execute query (parameterized to prevent SQL injection)
            conn = sqlite3.connect(self.db_path)
            cursor = conn.cursor()
            cursor.execute(
                "SELECT id, name, email, status, created_date, total_orders FROM customers WHERE LOWER(email) = ?",
                (email,)
            )
            
            row = cursor.fetchone()
            conn.close()
            
            if row:
                customer = CustomerRecord(*row)
                logger.info(f"Customer found: {customer.name}")
                return json.dumps(customer.to_dict())
            else:
                logger.info(f"No customer found for email: {email}")
                return json.dumps({"error": "Customer not found", "email": email})
                
        except ValidationError as e:
            logger.warning(f"Validation error: {e}")
            return json.dumps({"error": str(e)})
        except sqlite3.Error as e:
            logger.error(f"Database error: {e}")
            return json.dumps({"error": "Database query failed", "details": str(e)})
        except Exception as e:
            logger.error(f"Unexpected error: {e}")
            return json.dumps({"error": "Internal error occurred"})
    
    def search_customers_by_status(self, status: str) -> str:
        """
        Find all customers with a specific status.
        
        Args:
            status: Customer status ('active', 'inactive', or 'suspended')
        
        Returns:
            JSON string with list of matching customers or error
            
        Examples:
            >>> tool.search_customers_by_status("active")
            '[{"id": 1, "name": "Alice Johnson", ...}, ...]'
        """
        try:
            # Validate status
            if not status or not isinstance(status, str):
                raise ValidationError("Status must be a non-empty string")
            
            status = status.strip().lower()
            if status not in self.VALID_STATUSES:
                raise ValidationError(
                    f"Invalid status. Must be one of: {', '.join(self.VALID_STATUSES)}"
                )
            
            logger.info(f"Searching for customers with status: {status}")
            
            # Execute query
            conn = sqlite3.connect(self.db_path)
            cursor = conn.cursor()
            cursor.execute(
                """SELECT id, name, email, status, created_date, total_orders 
                   FROM customers 
                   WHERE status = ? 
                   LIMIT ?""",
                (status, self.MAX_RESULTS)
            )
            
            rows = cursor.fetchall()
            conn.close()
            
            customers = [CustomerRecord(*row).to_dict() for row in rows]
            logger.info(f"Found {len(customers)} customers with status '{status}'")
            
            return json.dumps({
                "customers": customers,
                "count": len(customers),
                "status_filter": status
            })
            
        except ValidationError as e:
            logger.warning(f"Validation error: {e}")
            return json.dumps({"error": str(e)})
        except sqlite3.Error as e:
            logger.error(f"Database error: {e}")
            return json.dumps({"error": "Database query failed"})
        except Exception as e:
            logger.error(f"Unexpected error: {e}")
            return json.dumps({"error": "Internal error occurred"})
    
    def get_customer_summary(self) -> str:
        """
        Get summary statistics about all customers.
        
        Returns:
            JSON string with customer statistics
        """
        try:
            logger.info("Generating customer summary")
            
            conn = sqlite3.connect(self.db_path)
            cursor = conn.cursor()
            
            # Get counts by status
            cursor.execute("""
                SELECT status, COUNT(*) as count, SUM(total_orders) as total_orders
                FROM customers
                GROUP BY status
            """)
            
            status_stats = {}
            for status, count, orders in cursor.fetchall():
                status_stats[status] = {"count": count, "total_orders": orders or 0}
            
            # Get overall stats
            cursor.execute("SELECT COUNT(*), SUM(total_orders) FROM customers")
            total_customers, total_orders = cursor.fetchone()
            
            conn.close()
            
            summary = {
                "total_customers": total_customers,
                "total_orders": total_orders or 0,
                "by_status": status_stats,
                "generated_at": datetime.now().isoformat()
            }
            
            logger.info(f"Summary generated: {total_customers} customers")
            return json.dumps(summary)
            
        except sqlite3.Error as e:
            logger.error(f"Database error: {e}")
            return json.dumps({"error": "Failed to generate summary"})
```

### 3. Comprehensive Test Suite

Create `test_database_tool.py`:

```python
import unittest
import os
import json
from database_tool import DatabaseTool, ValidationError
from setup_database import setup_database

class TestDatabaseTool(unittest.TestCase):
    """Comprehensive test suite for DatabaseTool."""
    
    @classmethod
    def setUpClass(cls):
        """Set up test database once for all tests."""
        cls.test_db = "test_database.db"
        if os.path.exists(cls.test_db):
            os.remove(cls.test_db)
        setup_database(cls.test_db)
        cls.tool = DatabaseTool(cls.test_db)
    
    @classmethod
    def tearDownClass(cls):
        """Clean up test database."""
        if os.path.exists(cls.test_db):
            os.remove(cls.test_db)
    
    def test_lookup_valid_email(self):
        """Test looking up customer with valid email."""
        result = self.tool.lookup_customer_by_email("alice@example.com")
        data = json.loads(result)
        self.assertIn("name", data)
        self.assertEqual(data["email"], "alice@example.com")
    
    def test_lookup_invalid_email_format(self):
        """Test lookup with invalid email format."""
        result = self.tool.lookup_customer_by_email("not-an-email")
        data = json.loads(result)
        self.assertIn("error", data)
    
    def test_lookup_nonexistent_customer(self):
        """Test lookup for customer that doesn't exist."""
        result = self.tool.lookup_customer_by_email("nobody@example.com")
        data = json.loads(result)
        self.assertIn("error", data)
        self.assertEqual(data["error"], "Customer not found")
    
    def test_search_active_status(self):
        """Test searching for active customers."""
        result = self.tool.search_customers_by_status("active")
        data = json.loads(result)
        self.assertIn("customers", data)
        self.assertGreater(data["count"], 0)
    
    def test_search_invalid_status(self):
        """Test search with invalid status."""
        result = self.tool.search_customers_by_status("invalid_status")
        data = json.loads(result)
        self.assertIn("error", data)
    
    def test_get_summary(self):
        """Test customer summary generation."""
        result = self.tool.get_customer_summary()
        data = json.loads(result)
        self.assertIn("total_customers", data)
        self.assertIn("by_status", data)
    
    def test_sql_injection_protection(self):
        """Test that SQL injection attempts are blocked."""
        malicious_input = "'; DROP TABLE customers; --"
        result = self.tool.lookup_customer_by_email(malicious_input)
        data = json.loads(result)
        # Should return error, not execute SQL
        self.assertIn("error", data)

if __name__ == "__main__":
    unittest.main()
```

### 4. Agent Integration

Create `main.py` that uses your production tool:

```python
import os
from dotenv import load_dotenv
from azure.ai.projects import AIProjectClient
from azure.ai.projects.models import FunctionTool, ToolSet
from azure.identity import DefaultAzureCredential
from database_tool import DatabaseTool
from setup_database import setup_database

# Initialize database
setup_database()
db_tool = DatabaseTool()

# Create tool from the methods
functions = FunctionTool([
    db_tool.lookup_customer_by_email,
    db_tool.search_customers_by_status,
    db_tool.get_customer_summary
])

toolset = ToolSet()
toolset.add(functions)

# TODO: Create agent, run test queries, demonstrate tool usage
```

### 5. Expected Behavior

Your tool must handle:
- ✅ Valid inputs correctly
- ✅ Invalid inputs gracefully
- ✅ SQL injection attempts safely
- ✅ Database errors without crashing
- ✅ Edge cases (empty strings, None, etc.)
- ✅ Logging for debugging
- ✅ Clear error messages

## Expected Output

```
=== Production Database Tool Demo ===

Initializing database...
Database initialized at database.db

Running test suite...
......
----------------------------------------------------------------------
Ran 7 tests in 0.123s

OK

All tests passed! ✓

Creating agent with production database tool...
Agent created with 3 database tools

=== Test Query 1 ===
User: Look up customer alice@example.com

[INFO] Looking up customer by email: alice@example.com
[INFO] Customer found: Alice Johnson
[TOOL CALL] lookup_customer_by_email(email='alice@example.com')

Agent: I found Alice Johnson in the system:
- Email: alice@example.com
- Status: active
- Total Orders: 45
- Customer since: 2024-03-15

=== Test Query 2 ===
User: Show me all suspended customers

[INFO] Searching for customers with status: suspended
[INFO] Found 1 customers with status 'suspended'
[TOOL CALL] search_customers_by_status(status='suspended')

Agent: There is currently 1 suspended customer:
- Eve Martinez (eve@example.com) - 12 orders

=== Test Query 3: SQL Injection Attempt ===
User: Look up '; DROP TABLE customers; --@example.com

[WARNING] Validation error: Invalid email format
[TOOL CALL] lookup_customer_by_email(email="'; DROP TABLE customers; --@example.com")
[TOOL RESULT] {"error": "Invalid email format"}

Agent: That doesn't appear to be a valid email address. Please provide a proper email format like name@domain.com.

✓ SQL injection blocked successfully!

=== Summary ===
✓ Tool handles valid inputs correctly
✓ Tool validates and rejects invalid inputs
✓ Tool prevents SQL injection
✓ Tool provides clear error messages
✓ Tool logs operations for debugging
✓ Agent uses tool appropriately

Production-grade tool complete!
```

## Hints

<details>
<summary>Hint 1: Validation patterns</summary>

Create reusable validation functions:

```python
def validate_email(email: str) -> str:
    """Validate and normalize email."""
    if not email or not isinstance(email, str):
        raise ValidationError("Email must be a non-empty string")
    
    email = email.strip().lower()
    
    # Basic email format check
    if "@" not in email:
        raise ValidationError("Email must contain @")
    
    local, domain = email.split("@", 1)
    if not local or not domain or "." not in domain:
        raise ValidationError("Invalid email format")
    
    return email
```

</details>

<details>
<summary>Hint 2: Error handling hierarchy</summary>

Create specific exception types:

```python
class ToolError(Exception):
    """Base exception."""
    pass

class ValidationError(ToolError):
    """Input validation failed."""
    pass

class DatabaseError(ToolError):
    """Database operation failed."""
    pass

class ConfigurationError(ToolError):
    """Tool misconfigured."""
    pass
```

</details>

<details>
<summary>Hint 3: Logging best practices</summary>

```python
import logging

logger = logging.getLogger(__name__)

# Different levels for different situations
logger.debug("Detailed diagnostic info")      # Development
logger.info("Normal operation logged")        # Audit trail
logger.warning("Unexpected but handled")      # Investigate
logger.error("Operation failed")              # Fix required
logger.critical("System unstable")            # Immediate action
```

</details>

<details>
<summary>Hint 4: Return structured data</summary>

Always return JSON strings from tools:

```python
# ✓ Good - structured, parseable
return json.dumps({
    "status": "success",
    "data": customer_dict,
    "timestamp": datetime.now().isoformat()
})

# ✗ Bad - unstructured string
return f"Customer: {name}, Email: {email}"
```

</details>

<details>
<summary>Hint 5: Testing edge cases</summary>

Test the extremes:

```python
# Null/empty
lookup_customer_by_email("")
lookup_customer_by_email(None)

# Invalid types
lookup_customer_by_email(123)
lookup_customer_by_email(["not", "a", "string"])

# Boundary values
search_customers_by_status("a" * 1000)  # Very long input

# Special characters
lookup_customer_by_email("test@domain.com'; DROP TABLE--")
```

</details>

## Testing

### Test Coverage Checklist
- [ ] Valid inputs produce correct results
- [ ] Invalid inputs produce clear error messages
- [ ] Edge cases handled (empty, None, wrong types)
- [ ] SQL injection attempts blocked
- [ ] Database errors caught and handled
- [ ] All functions have tests
- [ ] Test coverage > 80%

### Production Readiness
- [ ] Comprehensive documentation (docstrings)
- [ ] Type hints on all functions
- [ ] Logging at appropriate levels
- [ ] Error messages are user-friendly
- [ ] No secrets in code (use environment variables)
- [ ] Resource cleanup (database connections closed)

## Going Further

### Extension 1: Rate Limiting
Add rate limiting to prevent abuse:
```python
from collections import defaultdict
from datetime import datetime, timedelta

class RateLimitedTool:
    def __init__(self):
        self.calls = defaultdict(list)
        self.limit = 10  # calls per minute
    
    def check_rate_limit(self, user_id: str):
        now = datetime.now()
        minute_ago = now - timedelta(minutes=1)
        
        # Remove old calls
        self.calls[user_id] = [t for t in self.calls[user_id] if t > minute_ago]
        
        if len(self.calls[user_id]) >= self.limit:
            raise RateLimitError("Rate limit exceeded")
        
        self.calls[user_id].append(now)
```

### Extension 2: Caching
Add caching for frequently accessed data:
```python
from functools import lru_cache
from datetime import datetime, timedelta

@lru_cache(maxsize=100)
def lookup_customer_cached(email: str) -> str:
    # Cache results for 5 minutes
    return self._lookup_customer_internal(email)
```

### Extension 3: Metrics Collection
Track tool usage metrics:
```python
from dataclasses import dataclass
from typing import Dict

@dataclass
class ToolMetrics:
    total_calls: int = 0
    successful_calls: int = 0
    failed_calls: int = 0
    avg_response_time: float = 0.0
    
    def record_call(self, success: bool, duration: float):
        self.total_calls += 1
        if success:
            self.successful_calls += 1
        else:
            self.failed_calls += 1
        # Update average...
```

### Extension 4: Async Support
Make tools async for better performance:
```python
import aiosqlite

async def lookup_customer_async(self, email: str) -> str:
    """Async version for concurrent operations."""
    async with aiosqlite.connect(self.db_path) as db:
        cursor = await db.execute(
            "SELECT * FROM customers WHERE email = ?",
            (email,)
        )
        row = await cursor.fetchone()
        return json.dumps(row) if row else json.dumps({"error": "Not found"})
```

## Common Issues

| Issue | Cause | Solution |
|-------|-------|----------|
| SQL injection vulnerability | String concatenation in queries | Always use parameterized queries |
| Poor error messages | Generic exceptions | Create specific error types with details |
| Unclosed connections | No cleanup | Use context managers or try/finally |
| Type errors | No validation | Add type hints and runtime validation |
| No observability | Missing logging | Add structured logging throughout |

## References

- [SQL Injection Prevention](https://owasp.org/www-community/attacks/SQL_Injection)
- [Python Type Hints](https://docs.python.org/3/library/typing.html)
- [Logging Best Practices](https://docs.python.org/3/howto/logging.html)
- [Unit Testing in Python](https://docs.python.org/3/library/unittest.html)
- [SQLite3 Module](https://docs.python.org/3/library/sqlite3.html)

## Key Takeaways

✅ **Validation is critical** — Never trust inputs, always validate  
✅ **Use parameterized queries** — Prevents SQL injection  
✅ **Structured errors help debugging** — Clear error types and messages  
✅ **Logging enables observability** — Track what's happening in production  
✅ **Tests catch issues early** — Comprehensive test suites are essential  
✅ **Documentation serves everyone** — Helps agents AND developers  
✅ **Production tools need defense** — Handle edge cases, errors, and attacks  

## Next Steps

**Challenge 07: File Search** introduces vector databases and semantic search — you'll build RAG (Retrieval Augmented Generation) capabilities that let agents search through documents using meaning, not just keywords.

**Concept Preview:** Instead of exact keyword matching, semantic search finds "Python coding tutorials" when searching for "learn to program in Python" — game-changing for knowledge agents!
