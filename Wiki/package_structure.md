# Package Structure

A CPM package is a directory containing the code and configuration required to deploy an application on ColonyOS.

## Directory Layout

A standard package has the following structure:

```text
my-package/
├── colony.yaml       # Package manifest (Required)
├── values.yaml       # Default configuration values (Required)
└── templates/        # Template files (Required)
    ├── workflow.json
    └── ...
```

## files

### 1. colony.yaml (Manifest)
This file defines metadata about the package. It is used by CPM for versioning, search, and dependency management.

**Schema:**
```yaml
apiVersion: v1
name: my-package           # Name of the package (must be unique in registry)
version: 0.1.0             # SemVer version
description: A short description
maintainers:
  - name: Jane Doe
    email: jane@example.com
```

### 2. values.yaml (Values)
This file defines the default values (variables) that are passed to the templates. Users can override these values during installation using the `--set` flag.

**Example:**
```yaml
replicas: 1
image: "my-app:latest"
environment: "production"
```

### 3. templates/ (Templates)
This directory contains the actual ColonyOS resource definitions. Files here are processed by the Template Engine.

*   See [Templates](template.md) for more details on syntax and functions.

## CLI Usage

Commands related to creating and managing package structures:

- **Initialize a new package**:
  **Bash/PowerShell:**
  ```bash
  cpm init my-package
  ```
  *Generates the standard directory layout for you.*

- **Validate/Inspect**: 
  (Coming soon) `cpm lint my-package` to verify structure.
