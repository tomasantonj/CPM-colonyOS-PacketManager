# CPM Release Roadmap

This document outlines the strategic plan to move Colony Package Manager (CPM) from its current Beta state to a public v1.0 release.

## Phase 1: Stabilization & Core Reliability
**Goal:** Ensure the current feature set is bulletproof and production-ready.

- [ ] **Infrastructure Setup (Prerequisite)**
    - Deploy a live ColonyOS server for true integration testing.
    - **Requirements:**
        - **Compute:** Linux VM/Container (2 vCPU, 4GB RAM min).
        - **Database:** PostgreSQL 13+ (Persistent storage).
        - **Broker:** Redis 6+ (Queue management).
        - **Security:** TLS Certificates (Let's Encrypt/Self-signed), Ed25519 Root Keypair.
        - **Network:** Ingress for gRPC (Default 50051) and HTTP (Default 8080).
- [ ] **Comprehensive Testing Suite**
    - Implement unit tests for all `internal/usecase` logic.
    - Create integration tests for the full `pack` -> `publish` -> `install` cycle against a live ColonyOS server.
    - Add "chaos testing" for network failures and invalid configs.
- [ ] **Error Handling & UX Polish**
    - Replace generic Go errors with human-readable error messages and troubleshooting codes.
    - Add progress bars and spinners for long-running CLI operations (packaging, uploading).
    - Implement a `cpm doctor` command to diagnose environment issues (keys, connectivity).
- [ ] **Security Auditing**
    - Verify Ed25519 signing implementation.
    - Audit file permission handling during installation (ensure least privilege).
    - Validate input sanitization in template rendering to prevent injection attacks.

## Phase 2: Registry & Ecosystem
**Goal:** Expand from local/direct usage to a scalable, distributed ecosystem.

- [ ] **Remote Registry Server**
    - Develop a standalone HTTP registry server (or a ColonyOS plugin) to host packages.
    - Implement `cpm login` and `cpm logout` for registry authentication.
- [ ] **Advanced Registry Features**
    - `cpm search` with filters (tags, authors).
    - `cpm info` to view package metadata without downloading.
    - Support for semantic versioning (install `@latest`, `@v1.2`).
- [ ] **Developer Experience (DX)**
    - Create a VS Code extension for CPM (syntax highlighting for `colony.yaml`, snippets).
    - Add a "scaffold" feature to `cpm init` with multiple starter templates (e.g., Python, Node.js, Go workers).
- [ ] **CPM Web Dashboard (React)**
    - Develop a visual interface for creating `colony.yaml` templates (drag-and-drop workflow builder).
    - Implement an API gateway (or use Registry API) to fetch and publish templates.
    - **Note:** Depends on the *Remote Registry Server* for API access.

## Phase 3: Public Beta Launch
**Goal:** Prepare for public consumption and community uptake.

- [ ] **Documentation Overhaul**
    - Launch a dedicated documentation site (e.g., using Docusaurus or MkDocs).
    - Create video tutorials: "Zero to deployed in 5 minutes".
    - Write a "Migration Guide" for users transitioning from raw JSON/YAML dispatching.
- [ ] **CI/CD Integration**
    - Release official GitHub Actions and GitLab CI components for installing CPM and deploying packages.
    - Automate CPM binary releases (cross-platform builds for Linux, Mac, Windows).
- [ ] **Community Channels**
    - Set up a Discord or Slack community.
    - Create a "Package Index" where the community can list their public packages.
    - Define "Official" vs "Community" package tiers.
