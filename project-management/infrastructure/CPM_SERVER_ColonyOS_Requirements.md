# ColonyOS Server Specs

**Server Purpose:**
This infrastructure hosts the **ColonyOS Core** (Go).
*   **Role:** Execution Engine (Meta-Orchestrator).
*   **Function:** Manages the state of colonies, accepts workflow submissions, and assigns tasks to executors.
*   **Why we need it:** This is the "brain" of the entire system. Without it, CPM packages cannot be deployed or executed.

**Component:** Execution Engine
**Type:** Stateful VM / Hybrid

## 1. Hardware Limits
*   **CPU:** 2 vCPU+
*   **RAM:** 4GB+
*   **Disk:** 20GB+ SSD (Persistent)

## 2. Software Stack
*   **OS:** Ubuntu 20.04/22.04 LTS
*   **Runtime:** Docker Engine 20.10+
*   **Binary:** `colonies_server` (Run as user `colonyos`; ulimit 65535)

## 3. Services
*   **PostgreSQL:** v13+ (Persistent Volume).
*   **Redis:** v6+ (Auth enabled).

## 4. Network
| Service | Port | Protocol | Access |
| :--- | :--- | :--- | :--- |
| **gRPC** | 50051 | HTTP/2 | Ingress / Public |
| **REST** | 8080 | HTTP | Ingress / Public |
| **SSH** | 22 | TCP | Admin Only |

## 5. Security
*   **Identity:** Server Ed25519 PrvKey required at startup.
*   **TLS:** Mandatory for gRPC. Use Let's Encrypt.
