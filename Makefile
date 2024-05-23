.PHONY: default
default: build

IMAGE_TAG=$(shell git show --format='%h' --no-patch)
# Use CURDIR instead of $PWD. See https://www.gnu.org/software/make/manual/html_node/Quick-Reference.html about CURDIR
IMAGE_NAME=$(shell basename $(CURDIR)):$(IMAGE_TAG)

.PHONY: build
build:
	docker build -t $(IMAGE_NAME) .

.PHONY: run
run: build
	docker run --rm -p 8080:8080 $(IMAGE_NAME)

DOCKER_REPO_PREFIX=akima

.PHONY: push
push: build
	docker tag $(IMAGE_NAME) $(DOCKER_REPO_PREFIX)/$(IMAGE_NAME)
	docker push $(DOCKER_REPO_PREFIX)/$(IMAGE_NAME)
