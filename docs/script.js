// Mobile menu toggle
const mobileMenuToggle = document.querySelector(".mobile-menu-toggle");
const navLinks = document.querySelector(".nav-links");
const bodyElement = document.body;

const closeMobileMenu = () => {
    if (navLinks && navLinks.classList.contains("active")) {
        navLinks.classList.remove("active");
    }
    bodyElement.classList.remove("menu-open");
    if (mobileMenuToggle) {
        mobileMenuToggle.setAttribute("aria-expanded", "false");
    }
};

if (mobileMenuToggle && navLinks) {
    mobileMenuToggle.setAttribute("aria-expanded", "false");
    mobileMenuToggle.addEventListener("click", () => {
        const isActive = navLinks.classList.toggle("active");
        bodyElement.classList.toggle("menu-open", isActive);
        mobileMenuToggle.setAttribute(
            "aria-expanded",
            isActive ? "true" : "false"
        );
    });

    navLinks.querySelectorAll("a").forEach((link) => {
        link.addEventListener("click", () => {
            closeMobileMenu();
        });
    });

    window.addEventListener("resize", () => {
        if (window.innerWidth > 768) {
            closeMobileMenu();
        }
    });

    document.addEventListener("keydown", (event) => {
        if (event.key === "Escape") {
            closeMobileMenu();
        }
    });
}

// Tab switching
const tabButtons = document.querySelectorAll(".tab-btn");
const codeExamples = document.querySelectorAll(".code-example");

tabButtons.forEach((button) => {
    button.addEventListener("click", () => {
        const tabName = button.getAttribute("data-tab");

        // Remove active class from all tabs and examples
        tabButtons.forEach((btn) => btn.classList.remove("active"));
        codeExamples.forEach((example) => example.classList.remove("active"));

        // Add active class to clicked tab and corresponding example
        button.classList.add("active");
        const targetExample = document.getElementById(`${tabName}-example`);
        if (targetExample) {
            targetExample.classList.add("active");
        }
    });
});

// Copy code functionality
const copyButtons = document.querySelectorAll(".copy-btn");

copyButtons.forEach((button) => {
    button.addEventListener("click", async () => {
        const codeType = button.getAttribute("data-code");
        const example = document.getElementById(`${codeType}-example`);
        const codeBlock = example.querySelector("code");
        const codeText = codeBlock.textContent;

        try {
            await navigator.clipboard.writeText(codeText);
            button.textContent = "Copied!";
            button.classList.add("copied");

            setTimeout(() => {
                button.textContent = "Copy";
                button.classList.remove("copied");
            }, 2000);
        } catch (err) {
            console.error("Failed to copy:", err);
            button.textContent = "Failed";
            setTimeout(() => {
                button.textContent = "Copy";
            }, 2000);
        }
    });
});

// Smooth scroll for anchor links
document.querySelectorAll('a[href^="#"]').forEach((anchor) => {
    anchor.addEventListener("click", function (e) {
        e.preventDefault();
        const target = document.querySelector(this.getAttribute("href"));
        if (target) {
            target.scrollIntoView({
                behavior: "smooth",
                block: "start",
            });
        }
    });
});

// Navbar scroll effect
let lastScroll = 0;
const navbar = document.querySelector(".navbar");

window.addEventListener("scroll", () => {
    const currentScroll = window.pageYOffset;

    if (navbar) {
        if (currentScroll > 100) {
            navbar.style.boxShadow = "0 2px 10px rgba(0, 0, 0, 0.1)";
        } else {
            navbar.style.boxShadow = "none";
        }
    }

    if (navLinks && navLinks.classList.contains("active")) {
        closeMobileMenu();
    }

    lastScroll = currentScroll;
});

// Intersection Observer for fade-in animations
const observerOptions = {
    threshold: 0.1,
    rootMargin: "0px 0px -50px 0px",
};

const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
        if (entry.isIntersecting) {
            entry.target.style.opacity = "1";
            entry.target.style.transform = "translateY(0)";
        }
    });
}, observerOptions);

// Observe feature cards and keyword cards
document.querySelectorAll(".feature-card, .keyword-card").forEach((card) => {
    card.style.opacity = "0";
    card.style.transform = "translateY(20px)";
    card.style.transition = "opacity 0.6s ease, transform 0.6s ease";
    observer.observe(card);
});

// Add syntax highlighting for code blocks (simple version)
function highlightCode() {
    const codeBlocks = document.querySelectorAll("code.language-tau");

    codeBlocks.forEach((block) => {
        // Skip if already highlighted
        if (block.dataset.highlighted === "true") {
            return;
        }

        // Get original text content (before any HTML)
        // If already has HTML, extract text content
        let code = block.textContent || block.innerText;

        // Escape HTML first
        const escapeHtml = (text) => {
            const div = document.createElement("div");
            div.textContent = text;
            return div.innerHTML;
        };

        code = escapeHtml(code);

        // Process in correct order to avoid conflicts
        // 1. First, protect strings by replacing them with placeholders
        const stringPlaceholders = [];
        code = code.replace(/"([^"]*)"/g, (match, content) => {
            const placeholder = `__STRING_${stringPlaceholders.length}__`;
            stringPlaceholders.push(
                `<span class="code-string">"${escapeHtml(content)}"</span>`
            );
            return placeholder;
        });

        // 2. Highlight comments
        code = code.replace(/\/\/.*$/gm, (match) => {
            return `<span class="code-comment">${match}</span>`;
        });

        // 3. Highlight keywords
        const keywords = [
            "sun_liyo_tau",
            "tau_ka_jugaad",
            "agar_maan_lo",
            "na_toh",
            "laadle_ye_le",
            "jab_tak",
            "rok_diye",
            "jaan_de",
            "ne_bana_diye",
            "saccha",
            "jhootha",
            "print",
        ];

        keywords.forEach((keyword) => {
            // Use word boundaries, but avoid matching inside placeholders or HTML tags
            const regex = new RegExp(
                `\\b${keyword}\\b(?![^_]*__)(?![^<]*>)`,
                "g"
            );
            code = code.replace(
                regex,
                `<span class="code-keyword">${keyword}</span>`
            );
        });

        // 4. Highlight numbers (avoid matching inside placeholders or HTML tags)
        code = code.replace(
            /\b(\d+)\b(?![^_]*__)(?![^<]*>)/g,
            '<span class="code-number">$1</span>'
        );

        // 5. Restore strings
        stringPlaceholders.forEach((placeholder, index) => {
            code = code.replace(`__STRING_${index}__`, placeholder);
        });

        // Now apply our highlighting
        block.innerHTML = code;
        block.dataset.highlighted = "true";
    });
}

// Dark mode functionality
function initDarkMode() {
    const darkModeToggle = document.getElementById("dark-mode-toggle");
    const prefersDark = window.matchMedia(
        "(prefers-color-scheme: dark)"
    ).matches;
    const savedTheme = localStorage.getItem("theme");

    // Determine initial theme
    let isDark = savedTheme === "dark" || (!savedTheme && prefersDark);

    function applyTheme(dark) {
        document.documentElement.classList.toggle("dark-mode", dark);
        if (darkModeToggle) {
            darkModeToggle.textContent = dark ? "☀️" : "🌙";
            darkModeToggle.setAttribute(
                "aria-label",
                dark ? "Switch to light mode" : "Switch to dark mode"
            );
        }
        localStorage.setItem("theme", dark ? "dark" : "light");
    }

    // Apply initial theme
    applyTheme(isDark);

    // Toggle on button click
    if (darkModeToggle) {
        darkModeToggle.addEventListener("click", () => {
            isDark = !isDark;
            applyTheme(isDark);
        });
    }

    // Listen for system theme changes
    window
        .matchMedia("(prefers-color-scheme: dark)")
        .addEventListener("change", (e) => {
            if (!localStorage.getItem("theme")) {
                applyTheme(e.matches);
            }
        });
}

// Run highlighting after page load
document.addEventListener("DOMContentLoaded", () => {
    setTimeout(highlightCode, 100);
    initDarkMode();
});

// Re-highlight when tabs change (reset first)
tabButtons.forEach((button) => {
    button.addEventListener("click", () => {
        // Reset highlighted state for code blocks in the new tab
        setTimeout(() => {
            const activeExample = document.querySelector(
                ".code-example.active"
            );
            if (activeExample) {
                const codeBlock =
                    activeExample.querySelector("code.language-tau");
                if (codeBlock && codeBlock.dataset.highlighted === "true") {
                    // Get the original text before highlighting
                    const originalText = codeBlock.textContent;
                    codeBlock.dataset.highlighted = "false";
                    codeBlock.textContent = originalText;
                    highlightCode();
                }
            }
        }, 150);
    });
});
