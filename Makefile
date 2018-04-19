CONTAINER := agy/linuxkit
BINARIES := snapshot-import snapshot-import-poll snapshot-register snapshot-sfn


# container:
# 	docker build --tag $(CONTAINER) .

.PHONY: all
all: $(BINARIES)

%:
	go build -o bin/$@ $(wildcard cmd/$@/*.go)
