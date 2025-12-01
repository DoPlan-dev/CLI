// Theme Toggle Functionality
(function() {
    const themeToggle = document.getElementById('theme-toggle');
    const html = document.documentElement;

    // Get saved theme or default to dark
    const currentTheme = localStorage.getItem('theme') || 'dark';
    
    // Apply saved theme
    if (currentTheme === 'light') {
        html.setAttribute('data-theme', 'light');
    } else {
        html.removeAttribute('data-theme');
    }

    // Theme toggle handler
    themeToggle.addEventListener('click', () => {
        const currentTheme = html.getAttribute('data-theme');
        
        if (currentTheme === 'light') {
            // Switch to dark
            html.removeAttribute('data-theme');
            localStorage.setItem('theme', 'dark');
        } else {
            // Switch to light
            html.setAttribute('data-theme', 'light');
            localStorage.setItem('theme', 'light');
        }
    });
})();

