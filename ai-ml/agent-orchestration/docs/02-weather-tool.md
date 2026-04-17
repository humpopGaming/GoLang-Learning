# Challenge 02: Weather Tool

## Objective

Learn how to equip your agent with **function tools** that it can automatically invoke to get real-time information, transforming a pure text generator into an action-taking assistant.

## What You'll Learn

- What function tools are and why they're essential for practical agents
- How to define a Python function that an agent can call
- Creating function schemas that describe your tool to the agent
- Using `FunctionTool` and `ToolSet` from the Agent Framework
- Understanding the tool invocation lifecycle: user request → agent decides → calls function → uses result
- How to test and verify that tools are being called correctly

## Prerequisites

- **Challenge 01 completed** — You can create agents and run conversations
- Understanding of Python functions and type hints
- Familiarity with JSON schema concepts (helpful but not required)

## Background

Pure LLM agents (like Challenge 01) can only generate text based on their training data. They can't:
- Get current weather, stock prices, or news
- Look up information in your database
- Calculate complex math beyond token prediction
- Interact with external systems

**Function tools** solve this by letting the agent call Python functions you provide. Here's how it works:

1. You define a Python function (e.g., `get_weather(location: str)`)
2. The agent gets a description of what the function does
3. When a user asks "What's the weather in Seattle?", the agent:
   - Realizes it needs weather data
   - Calls your function with `location="Seattle"`
   - Receives the result
   - Generates a natural language response using that data

This is the foundation of **all** practical AI agents.

## What to Build

A weather assistant agent that:
1. Has access to a `get_weather()` function tool
2. Can answer questions about current weather in different cities
3. Automatically calls the function when needed (no manual tool invocation)
4. Incorporates the function results into natural responses

## File to Create

```
ai-ml/agent-orchestration/solutions/challenge02/main.py
ai-ml/agent-orchestration/solutions/challenge02/.env
ai-ml/agent-orchestration/solutions/challenge02/requirements.txt
```

## Requirements

### 1. Weather Function (Mock Implementation)

Create a `get_weather(location: str, unit: str = "celsius")` function that returns mock weather data:

```python
def get_weather(location: str, unit: str = "celsius") -> str:
    """
    Get the current weather for a location.
    
    Args:
        location: City name or location (e.g., "Seattle", "London")
        unit: Temperature unit - "celsius" or "fahrenheit"
    
    Returns:
        A string describing the current weather
    """
    # Mock implementation - returns fake but realistic data
    # In a real app, this would call a weather API
    
    weather_data = {
        "Seattle": "Cloudy with light rain, 12°C",
        "London": "Overcast, 10°C",
        "Tokyo": "Clear and sunny, 18°C",
        "Sydney": "Partly cloudy, 22°C",
        "New York": "Sunny, 15°C",
    }
    
    result = weather_data.get(location, f"Weather data not available for {location}")
    
    # Convert temperature if fahrenheit requested
    if unit == "fahrenheit" and "°C" in result:
        # Simple mock conversion for demo
        result = result.replace("°C", "°F").replace("12", "54").replace("10", "50")
        # etc...
    
    return f"Current weather in {location}: {result}"
```

### 2. Function Tool Registration

Create a `FunctionTool` from your function and add it to a `ToolSet`:

```python
from azure.ai.projects.models import FunctionTool, ToolSet

# Create function tool
functions = FunctionTool([get_weather])

# Create toolset and add the function tool
toolset = ToolSet()
toolset.add(functions)
```

### 3. Agent with Tools

Create an agent that has access to the toolset:

```python
agent = project_client.agents.create_agent(
    model=os.environ["MODEL_DEPLOYMENT_NAME"],
    name="weather-agent",
    instructions="""You are a helpful weather assistant. When users ask about 
    weather, use the get_weather function to get current information. 
    Always specify the location clearly.""",
    toolset=toolset
)
```

### 4. Conversation Flow

Test your agent with weather-related questions:
- "What's the weather like in Seattle?"
- "How's the weather in Tokyo right now?"
- "Is it raining in London?"

The agent should:
1. Recognize these require weather data
2. Call `get_weather()` with the correct location
3. Use the function's return value to answer naturally

### 5. Required Output

Your program must:
1. Print when the function is being called (for learning purposes)
2. Show the user's question
3. Show the agent's final response incorporating the weather data
4. Demonstrate at least 2 different weather queries

## Expected Output

```
Creating weather agent with function tools...
Agent created: weather-agent

=== Weather Query 1 ===
User: What's the weather like in Seattle?

[TOOL CALL] get_weather(location='Seattle', unit='celsius')
[TOOL RESULT] Current weather in Seattle: Cloudy with light rain, 12°C

Agent: The current weather in Seattle is cloudy with light rain, and the temperature is 12°C. You might want to bring an umbrella if you're heading out!

=== Weather Query 2 ===
User: How's Tokyo doing today?

[TOOL CALL] get_weather(location='Tokyo', unit='celsius')
[TOOL RESULT] Current weather in Tokyo: Clear and sunny, 18°C

Agent: Tokyo is enjoying beautiful weather today! It's clear and sunny with a pleasant temperature of 18°C. Perfect day to be outside!

Conversation complete!
```

## Hints

<details>
<summary>Hint 1: Function tool creation</summary>

The `FunctionTool` automatically generates schemas from your Python function's docstring and type hints:

```python
from azure.ai.projects.models import FunctionTool

def get_weather(location: str, unit: str = "celsius") -> str:
    """Get current weather for a location."""
    # Implementation here
    pass

# This automatically creates the schema from the function signature
functions = FunctionTool([get_weather])
```

The type hints (`str`) and default values (`= "celsius"`) tell the agent what parameters are required vs. optional.

</details>

<details>
<summary>Hint 2: Adding tools to the agent</summary>

Create a `ToolSet` and add your function tool to it:

```python
from azure.ai.projects.models import ToolSet

toolset = ToolSet()
toolset.add(functions)

agent = project_client.agents.create_agent(
    model=os.environ["MODEL_DEPLOYMENT_NAME"],
    name="weather-agent",
    instructions="Your instructions here...",
    toolset=toolset  # ← Add this parameter
)
```

</details>

<details>
<summary>Hint 3: Detecting tool calls</summary>

When you retrieve the run after completion, check if tools were called:

```python
run = project_client.agents.create_run(
    thread_id=thread.id,
    agent_id=agent.id
)

# Wait for completion
while run.status in ["queued", "in_progress", "requires_action"]:
    time.sleep(1)
    run = project_client.agents.get_run(thread_id=thread.id, run_id=run.id)

# Check for tool calls
if run.status == "completed":
    # Get messages to see the result
    messages = project_client.agents.list_messages(thread_id=thread.id)
```

For debugging, you might want to inspect `run.required_action` to see tool calls that were made.

</details>

<details>
<summary>Hint 4: Good instructions matter</summary>

Your agent's instructions should mention the tool availability:

```python
instructions = """You are a helpful weather assistant. 

When users ask about weather conditions, use the get_weather function to 
retrieve current weather data. Always be specific about the location.

Provide friendly, conversational responses that incorporate the weather 
information naturally."""
```

</details>

<details>
<summary>Hint 5: Testing multiple locations</summary>

Create a loop to test different queries:

```python
test_questions = [
    "What's the weather like in Seattle?",
    "How's Tokyo doing today?",
    "Is it raining in London?"
]

for i, question in enumerate(test_questions, 1):
    print(f"\n=== Weather Query {i} ===")
    print(f"User: {question}")
    
    # Create message
    project_client.agents.create_message(
        thread_id=thread.id,
        role="user",
        content=question
    )
    
    # Run agent
    run = project_client.agents.create_run(
        thread_id=thread.id,
        agent_id=agent.id
    )
    
    # Wait and get response
    # ... (your code to wait and print the response)
```

</details>

## Testing

Verify your solution:

### Functionality Tests
- [ ] Agent correctly identifies when weather data is needed
- [ ] Function is called with the right location parameter
- [ ] Agent incorporates function results into natural responses
- [ ] Works for multiple different locations
- [ ] Handles unknown locations gracefully

### Code Quality
- [ ] Function has proper type hints and docstring
- [ ] Tool calls are logged for debugging
- [ ] Clean output that's easy to follow
- [ ] No hardcoded responses (agent actually uses function data)

### Edge Cases to Try
```python
# What happens with these?
"What's the weather?"  # Missing location - agent should ask
"Paris weather please"  # Different phrasing
"Is it sunny in Miami?"  # Location not in your mock data
```

## Going Further

### Extension 1: Multiple Weather Functions
Add more weather-related functions:
```python
def get_forecast(location: str, days: int = 3) -> str:
    """Get weather forecast for the next N days."""
    pass

def get_weather_alert(location: str) -> str:
    """Get any active weather warnings or alerts."""
    pass
```

### Extension 2: Real Weather API
Replace the mock implementation with a real weather API:
- [OpenWeatherMap API](https://openweathermap.org/api) (free tier available)
- [WeatherAPI.com](https://www.weatherapi.com/) (free tier available)

```python
import requests

def get_weather(location: str, unit: str = "celsius") -> str:
    api_key = os.environ["WEATHER_API_KEY"]
    url = f"https://api.openweathermap.org/data/2.5/weather?q={location}&appid={api_key}"
    response = requests.get(url)
    data = response.json()
    # Parse and return weather data
```

### Extension 3: Error Handling
Add robust error handling to your function:
```python
def get_weather(location: str, unit: str = "celsius") -> str:
    try:
        # Your implementation
        if not location:
            return "Error: Please specify a location"
        # ... rest of code
    except Exception as e:
        return f"Error fetching weather: {str(e)}"
```

### Extension 4: Logging Tool Calls
Create a decorator to log all function tool calls:
```python
def log_tool_call(func):
    def wrapper(*args, **kwargs):
        print(f"[TOOL CALL] {func.__name__}({args}, {kwargs})")
        result = func(*args, **kwargs)
        print(f"[TOOL RESULT] {result}")
        return result
    return wrapper

@log_tool_call
def get_weather(location: str, unit: str = "celsius") -> str:
    # Implementation
```

## Common Issues

| Issue | Cause | Solution |
|-------|-------|----------|
| Tool never called | Instructions don't mention it | Update agent instructions to reference weather capability |
| Wrong parameters | Poor function docstring | Add detailed docstring with parameter descriptions |
| Agent hallucinates data | Not using function result | Check that agent actually incorporates the function's return value |
| `ToolSet` not found | Wrong import | Import from `azure.ai.projects.models` |

## References

- [Azure AI Agent Tools Documentation](https://learn.microsoft.com/azure/ai-studio/how-to/develop/agents)
- [Function Calling Patterns](https://platform.openai.com/docs/guides/function-calling)
- [Python Type Hints](https://docs.python.org/3/library/typing.html)
- [JSON Schema Basics](https://json-schema.org/learn/getting-started-step-by-step)

## Key Takeaways

By completing this challenge, you've learned:

✅ **Function tools** transform LLMs from text generators into action-takers  
✅ **Type hints and docstrings** automatically create tool schemas  
✅ **Agents decide when to use tools** based on user requests and instructions  
✅ **Tool results are incorporated** into natural language responses  
✅ **This pattern scales** — you can add dozens of functions the same way  

## Next Steps

**Challenge 03: Multi-Tool Agent** builds on this by giving your agent access to multiple different tools (weather, calculator, unit converter) and teaching you how to test that the agent selects the *right* tool for each task.

**Concept Preview:** With multiple tools, agents become even more powerful — but also need better instructions and testing to ensure they choose appropriately!
