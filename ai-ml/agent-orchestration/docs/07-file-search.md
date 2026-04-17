# Challenge 07: File Search

## Objective

Build a **semantic file search system** using vector embeddings and a local vector database, enabling your agent to find relevant documents based on meaning rather than exact keyword matches.

## What You'll Learn

- What vector embeddings are and why they enable semantic search
- Setting up a local vector database (ChromaDB)
- Creating embeddings from text documents
- Performing semantic similarity search
- Integrating vector search with agents
- When semantic search beats keyword search
- RAG (Retrieval Augmented Generation) fundamentals

## Prerequisites

- **Challenge 06 completed** — You can build production-grade tools
- Understanding of search concepts (keywords, ranking)
- Basic knowledge of machine learning concepts (helpful but not required)
- Familiarity with JSON and text processing

## Background

Traditional search uses **exact keywords**:
- User searches: "Python programming tutorial"
- Finds documents containing exactly those words
- Misses: "Learn to code in Python", "Python coding guide", "Introduction to Python development"

**Semantic search** uses **meaning**:
- Converts text to **vector embeddings** (arrays of numbers representing meaning)
- Searches for similar *meanings*, not just matching words
- Finds relevant documents even with different wording

Example:
```
Query: "How do I make coffee?"
Vector: [0.23, -0.15, 0.67, ..., 0.42]  (384 dimensions)

Documents:
"Brewing coffee guide"     → [0.25, -0.13, 0.65, ..., 0.45]  ← High similarity! ✓
"Espresso machine manual"  → [0.22, -0.11, 0.63, ..., 0.41]  ← High similarity! ✓
"Tea preparation steps"    → [-0.15, 0.32, -0.21, ..., 0.12] ← Low similarity
```

This is the foundation of **RAG (Retrieval Augmented Generation)**:
1. User asks a question
2. System finds relevant documents (retrieval)
3. Agent uses documents to generate answer (augmentation)
4. Response is grounded in your data, not just training data

## What to Build

A knowledge base system that:
1. Creates a vector database from a collection of documents
2. Generates embeddings for each document
3. Performs semantic search to find relevant documents
4. Integrates with an agent to answer questions using retrieved knowledge
5. Demonstrates the difference between keyword and semantic search

**Sample Knowledge Base**: Technology documentation covering Python, JavaScript, databases, cloud computing, etc.

## File to Create

```
ai-ml/agent-orchestration/solutions/challenge07/
├── main.py                    # Agent with semantic search
├── vector_store.py            # Vector database implementation
├── sample_docs.py             # Sample documents for knowledge base
├── test_search.py             # Search comparison tests
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
chromadb>=0.4.0
sentence-transformers>=2.2.0
```

### 2. Create Sample Documents

```python
# sample_docs.py
TECH_DOCS = [
    {
        "id": "doc001",
        "title": "Python Functions Tutorial",
        "content": """
        Python functions are reusable blocks of code defined using the 'def' keyword.
        Functions can accept parameters and return values. They help organize code
        and make it more maintainable. Example: def greet(name): return f"Hello {name}"
        """
    },
    {
        "id": "doc002",
        "title": "JavaScript Async/Await Guide",
        "content": """
        JavaScript async/await provides a cleaner syntax for handling asynchronous
        operations. It's built on promises and makes async code look synchronous.
        Use 'async' to declare async functions and 'await' to wait for promises.
        """
    },
    {
        "id": "doc003",
        "title": "SQL Database Basics",
        "content": """
        SQL (Structured Query Language) is used to manage relational databases.
        Common operations include SELECT for querying, INSERT for adding data,
        UPDATE for modifying records, and DELETE for removing data. Tables store
        data in rows and columns.
        """
    },
    {
        "id": "doc004",
        "title": "Cloud Computing Introduction",
        "content": """
        Cloud computing delivers computing services over the internet including
        storage, databases, servers, and software. Major providers are AWS, Azure,
        and Google Cloud. Benefits include scalability, cost-efficiency, and
        reduced infrastructure management.
        """
    },
    {
        "id": "doc005",
        "title": "REST API Design",
        "content": """
        REST APIs use HTTP methods (GET, POST, PUT, DELETE) to perform operations
        on resources. They are stateless and use URLs to identify resources.
        Common response formats include JSON and XML. Status codes indicate
        success (200) or errors (404, 500).
        """
    },
    {
        "id": "doc006",
        "title": "Git Version Control",
        "content": """
        Git is a distributed version control system for tracking code changes.
        Key commands include 'git commit' to save changes, 'git push' to upload,
        'git pull' to download, and 'git branch' for parallel development.
        GitHub and GitLab host Git repositories.
        """
    },
    {
        "id": "doc007",
        "title": "Docker Containers Guide",
        "content": """
        Docker packages applications and dependencies into containers that run
        consistently across environments. Containers are lightweight, isolated,
        and portable. Use Dockerfiles to define container images and docker-compose
        for multi-container applications.
        """
    },
    {
        "id": "doc008",
        "title": "Machine Learning Basics",
        "content": """
        Machine learning enables computers to learn from data without explicit
        programming. Types include supervised learning (labeled data), unsupervised
        learning (patterns in unlabeled data), and reinforcement learning (learning
        through rewards). Popular frameworks include TensorFlow and PyTorch.
        """
    },
]
```

### 3. Vector Store Implementation

```python
# vector_store.py
import chromadb
from chromadb.config import Settings
from sentence_transformers import SentenceTransformer
from typing import List, Dict, Any
import logging

logger = logging.getLogger(__name__)

class VectorStore:
    """Local vector database for semantic search."""
    
    def __init__(self, collection_name: str = "tech_docs", persist_directory: str = "./chroma_db"):
        """
        Initialize vector store with ChromaDB and sentence transformers.
        
        Args:
            collection_name: Name of the document collection
            persist_directory: Where to store the database
        """
        logger.info(f"Initializing vector store: {collection_name}")
        
        # Initialize ChromaDB client
        self.client = chromadb.PersistentClient(path=persist_directory)
        
        # Initialize embedding model (lightweight, runs locally)
        self.embedding_model = SentenceTransformer('all-MiniLM-L6-v2')
        logger.info("Loaded embedding model: all-MiniLM-L6-v2 (384 dimensions)")
        
        # Get or create collection
        self.collection = self.client.get_or_create_collection(
            name=collection_name,
            metadata={"description": "Technology documentation"}
        )
        
        logger.info(f"Collection '{collection_name}' ready with {self.collection.count()} documents")
    
    def add_documents(self, documents: List[Dict[str, str]]):
        """
        Add documents to the vector store.
        
        Args:
            documents: List of dicts with 'id', 'title', 'content' keys
        """
        logger.info(f"Adding {len(documents)} documents to vector store...")
        
        for doc in documents:
            # Create combined text for better search
            text = f"{doc['title']}\n\n{doc['content']}"
            
            # Generate embedding
            embedding = self.embedding_model.encode(text).tolist()
            
            # Add to collection
            self.collection.add(
                ids=[doc['id']],
                embeddings=[embedding],
                documents=[text],
                metadatas=[{"title": doc['title'], "doc_id": doc['id']}]
            )
        
        logger.info(f"Successfully added {len(documents)} documents")
    
    def semantic_search(self, query: str, top_k: int = 3) -> List[Dict[str, Any]]:
        """
        Perform semantic search to find relevant documents.
        
        Args:
            query: Search query
            top_k: Number of results to return
            
        Returns:
            List of matching documents with relevance scores
        """
        logger.info(f"Semantic search: '{query}' (top {top_k})")
        
        # Generate query embedding
        query_embedding = self.embedding_model.encode(query).tolist()
        
        # Search
        results = self.collection.query(
            query_embeddings=[query_embedding],
            n_results=top_k
        )
        
        # Format results
        matches = []
        for i in range(len(results['ids'][0])):
            matches.append({
                "doc_id": results['ids'][0][i],
                "title": results['metadatas'][0][i]['title'],
                "content": results['documents'][0][i],
                "distance": results['distances'][0][i] if 'distances' in results else None,
                "rank": i + 1
            })
        
        logger.info(f"Found {len(matches)} relevant documents")
        return matches
    
    def clear_collection(self):
        """Clear all documents from collection."""
        self.client.delete_collection(self.collection.name)
        logger.info("Collection cleared")
```

### 4. Search Tool for Agent

Create a function tool that wraps semantic search:

```python
def search_knowledge_base(query: str, max_results: int = 3) -> str:
    """
    Search the technical documentation knowledge base using semantic search.
    
    Args:
        query: Question or search query
        max_results: Maximum number of documents to return (default 3)
    
    Returns:
        JSON string with relevant documents
        
    Example:
        >>> search_knowledge_base("How do I write functions in Python?")
        '{"results": [{"title": "Python Functions Tutorial", "content": "...", "rank": 1}], ...}'
    """
    try:
        results = vector_store.semantic_search(query, top_k=max_results)
        
        return json.dumps({
            "query": query,
            "results": [
                {
                    "title": r['title'],
                    "content": r['content'],
                    "rank": r['rank']
                }
                for r in results
            ],
            "count": len(results)
        })
    except Exception as e:
        logger.error(f"Search error: {e}")
        return json.dumps({"error": "Search failed"})
```

### 5. Agent Integration

```python
# Create agent with knowledge base access
agent = project_client.agents.create_agent(
    model=os.environ["MODEL_DEPLOYMENT_NAME"],
    name="knowledge-agent",
    instructions="""You are a helpful technical assistant with access to 
    a knowledge base of technology documentation.
    
    When users ask questions about programming, databases, cloud computing, 
    or related topics, use the search_knowledge_base function to find relevant
    information before answering.
    
    Always cite the document titles when using retrieved information.
    If the knowledge base doesn't have relevant information, say so clearly.""",
    toolset=toolset  # Include search_knowledge_base function
)
```

### 6. Demonstrate Semantic vs. Keyword Search

Create a comparison showing how semantic search finds relevant documents that keyword search would miss:

```python
test_queries = [
    "How do I create reusable code blocks?",  # Should find Python Functions doc
    "Handling asynchronous operations in web development",  # Should find JS Async doc
    "Storing data in tables",  # Should find SQL doc
    "Managing code changes across team",  # Should find Git doc
]
```

## Expected Output

```
=== Semantic File Search Demo ===

Initializing vector store...
Loaded embedding model: all-MiniLM-L6-v2 (384 dimensions)
Adding 8 documents to vector store...
Successfully added 8 documents

=== Comparison: Keyword vs Semantic Search ===

Query: "How do I create reusable code blocks?"

[Keyword Search] (exact match only)
No documents contain "reusable code blocks"
Results: 0 matches

[Semantic Search] (meaning-based)
1. Python Functions Tutorial (distance: 0.42)
   "Python functions are reusable blocks of code defined using the 'def' keyword..."
   ✓ Found relevant document despite different wording!

=== Agent with Knowledge Base ===

User: How do I create reusable code blocks in Python?

[TOOL CALL] search_knowledge_base(query="reusable code blocks Python", max_results=3)
[SEARCH RESULTS] Found 3 relevant documents
  1. Python Functions Tutorial (rank 1)
  2. REST API Design (rank 2)
  3. Docker Containers Guide (rank 3)

Agent: Based on the Python Functions Tutorial from our knowledge base:

To create reusable code blocks in Python, you use functions defined with the 'def' keyword. 
Functions can accept parameters and return values, making your code more organized and maintainable.

Example:
```python
def greet(name):
    return f"Hello {name}"
```

This allows you to reuse the greeting logic throughout your code by calling greet("Alice"), 
greet("Bob"), etc.

=== Test Query 2 ===

User: What's the best way to handle async operations in JavaScript?

[TOOL CALL] search_knowledge_base(query="JavaScript async operations")
[SEARCH RESULTS] Found: "JavaScript Async/Await Guide"

Agent: According to our JavaScript Async/Await Guide, the best modern approach is using 
async/await syntax, which provides cleaner code for asynchronous operations. It's built 
on promises and makes async code look synchronous...

=== Semantic Search Advantages ===

✓ Finds documents by meaning, not just keywords
✓ Works with synonyms and paraphrasing  
✓ Better user experience (natural language queries)
✓ Foundation for RAG (Retrieval Augmented Generation)

Demo complete!
```

## Hints

<details>
<summary>Hint 1: Installing ChromaDB</summary>

ChromaDB is a lightweight vector database perfect for local development:

```bash
pip install chromadb sentence-transformers
```

First run will download the embedding model (~90MB) — this is normal and only happens once.

</details>

<details>
<summary>Hint 2: Understanding vector distances</summary>

ChromaDB returns distances (lower = more similar):
- **0.0-0.3**: Very similar (excellent match)
- **0.3-0.6**: Moderately similar (good match)
- **0.6-1.0**: Less similar (weak match)
- **>1.0**: Not similar

You can filter results by distance threshold:
```python
results = [r for r in results if r['distance'] < 0.5]
```

</details>

<details>
<summary>Hint 3: Choosing embedding models</summary>

Common local embedding models:

- **all-MiniLM-L6-v2**: Fast, 384 dimensions, good for most use cases (recommended)
- **all-mpnet-base-v2**: Slower, 768 dimensions, higher quality
- **paraphrase-MiniLM-L6-v2**: Optimized for paraphrase detection

For this challenge, use `all-MiniLM-L6-v2` (good balance of speed and quality).

</details>

<details>
<summary>Hint 4: Improving search quality</summary>

Tips for better semantic search:

```python
# Combine title and content for richer embeddings
text = f"Title: {doc['title']}\n\nContent: {doc['content']}"

# Add metadata for filtering
metadata = {
    "title": doc['title'],
    "category": doc.get('category', 'general'),
    "tags": doc.get('tags', [])
}

# Use where filters in queries
results = collection.query(
    query_embeddings=[embedding],
    where={"category": "programming"},  # Filter by metadata
    n_results=5
)
```

</details>

<details>
<summary>Hint 5: Testing semantic search</summary>

Create tests for search quality:

```python
test_cases = [
    {
        "query": "How to write functions?",
        "expected_doc_id": "doc001",  # Python Functions Tutorial
        "should_rank": 1  # Should be top result
    },
    {
        "query": "Database queries",
        "expected_doc_id": "doc003",  # SQL Basics
        "should_rank": [1, 2]  # Should be in top 2
    }
]

for test in test_cases:
    results = vector_store.semantic_search(test['query'])
    # Assert expected document is found
```

</details>

## Testing

### Functionality Checklist
- [ ] Vector store initializes correctly
- [ ] Documents are embedded and stored
- [ ] Semantic search returns relevant results
- [ ] Results are ranked by relevance
- [ ] Agent uses search tool appropriately
- [ ] Different phrasings find same documents

### Search Quality Tests
```python
# These should all find the Python Functions doc
queries = [
    "How do I write functions?",
    "Creating reusable code blocks",
    "Define a function in Python",
    "Python def keyword usage"
]

for q in queries:
    results = search(q)
    assert results[0]['doc_id'] == 'doc001'
```

### Performance Checks
- [ ] Search completes in < 1 second
- [ ] Embedding generation is reasonably fast
- [ ] Memory usage is acceptable
- [ ] Can handle 100+ documents

## Going Further

### Extension 1: Add More Documents
Expand your knowledge base:
```python
# Add your own documentation
# Parse markdown files, PDFs, web pages
# Build a knowledge base for your domain
```

### Extension 2: Metadata Filtering
Add category-based filtering:
```python
def search_by_category(query: str, category: str) -> str:
    results = collection.query(
        query_embeddings=[embedding],
        where={"category": category},
        n_results=5
    )
    return format_results(results)
```

### Extension 3: Hybrid Search
Combine keyword and semantic search:
```python
def hybrid_search(query: str):
    # Get semantic results
    semantic_results = vector_store.semantic_search(query)
    
    # Get keyword results  
    keyword_results = keyword_search(query)
    
    # Merge and re-rank
    combined = merge_results(semantic_results, keyword_results)
    return combined
```

### Extension 4: Document Chunking
Split large documents into chunks:
```python
def chunk_document(doc, chunk_size=500, overlap=50):
    """Split document into overlapping chunks for better retrieval."""
    chunks = []
    start = 0
    while start < len(doc):
        end = start + chunk_size
        chunk = doc[start:end]
        chunks.append(chunk)
        start = end - overlap
    return chunks
```

## Common Issues

| Issue | Cause | Solution |
|-------|-------|----------|
| Model download fails | Network or disk space | Check internet, ensure 500MB free |
| Slow first query | Model loading time | Normal on first run, cache used after |
| Poor search results | Bad embeddings | Try different model or improve document text |
| Out of memory | Too many documents | Use smaller model or process in batches |
| Results not relevant | Query too broad | Make queries more specific |

## References

- [ChromaDB Documentation](https://docs.trychroma.com/)
- [Sentence Transformers](https://www.sbert.net/)
- [Vector Embeddings Explained](https://www.pinecone.io/learn/vector-embeddings/)
- [RAG Fundamentals](https://www.anthropic.com/index/retrieval-augmented-generation-rag)
- [Semantic Search Best Practices](https://www.deepset.ai/blog/what-is-semantic-search)

## Key Takeaways

✅ **Vector embeddings capture meaning** — Text converted to numbers representing semantics  
✅ **Semantic search beats keywords** — Finds relevant content despite different wording  
✅ **ChromaDB enables local RAG** — No cloud required for vector search  
✅ **Foundation of knowledge agents** — Agents ground responses in your documents  
✅ **Quality depends on embeddings** — Better models = better search results  
✅ **This is RAG retrieval** — Next step is using retrieved docs to generate answers  

## Next Steps

**Challenge 08: Web Search** integrates external web search capabilities, teaching you how to ground agent responses in current, real-time information from the web rather than just your local knowledge base.

**Concept Preview:** Combine local knowledge (Challenge 07) with web search (Challenge 08) to create agents that can answer using both your proprietary docs AND current web information!
