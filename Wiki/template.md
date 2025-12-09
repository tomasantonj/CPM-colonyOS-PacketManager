# What is a Template?

In the context of the Colony Package Manager (CPM), a **template** is a file that serves as a blueprint for generating the configuration data (workflows, function specs, etc.) that ColonyOS executes.

## Technical Definition

1.  **File Format**: Templates are text files located in the `templates/` directory of a package. We support `.json`, `.yaml`, and `.tpl` extensions.
2.  **Engine**: They are processed by the Go `text/template` engine, which allows for dynamic content generation.
3.  **Inputs**: They receive a `Values` object (from `values.yaml` and CLI flags) to inject data.
4.  **Functions**: They have access to **Sprig** library functions (like `upper`, `trim`, `list`) and custom helpers (like `required`, `toYaml`, `toJson`) to perform logic and transformations.
5.  **Output**: All templates in a package are rendered and combined into a single JSON array `[...]`, which is then submitted to the ColonyOS backend.

## Example

A template file `workflow.json` allows you to write:

```json
{
  "name": "{{ .Values.appName }}-workflow",
  "replicas": {{ .Values.replicaCount }}
}
```

If `.Values.appName` is "my-app" and `.Values.replicaCount` is 3, CPM renders this into:

```json
{
  "name": "my-app-workflow",
  "replicas": 3
}
```

This allows users to reuse the same "package" for different environments (Dev, QA, Prod) by simply changing the input values.

## CLI Usage

You can control inputs to the template engine via the command line:

- **Override Values**:
  **Bash:**
  ```bash
  cpm install my-package --set replicas=5,appName="My Great App"
  ```
  **PowerShell:**
  ```powershell
  cpm install my-package --set "replicas=5,appName='My Great App'"
  ```
  *(Note: PowerShell quote handling can be tricky, ensure the argument is passed as a single string if it contains commas or spaces)*
  *Injects `replicas` and `appName` into `.Values`.*

- **Specify Version** (uses specific package version's templates):
  ```bash
  cpm install my-package --version 1.0.0
  ```
