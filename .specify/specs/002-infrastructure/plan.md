# Implementation Plan: Infrastructure & GitOps

**Branch**: `002-infrastructure` | **Date**: 2026-03-28 | **Spec**: [spec.md](./spec.md)

## Summary
Bootstrap Flux for GitOps management and upgrade the Spin API from mocks to real Kubernetes interaction.

## Technical Context
- **Flux CLI**: Latest stable
- **Kubernetes**: Vcluster (Vind)
- **Authentication**: K8s Service Account Token
- **API Endpoint**: `https://kubernetes.default.svc`

## Project Structure

### Documentation
```text
.specify/specs/002-infrastructure/
├── spec.md
├── plan.md
└── tasks.md
```

### Manifests
```text
deploy/
├── clusters/
│   └── my-cluster/
│       ├── flux-system/      # Flux components
│       └── vind-box.yaml     # GitRepository + Kustomization
└── apps/
    └── operator/             # Already exists
```

## Implementation Strategy

### Phase 1: Flux Bootstrap
1. Generate Flux bootstrap manifests using `flux bootstrap github` or manual equivalent for local vcluster.
2. Define `GitRepository` source pointing to this repo (public for now).
3. Define `Kustomization` pointing to `deploy/apps/operator`.

### Phase 2: Real K8s API in Spin
1. Implement token reading from `/var/run/secrets/kubernetes.io/serviceaccount/token`.
2. Implement CA cert loading for secure HTTPS calls to the API server.
3. Replace mock logic in `components/api/api/src/lib.rs` with `spin_sdk::http::send` calls.
4. Update `spin.toml` to authorize the API server host.

### Phase 3: Validation
1. Apply CRD manually.
2. Create "global" `UselessMachine` resource.
3. Run Spin locally (authorized via `kubectl proxy` or similar) to verify connectivity.

## Success Criteria
- [ ] Flux is reconciling the `deploy/` directory.
- [ ] Spin API can successfully `GET` and `PATCH` the `UselessMachine` resource.
