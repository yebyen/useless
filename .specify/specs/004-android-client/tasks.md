# Tasks: Android Client (Vind-Box)

**Input**: Design documents from `.specify/specs/004-android-client/`
**Prerequisites**: plan.md, spec.md

## Phase 1: Android Scaffolding

**Purpose**: Establish the mobile project foundation.

- [ ] T001 Create `components/android/` directory
- [ ] T002 Initialize `gradle` wrapper and project structure
- [ ] T003 Create `MainActivity.kt` with basic Compose layout

---

## Phase 2: API Integration (Kotlin)

**Purpose**: Connect the mobile app to the Cloud Gateway.

- [ ] T004 Add Ktor dependencies to `build.gradle.kts`
- [ ] T005 Implement `UselessMachineStatus` data class
- [ ] T006 Implement `VindBoxClient` for `/status` and `/push`

---

## Phase 3: UI Development

**Purpose**: Build the user interface.

- [ ] T007 Implement the main screen with `Button` and `Text` components
- [ ] T008 Add `ViewModel` to manage state and API calls
- [ ] T009 Implement nag state visual indicator (e.g., color change)

---

## Phase 4: Validation

**Purpose**: Verify the mobile piece of the puzzle.

- [ ] T010 Build the APK via `./gradlew assembleDebug`
- [ ] T011 Document emulator setup for Spin API connectivity

---

## Dependencies

- **Phase 1** is the entry point.
- **Phase 3** depends on the models from Phase 2.
