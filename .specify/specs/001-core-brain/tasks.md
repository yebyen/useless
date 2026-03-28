# Tasks: Core Brain (Vind-Box)

**Input**: Design documents from `.specify/specs/001-core-brain/`
**Prerequisites**: plan.md, spec.md

## Phase 1: Setup (The Trunk)

**Purpose**: Define the formal contracts that govern the system.

- [x] T001 Create `specs/core/state.yaml` with `UselessMachineState` (OpenAPI)
- [x] T002 Create `specs/core/brain.wit` with function signatures (WIT)
- [x] T003 [P] Create root `Makefile` with `generate` and `test` targets
- [x] T004 [P] Create `specify.yaml` to map specs to component directories

---

## Phase 2: Foundational (Component Initialization)

**Purpose**: Initialize the polyglot projects.

- [x] T005 Initialize Rust project in `components/brain`
- [x] T006 Initialize Go project in `components/operator` using `kubebuilder init`
- [x] T007 Configure `components/brain` for WASM/WASI target
- [x] T008 [P] Add `serde` and `serde_json` to `components/brain`
- [x] T009 [P] Add `extism` and `controller-runtime` to `components/operator`

---

## Phase 3: User Story 1 - Push Button (Priority: P1) 🎯 MVP

**Goal**: Implement the button-pushing logic in the WASM Brain.

**Independent Test**: `cargo test` in `components/brain` verifies `push_button` logic.

### Tests for User Story 1
- [x] T010 [US1] Write failing unit test for `push_button` in `components/brain/src/logic.rs`
- [x] T011 [US1] Write failing test for JSON serialization of `UselessMachineState`

### Implementation for User Story 1
- [x] T012 [US1] Generate Rust types from `specs/core/state.yaml`
- [x] T013 [US1] Implement `push_button` logic in `components/brain/src/logic.rs`
- [x] T014 [US1] Export `push_button` via Extism/WASM in `components/brain/src/lib.rs`
- [x] T015 [US1] Verify T010 and T011 pass

---

## Phase 4: User Story 2 - Check Nag State (Priority: P1)

**Goal**: Implement the nag-checking logic in the WASM Brain.

**Independent Test**: `cargo test` in `components/brain` verifies `nag_state` logic.

### Tests for User Story 2
- [x] T016 [US2] Write failing unit test for `is_nagging` logic
- [x] T017 [US2] Write failing test for threshold-based time calculation

### Implementation for User Story 2
- [x] T018 [US2] Implement `is_nagging` logic in `components/brain/src/logic.rs`
- [x] T019 [US2] Export `get_status` via Extism/WASM in `components/brain/src/lib.rs`
- [x] T020 [US2] Verify T016 and T017 pass

---

## Phase 5: Host Integration (Go Operator)

**Purpose**: Load and execute the WASM Brain from the Go Reconciler.

- [x] T021 Generate Go types from `specs/core/state.yaml`
- [x] T022 Implement `WasmRunner` in `components/operator/internal/wasm/runner.go`
- [x] T023 Integrate `WasmRunner` into the Operator Reconciler loop
- [x] T024 Write integration test in Go that loads the compiled `.wasm` and calls `push_button`

---

## Dependencies

- **Phase 1** must be complete before any code generation.
- **Phase 2** can run in parallel with Phase 1.
- **Phase 3 & 4** depend on Phase 1 & 2.
- **Phase 5** depends on Phase 3 & 4 (needs the compiled WASM).
