NAME := snapshot-register

GOOS := linux
GOARCH := amd64


.PHONY: build vet fmt
build: vet fmt
	GOOS=$(GOOS) GOARCH=$(GOARCH) go $@ .

vet:
	go $@ ./...

fmt:
	go $@ ./...

package: build
	zip -j9 $(NAME).zip $(NAME)
	rm -f $(NAME)

clean:
	rm -f $(NAME).zip
	rm -f $(NAME)
