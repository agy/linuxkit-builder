CONTAINER := agy/linuxkit
BINARIES := snapshot-import snapshot-import-poll snapshot-register snapshot-sfn


# container:
# 	docker build --tag $(CONTAINER) .

.PHONY: all
all: vet fmt test $(BINARIES)

.PHONY: test
test:
	go test ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

%:
	go build -o bin/$@ $(wildcard cmd/$@/*.go)
