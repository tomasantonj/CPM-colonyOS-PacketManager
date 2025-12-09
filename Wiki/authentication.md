# Authentication & Security

CPM secures communication with the ColonyOS backend using **Ed25519** digital signatures. This ensures that workflows submitted are authentic and have not been tampered with.

## Credentials

To interact with a secured ColonyOS instance, you need two pieces of information:

1.  **Colony ID**: The unique identifier of the Colony you are submitting to.
2.  **Private Key**: A 64-byte, hex-encoded Ed25519 private key authorized to submit workflows to that Colony.

These are passed to the CLI via flags:
```bash
cpm install ... --colonyid <ID> --prvkey <KEY>
```

## Protocol Details

When `cpm install` runs, it performs the following steps to authenticate the request:

### 1. Payload Creation
The rendered templates are combined into a JSON payload.
```json
[{ "name": "workflow-1", ... }]
```

### 2. Signing
The payload bytes are signed using the provided **Private Key**. This generates a cryptographic signature.

### 3. HTTP Request
CPM sends an HTTP `POST` request to `/api/workflows` with the following headers:

-   `Content-Type`: `application/json`
-   `X-Colony-ID`: The Colony ID.
-   `X-Colony-Signature`: The hex-encoded signature generated in step 2.

### 4. Verification (Server-Side)
ColonyOS receives the request, retrieves the *Public Key* associated with the Colony, and verifies that the signature matches the payload and the key. If verification fails, the request is rejected.

## CLI Usage

When installing a package that requires submission to a real Colony:

**Bash:**
```bash
cpm install my-package \
  --colonyid "af67-..." \
  --prvkey "e5b9..." \
  --host "colony.example.com" \
  --port 5432
```

**PowerShell:**
```powershell
cpm install my-package `
  --colonyid "af67-..." `
  --prvkey "e5b9..." `
  --host "colony.example.com" `
  --port 5432
```

-   `--colonyid`: The ID of the target Colony.
-   `--prvkey`: The 64-byte hex-encoded private key for signing.
-   `--host`: Hostname of the ColonyOS server (default: `localhost`).
-   `--port`: Port of the ColonyOS server (default: `50080`).
