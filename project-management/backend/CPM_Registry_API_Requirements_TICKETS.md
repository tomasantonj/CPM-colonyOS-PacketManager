# CPM Registry API Development Tickets
Based on: [CPM_Registry_API_Requirements.md](./CPM_Registry_API_Requirements.md)

## 1. Project Scaffolding

### [REGISTRY-01] Initialize .NET 8 Web API
**Priority:** High
**Description:**
Create the initial solution and project structure for the Registry API.
**Technical Details:**
- Use .NET 8 ASP.NET Core Web API template.
- Configure Entity Framework Core with Postgres.
- Setup Swagger/OpenAPI generation.
**Acceptance Criteria:**
- Solution builds.
- `/swagger` endpoint is accessible.
- Dockerfile is created and runs.

## 2. Authentication

### [REGISTRY-02] Implement User Authentication (JWT)
**Priority:** High
**Description:**
Implement endpoints for registration, login, and token generation.
**Technical Details:**
- `POST /api/auth/register`: Hashes password (Argon2/BCrypt) and saves user.
- `POST /api/auth/login`: Validates password and returns JWT.
- `GET /api/auth/me`: Returns user profile from claims.
**Acceptance Criteria:**
- User can register and login.
- JWT contains correct claims (UserId, Role).

### [REGISTRY-03] Implement API Key Support
**Priority:** Medium (CI/CD)
**Description:**
Allow generating long-lived API keys for automated pipelines.
**Technical Details:**
- `POST /api/auth/token`: Generates a distinct API Key entity linked to User.
- Middleware to accept both `Bearer <JWT>` and `X-API-Key <Key>`.
**Acceptance Criteria:**
- Endpoints accept API Key authentication.

## 3. Package Management

### [REGISTRY-04] Implement Package Publishing Endpoint
**Priority:** Critical
**Description:**
Handle multipart uploads of package artifacts (`.tar.gz`) and metadata.
**Technical Details:**
- `POST /api/packages`: Accepts file + json metadata.
- **CRITICAL:** Save file to Blob Storage (via `IBlobStorageService`).
- Save metadata to Postgres.
**Acceptance Criteria:**
- Uploaded file appears in S3/Disk.
- Database record is created.

### [REGISTRY-05] Implement Package Listing & Search
**Priority:** High
**Description:**
Allow CLI and Dashboard to search for packages.
**Technical Details:**
- `GET /api/packages`: Implement pagination (`page`, `pageSize`) and search (`q` filtering name/tags/description).
- Return standard JSON response.
**Acceptance Criteria:**
- Searching for "workflow" returns relevant packages.
- Pagination works correctly.

### [REGISTRY-06] Implement Package Yanking (Delete)
**Priority:** Low
**Description:**
Allow authors to deprecate/remove specific versions.
**Technical Details:**
- `DELETE /api/packages/{name}/{version}`.
- Soft delete (mark as `Yanked = true` in DB).
- Remove from search results but keep file accessible for existing downloads.
**Acceptance Criteria:**
- Yanked package does not appear in search.

## 4. Storage & Security

### [REGISTRY-07] Implement Blob Storage Abstraction
**Priority:** High
**Description:**
Create `IBlobStorageService` to decouple logic from AWS S3.
**Technical Details:**
- Implement `FileSystemStorage` (for local dev).
- Implement `S3Storage` (for production).
- Configurable via `appsettings.json`.
**Acceptance Criteria:**
- Switching config allows saving to local disk or S3 without code changes.

### [REGISTRY-08] Enforce Signature Verification
**Priority:** Critical
**Description:**
The server MUST verify the Ed25519 signature before accepting a package.
**Technical Details:**
- On `POST /packages`, read the `signature` field.
- Verify signature against the uploaded file content and provided public key.
- Reject upload if verification fails (HTTP 400).
**Acceptance Criteria:**
- Uploading a tampered file results in error.
