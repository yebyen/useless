# Known Limitations

This document tracks known architectural limitations and blockers.

## Kubernetes API Access from Spin Components (Resolved via Spintainer Concession)

In our initial experiments, we encountered a TLS verification failure when attempting to call the Kubernetes API (`kubernetes.default.svc`) from a Spin app. This was because the Spin/Wasm runtime lacked a way to verify the cluster's internal CA.

### The Solution (v0.0.2)
We have resolved this for the v0.0.2 release by utilizing the **Spintainer (`SpinAppExecutor`)** workaround. By deploying the Spin API as a standard Kubernetes container (using the `ghcr.io/spinframework/spin:v4.0.0` image), we are able to:
1.  **Mount the CA Certificate:** Inject the cluster's root CA into the container's standard certificate store.
2.  **ServiceAccount Token Injection:** Securely mount the Kubernetes ServiceAccount token for authentication.

### Long-term Goal
While the Spintainer approach unblocks us, it is a concession. We aim to contribute upstream to the **SpinKube/containerd-shim-spin** project to enable native CA certificate injection directly into the Wasm runtime, eliminating the need for a containerized wrapper.

---
*Last updated: May 2, 2026*
