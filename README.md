# ColonyOS Package Manager (CPM)

CPM is a package manager for [ColonyOS](https://github.com/colonyos/colonies), similar to Helm or npm, but designed for distributed computing workflows. It allows users to package, share, and deploy ColonyOS applications easily.

## Prerequisites

- **Go**: Version 1.20 or later.
- **ColonyOS**: A running ColonyOS instance (for submitting workflows).

## Installation

Build the binary from source:

```bash
go build -o cpm.exe ./cmd/cpm
```

## Usage

CPM supports the following core commands:

### 1. Initialize a Package
Create a new package with the required directory structure:

```bash
./cpm init my-package
```

This creates:
- `colony.yaml`: The package manifest.
- `values.yaml`: Default configuration values.
- `templates/`: Directory for template files (e.g., `workflow.json`).

### 2. Pack a Package
Compress a package directory into a distributable `.cpm` archive:

```bash
./cpm pack my-package
```

This generates a file named `{name}-{version}.cpm` (e.g., `my-package-0.1.0.cpm`).

### 3. Install a Package
Submit a package (either a directory or a `.cpm` archive) to ColonyOS:

```bash
./cpm install my-package-0.1.0.cpm
```

#### Overriding Values
You can override default values in `values.yaml` using the `--set` flag:

```bash
./cpm install my-package-0.1.0.cpm --set replicas=5 --set name="production-job"
```

## Package Requirements

A valid CPM package must contain:

1.  **`colony.yaml`**: Manifest file defining metadata.
    ```yaml
    apiVersion: v1
    name: my-package
    version: 0.1.0
    description: A sample package
    maintainers:
      - name: Your Name
    ```

2.  **`values.yaml`**: Default values available to templates.
    ```yaml
    replicas: 1
    image: ubuntu:latest
    ```

3.  **`templates/`**: Directory containing Go templates (usually JSON or YAML).
    Files in this directory are rendered using the values from `values.yaml` (and CLI overrides).
    
    *Example `templates/workflow.json`:*
    ```json
    {
      "name": "{{ .Values.name }}",
      "replicas": {{ .Values.replicas }}
    }
    ```
