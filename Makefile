.PHONY: default
default: build

# Use CURDIR instead of $PWD. See https://www.gnu.org/software/make/manual/html_node/Quick-Reference.html about CURDIR
IMAGE_NAME=$(shell basename $(CURDIR))

.PHONY: build
build:
	docker build -t $(IMAGE_NAME) .

.PHONY: run
run: build
	docker run --rm -p 8080:8080 $(IMAGE_NAME)
