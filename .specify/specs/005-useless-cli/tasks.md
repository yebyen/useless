# Tasks: Useless CLI (Vind-Box)

**Input**: Design documents from `.specify/specs/005-useless-cli/`
**Prerequisites**: plan.md, spec.md

## Phase 1: Scaffolding

**Purpose**: Build the CLI skeleton.

- [ ] T001 Create `components/cli/` directory and `go mod init`
- [ ] T002 Add `cobra`, `viper`, and `client-go` dependencies
- [ ] T003 Implement `cmd/root.go`, `cmd/push.go`, and `cmd/status.go`

---

## Phase 2: Configuration & Backends

**Purpose**: Connect the CLI to the sources of truth.

- [ ] T004 Implement `pkg/config/config.go` to load `~/.mecris.config`
- [ ] T005 Implement `pkg/backend/interface.go` for backend abstraction
- [ ] T006 Implement `pkg/backend/k8s.go` (direct CRD access)
- [ ] T007 Implement `pkg/backend/api.go` (Spin Gateway access)

---

## Phase 3: Assembly & Validation

**Purpose**: Make it usable and verified.

- [ ] T008 Update root `Makefile` with `make build-cli` target
- [ ] T009 Verify `useless status` works against local `vcluster`
- [ ] T010 Verify `useless status --backend api` works against local Spin

---

## Dependencies

- **Phase 1** is required for all logic.
- **Phase 2** is the core implementation.
