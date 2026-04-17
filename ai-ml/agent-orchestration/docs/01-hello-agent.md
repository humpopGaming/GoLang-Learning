# Challenge 01: Hello Agent

## Objective

Set up your local Python development environment, install the Microsoft Agent Framework SDK, and create your first AI agent that can hold a basic conversation.

## What You'll Learn

- Setting up a Python virtual environment for agent development
- Installing and importing the Azure AI Projects SDK (Agent Framework)
- Creating an `AzureAIProjectClient` for local agent development
- Creating your first agent with instructions
- Starting a conversation thread and running the agent
- Understanding the basic agent lifecycle: create → thread → run → response

## Prerequisites

- **Python 3.10 or higher** installed on your system
- **Basic Python knowledge** (imports, functions, virtual environments)
- **No prior AI agent experience required** — this is your starting point!

## Background

Every AI agent system starts with the fundamentals: connecting to an agent runtime and creating a simple agent that can respond to messages. Unlike manually prompting an LLM through a chat interface, programmatic agents let you automate conversations, integrate tools, and orchestrate complex workflows.

In this challenge, you'll use the **Microsoft Agent Framework SDK** to create an agent locally. Think of the agent as a configured LLM instance with specific instructions (system prompt) that defines its behavior.

## What to Build

A Python program that:
1. Sets up a connection to a local or cloud agent project
2. Creates an agent with custom instructions (personality/role)
3. Creates a conversation thread
4. Sends a message and gets the agent's response
5. Prints the conversation to the console

## File to Create

```
ai-ml/agent-orchestration/solutions/challenge01/main.py
ai-ml/agent-orchestration/solutions/challenge01/.env
ai-ml/agent-orchestration/solutions/challenge01/requirements.txt
```

## Requirements

### 1. Environment Setup (`requirements.txt`)
Create a requirements file with:
```
azure-ai-projects
azure-identity
python-dotenv
```

### 2. Configuration (`.env`)
Create an environment file with these variables:
```
# For local development, you can use mock values or connect to Azure AI Foundry
PROJECT_CONNECTION_STRING=<your-connection-string-or-mock-value>
MODEL_DEPLOYMENT_NAME=gpt-4o
```

**Note:** For truly local development without cloud resources, we'll use environment variables that the SDK can work with in offline mode, or you can connect to Azure AI Foundry if you have access.

### 3. Agent Implementation (`main.py`)

Your program must:

1. **Load environment variables** using `python-dotenv`
2. **Create a project client** using `AIProjectClient`
3. **Create an agent** with:
   - Model: from environment variable `MODEL_DEPLOYMENT_NAME`
   - Name: `"hello-agent"`
   - Instructions: `"You are a friendly AI assistant helping someone learn about AI agents. Keep responses concise and encouraging."`
4. **Create a thread** for the conversation
5. **Send a message** to the thread: `"Hello! What can you tell me about AI agents?"`
6. **Run the agent** on the thread
7. **Retrieve and print** the agent's response
8. **Clean up** by deleting the thread and agent (optional but good practice)

## Expected Output

When you run `python main.py`, you should see output similar to:

```
Creating agent...
Agent created: hello-agent

Starting conversation thread...
Thread created: thread_abc123

User: Hello! What can you tell me about AI agents?

Running agent...
Agent: AI agents are programs that can understand natural language, make decisions, and take actions to help accomplish tasks. Think of them as intelligent assistants that can:

- Have conversations and answer questions
- Use tools and functions to get real-time information
- Work together with other agents to solve complex problems
- Learn from context and provide personalized responses

You're taking your first step into agent development - exciting! What would you like to explore?

Conversation complete!
```

**Note:** The exact response will vary based on the model, but it should be conversational and relevant to AI agents.

## Hints

<details>
<summary>Hint 1: Setting up the project client</summary>

The `AIProjectClient` is your gateway to the agent runtime. Import it from `azure.ai.projects`:

```python
from azure.ai.projects import AIProjectClient
from azure.identity import DefaultAzureCredential
import os

project_client = AIProjectClient.from_connection_string(
    credential=DefaultAzureCredential(),
    conn_str=os.environ["PROJECT_CONNECTION_STRING"]
)
```

For local development without Azure, you might need to use a different initialization pattern or mock implementation.

</details>

<details>
<summary>Hint 2: Creating the agent</summary>

Use the `agents` property of your project client:

```python
agent = project_client.agents.create_agent(
    model=os.environ["MODEL_DEPLOYMENT_NAME"],
    name="hello-agent",
    instructions="Your system prompt here..."
)
print(f"Agent created: {agent.id}")
```

</details>

<details>
<summary>Hint 3: Running a conversation</summary>

The pattern is: create thread → add message → create run → wait for completion → get messages:

```python
# Create thread
thread = project_client.agents.create_thread()

# Add user message
message = project_client.agents.create_message(
    thread_id=thread.id,
    role="user",
    content="Your message here"
)

# Run the agent
run = project_client.agents.create_run(
    thread_id=thread.id,
    agent_id=agent.id
)

# Wait for completion and get response
# (You'll need to poll or use the SDK's wait methods)
```

</details>

<details>
<summary>Hint 4: Getting the response</summary>

After the run completes, retrieve messages from the thread:

```python
messages = project_client.agents.list_messages(thread_id=thread.id)

# Messages are returned in reverse chronological order
for message in messages:
    print(f"{message.role}: {message.content[0].text.value}")
```

</details>

<details>
<summary>Hint 5: Complete skeleton</summary>

```python
import os
from dotenv import load_dotenv
from azure.ai.projects import AIProjectClient
from azure.identity import DefaultAzureCredential

# Load environment variables
load_dotenv()

# Create project client
project_client = AIProjectClient.from_connection_string(
    credential=DefaultAzureCredential(),
    conn_str=os.environ["PROJECT_CONNECTION_STRING"]
)

# TODO: Create agent with create_agent()
# TODO: Create thread with create_thread()
# TODO: Add message with create_message()
# TODO: Run agent with create_run()
# TODO: Wait for completion
# TODO: Get and print messages
# TODO: Clean up (delete thread, agent)
```

</details>

## Testing

To verify your solution works:

1. **Run the program**: `python main.py`
2. **Check output**: You should see the conversation printed clearly
3. **Verify agent response**: The response should be relevant to AI agents
4. **No errors**: Program should complete without exceptions
5. **Clean execution**: Resources cleaned up (if you added cleanup code)

### Troubleshooting

| Issue | Solution |
|-------|----------|
| `ModuleNotFoundError` | Activate venv and install requirements: `pip install -r requirements.txt` |
| `Connection refused` | Check PROJECT_CONNECTION_STRING in .env |
| `Model not found` | Verify MODEL_DEPLOYMENT_NAME exists in your project |
| `Authentication failed` | Run `az login` or check DefaultAzureCredential setup |

## Going Further

Once you have the basic version working, try these extensions:

### Extension 1: Multi-Turn Conversation
Modify your program to have a multi-turn conversation:
```python
questions = [
    "Hello! What can you tell me about AI agents?",
    "How do agents differ from regular chatbots?",
    "What are some real-world applications?"
]

for question in questions:
    # Send message, run agent, print response
```

### Extension 2: Different Personalities
Create multiple agents with different instructions and see how they respond differently to the same question:
- A formal, technical assistant
- A casual, friendly tutor
- A creative, enthusiastic teacher

### Extension 3: Streaming Responses
Instead of waiting for the complete response, stream it token-by-token (check SDK docs for streaming support).

## References

- [Azure AI Projects Python SDK Documentation](https://learn.microsoft.com/python/api/overview/azure/ai-projects-readme)
- [Microsoft Agent Framework GitHub](https://github.com/microsoft/agent-framework)
- [Azure AI Foundry Documentation](https://learn.microsoft.com/azure/ai-studio/)
- [Python Virtual Environments Guide](https://docs.python.org/3/tutorial/venv.html)
- [python-dotenv Documentation](https://pypi.org/project/python-dotenv/)

## Next Steps

Once you complete this challenge, you'll be ready for **Challenge 02: Weather Tool**, where you'll learn how agents can automatically use functions (tools) to get real-time information instead of just generating text responses.

**Concept Preview:** Right now, your agent only knows what was in its training data. In the next challenge, you'll give it the ability to call a weather function — this is where agents become truly powerful!
