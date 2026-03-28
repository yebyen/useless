# Feature Specification: Core Brain (Vind-Box)

**Feature Branch**: `001-core-brain`  
**Created**: 2026-03-28  
**Status**: Draft  
**Input**: User description: "Build the WASM Brain that tracks daily count and last pushed timestamp."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Push the Button (Priority: P1)

As a user, I want to push a button so that my daily activity is recorded and my streak/nag timer is reset.

**Why this priority**: This is the core interaction of the "Useless Machine." Without it, the system has no purpose.

**Independent Test**: The WASM module can be invoked with a `push_button` action and a current state, returning an updated state with an incremented count and a refreshed timestamp.

**Acceptance Scenarios**:

1. **Given** a state with `daily_count: 5`, **When** `push_button` is called, **Then** `daily_count` becomes `6`.
2. **Given** a state with an old `last_pushed` timestamp, **When** `push_button` is called, **Then** `last_pushed` is updated to the current time provided by the host.

---

### User Story 2 - Check Nag State (Priority: P1)

As a user, I want the system to know if I've been lazy so that it can "nag" me to push the button.

**Why this priority**: This drives the "Useless Machine" behavior and justifies the persistence of the timestamp.

**Independent Test**: The WASM module can be queried for its `nag_state` given a state and the current time, returning a boolean or status string.

**Acceptance Scenarios**:

1. **Given** `last_pushed` was 25 hours ago, **When** state is checked, **Then** `nag_state` is `true`.
2. **Given** `last_pushed` was 1 hour ago, **When** state is checked, **Then** `nag_state` is `false`.

---

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST encapsulate all business logic in a WASM component.
- **FR-002**: System MUST track `daily_count` as an integer.
- **FR-003**: System MUST track `last_pushed` as an ISO8601 timestamp or Unix epoch.
- **FR-004**: System MUST determine `nag_state` based on a 24-hour threshold.
- **FR-005**: WASM module MUST be stateless and receive all context (state + current time) via host parameters.
- **FR-006**: WASM module MUST communicate via JSON strings for input and output.

### Technical Requirements (The Trunk)

- **TR-001**: Data models MUST be defined in an OpenAPI 3.0+ specification.
- **TR-002**: Function signatures MUST be defined in a WIT (Wasm Interface Type) file.
- **TR-003**: Host functions (Time, Logging) MUST be explicitly defined in the WIT file.

### Key Entities

- **UselessMachineState**: The primary state object containing `daily_count` and `last_pushed`.
- **NagStatus**: A calculated view of the state determining if nagging is required.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Rust-based WASM module compiles and passes all unit tests for `push_button` and `nag_state`.
- **SC-002**: Go-based host can successfully load the WASM module and execute the logic using a mock state.
- **SC-003**: Spec-generated types match between Rust and Go implementations.

## Assumptions

- We assume the Host will provide the "current time" to the Brain to keep the Brain deterministic and easy to test.
- We assume Extism or the WASM Component Model will be used as the runtime.
- We assume JSON serialization/deserialization overhead is acceptable for this MVP.
