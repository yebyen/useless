# Research & Architectural Decisions: Core Brain

This document records the key architectural decisions made during the initialization of the Vind-Box project.

## ⚖️ Core Decisions

### 1. Functional & Stateless Brain
-   **Decision:** The WASM Brain will be strictly functional.
-   **Rationale:** To prevent "split-brain" issues, the host must read state from Kubernetes and pass it as a JSON blob to the WASM module. The WASM module returns the updated state. This makes the logic portable across Go (Operator), Rust (Spin), and Kotlin (Android).

### 2. Hybrid Specification ("Trunk")
-   **Decision:** Use **OpenAPI** for data models and **WIT** for function signatures.
-   **Rationale:** OpenAPI is the standard for Kubernetes CRDs and MCP tools. WIT provides the strict calling convention required for the WASM guest/host boundary. This ensures data consistency while maintaining type-safe binary interfaces.

### 3. OCI-First Distribution
-   **Decision:** The WASM Brain is the only versioned binary artifact released via OCI.
-   **Rationale:** Since the Brain is the single source of truth for logic, it must be versioned and immutable. OCI registries (GHCR) allow all host environments to pull the exact same logic at runtime or build time.

### 4. Flux GitOps Model
-   **Decision:** Follow the `flux2-kustomize-helm-example` and OCI-centric pattern.
-   **Rationale:** We utilize a "Scripted Assembly" approach where component-specific YAMLs are collected into a top-level `deploy/` directory for Flux to consume.

### 5. Test-Driven Generation (TDG)
-   **Decision:** Implement unit tests (the "Leaves") before or alongside code.
-   **Rationale:** Spec-driven development defines the "Trunk," but TDG ensures that individual components are functionally correct. Implementation hindsight from failed tests is used to refine the Spec.

## 🛠️ Technology Stack
-   **Runtime:** Extism (initially) with a path toward the WASM Component Model.
-   **Orchestration:** `controller-runtime` (Go) for the Operator.
-   **API:** Rust (Spin) for the Cloud Gateway.
-   **MCP:** Python (FastMCP) for LLM tools.
-   **E2E:** GitHub Actions + `kind` for automated cluster testing.

## ⚖️ MCP & Hosting Decisions

### 1. K8s-Aware MCP Server
-   **Decision:** The MCP server interacts directly with the Kubernetes API using the `kubernetes` Python library.
-   **Rationale:** To ensure the LLM always sees the true state of the machine (and not a cached or out-of-sync value), the tools fetch the current state from the CRD before invoking the WASM Brain.

### 2. Platform-as-a-Product (Init Containers)
-   **Decision:** Deliver the WASM Brain to Kubernetes-hosted components (Operator, MCP, API) via Init Containers.
-   **Rationale:** Using an Init Container to pull the `vind-box-brain` OCI artifact into a shared volume allows the business logic to be updated independently of the host's container image, ensuring all parts of the system pick up new logic simultaneously.

### 3. Abstract Tool Interface
-   **Decision:** MCP tools (e.g., `push_button`) are abstract and do not require Kubernetes-specific arguments (like namespace or resource name).
-   **Rationale:** The LLM agent should focus on business logic. The MCP server handles the mapping to the `global` instance in Kubernetes internally.
