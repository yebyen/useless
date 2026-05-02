# Useless Machine Architecture

The "Useless Machine" is a distributed, "Zero-Split-Brain" cloud-native system built to explore modern cloud patterns including WebAssembly (WASM), Kubernetes Operators, and the Spin framework.

A physical "Useless Machine" is a box with a switch; when you flip the switch on, a mechanical finger immediately pops out and flips it back off. This project translates that concept into a resilient, distributed software architecture.

## System Components

### 1. The WASM Brain (`components/brain`)
*   **What it is:** The pure, isolated logic core written in Rust.
*   **How it works:** It uses the Extism PDK. It doesn't know about Kubernetes, HTTP, or Spin. It takes in a JSON payload representing the current state of the machine (e.g., `daily_count`, `last_pushed`) and the current time. It applies pure functional logic:
    *   *Push:* Increments the daily count and updates the timestamp.
    *   *Status:* Checks if you've gone more than 24 hours without pushing the button. If so, it flags `is_nagging: true` (because you're neglecting your duties to the useless machine).
*   **Why it's cool:** By compiling this to WASM, the logic is universally portable and executed in a secure, sandboxed environment.

### 2. The Kubernetes Operator (`components/operator`)
*   **What it is:** A custom Kubernetes controller written in Go using Kubebuilder.
*   **How it works:** It introduces a Custom Resource Definition (CRD) called `UselessMachine`. The operator continuously watches the Kubernetes API for changes to this resource. 
    *   When it sees a `UselessMachine` where the `spec.action` has been patched to `"push"`, it spins up the WASM Brain using the `extism/go-sdk`.
    *   It feeds the current status to the Brain, receives the updated state, updates the `status` block in Kubernetes, and clears the `action` request.
*   **Why it's cool:** It turns the "finger" of the useless machine into an autonomous reconcile loop. Kubernetes itself enforces the desired state (turning the switch back off).

### 3. The Spin API Server (`components/api/api`)
*   **What it is:** A WebAssembly microservice built with the Spin framework.
*   **How it works:** It acts as the HTTP gateway to the machine, exposing `/status` and `/push` endpoints. When a user hits `/push`, the API server makes an authenticated HTTP PATCH request to the Kubernetes API server, setting the `UselessMachine`'s `action` to `"push"`. 
*   **Deployment:** When running inside Kubernetes, it uses the `Spintainer` (Spin container executor) with Spin v4.0.0. This allows it to run as a standard Kubernetes container, mount the Kubernetes ServiceAccount token, and trust the cluster's CA Certificate (`cluster-ca`), allowing the Spin HTTP client to securely talk to the cluster's internal API. When deployed to Fermyon Cloud, it requires a publicly accessible Kubernetes API URL and token.

### 4. The CLI (`components/cli`)
*   **What it is:** A Go-based command-line interface for the user.
*   **How it works:** It provides simple commands (`useless push`, `useless status`). It supports swappable backends:
    *   `--backend k8s`: Talks directly to Kubernetes (like `kubectl`).
    *   `--backend api`: Talks to the Spin API Gateway.
*   **Why it's cool:** It hides all the distributed complexity from the end user, giving them a simple remote control.

## The Request Lifecycle

When a user executes `./useless push --backend api`, the following sequence occurs:

1.  **CLI:** Sends an HTTP POST to the **Spin API** `/push` endpoint.
2.  **Spin API:** Uses its injected K8s ServiceAccount token (or configured Fermyon Cloud variables) to issue an HTTP PATCH to the K8s API server, setting `spec.action: "push"` on the `global` UselessMachine CRD.
3.  **K8s API:** Accepts the patch and notifies watchers.
4.  **Operator:** Detects the change in the CRD. It loads `brain.wasm` into the Extism runtime, passing the current count and timestamp.
5.  **WASM Brain:** Computes the new count (e.g., `1 -> 2`) and returns the state.
6.  **Operator:** Updates the CRD's `status` with the new count and resets `spec.action` to empty (flipping the switch back off).
7.  **CLI:** The user runs `./useless status`, which asks the API, which queries K8s, and shows the newly incremented count.

## Concessions & Upstream Contributions (v0.0.2 to v0.1.0)

During the implementation of version 0.0.2, we discovered a gap in the SpinKube ecosystem regarding Kubernetes API authentication.

*   **The Spintainer Concession:** To allow the Spin API (compiled to WASM) to securely communicate with the Kubernetes API, it requires the cluster's internal CA Certificate and a ServiceAccount token. Injecting these via standard Kubernetes volumes into the pure `containerd-shim-spin` (Runtime Class Manager) proved difficult. As a concession, we utilize the **Spintainer (`spintainer-v4`) SpinAppExecutor**, which runs the `ghcr.io/spinframework/spin:v4.0.0` image as a standard container. This allows us to use native Kubernetes volume mounts for the CA cert and token.
*   **Upstream Goal:** We aim to work with the Spin and Kubernetes communities to natively support certificate and OIDC/token injection in the shim, allowing WASM apps to interact with the K8s API seamlessly without needing a containerized executor wrapper.
*   **WebSockets (The SpinKube Advantage):** The v0.1.0 roadmap includes replacing client polling with WebSockets or Server-Sent Events (SSE). While serverless platforms (like Spin Cloud) often discourage or aggressively terminate long-lived connections to scale to zero, deploying our Spin API in **SpinKube** means we control the infrastructure. We can easily support persistent real-time connections for our clients, demonstrating the flexibility of self-hosted WASM workloads.
