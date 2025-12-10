# ColonyOS Automation Specs

# ColonyOS Automation Specs

**Server Purpose:**
This document defines the **Deployment Logic** (Ansible/Bash).
*   **Role:** Configuration Management.
*   **Function:** Bootstraps the server from a fresh OS to a running ColonyOS state.
*   **Why we need it:** To ensure reproducible, idempotent deployments rather than manual "clicking around".

**Component:** Provisioning
**Tooling:** Ansible (Preferred) or Bash

## 1. Deployment Targets
*   Bootstrap standard Linux VM (see `CPM_SERVER_ColonyOS_Requirements.md`).
*   Install: `docker.io`, `postgresql-client`, `redis-tools`, `jq`.

## 2. Orchestration
**Strategy:** Docker Compose.
*   `postgres`: Bind mount `/var/lib/postgresql/data`.
*   `redis`: Configured with `requirepass`.
*   `colonyos-server`: Env vars for DB/Redis connection.

## 3. Post-Install Bootstrapping (`init_colony.sh`)
Must run exactly once:
1.  Check for `server.key`. If missing, generate (`colonyos key`).
2.  Register Root Colony (`cpm_colony`).
3.  Register default Executor.
4.  Export Generated Credentials -> `.env`.

## 4. Maintenance
*   **Backup:** `pg_dump` of `colonyos` db to S3/Blob.
*   **Logs:** Ship Docker logs to centralized aggregation.
