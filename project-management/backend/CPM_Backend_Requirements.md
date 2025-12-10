# CPM Backend Requirements & Upgrade Plan

### Architecture Context
**Component:** CPM CLI / Go Core
**Role:** Client & Local Logic
This document specifies changes required in the **existing Go codebase**. It focuses on data structures (`pkg/domain`) and CLI commands (`cpm login`) needed to support the wider ecosystem.

**Target:** Go Backend (CLI & Registry Server)
**Version:** v1.0 Preparation

This document outlines the required upgrades to the core Go codebase to support the Release Roadmap and the new Web Dashboard.

## 1. Core Data Model Upgrades
To support the Visual Template Builder in the Web UI, the manifest structure must be strictly typed.

### A. Manifest Schema Expansion
**File:** `pkg/domain/manifest.go`

Current state:
```go
type ColonyManifest struct {
    // ...
    // No explicit inputs definition
}
```

**Requirement:**
Add an `Inputs` definition to allow the frontend to generate dynamic forms.

```go
type ColonyManifest struct {
    // ...
    Inputs []InputVariable `yaml:"inputs,omitempty"`
}

type InputVariable struct {
    Key         string   `yaml:"key"`                   // e.g., "gpu_count"
    Label       string   `yaml:"label"`                 // e.g., "GPU Allocation"
    Type        string   `yaml:"type"`                  // "string", "int", "bool", "select"
    Description string   `yaml:"description,omitempty"`
    Required    bool     `yaml:"required"`
    Default     any      `yaml:"default,omitempty"`     // Generic default value
    Options     []string `yaml:"options,omitempty"`     // For "select" type
}
```

## 2. Remote Registry Server
The "Phase 2" roadmap requires moving from a local-only architecture to a distributed one.

### A. API Endpoints (REST/JSON)
A new server component (or mode `cpm serve`) is needed to handle these requests.

*   `POST /api/v1/auth/login`: Authenticate User (Issue JWT).
*   `GET /api/v1/packages`: Search/List packages (Query params: `q`, `tag`, `author`).
*   `GET /api/v1/packages/:name`: Get package metadata (including new `Inputs` schema).
*   `POST /api/v1/packages`: Publish a new package (Requires JWT & Valid Signature).

### B. Storage Interface
*   **Abstract logic:** Define a `Repository` interface for storing packages.
*   **Implementations:**
    *   `FSRepository` (Local filesystem - current).
    *   `S3Repository` (AWS S3 / MinIO - for production registry).
    *   `PostgresRepository` (Metadata indexing).

## 3. CLI Enhancements

### A. Authentication
*   **Command:** `cpm login <registry-url>`
*   **Logic:** Authenticate with the remote registry and store the returned JWT in `~/.cpm/config`.

### B. Validation
*   **Command:** `cpm validate`
*   **Logic:** Validate `colony.yaml` against the new `InputVariable` schema. Ensure provided `values.yaml` matches the defined inputs.

## 4. Integration & Security

### A. Signing
*   Ensure the Ed25519 signing mechanism validates not just the workflow content, but the *integrity* of the package being uploaded to the registry.

### B. CORS
*   The Registry Server must enable CORS to allow requests from the Web Dashboard (running on a different domain/port).
