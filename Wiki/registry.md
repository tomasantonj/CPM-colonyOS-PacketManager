# Registry

The CPM Registry is the storage backend where packages are published and retrieved from.

## Storage Location
By default, CPM uses a **local filesystem registry**.

- **Default Path**: `~/.cpm/registry` (on Linux/macOS) or `%USERPROFILE%\.cpm\registry` (on Windows).
- **Custom Path**: You can override the root location by setting the `CPM_HOME` environment variable. The registry will be at `$CPM_HOME/registry`.

## Internal Layout
The registry organizes packages by name and version to support fast lookups.

```text
~/.cpm/registry/
├── my-package/
│   ├── 0.1.0.cpm
│   ├── 0.1.1.cpm
│   └── 0.2.0.cpm
└── postgres/
    └── 14.1.0.cpm
```

## CLI Usage

Here are the primary commands for interacting with the registry:

### Publishing
Pack a directory and upload it to the registry.

**Bash:**
```bash
cpm publish ./my-package
```

**PowerShell:**
```powershell
cpm publish .\my-package
```
*Effect: Creates `~/.cpm/registry/my-package/0.1.0.cpm`*

### Search
Find packages in the registry.

**Bash/PowerShell:**
```bash
cpm search postgres
```
*Effect: Lists matching packages (e.g., `postgres - 14.1.0`)*

### Installing
Install a package from the registry.

**Bash:**
```bash
cpm install my-package --version 0.1.0
```

**PowerShell:**
```powershell
cpm install my-package --version 0.1.0
```
*Effect: Fetches, unpacks, renders, and installs the package.*
