// Tour Navigation Data
const tourStructure = {
    'basics': {
        title: 'Basics',
        pages: [
            { id: '01-packages', title: 'Packages' },
            { id: '02-imports', title: 'Imports' },
            { id: '03-exported-names', title: 'Exported Names' },
            { id: '04-functions', title: 'Functions' },
            { id: '05-functions-continued', title: 'Functions Continued' },
            { id: '06-multiple-results', title: 'Multiple Results' },
            { id: '07-named-return-values', title: 'Named Return Values' },
            { id: '08-variables', title: 'Variables' },
            { id: '09-variables-with-initializers', title: 'Variables with Initializers' },
            { id: '10-short-variable-declarations', title: 'Short Variable Declarations' },
            { id: '11-basic-types', title: 'Basic Types' },
            { id: '12-zero-values', title: 'Zero Values' },
            { id: '13-type-conversions', title: 'Type Conversions' },
            { id: '14-type-inference', title: 'Type Inference' },
            { id: '15-constants', title: 'Constants' },
            { id: '16-numeric-constants', title: 'Numeric Constants' }
        ]
    },
    'flowcontrol': {
        title: 'Flow Control',
        pages: [
            { id: '01-for', title: 'For' },
            { id: '02-for-continued', title: 'For Continued' },
            { id: '03-for-is-while', title: 'For is While' },
            { id: '04-forever', title: 'Forever' },
            { id: '05-if', title: 'If' },
            { id: '06-if-with-short-statement', title: 'If with Short Statement' },
            { id: '07-if-and-else', title: 'If and Else' },
            { id: '08-exercise-loops-and-functions', title: 'Exercise: Loops and Functions' },
            { id: '09-switch', title: 'Switch' },
            { id: '10-switch-evaluation-order', title: 'Switch Evaluation Order' },
            { id: '11-switch-with-no-condition', title: 'Switch with No Condition' },
            { id: '12-defer', title: 'Defer' },
            { id: '13-stacking-defers', title: 'Stacking Defers' }
        ]
    },
    'moretypes': {
        title: 'More Types',
        pages: [
            { id: '01-pointers', title: 'Pointers' },
            { id: '02-structs', title: 'Structs' },
            { id: '03-struct-fields', title: 'Struct Fields' },
            { id: '04-pointers-to-structs', title: 'Pointers to Structs' },
            { id: '05-struct-literals', title: 'Struct Literals' },
            { id: '06-arrays', title: 'Arrays' },
            { id: '07-slices', title: 'Slices' },
            { id: '08-slices-are-like-references', title: 'Slices are like References' },
            { id: '09-slice-literals', title: 'Slice Literals' },
            { id: '10-slice-defaults', title: 'Slice Defaults' },
            { id: '11-slice-length-and-capacity', title: 'Slice Length and Capacity' },
            { id: '12-nil-slices', title: 'Nil Slices' },
            { id: '13-creating-slices-with-make', title: 'Creating Slices with Make' },
            { id: '14-slices-of-slices', title: 'Slices of Slices' },
            { id: '15-appending-to-slices', title: 'Appending to Slices' },
            { id: '16-range', title: 'Range' },
            { id: '17-range-continued', title: 'Range Continued' },
            { id: '18-exercise-slices', title: 'Exercise: Slices' },
            { id: '19-maps', title: 'Maps' },
            { id: '20-map-literals', title: 'Map Literals' },
            { id: '21-map-literals-continued', title: 'Map Literals Continued' },
            { id: '22-mutating-maps', title: 'Mutating Maps' },
            { id: '23-exercise-maps', title: 'Exercise: Maps' },
            { id: '24-function-values', title: 'Function Values' },
            { id: '25-function-closures', title: 'Function Closures' }
        ]
    },
    'methods': {
        title: 'Methods',
        pages: [
            { id: '01-methods', title: 'Methods' },
            { id: '02-methods-are-functions', title: 'Methods Are Functions' },
            { id: '03-methods-continued', title: 'Methods Continued' },
            { id: '04-pointer-receivers', title: 'Pointer Receivers' },
            { id: '05-pointers-and-functions', title: 'Pointers and Functions' },
            { id: '06-methods-and-pointer-indirection', title: 'Methods and Pointer Indirection' },
            { id: '07-methods-and-pointer-indirection-2', title: 'Methods and Pointer Indirection (2)' },
            { id: '08-choosing-value-or-pointer-receiver', title: 'Choosing Value or Pointer Receiver' },
            { id: '09-interfaces', title: 'Interfaces' },
            { id: '10-interfaces-are-implemented-implicitly', title: 'Interfaces Are Implemented Implicitly' },
            { id: '11-interface-values', title: 'Interface Values' },
            { id: '12-interface-values-with-nil-underlying-values', title: 'Interface Values with Nil Underlying Values' },
            { id: '13-nil-interface-values', title: 'Nil Interface Values' },
            { id: '14-the-empty-interface', title: 'The Empty Interface' },
            { id: '15-type-assertions', title: 'Type Assertions' },
            { id: '16-type-switches', title: 'Type Switches' },
            { id: '17-stringers', title: 'Stringers' },
            { id: '18-exercise-stringers', title: 'Exercise: Stringers' },
            { id: '19-errors', title: 'Errors' },
            { id: '20-exercise-errors', title: 'Exercise: Errors' },
            { id: '21-readers', title: 'Readers' },
            { id: '22-exercise-readers', title: 'Exercise: Readers' },
            { id: '23-exercise-rot13reader', title: 'Exercise: rot13Reader' }
        ]
    },
    'generics': {
        title: 'Generics',
        pages: [
            { id: '01-type-parameters', title: 'Type Parameters' },
            { id: '02-generic-types', title: 'Generic Types' }
        ]
    },
    'concurrency': {
        title: 'Concurrency',
        pages: [
            { id: '01-goroutines', title: 'Goroutines' },
            { id: '02-channels', title: 'Channels' },
            { id: '03-buffered-channels', title: 'Buffered Channels' },
            { id: '04-range-and-close', title: 'Range and Close' },
            { id: '05-select', title: 'Select' },
            { id: '06-default-selection', title: 'Default Selection' },
            { id: '07-exercise-equivalent-binary-trees', title: 'Exercise: Equivalent Binary Trees' },
            { id: '08-exercise-equivalent-binary-trees-2', title: 'Exercise: Equivalent Binary Trees (2)' },
            { id: '09-sync-mutex', title: 'sync.Mutex' }
        ]
    }
};

// Initialize the tour page
function initTour() {
    renderNavigation();
    highlightCurrentPage();
    setupMobileMenu();
}

// Render sidebar navigation
function renderNavigation() {
    const nav = document.querySelector('.nav-sidebar');
    if (!nav) return;

    // Determine current section from the URL
    const currentPath = window.location.pathname;
    const pathMatch = currentPath.match(/\/([^\/]+)\/([^\/]+)\.html/);
    const currentSection = pathMatch ? pathMatch[1] : null;
    
    // Check if we're on the index page
    const isIndexPage = currentPath.includes('index.html') && !pathMatch;

    // Set home link based on current location
    const homeLink = isIndexPage ? 'index.html' : '../index.html';
    let html = `<h1><a href="${homeLink}" style="color: inherit; text-decoration: none;">Go Tour</a></h1>`;
    
    for (const [section, data] of Object.entries(tourStructure)) {
        html += `<div class="nav-section">`;
        html += `<h3>${data.title}</h3>`;
        html += `<ul>`;
        
        for (const page of data.pages) {
            // Generate relative path based on current location
            let href;
            if (isIndexPage) {
                // On index page - direct path to section folder
                href = `${section}/${page.id}.html`;
            } else if (currentSection === section) {
                // Same section - just the filename
                href = `${page.id}.html`;
            } else {
                // Different section - need to go up and then into target section
                href = `../${section}/${page.id}.html`;
            }
            html += `<li><a href="${href}" data-section="${section}" data-page="${page.id}">${page.title}</a></li>`;
        }
        
        html += `</ul></div>`;
    }
    
    nav.innerHTML = html;
}

// Highlight current page in navigation
function highlightCurrentPage() {
    const currentPath = window.location.pathname;
    const links = document.querySelectorAll('.nav-sidebar a');
    
    links.forEach(link => {
        if (currentPath.includes(link.getAttribute('href'))) {
            link.classList.add('active');
        }
    });
}

// Find next and previous pages
function getAdjacentPages() {
    const currentPath = window.location.pathname;
    const pathMatch = currentPath.match(/\/([^\/]+)\/([^\/]+)\.html/);
    
    if (!pathMatch) return { prev: null, next: null };
    
    const [, currentSection, currentId] = pathMatch;
    const sections = Object.keys(tourStructure);
    const currentSectionIndex = sections.indexOf(currentSection);
    const currentSectionData = tourStructure[currentSection];
    const currentPageIndex = currentSectionData.pages.findIndex(p => p.id === currentId);
    
    let prev = null;
    let next = null;
    
    // Find next page
    if (currentPageIndex < currentSectionData.pages.length - 1) {
        const nextPage = currentSectionData.pages[currentPageIndex + 1];
        next = {
            href: `${nextPage.id}.html`,
            title: nextPage.title
        };
    } else if (currentSectionIndex < sections.length - 1) {
        const nextSection = sections[currentSectionIndex + 1];
        const nextPage = tourStructure[nextSection].pages[0];
        next = {
            href: `../${nextSection}/${nextPage.id}.html`,
            title: nextPage.title
        };
    }
    
    // Find previous page
    if (currentPageIndex > 0) {
        const prevPage = currentSectionData.pages[currentPageIndex - 1];
        prev = {
            href: `${prevPage.id}.html`,
            title: prevPage.title
        };
    } else if (currentSectionIndex > 0) {
        const prevSection = sections[currentSectionIndex - 1];
        const prevPages = tourStructure[prevSection].pages;
        const prevPage = prevPages[prevPages.length - 1];
        prev = {
            href: `../${prevSection}/${prevPage.id}.html`,
            title: prevPage.title
        };
    }
    
    return { prev, next };
}

// Setup mobile menu toggle
function setupMobileMenu() {
    const toggle = document.querySelector('.menu-toggle');
    const nav = document.querySelector('.nav-sidebar');
    
    if (toggle && nav) {
        toggle.addEventListener('click', () => {
            nav.classList.toggle('open');
        });
        
        // Close menu when clicking outside
        document.addEventListener('click', (e) => {
            if (!nav.contains(e.target) && !toggle.contains(e.target)) {
                nav.classList.remove('open');
            }
        });
    }
}

// Initialize when DOM is ready
if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', initTour);
} else {
    initTour();
}
