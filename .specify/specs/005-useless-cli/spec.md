# Feature Specification: Useless CLI (Vind-Box)

**Feature Branch**: `005-useless-cli`  
**Created**: 2026-03-28  
**Status**: Draft  
**Input**: User description: "Build a CLI that can interact with K8s or Spin API, using the shared WASM Brain for logic."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Command-Line Button Pushing (Priority: P1)

As a power user, I want to run `useless push` so that I can interact with the machine without opening a browser or an app.

**Why this priority**: Provides a low-overhead interface for developers and scripts.

**Independent Test**: Running the CLI with `push` command updates the Kubernetes state.

**Acceptance Scenarios**:

1. **Given** a valid configuration, **When** I run `useless push`, **Then** the CLI patches the K8s CRD (or calls Spin API) and reports success.

---

### User Story 2 - Status Query (Priority: P1)

As a user, I want to run `useless status` so that I can see the machine's state in my terminal.

**Why this priority**: Essential for visibility and debugging.

**Independent Test**: `useless status` returns the correct `dailyCount` and `isNagging` state.

---

## Requirements *(mandatory)*

### Functional Requirements

- **FR-015**: System MUST provide a Go-based CLI named `useless`.
- **FR-016**: CLI MUST support configuration via a dotfile (`~/.mecris.config` or `~/.config/useless/backend.conf`).
- **FR-017**: CLI MUST be capable of interacting with the Kubernetes API directly.
- **FR-018**: CLI MUST be capable of interacting with the Spin API via HTTP.
- **FR-019**: CLI SHOULD eventually support "offline" mode using the local WASM Brain.

### Technical Requirements

- **TR-011**: CLI MUST use `spf13/cobra` for command handling.
- **TR-012**: CLI MUST share the Go types used by the Operator.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-013**: `useless status` successfully displays JSON or formatted text from K8s.
- **SC-014**: Switching between "k8s" and "api" backends in the config file works correctly.
- **SC-015**: The CLI uses the same `UselessMachineState` definition as the rest of the system.

## Assumptions

- We assume the user has a valid `kubeconfig` if using the K8s backend.
- We assume the Spin API is reachable via HTTP if using the API backend.
