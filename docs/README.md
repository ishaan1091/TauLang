# TauLang Documentation Site

This directory contains the static site for TauLang, which is automatically deployed to GitHub Pages.

## Local Development

To view the site locally:

1. Open `index.html` in a web browser, or
2. Use a local server:

    ```bash
    # Using Python
    python -m http.server 8000

    # Using Node.js (if you have http-server installed)
    npx http-server -p 8000
    ```

3. Navigate to `http://localhost:8000`

## Deployment

The site is automatically deployed to GitHub Pages when changes are pushed to the `main` branch in the `docs/` directory.

### Setting up GitHub Pages

1. Go to your repository settings on GitHub
2. Navigate to "Pages" in the left sidebar
3. Under "Source", select "GitHub Actions"
4. The workflow will automatically deploy when you push changes

### Manual Deployment

If you need to manually trigger a deployment:

1. Go to the "Actions" tab in your repository
2. Select "Deploy to GitHub Pages" workflow
3. Click "Run workflow"

## File Structure

```
docs/
├── index.html      # Main landing page
├── styles.css      # Stylesheet
├── script.js       # JavaScript for interactivity
├── .nojekyll       # Prevents Jekyll processing
└── README.md       # This file
```

## Customization

To customize the site:

-   **Colors**: Edit CSS variables in `styles.css` (`:root` section)
-   **Content**: Edit `index.html`
-   **Functionality**: Edit `script.js`

Make sure to test locally before pushing changes!
