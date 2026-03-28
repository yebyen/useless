# Vind-Box Universal Remote Control

BRAIN_DIR = components/brain
OPERATOR_DIR = components/operator
API_DIR = components/api/api
SPECS_DIR = specs/core

.PHONY: all
all: generate build test

.PHONY: generate
generate: generate-rust generate-go

.PHONY: generate-rust
generate-rust:
	@echo "Generating Rust types from OpenAPI..."
	# Placeholder for openapi-generator-cli
	@echo "Generating Rust bindings from WIT..."
	# Placeholder for wit-bindgen

.PHONY: generate-go
generate-go:
	@echo "Generating Go types from OpenAPI..."
	# Placeholder for openapi-generator-cli
	@echo "Generating Go CRDs and controllers..."
	cd $(OPERATOR_DIR) && make generate

.PHONY: build
build: build-brain build-operator build-api build-cli

.PHONY: build-brain
build-brain:
	cd $(BRAIN_DIR) && cargo build --target wasm32-wasip1 --release

.PHONY: build-operator
build-operator:
	cd $(OPERATOR_DIR) && go build -o bin/manager main.go

.PHONY: build-api
build-api:
	cd $(API_DIR) && spin build

.PHONY: build-cli
build-cli:
	cd components/cli && go build -o useless main.go

.PHONY: test
test: test-brain test-operator

.PHONY: test-brain
test-brain:
	cd $(BRAIN_DIR) && cargo test

.PHONY: test-operator
test-operator:
	cd $(OPERATOR_DIR) && go test ./...

.PHONY: assemble
assemble:
	@echo "Assembling deployment artifacts into deploy/..."
	mkdir -p deploy/apps/operator
	cp $(OPERATOR_DIR)/config/crd/bases/* deploy/apps/operator/crd.yaml
	# Ensure the Kustomization is up to date
	@echo "apiVersion: kustomize.config.k8s.io/v1beta1\nkind: Kustomization\nresources:\n  - crd.yaml\n  - deployment.yaml" > deploy/apps/operator/kustomization.yaml

.PHONY: flux-bootstrap
flux-bootstrap:
	@echo "Bootstrapping Flux into the local cluster..."
	kubectl apply -f deploy/clusters/my-cluster/vind-box.yaml

.PHONY: run-mcp
run-mcp:
	cd components/mcp && uv run python server.py

.PHONY: e2e-test
e2e-test:
	@echo "Running E2E tests in Kind cluster..."
	# Placeholder for kind-driven test script
