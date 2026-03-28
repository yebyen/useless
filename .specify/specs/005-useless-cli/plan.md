# Implementation Plan: Useless CLI

**Branch**: `005-useless-cli` | **Date**: 2026-03-28 | **Spec**: [spec.md](./spec.md)

## Summary
Build a Go-based CLI tool named `useless` that allows users to interact with the Useless Machine via direct Kubernetes API calls or the Spin Cloud Gateway.

## Technical Context
- **Language**: Go 1.22+
- **CLI Framework**: `spf13/cobra`
- **Config Management**: `spf13/viper`
- **K8s Client**: `sigs.k8s.io/controller-runtime/pkg/client`

## Project Structure

### Documentation
```text
.specify/specs/005-useless-cli/
├── spec.md
├── plan.md
└── tasks.md
```

### Source Code
```text
components/cli/
├── cmd/
│   ├── root.go      # Root command and config loading
│   ├── push.go      # 'push' command
│   └── status.go    # 'status' command
├── pkg/
│   ├── backend/     # Backend interfaces (K8s, Spin)
│   └── config/      # Config file handling
└── main.go
```

## Implementation Strategy

### Phase 1: CLI Scaffolding
1. Create `components/cli/` directory.
2. Initialize Go module.
3. Scaffold `cobra` root and subcommands.

### Phase 2: Configuration
1. Implement logic to load `~/.mecris.config` (YAML/TOML).
2. Support setting the default backend (k8s vs api).

### Phase 3: Backend Implementation
1. **K8s Backend:** Re-use the patching logic from the MCP server (translated to Go).
2. **Spin Backend:** Simple HTTP client calling `/status` and `/push`.

### Phase 4: Offline Mode (Future Proofing)
1. Add a placeholder for `local` backend that uses the WASM Brain via `extism-go`.

## Success Criteria
- [ ] CLI compiles to a single binary.
- [ ] `useless push` successfully triggers an action in K8s.
- [ ] `useless status` displays the current count and nag state.
