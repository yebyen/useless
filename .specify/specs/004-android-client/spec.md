# Feature Specification: Android Client (Vind-Box)

**Feature Branch**: `004-android-client`  
**Created**: 2026-03-28  
**Status**: Draft  
**Input**: User description: "Build a Kotlin-based Android client that uses the shared WASM Brain."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Physical Button Push (Priority: P1)

As a user, I want a large button on my phone so that I can push it and see my daily count increase instantly.

**Why this priority**: This is the mobile manifestation of the Useless Machine.

**Independent Test**: Running the app in Android Studio, clicking the button triggers an API call to the Spin Gateway, which updates the Kubernetes state.

**Acceptance Scenarios**:

1. **Given** the app is open, **When** I click "Push Button," **Then** the screen updates to reflect the new count returned by the API.

### User Story 2 - Physical Box Feedback (Priority: P1)

As a user, I want to see when the machine is "flipping the switch back" so that it feels like a real useless box.

**Why this priority**: Without the back-and-forth reaction, it's just a counter, not a useless machine.

**Acceptance Scenarios**:
1. **Given** the switch is idle, **When** I push it, **Then** the UI shows a "Switch Flipped" state.
2. **Given** the Switch is flipped, **When** the Operator processes the action, **Then** the UI shows a "Machine Reacting" state before returning to idle.

### User Story 3 - UI Modes (Priority: P2)

As a user, I want to toggle between "Verbose" and "Quiet" modes.

**Why this priority**: Allows for both deep status visibility and a simple, playful interaction.

---

## Requirements *(mandatory)*

### Functional Requirements

- **FR-011**: System MUST provide an Android application built with Kotlin/Compose.
- **FR-012**: Android app MUST consume the Spin API (`/push` and `/status` endpoints).
- **FR-013**: Android app SHOULD support a "Verbose" mode (showing all K8s metadata) and a "Quiet" mode (pure useless box experience).
- **FR-014**: Android app MUST display the `machine_state` (idle, requested, flippingBack).

### Technical Requirements

- **TR-009**: Android app MUST use `Retrofit` or `Ktor` for HTTP communication.
- **TR-010**: Android app MUST be side-loadable via Android Studio.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-010**: The app compiles and runs in the Android Emulator.
- **SC-011**: The app successfully fetches and displays data from the `http://localhost:3000` (via emulator loopback).
- **SC-012**: UI reflects the "Zero-Split-Brain" logic (same status as K8s/MCP).

## Assumptions

- We assume the Spin API is reachable from the Android Emulator (typically `10.0.2.2`).
- We assume local side-loading is sufficient for the MVP (no marketplace).
- Offline WASM logic is deferred to v0.2.0.
