# CPM Backend Development Tickets
Based on: [CPM_Backend_Requirements.md](./CPM_Backend_Requirements.md)

## 1. Domain & Data Model

### [BACKEND-01] Implement `Inputs` in Manifest Schema
**Priority:** High

**Description:**
Update `pkg/domain/manifest.go` to support dynamic input definitions. This is required for the Frontend Template Builder.

**Technical Details:**
- Add `Inputs []InputVariable` struct with `Key`, `Label`, `Type`, `Required`, `Default`.
- Update YAML parser to handle this new field.

**Acceptance Criteria:**
- `colony.yaml` with an `inputs` section parses correctly.
- Unit tests verify struct mapping.

## 2. CLI Enhancements

### [BACKEND-02] Implement `cpm login` Command
**Priority:** High

**Description:**
Add authentication support to the CLI to interact with the Remote Registry.

**Technical Details:**
- Implement `cpm login <registry_url>`.
- Prompt for username/password (or token).
- Store resulting JWT in `~/.cpm/config` (secure file storage).

**Acceptance Criteria:**
- Command saves the key/token to disk.
- Subsequent commands can read this token.

### [BACKEND-03] Implement `cpm publish --token` Support
**Priority:** Medium (CI/CD)

**Description:**
Allow non-interactive authentication for CI/CD pipelines.

**Technical Details:**
- Check for `--token` flag or `CPM_API_TOKEN` env var.
- Bypass `~/.cpm/config` lookup if token is provided directly.

**Acceptance Criteria:**
- `cpm publish` works without a config file if env var is set.

### [BACKEND-04] Implement `cpm yank` Command
**Priority:** Low

**Description:**
Allow developers to deprecate broken versions.

**Technical Details:**
- CALL `DELETE /api/packages/{name}/{version}` on the Registry API.
- Use the stored JWT for auth.

**Acceptance Criteria:**
- Successfully calls the endpoint.
- Handles 401/403/404 errors gracefully.

### [BACKEND-05] Implement `cpm validate` Command
**Priority:** Medium

**Description:**
Verify that a `values.yaml` file matches the `inputs` defined in `colony.yaml`.

**Technical Details:**
- Check required fields are present.
- Check types match (int vs string).

**Acceptance Criteria:**
- Returns exit code 1 if validation fails.
- Prints specific error messages (e.g., "Missing required input: gpu_count").

## 3. Registry & Integration

### [BACKEND-06] Update Signing Logic for Registry
**Priority:** Critical

**Description:**
Ensure `cpm publish` signs the entire artifact (checksum), not just the `workflow.json`.

**Technical Details:**
- Generate SHA256 of the `.tar.gz`.
- Sign the hash with Ed25519.
- Include signature in the POST request header/body.

**Acceptance Criteria:**
- Registry can verify the upload integrity.
