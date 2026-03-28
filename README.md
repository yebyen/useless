# Vind-Box: The Next-Generation Useless Machine

Vind-Box is an architectural blueprint for building **Zero-Split-Brain** agent systems. It uses WebAssembly (WASM) to ensure that core business logic runs identically across every part of the stack—from Kubernetes Operators to Android apps.

## 🚀 Key Concepts

-   **Zero-Split-Brain:** One Rust-based WASM module ("The Brain") contains all logic. No more re-implementing rules in Go for the backend and Kotlin for the mobile app.
-   **Spec-Driven Development (SDD):** The system is built from the "Trunk" out. Data models (OpenAPI) and interfaces (WIT) are defined before code is written.
-   **Vind (vcluster-in-docker):** Uses local Kubernetes clusters for high-fidelity testing and GitOps-ready deployments.

## 🏗️ Architecture

1.  **The Brain (Rust/WASM):** Stateless business logic (tracks `daily_count` and `nag_state`).
2.  **The Operator (Go):** A Kubernetes controller that loads the WASM Brain to reconcile state stored in a Custom Resource (CRD).
3.  **The API (Spin/Rust):** A serverless-style web service that hits the same K8s state and logic.
4.  **The Client (Android/Kotlin):** The user interface (Planned).
5.  **Infrastructure (Flux):** GitOps manifests for automated deployment.

## 🛠️ Getting Started

### Prerequisites
-   Rust & Cargo (with `wasm32-wasip1` target)
-   Go 1.22+
-   Spin CLI
-   A local Kubernetes cluster (e.g., Kind or Vcluster)

### Build & Test
```bash
make generate  # (Upcoming: Full automation of OpenAPI/WIT generation)
make build     # Compiles the WASM Brain and Go Operator
make test      # Runs unit and integration tests
```

## 📈 Roadmap

- [x] **Phase 1: Core Brain Spec & Rust Implementation**
- [x] **Phase 2: Go Operator Scaffolding & WASM Integration**
- [x] **Phase 3: Spin API Layer (Mocked)**
- [ ] **Phase 4: Flux GitOps Bootstrap**
- [ ] **Phase 5: Android Client Prototype**
- [ ] **Phase 6: MCP Toolset for LLM Interaction**

## 📄 Design Artifacts
Detailed specifications, plans, and task lists are maintained in the `.specify/` directory:
- `Constitution`: [Project Governance](.specify/memory/constitution.md)
- `Core Brain Spec`: [.specify/specs/001-core-brain/spec.md](.specify/specs/001-core-brain/spec.md)
- `Implementation Plan`: [.specify/specs/001-core-brain/plan.md](.specify/specs/001-core-brain/plan.md)

---
*Built with ❤️ using Gemini CLI and Specify-CLI.*
