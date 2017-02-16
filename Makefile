SHELL = /bin/sh
IMAGE_REPO = kassisol

default: binary

dep:
	if [ ! -d "/tmp/hbm" ]; then mkdir /tmp/hbm; fi

scripts:
	if [ ! -d "tmp" ]; then mkdir tmp; fi

	if [ ! -d tmp/makefile-scripts ]; then \
		git clone https://github.com/juliengk/makefile-scripts.git tmp/makefile-scripts; \
	fi

clean:
	git reset --hard origin/${GIT_BRANCH}

binary: dep scripts Dockerfile
	$(eval TAG_VERSION := $(shell tmp/makefile-scripts/git-version.sh))
	$(eval GIT_BRANCH := $(shell tmp/makefile-scripts/git-branch.sh))

	@tmp/makefile-scripts/docker-check-image-exist.sh ${IMAGE_REPO} "hbm-binary" ${TAG_VERSION}

	docker build --rm --no-cache --build-arg version=${TAG_VERSION} -t ${IMAGE_REPO}/hbm-binary:${TAG_VERSION} .
	docker run -t --rm -v /tmp/hbm:/tmp/hbm ${IMAGE_REPO}/hbm-binary:${TAG_VERSION}

centos7: binary hack/packages/rpm/centos7/Dockerfile
	$(eval TAG_VERSION := $(shell tmp/makefile-scripts/git-version.sh))
	$(eval GIT_BRANCH := $(shell tmp/makefile-scripts/git-branch.sh))

	TMP_PATH = hack/packages/rpm/centos7/tmp

	@tmp/makefile-scripts/docker-check-image-exist.sh ${IMAGE_REPO} "hbm-centos7" ${TAG_VERSION}

	if [ ! -d ${TMP_PATH} ]; then mkdir ${TMP_PATH}; fi
	cp contrib/init/systemd/hbm.service ${TMP_PATH}/
	cp hack/packages/rpm/build-rpm.sh ${TMP_PATH}/

	cp /tmp/hbm/hbm-*-linux-amd64.tar.gz ${TMP_PATH}/

	docker build --rm --no-cache -t ${IMAGE_REPO}/hbm-centos7:${TAG_VERSION} hack/packages/rpm/centos7
	docker run -t --rm -v /tmp/hbm:/tmp/hbm ${IMAGE_REPO}/hbm-centos7:${TAG_VERSION}

.PHONY: default clean binary centos7
