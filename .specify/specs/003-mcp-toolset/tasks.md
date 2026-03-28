# Tasks: MCP Toolset (Vind-Box)

**Input**: Design documents from `.specify/specs/003-mcp-toolset/`
**Prerequisites**: plan.md, spec.md

## Phase 1: Environment Setup

**Purpose**: Prepare the Python environment.

- [ ] T001 Create `components/mcp/` directory
- [ ] T002 Create `requirements.txt` with `mcp`, `extism`, `pydantic`
- [ ] T003 Initialize `uv` or `venv` for the MCP server

---

## Phase 2: WASM Bridge (Python)

**Purpose**: Build the Python interface to the WASM Brain.

- [ ] T004 Implement `WasmBridge` class in `components/mcp/wasm_bridge.py`
- [ ] T005 Implement Pydantic models in `components/mcp/models.py` matching `specs/core/state.yaml`
- [ ] T006 Implement Extism call logic for `push_button_wasm` and `get_status_wasm`

---

## Phase 3: MCP Server Implementation

**Purpose**: Expose the bridge via the Model Context Protocol.

- [ ] T007 Initialize MCP `Server` in `components/mcp/server.py`
- [ ] T008 Implement tool registration for `push_button`
- [ ] T009 Implement tool registration for `get_status`
- [ ] T010 Implement stdin/stdout server execution

---

## Phase 4: Validation & Tooling

**Purpose**: Verify the LLM can use the tools.

- [ ] T011 Run `mcp-inspector` to discover and test the tools
- [ ] T012 Update root `Makefile` with `make run-mcp` target

---

## Dependencies

- **Phase 1** is required for all subsequent phases.
- **Phase 3** depends on Phase 2.
- **Phase 4** depends on completion of Phase 3.
