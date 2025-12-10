# CPM Web Dashboard Requirements

This document outlines the requirements for the **CPM Web Dashboard**, a visual interface for managing packages and building `colony.yaml` templates.

## Functional Requirements

### 1. Template Builder
*   **Visual Editor:** A drag-and-drop or form-based interface to generate valid `colony.yaml` files.
*   **Validation:** Real-time validation of fields (e.g., checking if required variables like `${{gpu}}` are defined).
*   **Export:** Ability to download the generated YAML or trigger a `cpm pack` (via API).

### 2. Registry Integration
*   **Search & Browse:** Interface to search the generic Registry API for published packages.
*   **Package Details:** View `README.md`, version history, and input variables for specific packages.

### 3. ColonyOS Interaction (Optional / Phase 2)
*   **Deployment:** Trigger a workflow submission directly from the UI (requires user private key input or delegation).

## Technical Architecture

### 1. Tech Stack
*   **Frontend:** React 18+ (using Vite).
*   **UI Library:** Tailwind CSS (for styling) + Shadcn/UI (for components).
*   **State Management:** React Query (for API caching).
*   **Visuals:** React Flow (for node-based workflow visualization).

### 2. Docker containerization
The application will be packaged as a stateless Docker container.

*   **Base Image:** `nginx:alpine` (serving a static production build) or `node:18-alpine` (if SSR is needed).
*   **Port:** Exposes port `80` (internal) mapped to host port (e.g., `3000`).
*   **Environment Variables:**
    *   `REACT_APP_REGISTRY_URL`: URL of the CPM Registry API.
    *   `REACT_APP_COLONYOS_URL`: URL of the ColonyOS Server (grpc-web or proxy).

## Hosting / Server Requirements

Since this runs as a Docker container, it can co-exist on the main ColonyOS server or run on a separate instance.

### 1. Compute Resources (Container)
*   **CPU:** 0.5 vCPU (Low usage, mostly static asset serving).
*   **RAM:** 512MB (Minimum) - 1GB (Recommended).

### 2. Host System
*   **Container Runtime:** Docker Engine 20.10+ or Podman.
*   **Network:**
    *   **Ingress:** Reverse proxy (Nginx/Traefik) handling SSL termination.
    *   **Connectivity:** Must be able to reach the *Registry API* endpoint.

### 3. Deployment Command (Example)
```bash
docker run -d \
  --name cpm-dashboard \
  -p 3000:80 \
  -e REACT_APP_REGISTRY_URL="https://registry.example.com" \
  cpm-dashboard:latest
```
