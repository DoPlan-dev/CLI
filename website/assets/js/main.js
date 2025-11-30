ooooo// Smooth scrolling for anchor links
document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();
        const target = document.querySelector(this.getAttribute('href'));
        if (target) {
            target.scrollIntoView({
                behavior: 'smooth',
                block: 'start'
            });
        }
    });
});

// Highlight active nav link on scroll
const sections = document.querySelectorAll('section[id]');
const navLinks = document.querySelectorAll('.nav-links a, .docs-nav a');

window.addEventListener('scroll', () => {
    let current = '';
    sections.forEach(section => {
        const sectionTop = section.offsetTop;
        const sectionHeight = section.clientHeight;
        if (window.pageYOffset >= sectionTop - 200) {
            current = section.getAttribute('id');
        }
    });

    navLinks.forEach(link => {
        link.classList.remove('active');
        if (link.getAttribute('href') === `#${current}`) {
            link.classList.add('active');
        }
    });
});

// Mobile menu toggle (if needed in future)
// Add mobile menu functionality here if needed

// Typing animation for terminal command with loop
const terminalCommand = document.getElementById('terminal-command');
const terminalCursor = document.getElementById('terminal-cursor');
const terminalPrompt = document.getElementById('terminal-prompt');
const copyButton = document.getElementById('copy-button');
const installLabel = document.getElementById('install-label');
const terminalWindow = document.getElementById('terminal-window');
const installNote = document.getElementById('install-note');
const pageCards = document.getElementById('page-cards');

const initCommand = 'npx @doplan-dev/cli init';
const initialPath = '~/work/ $';

let isLooping = true;
let loopCount = 0;
let terminalOpened = false;
let copyButtonShown = false;

function resetAnimation() {
    // Reset all elements (except terminal window)
    if (terminalCommand) terminalCommand.textContent = '';
    if (terminalCursor) terminalCursor.style.opacity = '1';
    if (terminalPrompt) terminalPrompt.textContent = initialPath;
    if (copyButton && !copyButtonShown) {
        copyButton.style.opacity = '0';
        copyButton.style.transform = 'scale(0.8)';
        copyButton.classList.remove('pulse-attention', 'copied');
        const icon = copyButton.querySelector('i');
        if (icon) icon.className = 'fas fa-copy';
    }
}

function typeText(text, callback) {
    if (!terminalCommand || !terminalCursor) return;
    
    let charIndex = 0;
    
    function type() {
        if (charIndex < text.length) {
            terminalCommand.textContent += text[charIndex];
            charIndex++;
            setTimeout(type, 120);
        } else {
            // Hide cursor when typing is complete
            terminalCursor.style.opacity = '0';
            if (callback) {
                setTimeout(callback, 500);
            }
        }
    }
    
    type();
}

function startDemoLoop() {
    if (!isLooping) return;
    
    // Step 1: Show "Install in seconds" label
    if (installLabel) {
        installLabel.classList.add('visible');
    }
    
    // Step 2: Show terminal window after label appears (only once, never close)
    setTimeout(() => {
        if (terminalWindow && !terminalOpened) {
            terminalWindow.classList.add('visible');
            terminalOpened = true;
            
            // Step 2.5: Show install note after terminal window fades in
            setTimeout(() => {
                if (installNote) {
                    installNote.classList.add('visible');
                }
            }, 500); // Wait for terminal fade-in to complete
        }
        
        // Step 3: Start typing sequence after terminal appears
        setTimeout(() => {
            startCommandSequence();
        }, 1500);
    }, 2000);
}

function startCommandSequence() {
    // Reset to initial state
    resetAnimation();
    
    setTimeout(() => {
        if (terminalCursor) terminalCursor.style.opacity = '1';
        typeText(initCommand, () => {
            setTimeout(() => {
                if (copyButton) {
                    copyButton.style.opacity = '1';
                    copyButton.style.transform = 'scale(1)';
                    copyButton.classList.add('pulse-attention');
                    copyButtonShown = true;
                }
                
                // Show cards after copy button appears (one by one)
                setTimeout(() => {
                    if (pageCards) {
                        pageCards.classList.add('visible');
                        
                        const cards = document.querySelectorAll('.category-card');
                        cards.forEach((card, index) => {
                            setTimeout(() => {
                                card.classList.add('visible');
                                
                                const commandItems = card.querySelectorAll('.command-item');
                                commandItems.forEach((item, itemIndex) => {
                                    setTimeout(() => {
                                        const cmd = item.querySelector('code');
                                        if (cmd) {
                                            cmd.classList.add('visible');
                                        }
                                    }, 200 + (itemIndex * 200));
                                });
                            }, index * 400);
                        });
                    }
                }, 300);
                
                // Loop typing animation while keeping copy button visible after first reveal
                loopCount += 1;
                if (loopCount < 2) {
                    setTimeout(() => {
                        startCommandSequence();
                    }, 5000);
                } else {
                    isLooping = false;
                }
            }, 500);
        });
    }, 500);
}

// Start the demo loop
setTimeout(() => {
    startDemoLoop();
}, 1000);

// Copy terminal command to clipboard
if (copyButton && terminalCommand) {
    copyButton.addEventListener('click', async () => {
        const commandText = 'npx @doplan-dev/cli init';
        
        try {
            await navigator.clipboard.writeText(commandText);
            
            // Update button state
            copyButton.classList.remove('pulse-attention');
            copyButton.classList.add('copied');
            const icon = copyButton.querySelector('i');
            if (icon) {
                icon.className = 'fas fa-check';
            }
            
            // Reset button after 2 seconds
            setTimeout(() => {
                copyButton.classList.remove('copied');
                if (icon) {
                    icon.className = 'fas fa-copy';
                }
            }, 2000);
        } catch (err) {
            console.error('Failed to copy command:', err);
            // Fallback for older browsers
            const textArea = document.createElement('textarea');
            textArea.value = commandText;
            textArea.style.position = 'fixed';
            textArea.style.opacity = '0';
            document.body.appendChild(textArea);
            textArea.select();
            try {
                document.execCommand('copy');
                copyButton.classList.remove('pulse-attention');
                copyButton.classList.add('copied');
                setTimeout(() => {
                    copyButton.classList.remove('copied');
                    const fallbackIcon = copyButton.querySelector('i');
                    if (fallbackIcon) fallbackIcon.className = 'fas fa-copy';
                }, 2000);
            } catch (fallbackErr) {
                console.error('Fallback copy failed:', fallbackErr);
            }
            document.body.removeChild(textArea);
        }
    });
}

