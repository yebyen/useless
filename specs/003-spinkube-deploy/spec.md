# Feature Specification: SpinKube Deployment

**Feature Branch**: `003-spinkube-deploy`  
**Created**: 2026-03-28  
**Status**: Draft  

## Implementation Findings

### Successful Cluster Setup (Validated 2026-03-28)
- **Tooling**: `k3d` is the preferred local cluster provider.
- **Node Image**: `ghcr.io/spinframework/containerd-shim-spin/k3d:v0.23.0` (Bundles the `containerd-shim-spin-v2`).
- **Operator**: `spin-operator` v0.6.1 via `oci://ghcr.io/spinframework/charts/spin-operator`.
- **Runtime Class**: `wasmtime-spin-v2` (Handler: `spin`).
- **Lesson Learned**: `kwasm-operator` is deprecated/incompatible with minimal k3s node images used by k3d; use the official adapter image or the new `runtime-class-manager`.

## User Scenarios & Testing

### User Story 1 - Deploying Spin App to Kubernetes (Priority: P1)
As a cluster operator, I want the Spin API to be automatically deployed into the cluster by Flux using a SpinApp custom resource, so that it runs alongside the real Kubernetes CRDs.

**Independent Test**: I can see the SpinApp running in the cluster by running `kubectl get spinapps -n default` and it should show `Ready`.

### User Story 2 - Local Network Access (Priority: P1)
As an Android user, I want to access the Spin API from my physical Android device connected to the local Wi-Fi, rather than relying on emulator port forwarding.

**Independent Test**: Hitting `http://<local-ip>:<node-port>/status` from a phone's browser should return the Useless Machine status.

## Requirements

### Functional Requirements
- **FR-001**: System MUST package the Spin application as an OCI artifact and publish it to a registry (e.g. `ttl.sh` for testing, or `ghcr.io`).
- **FR-002**: System MUST include a `SpinApp` Kubernetes manifest in the `deploy/apps/api` directory.
- **FR-003**: System MUST expose the `SpinApp` via a Kubernetes `Service` (NodePort or LoadBalancer) to the local network.

### Technical Requirements
- **TR-001**: OCI artifact pushing should be orchestrated via the `Makefile`.
- **TR-002**: Flux Kustomization must recognize and apply the new manifests in `deploy/apps/api`.

## Success Criteria

- **SC-001**: Flux correctly synchronizes the `deploy/apps/api` folder.
- **SC-002**: Real Android phone can push the button.
