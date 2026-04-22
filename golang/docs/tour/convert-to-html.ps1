# Convert Markdown Tour Pages to HTML
param(
    [string]$SourceDir = "c:\Repos\TestsAndIdeas\Personal-Learning\golang\docs\tour",
    [string]$OutputDir = "c:\Repos\TestsAndIdeas\Personal-Learning\golang\docs\tour"
)

function Convert-MarkdownToHtml {
    param(
        [string]$MarkdownFile,
        [string]$OutputFile,
        [string]$Section
    )
    
    $content = Get-Content $MarkdownFile -Raw
    
    # Extract title from first # heading
    $title = ""
    if ($content -match '^#\s+(.+)$') {
        $title = $matches[1]
    }
    
    # Convert markdown to HTML sections
    $htmlContent = ""
    
    # Split by ## headings (Go Concept, C# Equivalent, Key Differences)
    $sections = $content -split '(?=^## )'
    
    foreach ($section in $sections) {
        if ([string]::IsNullOrWhiteSpace($section)) { continue }
        
        # Check section type
        if ($section -match '^## Go Concept') {
            $htmlContent += '<div class="content-section go-section">'
            $htmlContent += '<h2>Go Concept</h2>'
            $sectionContent = $section -replace '^## Go Concept\s*', ''
        }
        elseif ($section -match '^## C# Equivalent') {
            $htmlContent += '<div class="content-section csharp-section">'
            $htmlContent += '<h2>C# Equivalent</h2>'
            $sectionContent = $section -replace '^## C# Equivalent\s*', ''
        }
        elseif ($section -match '^## Key Differences') {
            $htmlContent += '<div class="content-section key-differences">'
            $htmlContent += '<h2>Key Differences</h2>'
            $sectionContent = $section -replace '^## Key Differences\s*', ''
        }
        else {
            # Handle title section
            if ($section -match '^#\s+(.+)') {
                continue  # Skip title, we'll add it separately
            }
            $sectionContent = $section
        }
        
        # Convert markdown elements to HTML
        # Code blocks
        $sectionContent = $sectionContent -replace '```go\s*\n([\s\S]*?)\n```', '<pre><code class="language-go">$1</code></pre>'
        $sectionContent = $sectionContent -replace '```csharp\s*\n([\s\S]*?)\n```', '<pre><code class="language-csharp">$1</code></pre>'
        $sectionContent = $sectionContent -replace '```\s*\n([\s\S]*?)\n```', '<pre><code>$1</code></pre>'
        
        # Inline code
        $sectionContent = $sectionContent -replace '`([^`]+)`', '<code>$1</code>'
        
        # Bold
        $sectionContent = $sectionContent -replace '\*\*([^*]+)\*\*', '<strong>$1</strong>'
        
        # ### headings
        $sectionContent = $sectionContent -replace '^### (.+)$', '<h3>$1</h3>' -replace '\r?\n', "`n"
        
        # Bullet lists
        $sectionContent = $sectionContent -replace '(?m)^- (.+)$', '<li>$1</li>'
        $sectionContent = $sectionContent -replace '(<li>.*</li>)', '<ul>$1</ul>'
        
        # Paragraphs
        $lines = $sectionContent -split '\r?\n\r?\n'
        $processedLines = @()
        foreach ($line in $lines) {
            $line = $line.Trim()
            if ([string]::IsNullOrWhiteSpace($line)) { continue }
            if ($line -match '^<(h3|pre|ul|code)') {
                $processedLines += $line
            }
            else {
                $processedLines += "<p>$line</p>"
            }
        }
        $sectionContent = $processedLines -join "`n"
        
        $htmlContent += $sectionContent
        $htmlContent += '</div>'
    }
    
    # Determine relative path for assets
    $depth = ".."
    if ($Section -ne "root") {
        $depth = "../.."
    }
    
    # Create full HTML document
    $html = @"
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>$title - Go Tour with C# Comparison</title>
    <link rel="stylesheet" href="$depth/assets/style.css">
</head>
<body>
    <button class="menu-toggle">☰ Menu</button>
    
    <nav class="nav-sidebar">
        <!-- Navigation will be populated by JavaScript -->
    </nav>
    
    <main class="main-content">
        <div class="content-header">
            <h1>$title</h1>
        </div>
        
        $htmlContent
        
        <div class="page-nav" id="pageNav">
            <!-- Navigation buttons will be populated by JavaScript -->
        </div>
    </main>
    
    <script src="$depth/assets/tour.js"></script>
    <script>
        // Add page navigation
        const adjacent = getAdjacentPages();
        const navContainer = document.getElementById('pageNav');
        let navHtml = '';
        
        if (adjacent.prev) {
            navHtml += '<a href="' + adjacent.prev.href + '" class="prev">' + adjacent.prev.title + '</a>';
        } else {
            navHtml += '<span></span>';
        }
        
        if (adjacent.next) {
            navHtml += '<a href="' + adjacent.next.href + '" class="next">' + adjacent.next.title + '</a>';
        }
        
        navContainer.innerHTML = navHtml;
    </script>
</body>
</html>
"@
    
    # Write HTML file
    $html | Out-File -FilePath $OutputFile -Encoding UTF8
    Write-Host "Converted: $OutputFile"
}

# Convert all markdown files
$sections = @('basics', 'flowcontrol', 'moretypes', 'methods', 'generics', 'concurrency')

foreach ($section in $sections) {
    $sectionPath = Join-Path $SourceDir $section
    $mdFiles = Get-ChildItem -Path $sectionPath -Filter "*.md"
    
    foreach ($mdFile in $mdFiles) {
        $outputFile = Join-Path $sectionPath ($mdFile.BaseName + ".html")
        Convert-MarkdownToHtml -MarkdownFile $mdFile.FullName -OutputFile $outputFile -Section $section
    }
}

Write-Host "`nConversion complete! All HTML files created."
