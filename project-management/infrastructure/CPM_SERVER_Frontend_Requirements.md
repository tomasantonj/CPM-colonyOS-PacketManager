# CPM Frontend Hosting Specs

**Server Purpose:**
This infrastructure hosts the **CPM Web Dashboard** (React / Nginx).
*   **Role:** Visual User Interface.
*   **Function:** Serves static assets (`.js`, `.css`) to the browser.
*   **Why we need it:** To provide a drag-and-drop template builder and package browser for users who prefer UI over CLI.

**Component:** CPM Content Delivery (Nginx)
**Type:** Stateless Container

## 1. Container Specs
*   **Image:** `cpm-dashboard:latest`
*   **Base:** `nginx:alpine`
*   **Port:** `80` (Internal)
*   **Resources:** 0.25 vCPU, 128MB RAM.

## 2. Configuration

### Environment Variables
Since this is a static SPA, these must be injected at runtime (or build time).
*   `REACT_APP_REGISTRY_API_URL`: Registry API Public Endpoint.
*   `REACT_APP_COLONYOS_URL`: ColonyOS Server Public Endpoint.

### Nginx Config
Must handle SPA routing (History Mode) and caching.

```nginx
server {
    listen 80;
    location / {
        try_files $uri $uri/ /index.html;
    }
    location /static {
        expires 1y;
        add_header Cache-Control "public";
    }
}
```

## 3. Network & Security
*   **Ingress:** HTTPS (443) required (SubtleCrypto API dependency).
*   **CORS:** Registry & ColonyOS must accept this origin.
