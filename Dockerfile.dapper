FROM debian:stable-slim

RUN apt-get update \
	&& apt-get install -y --no-install-recommends \
		build-essential \
		ca-certificates \
		curl \
		gccgo \
		git

########## Dapper Configuration #####################
ENV DAPPER_DOCKER_SOCKET true
ENV DAPPER_SOURCE /go/src/github.com/kassisol/hbm
ENV DAPPER_OUTPUT ./bin ./dist
ENV SHELL /bin/bash

WORKDIR ${DAPPER_SOURCE}

########## General Configuration #####################
ARG DAPPER_HOST_ARCH=amd64
ARG HOST_ARCH=${DAPPER_HOST_ARCH}
ARG ARCH=${HOST_ARCH}

ARG APP_REPO=kassisol

ARG docker_version=18.03.0

ARG DOCKER_URL_amd64=https://download.docker.com/linux/static/stable/x86_64/docker-${docker_version}-ce.tgz

# Set up environment and export all ARGS as ENV
ENV ARCH=${ARCH} \
	HOST_ARCH=${HOST_ARCH}

ENV DOCKER_URL=${DOCKER_URL_amd64} \
	DAPPER_HOST_ARCH=${DAPPER_HOST_ARCH} \
	GOPATH=/go \
	GOARCH=$ARCH \
	GO_VERSION=1.8.3

ENV PATH=/go/bin:/usr/local/go/bin:$PATH

# Install Go
RUN curl -sfL https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz | tar -xzC /usr/local

# Install Host Docker
RUN curl -sfL ${DOCKER_URL} | tar -xzC /usr/local/src \
	&& cp /usr/local/src/docker/docker /usr/bin/docker \
	&& chmod +x /usr/bin/docker

# Install dapper
RUN curl -sL "https://releases.rancher.com/dapper/latest/dapper-$(uname -s)-$(uname -m | sed 's/arm.*/arm/')" > /usr/bin/dapper \
	&& chmod +x /usr/bin/dapper

ENTRYPOINT ["./scripts/entry"]
CMD ["ci"]
