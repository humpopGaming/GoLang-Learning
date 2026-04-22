# Quick Start Guide: Holiday Finder Platform

**Feature**: 001-ai-holiday-booking  
**Date**: 2026-04-20  
**Target**: Developers setting up the project for the first time

## Overview

This guide will help you set up the Holiday Finder platform locally, including the ASP.NET Core Web API with HATEOAS, Blazor WebAssembly client, and Azure Cosmos DB.

## Prerequisites

### Required Software

- **.NET 8 SDK** or later - [Download](https://dotnet.microsoft.com/download/dotnet/8.0)
- **Visual Studio 2022** (17.8+) or **VS Code** with C# Dev Kit
- **Azure Cosmos DB Emulator** - [Download](https://learn.microsoft.com/azure/cosmos-db/emulator) or Azure subscription
- **Git** for version control
- **Node.js 18+** (optional, for additional tooling)

### Recommended Tools

- **Postman** or **Thunder Client** for API testing
- **Azure Cosmos DB VS Code Extension** for database management
- **Azure Storage Emulator** (if using blob storage for future features)

---

## Project Structure

```
holiday-finder/
├── src/
│   ├── HolidayFinder.Api/              # ASP.NET Core Web API
│   │   ├── Controllers/                 # API controllers
│   │   ├── Services/                    # Business logic services
│   │   ├── Models/                      # Cosmos DB models
│   │   ├── Middleware/                  # Custom middleware
│   │   ├── Infrastructure/              # Cosmos DB setup, auth config
│   │   └── Program.cs                   # API startup
│   │
│   ├── HolidayFinder.Client/            # Blazor WebAssembly
│   │   ├── Pages/                       # Blazor pages (Home, Search, etc.)
│   │   ├── Components/                  # Reusable components
│   │   ├── Services/                    # HTTP client services
│   │   ├── wwwroot/                     # Static assets
│   │   └── Program.cs                   # Client startup
│   │
│   └── HolidayFinder.Shared/            # Shared DTOs and contracts
│       ├── DTOs/                        # Request/Response models
│       ├── Enums/                       # Shared enumerations
│       └── Constants/                   # Shared constants
│
├── tests/
│   ├── HolidayFinder.Api.Tests/         # xUnit tests for API
│   │   ├── Controllers/                 # Controller tests
│   │   ├── Services/                    # Service tests
│   │   └── Integration/                 # Integration tests
│   │
│   └── HolidayFinder.Client.Tests/      # bUnit tests for Blazor
│       └── Components/                  # Component tests
│
├── specs/                               # Specification and planning docs
│   └── 001-ai-holiday-booking/
│       ├── spec.md                      # Feature specification
│       ├── plan.md                      # Implementation plan
│       ├── research.md                  # Technology research
│       ├── data-model.md                # Database schema
│       └── contracts/                   # API contracts
│
└── .specify/                            # Spec Kit configuration
```

---

## Step 1: Clone and Setup

### 1.1 Create Solution and Projects

```powershell
# Create solution
dotnet new sln -n HolidayFinder

# Create API project
dotnet new webapi -n HolidayFinder.Api -o src/HolidayFinder.Api
dotnet sln add src/HolidayFinder.Api/HolidayFinder.Api.csproj

# Create Blazor WASM project
dotnet new blazorwasm -n HolidayFinder.Client -o src/HolidayFinder.Client
dotnet sln add src/HolidayFinder.Client/HolidayFinder.Client.csproj

# Create Shared library
dotnet new classlib -n HolidayFinder.Shared -o src/HolidayFinder.Shared
dotnet sln add src/HolidayFinder.Shared/HolidayFinder.Shared.csproj

# Create Test projects
dotnet new xunit -n HolidayFinder.Api.Tests -o tests/HolidayFinder.Api.Tests
dotnet new bunit -n HolidayFinder.Client.Tests -o tests/HolidayFinder.Client.Tests
dotnet sln add tests/HolidayFinder.Api.Tests/HolidayFinder.Api.Tests.csproj
dotnet sln add tests/HolidayFinder.Client.Tests/HolidayFinder.Client.Tests.csproj
```

### 1.2 Add Project References

```powershell
# API references Shared
dotnet add src/HolidayFinder.Api/HolidayFinder.Api.csproj reference src/HolidayFinder.Shared/HolidayFinder.Shared.csproj

# Client references Shared
dotnet add src/HolidayFinder.Client/HolidayFinder.Client.csproj reference src/HolidayFinder.Shared/HolidayFinder.Shared.csproj

# Test projects reference their respective projects
dotnet add tests/HolidayFinder.Api.Tests/HolidayFinder.Api.Tests.csproj reference src/HolidayFinder.Api/HolidayFinder.Api.csproj
dotnet add tests/HolidayFinder.Client.Tests/HolidayFinder.Client.Tests.csproj reference src/HolidayFinder.Client/HolidayFinder.Client.csproj
```

---

## Step 2: Install NuGet Packages

### 2.1 API Project Packages

```powershell
cd src/HolidayFinder.Api

# Cosmos DB
dotnet add package Microsoft.Azure.Cosmos --version 3.39.0

# Authentication
dotnet add package Microsoft.AspNetCore.Authentication.JwtBearer --version 8.0.0
dotnet add package Microsoft.AspNetCore.Identity.EntityFrameworkCore --version 8.0.0

# Swagger/OpenAPI (included in webapi template)
# dotnet add package Swashbuckle.AspNetCore --version 6.5.0

# Application Insights (optional for production)
dotnet add package Microsoft.ApplicationInsights.AspNetCore --version 2.22.0

# For HATEOAS (custom implementation or library)
# Option: Build custom or use a community package like HAL.AspNetCore
```

### 2.2 Client Project Packages

```powershell
cd ../HolidayFinder.Client

# HTTP Client extensions
dotnet add package Microsoft.Extensions.Http --version 8.0.0

# Blazor Authentication
dotnet add package Microsoft.AspNetCore.Components.WebAssembly.Authentication --version 8.0.0

# Local storage (for JWT storage)
dotnet add package Blazored.LocalStorage --version 4.5.0

# Optional: UI component library
# dotnet add package MudBlazor --version 6.19.0
# Or: dotnet add package Blazorise --version 5.1.0
```

### 2.3 Test Project Packages

```powershell
cd ../../tests/HolidayFinder.Api.Tests

# Web application testing
dotnet add package Microsoft.AspNetCore.Mvc.Testing --version 8.0.0

# Mocking
dotnet add package Moq --version 4.20.0

# Fluent Assertions
dotnet add package FluentAssertions --version 6.12.0

# For Cosmos DB integration tests
dotnet add package Microsoft.Azure.Cosmos --version 3.39.0
```

---

## Step 3: Azure Cosmos DB Setup

### Option A: Local Emulator (Recommended for Development)

1. **Install Cosmos DB Emulator**:
   - Download from [Microsoft Docs](https://learn.microsoft.com/azure/cosmos-db/emulator)
   - Or use Docker:
   ```powershell
   docker run -p 8081:8081 -p 10250-10255:10250-10255 `
     --name cosmosdb-emulator `
     -e AZURE_COSMOS_EMULATOR_PARTITION_COUNT=10 `
     -e AZURE_COSMOS_EMULATOR_ENABLE_DATA_PERSISTENCE=true `
     mcr.microsoft.com/cosmosdb/emulator
   ```

2. **Access Emulator**:
   - Navigate to `https://localhost:8081/_explorer/index.html`
   - Accept the SSL certificate warning (self-signed certificate)

3. **Connection String** (use in appsettings.Development.json):
   ```json
   {
     "CosmosDb": {
       "AccountEndpoint": "https://localhost:8081",
       "AccountKey": "C2y6yDjf5/R+ob0N8A7Cgv30VRDJIWEHLM+4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw/Jw==",
       "DatabaseName": "HolidayFinderDb"
     }
   }
   ```

### Option B: Azure Cosmos DB (Production/Shared Dev)

1. **Create Cosmos DB Account**:
   ```powershell
   # Login to Azure
   az login
   
   # Create resource group
   az group create --name rg-holidayfinder --location eastus
   
   # Create Cosmos DB account
   az cosmosdb create `
     --name cosmos-holidayfinder `
     --resource-group rg-holidayfinder `
     --default-consistency-level Session `
     --locations regionName=eastus failoverPriority=0 isZoneRedundant=False
   
   # Get connection string
   az cosmosdb keys list `
     --name cosmos-holidayfinder `
     --resource-group rg-holidayfinder `
     --type connection-strings
   ```

2. **Update appsettings.json**:
   ```json
   {
     "CosmosDb": {
       "AccountEndpoint": "https://cosmos-holidayfinder.documents.azure.com:443/",
       "AccountKey": "YOUR_ACCOUNT_KEY_HERE",
       "DatabaseName": "HolidayFinderDb"
     }
   }
   ```

### 3.1 Initialize Database Schema

The API will auto-create database and containers on first run. See `Infrastructure/CosmosDbService.cs` for initialization code.

**Containers to Create**:
- `Users` (partition key: `/id`)
- `SearchCriteria` (partition key: `/userId`)
- `SearchHistory` (partition key: `/userId`)
- `MembershipTiers` (partition key: `/id`)

---

## Step 4: Configure API Settings

### 4.1 Create appsettings.Development.json

```json
{
  "Logging": {
    "LogLevel": {
      "Default": "Information",
      "Microsoft.AspNetCore": "Warning"
    }
  },
  "CosmosDb": {
    "AccountEndpoint": "https://localhost:8081",
    "AccountKey": "C2y6yDjf5/R+ob0N8A7Cgv30VRDJIWEHLM+4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw/Jw==",
    "DatabaseName": "HolidayFinderDb"
  },
  "Jwt": {
    "Key": "YOUR_SECRET_KEY_HERE_MIN_32_CHARS",
    "Issuer": "https://localhost:7001",
    "Audience": "holidayfinder-api",
    "ExpiryMinutes": 60
  },
  "Cors": {
    "AllowedOrigins": [
      "https://localhost:7002",
      "http://localhost:5000"
    ]
  },
  "AIService": {
    "Endpoint": "https://your-ai-service.azure.com",
    "ApiKey": "YOUR_AI_SERVICE_KEY",
    "Timeout": 10
  }
}
```

### 4.2 Generate JWT Secret Key

```powershell
# Generate a random 32-character key
$bytes = New-Object byte[] 32
[Security.Cryptography.RandomNumberGenerator]::Create().GetBytes($bytes)
[Convert]::ToBase64String($bytes)
```

---

## Step 5: Run the Application

### 5.1 Start the API

```powershell
cd src/HolidayFinder.Api
dotnet run
```

API will be available at:
- HTTPS: `https://localhost:7001`
- HTTP: `http://localhost:5000`
- Swagger: `https://localhost:7001/swagger`

### 5.2 Start the Blazor Client

```powershell
cd src/HolidayFinder.Client
dotnet run
```

Client will be available at:
- HTTPS: `https://localhost:7002`
- HTTP: `http://localhost:5001`

### 5.3 Run Both in Visual Studio

1. Right-click solution → **Properties**
2. Select **Multiple startup projects**
3. Set both `HolidayFinder.Api` and `HolidayFinder.Client` to **Start**
4. Press **F5** to debug

---

## Step 6: Verify Installation

### 6.1 Test API

```powershell
# Health check
curl https://localhost:7001/api/health

# Get membership tiers (public endpoint)
curl https://localhost:7001/api/memberships
```

### 6.2 Test Swagger UI

1. Navigate to `https://localhost:7001/swagger`
2. Expand `POST /api/auth/register`
3. Click **Try it out**
4. Submit sample registration data
5. Verify 201 Created response

### 6.3 Test Blazor Client

1. Navigate to `https://localhost:7002`
2. Click "Get Started" or "Register"
3. Create a test account
4. Verify redirect to search page

---

## Step 7: Seed Sample Data (Optional)

### 7.1 Create Seed Data Script

Create `src/HolidayFinder.Api/Infrastructure/DataSeeder.cs`:

```csharp
public class DataSeeder
{
    public static async Task SeedAsync(CosmosClient client, string databaseName)
    {
        var database = client.GetDatabase(databaseName);
        
        // Seed membership tiers
        var tiersContainer = database.GetContainer("MembershipTiers");
        
        var freeTier = new MembershipTier
        {
            Id = "tier-free",
            Name = "Free",
            DisplayName = "Free Explorer",
            // ... other properties
        };
        
        await tiersContainer.UpsertItemAsync(freeTier, 
            new PartitionKey(freeTier.Id));
        
        // Add Premium and VIP tiers similarly
    }
}
```

### 7.2 Run Seeder

Add to `Program.cs`:

```csharp
// After app.Build()
if (app.Environment.IsDevelopment())
{
    await DataSeeder.SeedAsync(cosmosClient, "HolidayFinderDb");
}
```

---

## Step 8: Run Tests

```powershell
# Run all tests
dotnet test

# Run with code coverage
dotnet test /p:CollectCoverage=true /p:CoverageReportsDirectory=./coverage

# Run specific test project
dotnet test tests/HolidayFinder.Api.Tests/HolidayFinder.Api.Tests.csproj
```

---

## Development Workflow

### Daily Development

1. **Pull latest changes**: `git pull`
2. **Restore packages**: `dotnet restore`
3. **Build solution**: `dotnet build`
4. **Run tests**: `dotnet test`
5. **Start emulator**: Ensure Cosmos DB emulator is running
6. **Launch apps**: Start API and Client

### Adding New Features

1. **Create spec**: Use `/speckit.specify` command
2. **Plan implementation**: Use `/speckit.plan` command
3. **Create tests**: Write tests first (TDD from constitution)
4. **Implement**: Follow plan and contracts
5. **Verify**: Run tests, check code coverage

### Database Changes

1. **Update data-model.md**: Document schema changes
2. **Create migration script**: If needed
3. **Test locally**: Verify with emulator
4. **Update seeder**: If adding new reference data

---

## Troubleshooting

### Cosmos DB Emulator Issues

**Problem**: "The SSL connection could not be established"

**Solution**:
```powershell
# Trust the emulator certificate
$cert = Get-ChildItem Cert:\LocalMachine\Root | Where-Object { $_.Subject -like "*Cosmos DB Emulator*" }
Export-Certificate -Cert $cert -FilePath "$env:TEMP\CosmosDbEmulator.cer"
Import-Certificate -FilePath "$env:TEMP\CosmosDbEmulator.cer" -CertStoreLocation Cert:\CurrentUser\Root
```

### Port Conflicts

**Problem**: Port 7001 or 7002 already in use

**Solution**: Update `launchSettings.json` in both projects:
```json
{
  "applicationUrl": "https://localhost:7011;http://localhost:5010"
}
```

### CORS Errors

**Problem**: Blazor client gets CORS errors when calling API

**Solution**: Verify `appsettings.Development.json` includes client URL in `Cors.AllowedOrigins`

### JWT Token Issues

**Problem**: 401 Unauthorized on protected endpoints

**Solution**:
1. Verify JWT secret key is same in config
2. Check token expiration
3. Ensure `Authorization: Bearer {token}` header is included

---

## Next Steps

1. ✅ **Setup Complete** - You should now have a working development environment
2. 📝 **Read Contracts** - Review API contracts in `specs/001-ai-holiday-booking/contracts/`
3. 🧪 **Write Tests** - Start with user registration tests (TDD approach)
4. 🔨 **Implement Features** - Follow the implementation plan in `plan.md`
5. 📊 **Monitor Progress** - Use `/speckit.tasks` to generate task list

---

## Additional Resources

- [ASP.NET Core Docs](https://learn.microsoft.com/aspnet/core/)
- [Blazor WebAssembly Docs](https://learn.microsoft.com/aspnet/core/blazor/)
- [Azure Cosmos DB .NET SDK](https://learn.microsoft.com/azure/cosmos-db/nosql/sdk-dotnet-v3)
- [xUnit Documentation](https://xunit.net/)
- [bUnit Documentation](https://bunit.dev/)
- [Constitution](../../.specify/memory/constitution.md)

---

**Questions?** Check the specification artifacts in `specs/001-ai-holiday-booking/` or refer to the plan.md file.
