# CPM Registry Hosting Specs

**Server Purpose:**
This infrastructure hosts the **CPM Registry API** (ASP.NET Core / C#).
*   **Role:** Central Hub & Remote Backend.
*   **Function:** Handles package publishing (`POST`), searching (`GET`), and metadata storage.
*   **Why we need it:** To allow users to share packages remotely and authenticate securely (JWT).

## 1. Container Specs
*   **Image:** `cpm-registry:latest`
*   **Base:** `mcr.microsoft.com/dotnet/aspnet:8.0`
*   **Port:** `80` (Internal) mapped via Proxy.
*   **Resources:** 1 vCPU, 512MB RAM.

## 2. Configuration (`appsettings.json`)
Inject via Environment Variables or mount JSON:
*   `ConnectionStrings__DefaultConnection`: Postgres connection string.
*   `Storage__AccessKey/SecretKey`: S3/Blob storage credentials.
*   `Jwt__Key`: Secret key for token signing.

## 3. Services Dependencies
*   **PostgreSQL:** Stores user accounts and package metadata.
*   **Blob Storage (S3/MinIO):** Stores the actual `.tar.gz` package files.

## 4. Network & Security
*   **Ingress:** HTTPS (443).
*   **CORS:** Must allow requests from the **Frontend Dashboard** domain.
