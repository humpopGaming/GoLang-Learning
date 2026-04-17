# Challenge 03: Multi-Tool Agent

## Objective

Build an agent with **multiple function tools** and learn how agents intelligently select the right tool for each task, plus establish patterns for testing tool selection accuracy.

## What You'll Learn

- Managing multiple function tools in a single ToolSet
- Writing clear tool descriptions so agents choose correctly
- Testing that agents select the appropriate tool for different queries
- Handling scenarios where multiple tools could apply
- Best practices for tool naming and documentation
- Creating a systematic testing framework for tool-enabled agents

## Prerequisites

- **Challenge 02 completed** — You understand function tools and basic tool usage
- Comfort with Python functions and testing concepts
- Understanding that agents make decisions based on tool descriptions

## Background

In Challenge 02, you created an agent with one tool. Real-world agents typically have access to many tools — think of a customer service agent that can:
- Look up order status
- Check inventory
- Process refunds
- Send notifications
- Update account information

The agent's **tool selection ability** is critical. If a customer asks "What's my order status?", you want the agent to call `check_order_status()`, not `process_refund()`.

This challenge teaches you:
1. How to design tool collections that work well together
2. How to write descriptions that guide correct tool selection
3. How to test and verify tool selection reliability

## What to Build

A multi-talented assistant agent with three distinct tools:
1. **Weather Tool** — Get current weather (from Challenge 02)
2. **Calculator Tool** — Perform mathematical calculations
3. **Unit Converter Tool** — Convert between different units (temperature, distance, weight)

Your agent must:
- Automatically choose the right tool for each user request
- Handle ambiguous requests gracefully
- Combine multiple tools when needed
- Fall back to pure LLM responses when no tool is appropriate

## File to Create

```
ai-ml/agent-orchestration/solutions/challenge03/main.py
ai-ml/agent-orchestration/solutions/challenge03/test_suite.py
ai-ml/agent-orchestration/solutions/challenge03/.env
ai-ml/agent-orchestration/solutions/challenge03/requirements.txt
```

## Requirements

### 1. Tool Implementations

Create three well-documented functions:

```python
def get_weather(location: str, unit: str = "celsius") -> str:
    """
    Get current weather information for a specified location.
    
    Args:
        location: City or location name (e.g., "Seattle", "Tokyo")
        unit: Temperature unit - "celsius" or "fahrenheit"
    
    Returns:
        Current weather description including conditions and temperature
    """
    # Your implementation from Challenge 02
    pass

def calculate(expression: str) -> str:
    """
    Perform mathematical calculations safely.
    
    Args:
        expression: Mathematical expression to evaluate (e.g., "2 + 2", "15 * 7", "100 / 4")
    
    Returns:
        The calculated result as a string, or an error message if invalid
        
    Examples:
        calculate("5 + 3") → "8"
        calculate("12 * 12") → "144"
    """
    # Use eval() safely or implement a proper expression parser
    pass

def convert_units(value: float, from_unit: str, to_unit: str) -> str:
    """
    Convert values between different units of measurement.
    
    Args:
        value: Numerical value to convert
        from_unit: Source unit (e.g., "celsius", "fahrenheit", "km", "miles", "kg", "lbs")
        to_unit: Target unit (e.g., "celsius", "fahrenheit", "km", "miles", "kg", "lbs")
    
    Returns:
        Conversion result with units, or error message if conversion not supported
        
    Supported conversions:
        - Temperature: celsius ↔ fahrenheit
        - Distance: km ↔ miles
        - Weight: kg ↔ lbs
        
    Examples:
        convert_units(100, "celsius", "fahrenheit") → "212.0°F"
        convert_units(10, "km", "miles") → "6.21 miles"
    """
    pass
```

### 2. Agent with Multiple Tools

Create a ToolSet with all three functions:

```python
from azure.ai.projects.models import FunctionTool, ToolSet

# Create function tools
functions = FunctionTool([get_weather, calculate, convert_units])

# Add to toolset
toolset = ToolSet()
toolset.add(functions)

# Create agent
agent = project_client.agents.create_agent(
    model=os.environ["MODEL_DEPLOYMENT_NAME"],
    name="multi-tool-agent",
    instructions="""You are a versatile assistant with multiple capabilities:
    
    1. Weather Information: Use get_weather to provide current weather data
    2. Calculations: Use calculate for mathematical operations
    3. Unit Conversions: Use convert_units to convert between different units
    
    Always choose the most appropriate tool for the user's request. If no tool 
    is needed, respond conversationally. Be helpful and accurate.""",
    toolset=toolset
)
```

### 3. Test Suite

Create a comprehensive test suite that verifies correct tool selection:

```python
test_cases = [
    # Weather queries
    {"query": "What's the weather in Paris?", "expected_tool": "get_weather"},
    {"query": "Is it raining in Seattle?", "expected_tool": "get_weather"},
    
    # Calculator queries
    {"query": "What's 234 times 567?", "expected_tool": "calculate"},
    {"query": "Calculate 15% of 200", "expected_tool": "calculate"},
    
    # Unit conversion queries
    {"query": "Convert 100 celsius to fahrenheit", "expected_tool": "convert_units"},
    {"query": "How many miles is 50 kilometers?", "expected_tool": "convert_units"},
    
    # Ambiguous or no-tool queries
    {"query": "What's the capital of France?", "expected_tool": None},  # Pure LLM
    {"query": "Tell me a joke", "expected_tool": None},  # Pure LLM
    
    # Multi-tool queries (optional advanced challenge)
    {"query": "What's the weather in Seattle and convert 70 fahrenheit to celsius?", 
     "expected_tools": ["get_weather", "convert_units"]},
]
```

### 4. Test Runner

Implement a test runner that:
- Runs each test case
- Detects which tool(s) were called
- Reports success/failure
- Provides summary statistics

## Expected Output

```
=== Multi-Tool Agent Test Suite ===

Creating agent with 3 tools: get_weather, calculate, convert_units
Agent created: multi-tool-agent

Running 10 test cases...

[1/10] Test: "What's the weather in Paris?"
Expected tool: get_weather
Tool called: get_weather ✓
Response: The current weather in Paris is overcast with a temperature of 10°C...

[2/10] Test: "What's 234 times 567?"
Expected tool: calculate
Tool called: calculate ✓
Response: 234 times 567 equals 132,678.

[3/10] Test: "Convert 100 celsius to fahrenheit"
Expected tool: convert_units
Tool called: convert_units ✓
Response: 100°C converts to 212.0°F.

[4/10] Test: "What's the capital of France?"
Expected tool: None (pure LLM response)
Tool called: None ✓
Response: The capital of France is Paris, known for its iconic Eiffel Tower...

...

=== Test Results ===
Passed: 9/10 (90%)
Failed: 1/10 (10%)

Failed tests:
- Test #7: Expected 'convert_units', got 'calculate'
  Query: "What's 70 degrees fahrenheit in celsius?"
  Issue: Agent may have misunderstood as a calculation rather than conversion

=== Summary ===
Tool selection accuracy is good! Consider clarifying tool descriptions or
instructions for edge cases like temperature conversion questions.
```

## Hints

<details>
<summary>Hint 1: Safe calculation implementation</summary>

Use Python's `eval()` carefully or implement a safer approach:

```python
import ast
import operator

def calculate(expression: str) -> str:
    """Safely evaluate mathematical expressions."""
    try:
        # Remove whitespace
        expression = expression.strip()
        
        # Simple approach: use eval with limited namespace
        # ONLY for demonstration - not production-safe!
        allowed_chars = set("0123456789+-*/(). ")
        if not all(c in allowed_chars for c in expression):
            return "Error: Invalid characters in expression"
        
        result = eval(expression, {"__builtins__": {}}, {})
        return str(result)
    except Exception as e:
        return f"Error calculating: {str(e)}"
```

**Better approach**: Use the `ast` module or a proper expression parser library.

</details>

<details>
<summary>Hint 2: Unit conversion implementation</summary>

Create a conversion dictionary:

```python
def convert_units(value: float, from_unit: str, to_unit: str) -> str:
    """Convert between units."""
    conversions = {
        ("celsius", "fahrenheit"): lambda x: (x * 9/5) + 32,
        ("fahrenheit", "celsius"): lambda x: (x - 32) * 5/9,
        ("km", "miles"): lambda x: x * 0.621371,
        ("miles", "km"): lambda x: x / 0.621371,
        ("kg", "lbs"): lambda x: x * 2.20462,
        ("lbs", "kg"): lambda x: x / 2.20462,
    }
    
    key = (from_unit.lower(), to_unit.lower())
    if key in conversions:
        result = conversions[key](value)
        return f"{result:.2f}{to_unit}"
    else:
        return f"Error: Conversion from {from_unit} to {to_unit} not supported"
```

</details>

<details>
<summary>Hint 3: Detecting tool calls in tests</summary>

After running the agent, check what tool was used:

```python
# After creating and completing the run
run = project_client.agents.get_run(thread_id=thread.id, run_id=run.id)

tools_called = []
if hasattr(run, 'required_action') and run.required_action:
    for tool_call in run.required_action.submit_tool_outputs.tool_calls:
        tools_called.append(tool_call.function.name)

# Compare with expected
if expected_tool in tools_called:
    print("✓ Correct tool selected")
else:
    print("✗ Wrong tool selected")
```

</details>

<details>
<summary>Hint 4: Test suite structure</summary>

```python
def run_test_suite(agent, project_client, test_cases):
    results = {"passed": 0, "failed": 0, "failures": []}
    
    for i, test in enumerate(test_cases, 1):
        print(f"\n[{i}/{len(test_cases)}] Test: \"{test['query']}\"")
        
        # Create thread for this test
        thread = project_client.agents.create_thread()
        
        # Send message
        project_client.agents.create_message(
            thread_id=thread.id,
            role="user",
            content=test['query']
        )
        
        # Run agent
        run = project_client.agents.create_run(
            thread_id=thread.id,
            agent_id=agent.id
        )
        
        # Wait for completion and check tool calls
        # ... (your code to detect tool calls)
        
        # Compare with expected
        # ... (your verification logic)
        
        # Clean up
        project_client.agents.delete_thread(thread.id)
    
    return results
```

</details>

<details>
<summary>Hint 5: Improving tool selection accuracy</summary>

If your agent frequently selects the wrong tool, try:

1. **Better tool descriptions** — Make each tool's purpose crystal clear
2. **Better instructions** — Give examples of when to use each tool
3. **Better function names** — `convert_temperature()` is clearer than `convert()`
4. **Add examples to docstrings** — Show the agent what good inputs/outputs look like

Example improved instructions:
```python
instructions = """You are a versatile assistant with three specialized tools:

WEATHER: When users ask about weather conditions, forecasts, or temperature 
outside, use the get_weather function. Examples: "What's the weather?", 
"Is it raining?", "How's the weather in London?"

CALCULATIONS: When users ask you to compute, calculate, or do math, use the 
calculate function. Examples: "What's 5 times 7?", "Calculate 20% of 500"

UNIT CONVERSION: When users want to convert measurements between units, use the 
convert_units function. Examples: "Convert 10 km to miles", "100 celsius in 
fahrenheit", "How many lbs is 50 kg?"

Choose the most appropriate tool for each request. If multiple tools are needed, 
use them sequentially. If no tool applies, respond conversationally."""
```

</details>

## Testing

### Functionality Checklist
- [ ] All three tools are properly registered and callable
- [ ] Agent correctly selects `get_weather` for weather queries
- [ ] Agent correctly selects `calculate` for math queries
- [ ] Agent correctly selects `convert_units` for conversion queries
- [ ] Agent can handle queries that don't need any tool
- [ ] Test suite runs and reports results accurately

### Tool Selection Accuracy Goals
- **90%+ accuracy** on clear, unambiguous queries (Excellent)
- **75-90% accuracy** on mixed queries (Good — room for improvement)
- **Below 75%** — Review tool descriptions and instructions

### Edge Cases to Test
```python
# Ambiguous queries that could go either way
"What's 100 fahrenheit in celsius?"  # conversion or calculation?
"The temperature is 25 degrees"  # weather or conversion?

# Multi-step reasoning
"If it's 70F outside and I convert that to celsius, what do I get?"

# Negations
"I don't need weather, but what's 5 + 5?"
```

## Going Further

### Extension 1: Add More Tools
Expand your agent's capabilities:
```python
def get_time(timezone: str = "UTC") -> str:
    """Get current time in a timezone."""
    pass

def translate_text(text: str, target_language: str) -> str:
    """Translate text to another language."""
    pass

def search_web(query: str) -> str:
    """Search the web and return a summary."""
    pass
```

### Extension 2: Tool Combination Testing
Test scenarios that require using multiple tools:
```python
test_cases = [
    {
        "query": "What's the weather in Tokyo, and if it's 18 celsius, what's that in fahrenheit?",
        "expected_tools": ["get_weather", "convert_units"],
        "expected_order": ["get_weather", "convert_units"]
    }
]
```

### Extension 3: Tool Call Logging Dashboard
Create a visual summary of tool usage:
```
=== Tool Usage Statistics ===
get_weather:    ████████░░  45%  (9/20 calls)
calculate:      ██████░░░░  30%  (6/20 calls)
convert_units:  ████░░░░░░  20%  (4/20 calls)
no_tool:        ██░░░░░░░░   5%  (1/20 calls)
```

### Extension 4: Automated Regression Testing
Save test results and compare across agent instruction changes:
```python
# Run tests, save results
baseline_results = run_test_suite(agent_v1, test_cases)
save_results("baseline_v1.json", baseline_results)

# Modify instructions, run again
new_results = run_test_suite(agent_v2, test_cases)

# Compare
print_diff(baseline_results, new_results)
```

## Common Issues

| Issue | Cause | Solution |
|-------|-------|----------|
| Agent never uses tools | Instructions don't mention tools | Explicitly describe each tool in instructions |
| Wrong tool selected frequently | Tool descriptions too similar | Make tool purposes more distinct |
| Tools work but responses are poor | Agent not using tool results properly | Check that results are clear and formatted well |
| Calculate errors | Unsafe eval() usage | Implement proper expression parser |
| Conversion errors | Missing conversion pairs | Add bidirectional conversions |

## References

- [Agent Function Calling Best Practices](https://platform.openai.com/docs/guides/function-calling/best-practices)
- [Python `ast` Module (Safe Eval)](https://docs.python.org/3/library/ast.html)
- [Unit Testing in Python](https://docs.python.org/3/library/unittest.html)
- [Type Hints for Better Tool Schemas](https://docs.python.org/3/library/typing.html)

## Key Takeaways

✅ **Multiple tools require clear differentiation** — Distinct names, descriptions, and use cases  
✅ **Instructions guide tool selection** — Explicit guidance improves accuracy  
✅ **Testing is essential** — Systematic test suites catch tool selection issues early  
✅ **Tool descriptions are documentation** — They serve both the agent AND human developers  
✅ **Edge cases reveal weaknesses** — Ambiguous queries help improve your design  

## Next Steps

**Challenge 04: MCP Explorer** introduces you to **Model Context Protocol (MCP)** servers — a powerful way to give agents access to entire tool ecosystems (GitHub, file systems, databases) rather than building individual functions.

**Concept Preview:** Instead of writing `get_file()`, `list_directory()`, `search_files()` separately, MCP servers provide whole categories of tools at once!
