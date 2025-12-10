# CPM Registry API Requirements

**Target:** Standalone Registry Server
**Tech Stack:** C# / .NET 8 (ASP.NET Core)
**Goal:** Provide a robust, high-performance HTTP API for publishing, searching, and managing CPM packages.

## 1. Architecture & Tech Stack

Using C# and .NET 8 is an excellent choice for the API due to its strong typing, performance, and robust tooling.

*   **Framework:** ASP.NET Core Web API (.NET 8).
*   **Language:** C# 12.
*   **Database ORM:** Entity Framework Core (Code-first approach).
*   **Database:** PostgreSQL (Can share the same instance as ColonyOS but use a generic `cpm_registry` database).
*   **Serialization:** System.Text.Json.
*   **Documentation:** Swagger/OpenAPI (Built-in).

## 2. Core Responsibilities

The API will serve two main clients:
1.  **CPM CLI:** For publishing (`cpm publish`) and searching.
2.  **Web Dashboard:** For the visual package browser and authentication.

## 3. API Endpoints

### A. Authentication (JWT)
*   `POST /api/auth/register`: Create a new user/maintainer.
*   `POST /api/auth/login`: Exchange credentials for a Bearer Token (JWT).
*   `GET /api/auth/me`: Get current user profile.

### B. Packages
*   `POST /api/packages`: Publish a new package version.
    *   *Payload:* Multipart form data containing the `tar.gz` archive and the `signature`.
    *   *Validation:* Server *must* verify the Ed25519 signature against the public key in the payload/manifest.
*   `GET /api/packages`: List packages (Pagination + Search support).
    *   *Query Params:* `search`, `tags`, `page`, `pageSize`.
*   `GET /api/packages/{name}`: Get package details (aggregated versions).
*   `GET /api/packages/{name}/{version}`: Get specific version details.
*   `GET /api/packages/{name}/{version}/download`: Download the `tar.gz` artifact.

## 4. Storage Handling

The API needs to store the actual package binaries (`.tar.gz` files).

*   **Abstraction:** `IBlobStorageService`.
*   **Implementations:**
    *   **FileSystem (Dev):** `App_Data/packages/`.
    *   **S3 Compatible (Prod):** AWS S3, MinIO, or DigitalOcean Spaces.
    *   *Note:* Do not store binaries in the Database. Store them in object storage and save the URL/Path in Postgres.

## 5. Hosting & Deployment

This service should run as its own Docker container, separate from the ColonyOS core.

### Dockerfile Requirements
*   **Base Image:** `mcr.microsoft.com/dotnet/aspnet:8.0` (Alpine/Debian).
*   **Build SDK:** `mcr.microsoft.com/dotnet/sdk:8.0`.
*   **Port:** Expose `80` (Internal). Map to `5000` or handle via Nginx proxy.

### Configuration (`appsettings.json`)
```json
{
  "ConnectionStrings": {
    "DefaultConnection": "Host=postgres;Database=cpm_registry;Username=postgres;Password=..."
  },
  "Storage": {
    "Type": "S3",
    "BucketName": "cpm-packages",
    "AccessKey": "...",
    "SecretKey": "..."
  },
  "Jwt": {
    "Key": "super-secret-key-...",
    "Issuer": "cpm-registry"
  }
}
```

## 6. Integration Checklist
*   **CORS:** Must act as the bridge. Enable CORS for the Web Dashboard URL.
*   **Health Checks:** Expose `/health` for Kubernetes/Docker probes.
