# Colony Package Manager (CPM)

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Status](https://img.shields.io/badge/status-beta-orange.svg)
![Go Version](https://img.shields.io/badge/go-1.20%2B-blue)

**CPM** is the distributed package manager for [ColonyOS](https://github.com/colonyos/colonies). It streamlines the process of packaging, sharing, and deploying computational workflows across decentralized "Colonies". Think of it as **Helm for global computing**.

---

## 🚀 Features

- **📦 Package Management**: Initialize, pack, and install workflow applications.
- **🏛 Registry**: Built-in support for publishing and searching packages from a local or remote registry.
- **📝 Advanced Templating**: Dynamic configuration using Go templates and [Sprig](http://masterminds.github.io/sprig/) functions.
- **🔐 Secure**: Authenticated workflow submission using Ed25519 digital signatures.
- **🛠 Extensible**: Simple YAML manifests and JSON workflow definitions.

---

## 📚 Documentation

Detailed documentation is available in the **[Wiki](Wiki/)**:

- **[Package Structure](Wiki/package_structure.md)**: Understanding `colony.yaml`, `values.yaml`, and templates.
- **[Registry & Storage](Wiki/registry.md)**: How packaging and versioning works.
- **[Templating Engine](Wiki/template.md)**: Using variables, functions, and logic in your workflows.
- **[Authentication](Wiki/authentication.md)**: Security details and signing protocol.

---

## ⚡ Quick Start

### 1. Installation

Build the binary from source:

```bash
go build -o cpm.exe ./cmd/cpm
```

### 2. Initialize a Package

Create a new package skeleton:

```bash
./cpm init my-app
```

### 3. Deploy

Install the package to your Colony:

**Bash:**
```bash
./cpm install my-app --colonyid "af67..." --prvkey "e5b9..."
```

**PowerShell:**
```powershell
.\cpm install my-app --colonyid "af67..." --prvkey "e5b9..."
```

---

---
## 🗺️ Roadmap

We are actively working towards v1.0. Here is our high-level plan:

*   **Phase 1: Stabilization** (Current) - Comprehensive testing, Error handling polish, and Security auditing.
*   **Phase 2: Ecosystem** - Launching the **Remote Registry API** (C#), **Web Dashboard** (React), and `cpm publish` support.
*   **Phase 3: Public Beta** - CI/CD Integration, VS Code Extension, and Community Launch.

See the full [Release Roadmap](project-management/roadmap/release_roadmap.md) for details.

---

## 🔧 Architecture

CPM operates on a client-side architecture that interacts with a ColonyOS server.

1.  **Resolution**: CPM looks up packages in the registry (`$CPM_HOME/registry`).
2.  **Rendering**: It combines `templates/` with `values.yaml` (overridden by CLI flags) to generate the final workflow spec.
3.  **Submission**: The spec is signed with your private key and submitted to the ColonyOS `/api/workflows` endpoint.

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.
