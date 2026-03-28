# Tasks: Infrastructure & GitOps (Vind-Box)

**Input**: Design documents from `.specify/specs/002-infrastructure/`
**Prerequisites**: plan.md, spec.md

## Phase 1: Flux Setup

**Purpose**: Establish the GitOps control plane.

- [x] T001 [P] Create `deploy/clusters/my-cluster/flux-system` (Placeholder for `flux bootstrap`)
- [x] T002 Create `deploy/clusters/my-cluster/vind-box.yaml` with `GitRepository` and `Kustomization`
- [ ] T003 Verify `vcluster` connectivity and Flux installation

---

## Phase 2: Spin API (Real K8s Integration)

**Purpose**: Upgrade the Cloud Gateway to talk to the real source of truth.

- [x] T004 Implement helper to read K8s Service Account token in `components/api/api/src/lib.rs`
- [x] T005 Implement K8s API request logic using `spin_sdk::http::send`
- [x] T006 Update `/status` endpoint to fetch real CRD status
- [x] T007 Update `/push` endpoint to patch the CRD spec
- [x] T008 Update `spin.toml` with `allowed_outbound_hosts = ["https://kubernetes.default.svc"]`

---

## Phase 3: Validation & Assembly

**Purpose**: Ensure the entire pipeline works end-to-end.

- [x] T009 Update root `Makefile` with `make assemble` to collect manifests
- [ ] T010 Manually create the `global` UselessMachine resource for testing
- [ ] T011 Run end-to-end smoke test (API -> K8s -> Operator -> WASM)

---

## Dependencies

- **Phase 1** can run independently.
- **Phase 2** depends on having a functioning K8s API (vcluster).
- **Phase 3** depends on completion of Phase 1 and 2.
