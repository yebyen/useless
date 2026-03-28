# Vind-Box: The Next-Generation Useless Machine

Vind-Box is an architectural blueprint for building **Zero-Split-Brain** agent systems. It uses WebAssembly (WASM) to ensure that core business logic—the "Brain"—is shared identically across every part of the stack, from Kubernetes Operators to Android apps, eliminating logic duplication and synchronization bugs.

## 🚀 The Ecosystem

1.  **The Brain (Rust/WASM):** The single source of truth for logic (`push_button`, `is_nagging`).
2.  **The Operator (Go):** Kubernetes controller that loads the WASM Brain to reconcile state.
3.  **The API (Spin/Rust):** Cloud gateway providing web and mobile access to the machine.
4.  **The MCP Server (Python):** Exposes the machine logic as tools for AI agents.
5.  **The CLI (Go):** Robust terminal interface for power users.
6.  **The Client (Android/Kotlin):** Modern mobile UI with real-time status.

## 🛠️ Usage

### Build Everything
```bash
make build
```

### Run the Operator
```bash
cd components/operator
go run main.go --wasm-path=../brain/target/wasm32-wasip1/release/brain.wasm
```

### Run the API (Spin)
```bash
cd components/api/api
spin up --variable k8s_api_url=http://localhost:8001
```

### Use the CLI
```bash
./components/cli/useless status
./components/cli/useless push
```

### AI Integration (MCP)
Add the following to your Gemini/Claude settings:
```json
"vind-box": {
  "command": "uv",
  "args": ["run", "--project", "./components/mcp", "python3", "server.py", "--stdio"]
}
```

## 📄 Documentation
Detailed specs and architectural records are in the `.specify/` directory.

---
**Version:** 0.0.1 | **License:** MIT | *Built with Gemini CLI*
