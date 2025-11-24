// Mobile menu toggle
const mobileMenuToggle = document.querySelector(".mobile-menu-toggle");
const navLinks = document.querySelector(".nav-links");

if (mobileMenuToggle) {
    mobileMenuToggle.addEventListener("click", () => {
        navLinks.classList.toggle("active");
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

    if (currentScroll > 100) {
        navbar.style.boxShadow = "0 2px 10px rgba(0, 0, 0, 0.1)";
    } else {
        navbar.style.boxShadow = "none";
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
        let code = block.textContent;

        // Highlight comments
        code = code.replace(/\/\/.*$/gm, '<span class="comment">$&</span>');

        // Highlight keywords
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
        ];

        keywords.forEach((keyword) => {
            const regex = new RegExp(`\\b${keyword}\\b`, "g");
            code = code.replace(
                regex,
                `<span style="color: #f59e0b; font-weight: 600;">${keyword}</span>`
            );
        });

        // Highlight strings
        code = code.replace(
            /"([^"]*)"/g,
            '<span style="color: #10b981;">"$1"</span>'
        );

        // Highlight numbers
        code = code.replace(
            /\b(\d+)\b/g,
            '<span style="color: #3b82f6;">$1</span>'
        );

        block.innerHTML = code;
    });
}

// Run highlighting after page load
document.addEventListener("DOMContentLoaded", highlightCode);

// Re-highlight when tabs change
tabButtons.forEach((button) => {
    button.addEventListener("click", () => {
        setTimeout(highlightCode, 100);
    });
});
