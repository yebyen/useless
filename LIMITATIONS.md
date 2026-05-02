# Known Limitations

This document tracks known architectural limitations and blockers.

## Kubernetes API Access from Spin Components (TLS Verification Failure)

In our experiment (preserved in `5851e04`), we attempted to call the Kubernetes API directly from a Spin app deployed via SpinKube.

### The Problem
Spin enforces secure TLS for all outbound HTTPS calls. When running inside a Kubernetes cluster, the API server (`kubernetes.default.svc`) uses a certificate signed by the internal cluster CA. 

Currently, there is no standardized way to inject this CA certificate into the Spin/Wasm runtime environment such that the HTTP client can verify the API server's identity. This prevents Spin components from interacting directly with the Kubernetes API within the cluster.

### Impact
This blocks deploying the full "Zero-Split-Brain" architecture that relies on Spin components to interact with the cluster state directly. It is not an urgent blocker for local development but will prevent production-ready deployments of these specific components.

### Possible Workarounds
1.  **TLS-Terminated Proxy:** Use a proxy for the Kube API (e.g., a sidecar or a dedicated gateway) that is appropriately TLS-terminated with a certificate from a known, trusted CA.
2.  **CA Injection:** Wait for or implement support in SpinKube/containerd-shim-spin for CA certificate injection into the runtime.
3.  **Bridge Component:** Use a traditional container-based component (e.g., a Go operator or Python service) that has access to the cluster CA to act as a bridge for Spin components.

---
*Last updated: March 29, 2026*
