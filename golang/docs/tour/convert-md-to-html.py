import os
import re
from pathlib import Path

def escape_html(text):
    """Escape HTML special characters"""
    return (text
            .replace('&', '&amp;')
            .replace('<', '&lt;')
            .replace('>', '&gt;')
            .replace('"', '&quot;'))

def process_markdown_block(text, in_list=False):
    """Convert markdown formatting to HTML"""
    # Process inline code first (before escaping)
    text = re.sub(r'`([^`]+)`', lambda m: f'<code>{escape_html(m.group(1))}</code>', text)
    
    # Bold text
    text = re.sub(r'\*\*([^*]+)\*\*', r'<strong>\1</strong>', text)
    
    # Italic text
    text = re.sub(r'\*([^*]+)\*', r'<em>\1</em>', text)
    
    return text

def convert_markdown_to_html(md_file, output_file, section):
    """Convert a markdown file to HTML"""
    with open(md_file, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # Extract title
    title_match = re.match(r'^#\s+(.+)$', content, re.MULTILINE)
    title = title_match.group(1) if title_match else "Go Tour"
    
    # Determine asset path depth
    depth = "../.."
    
    html_sections = []
    
    # Split content by ## headers
    parts = re.split(r'^## ', content, flags=re.MULTILINE)
    
    for i, part in enumerate(parts):
        if i == 0:
            continue  # Skip content before first ##
        
        # Determine section type
        section_class = ""
        if part.startswith('Go Concept'):
            section_class = "go-section"
            heading = "Go Concept"
            section_content = part[10:].strip()
        elif part.startswith('C# Equivalent'):
            section_class = "csharp-section"
            heading = "C# Equivalent"
            section_content = part[13:].strip()
        elif part.startswith('Key Differences'):
            section_class = "key-differences"
            heading = "Key Differences"
            section_content = part[15:].strip()
        else:
            continue
        
        # Start section div
        section_html = f'<div class="content-section {section_class}">\n'
        section_html += f'<h2>{heading}</h2>\n'
        
        # Process content
        lines = section_content.split('\n')
        in_code_block = False
        in_list = False
        code_lang = ''
        code_lines = []
        paragraph_lines = []
        list_items = []
        
        for line in lines:
            # Handle code blocks
            if line.strip().startswith('```'):
                if in_code_block:
                    # End code block
                    code_text = '\n'.join(code_lines)
                    section_html += f'<pre><code class="language-{code_lang}">{escape_html(code_text)}</code></pre>\n'
                    code_lines = []
                    in_code_block = False
                    code_lang = ''
                else:
                    # Start code block
                    if paragraph_lines:
                        section_html += f'<p>{process_markdown_block(" ".join(paragraph_lines))}</p>\n'
                        paragraph_lines = []
                    if list_items:
                        section_html += '<ul>\n'
                        for item in list_items:
                            section_html += f'<li>{process_markdown_block(item)}</li>\n'
                        section_html += '</ul>\n'
                        list_items = []
                        in_list = False
                    
                    in_code_block = True
                    lang_match = re.match(r'```(\w+)', line.strip())
                    code_lang = lang_match.group(1) if lang_match else ''
                continue
            
            if in_code_block:
                code_lines.append(line)
                continue
            
            # Handle ### headers
            if line.strip().startswith('###'):
                if paragraph_lines:
                    section_html += f'<p>{process_markdown_block(" ".join(paragraph_lines))}</p>\n'
                    paragraph_lines = []
                if list_items:
                    section_html += '<ul>\n'
                    for item in list_items:
                        section_html += f'<li>{process_markdown_block(item)}</li>\n'
                    section_html += '</ul>\n'
                    list_items = []
                    in_list = False
                
                h3_text = line.strip()[3:].strip()
                section_html += f'<h3>{process_markdown_block(h3_text)}</h3>\n'
                continue
            
            # Handle bullet points
            if line.strip().startswith('- '):
                if paragraph_lines:
                    section_html += f'<p>{process_markdown_block(" ".join(paragraph_lines))}</p>\n'
                    paragraph_lines = []
                
                in_list = True
                item_text = line.strip()[2:]
                list_items.append(item_text)
                continue
            
            # Handle empty lines
            if not line.strip():
                if paragraph_lines:
                    section_html += f'<p>{process_markdown_block(" ".join(paragraph_lines))}</p>\n'
                    paragraph_lines = []
                if list_items:
                    section_html += '<ul>\n'
                    for item in list_items:
                        section_html += f'<li>{process_markdown_block(item)}</li>\n'
                    section_html += '</ul>\n'
                    list_items = []
                    in_list = False
                continue
            
            # Regular paragraph text
            if not in_list:
                paragraph_lines.append(line.strip())
            else:
                # Continuation of list item
                list_items[-1] += ' ' + line.strip()
        
        # Flush remaining content
        if paragraph_lines:
            section_html += f'<p>{process_markdown_block(" ".join(paragraph_lines))}</p>\n'
        if list_items:
            section_html += '<ul>\n'
            for item in list_items:
                section_html += f'<li>{process_markdown_block(item)}</li>\n'
            section_html += '</ul>\n'
        
        section_html += '</div>\n'
        html_sections.append(section_html)
    
    # Build complete HTML document
    html = f'''<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{title} - Go Tour with C# Comparison</title>
    <link rel="stylesheet" href="{depth}/assets/style.css">
</head>
<body>
    <button class="menu-toggle">☰ Menu</button>
    
    <nav class="nav-sidebar">
        <!-- Navigation will be populated by JavaScript -->
    </nav>
    
    <main class="main-content">
        <div class="content-header">
            <h1>{title}</h1>
        </div>
        
        {''.join(html_sections)}
        
        <div class="page-nav" id="pageNav">
            <!-- Navigation buttons will be populated by JavaScript -->
        </div>
    </main>
    
    <script src="{depth}/assets/tour.js"></script>
    <script>
        // Add page navigation
        const adjacent = getAdjacentPages();
        const navContainer = document.getElementById('pageNav');
        let navHtml = '';
        
        if (adjacent.prev) {{
            navHtml += '<a href="' + adjacent.prev.href + '" class="prev">' + adjacent.prev.title + '</a>';
        }} else {{
            navHtml += '<span></span>';
        }}
        
        if (adjacent.next) {{
            navHtml += '<a href="' + adjacent.next.href + '" class="next">' + adjacent.next.title + '</a>';
        }}
        
        navContainer.innerHTML = navHtml;
    </script>
</body>
</html>
'''
    
    # Write HTML file
    with open(output_file, 'w', encoding='utf-8') as f:
        f.write(html)
    
    print(f"Converted: {output_file}")

def main():
    base_dir = Path(r"c:\Repos\TestsAndIdeas\Personal-Learning\golang\docs\tour")
    sections = ['basics', 'flowcontrol', 'moretypes', 'methods', 'generics', 'concurrency']
    
    for section in sections:
        section_dir = base_dir / section
        md_files = list(section_dir.glob('*.md'))
        
        for md_file in md_files:
            output_file = md_file.with_suffix('.html')
            convert_markdown_to_html(md_file, output_file, section)
    
    print("\nConversion complete!")

if __name__ == '__main__':
    main()
