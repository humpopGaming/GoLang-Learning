# Challenge 08: Web Search

## Objective

Integrate **web search capabilities** into your agent, enabling it to find current information from the internet, handle search results, and properly cite sources.

## What You'll Learn

- Integrating web search tools with agents
- Handling search results and extracting relevant information
- Citation and source attribution best practices
- Combining web search with local knowledge (hybrid approach)
- When to use web search vs. local knowledge base
- Managing search result quality and relevance

## Prerequisites

- **Challenge 07 completed** — You understand semantic search and vector databases
- Basic understanding of web APIs and HTTP requests
- Familiarity with search engines and result ranking

## Background

Your local knowledge base (Challenge 07) has limitations:
- ❌ No current events or real-time data
- ❌ Limited to documents you've indexed
- ❌ Becomes stale over time
- ❌ May not cover all topics

**Web search** solves these problems by:
- ✅ Accessing billions of current web pages
- ✅ Finding information on any topic
- ✅ Always up-to-date (real-time data)
- ✅ Complementing your local knowledge

However, web search introduces new challenges:
- **Quality control** — Not all web content is reliable
- **Citation requirements** — Must credit sources
- **Ranking/relevance** — Choosing best results
- **Rate limits** — API quotas and costs

This challenge teaches you to build agents that responsibly use web search to ground their responses in current, cited information.

## What to Build

An agent with web search capabilities that:
1. Determines when web search is needed vs. using local knowledge
2. Performs web searches using an API
3. Evaluates and ranks search results
4. Extracts relevant information from results
5. Properly cites sources in responses
6. Combines web search with local knowledge

**Use Cases:**
- "What's the latest version of Python?" (current info needed)
- "Who won the 2024 Nobel Prize in Physics?" (recent events)
- "Current weather in Paris" (real-time data)
- "Best practices for React hooks in 2026" (current best practices)

## File to Create

```
ai-ml/agent-orchestration/solutions/challenge08/
├── main.py                    # Agent with web search
├── web_search_tool.py         # Web search implementation
├── search_api.py              # API integration (Bing/Google/mock)
├── citation_formatter.py      # Source citation utilities
├── .env
└── requirements.txt
```

## Requirements

### 1. Install Dependencies

```
# requirements.txt
azure-ai-projects
azure-identity
python-dotenv
requests
beautifulsoup4  # For parsing search results (optional)
```

### 2. Web Search Implementation

You have several options for web search:

**Option A: Mock Search API** (for learning without API keys)
```python
# search_api.py
class MockSearchAPI:
    """Mock search API for development/testing."""
    
    MOCK_RESULTS = {
        "python latest version": [
            {
                "title": "Python 3.13 Released - Python.org",
                "url": "https://www.python.org/downloads/release/python-3130/",
                "snippet": "Python 3.13.0 is the latest major release featuring improved performance..."
            },
            {
                "title": "What's New in Python 3.13",
                "url": "https://docs.python.org/3.13/whatsnew/3.13.html",
                "snippet": "Python 3.13 includes enhanced error messages, performance improvements..."
            }
        ],
        "react hooks best practices": [
            {
                "title": "React Hooks Best Practices 2026 - Dev.to",
                "url": "https://dev.to/react-hooks-2026",
                "snippet": "Modern React development relies heavily on hooks. Here are the top practices..."
            }
        ],
        # Add more mock results for different queries
    }
    
    def search(self, query: str, num_results: int = 5) -> List[Dict]:
        """
        Simulate web search.
        
        Args:
            query: Search query
            num_results: Number of results to return
            
        Returns:
            List of search result dictionaries
        """
        # Find closest matching mock results
        query_lower = query.lower()
        for key in self.MOCK_RESULTS:
            if key in query_lower or query_lower in key:
                return self.MOCK_RESULTS[key][:num_results]
        
        # Return generic result if no match
        return [{
            "title": f"Search results for: {query}",
            "url": "https://example.com/search",
            "snippet": "No specific mock data available for this query."
        }]
```

**Option B: Bing Search API** (if you have access)
```python
import requests
import os

class BingSearchAPI:
    """Bing Web Search API integration."""
    
    def __init__(self):
        self.api_key = os.environ.get("BING_SEARCH_API_KEY")
        self.endpoint = "https://api.bing.microsoft.com/v7.0/search"
    
    def search(self, query: str, num_results: int = 5) -> List[Dict]:
        headers = {"Ocp-Apim-Subscription-Key": self.api_key}
        params = {
            "q": query,
            "count": num_results,
            "textDecorations": False,
            "textFormat": "Raw"
        }
        
        response = requests.get(self.endpoint, headers=headers, params=params)
        response.raise_for_status()
        
        data = response.json()
        results = []
        
        for item in data.get("webPages", {}).get("value", []):
            results.append({
                "title": item.get("name"),
                "url": item.get("url"),
                "snippet": item.get("snippet", "")
            })
        
        return results
```

### 3. Web Search Tool for Agent

```python
# web_search_tool.py
import json
import logging
from typing import List, Dict, Any
from search_api import MockSearchAPI  # or BingSearchAPI

logger = logging.getLogger(__name__)

class WebSearchTool:
    """Web search tool with citation management."""
    
    def __init__(self, search_api=None):
        self.search_api = search_api or MockSearchAPI()
        logger.info("WebSearchTool initialized")
    
    def search_web(self, query: str, num_results: int = 5) -> str:
        """
        Search the web and return formatted results with citations.
        
        Args:
            query: Search query
            num_results: Number of results to return (default 5, max 10)
            
        Returns:
            JSON string with search results and citations
            
        Example:
            >>> search_web("Python latest version")
            '{
                "query": "Python latest version",
                "results": [
                    {
                        "title": "Python 3.13 Released",
                        "url": "https://python.org/...",
                        "snippet": "...",
                        "citation_id": 1
                    }
                ],
                "citations": {...}
            }'
        """
        try:
            # Validate inputs
            if not query or not isinstance(query, str):
                raise ValueError("Query must be a non-empty string")
            
            num_results = max(1, min(num_results, 10))  # Clamp between 1-10
            
            logger.info(f"Web search: '{query}' (requesting {num_results} results)")
            
            # Perform search
            raw_results = self.search_api.search(query, num_results)
            
            # Format results with citations
            formatted_results = []
            citations = {}
            
            for i, result in enumerate(raw_results, 1):
                citation_id = i
                
                formatted_results.append({
                    "title": result['title'],
                    "snippet": result['snippet'],
                    "citation_id": citation_id,
                    "relevance_rank": i
                })
                
                citations[citation_id] = {
                    "title": result['title'],
                    "url": result['url'],
                    "source": self._extract_domain(result['url'])
                }
            
            logger.info(f"Found {len(formatted_results)} results")
            
            return json.dumps({
                "query": query,
                "results": formatted_results,
                "citations": citations,
                "result_count": len(formatted_results)
            })
            
        except Exception as e:
            logger.error(f"Web search error: {e}")
            return json.dumps({"error": str(e)})
    
    def _extract_domain(self, url: str) -> str:
        """Extract domain from URL for citation."""
        try:
            from urllib.parse import urlparse
            parsed = urlparse(url)
            return parsed.netloc
        except:
            return "unknown"
```

### 4. Agent with Web Search

```python
from azure.ai.projects.models import FunctionTool, ToolSet

# Create web search tool instance
web_tool = WebSearchTool()

# Create function tool
functions = FunctionTool([web_tool.search_web])

toolset = ToolSet()
toolset.add(functions)

# Create agent
agent = project_client.agents.create_agent(
    model=os.environ["MODEL_DEPLOYMENT_NAME"],
    name="web-search-agent",
    instructions="""You are a helpful research assistant with web search capabilities.
    
    When users ask about:
    - Current events or recent information
    - Latest versions of software or technology
    - Real-time data (weather, stock prices, news)
    - Topics not in your training data
    
    Use the search_web function to find current information.
    
    IMPORTANT CITATION RULES:
    1. Always cite your sources using the provided citation_ids
    2. Format citations as [1], [2], etc. after relevant statements
    3. Include a "Sources:" section at the end listing all citations
    4. Never claim to have searched without actually calling search_web
    5. Be clear about what information comes from search vs your knowledge
    
    Example response format:
    "Python 3.13 is the latest version [1]. It includes performance improvements 
    and enhanced error messages [2].
    
    Sources:
    [1] Python 3.13 Released - python.org
    [2] What's New in Python 3.13 - docs.python.org"
    """,
    toolset=toolset
)
```

### 5. Hybrid Search: Local + Web

Create a tool that decides when to use local knowledge vs. web search:

```python
def smart_search(query: str, search_type: str = "auto") -> str:
    """
    Intelligently search using local knowledge base or web.
    
    Args:
        query: Search query
        search_type: "local", "web", or "auto" (default)
        
    Returns:
        JSON with search results and source type
    """
    if search_type == "auto":
        # Heuristics to determine search type
        current_info_keywords = ["latest", "current", "recent", "today", "2024", "2025", "2026"]
        needs_web = any(keyword in query.lower() for keyword in current_info_keywords)
        
        if needs_web:
            search_type = "web"
        else:
            search_type = "local"
    
    if search_type == "local":
        # Use vector store from Challenge 07
        results = vector_store.semantic_search(query)
        source = "local_knowledge_base"
    else:
        # Use web search
        results = web_tool.search_web(query)
        source = "web_search"
    
    return json.dumps({
        "source": source,
        "results": results,
        "query": query
    })
```

## Expected Output

```
=== Web Search Agent Demo ===

Initializing web search tool...
WebSearchTool initialized with MockSearchAPI

Creating agent with web search capabilities...
Agent created: web-search-agent

=== Test 1: Current Information Query ===

User: What's the latest version of Python?

[Determining search strategy]
Keywords detected: "latest" → Web search recommended

[TOOL CALL] search_web(query="latest Python version", num_results=5)

[WEB SEARCH] Query: "latest Python version"
Found 2 results:
  [1] Python 3.13 Released - python.org
  [2] What's New in Python 3.13 - docs.python.org

Agent: The latest version of Python is Python 3.13.0 [1]. This major release features 
improved performance, enhanced error messages, and several new language features [2].

Sources:
[1] Python 3.13 Released - python.org (https://www.python.org/downloads/release/python-3130/)
[2] What's New in Python 3.13 - docs.python.org (https://docs.python.org/3.13/whatsnew/3.13.html)

=== Test 2: Historical Query (No Web Search Needed) ===

User: What are functions in Python?

[Determining search strategy]
No current-info keywords detected → Local knowledge preferred

[TOOL CALL] search_knowledge_base(query="Python functions")
[LOCAL SEARCH] Found: "Python Functions Tutorial"

Agent: Functions in Python are reusable blocks of code defined using the 'def' keyword. 
They can accept parameters and return values, helping organize code and make it more 
maintainable [Local KB: Python Functions Tutorial].

=== Test 3: Hybrid Query ===

User: What are the best practices for React hooks in 2026?

[Determining search strategy]
Keywords: "2026", "best practices" → Hybrid approach

[TOOL CALL] search_web(query="React hooks best practices 2026")
[WEB SEARCH] Found current information
[TOOL CALL] search_knowledge_base(query="React hooks")
[LOCAL SEARCH] Found basic hook concepts

Agent: Based on current best practices in 2026 [1]:

1. **Always call hooks at the top level** - Never inside loops or conditions
2. **Use custom hooks for reusable logic** - Extract common patterns
3. **Optimize with useMemo and useCallback** - Prevent unnecessary re-renders
4. **Prefer composition over complex hooks** - Keep hooks simple and focused

These build on the fundamental React hooks concepts: useState for state management, 
useEffect for side effects, and useContext for sharing data [Local KB].

Sources:
[1] React Hooks Best Practices 2026 - dev.to

=== Citation Quality Check ===

✓ All factual claims have citations
✓ Citation IDs are properly formatted [1], [2]
✓ Sources section includes URLs
✓ Clear distinction between web and local sources
✓ No hallucinated search results

Demo complete!
```

## Hints

<details>
<summary>Hint 1: Choosing when to search</summary>

Create heuristics for search strategy:

```python
def should_use_web_search(query: str) -> bool:
    """Determine if web search is needed."""
    
    # Indicators of current information need
    time_indicators = ["latest", "current", "recent", "today", "now"]
    year_indicators = ["2024", "2025", "2026"]
    version_indicators = ["new version", "updated", "release"]
    
    query_lower = query.lower()
    
    return any(
        indicator in query_lower 
        for indicator in time_indicators + year_indicators + version_indicators
    )
```

</details>

<details>
<summary>Hint 2: Rate limiting web searches</summary>

Protect against excessive API calls:

```python
from datetime import datetime, timedelta
from collections import defaultdict

class RateLimitedSearchTool:
    def __init__(self, max_calls_per_minute=10):
        self.max_calls = max_calls_per_minute
        self.call_times = defaultdict(list)
    
    def search_web(self, query: str) -> str:
        now = datetime.now()
        minute_ago = now - timedelta(minutes=1)
        
        # Remove old calls
        self.call_times['search'] = [
            t for t in self.call_times['search'] if t > minute_ago
        ]
        
        # Check limit
        if len(self.call_times['search']) >= self.max_calls:
            return json.dumps({"error": "Rate limit exceeded. Try again in a minute."})
        
        # Record this call
        self.call_times['search'].append(now)
        
        # Perform search...
```

</details>

<details>
<summary>Hint 3: Citation formatting</summary>

Create consistent citation formats:

```python
def format_citation(citation_id: int, citation_data: Dict) -> str:
    """Format citation for output."""
    return f"[{citation_id}] {citation_data['title']} - {citation_data['source']}"

def create_sources_section(citations: Dict) -> str:
    """Create formatted sources section."""
    sources = ["Sources:"]
    for cid, data in sorted(citations.items()):
        sources.append(f"[{cid}] {data['title']} - {data['source']} ({data['url']})")
    return "\n".join(sources)
```

</details>

<details>
<summary>Hint 4: Result quality filtering</summary>

Filter low-quality results:

```python
def filter_results(results: List[Dict]) -> List[Dict]:
    """Remove low-quality search results."""
    filtered = []
    
    for result in results:
        # Skip if snippet is too short (likely low quality)
        if len(result.get('snippet', '')) < 50:
            continue
        
        # Skip certain domains (optional)
        blocked_domains = ['example.com', 'spam-site.com']
        if any(domain in result['url'] for domain in blocked_domains):
            continue
        
        filtered.append(result)
    
    return filtered
```

</details>

<details>
<summary>Hint 5: Testing web search agents</summary>

```python
test_cases = [
    {
        "query": "latest Python version",
        "should_use_web": True,
        "must_have_citations": True
    },
    {
        "query": "what are functions",
        "should_use_web": False,
        "can_use_local": True
    },
    {
        "query": "React best practices 2026",
        "should_use_web": True,
        "must_cite_year": True
    }
]

for test in test_cases:
    # Run query and validate
    response = run_agent_query(test['query'])
    if test['should_use_web']:
        assert "search_web" in tool_calls
    if test['must_have_citations']:
        assert re.search(r'\[\d+\]', response)  # Has [1], [2], etc.
```

</details>

## Testing

### Functionality Checklist
- [ ] Web search executes successfully
- [ ] Results are properly formatted
- [ ] Citations are included and correct
- [ ] Agent uses search when appropriate
- [ ] Agent avoids search when unnecessary
- [ ] Rate limiting works (if implemented)

### Citation Quality
- [ ] Every factual claim has a citation
- [ ] Citation IDs are consistent [1], [2], [3]
- [ ] Sources section lists all citations
- [ ] URLs are included and accessible
- [ ] No hallucinated sources

### Search Strategy
- [ ] Current-info queries trigger web search
- [ ] General knowledge queries use local KB
- [ ] Hybrid queries use both sources appropriately
- [ ] Clear indication of information source

## Going Further

### Extension 1: News Search
Add specialized news search:
```python
def search_news(query: str, days_back: int = 7) -> str:
    """Search recent news articles."""
    # Use news-specific API or filter by date
    pass
```

### Extension 2: Source Credibility Scoring
Rank sources by credibility:
```python
trusted_domains = {
    "python.org": 10,
    "github.com": 9,
    "stackoverflow.com": 8,
    # etc.
}

def get_credibility_score(url: str) -> int:
    domain = extract_domain(url)
    return trusted_domains.get(domain, 5)  # Default: 5/10
```

### Extension 3: Caching Search Results
Cache results to reduce API calls:
```python
from functools import lru_cache
import hashlib

@lru_cache(maxsize=100)
def cached_search(query_hash: str) -> str:
    # Cache results for repeated queries
    pass
```

### Extension 4: Multi-Source Aggregation
Combine results from multiple search APIs:
```python
def aggregate_search(query: str) -> str:
    bing_results = bing_api.search(query, 3)
    google_results = google_api.search(query, 3)
    
    # Merge, deduplicate, and re-rank
    combined = merge_and_rank(bing_results, google_results)
    return combined
```

## Common Issues

| Issue | Cause | Solution |
|-------|-------|----------|
| API quota exceeded | Too many calls | Implement rate limiting and caching |
| Poor result relevance | Generic queries | Make queries more specific |
| Missing citations | Agent hallucinating | Improve instructions, add validation |
| Slow searches | API latency | Add timeout, cache results |
| Wrong search strategy | Bad heuristics | Improve decision logic |

## References

- [Bing Web Search API](https://learn.microsoft.com/azure/cognitive-services/bing-web-search/)
- [Citation Best Practices](https://www.scribbr.com/citing-sources/)
- [Web Scraping Ethics](https://www.scraperapi.com/blog/web-scraping-best-practices/)
- [Search Quality Metrics](https://en.wikipedia.org/wiki/Evaluation_measures_(information_retrieval))

## Key Takeaways

✅ **Web search provides current information** — Complementing local knowledge bases  
✅ **Citation is critical** — Always attribute information to sources  
✅ **Search strategy matters** — Choose between local/web/hybrid intelligently  
✅ **Quality filtering is essential** — Not all search results are equally valuable  
✅ **Rate limiting protects resources** — Prevent excessive API usage  
✅ **Hybrid approach is powerful** — Combine local knowledge with web search  

## Next Steps

**Challenge 09: RAG Agent** brings everything together — combining file search (Challenge 07) and web search (Challenge 08) into a complete Retrieval Augmented Generation system with sophisticated retrieval strategies and response generation.

**Concept Preview:** Build an agent that automatically chooses the best information sources, retrieves relevant content, and generates comprehensive, well-cited answers grounded in both your knowledge base and the web!
