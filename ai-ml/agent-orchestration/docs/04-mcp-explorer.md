# Challenge 04: MCP Explorer

## Objective

Learn to connect agents to **Model Context Protocol (MCP) servers**, unlocking access to entire tool ecosystems rather than building individual functions one at a time.

## What You'll Learn

- What MCP (Model Context Protocol) is and why it's powerful for agent development
- How to connect an agent to an MCP server
- Understanding MCP approval workflows for security
- Using MCP tools vs. custom function tools
- Filtering which MCP tools your agent can access
- Real-world MCP servers: filesystem, GitHub, databases, and more

## Prerequisites

- **Challenge 03 completed** — You understand function tools and tool selection
- Basic understanding of APIs and servers
- Familiarity with Git/GitHub (helpful for one example, not required)

## Background

In Challenges 02-03, you created custom function tools like `get_weather()` and `calculate()`. This works great for a few tools, but imagine building a GitHub integration:

❌ **Manual approach**: Write 50+ functions:
```python
def get_repository(owner, repo) -> str: ...
def list_issues(owner, repo) -> str: ...
def create_issue(owner, repo, title, body) -> str: ...
def get_pull_request(owner, repo, number) -> str: ...
# ... 46 more functions ...
```

✅ **MCP approach**: Connect to a GitHub MCP server:
```python
mcp_tool = MCPTool(
    server_label="github",
    server_url="https://mcp-server-github.com",
    allowed_tools=["*"]  # Or specify which tools
)
# Instantly get access to all 50+ tools!
```

**Model Context Protocol (MCP)** is an open standard that lets MCP servers expose tools, resources, and prompts to AI agents. Many pre-built MCP servers exist:
- **Filesystem MCP** — Read/write files, list directories, search contents
- **GitHub MCP** — Manage repos, issues, PRs, code search
- **Database MCP** — Query databases, manage schemas
- **Web MCP** — Fetch web pages, parse content

This challenge teaches you how to connect to MCP servers and control agent access for security.

## What to Build

An agent that explores your filesystem using the **Filesystem MCP server**, with abilities to:
1. List directory contents
2. Read file contents
3. Search for files
4. Understand MCP approval workflows (for write operations)

You'll also create a comparison showing the difference between custom function tools and MCP tools.

## File to Create

```
ai-ml/agent-orchestration/solutions/challenge04/main.py
ai-ml/agent-orchestration/solutions/challenge04/mcp_server.py  # Local MCP server implementation
ai-ml/agent-orchestration/solutions/challenge04/.env
ai-ml/agent-orchestration/solutions/challenge04/requirements.txt
```

## Requirements

### 1. MCP Server Setup

For this challenge, we'll use a simple local MCP server. You have two options:

**Option A: Use an existing MCP server** (recommended if available)
- Filesystem MCP server
- Public MCP servers with examples

**Option B: Create a minimal local MCP server** (for learning)

```python
# mcp_server.py
from typing import Any
import json
import os

class SimpleMCPServer:
    """Minimal MCP server implementing filesystem tools."""
    
    def __init__(self, base_path: str = "."):
        self.base_path = base_path
        self.tools = {
            "list_directory": self.list_directory,
            "read_file": self.read_file,
            "search_files": self.search_files,
        }
    
    def list_directory(self, path: str = ".") -> str:
        """List contents of a directory."""
        full_path = os.path.join(self.base_path, path)
        try:
            items = os.listdir(full_path)
            return json.dumps({"items": items, "count": len(items)})
        except Exception as e:
            return json.dumps({"error": str(e)})
    
    def read_file(self, file_path: str) -> str:
        """Read contents of a file."""
        full_path = os.path.join(self.base_path, file_path)
        try:
            with open(full_path, 'r', encoding='utf-8') as f:
                content = f.read(1000)  # Limit to first 1000 chars
                return json.dumps({"content": content})
        except Exception as e:
            return json.dumps({"error": str(e)})
    
    def search_files(self, pattern: str) -> str:
        """Search for files matching a pattern."""
        # Simple implementation
        matches = []
        for root, dirs, files in os.walk(self.base_path):
            for file in files:
                if pattern.lower() in file.lower():
                    matches.append(os.path.join(root, file))
        return json.dumps({"matches": matches[:10]})  # Limit results
    
    def list_tools(self) -> list[dict[str, Any]]:
        """Return tool descriptions for MCP protocol."""
        return [
            {
                "name": "list_directory",
                "description": "List all files and directories in a path",
                "parameters": {
                    "type": "object",
                    "properties": {
                        "path": {"type": "string", "description": "Directory path"}
                    }
                }
            },
            {
                "name": "read_file",
                "description": "Read contents of a text file",
                "parameters": {
                    "type": "object",
                    "properties": {
                        "file_path": {"type": "string", "description": "Path to file"}
                    },
                    "required": ["file_path"]
                }
            },
            {
                "name": "search_files",
                "description": "Search for files by name pattern",
                "parameters": {
                    "type": "object",
                    "properties": {
                        "pattern": {"type": "string", "description": "Search pattern"}
                    },
                    "required": ["pattern"]
                }
            }
        ]
```

### 2. MCP Tool Integration

Connect your agent to the MCP server:

```python
from azure.ai.projects.models import MCPTool

# For a real MCP server
mcp_tool = MCPTool(
    server_label="filesystem",
    server_url="http://localhost:8080/mcp",  # Your MCP server URL
    require_approval="always",  # Security: always ask before executing
    allowed_tools=["list_directory", "read_file", "search_files"]  # Whitelist
)

# For local development, you might integrate differently
# (Check Agent Framework SDK docs for local MCP integration)
```

### 3. Agent with MCP Tools

Create an agent that can explore the filesystem:

```python
agent = project_client.agents.create_agent(
    model=os.environ["MODEL_DEPLOYMENT_NAME"],
    name="filesystem-explorer",
    instructions="""You are a helpful filesystem assistant. You can:
    
    - List directory contents using list_directory
    - Read file contents using read_file
    - Search for files using search_files
    
    When asked about files or directories, use the appropriate MCP tool.
    Be cautious with file operations and always confirm the path first.
    
    For security, you should always explain what you're about to do before 
    using a tool that modifies files.""",
    tools=[mcp_tool]  # or however MCP tools are added in your SDK version
)
```

### 4. Approval Workflow Demonstration

Show how MCP approval workflows work:

```python
# When agent wants to use an MCP tool with require_approval="always",
# it will return a request for approval instead of executing immediately

run = project_client.agents.create_run(
    thread_id=thread.id,
    agent_id=agent.id
)

# Check if approval is needed
if run.status == "requires_action":
    if run.required_action.type == "mcp_approval_request":
        approval_request = run.required_action.mcp_approval_request
        
        print(f"Agent wants to call: {approval_request.tool_name}")
        print(f"With arguments: {approval_request.arguments}")
        
        # In a real app, you'd ask the user for approval
        user_approves = input("Approve this action? (y/n): ")
        
        if user_approves.lower() == 'y':
            # Submit approval
            project_client.agents.submit_mcp_approval(
                thread_id=thread.id,
                run_id=run.id,
                tool_call_id=approval_request.tool_call_id,
                approve=True
            )
        else:
            # Deny the tool call
            project_client.agents.submit_mcp_approval(
                thread_id=thread.id,
                run_id=run.id,
                tool_call_id=approval_request.tool_call_id,
                approve=False
            )
```

### 5. Required Demonstrations

Your program must demonstrate:

1. **Listing directories** — Ask the agent to list files in a directory
2. **Reading files** — Ask the agent to read a specific file
3. **Searching files** — Ask the agent to find files matching a pattern
4. **Approval workflow** — Show what happens when approval is required
5. **Comparison** — Show the difference between MCP tools and custom function tools

## Expected Output

```
=== MCP Explorer Demo ===

Setting up Filesystem MCP server at: C:\Users\YourName\Documents
MCP Server running with 3 tools available:
  - list_directory
  - read_file
  - search_files

Creating filesystem-explorer agent with MCP tools...
Agent created with MCP integration enabled

=== Test 1: List Directory ===
User: What files are in the current directory?

[MCP TOOL REQUEST] list_directory(path='.')
[APPROVAL REQUIRED] Agent wants to list directory '.'
Auto-approving for demo... ✓

[MCP TOOL RESULT] {"items": ["main.py", "mcp_server.py", ".env", "README.md"], "count": 4}

Agent: The current directory contains 4 items:
- main.py
- mcp_server.py  
- .env
- README.md

=== Test 2: Read File ===
User: Can you show me what's in README.md?

[MCP TOOL REQUEST] read_file(file_path='README.md')
[APPROVAL REQUIRED] Agent wants to read file 'README.md'
Auto-approving for demo... ✓

[MCP TOOL RESULT] {"content": "# Challenge 04: MCP Explorer\n\nThis challenge..."}

Agent: Here's the content of README.md:

# Challenge 04: MCP Explorer

This challenge demonstrates MCP integration...

=== Test 3: Search Files ===
User: Find all Python files

[MCP TOOL REQUEST] search_files(pattern='py')
[APPROVAL REQUIRED] Agent wants to search for files matching 'py'
Auto-approving for demo... ✓

[MCP TOOL RESULT] {"matches": ["main.py", "mcp_server.py", "tests/test_mcp.py"]}

Agent: I found 3 Python files:
1. main.py
2. mcp_server.py
3. tests/test_mcp.py

=== Comparison: MCP vs Custom Functions ===

Custom Function Tools:
✓ Full control over implementation
✓ Simple for a few tools
✗ Manual work for each function
✗ No standardization
✗ Hard to maintain at scale

MCP Tools:
✓ Access to entire tool ecosystems instantly
✓ Standardized protocol
✓ Built-in security (approval workflows)
✓ Community servers available (GitHub, DB, etc.)
✗ Requires MCP server setup
✗ Less control over tool implementation

Demo complete!
```

## Hints

<details>
<summary>Hint 1: MCP server simplification</summary>

For learning purposes, you can simulate an MCP server with regular functions first:

```python
# Instead of a real MCP server, use function tools that mimic MCP behavior
def mcp_list_directory(path: str = ".") -> str:
    """[MCP] List directory contents."""
    items = os.listdir(path)
    return json.dumps({"items": items})

# Mark them as MCP-style in naming/documentation
functions = FunctionTool([mcp_list_directory, mcp_read_file, mcp_search_files])
```

Then in your output, explain: "This simulates MCP tools. In production, you'd connect to a real MCP server."

</details>

<details>
<summary>Hint 2: Approval workflow simulation</summary>

If your SDK doesn't support MCP approval yet, simulate it:

```python
def require_approval(tool_name: str, args: dict) -> bool:
    """Simulate MCP approval workflow."""
    print(f"\n[APPROVAL REQUIRED]")
    print(f"Tool: {tool_name}")
    print(f"Arguments: {json.dumps(args, indent=2)}")
    
    # Auto-approve for demo, but show the concept
    print("Auto-approving for demo purposes...")
    print("In production, you'd require user confirmation here.\n")
    
    return True  # or prompt user: input("Approve? (y/n): ") == 'y'
```

</details>

<details>
<summary>Hint 3: Safe file operations</summary>

When implementing file tools, add safety checks:

```python
def read_file(file_path: str) -> str:
    """Read file with safety checks."""
    # Prevent path traversal
    if ".." in file_path or file_path.startswith("/"):
        return json.dumps({"error": "Invalid path"})
    
    # Limit to certain directories
    base_dir = os.path.abspath("./safe_directory")
    full_path = os.path.abspath(os.path.join(base_dir, file_path))
    
    if not full_path.startswith(base_dir):
        return json.dumps({"error": "Access denied"})
    
    # Limit file size
    if os.path.getsize(full_path) > 100000:  # 100KB
        return json.dumps({"error": "File too large"})
    
    with open(full_path, 'r') as f:
        return json.dumps({"content": f.read()})
```

</details>

<details>
<summary>Hint 4: Testing MCP tool selection</summary>

Create test cases like Challenge 03:

```python
test_cases = [
    {"query": "What files are in the docs folder?", "expected_tool": "list_directory"},
    {"query": "Show me the contents of config.json", "expected_tool": "read_file"},
    {"query": "Find all markdown files", "expected_tool": "search_files"},
]

# Run tests and verify correct MCP tool is chosen
```

</details>

<details>
<summary>Hint 5: Real MCP server resources</summary>

If you want to use real MCP servers instead of building your own:

1. **Filesystem MCP**: https://github.com/modelcontextprotocol/servers/tree/main/src/filesystem
2. **GitHub MCP**: Available through GitHub Copilot Chat
3. **MCP Specification**: https://spec.modelcontextprotocol.io/

For this challenge, the learning goal is understanding the **concept** — using real servers is optional.

</details>

## Testing

### Functionality Checklist
- [ ] Agent can successfully list directories
- [ ] Agent can read file contents
- [ ] Agent can search for files
- [ ] Approval workflow is demonstrated (or simulated)
- [ ] Security measures are in place (path validation, etc.)
- [ ] Clear comparison between custom tools and MCP tools

### Security Checklist (Important!)
- [ ] No path traversal vulnerabilities (`../../../etc/passwd`)
- [ ] File operations restricted to safe directories
- [ ] File size limits to prevent memory issues
- [ ] Approval required for sensitive operations
- [ ] Error handling for invalid inputs

### Conceptual Understanding
- [ ] Can explain what MCP is
- [ ] Can explain benefits of MCP vs. custom tools
- [ ] Understands approval workflows and why they matter
- [ ] Can identify good use cases for MCP

## Going Further

### Extension 1: Multiple MCP Servers
Connect to multiple MCP servers simultaneously:
```python
tools = [
    MCPTool(server_label="filesystem", ...),
    MCPTool(server_label="github", ...),
    MCPTool(server_label="database", ...),
]

# Agent can now use tools from all three servers!
```

### Extension 2: Dynamic Tool Filtering
Let users choose which MCP tools to enable:
```python
allowed_tools = ["list_directory", "read_file"]  # read-only
# vs.
allowed_tools = ["*"]  # all tools including write operations
```

### Extension 3: Build a Real MCP Server
Create a proper MCP server using the MCP Python SDK:
```python
from mcp.server import Server, Tool

server = Server("my-custom-mcp")

@server.tool()
async def custom_tool(param: str) -> str:
    """Your custom tool."""
    return f"Result: {param}"

if __name__ == "__main__":
    server.run()
```

### Extension 4: MCP Tool Usage Analytics
Track which MCP tools are used most:
```python
tool_usage = defaultdict(int)

# After each run
for tool_call in tool_calls:
    tool_usage[tool_call.function.name] += 1

print("MCP Tool Usage:")
for tool, count in sorted(tool_usage.items(), key=lambda x: x[1], reverse=True):
    print(f"  {tool}: {count} calls")
```

## Common Issues

| Issue | Cause | Solution |
|-------|-------|----------|
| MCP server connection fails | Server not running or wrong URL | Check server status, verify URL in config |
| Tools not appearing | Tool filtering too restrictive | Check `allowed_tools` list |
| Approval workflow breaks | SDK version mismatch | Check SDK docs for approval API |
| Path traversal security issue | Missing validation | Add path sanitization |
| File read errors | Encoding issues | Use `encoding='utf-8'` and handle errors |

## References

- [Model Context Protocol Specification](https://spec.modelcontextprotocol.io/)
- [MCP Server Directory](https://github.com/modelcontextprotocol/servers)
- [Azure AI Agent Tools Documentation](https://learn.microsoft.com/azure/ai-studio/)
- [Building Secure File Tools](https://owasp.org/www-community/attacks/Path_Traversal)

## Key Takeaways

✅ **MCP provides instant access to tool ecosystems** — No need to build 50 functions manually  
✅ **Approval workflows enable security** — Control which tools can execute without confirmation  
✅ **MCP is a standard protocol** — Tools from different servers work the same way  
✅ **Trade-off: convenience vs. control** — MCP = less control, custom functions = more work  
✅ **Security is critical** — Always validate inputs and require approval for sensitive operations  

## Next Steps

**Challenge 05: Code Interpreter** introduces another powerful pre-built tool: the Code Interpreter, which lets your agent write and execute Python code in a sandboxed environment for calculations, data analysis, and more.

**Concept Preview:** Instead of writing a `calculate()` function that handles one expression, Code Interpreter lets the agent write and run arbitrary Python — much more powerful!
