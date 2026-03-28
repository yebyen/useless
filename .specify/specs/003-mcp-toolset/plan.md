# Implementation Plan: MCP Toolset

**Branch**: `003-mcp-toolset` | **Date**: 2026-03-28 | **Spec**: [spec.md](./spec.md)

## Summary
Expose the WASM Brain logic as Model Context Protocol (MCP) tools using a Python-based server and the Extism SDK.

## Technical Context
- **Language**: Python 3.10+
- **Primary Libraries**: `mcp`, `extism`, `pydantic`
- **WASM Interop**: Extism Python SDK
- **Data Schemas**: Pydantic models derived from `specs/core/state.yaml`

## Project Structure

### Documentation
```text
.specify/specs/003-mcp-toolset/
├── spec.md
├── plan.md
└── tasks.md
```

### Source Code
```text
components/mcp/
├── server.py        # MCP Server entry point
├── wasm_bridge.py   # Extism bridge
└── requirements.txt # Python dependencies
```

## Implementation Strategy

### Phase 1: Environment Setup
1. Create `components/mcp/` directory.
2. Initialize virtual environment or `uv` project.
3. Install `mcp`, `extism`, and `pydantic`.

### Phase 2: WASM Bridge
1. Implement `WasmBridge` class to load `brain.wasm`.
2. Map Python tool calls to WASM exports (`push_button_wasm`, `get_status_wasm`).
3. Ensure JSON serialization/deserialization matches the Spec.

### Phase 3: MCP Server
1. Initialize the MCP `Server` object.
2. Define `@server.list_tools()` to advertise `push_button` and `get_status`.
3. Define `@server.call_tool()` to execute the logic via the `WasmBridge`.
4. Implement standard MCP I/O (stdin/stdout).

### Phase 4: Validation
1. Use `mcp-inspector` to verify tools are discovered.
2. Run tool calls through the inspector to see logic execution.

## Success Criteria
- [ ] `mcp-inspector` successfully discover tools.
- [ ] Tool calls trigger WASM execution and return correct JSON.
- [ ] Zero business logic is written in Python.
