# Challenge 09: RAG Agent

## Objective

Build a complete **Retrieval Augmented Generation (RAG) system** that intelligently combines local vector search and web search to provide accurate, well-cited answers grounded in retrieved information.

## What You'll Learn

- Complete RAG architecture: retrieve → rank → generate
- Intelligent retrieval strategy selection
- Combining multiple information sources (local + web)
- Re-ranking and relevance scoring
- Context window management for retrieved content
- Confidence scoring and uncertainty handling
- Production RAG patterns and best practices

## Prerequisites

- **Challenge 07 completed** — Vector search and semantic embeddings
- **Challenge 08 completed** — Web search and citation handling
- Strong understanding of information retrieval concepts
- Experience with the Agent Framework SDK

## Background

**RAG (Retrieval Augmented Generation)** solves a critical problem: LLMs have limited, static knowledge from their training data. RAG extends agents with dynamic, up-to-date, domain-specific knowledge.

**Traditional LLM:**
```
User: "What's our Q4 2025 revenue?"
Agent: "I don't have access to your company data."
```

**RAG Agent:**
```
User: "What's our Q4 2025 revenue?"
Agent: [retrieves from company database]
      "According to the Q4 2025 Financial Report [1], revenue was $12.5M..."
```

### RAG Pipeline

```
┌─────────────┐
│ User Query  │
└──────┬──────┘
       │
       ▼
┌─────────────────────────┐
│  Retrieval Strategy     │  ← Decide: local KB? web? both?
│  Selection              │
└──────┬──────────────────┘
       │
       ▼
┌─────────────────────────┐
│  Multi-Source Retrieval │  ← Query vector store + web search
│  (Parallel)             │
└──────┬──────────────────┘
       │
       ▼
┌─────────────────────────┐
│  Re-ranking &           │  ← Score relevance, combine sources
│  Filtering              │
└──────┬──────────────────┘
       │
       ▼
┌─────────────────────────┐
│  Context Injection      │  ← Build prompt with retrieved docs
│  (Top-K results)        │
└──────┬──────────────────┘
       │
       ▼
┌─────────────────────────┐
│  LLM Generation         │  ← Generate response using context
│  with Citations         │
└──────┬──────────────────┘
       │
       ▼
┌─────────────┐
│   Response  │
└─────────────┘
```

This challenge implements the complete pipeline.

## What to Build

A production-grade RAG agent that:
1. **Analyzes queries** to determine information needs
2. **Retrieves from multiple sources** (local vector DB + web) in parallel
3. **Re-ranks results** by relevance and credibility
4. **Manages context** to stay within token limits
5. **Generates responses** using top-ranked retrieved content
6. **Cites sources** properly for all factual claims
7. **Handles uncertainty** when information is insufficient

## File to Create

```
ai-ml/agent-orchestration/solutions/challenge09/
├── main.py                       # RAG agent implementation
├── rag_pipeline.py               # Complete RAG pipeline
├── retrieval_strategy.py         # Strategy selection logic
├── reranker.py                   # Result re-ranking
├── context_manager.py            # Context window management
├── setup_knowledge_base.py       # Initialize local KB
├── .env
└── requirements.txt
```

## Requirements

### 1. Enhanced Knowledge Base

Expand your knowledge base with more documents:

```python
# setup_knowledge_base.py
KNOWLEDGE_BASE = [
    # Technology (from Challenge 07)
    # ... existing tech docs ...
    
    # Company-specific (private knowledge)
    {
        "id": "company001",
        "title": "Product Roadmap 2025-2026",
        "content": """Our product roadmap focuses on three pillars:
        1. AI Integration - Adding ML capabilities to core features
        2. Performance - 50% latency reduction target
        3. Security - SOC 2 Type II compliance
        Expected launch: Q2 2026.""",
        "category": "company",
        "last_updated": "2025-12-15"
    },
    {
        "id": "company002",
        "title": "Support Guidelines",
        "content": """Customer support response times:
        - P0 (Critical): 1 hour
        - P1 (High): 4 hours
        - P2 (Medium): 24 hours
        - P3 (Low): 72 hours
        Escalation to engineering if unresolved in 48 hours.""",
        "category": "company",
        "last_updated": "2025-11-01"
    },
    # Add 10-15 more documents across categories
]
```

### 2. Retrieval Strategy

```python
# retrieval_strategy.py
from enum import Enum
from typing import List, Tuple
import logging

logger = logging.getLogger(__name__)

class RetrievalStrategy(Enum):
    """Available retrieval strategies."""
    LOCAL_ONLY = "local"
    WEB_ONLY = "web"
    HYBRID = "hybrid"
    LOCAL_FIRST_FALLBACK_WEB = "local_first"

class StrategySelector:
    """Intelligent retrieval strategy selection."""
    
    CURRENT_INFO_KEYWORDS = [
        "latest", "current", "recent", "today", "now",
        "2025", "2026", "breaking", "news"
    ]
    
    COMPANY_KEYWORDS = [
        "our", "company", "internal", "roadmap", "policy",
        "guidelines", "team", "support"
    ]
    
    def select_strategy(self, query: str) -> RetrievalStrategy:
        """
        Determine optimal retrieval strategy based on query.
        
        Args:
            query: User query
            
        Returns:
            RetrievalStrategy enum
        """
        query_lower = query.lower()
        
        # Check for company/internal information
        has_company_terms = any(kw in query_lower for kw in self.COMPANY_KEYWORDS)
        
        # Check for current information needs
        needs_current = any(kw in query_lower for kw in self.CURRENT_INFO_KEYWORDS)
        
        # Decision logic
        if has_company_terms and not needs_current:
            logger.info("Strategy: LOCAL_ONLY (company information)")
            return RetrievalStrategy.LOCAL_ONLY
        
        elif needs_current and not has_company_terms:
            logger.info("Strategy: WEB_ONLY (current information)")
            return RetrievalStrategy.WEB_ONLY
        
        elif needs_current and has_company_terms:
            logger.info("Strategy: HYBRID (current + company)")
            return RetrievalStrategy.HYBRID
        
        else:
            logger.info("Strategy: LOCAL_FIRST_FALLBACK_WEB (default)")
            return RetrievalStrategy.LOCAL_FIRST_FALLBACK_WEB
```

### 3. Re-Ranking System

```python
# reranker.py
from typing import List, Dict, Any
import logging

logger = logging.getLogger(__name__)

class ResultReranker:
    """Re-rank retrieval results by relevance and credibility."""
    
    # Source credibility scores
    CREDIBILITY_SCORES = {
        # Official/trusted sources
        "python.org": 10,
        "docs.python.org": 10,
        "github.com": 9,
        "stackoverflow.com": 8,
        
        # Company internal (highest trust)
        "company_kb": 10,
        
        # General web (medium trust)
        "medium.com": 6,
        "dev.to": 6,
        
        # Default
        "unknown": 5
    }
    
    def rerank(
        self,
        local_results: List[Dict],
        web_results: List[Dict],
        query: str
    ) -> List[Dict]:
        """
        Re-rank combined results.
        
        Args:
            local_results: Results from vector store
            web_results: Results from web search
            query: Original query for relevance scoring
            
        Returns:
            Combined and re-ranked results
        """
        all_results = []
        
        # Process local results
        for result in local_results:
            score = self._calculate_score(
                result,
                source_type="local",
                query=query
            )
            result['final_score'] = score
            result['source_type'] = 'local'
            all_results.append(result)
        
        # Process web results
        for result in web_results:
            score = self._calculate_score(
                result,
                source_type="web",
                query=query
            )
            result['final_score'] = score
            result['source_type'] = 'web'
            all_results.append(result)
        
        # Sort by score (highest first)
        ranked = sorted(all_results, key=lambda x: x['final_score'], reverse=True)
        
        logger.info(f"Re-ranked {len(ranked)} results")
        return ranked
    
    def _calculate_score(
        self,
        result: Dict,
        source_type: str,
        query: str
    ) -> float:
        """
        Calculate relevance score for a result.
        
        Combines:
        - Base relevance (distance/rank)
        - Source credibility
        - Recency (if available)
        - Query term overlap
        """
        score = 0.0
        
        # Base relevance (40% weight)
        if source_type == "local":
            # Lower distance = higher relevance
            distance = result.get('distance', 0.5)
            base_relevance = max(0, 1.0 - distance)
        else:
            # Web results: rank-based
            rank = result.get('rank', 5)
            base_relevance = 1.0 / rank
        
        score += base_relevance * 0.4
        
        # Source credibility (30% weight)
        source = result.get('source', 'unknown')
        credibility = self.CREDIBILITY_SCORES.get(source, 5) / 10.0
        score += credibility * 0.3
        
        # Recency (20% weight) - if available
        if 'last_updated' in result:
            recency_score = self._calculate_recency(result['last_updated'])
            score += recency_score * 0.2
        else:
            score += 0.1  # Neutral if no date
        
        # Query overlap (10% weight)
        content = (result.get('title', '') + ' ' + result.get('content', '')).lower()
        query_terms = set(query.lower().split())
        overlap = len([term for term in query_terms if term in content]) / max(len(query_terms), 1)
        score += overlap * 0.1
        
        return score
    
    def _calculate_recency(self, date_str: str) -> float:
        """Score based on how recent the information is."""
        from datetime import datetime, timedelta
        try:
            doc_date = datetime.fromisoformat(date_str)
            now = datetime.now()
            days_old = (now - doc_date).days
            
            # Recent (< 30 days) = 1.0
            # 30-180 days = 0.7
            # > 180 days = 0.3
            if days_old < 30:
                return 1.0
            elif days_old < 180:
                return 0.7
            else:
                return 0.3
        except:
            return 0.5  # Unknown date
```

### 4. Context Manager

```python
# context_manager.py
from typing import List, Dict
import tiktoken

class ContextManager:
    """Manage context window for LLM input."""
    
    def __init__(self, model: str = "gpt-4", max_context_tokens: int = 8000):
        """
        Initialize context manager.
        
        Args:
            model: Model name for token counting
            max_context_tokens: Maximum tokens for retrieved context
        """
        self.tokenizer = tiktoken.encoding_for_model(model)
        self.max_tokens = max_context_tokens
    
    def build_context(
        self,
        query: str,
        results: List[Dict],
        max_sources: int = 5
    ) -> Tuple[str, List[int]]:
        """
        Build context from retrieval results, respecting token limits.
        
        Args:
            query: User query
            results: Ranked retrieval results
            max_sources: Maximum number of sources to include
            
        Returns:
            (context_text, list_of_citation_ids)
        """
        context_parts = []
        citations = []
        total_tokens = 0
        
        for i, result in enumerate(results[:max_sources], 1):
            # Build source text
            source_text = f"""
[Source {i}]
Title: {result.get('title', 'Unknown')}
Content: {result.get('content', result.get('snippet', ''))}
"""
            
            # Count tokens
            tokens = len(self.tokenizer.encode(source_text))
            
            # Check if adding this source exceeds limit
            if total_tokens + tokens > self.max_tokens:
                break
            
            context_parts.append(source_text)
            citations.append(i)
            total_tokens += tokens
        
        context = "\n".join(context_parts)
        
        return context, citations
```

### 5. Complete RAG Pipeline

```python
# rag_pipeline.py
import logging
from typing import Dict, Any
from retrieval_strategy import StrategySelector, RetrievalStrategy
from reranker import ResultReranker
from context_manager import ContextManager

logger = logging.getLogger(__name__)

class RAGPipeline:
    """Complete RAG pipeline orchestration."""
    
    def __init__(self, vector_store, web_search_tool):
        self.vector_store = vector_store
        self.web_search = web_search_tool
        self.strategy_selector = StrategySelector()
        self.reranker = ResultReranker()
        self.context_manager = ContextManager()
    
    def retrieve_and_generate(self, query: str) -> Dict[str, Any]:
        """
        Execute complete RAG pipeline.
        
        Args:
            query: User query
            
        Returns:
            Dict with context, citations, and metadata
        """
        logger.info(f"RAG Pipeline: {query}")
        
        # Step 1: Select retrieval strategy
        strategy = self.strategy_selector.select_strategy(query)
        
        # Step 2: Retrieve from sources (parallel when hybrid)
        local_results = []
        web_results = []
        
        if strategy in [RetrievalStrategy.LOCAL_ONLY, RetrievalStrategy.HYBRID, 
                        RetrievalStrategy.LOCAL_FIRST_FALLBACK_WEB]:
            local_results = self.vector_store.semantic_search(query, top_k=5)
        
        if strategy in [RetrievalStrategy.WEB_ONLY, RetrievalStrategy.HYBRID]:
            web_results = self.web_search.search(query, num_results=5)
        
        elif strategy == RetrievalStrategy.LOCAL_FIRST_FALLBACK_WEB:
            # If local results are poor, fall back to web
            if not local_results or local_results[0].get('distance', 1.0) > 0.6:
                logger.info("Local results insufficient, falling back to web")
                web_results = self.web_search.search(query, num_results=5)
        
        # Step 3: Re-rank combined results
        ranked_results = self.reranker.rerank(local_results, web_results, query)
        
        # Step 4: Build context within token limits
        context, citation_ids = self.context_manager.build_context(
            query, ranked_results, max_sources=5
        )
        
        # Step 5: Return pipeline output
        return {
            "context": context,
            "citations": citation_ids,
            "sources": ranked_results[:len(citation_ids)],
            "strategy_used": strategy.value,
            "total_results_retrieved": len(ranked_results)
        }
```

### 6. RAG Agent

```python
# main.py - Agent integration
from azure.ai.projects.models import FunctionTool, ToolSet

# Initialize RAG pipeline
rag_pipeline = RAGPipeline(vector_store, web_search_tool)

def rag_search(query: str) -> str:
    """
    RAG search tool for agent.
    
    Executes complete RAG pipeline and returns formatted context.
    """
    try:
        result = rag_pipeline.retrieve_and_generate(query)
        
        return json.dumps({
            "context": result['context'],
            "sources": result['sources'],
            "strategy": result['strategy_used']
        })
    except Exception as e:
        logger.error(f"RAG search error: {e}")
        return json.dumps({"error": str(e)})

# Create agent with RAG capability
agent = project_client.agents.create_agent(
    model=os.environ["MODEL_DEPLOYMENT_NAME"],
    name="rag-agent",
    instructions="""You are an intelligent research assistant powered by RAG.
    
    When users ask questions:
    1. Use the rag_search function to retrieve relevant information
    2. The context will contain numbered sources [Source 1], [Source 2], etc.
    3. Generate your response based ONLY on the provided context
    4. Cite sources using [1], [2], etc. after relevant statements
    5. If the context doesn't contain sufficient information, say so clearly
    6. Never hallucinate information not in the retrieved context
    
    Response format:
    - Answer the question using retrieved information
    - Cite sources inline [1], [2]
    - Include Sources section at the end
    - If uncertain, acknowledge limitations
    
    CRITICAL: Only use information from the retrieved context. If context is 
    insufficient, say "I don't have enough information" rather than guessing.
    """,
    tools=[FunctionTool([rag_search])]
)
```

## Expected Output

```
=== Production RAG Agent Demo ===

Initializing RAG pipeline...
✓ Vector store loaded (15 documents)
✓ Web search tool ready
✓ Re-ranker configured
✓ Context manager initialized (8K token limit)

Creating RAG agent...
Agent created: rag-agent

=== Test 1: Company-Specific Query ===

User: What's our product roadmap for 2026?

[RAG PIPELINE]
→ Strategy Selection: LOCAL_ONLY (company keywords detected)
→ Retrieval: Vector store (5 results)
→ Re-ranking: Scored by relevance
→ Context Building: 3 sources (1,247 tokens)

[TOP SOURCES]
1. Product Roadmap 2025-2026 (score: 0.92, local)
2. Engineering Priorities Q1 2026 (score: 0.78, local)
3. Feature Launch Schedule (score: 0.71, local)

Agent: Based on our Product Roadmap 2025-2026 [1], we're focusing on three key pillars:

1. **AI Integration** - Adding machine learning capabilities to core features
2. **Performance Optimization** - Targeting a 50% reduction in latency
3. **Security Compliance** - Achieving SOC 2 Type II certification

These initiatives are expected to launch in Q2 2026 [1]. Our engineering team has 
prioritized the AI integration work for Q1 [2].

Sources:
[1] Product Roadmap 2025-2026 - Company KB
[2] Engineering Priorities Q1 2026 - Company KB

=== Test 2: Current Events Query ===

User: What are the latest AI breakthroughs announced this month?

[RAG PIPELINE]
→ Strategy Selection: WEB_ONLY (current + latest keywords)
→ Retrieval: Web search (5 results)
→ Re-ranking: Prioritizing credibility
→ Context Building: 4 sources (2,103 tokens)

[TOP SOURCES]
1. OpenAI Announces GPT-5 (score: 0.89, web - openai.com)
2. Google's Gemini 2.0 Released (score: 0.85, web - google.com)
3. Meta's LLaMA 4 Open Source (score: 0.82, web - meta.com)

Agent: The most significant AI breakthroughs announced recently include:

1. **OpenAI's GPT-5** [1] - Featuring multimodal capabilities and enhanced reasoning
2. **Google's Gemini 2.0** [2] - 2x performance improvement over previous version
3. **Meta's LLaMA 4** [3] - Open-source model rivaling proprietary competitors

These announcements represent major advances in large language model capabilities.

Sources:
[1] OpenAI Announces GPT-5 - openai.com
[2] Google's Gemini 2.0 Released - google.com
[3] Meta's LLaMA 4 Open Source - meta.com

=== Test 3: Hybrid Query ===

User: How should our team handle support escalations based on industry best practices?

[RAG PIPELINE]
→ Strategy Selection: HYBRID (our + best practices keywords)
→ Retrieval: Vector store (5) + Web search (5) in parallel
→ Re-ranking: Combined 10 results
→ Context Building: 5 sources (3,421 tokens)

[TOP SOURCES]
1. Support Guidelines (score: 0.94, local - company_kb)
2. Customer Support Best Practices 2026 (score: 0.81, web - zendesk.com)
3. Escalation Procedures (score: 0.79, local - company_kb)
4. SLA Management Guide (score: 0.73, web - freshdesk.com)

Agent: Our support escalation process [1] aligns with industry best practices [2]:

**Escalation Timeline:**
- P0 (Critical): 1 hour response
- P1 (High): 4 hours response
- P2 (Medium): 24 hours response
- P3 (Low): 72 hours response

If issues remain unresolved after 48 hours, escalate to engineering [1]. This follows 
industry standards where 80% of critical issues should be resolved within first contact [2].

**Best Practices:**
- Document escalation reasons clearly [2]
- Include reproduction steps and impact assessment [1]
- Maintain communication with customer during escalation [2]

Sources:
[1] Support Guidelines - Company KB
[2] Customer Support Best Practices 2026 - zendesk.com
[3] Escalation Procedures - Company KB

=== Test 4: Insufficient Information ===

User: What is the meaning of life?

[RAG PIPELINE]
→ Strategy Selection: LOCAL_FIRST_FALLBACK_WEB
→ Retrieval: Vector store (low relevance, distance > 0.7)
→ Fallback: Web search (philosophical, no definitive answer)
→ Context Building: 2 sources (821 tokens)

Agent: I don't have sufficient information in my knowledge base to provide a definitive 
answer to this philosophical question. While I found some general philosophical discussions 
[1][2], this question doesn't have a factual answer that can be retrieved from sources.

This is a deeply personal and philosophical question that different people and cultures 
answer differently based on their values, beliefs, and experiences.

=== RAG Quality Metrics ===

Total Queries: 4
Successful Retrievals: 4/4 (100%)
Average Sources per Response: 3.5
Citation Coverage: 100% (all claims cited)
Hallucination Rate: 0% (no unsourced claims)
Average Response Time: 2.3 seconds

✓ All responses grounded in retrieved information
✓ No hallucinated facts
✓ Proper source attribution
✓ Intelligent strategy selection
✓ Graceful handling of insufficient information

Production RAG system complete!
```

## Hints

<details>
<summary>Hint 1: Parallel retrieval</summary>

Use threading for parallel source retrieval:

```python
from concurrent.futures import ThreadPoolExecutor, as_completed

def retrieve_parallel(query):
    with ThreadPoolExecutor(max_workers=2) as executor:
        future_local = executor.submit(vector_store.search, query)
        future_web = executor.submit(web_search.search, query)
        
        local_results = future_local.result()
        web_results = future_web.result()
    
    return local_results, web_results
```

</details>

<details>
<summary>Hint 2: Confidence scoring</summary>

Add confidence scores to responses:

```python
def calculate_confidence(results, query):
    """
    Estimate confidence in retrieved results.
    
    High confidence: Top result distance < 0.3, multiple supporting sources
    Medium: Distance 0.3-0.6, some supporting sources
    Low: Distance > 0.6, contradictory sources
    """
    if not results:
        return "very_low"
    
    top_score = results[0]['final_score']
    
    if top_score > 0.8:
        return "high"
    elif top_score > 0.5:
        return "medium"
    else:
        return "low"
```

</details>

<details>
<summary>Hint 3: Hallucination detection</summary>

```python
def validate_response_against_context(response: str, context: str) -> bool:
    """Check if response contains information from context."""
    # Extract factual claims from response
    # Verify each claim appears in context
    # Flag potential hallucinations
    pass
```

</details>

<details>
<summary>Hint 4: Context compression</summary>

For long documents, use summarization:

```python
def compress_context(long_text: str, max_tokens: int) -> str:
    """Compress long context using extractive summarization."""
    sentences = sent_tokenize(long_text)
    
    # Score sentences by importance
    scored = score_sentences(sentences, query)
    
    # Select top sentences within token budget
    selected = select_within_budget(scored, max_tokens)
    
    return " ".join(selected)
```

</details>

<details>
<summary>Hint 5: Testing RAG quality</summary>

```python
test_cases = [
    {
        "query": "Product roadmap",
        "expected_strategy": "local",
        "must_cite_sources": True,
        "minimum_sources": 2,
        "no_hallucination": True
    }
]

def evaluate_rag_response(query, response, expected):
    # Check strategy selection
    # Verify citations present
    # Count sources
    # Check for hallucinations
    # Calculate accuracy score
    pass
```

</details>

## Testing

### RAG Pipeline Testing
- [ ] Strategy selection is appropriate for query types
- [ ] Retrieval returns relevant results
- [ ] Re-ranking improves result quality
- [ ] Context stays within token limits
- [ ] All sources are cited in responses

### Quality Metrics
- [ ] Precision: Retrieved results are relevant
- [ ] Recall: Important information is retrieved
- [ ] Citation coverage: All claims are sourced
- [ ] Hallucination rate: Near zero unsourced claims
- [ ] Response time: < 3 seconds end-to-end

### Edge Cases
- [ ] No relevant results found
- [ ] Contradictory information from sources
- [ ] Very long documents (context overflow)
- [ ] Ambiguous queries
- [ ] Multilingual queries (if supported)

## Going Further

### Extension 1: Advanced Re-ranking
Use cross-encoder models for better re-ranking:
```python
from sentence_transformers import CrossEncoder

cross_encoder = CrossEncoder('cross-encoder/ms-marco-MiniLM-L-6-v2')

def rerank_with_cross_encoder(query, results):
    pairs = [[query, r['content']] for r in results]
    scores = cross_encoder.predict(pairs)
    
    for result, score in zip(results, scores):
        result['cross_encoder_score'] = score
    
    return sorted(results, key=lambda x: x['cross_encoder_score'], reverse=True)
```

### Extension 2: Query Expansion
Expand queries for better retrieval:
```python
def expand_query(original_query):
    # Add synonyms
    # Generate related queries
    # Use query reformulation
    expanded = generate_variations(original_query)
    return expanded
```

### Extension 3: Conversational RAG
Handle multi-turn conversations:
```python
class ConversationalRAG:
    def __init__(self):
        self.conversation_history = []
    
    def retrieve_with_history(self, query, history):
        # Reformulate query considering conversation context
        # Retrieve relevant information
        # Update history with new turn
        pass
```

### Extension 4: RAG Evaluation Framework
```python
def evaluate_rag_system(test_dataset):
    metrics = {
        "retrieval_precision": [],
        "retrieval_recall": [],
        "answer_accuracy": [],
        "citation_f1": []
    }
    
    for test_case in test_dataset:
        # Run RAG pipeline
        # Compare to ground truth
        # Calculate metrics
        pass
    
    return aggregate_metrics(metrics)
```

## Common Issues

| Issue | Cause | Solution |
|-------|-------|----------|
| Poor retrieval quality | Bad embeddings or query | Improve document chunking, use better embedding model |
| Context overflow | Too many/long sources | Implement better context compression |
| Hallucinations | Agent not following instructions | Strengthen instructions, add validation |
| Slow response | Sequential retrieval | Parallelize local + web retrieval |
| Contradictory sources | Multiple conflicting results | Add source agreement scoring |

## References

- [RAG Survey Paper](https://arxiv.org/abs/2312.10997)
- [LlamaIndex Documentation](https://docs.llamaindex.ai/)
- [LangChain RAG Guide](https://python.langchain.com/docs/use_cases/question_answering/)
- [Evaluating RAG Systems](https://www.anthropic.com/index/evaluating-rag)
- [Advanced RAG Techniques](https://www.pinecone.io/learn/advanced-rag/)

## Key Takeaways

✅ **RAG extends LLM knowledge** — Dynamic, up-to-date, domain-specific information  
✅ **Pipeline is critical** — Retrieve → rank → generate with quality controls  
✅ **Multi-source retrieval is powerful** — Combine local KB + web for comprehensive answers  
✅ **Context management is essential** — Stay within token limits, prioritize quality  
✅ **Citation prevents hallucination** — Every claim must be sourced  
✅ **Strategy selection matters** — Choose retrieval approach based on query type  
✅ **Production RAG needs testing** — Evaluate precision, recall, hallucination rate  

## Next Steps

**Challenge 10: Sequential Workflow** begins the Multi-Agent Workflows phase. You'll learn to chain multiple specialized agents together using graph-based orchestration, creating powerful agent workflows that exceed what single agents can achieve.

**Concept Preview:** Instead of one agent doing everything, build workflows like: Document → Analyzer Agent → Summarizer Agent → Translation Agent → Final Output. Each agent specializes in one task!
