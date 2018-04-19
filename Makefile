CONTAINER := agy/linuxkit


# container:
# 	docker build --tag $(CONTAINER) .

.PHONY: all
all: snapshot-import snapshot-import-poll snapshot-register snapshot-sfn

.PHONY: snapshot-import
snapshot-import:
	go build -o bin/$@ cmd/$@/main.go

.PHONY: snapshot-import-poll
snapshot-import-poll:
	go build -o bin/$@ cmd/$@/main.go

.PHONY: snapshot-register
snapshot-register:
	go build -o bin/$@ cmd/$@/main.go

.PHONY: snapshot-sfn
snapshot-sfn:
	go build -o bin/$@ cmd/$@/main.go
