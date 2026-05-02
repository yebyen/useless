# Feature Specification: Infrastructure & GitOps (Vind-Box)

**Feature Branch**: `002-infrastructure`  
**Created**: 2026-03-28  
**Status**: Draft  
**Input**: User description: "Setup Flux GitOps bootstrap and integrate real K8s API into the Spin gateway."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Self-Healing Deployment (Priority: P1)

As a maintainer, I want the system to be managed by Flux so that any changes to the `deploy/` directory are automatically applied to the cluster.

**Why this priority**: This is the foundation of our GitOps approach. It ensures the environment is reproducible and "self-heals" if resources are deleted.

**Independent Test**: Deleting the `vind-box-operator` deployment manually in the cluster should result in Flux recreating it within minutes.

**Acceptance Scenarios**:

1. **Given** a functioning K8s cluster, **When** I apply the Flux bootstrap manifests, **Then** a `GitRepository` source and `Kustomization` are created.
2. **Given** Flux is running, **When** I push a change to `deploy/apps/operator`, **Then** the operator is updated in the cluster.

---

### User Story 2 - Real K8s Integration for Spin API (Priority: P1)

As a user, I want the Spin API to show the *actual* status of the `UselessMachine` CRD from Kubernetes, not a mock.

**Why this priority**: This proves the Cloud Gateway can talk to the "Database" (K8s API) securely using standard service account tokens.

**Independent Test**: Pushing the button via the API should update the `UselessMachine` status in Kubernetes, visible via `kubectl get uselessmachine global -o yaml`.

---

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST bootstrap Flux into the `vind` cluster.
- **FR-002**: Flux MUST monitor the `main` branch of this repository.
- **FR-003**: Spin API MUST read the `UselessMachine` status from the K8s API.
- **FR-004**: Spin API MUST patch the `UselessMachine` spec to trigger a "push" action.
- **FR-005**: Spin API SHOULD support a long-polling or webhook mechanism to notify clients of state changes instantly.

### Technical Requirements

- **TR-004**: Flux manifests MUST be stored in `deploy/clusters/my-cluster`.
- **TR-005**: Spin API MUST use the `kubernetes.default.svc` endpoint.
- **TR-006**: Spin API MUST authorize the Kubernetes host in `spin.toml`.

## Implementation Notes (v0.0.2 -> v0.1.0)

- **CA Certificate Injection (Resolved via Concession):** The `containerd-shim-spin` currently lacks an ergonomic way to mount the cluster's CA certificate necessary for verifying `https://kubernetes.default.svc`. To unblock v0.0.2, we use the `spintainer-v4` (`SpinAppExecutor`), which deploys the app as a standard container (`ghcr.io/spinframework/spin:v4.0.0`). This allows standard volume mounting of the CA cert and service account token. *Future Goal: Contribute upstream to support this natively in the shim.*
- **Real-Time Client Updates (v0.1.0):** To address **FR-005**, the Android client currently relies on 5-second HTTP polling. Version 0.1.0 will implement WebSockets or SSE in the Spin API. Because we run this in SpinKube (rather than a strictly scale-to-zero serverless cloud), we can fully support long-lived persistent connections for real-time reactivity.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-004**: `flux get kustomizations` shows "Ready" for the `vind-box` app.
- **SC-005**: `curl http://localhost:3000/status` returns the real JSON from the CRD.
- **SC-006**: `curl http://localhost:3000/push` successfully updates the CRD action.

## Assumptions

- We assume the `vcluster` (Vind) environment provides a standard service account token at `/var/run/secrets/kubernetes.io/serviceaccount/`.
- We assume the Spin app will be deployed *inside* the cluster (SpinKube) to access the internal K8s service.
