# Challenge 05: Code Interpreter

## Objective

Learn to use the built-in **Code Interpreter** tool that lets your agent write and execute Python code dynamically in a sandboxed environment for calculations, data processing, and analysis.

## What You'll Learn

- What the Code Interpreter tool is and when to use it
- How to enable Code Interpreter in your agent
- The difference between function tools and dynamic code execution
- Handling file inputs and outputs with Code Interpreter
- Use cases: calculations, data analysis, plotting, file processing
- Security implications of code execution tools

## Prerequisites

- **Challenge 03 completed** — You understand function tools
- **Challenge 04 completed** — You understand built-in tools vs. custom tools
- Basic Python programming knowledge
- Understanding of sandboxed environments (helpful but not required)

## Background

In Challenge 02, you created a `calculate()` function that could evaluate simple math expressions like `"5 + 3"`. But what if the user asks:

- "Calculate the compound interest on $10,000 at 5% annual rate for 30 years"
- "Generate a bar chart showing sales data"
- "Read this CSV file and calculate the average of column B"
- "Write a function to find prime numbers up to 1000"

Your simple `calculate()` function can't handle these. You'd need to write dozens of specialized functions.

**Code Interpreter** solves this by letting the agent **write and execute Python code** dynamically. The agent:
1. Receives a request requiring computation
2. Writes Python code to solve it
3. Executes the code in a sandboxed environment
4. Returns the result to the user

**Sandboxed environment** means the code runs isolated from your system:
- ✓ Can't access your files (unless explicitly uploaded)
- ✓ Can't make network requests
- ✓ Can't modify your system
- ✓ Has access to common Python libraries (numpy, pandas, matplotlib, etc.)

This is incredibly powerful for data analysis, scientific computing, and complex calculations.

## What to Build

An agent with Code Interpreter that can:
1. Perform complex mathematical calculations
2. Analyze data (process lists, calculate statistics)
3. Generate visualizations (charts/graphs)
4. Process files (CSVs, text files)
5. Write and run custom Python functions

## File to Create

```
ai-ml/agent-orchestration/solutions/challenge05/main.py
ai-ml/agent-orchestration/solutions/challenge05/sample_data.csv
ai-ml/agent-orchestration/solutions/challenge05/.env
ai-ml/agent-orchestration/solutions/challenge05/requirements.txt
```

## Requirements

### 1. Sample Data File

Create a CSV file for testing file processing:

```csv
# sample_data.csv
Month,Sales,Expenses
January,10000,7500
February,12000,8000
March,15000,9500
April,13000,8200
May,16000,10000
June,18000,11000
```

### 2. Enable Code Interpreter

Create an agent with Code Interpreter enabled:

```python
from azure.ai.projects.models import CodeInterpreterTool

agent = project_client.agents.create_agent(
    model=os.environ["MODEL_DEPLOYMENT_NAME"],
    name="code-interpreter-agent",
    instructions="""You are a helpful data analyst and mathematician. 
    
    You have access to Python code execution through Code Interpreter. Use it to:
    - Perform complex calculations and mathematical operations
    - Analyze data and calculate statistics
    - Generate charts and visualizations
    - Process CSV files and data
    - Write and test Python functions
    
    Always show your Python code to the user so they can learn.
    Explain your approach before executing code.
    """,
    tools=[CodeInterpreterTool()]
)
```

### 3. Test Cases

Demonstrate Code Interpreter with these scenarios:

**Test 1: Complex Calculation**
```python
query = "Calculate compound interest: principal $10,000, annual rate 5%, compounded monthly for 30 years. Show the formula and final amount."
```

**Test 2: Data Analysis**
```python
query = "Generate a list of the first 20 Fibonacci numbers and calculate their sum."
```

**Test 3: Statistics**
```python
query = "Create a list of 100 random numbers between 1 and 100, then calculate mean, median, mode, and standard deviation."
```

**Test 4: File Processing** (if supported in your SDK version)
```python
# Upload sample_data.csv to the agent
query = "Analyze the sales data CSV file. Calculate total sales, average monthly sales, and profit margin (sales - expenses)."
```

**Test 5: Visualization**
```python
query = "Create a bar chart comparing sales vs expenses by month from the CSV data. Save it as a PNG file."
```

### 4. Code Visibility

Your program should show:
- The user's query
- The Python code the agent wrote (for learning)
- The execution result
- Any files generated (plots, outputs)

## Expected Output

```
=== Code Interpreter Agent Demo ===

Creating agent with Code Interpreter enabled...
Agent created: code-interpreter-agent

=== Test 1: Compound Interest ===
User: Calculate compound interest: principal $10,000, annual rate 5%, compounded monthly for 30 years. Show the formula and final amount.

Agent: I'll calculate the compound interest using the formula:
A = P(1 + r/n)^(nt)

Where:
- P = principal ($10,000)
- r = annual rate (0.05)
- n = compounds per year (12)
- t = time in years (30)

[EXECUTING PYTHON CODE]
```python
# Compound interest calculation
P = 10000  # Principal
r = 0.05   # Annual rate
n = 12     # Compounds per year
t = 30     # Time in years

A = P * (1 + r/n) ** (n*t)
total_interest = A - P

print(f"Final amount: ${A:,.2f}")
print(f"Total interest earned: ${total_interest:,.2f}")
```
[CODE OUTPUT]
Final amount: $44,677.44
Total interest earned: $34,677.44

Agent: After 30 years, your $10,000 investment will grow to $44,677.44, earning $34,677.44 in interest!

=== Test 2: Fibonacci Numbers ===
User: Generate a list of the first 20 Fibonacci numbers and calculate their sum.

[EXECUTING PYTHON CODE]
```python
def fibonacci(n):
    fib_list = [0, 1]
    for i in range(2, n):
        fib_list.append(fib_list[-1] + fib_list[-2])
    return fib_list

fib_20 = fibonacci(20)
fib_sum = sum(fib_20)

print(f"First 20 Fibonacci numbers: {fib_20}")
print(f"Sum: {fib_sum}")
```
[CODE OUTPUT]
First 20 Fibonacci numbers: [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181]
Sum: 10945

Agent: The sum of the first 20 Fibonacci numbers is 10,945.

=== Test 3: File Analysis ===
User: Analyze the sales data. Calculate total sales and profit margin.

[FILE UPLOADED] sample_data.csv

[EXECUTING PYTHON CODE]
```python
import pandas as pd

# Read the CSV file
df = pd.read_csv('sample_data.csv')

# Calculate totals
total_sales = df['Sales'].sum()
total_expenses = df['Expenses'].sum()
profit = total_sales - total_expenses
profit_margin = (profit / total_sales) * 100

print(f"Total Sales: ${total_sales:,}")
print(f"Total Expenses: ${total_expenses:,}")
print(f"Profit: ${profit:,}")
print(f"Profit Margin: {profit_margin:.2f}%")
print(f"\nAverage Monthly Sales: ${df['Sales'].mean():,.2f}")
```
[CODE OUTPUT]
Total Sales: $84,000
Total Expenses: $54,200
Profit: $29,800
Profit Margin: 35.48%

Average Monthly Sales: $14,000.00

Agent: Your business generated $84,000 in sales with a healthy profit margin of 35.48%!

=== Test 4: Visualization ===
User: Create a bar chart showing sales vs expenses.

[EXECUTING PYTHON CODE]
```python
import matplotlib.pyplot as plt

months = df['Month']
sales = df['Sales']
expenses = df['Expenses']

fig, ax = plt.subplots(figsize=(10, 6))
x = range(len(months))
width = 0.35

ax.bar([i - width/2 for i in x], sales, width, label='Sales', color='green')
ax.bar([i + width/2 for i in x], expenses, width, label='Expenses', color='red')

ax.set_xlabel('Month')
ax.set_ylabel('Amount ($)')
ax.set_title('Sales vs Expenses by Month')
ax.set_xticks(x)
ax.set_xticklabels(months, rotation=45)
ax.legend()
ax.grid(axis='y', alpha=0.3)

plt.tight_layout()
plt.savefig('sales_vs_expenses.png', dpi=300)
print("Chart saved as sales_vs_expenses.png")
```
[CODE OUTPUT]
Chart saved as sales_vs_expenses.png

[FILE GENERATED] sales_vs_expenses.png

Agent: I've created a bar chart comparing sales and expenses. You can see that sales consistently exceed expenses, with the gap growing over time!

All tests complete!
```

## Hints

<details>
<summary>Hint 1: Enabling Code Interpreter</summary>

Import and add the tool:

```python
from azure.ai.projects.models import CodeInterpreterTool

tools = [CodeInterpreterTool()]

agent = project_client.agents.create_agent(
    model=os.environ["MODEL_DEPLOYMENT_NAME"],
    name="code-agent",
    instructions="...",
    tools=tools  # Add Code Interpreter
)
```

The agent will automatically use it when code execution is helpful.

</details>

<details>
<summary>Hint 2: Extracting code from responses</summary>

To show the Python code the agent wrote, inspect the run details:

```python
# After run completes
run = project_client.agents.get_run(thread_id=thread.id, run_id=run.id)

# Check for tool calls
if hasattr(run, 'required_action'):
    for tool_call in run.required_action.submit_tool_outputs.tool_calls:
        if tool_call.type == "code_interpreter":
            print("[EXECUTING PYTHON CODE]")
            print(f"```python\n{tool_call.code_interpreter.input}\n```")
            print(f"[OUTPUT]\n{tool_call.code_interpreter.outputs}")
```

</details>

<details>
<summary>Hint 3: File uploads (if supported)</summary>

Some SDK versions support uploading files for Code Interpreter:

```python
# Upload file
with open("sample_data.csv", "rb") as f:
    file = project_client.agents.upload_file(
        file=f,
        purpose="assistants"
    )

# Create agent with file access
agent = project_client.agents.create_agent(
    model=os.environ["MODEL_DEPLOYMENT_NAME"],
    name="data-agent",
    tools=[CodeInterpreterTool()],
    file_ids=[file.id]  # Give agent access to the file
)
```

</details>

<details>
<summary>Hint 4: When to use Code Interpreter vs. function tools</summary>

**Use Code Interpreter when:**
- Complex calculations requiring multiple steps
- Data analysis and statistics
- Generating visualizations
- Processing structured data (CSV, JSON)
- User needs vary widely (can't predict all cases)

**Use Function Tools when:**
- Simple, predictable operations
- Need to integrate with external APIs/systems
- Security-critical operations (database writes, API calls)
- Performance is critical (function calls are faster)
- Need precise control over execution

</details>

<details>
<summary>Hint 5: Testing systematic progression</summary>

Structure your tests to build complexity:

```python
test_cases = [
    # Simple calculation
    {"query": "What's 123 * 456?", "complexity": "basic"},
    
    # Multi-step calculation
    {"query": "Calculate compound interest...", "complexity": "intermediate"},
    
    # Data structure
    {"query": "Generate Fibonacci numbers...", "complexity": "intermediate"},
    
    # File processing
    {"query": "Analyze the CSV file...", "complexity": "advanced"},
    
    # Visualization
    {"query": "Create a bar chart...", "complexity": "advanced"},
]

for test in test_cases:
    run_test(test)
```

</details>

## Testing

### Functionality Checklist
- [ ] Agent uses Code Interpreter for calculations
- [ ] Complex multi-step calculations work correctly
- [ ] Data analysis (statistics, aggregations) works
- [ ] File processing works (if supported)
- [ ] Visualizations can be generated (if supported)
- [ ] Code is shown in output for learning

### Code Quality in Agent Responses
- [ ] Agent's Python code is syntactically correct
- [ ] Code includes appropriate comments
- [ ] Code follows Python best practices
- [ ] Agent explains its approach before coding

### Safety Checks
- [ ] Code runs in isolated environment
- [ ] No system access attempts succeed
- [ ] File operations are restricted to uploaded files
- [ ] Error handling works for invalid code

## Going Further

### Extension 1: Data Science Workflow
Create a complete data science workflow:
```python
query = """
1. Load the sales data CSV
2. Clean the data (check for missing values)
3. Calculate descriptive statistics
4. Perform trend analysis (is sales growing?)
5. Create visualizations:
   - Line chart of sales over time
   - Pie chart of expense categories
   - Correlation heatmap
6. Provide recommendations based on data
"""
```

### Extension 2: Interactive Calculations
Let users provide parameters:
```python
principal = input("Enter principal amount: ")
rate = input("Enter annual interest rate (%): ")
years = input("Enter number of years: ")

query = f"""Calculate compound interest with:
- Principal: ${principal}
- Rate: {rate}%
- Time: {years} years
- Compounded monthly
"""
```

### Extension 3: Code Explanation Mode
Add a mode where the agent explains every line:
```python
instructions = """...
When writing code, add detailed comments explaining each step.
After execution, walk through the code line-by-line to teach the user.
"""
```

### Extension 4: Comparison: Code Interpreter vs Function
Create both approaches for the same task and compare:
```python
# Approach 1: Function tool
def calculate_compound_interest(P, r, n, t):
    return P * (1 + r/n) ** (n*t)

# Approach 2: Code Interpreter
# Agent writes code dynamically

# Compare: speed, flexibility, accuracy
```

## Common Issues

| Issue | Cause | Solution |
|-------|-------|----------|
| Code never executes | Agent doesn't recognize need | Improve instructions, be explicit in query |
| Code execution fails | Syntax error in generated code | Check agent's code quality, provide examples |
| Libraries not available | Limited sandbox environment | Check which libraries are available in sandbox |
| Files not accessible | File not uploaded properly | Verify file upload process |
| Visualization doesn't appear | No file output handling | Check how generated files are retrieved |

## References

- [Azure AI Code Interpreter Documentation](https://learn.microsoft.com/azure/ai-studio/)
- [Python Code Execution Best Practices](https://docs.python.org/3/library/functions.html#eval)
- [Matplotlib Documentation](https://matplotlib.org/stable/contents.html)
- [Pandas Documentation](https://pandas.pydata.org/docs/)
- [Sandboxed Python Environments](https://docs.python.org/3/library/sandbox.html)

## Key Takeaways

✅ **Code Interpreter enables dynamic problem-solving** — Agent writes code for each specific request  
✅ **More flexible than function tools** — Handles unpredictable, complex tasks  
✅ **Sandboxing provides security** — Code runs isolated from your system  
✅ **Best for data analysis and computation** — Statistics, visualizations, file processing  
✅ **Trade-off: flexibility vs. predictability** — Less control over exact execution  

## Next Steps

**Challenge 06: Custom Tool Builder** brings it all together — you'll design and build a complete custom tool with validation, error handling, comprehensive documentation, and testing, following production-grade best practices.

**Concept Preview:** Now that you've seen function tools, MCP tools, and Code Interpreter, you'll learn how to build industrial-strength custom tools that are reliable, secure, and maintainable!
