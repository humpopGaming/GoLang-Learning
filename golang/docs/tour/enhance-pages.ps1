# Script to enhance all tutorial pages with breadcrumb navigation and better structure

$sections = @{
    "basics" = "Basics"
    "flowcontrol" = "Flow Control"
    "moretypes" = "More Types"
    "methods" = "Methods"
    "generics" = "Generics"
    "concurrency" = "Concurrency"
}

$tourRoot = $PSScriptRoot
Write-Host "Processing tour pages in: $tourRoot" -ForegroundColor Cyan

foreach ($section in $sections.Keys) {
    $sectionPath = Join-Path $tourRoot $section
    if (-not (Test-Path $sectionPath)) {
        Write-Host "Section not found: $sectionPath" -ForegroundColor Yellow
        continue
    }
    
    $htmlFiles = Get-ChildItem -Path $sectionPath -Filter "*.html" | Sort-Object Name
    $pageCount = $htmlFiles.Count
    
    Write-Host "`nProcessing $pageCount pages in $($sections[$section])..." -ForegroundColor Green
    
    for ($i = 0; $i -lt $htmlFiles.Count; $i++) {
        $file = $htmlFiles[$i]
        $pageNumber = $i + 1
        
        # Read the file content
        $content = Get-Content $file.FullName -Raw -Encoding UTF8
        
        # Extract the title from the h1 tag
        if ($content -match '<h1>(.*?)</h1>') {
            $pageTitle = $matches[1]
        } else {
            $pageTitle = $file.BaseName
        }
        
        # Check if breadcrumb already exists
        if ($content -notmatch 'breadcrumb') {
            # Create breadcrumb HTML
            $breadcrumb = "        <div class=`"breadcrumb`">`r`n"
            $breadcrumb += "            <a href=`"../../index.html`">Home</a>`r`n"
            $breadcrumb += "            <span>/</span>`r`n"
            $breadcrumb += "            <a href=`"$(($htmlFiles[0]).Name)`">$($sections[$section])</a>`r`n"
            $breadcrumb += "            <span>/</span>`r`n"
            $breadcrumb += "            <span>$pageTitle</span>`r`n"
            $breadcrumb += "        </div>`r`n        `r`n"
            
            # Create page meta badges
            $pageMeta = "            <div class=`"page-meta`">`r`n"
            $pageMeta += "                <span class=`"page-badge section`">$($sections[$section])</span>`r`n"
            $pageMeta += "                <span class=`"page-badge page-number`">Page $pageNumber of $pageCount</span>`r`n"
            $pageMeta += "            </div>"
            
            # Insert breadcrumb after <main class="main-content">
            $content = $content -replace '(<main class="main-content">)', "`$1`r`n$breadcrumb"
            
            # Insert page meta after <h1>
            $escapedTitle = [regex]::Escape($pageTitle)
            $content = $content -replace "(<h1>$escapedTitle</h1>)", "`$1`r`n$pageMeta"
            
            # Save the modified content
            Set-Content -Path $file.FullName -Value $content -Encoding UTF8 -NoNewline
            
            Write-Host "  Enhanced: $($file.Name)" -ForegroundColor Gray
        } else {
            Write-Host "  Skipped: $($file.Name)" -ForegroundColor DarkGray
        }
    }
}

Write-Host "`nAll pages have been enhanced!" -ForegroundColor Green

