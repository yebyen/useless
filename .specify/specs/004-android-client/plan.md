# Implementation Plan: Android Client

**Branch**: `004-android-client` | **Date**: 2026-03-28 | **Spec**: [spec.md](./spec.md)

## Summary
Scaffold a Kotlin-based Android application using Jetpack Compose to provide a mobile interface for the Useless Machine.

## Technical Context
- **Language**: Kotlin 1.9+
- **UI Framework**: Jetpack Compose
- **Networking**: Ktor Client or Retrofit
- **Minimum SDK**: API 26 (Android 8.0)

## Project Structure

### Documentation
```text
.specify/specs/004-android-client/
├── spec.md
├── plan.md
└── tasks.md
```

### Source Code
```text
components/android/
├── app/
│   ├── src/main/java/io/mecris/vindbox/
│   │   ├── MainActivity.kt
│   │   ├── api/             # API Models and Clients
│   │   └── ui/              # Compose Screens
│   └── build.gradle.kts
└── build.gradle.kts
```

## Implementation Strategy

### Phase 1: Android Scaffolding
1. Create `components/android/` directory structure.
2. Initialize Gradle wrapper and root `build.gradle.kts`.
3. Scaffold the basic `MainActivity` with a Compose "Hello World."

### Phase 2: API Integration
1. Define Kotlin models for `UselessMachineStatus` (matching OpenAPI).
2. Implement a `VindBoxClient` using Ktor to call the Spin Gateway.
3. Handle the emulator loopback address (`10.0.2.2`).

### Phase 3: UI Implementation
1. Build a simple UI with:
    - A large, centered "Push" button.
    - A counter display for `dailyCount`.
    - A status indicator for `isNagging`.
2. Implement basic polling or a refresh button to update the state.

### Phase 4: Validation
1. Verify the app compiles.
2. Verify it can reach the Spin API when running locally.

## Success Criteria
- [ ] Android app builds and deploys to an emulator.
- [ ] "Push" button triggers a successful API call.
- [ ] UI accurately reflects the state stored in Kubernetes.
