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
Submit a package to ColonyOS. You can install from a local file or from the configured registry.

**Authentication:**
To submit workflows to a real ColonyOS server, you must provide credentials:
```bash
./cpm install my-package --colonyid <ID> --prvkey <HEX_KEY> --host <HOST> --port <PORT>
```

**From Local File:**
```bash
./cpm install my-package-0.1.0.cpm
```

**From Registry:**
(Requires `CPM_HOME` to point to a valid registry location)
```bash
./cpm install my-package --version 0.1.0
```

#### Overriding Values
You can override default values in `values.yaml` using the `--set` flag:

```bash
./cpm install my-package --version 0.1.0 --set replicas=5
```

### 4. Publish a Package
Publish a package directory to the registry:

```bash
./cpm publish ./my-package
```

### 5. Search for Packages
Search the registry for available packages:

```bash
./cpm search my-package
```

### 6. List Installed Packages
View currently installed packages and their status:

```bash
./cpm list
```

### 7. Uninstall a Package
Remove a package from the local state (and ColonyOS):

```bash
./cpm uninstall my-package
```

## Configuration

### State Directory (`CPM_HOME`)
By default, CPM stores state in `~/.cpm`. You can override this using the `CPM_HOME` environment variable.
- **Registry**: stored in `$CPM_HOME/registry` (currently a local directory mock).
- **Installed State**: stored in `$CPM_HOME/state.json`.

**PowerShell:**
```powershell
$env:CPM_HOME=".\.cpm"
./cpm install ...
```

**Bash:**
```bash
export CPM_HOME=./.cpm
./cpm install ...
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

3.  **`templates/`**: Directory containing Go templates (`.json`, `.yaml`, `.tpl`).
    Files in this directory are rendered using the values from `values.yaml` (and CLI overrides).

    *Example `templates/workflow.json`:*
    ```json
    {
      "name": "{{ .Values.name }}",
      "replicas": {{ .Values.replicas }}
    }
    ```

## Advanced Templating

CPM uses the Go template engine extended with **Sprig** functions and custom helpers:

- **Sprig Functions**: standard functions like `upper`, `lower`, `trim`, `list`, `dict`, `default`, etc.
- **`required "message" <value>`**: Fails rendering if value is empty.
- **`toYaml <value>`**: Converts complex objects to YAML string.
- **`toJson <value>`**: Converts complex objects to JSON string.

*Example:*
```json
{
  "name": "{{ .Values.name | upper }}",
  "config": {{ .Values.extraConfig | toJson }}
}
```
