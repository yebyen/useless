# Vind-Box Constitution

## Core Principles

### I. Zero-Split-Brain
The core business logic MUST be encapsulated in a single WASM component that runs identically across all host environments (Go Operator, Spin API, MCP Server, Android App).

### II. Spec-Driven Development (SDD)
All interfaces and data models MUST be defined in formal specifications (OpenAPI for data, WIT for function signatures) before implementation. These specs are the "Source of Truth" (the Trunk).

### III. Functional & Stateless Brain
The WASM Brain MUST be strictly functional. It consumes JSON and produces JSON. It has no knowledge of its host. All I/O (persistence, time, network) is handled by the Host via standard environment interfaces.

### IV. Test-Driven Generation (TDG)
Implementations MUST follow a Test-Driven Generation approach. Unit tests (the Leaves) are built first and must pass to validate functionality. The Spec is refined by implementation hindsight.

### V. OCI-First Distribution
The WASM Brain is the primary versioned binary artifact, distributed via OCI registries. Other components are source-distributed in the early stages.

### VI. Flux-Managed Infrastructure
Kubernetes deployments MUST follow the Flux GitOps model, specifically the OCI-centric pattern, utilizing a scripted assembly into a top-level `deploy/` directory.

## Technology Stack
- **Core Brain:** Rust (WASM)
- **Operator:** Go (controller-runtime / Kubebuilder)
- **API:** Rust (Spin)
- **Tools:** Python (MCP)
- **Client:** Kotlin (Android)
- **Infrastructure:** Kubernetes (Kind/Vcluster), Flux, Kustomize

## Development Workflow
1. **Specify:** Update OpenAPI/WIT specs in `specs/core/`.
2. **Plan:** Outline implementation and testing strategy.
3. **Generate:** Use `make generate` to update bindings and types.
4. **Implement:** Follow TDG to build functionality.
5. **Assemble:** Use `make assemble` to prepare `deploy/` artifacts.
6. **Validate:** Run `make e2e-test` in a local `kind` cluster.

## Governance
This Constitution supersedes all other practices. Amendments require a deliberate update to this document and a corresponding migration plan for affected components.

**Version**: 0.1.0 | **Ratified**: 2026-03-28 | **Last Amended**: 2026-03-28
