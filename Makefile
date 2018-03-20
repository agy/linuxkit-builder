CONTAINER := agy/linuxkit


container:
	docker build --tag $(CONTAINER) .
