# Implementation Plan: Core Brain

**Branch**: `001-core-brain` | **Date**: 2026-03-28 | **Spec**: [spec.md](./spec.md)

## Summary
Build the "Useless Machine" logic in a Rust-based WASM component. The Brain will be stateless, receiving its state and current time from the Host (Go Operator or Spin API) and returning an updated state.

## Technical Context
- **Language/Version**: Rust 1.75+ (Brain), Go 1.22+ (Operator)
- **Primary Dependencies**: `serde` (Rust), `extism` or `wasmtime` (Host), `controller-runtime` (Go)
- **Storage**: Kubernetes ETCD via Custom Resource Definition (CRD)
- **Testing**: `cargo test` (Rust Unit Tests), `go test` (Host Integration Tests)
- **Target Platform**: WASM (WASI or Extism)

## Constitution Check
- **Zero-Split-Brain**: PASSED (One WASM module for all logic)
- **Spec-Driven**: PASSED (OpenAPI/WIT will be created first)
- **Functional Brain**: PASSED (Stateless logic)
- **Test-Driven**: PASSED (Unit tests for logic before host integration)

## Project Structure

### Documentation (this feature)
```text
.specify/specs/001-core-brain/
├── spec.md              # Feature specification
├── plan.md              # This file
└── tasks.md             # Actionable tasks
```

### Source Code
```text
specs/
└── core/
    ├── state.yaml       # OpenAPI Spec (The Trunk)
    └── brain.wit        # WIT Spec (The Trunk)

components/
├── brain/               # Rust WASM Project
│   ├── src/
│   │   ├── lib.rs       # WASM Exports
│   │   └── logic.rs     # Pure business logic
│   └── tests/           # Rust Unit Tests
└── operator/            # Go Kubernetes Operator
    ├── api/             # CRD Definitions (Generated)
    ├── internal/
    │   └── controller/  # WASM-loading Reconciler
    └── main.go

Makefile                 # The Universal Remote Control
```

## Implementation Strategy

### Phase 1: The Trunk (Specification)
1. Define `UselessMachineState` in `specs/core/state.yaml`.
2. Define function exports (`push_button`, `get_status`) in `specs/core/brain.wit`.
3. Define host imports (logging) in `specs/core/brain.wit`.

### Phase 2: The Leaves (Brain Implementation)
1. Initialize Rust project in `components/brain`.
2. Generate Rust types from OpenAPI spec.
3. Implement business logic using TDD (Red-Green-Refactor).
4. Export logic via WASM/Extism functions.

### Phase 3: The Host (Operator Integration)
1. Initialize Go Operator using Kubebuilder in `components/operator`.
2. Integrate WASM loading logic in the Reconciler.
3. Validate with a mock CRD and local WASM execution.

## Success Criteria
- [ ] `specs/core/state.yaml` and `brain.wit` are valid.
- [ ] `cargo test` passes in `components/brain`.
- [ ] `make generate` produces consistent types in Rust and Go.
- [ ] Go Reconciler can call `push_button` in the WASM Brain and see the count increment.
