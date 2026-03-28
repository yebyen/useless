# Feature Specification: MCP Toolset (Vind-Box)

**Feature Branch**: `003-mcp-toolset`  
**Created**: 2026-03-28  
**Status**: Draft  
**Input**: User description: "Expose the WASM Brain logic as Model Context Protocol (MCP) tools for an LLM."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - LLM Button Pushing (Priority: P1)

As an LLM agent, I want to use a `push_button` tool so that I can interact with the Useless Machine on behalf of the user.

**Why this priority**: This is the primary way for an AI agent to become a "user" of the system.

**Independent Test**: Running the MCP server and invoking the `push_button` tool results in the WASM logic being executed and returning the updated state.

**Acceptance Scenarios**:

1. **Given** the MCP server is running, **When** the LLM calls `push_button`, **Then** the server loads the WASM Brain, executes `push_button_wasm`, and returns the JSON result.

---

### User Story 2 - LLM Nag Check (Priority: P1)

As an LLM agent, I want to use a `get_status` tool so that I can determine if the user needs to be nagged.

**Why this priority**: Enables the LLM to proactively assist the user based on the machine's state.

**Independent Test**: Invoking `get_status` through the MCP inspector/client returns the `is_nagging` boolean from the WASM Brain.

---

## Requirements *(mandatory)*

### Functional Requirements

- **FR-007**: System MUST provide a Python-based MCP server.
- **FR-008**: MCP server MUST load the *exact same* WASM Brain (`brain.wasm`) used by the Operator and API.
- **FR-009**: MCP server MUST expose `push_button` and `get_status` as Model Tools.
- **FR-010**: MCP server MUST follow the "Zero-Split-Brain" principle (no logic re-implementation in Python).

### Technical Requirements

- **TR-007**: MCP server MUST use the `mcp` and `extism` Python libraries.
- **TR-008**: Tool schemas MUST be derived from the `specs/core/state.yaml` OpenAPI models.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-007**: `mcp-inspector` shows two tools: `push_button` and `get_status`.
- **SC-008**: Invoking tools via the MCP server produces results identical to the Go Operator's integration tests.
- **SC-009**: Zero logic is implemented in the Python server; it is purely a proxy to the WASM Brain.

## Assumptions

- We assume the `brain.wasm` is available at a path known to the MCP server.
- We assume the LLM client (e.g., Gemini CLI, Claude Desktop) supports MCP tool calling.
