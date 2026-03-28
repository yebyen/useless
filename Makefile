# Vind-Box Universal Remote Control

BRAIN_DIR = components/brain
OPERATOR_DIR = components/operator
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
	cd $(OPERATOR_DIR) && controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases

.PHONY: build
build: build-brain build-operator

.PHONY: build-brain
build-brain:
	cd $(BRAIN_DIR) && cargo build --target wasm32-wasi --release

.PHONY: build-operator
build-operator:
	cd $(OPERATOR_DIR) && go build -o bin/manager main.go

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
	mkdir -p deploy/manifests
	# Placeholder for kustomize build
	cp -r $(OPERATOR_DIR)/config/crd/bases/* deploy/manifests/

.PHONY: e2e-test
e2e-test:
	@echo "Running E2E tests in Kind cluster..."
	# Placeholder for kind-driven test script
