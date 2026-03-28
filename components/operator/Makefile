# Operator Makefile

OPERATOR_GEN = go run sigs.k8s.io/controller-tools/cmd/controller-gen

.PHONY: generate
generate:
	$(OPERATOR_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."
	$(OPERATOR_GEN) crd paths="./..." output:crd:dir=config/crd/bases

.PHONY: build
build: generate
	go build -o bin/manager main.go

.PHONY: test
test:
	go test ./...
