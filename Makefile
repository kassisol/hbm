SHELL = /bin/sh
IMAGE_REPO = kassisol
IMAGE_NAME = hbm-build

TAG_VERSION := $(shell hack/git-version.sh)
GIT_BRANCH := $(shell hack/git-branch.sh)

default: build

dep:
	mkdir /tmp/hbm

all: dep
	docker run -t --rm -v /tmp/hbm:/tmp/hbm ${IMAGE_REPO}/${IMAGE_NAME}:${TAG_VERSION} all

binary: dep
	docker run -t --rm -v /tmp/hbm:/tmp/hbm ${IMAGE_REPO}/${IMAGE_NAME}:${TAG_VERSION} binary


build: Dockerfile
	@hack/docker-check-image-exist.sh ${IMAGE_REPO} ${IMAGE_NAME} ${TAG_VERSION}

	docker build --rm --no-cache --build-arg version=${TAG_VERSION} -t ${IMAGE_REPO}/${IMAGE_NAME}:${TAG_VERSION} .

clean: scripts
	git reset --hard origin/${GIT_BRANCH}

rpm: dep
	docker run -t --rm -v /tmp/hbm:/tmp/hbm ${IMAGE_REPO}/${IMAGE_NAME}:${TAG_VERSION} rpm

.PHONY: default all binary build clean rpm
