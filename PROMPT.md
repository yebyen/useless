🚀 PROMPT: The Unified Mecris Framework (Codename: "Vind-Box")

Objective: Build a "Useless Machine" MVP that serves as the architectural blueprint for a next-generation, zero-split-brain agent system. The core business
logic must be encapsulated in a WASM component that runs identically across four distinct host environments.

1. The Core Components
 1. The Brain (WASM/Rust): The single source of truth for business logic.
     - Tracks a last_pushed timestamp and daily_count.
     - Determines nag_state (if now - last_pushed > 24h).
     - Logic: push_button() increments count and resets the timer.
 2. The Database (K8s CRD): A UselessMachine Custom Resource in a Vind (vcluster-in-docker) cluster.
     - State is persisted in the cluster's etcd.
 3. The Reconciler (Go Operator): A controller (using kubebuilder) that loads the WASM Brain to reconcile the CRD state.
 4. The Cloud Gateway (Spin App): A Rust-based Spin service deployed via SpinKube.
     - Exposes a private API (/status, /push).
     - Uses a Kubernetes Service Account to read/write the UselessMachine CRD.
     - Loads the same WASM Brain to ensure API responses match the Operator's logic.
 5. The Tools (MCP/Python): An MCP server for LLMs to interact with the system via the same WASM Brain.
 6. The Client (Android/Kotlin): Hits the Spin API for primary interaction, with a skeleton for local WASM failover.

2. Constraints & Scope
 - Single User: No multi-tenancy. One CRD represents the global state.
 - No Auth: Private network only. Authentication (like Pocket-ID) is out of scope for this MVP.
 - Framework-First: Use Extism for Host Functions (logging, time, HTTP). Ensure the guest-host interface is clean enough to pivot to the WASM Component
   Model (WIT) if needed.
 - Decoupled Roadmap: Components can be implemented iteratively. It is OK if the Android app starts with a mock before the WASM is fully integrated.

3. Initial Deliverables (Phase 1)
 1. Infra: A Makefile to bootstrap the Vind cluster.
 2. Guest: The Rust-based WASM Brain.
 3. Operator: The Go-based CRD controller loading the WASM.
 4. API: The Spin-based web service talking to the K8s API.

Mandate: The WASM should have no knowledge of its host. It consumes JSON and produces JSON. All I/O (K8s API, Network, Time) is handled by the Host via
Extism host functions or parameter passing.
